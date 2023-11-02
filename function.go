package backendgcp

import (
	"context"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertDataGeojson(MongoConn *mongo.Database, colname string, coordinate []float64, name, volume, tipe string) (InsertedID interface{}) {
	req := new(LonLatProperties)
	req.Type = tipe
	req.Coordinates = coordinate
	req.Name = name
	req.Volume = volume

	ins := atdb.InsertOneDoc(MongoConn, colname, req)
	return ins
}

func UpdateNameGeojson(Mongoenv, dbname string, ctx context.Context, val LonLatProperties) (UpdateID interface{}) {
	conn := GetConnectionMongo(Mongoenv, dbname)
	filter := bson.D{{"volume", val.Volume}}
	update := bson.D{{"$set", bson.D{
		{"name", val.Name},
	}}}
	res, err := conn.Collection("post").UpdateOne(ctx, filter, update)
	if err != nil {
		return "Updatenya Gagal Nich"
	}
	return res
}

func DeleteDataGeojson(Mongoenv, dbname string, ctx context.Context, val LonLatProperties) (DeletedId interface{}) {
	conn := GetConnectionMongo(Mongoenv, dbname)
	filter := bson.D{{"volume", val.Volume}}
	res, err := conn.Collection("post").DeleteOne(ctx, filter)
	if err != nil {
		return "Deletenya Gagal Nich"
	}
	return res
}
