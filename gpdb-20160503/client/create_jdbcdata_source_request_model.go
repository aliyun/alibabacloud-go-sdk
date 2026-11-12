// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJDBCDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateJDBCDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceDescription(v string) *CreateJDBCDataSourceRequest
	GetDataSourceDescription() *string
	SetDataSourceName(v string) *CreateJDBCDataSourceRequest
	GetDataSourceName() *string
	SetDataSourceType(v string) *CreateJDBCDataSourceRequest
	GetDataSourceType() *string
	SetJDBCConnectionString(v string) *CreateJDBCDataSourceRequest
	GetJDBCConnectionString() *string
	SetJDBCPassword(v string) *CreateJDBCDataSourceRequest
	GetJDBCPassword() *string
	SetJDBCUserName(v string) *CreateJDBCDataSourceRequest
	GetJDBCUserName() *string
	SetRegionId(v string) *CreateJDBCDataSourceRequest
	GetRegionId() *string
}

type CreateJDBCDataSourceRequest struct {
	// The instance ID.
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
	// The name of data soruce
	//
	// example:
	//
	// jdbc_pxf
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The type of the data source.
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
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateJDBCDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJDBCDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateJDBCDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateJDBCDataSourceRequest) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *CreateJDBCDataSourceRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateJDBCDataSourceRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CreateJDBCDataSourceRequest) GetJDBCConnectionString() *string {
	return s.JDBCConnectionString
}

func (s *CreateJDBCDataSourceRequest) GetJDBCPassword() *string {
	return s.JDBCPassword
}

func (s *CreateJDBCDataSourceRequest) GetJDBCUserName() *string {
	return s.JDBCUserName
}

func (s *CreateJDBCDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateJDBCDataSourceRequest) SetDBInstanceId(v string) *CreateJDBCDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateJDBCDataSourceRequest) SetDataSourceDescription(v string) *CreateJDBCDataSourceRequest {
	s.DataSourceDescription = &v
	return s
}

func (s *CreateJDBCDataSourceRequest) SetDataSourceName(v string) *CreateJDBCDataSourceRequest {
	s.DataSourceName = &v
	return s
}

func (s *CreateJDBCDataSourceRequest) SetDataSourceType(v string) *CreateJDBCDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateJDBCDataSourceRequest) SetJDBCConnectionString(v string) *CreateJDBCDataSourceRequest {
	s.JDBCConnectionString = &v
	return s
}

func (s *CreateJDBCDataSourceRequest) SetJDBCPassword(v string) *CreateJDBCDataSourceRequest {
	s.JDBCPassword = &v
	return s
}

func (s *CreateJDBCDataSourceRequest) SetJDBCUserName(v string) *CreateJDBCDataSourceRequest {
	s.JDBCUserName = &v
	return s
}

func (s *CreateJDBCDataSourceRequest) SetRegionId(v string) *CreateJDBCDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateJDBCDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
