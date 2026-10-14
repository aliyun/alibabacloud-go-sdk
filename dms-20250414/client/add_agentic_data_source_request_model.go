// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgenticDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *AddAgenticDataSourceRequest
	GetAlias() *string
	SetAuthAccount(v string) *AddAgenticDataSourceRequest
	GetAuthAccount() *string
	SetAuthPassword(v string) *AddAgenticDataSourceRequest
	GetAuthPassword() *string
	SetAuthType(v int32) *AddAgenticDataSourceRequest
	GetAuthType() *int32
	SetCatalogName(v string) *AddAgenticDataSourceRequest
	GetCatalogName() *string
	SetCreateMode(v string) *AddAgenticDataSourceRequest
	GetCreateMode() *string
	SetDbType(v string) *AddAgenticDataSourceRequest
	GetDbType() *string
	SetEnvType(v string) *AddAgenticDataSourceRequest
	GetEnvType() *string
	SetHost(v string) *AddAgenticDataSourceRequest
	GetHost() *string
	SetInstanceResourceId(v string) *AddAgenticDataSourceRequest
	GetInstanceResourceId() *string
	SetInstanceSource(v string) *AddAgenticDataSourceRequest
	GetInstanceSource() *string
	SetKmsSecret(v string) *AddAgenticDataSourceRequest
	GetKmsSecret() *string
	SetPolicyId(v string) *AddAgenticDataSourceRequest
	GetPolicyId() *string
	SetPort(v int32) *AddAgenticDataSourceRequest
	GetPort() *int32
	SetRoleArn(v string) *AddAgenticDataSourceRequest
	GetRoleArn() *string
	SetSchema(v string) *AddAgenticDataSourceRequest
	GetSchema() *string
	SetSnowflakeRole(v string) *AddAgenticDataSourceRequest
	GetSnowflakeRole() *string
	SetSnowflakeWarehouse(v string) *AddAgenticDataSourceRequest
	GetSnowflakeWarehouse() *string
	SetSyncAlias(v int32) *AddAgenticDataSourceRequest
	GetSyncAlias() *int32
	SetUseSsl(v int32) *AddAgenticDataSourceRequest
	GetUseSsl() *int32
	SetVSwitchId(v string) *AddAgenticDataSourceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *AddAgenticDataSourceRequest
	GetVpcId() *string
	SetZoneId(v string) *AddAgenticDataSourceRequest
	GetZoneId() *string
}

