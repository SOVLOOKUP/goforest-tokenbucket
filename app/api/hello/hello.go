package hello

import (
	"gf-app/app/service"
	"gf-app/boot"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

type Config struct {
	ID			string			`json:"id" v:"required#缺少id(令牌桶标识)"`
	Rate		float64			`json:"rate" v:"required|min:0.01#缺少rate(令牌桶每秒放令牌个数)|速率不能为0"`
	Capacity	int64			`json:"capacity" v:"required|min:1#缺少capacity(令牌桶大小)|令牌桶大小最小为1"`
	MaxWait		time.Duration	`json:"max_wait" v:"required#缺少max_wait(取令牌最长等待时间)"`
}

type Res struct {
	Code int			`json:"code"`
	Msg string			`json:"msg"`
	Detail interface{} 	`json:"detail"`
}
func GetToken(r *ghttp.Request) {
	if bucket , ok := boot.BucketMap.Get(r.GetString("id")).(*service.Bucket); ok {
		r.Response.WriteExit(bucket.GetToken(1))
	} else {
		r.Response.WriteExit(&Res{-1,"IDNotFound",boot.BucketMap.Keys()})
	}
}

func CreateBucket(r *ghttp.Request) {
	var config *Config

	if err := r.Parse(&config); err != nil {
		r.Response.WriteExit(&Res{-1,"ParamsErr",err.Error()})
	}

	boot.BucketMap.Set(config.ID,service.NewBucket(config.Rate,config.Capacity,config.MaxWait))

	r.Response.WriteExit(&Res{1,"success","token-bucket单元创建成功"})
}

func RemoveBucket(r *ghttp.Request) {

	r.Response.WriteExit(&Res{1,"removed",boot.BucketMap.Remove(r.GetString("id"))})
}
