package main

import (
	"CivetCompanyCenter"
	"Civets/CivetTarsDataBase"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"os"
)

// CompanyManProcessImp servant implementation
type CompanyManProcessImp struct {
}

func (imp *CompanyManProcessImp) CreateCompany(tarsCtx context.Context, newCompany *CivetCompanyCenter.Company, result *int32) (ret int32, err error) {
	res := ins.CreateRow(newCompany)
	if res == true {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("CreateCompany Fail")
	}
}

func (imp *CompanyManProcessImp) DeleteCompany(tarsCtx context.Context, CompanyID int32, result *int32) (ret int32, err error) {
	if ins.DelRowByCondition(&CivetCompanyCenter.Company{}, "id", CompanyID) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *CompanyManProcessImp) UpdateCompany(tarsCtx context.Context, Company *CivetCompanyCenter.Company, keys string, value string, result *int32) (ret int32, err error) {
	if ins.EditRowByConditionbyModel(Company, keys, value) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

func (imp *CompanyManProcessImp) QueryCompany(tarsCtx context.Context, offset int32, limit int32, Companys *[]CivetCompanyCenter.Company, count *int32, result *int32) (ret int32, err error) {
	if ins.QueryRowsAllWithOSLM(int(offset), int(limit), count, Companys) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *CompanyManProcessImp) QueryCompanyByCondition(tarsCtx context.Context, offset int32, limit int32, keys string, value string, Companys *[]CivetCompanyCenter.Company, count *int32, result *int32) (ret int32, err error) {
	if ins.QueryRowsAllWithConditionOSLM(int(offset), int(limit), Companys, count, keys, value) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *CompanyManProcessImp) BindCompanyAndRole(tarsCtx context.Context, newCompanyHasRole *CivetCompanyCenter.CompanyHasRole, result *int32) (ret int32, err error) {
	res := ins.CreateRow(newCompanyHasRole)
	if res == true {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("newCompanyHasRole Fail")
	}
}

func (imp *CompanyManProcessImp) UnBindCompanyAndRole(tarsCtx context.Context, CompanyHasRoleID int32, result *int32) (ret int32, err error) {
	if ins.DelRowByCondition(&CivetCompanyCenter.CompanyHasRole{}, "id", CompanyHasRoleID) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("del Fail")
	}
}

func (imp *CompanyManProcessImp) QueryCompanyRole(tarsCtx context.Context, offset int32, limit int32, Companys *[]CivetCompanyCenter.CompanyHasRole, count *int32, result *int32) (ret int32, err error) {
	if ins.QueryRowsAllWithOSLM(int(offset), int(limit), count, Companys) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *CompanyManProcessImp) QueryCompanyRoleByCondition(tarsCtx context.Context, offset int32, limit int32, keys string, value string, binds *[]CivetCompanyCenter.CompanyHasRole, count *int32, result *int32) (ret int32, err error) {
	if ins.QueryRowsAllWithConditionOSLM(int(offset), int(limit), binds, count, keys, value) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *CompanyManProcessImp) CreateRole(tarsCtx context.Context, newRole *CivetCompanyCenter.Role, c *int32) (ret int32, err error) {
	res := ins.CreateRow(newRole)
	if res == true {
		*c = 1
		return 0, nil
	} else {
		*c = -1
		return 0, errors.New("CreateRole Fail")
	}
}

