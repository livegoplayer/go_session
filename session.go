package go_session

import (
	"time"

	myHelper "github.com/livegoplayer/go_helper"
	redisHelper "github.com/livegoplayer/go_redis_helper"
)

func InitRedisSession(host string, port string, password string, db int, pre string) {
	//设置最大超时容忍度
	redisHelper.InitRedisHelper(host, port, password, db, pre, 2*time.Second)
}

func CheckUserSession(key string) int64 {
	return redisHelper.IsNameExits(key)
}

func SetUserSession(key string, expireTime time.Duration) {
	redisHelper.SetCacheData(key, myHelper.StringToBytes("1"), expireTime)
}
