// Package main mongodb官方驱动的使用demo
package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mgoURI = `mongodb://root123:abcABC123@127.0.0.1:27017/admin`
	// note: 服务器是副本集，但是找不到Primary连接错误时使用connect=direct
	database   = "test1"
	collection = "t1"
)

func main() {
	//
	ops := options.Client().ApplyURI(mgoURI)
	if err := ops.Validate(); err != nil {
		log.Fatalln("解析mongodb连接字符串发生错误: ", err)
	}

	ops.SetAppName("demo")

	client, err := mongo.NewClient(ops)
	if err != nil {
		log.Fatalln("初始化mongodb客户端发生错误: ", err)
	}

	ctx := context.Background()
	// 开始连接数据库
	if err = client.Connect(ctx); err != nil {
		log.Fatalln("连接mongodb服务器发生错误: ", err)
	}

	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)

	if err = client.Ping(timeout, nil); err != nil {
		log.Fatalln("ping", err)
	}

	coll := client.Database(database).Collection(collection)
	var d data

	timeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	err = coll.FindOne(timeout, bson.M{}).Decode(&d)
	if err != nil {
		log.Fatalln("查询一条数据发生错误: ", err)
	}
	defer cancel()
	log.Printf("查询结果: %+v", d)
}

//
type data struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}