func (imp *CompanyManProcessImp) DeleteRole(tarsCtx context.Context, roleID int32, result *int32) (ret int32, err error) {
	if ins.DelRowByCondition(&CivetCompanyCenter.Role{}, "id", roleID) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *CompanyManProcessImp) UpdateRole(tarsCtx context.Context, role *CivetCompanyCenter.Role, keys string, value string, result *int32) (ret int32, err error) {
	if ins.EditRowByConditionbyModel(role, keys, value) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

func (imp *CompanyManProcessImp) QueryRole(tarsCtx context.Context, offset int32, limit int32, role *[]CivetCompanyCenter.Role, count *int32, result *int32) (ret int32, err error) {
	if ins.QueryRowsAllWithOSLM(int(offset), int(limit), count, role) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *CompanyManProcessImp) QueryRoleByCondition(tarsCtx context.Context, offset int32, limit int32, keys string, value string, role *[]CivetCompanyCenter.Role, count *int32, result *int32) (ret int32, err error) {
	if ins.QueryRowsAllWithConditionOSLM(int(offset), int(limit), role, count, keys, value) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *CompanyManProcessImp) BindRoleAndPermissions(tarsCtx context.Context, newRoleHasPermission *CivetCompanyCenter.RoleHasPermission, result *int32) (ret int32, err error) {
	res := ins.CreateRow(newRoleHasPermission)
	if res == true {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("Create BindRoleAndPermissions Fail")
	}
}

func (imp *CompanyManProcessImp) UnBindRoleAndPermissions(tarsCtx context.Context, RoleID int32, result *int32) (ret int32, err error) {
	if ins.DelRowByCondition(&CivetCompanyCenter.RoleHasPermission{}, "id", RoleID) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("del Fail")
	}
}

func (imp *CompanyManProcessImp) QueryRolePermissions(tarsCtx context.Context, offset int32, limit int32, RoleHasPermission *[]CivetCompanyCenter.RoleHasPermission, count *int32, result *int32) (ret int32, err error) {
	if ins.QueryRowsAllWithOSLM(int(offset), int(limit), count, RoleHasPermission) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *CompanyManProcessImp) QueryRolePermissionsByCondition(tarsCtx context.Context, offset int32, limit int32, keys string, value string, binds *[]CivetCompanyCenter.RoleHasPermission, count *int32, result *int32) (ret int32, err error) {
	if ins.QueryRowsAllWithConditionOSLM(int(offset), int(limit), binds, count, keys, value) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *CompanyManProcessImp) CreatePermissions(tarsCtx context.Context, newCPermission *CivetCompanyCenter.CPermission, c *int32) (ret int32, err error) {
	res := ins.CreateRow(newCPermission)
	if res == true {
		*c = 1
		return 0, nil
	} else {
		*c = -1
		return 0, errors.New("CreatePermissions Fail")
	}
}

