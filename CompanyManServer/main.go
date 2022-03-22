package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"CivetCompanyCenter"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(CompanyManProcessImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("CompanyManProcessImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(CivetCompanyCenter.CompanyManProcess)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".CompanyManProcessObj")
	
	// Run application
	tars.Run()
}