type AddAgenticDataSourceRequest struct {
	// The display name of the data source. If you do not specify this parameter, the server generates a name.
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
	// The logon password. This is sensitive information. Do not record or disclose it.
	//
	// example:
	//
	// P@ss****
	AuthPassword *string `json:"AuthPassword,omitempty" xml:"AuthPassword,omitempty"`
	// Optional. The authentication type. Valid values:
	//
	// - -1: No authentication required.
	//
	// - 0: Username and password.
	//
	// - 1: Password only.
	//
	// - 2: Password-free.
	//
	// example:
	//
	// 0
	AuthType *int32 `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The name of the metadata catalog. If you do not specify this parameter, the name is automatically generated. The name can be up to 128 characters in length. It must start with a Chinese character, a letter, or an underscore, and can contain digits and hyphens. It cannot contain only digits.
	//
	// example:
	//
	// demo_catalog
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The entry method. Valid values:
	//
	// - INSTANCE_ID: Select an existing cloud instance.
	//
	// - HOST_PORT: Enter the endpoint and port.
	//
	// If you do not specify this parameter, HOST_PORT is used.
	//
	// example:
	//
	// HOST_PORT
	CreateMode *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// The type of the data source. Use uppercase letters. Do not pass numbers. Valid values: MYSQL, POSTGRESQL, ORACLE, SQLSERVER, MARIADB, POLARDB, POLARDB_POSTGRESQL, POLARDB_ORACLE, OCEANBASEX, DRDS, ADB_MYSQL, ADB_POSTGRESQL, CLICKHOUSE, SELECTDB, STARROCKS, HOLOGRES, MAXCOMPUTE, ES_REST, and TRINO.
	//
	// This parameter is required.
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The environment type. Valid values: product, pre, dev, test, sit, uat, pet, and stag.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Optional. The domain name or IP address of the database. Do not include the protocol, path, or port. You can leave this parameter empty if you select an existing cloud instance and the address is automatically populated.
	//
	// example:
	//
	// mysql.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the cloud resource instance. This parameter is required when you select an existing cloud instance. You can leave this parameter empty when you enter an endpoint.
	//
	// example:
	//
	// rm-bp1234567890
	InstanceResourceId *string `json:"InstanceResourceId,omitempty" xml:"InstanceResourceId,omitempty"`
	// The source of the data source. Valid values:
	//
	// - RDS: An ApsaraDB instance.
	//
	// - PUBLIC_OWN: A self-managed database on the Internet.
	//
	// - ECS_OWN: A self-managed database on an ECS instance.
	//
	// - GATEWAY: A database accessed through Database Gateway.
	//
	// - VPC_IDC: A database accessed through a VPC or data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The name of the credential. Specify this parameter when you use Key Management Service (KMS) to manage database credentials. Leave this parameter empty in other scenarios.
	//
	// example:
	//
	// datasource-secret
	KmsSecret *string `json:"KmsSecret,omitempty" xml:"KmsSecret,omitempty"`
	// Optional. The ID of the data access security policy. This policy is used to control the access and protection rules for the data source.
	//
	// example:
	//
	// sp-usr-example
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Optional. The service port of the database. You can leave this parameter empty if you select an existing cloud instance and the port is automatically populated.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The complete ARN of the Resource Access Management (RAM) role. Specify this parameter when you use a RAM role to access the data source. Leave this parameter empty in other scenarios.
	//
	// example:
	//
	// acs:ram::123456789012****:role/DataSourceRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// Optional. For ADB_POSTGRESQL, DRDS, POSTGRESQL, and POLARDB_POSTGRESQL, enter the database name. For ORACLE, enter the SID or service name, and prefix the service name with a forward slash (/). For POLARDB_ORACLE, enter the SID. For MAXCOMPUTE, enter the project name. Leave this parameter empty for other types.
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
	// Specifies whether to use the cloud instance name as the data source name. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// This parameter takes effect only when you select a cloud database instance.
	//
	// example:
	//
	// 0
	SyncAlias *int32 `json:"SyncAlias,omitempty" xml:"SyncAlias,omitempty"`
	// Specifies whether to use SSL. Valid values:
	//
	// - 1: Enable.
	//
	// - 0: Disable.
	//
	// - -1: Use the default setting.
	//
	// Only ADB_MYSQL, CLICKHOUSE, MARIADB, MYSQL, OCEANBASEX, POLARDB, SELECTDB, and STARROCKS support SSL. Leave this parameter empty for other types.
	//
	// example:
	//
	// 1
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	// The ID of the destination vSwitch. Specify this parameter when you access MaxCompute through a VPC. Leave this parameter empty in other scenarios.
	//
	// example:
	//
	// vsw-bp1234567890
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the destination VPC. Specify this parameter when you access MaxCompute through a VPC. Leave this parameter empty in other scenarios.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the destination zone. Specify this parameter when you access MaxCompute through a VPC. Leave this parameter empty in other scenarios.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddAgenticDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAgenticDataSourceRequest) GoString() string {
	return s.String()
}

func (s *AddAgenticDataSourceRequest) GetAlias() *string {
	return s.Alias
}

func (s *AddAgenticDataSourceRequest) GetAuthAccount() *string {
	return s.AuthAccount
}

func (s *AddAgenticDataSourceRequest) GetAuthPassword() *string {
	return s.AuthPassword
}

func (s *AddAgenticDataSourceRequest) GetAuthType() *int32 {
	return s.AuthType
}

func (s *AddAgenticDataSourceRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *AddAgenticDataSourceRequest) GetCreateMode() *string {
	return s.CreateMode
}

func (s *AddAgenticDataSourceRequest) GetDbType() *string {
	return s.DbType
}

func (s *AddAgenticDataSourceRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *AddAgenticDataSourceRequest) GetHost() *string {
	return s.Host
}

func (s *AddAgenticDataSourceRequest) GetInstanceResourceId() *string {
	return s.InstanceResourceId
}

func (s *AddAgenticDataSourceRequest) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *AddAgenticDataSourceRequest) GetKmsSecret() *string {
	return s.KmsSecret
}

func (s *AddAgenticDataSourceRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *AddAgenticDataSourceRequest) GetPort() *int32 {
	return s.Port
}

func (s *AddAgenticDataSourceRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *AddAgenticDataSourceRequest) GetSchema() *string {
	return s.Schema
}

func (s *AddAgenticDataSourceRequest) GetSnowflakeRole() *string {
	return s.SnowflakeRole
}

func (s *AddAgenticDataSourceRequest) GetSnowflakeWarehouse() *string {
	return s.SnowflakeWarehouse
}

func (s *AddAgenticDataSourceRequest) GetSyncAlias() *int32 {
	return s.SyncAlias
}

func (s *AddAgenticDataSourceRequest) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *AddAgenticDataSourceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AddAgenticDataSourceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AddAgenticDataSourceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddAgenticDataSourceRequest) SetAlias(v string) *AddAgenticDataSourceRequest {
	s.Alias = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetAuthAccount(v string) *AddAgenticDataSourceRequest {
	s.AuthAccount = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetAuthPassword(v string) *AddAgenticDataSourceRequest {
	s.AuthPassword = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetAuthType(v int32) *AddAgenticDataSourceRequest {
	s.AuthType = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetCatalogName(v string) *AddAgenticDataSourceRequest {
	s.CatalogName = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetCreateMode(v string) *AddAgenticDataSourceRequest {
	s.CreateMode = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetDbType(v string) *AddAgenticDataSourceRequest {
	s.DbType = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetEnvType(v string) *AddAgenticDataSourceRequest {
	s.EnvType = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetHost(v string) *AddAgenticDataSourceRequest {
	s.Host = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetInstanceResourceId(v string) *AddAgenticDataSourceRequest {
	s.InstanceResourceId = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetInstanceSource(v string) *AddAgenticDataSourceRequest {
	s.InstanceSource = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetKmsSecret(v string) *AddAgenticDataSourceRequest {
	s.KmsSecret = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetPolicyId(v string) *AddAgenticDataSourceRequest {
	s.PolicyId = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetPort(v int32) *AddAgenticDataSourceRequest {
	s.Port = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetRoleArn(v string) *AddAgenticDataSourceRequest {
	s.RoleArn = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetSchema(v string) *AddAgenticDataSourceRequest {
	s.Schema = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetSnowflakeRole(v string) *AddAgenticDataSourceRequest {
	s.SnowflakeRole = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetSnowflakeWarehouse(v string) *AddAgenticDataSourceRequest {
	s.SnowflakeWarehouse = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetSyncAlias(v int32) *AddAgenticDataSourceRequest {
	s.SyncAlias = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetUseSsl(v int32) *AddAgenticDataSourceRequest {
	s.UseSsl = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetVSwitchId(v string) *AddAgenticDataSourceRequest {
	s.VSwitchId = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetVpcId(v string) *AddAgenticDataSourceRequest {
	s.VpcId = &v
	return s
}

func (s *AddAgenticDataSourceRequest) SetZoneId(v string) *AddAgenticDataSourceRequest {
	s.ZoneId = &v
	return s
}

func (s *AddAgenticDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
