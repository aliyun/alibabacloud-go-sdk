// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckJDBCSourceNetConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CheckJDBCSourceNetConnectionRequest
	GetDBInstanceId() *string
	SetDataSourceId(v string) *CheckJDBCSourceNetConnectionRequest
	GetDataSourceId() *string
	SetJdbcConnectionString(v string) *CheckJDBCSourceNetConnectionRequest
	GetJdbcConnectionString() *string
	SetRegionId(v string) *CheckJDBCSourceNetConnectionRequest
	GetRegionId() *string
}

type CheckJDBCSourceNetConnectionRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Either DataSourceId or JdbcConnectionString must be specified as input, otherwise an error will occur. If both parameters are specified, JdbcConnectionString will be used preferentially.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// JDBC connection string.
	//
	// example:
	//
	// jdbc:mysql://rm-xxx.mysql.rds.aliyuncs.com:3306/testadmin
	JdbcConnectionString *string `json:"JdbcConnectionString,omitempty" xml:"JdbcConnectionString,omitempty"`
	// The ID of the region where the instance is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckJDBCSourceNetConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckJDBCSourceNetConnectionRequest) GoString() string {
	return s.String()
}

func (s *CheckJDBCSourceNetConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CheckJDBCSourceNetConnectionRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CheckJDBCSourceNetConnectionRequest) GetJdbcConnectionString() *string {
	return s.JdbcConnectionString
}

func (s *CheckJDBCSourceNetConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckJDBCSourceNetConnectionRequest) SetDBInstanceId(v string) *CheckJDBCSourceNetConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckJDBCSourceNetConnectionRequest) SetDataSourceId(v string) *CheckJDBCSourceNetConnectionRequest {
	s.DataSourceId = &v
	return s
}

func (s *CheckJDBCSourceNetConnectionRequest) SetJdbcConnectionString(v string) *CheckJDBCSourceNetConnectionRequest {
	s.JdbcConnectionString = &v
	return s
}

func (s *CheckJDBCSourceNetConnectionRequest) SetRegionId(v string) *CheckJDBCSourceNetConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CheckJDBCSourceNetConnectionRequest) Validate() error {
	return dara.Validate(s)
}
