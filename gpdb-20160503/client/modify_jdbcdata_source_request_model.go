// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJDBCDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyJDBCDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceDescription(v string) *ModifyJDBCDataSourceRequest
	GetDataSourceDescription() *string
	SetDataSourceId(v string) *ModifyJDBCDataSourceRequest
	GetDataSourceId() *string
	SetDataSourceType(v string) *ModifyJDBCDataSourceRequest
	GetDataSourceType() *string
	SetJDBCConnectionString(v string) *ModifyJDBCDataSourceRequest
	GetJDBCConnectionString() *string
	SetJDBCPassword(v string) *ModifyJDBCDataSourceRequest
	GetJDBCPassword() *string
	SetJDBCUserName(v string) *ModifyJDBCDataSourceRequest
	GetJDBCUserName() *string
	SetRegionId(v string) *ModifyJDBCDataSourceRequest
	GetRegionId() *string
}

type ModifyJDBCDataSourceRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Data source description.
	//
	// example:
	//
	// test
	DataSourceDescription *string `json:"DataSourceDescription,omitempty" xml:"DataSourceDescription,omitempty"`
	// Data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Database type: - mysql - postgresql - sqlserver
	//
	// example:
	//
	// mysql
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The JDBC connection string.
	//
	// example:
	//
	// xxxxxx
	JDBCConnectionString *string `json:"JDBCConnectionString,omitempty" xml:"JDBCConnectionString,omitempty"`
	// The password of the database account.
	//
	// example:
	//
	// xxxxxx
	JDBCPassword *string `json:"JDBCPassword,omitempty" xml:"JDBCPassword,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// xxxxxx
	JDBCUserName *string `json:"JDBCUserName,omitempty" xml:"JDBCUserName,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyJDBCDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyJDBCDataSourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyJDBCDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyJDBCDataSourceRequest) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *ModifyJDBCDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ModifyJDBCDataSourceRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ModifyJDBCDataSourceRequest) GetJDBCConnectionString() *string {
	return s.JDBCConnectionString
}

func (s *ModifyJDBCDataSourceRequest) GetJDBCPassword() *string {
	return s.JDBCPassword
}

func (s *ModifyJDBCDataSourceRequest) GetJDBCUserName() *string {
	return s.JDBCUserName
}

func (s *ModifyJDBCDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyJDBCDataSourceRequest) SetDBInstanceId(v string) *ModifyJDBCDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyJDBCDataSourceRequest) SetDataSourceDescription(v string) *ModifyJDBCDataSourceRequest {
	s.DataSourceDescription = &v
	return s
}

func (s *ModifyJDBCDataSourceRequest) SetDataSourceId(v string) *ModifyJDBCDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *ModifyJDBCDataSourceRequest) SetDataSourceType(v string) *ModifyJDBCDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *ModifyJDBCDataSourceRequest) SetJDBCConnectionString(v string) *ModifyJDBCDataSourceRequest {
	s.JDBCConnectionString = &v
	return s
}

func (s *ModifyJDBCDataSourceRequest) SetJDBCPassword(v string) *ModifyJDBCDataSourceRequest {
	s.JDBCPassword = &v
	return s
}

func (s *ModifyJDBCDataSourceRequest) SetJDBCUserName(v string) *ModifyJDBCDataSourceRequest {
	s.JDBCUserName = &v
	return s
}

func (s *ModifyJDBCDataSourceRequest) SetRegionId(v string) *ModifyJDBCDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyJDBCDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
