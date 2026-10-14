// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestAgenticDataSourceConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *TestAgenticDataSourceConnectionRequest
	GetAlias() *string
	SetAuthAccount(v string) *TestAgenticDataSourceConnectionRequest
	GetAuthAccount() *string
	SetAuthPassword(v string) *TestAgenticDataSourceConnectionRequest
	GetAuthPassword() *string
	SetAuthType(v int32) *TestAgenticDataSourceConnectionRequest
	GetAuthType() *int32
	SetCatalogName(v string) *TestAgenticDataSourceConnectionRequest
	GetCatalogName() *string
	SetDataSourceUuid(v string) *TestAgenticDataSourceConnectionRequest
	GetDataSourceUuid() *string
	SetDbType(v string) *TestAgenticDataSourceConnectionRequest
	GetDbType() *string
	SetEnvType(v string) *TestAgenticDataSourceConnectionRequest
	GetEnvType() *string
	SetHost(v string) *TestAgenticDataSourceConnectionRequest
	GetHost() *string
	SetInstanceResourceId(v string) *TestAgenticDataSourceConnectionRequest
	GetInstanceResourceId() *string
	SetInstanceSource(v string) *TestAgenticDataSourceConnectionRequest
	GetInstanceSource() *string
	SetKmsSecret(v string) *TestAgenticDataSourceConnectionRequest
	GetKmsSecret() *string
	SetPort(v int32) *TestAgenticDataSourceConnectionRequest
	GetPort() *int32
	SetRoleArn(v string) *TestAgenticDataSourceConnectionRequest
	GetRoleArn() *string
	SetSchema(v string) *TestAgenticDataSourceConnectionRequest
	GetSchema() *string
	SetSnowflakeRole(v string) *TestAgenticDataSourceConnectionRequest
	GetSnowflakeRole() *string
	SetSnowflakeWarehouse(v string) *TestAgenticDataSourceConnectionRequest
	GetSnowflakeWarehouse() *string
	SetSyncAlias(v int32) *TestAgenticDataSourceConnectionRequest
	GetSyncAlias() *int32
	SetUseSsl(v int32) *TestAgenticDataSourceConnectionRequest
	GetUseSsl() *int32
	SetVSwitchId(v string) *TestAgenticDataSourceConnectionRequest
	GetVSwitchId() *string
	SetVpcId(v string) *TestAgenticDataSourceConnectionRequest
	GetVpcId() *string
	SetZoneId(v string) *TestAgenticDataSourceConnectionRequest
	GetZoneId() *string
}

