package backendgcp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GCHandlerFunc(Mongostring, dbname, colname string) string {
	koneksyen := GetConnectionMongo(Mongostring, dbname)
	datageo := GetAllData(koneksyen, colname)

	jsoncihuy, _ := json.Marshal(datageo)

	return string(jsoncihuy)
}

func GCFPostCoordinate(Mongostring, dbname, colname string, r *http.Request) string {
	req := new(Credents)
	conn := GetConnectionMongo(Mongostring, dbname)
	resp := new(LonLatProperties)
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		req.Status = strconv.Itoa(http.StatusNotFound)
		req.Message = "error parsing application/json: " + err.Error()
	} else {
		req.Status = strconv.Itoa(http.StatusOK)
		Ins := InsertDataGeojson(conn, colname,
			resp.Coordinates,
			resp.Name,
			resp.Volume,
			resp.Type)
		req.Message = fmt.Sprintf("%v:%v", "Berhasil Input data", Ins)
	}
	return ReturnStringStruct(req)
}

func ReturnStringStruct(Data any) string {
	jsonee, _ := json.Marshal(Data)
	return string(jsonee)
}

func GCFUpdateNameGeojson(Mongostring, dbname, colname string, r *http.Request) string {
	req := new(Credents)
	resp := new(LonLatProperties)
	conn := GetConnectionMongo(Mongostring, dbname)
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		req.Status = strconv.Itoa(http.StatusNotFound)
		req.Message = "error parsing application/json: " + err.Error()
	} else {
		req.Status = strconv.Itoa(http.StatusOK)
		Ins := UpdateDataGeojson(conn, colname,
			resp.Name,
			resp.Volume,
			resp.Type)
		req.Message = fmt.Sprintf("%v:%v", "Berhasil Input data", Ins)
	}
	return ReturnStringStruct(req)
}

func GCFDeleteDataGeojson(Mongostring, dbname, colname string, r *http.Request) string {
    req := new(Credents)
    resp := new(LonLatProperties)
    conn := GetConnectionMongo(Mongostring, dbname)
    err := json.NewDecoder(r.Body).Decode(&resp)
    if err != nil {
        req.Status = strconv.Itoa(http.StatusNotFound)
        req.Message = "error parsing application/json: " + err.Error()
    } else {
        req.Status = strconv.Itoa(http.StatusOK)
        delResult, delErr := DeleteDataGeojson(conn, colname, resp.Name)
        if delErr != nil {
            req.Status = strconv.Itoa(http.StatusInternalServerError)
            req.Message = "error deleting data: " + delErr.Error()
        } else {
            req.Message = fmt.Sprintf("Berhasil menghapus data. Jumlah data terhapus: %v", delResult.DeletedCount)
        }
    }
    return ReturnStringStruct(req)
}

