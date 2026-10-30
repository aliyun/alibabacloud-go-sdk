// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourcePostgreSQLParameters interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *SourcePostgreSQLParameters
	GetDatabaseName() *string
	SetHostName(v string) *SourcePostgreSQLParameters
	GetHostName() *string
	SetNetworkType(v string) *SourcePostgreSQLParameters
	GetNetworkType() *string
	SetPassword(v string) *SourcePostgreSQLParameters
	GetPassword() *string
	SetPort(v int32) *SourcePostgreSQLParameters
	GetPort() *int32
	SetRegionId(v string) *SourcePostgreSQLParameters
	GetRegionId() *string
	SetSchemaName(v string) *SourcePostgreSQLParameters
	GetSchemaName() *string
	SetSecurityGroupId(v string) *SourcePostgreSQLParameters
	GetSecurityGroupId() *string
	SetSnapshotMode(v string) *SourcePostgreSQLParameters
	GetSnapshotMode() *string
	SetTableNames(v string) *SourcePostgreSQLParameters
	GetTableNames() *string
	SetUser(v string) *SourcePostgreSQLParameters
	GetUser() *string
	SetVSwitchIds(v string) *SourcePostgreSQLParameters
	GetVSwitchIds() *string
	SetVpcId(v string) *SourcePostgreSQLParameters
	GetVpcId() *string
}

type SourcePostgreSQLParameters struct {
	// The name of the PostgreSQL database to connect to.
	//
	// example:
	//
	// analytics_db
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The hostname or IP address of the PostgreSQL database server.
	//
	// example:
	//
	// pgm-bp1xxxxx.pg.rds.aliyuncs.com
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The network type used to connect to the PostgreSQL data source. Valid values: PublicNetwork (public network) and PrivateNetwork (VPC).
	//
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The password of the database connection user.
	//
	// example:
	//
	// your_password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The connection port of the PostgreSQL database server.
	//
	// example:
	//
	// 5432
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region where the PostgreSQL data source resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the PostgreSQL schema to synchronize.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The ID of the security group used to connect to the PostgreSQL instance.
	//
	// example:
	//
	// sg-a1b2c3d4e5f6g7h8i9j0
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The snapshot mode used when PostgreSQL data synchronization is started for the first time.
	//
	// example:
	//
	// when_needed
	SnapshotMode *string `json:"SnapshotMode,omitempty" xml:"SnapshotMode,omitempty"`
	// The names of the PostgreSQL tables to synchronize. Separate multiple table names with commas. Each table name must be in the {SchemaName}.{TableName} format.
	//
	// example:
	//
	// test_schema.test_table
	TableNames *string `json:"TableNames,omitempty" xml:"TableNames,omitempty"`
	// The username for the database connection.
	//
	// example:
	//
	// readonly_user
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The ID of the vSwitch in the VPC.
	//
	// example:
	//
	// [vsw-a1b2c3d4e5f6g7h8i9j0k]
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the VPC where the PostgreSQL instance resides.
	//
	// example:
	//
	// vpc-a1b2c3d4e5f6g7h8i9j0k
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SourcePostgreSQLParameters) String() string {
	return dara.Prettify(s)
}

func (s SourcePostgreSQLParameters) GoString() string {
	return s.String()
}

func (s *SourcePostgreSQLParameters) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *SourcePostgreSQLParameters) GetHostName() *string {
	return s.HostName
}

func (s *SourcePostgreSQLParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SourcePostgreSQLParameters) GetPassword() *string {
	return s.Password
}

func (s *SourcePostgreSQLParameters) GetPort() *int32 {
	return s.Port
}

func (s *SourcePostgreSQLParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *SourcePostgreSQLParameters) GetSchemaName() *string {
	return s.SchemaName
}

func (s *SourcePostgreSQLParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SourcePostgreSQLParameters) GetSnapshotMode() *string {
	return s.SnapshotMode
}

func (s *SourcePostgreSQLParameters) GetTableNames() *string {
	return s.TableNames
}

func (s *SourcePostgreSQLParameters) GetUser() *string {
	return s.User
}

func (s *SourcePostgreSQLParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *SourcePostgreSQLParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *SourcePostgreSQLParameters) SetDatabaseName(v string) *SourcePostgreSQLParameters {
	s.DatabaseName = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetHostName(v string) *SourcePostgreSQLParameters {
	s.HostName = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetNetworkType(v string) *SourcePostgreSQLParameters {
	s.NetworkType = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetPassword(v string) *SourcePostgreSQLParameters {
	s.Password = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetPort(v int32) *SourcePostgreSQLParameters {
	s.Port = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetRegionId(v string) *SourcePostgreSQLParameters {
	s.RegionId = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetSchemaName(v string) *SourcePostgreSQLParameters {
	s.SchemaName = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetSecurityGroupId(v string) *SourcePostgreSQLParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetSnapshotMode(v string) *SourcePostgreSQLParameters {
	s.SnapshotMode = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetTableNames(v string) *SourcePostgreSQLParameters {
	s.TableNames = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetUser(v string) *SourcePostgreSQLParameters {
	s.User = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetVSwitchIds(v string) *SourcePostgreSQLParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SourcePostgreSQLParameters) SetVpcId(v string) *SourcePostgreSQLParameters {
	s.VpcId = &v
	return s
}

func (s *SourcePostgreSQLParameters) Validate() error {
	return dara.Validate(s)
}
