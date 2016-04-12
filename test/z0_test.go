package efs

import (
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/json"
	_ "github.com/eaciit/dbox/dbc/jsons"
	_ "github.com/eaciit/dbox/dbc/mongo"
	"github.com/eaciit/efs"
	"github.com/eaciit/toolkit"
	"os"
	"testing"
)

var wd, _ = os.Getwd()

func prepareconnection() (conn dbox.IConnection, err error) {
	conn, err = dbox.NewConnection("mongo",
		&dbox.ConnectionInfo{"192.168.0.200:27017", "efspttgcc", "", "", toolkit.M{}.Set("timeout", 3)})
	// conn, err = dbox.NewConnection("jsons",
	// 	&dbox.ConnectionInfo{wd, "", "", "", toolkit.M{}.Set("newfile", true)})
	if err != nil {
		return
	}

	err = conn.Connect()
	return
}

func TestInitialSetDatabase(t *testing.T) {
	conn, err := prepareconnection()

	if err != nil {
		t.Errorf("Error connecting to database: %s \n", err.Error())
	}

	err = efs.SetDb(conn)
	if err != nil {
		t.Errorf("Error set database to efs: %s \n", err.Error())
	}
}

func loaddatasample() (arrtkm []efs.StatementElement, err error) {

	conn, err := dbox.NewConnection("json",
		&dbox.ConnectionInfo{toolkit.Sprintf("%v/sample.json", wd), "", "", "", nil})
	if err != nil {
		return
	}

	err = conn.Connect()
	if err != nil {
		return
	}

	c, err := conn.NewQuery().Select().Cursor(nil)
	if err != nil {
		return
	}

	arrtkm = make([]efs.StatementElement, 0, 0)
	err = c.Fetch(&arrtkm, 0, false)

	return
}

func TestCreateStatement(t *testing.T) {
	t.Skip("Skip : Comment this line to do test")
	arrdata, err := loaddatasample()
	if err != nil {
		t.Errorf("Error to load data sample: %s \n", err.Error())
		return
	}

	// toolkit.Println(arrdata)
	ds := efs.NewStatements()
	// ds.ID = toolkit.RandomString(32)
	ds.ID = "bid1EWFRZwL-at1uyFvzJYUjPu3yuh3j"
	ds.Title = "donation"
	ds.Elements = make([]efs.StatementElement, 0, 0)

	ds.Elements = append(ds.Elements, arrdata...)

	err = efs.Save(ds)
	if err != nil {
		t.Errorf("Error to save efs: %s \n", err.Error())
		return
	}
}

func TestRunStatement(t *testing.T) {
	// t.Skip("Skip : Comment this line to do test")
	sid := "bid1EWFRZwL-at1uyFvzJYUjPu3yuh3j"

	ds := new(efs.Statements)
	err := efs.Get(ds, sid)
	if err != nil {
		t.Errorf("Error to get statement by id, found : %s \n", err.Error())
		return
	}

	ins := toolkit.M{}.Set("title", "base-v1")
	sv := ds.Run(ins)
	toolkit.Printf("%#v\n", sv)

}