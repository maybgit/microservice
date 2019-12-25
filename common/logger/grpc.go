package logger

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/limitedlee/microservice/common"

	"github.com/limitedlee/microservice/common/config"
	"google.golang.org/grpc"
)

var logGrpcUrl string

func init() {
	//获取项目配置中的数据
	var err error
	logGrpcUrl, err = config.Get("LogGrpc")
	if err != nil {
		log.Printf("get config info fail: %v", err)
	}
}

func writeLog(data string, level int) (r *Reply) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	conn, err := grpc.Dial(logGrpcUrl, grpc.WithInsecure())
	defer conn.Close()
	client := NewLogClient(conn)

	ctx, cf := context.WithTimeout(context.Background(), time.Second*60)
	defer cf()

	sps := strings.Split(data, "]")

	m := LogRequest{}
	m.Logger = sps[0]
	m.Appid = common.PbConfig.Grpc.Appid
	m.Message = strings.Join(sps[1:], "]")

	switch level {
	case 0:
		r, err = client.Info(ctx, &m)
	case 1:
		r, err = client.Warn(ctx, &m)
	case 2:
		r, err = client.Error(ctx, &m)
	case 3:
		r, err = client.Fatal(ctx, &m)
	}
	if err != nil {
		panic(err)
	}
	return
}