func (imp *CompanyManProcessImp) DeletePermissions(tarsCtx context.Context, PermissionsID int32, result *int32) (ret int32, err error) {
	if ins.DelRowByCondition(&CivetCompanyCenter.CPermission{}, "id", PermissionsID) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *CompanyManProcessImp) UpdatePermissions(tarsCtx context.Context, PermissionsID *CivetCompanyCenter.CPermission, keys string, value string, result *int32) (ret int32, err error) {
	if ins.EditRowByConditionbyModel(PermissionsID, keys, value) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

func (imp *CompanyManProcessImp) QueryPermissions(tarsCtx context.Context, offset int32, limit int32, Permissions *[]CivetCompanyCenter.CPermission, count *int32, result *int32) (ret int32, err error) {
	if ins.QueryRowsAllWithOSLM(int(offset), int(limit), count, Permissions) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *CompanyManProcessImp) QueryPermissionsByCondition(tarsCtx context.Context, offset int32, limit int32, keys string, value string, Permissions *[]CivetCompanyCenter.CPermission, count *int32, result *int32) (ret int32, err error) {
	if ins.QueryRowsAllWithConditionOSLM(int(offset), int(limit), Permissions, count, keys, value) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *CompanyManProcessImp) CheckCompanyHasPermissionsByID(tarsCtx context.Context, CompanyID int32, PermissionID int32, result *int32) (ret int32, err error) {
	var ressult []CivetCompanyCenter.CPermission
	var res int32
	imp.QueryCompanyPermissions(tarsCtx, CompanyID, &ressult, &res)
	if res == 1 {
		for _, right := range ressult {
			if right.Id == PermissionID {
				fmt.Println(right.Tag)
				fmt.Println(PermissionID)
				*result = 1
				return 0, nil
			}
		}
		*result = -1
		return 0, nil
	}
	return 0, nil
}

func (imp *CompanyManProcessImp) CheckCompanyHasPermissionsByTag(tarsCtx context.Context, CompanyID int32, PermissionsTag string, result *int32) (ret int32, err error) {
	var ressult []CivetCompanyCenter.CPermission
	var res int32
	imp.QueryCompanyPermissions(tarsCtx, CompanyID, &ressult, &res)
	if res == 1 {
		for _, right := range ressult {
			if right.Tag == PermissionsTag {
				fmt.Println(right.Tag)
				fmt.Println(PermissionsTag)
				*result = 1
				return 0, nil
			}
		}
		*result = -1
		return 0, nil
	}
	return 0, nil
}

func (imp *CompanyManProcessImp) QueryCompanyPermissions(tarsCtx context.Context, CompanyID int32, Permissions *[]CivetCompanyCenter.CPermission, result *int32) (ret int32, err error) {
	var ressult []CivetCompanyCenter.CPermission
	rows, err := ins.Db.Raw("SELECT c_rights.id,c_rights.name,c_rights.summarize,c_rights.tag FROM companies INNER JOIN company_has_roles ON company_has_roles.company_id=companies.role_id INNER JOIN roles on roles.id = company_has_roles.role_id INNER JOIN role_has_rights ON roles.id = role_has_rights.role_id INNER JOIN c_rights on c_rights.id = role_has_rights.right_id AND companies.id =?", CompanyID).Rows()
	defer rows.Close()

	fmt.Println(err)
	for rows.Next() {
		// ScanRows 将一行扫描至 user
		tmp := CivetCompanyCenter.CPermission{}
		ins.Db.ScanRows(rows, &tmp)
		ressult = append(ressult, tmp)
		// 业务逻辑...
	}
	fmt.Println(rows)
	*Permissions = ressult
	*result = 1
	fmt.Println(ressult)
	return 0, nil
}

type Config struct {
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBHost     string `json:"DBHost"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
	RDUserName string `json:"RDUserName"`
	RDPassword string `json:"RDPassword"`
	RDHost     string `json:"RDHost"`
	RDPort     string `json:"RDPort"`
	RDSize     int    `json:"RDSize"`
}

var ins *CivetTarsDataBase.CivetData

// Init servant init
func (imp *CompanyManProcessImp) Init() error {
	var casd = tars.NewRConf("CivetCompanyCenter", "CompanyManServer", tars.GetServerConfig().BasePath)
	config, _ := casd.GetConfig("CompanyManServer.conf")
	var mc Config
	err := json.Unmarshal([]byte(config), &mc)
	if err != nil {
		print(err)
		return err
	}
	if mc.DBName == "" || mc.DBHost == "" || mc.DBUserName == "" {
		fmt.Println("The Mysql Config error")
		os.Exit(1)
	} else {
		fmt.Println("the config is", config, mc.DBHost, mc.DBPassword, mc.DBUserName, mc.DBName)
	}

	ins = CivetTarsDataBase.CreateCivetData(mc.DBUserName, mc.DBPassword, mc.DBHost, mc.DBPort)
	res := ins.ConnectOrCreateDataBase(mc.DBName)
	if res == true {
		if ins.CheckTableExist("companies") == true {
			tars.TLOG.Debug("Init Done")
		} else {
			createResult := ins.CreateTable(&CivetCompanyCenter.CPermission{}) &&
				ins.CreateTable(&CivetCompanyCenter.Role{}) &&
				ins.CreateTable(&CivetCompanyCenter.Company{}) &&
				ins.CreateTable(&CivetCompanyCenter.RoleHasPermission{}) &&
				ins.CreateTable(&CivetCompanyCenter.CompanyHasRole{})
			if createResult != true {
				fmt.Println("Create TB Fail")
				tars.TLOG.Error("Create TB Fail")
				os.Exit(1)
			}
		}

	} else {
		tars.TLOG.Error("Connect DB Fail")
		os.Exit(1)
	}
	return nil
}

// Destroy servant destory
func (imp *CompanyManProcessImp) Destroy() {
	//destroy servant here:
	//...
}
