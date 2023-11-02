package backendgcp

import (
	"context"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"

)

func GetConnectionMongo(MongoString, dbname string) *mongo.Database {
	MongoInfo := atdb.DBInfo{
		DBString: os.Getenv(MongoString),
		DBName:   dbname,
	}
	conn := atdb.MongoConnect(MongoInfo)
	return conn
}

func GetAllData(MongoConnect *mongo.Database, colname string) []GeoJson {
	data := atdb.GetAllDoc[[]GeoJson](MongoConnect, colname)
	return data
}

func InsertDataGeojson(MongoConn *mongo.Database, colname string, coordinate []float64, name, volume, tipe string) (InsertedID interface{}) {
	req := new(LonLatProperties)
	req.Type = tipe
	req.Coordinates = coordinate
	req.Name = name
	req.Volume = volume

	ins := atdb.InsertOneDoc(MongoConn, colname, req)
	return ins
}

func UpdateDataGeojson(MongoConn *mongo.Database, colname, name, newVolume, newTipe string) error {
    // Filter berdasarkan nama
    filter := bson.M{"name": name}

    // Update data yang akan diubah
    update := bson.M{
        "$set": bson.M{
            "volume": newVolume,
            "tipe":   newTipe,
        },
    }

    // Mencoba untuk mengupdate dokumen
    _, err := MongoConn.Collection(colname).UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return err
    }

    return nil
}

func DeleteDataGeojson(MongoConn *mongo.Database, colname string, name string) (*mongo.DeleteResult, error) {
    filter := bson.M{"name": name}
    del, err := MongoConn.Collection(colname).DeleteOne(context.TODO(), filter)
    if err != nil {
        return nil, err
    }
    return del, nil
}