type TestAgenticDataSourceConnectionRequest struct {
	// The alias. You can leave this parameter empty when testing a connection. This parameter does not create or modify a data source.
	//
	// example:
	//
	// demo_mysql
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The logon username. You can leave this parameter empty for authentication methods that do not require a username.
	//
	// example:
	//
	// db_user
	AuthAccount *string `json:"AuthAccount,omitempty" xml:"AuthAccount,omitempty"`
	// The logon password. When testing an existing data source, you can leave this parameter empty to reuse the saved password. If a new value is passed, the new value takes precedence.
	//
	// example:
	//
	// P@ss****
	AuthPassword *string `json:"AuthPassword,omitempty" xml:"AuthPassword,omitempty"`
	// The authentication type. This parameter is optional. Valid values: -1 (no authentication required), 0 (username and password), 1 (password only), and 2 (password-free).
	//
	// example:
	//
	// 0
	AuthType *int32 `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The catalog name. You can leave this parameter empty when testing a connection. This parameter does not create a metadata catalog.
	//
	// example:
	//
	// demo_catalog
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The UUID of the data source. This parameter is required when testing an existing data source. If no new password is passed, the saved password is reused. Leave this parameter empty when testing a new data source.
	//
	// example:
	//
	// cn-hangzhou_a1b2c3d4e5f6
	DataSourceUuid *string `json:"DataSourceUuid,omitempty" xml:"DataSourceUuid,omitempty"`
	// The type of the data source. Use uppercase letters. Do not pass numbers. Valid values: MYSQL, POSTGRESQL, ORACLE, SQLSERVER, MARIADB, POLARDB, POLARDB_POSTGRESQL, POLARDB_ORACLE, OCEANBASEX, DRDS, ADB_MYSQL, ADB_POSTGRESQL, CLICKHOUSE, SELECTDB, STARROCKS, HOLOGRES, MAXCOMPUTE, ES_REST, and TRINO.
	//
	// This parameter is required.
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The environment type. This parameter is optional when testing a connection. Valid values: product, pre, dev, test, sit, uat, pet, and stag.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The domain name or IP address of the database. Do not include the protocol, path, or port. This parameter is optional. You can leave this parameter empty when you select an existing cloud instance and the endpoint is automatically populated.
	//
	// example:
	//
	// mysql.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the cloud resource instance. This parameter is required when you select an existing cloud instance. You can leave this parameter empty when you enter a connection endpoint.
	//
	// example:
	//
	// rm-bp1234567890
	InstanceResourceId *string `json:"InstanceResourceId,omitempty" xml:"InstanceResourceId,omitempty"`
	// The source of the data source. Valid values:
	//
	// - RDS: ApsaraDB instance.
	//
	// - PUBLIC_OWN: self-managed database connected over the Internet.
	//
	// - ECS_OWN: self-managed database hosted on an ECS instance.
	//
	// - GATEWAY: database connected through Database Gateway.
	//
	// - VPC_IDC: database connected over a VPC or data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The name of the KMS secret. This parameter is required when you use KMS to manage database credentials. Leave this parameter empty in other scenarios.
	//
	// example:
	//
	// datasource-secret
	KmsSecret *string `json:"KmsSecret,omitempty" xml:"KmsSecret,omitempty"`
	// The service port of the database. This parameter is optional. You can leave this parameter empty when you select an existing cloud instance and the port is automatically populated.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The complete ARN of the RAM role. This parameter is required when you use an Alibaba Cloud RAM role to access the data source. Leave this parameter empty in other scenarios.
	//
	// example:
	//
	// acs:ram::123456789012****:role/DataSourceRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The schema name. This parameter is optional. For ADB_POSTGRESQL, DRDS, POSTGRESQL, and POLARDB_POSTGRESQL, enter the database name. For ORACLE, enter the SID or service name, and prefix the service name with a forward slash (/). For POLARDB_ORACLE, enter the SID. For MAXCOMPUTE, enter the project name. Leave this parameter empty for other types.
	//
	// example:
	//
	// postgres
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The name of the Snowflake role. This parameter takes effect only when DbType is set to SNOWFLAKE.
	//
	// example:
	//
	// ROLE
	SnowflakeRole *string `json:"SnowflakeRole,omitempty" xml:"SnowflakeRole,omitempty"`
	// The name of the Snowflake compute warehouse. This parameter takes effect only when DbType is set to SNOWFLAKE.
	//
	// example:
	//
	// WAREHOUSE
	SnowflakeWarehouse *string `json:"SnowflakeWarehouse,omitempty" xml:"SnowflakeWarehouse,omitempty"`
	// Specifies whether to synchronize the alias. This parameter is optional when testing a connection. This parameter does not modify the data source name.
	//
	// example:
	//
	// 0
	SyncAlias *int32 `json:"SyncAlias,omitempty" xml:"SyncAlias,omitempty"`
	// Specifies whether to enable SSL. Valid values:
	//
	// - 1: Enable.
	//
	// - 0: Disable.
	//
	// - -1: Use the default setting.
	//
	// SSL is supported only for ADB_MYSQL, CLICKHOUSE, MARIADB, MYSQL, OCEANBASEX, POLARDB, SELECTDB, and STARROCKS. Leave this parameter empty for other types.
	//
	// example:
	//
	// 1
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	// The ID of the vSwitch. This parameter is required when you access MaxCompute over a VPC. Leave this parameter empty in other scenarios.
	//
	// example:
	//
	// vsw-bp1234567890
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC. This parameter is required when you access MaxCompute over a VPC. Leave this parameter empty in other scenarios.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID. This parameter is required when you access MaxCompute over a VPC. Leave this parameter empty in other scenarios.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s TestAgenticDataSourceConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s TestAgenticDataSourceConnectionRequest) GoString() string {
	return s.String()
}

func (s *TestAgenticDataSourceConnectionRequest) GetAlias() *string {
	return s.Alias
}

func (s *TestAgenticDataSourceConnectionRequest) GetAuthAccount() *string {
	return s.AuthAccount
}

func (s *TestAgenticDataSourceConnectionRequest) GetAuthPassword() *string {
	return s.AuthPassword
}

func (s *TestAgenticDataSourceConnectionRequest) GetAuthType() *int32 {
	return s.AuthType
}

func (s *TestAgenticDataSourceConnectionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *TestAgenticDataSourceConnectionRequest) GetDataSourceUuid() *string {
	return s.DataSourceUuid
}

func (s *TestAgenticDataSourceConnectionRequest) GetDbType() *string {
	return s.DbType
}

func (s *TestAgenticDataSourceConnectionRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *TestAgenticDataSourceConnectionRequest) GetHost() *string {
	return s.Host
}

func (s *TestAgenticDataSourceConnectionRequest) GetInstanceResourceId() *string {
	return s.InstanceResourceId
}

func (s *TestAgenticDataSourceConnectionRequest) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *TestAgenticDataSourceConnectionRequest) GetKmsSecret() *string {
	return s.KmsSecret
}

func (s *TestAgenticDataSourceConnectionRequest) GetPort() *int32 {
	return s.Port
}

func (s *TestAgenticDataSourceConnectionRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *TestAgenticDataSourceConnectionRequest) GetSchema() *string {
	return s.Schema
}

func (s *TestAgenticDataSourceConnectionRequest) GetSnowflakeRole() *string {
	return s.SnowflakeRole
}

func (s *TestAgenticDataSourceConnectionRequest) GetSnowflakeWarehouse() *string {
	return s.SnowflakeWarehouse
}

func (s *TestAgenticDataSourceConnectionRequest) GetSyncAlias() *int32 {
	return s.SyncAlias
}

func (s *TestAgenticDataSourceConnectionRequest) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *TestAgenticDataSourceConnectionRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *TestAgenticDataSourceConnectionRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *TestAgenticDataSourceConnectionRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *TestAgenticDataSourceConnectionRequest) SetAlias(v string) *TestAgenticDataSourceConnectionRequest {
	s.Alias = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetAuthAccount(v string) *TestAgenticDataSourceConnectionRequest {
	s.AuthAccount = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetAuthPassword(v string) *TestAgenticDataSourceConnectionRequest {
	s.AuthPassword = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetAuthType(v int32) *TestAgenticDataSourceConnectionRequest {
	s.AuthType = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetCatalogName(v string) *TestAgenticDataSourceConnectionRequest {
	s.CatalogName = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetDataSourceUuid(v string) *TestAgenticDataSourceConnectionRequest {
	s.DataSourceUuid = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetDbType(v string) *TestAgenticDataSourceConnectionRequest {
	s.DbType = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetEnvType(v string) *TestAgenticDataSourceConnectionRequest {
	s.EnvType = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetHost(v string) *TestAgenticDataSourceConnectionRequest {
	s.Host = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetInstanceResourceId(v string) *TestAgenticDataSourceConnectionRequest {
	s.InstanceResourceId = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetInstanceSource(v string) *TestAgenticDataSourceConnectionRequest {
	s.InstanceSource = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetKmsSecret(v string) *TestAgenticDataSourceConnectionRequest {
	s.KmsSecret = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetPort(v int32) *TestAgenticDataSourceConnectionRequest {
	s.Port = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetRoleArn(v string) *TestAgenticDataSourceConnectionRequest {
	s.RoleArn = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetSchema(v string) *TestAgenticDataSourceConnectionRequest {
	s.Schema = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetSnowflakeRole(v string) *TestAgenticDataSourceConnectionRequest {
	s.SnowflakeRole = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetSnowflakeWarehouse(v string) *TestAgenticDataSourceConnectionRequest {
	s.SnowflakeWarehouse = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetSyncAlias(v int32) *TestAgenticDataSourceConnectionRequest {
	s.SyncAlias = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetUseSsl(v int32) *TestAgenticDataSourceConnectionRequest {
	s.UseSsl = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetVSwitchId(v string) *TestAgenticDataSourceConnectionRequest {
	s.VSwitchId = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetVpcId(v string) *TestAgenticDataSourceConnectionRequest {
	s.VpcId = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) SetZoneId(v string) *TestAgenticDataSourceConnectionRequest {
	s.ZoneId = &v
	return s
}

func (s *TestAgenticDataSourceConnectionRequest) Validate() error {
	return dara.Validate(s)
}
