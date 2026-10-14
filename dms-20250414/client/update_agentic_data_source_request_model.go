// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgenticDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *UpdateAgenticDataSourceRequest
	GetAlias() *string
	SetAuthAccount(v string) *UpdateAgenticDataSourceRequest
	GetAuthAccount() *string
	SetAuthPassword(v string) *UpdateAgenticDataSourceRequest
	GetAuthPassword() *string
	SetAuthType(v int32) *UpdateAgenticDataSourceRequest
	GetAuthType() *int32
	SetCatalogName(v string) *UpdateAgenticDataSourceRequest
	GetCatalogName() *string
	SetCreateMode(v string) *UpdateAgenticDataSourceRequest
	GetCreateMode() *string
	SetDataSourceUuid(v string) *UpdateAgenticDataSourceRequest
	GetDataSourceUuid() *string
	SetDbType(v string) *UpdateAgenticDataSourceRequest
	GetDbType() *string
	SetEnvType(v string) *UpdateAgenticDataSourceRequest
	GetEnvType() *string
	SetHost(v string) *UpdateAgenticDataSourceRequest
	GetHost() *string
	SetInstanceResourceId(v string) *UpdateAgenticDataSourceRequest
	GetInstanceResourceId() *string
	SetKmsSecret(v string) *UpdateAgenticDataSourceRequest
	GetKmsSecret() *string
	SetPolicyId(v string) *UpdateAgenticDataSourceRequest
	GetPolicyId() *string
	SetPort(v int32) *UpdateAgenticDataSourceRequest
	GetPort() *int32
	SetRoleArn(v string) *UpdateAgenticDataSourceRequest
	GetRoleArn() *string
	SetSchema(v string) *UpdateAgenticDataSourceRequest
	GetSchema() *string
	SetSnowflakeRole(v string) *UpdateAgenticDataSourceRequest
	GetSnowflakeRole() *string
	SetSnowflakeWarehouse(v string) *UpdateAgenticDataSourceRequest
	GetSnowflakeWarehouse() *string
	SetSyncAlias(v int32) *UpdateAgenticDataSourceRequest
	GetSyncAlias() *int32
	SetUseSsl(v int32) *UpdateAgenticDataSourceRequest
	GetUseSsl() *int32
}

type UpdateAgenticDataSourceRequest struct {
	// The new display name of the data source. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// demo_mysql
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The new logon username. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// db_user
	AuthAccount *string `json:"AuthAccount,omitempty" xml:"AuthAccount,omitempty"`
	// The new logon password. If this parameter is not specified, the value is not updated. This is sensitive information. Do not record or disclose it.
	//
	// example:
	//
	// P@ss****
	AuthPassword *string `json:"AuthPassword,omitempty" xml:"AuthPassword,omitempty"`
	// Optional. The authentication type. Valid values: -1: No authentication required. 0: Username and password. 1: Password only. 2: Password-free. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// 0
	AuthType *int32 `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The new metadata catalog name. If this parameter is not specified, the value is not updated. The name can be up to 128 characters in length. It must start with a Chinese character, a letter, or an underscore, and can contain digits and hyphens. It cannot contain only digits.
	//
	// example:
	//
	// demo_catalog
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The registration method. Valid values: INSTANCE_ID: Select an existing cloud instance. HOST_PORT: Register by endpoint and port. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// HOST_PORT
	CreateMode *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// The UUID of the data source to be updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou_a1b2c3d4e5f6
	DataSourceUuid *string `json:"DataSourceUuid,omitempty" xml:"DataSourceUuid,omitempty"`
	// Modifying the data source type is not supported. Do not specify this parameter.
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The environment type. Valid values: product, pre, dev, test, sit, uat, pet, and stag. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Optional. The new database domain name or IP address. Do not include the protocol, path, or port. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// mysql.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The new cloud resource instance ID. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// rm-bp1234567890
	InstanceResourceId *string `json:"InstanceResourceId,omitempty" xml:"InstanceResourceId,omitempty"`
	// The new KMS managed credential name. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// datasource-secret
	KmsSecret *string `json:"KmsSecret,omitempty" xml:"KmsSecret,omitempty"`
	// Optional. The new data access security policy ID. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// sp-usr-example
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Optional. The new database service port. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The new Alibaba Cloud RAM role ARN. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// acs:ram::123456789012****:role/DataSourceRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// Optional. For ADB_POSTGRESQL, DRDS, POSTGRESQL, and POLARDB_POSTGRESQL, specify the database name. For ORACLE, specify the SID or service name, and prefix the service name with a forward slash (/). For POLARDB_ORACLE, specify the SID. For MAXCOMPUTE, specify the project name. For other types, leave this parameter empty. If this parameter is not specified, the value is not updated.
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
	// Specifies whether to use the cloud instance name as the data source name. Valid values: 1: Yes. 0: No. If this parameter is not specified, the value is not updated.
	//
	// example:
	//
	// 0
	SyncAlias *int32 `json:"SyncAlias,omitempty" xml:"SyncAlias,omitempty"`
	// The new SSL setting. Valid values: 1: Enable. 0: Disable. -1: Use the default setting. If this parameter is not specified, the value is not updated. SSL is supported only for ADB_MYSQL, CLICKHOUSE, MARIADB, MYSQL, OCEANBASEX, POLARDB, SELECTDB, and STARROCKS.
	//
	// example:
	//
	// 1
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
}

func (s UpdateAgenticDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgenticDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgenticDataSourceRequest) GetAlias() *string {
	return s.Alias
}

func (s *UpdateAgenticDataSourceRequest) GetAuthAccount() *string {
	return s.AuthAccount
}

func (s *UpdateAgenticDataSourceRequest) GetAuthPassword() *string {
	return s.AuthPassword
}

func (s *UpdateAgenticDataSourceRequest) GetAuthType() *int32 {
	return s.AuthType
}

func (s *UpdateAgenticDataSourceRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateAgenticDataSourceRequest) GetCreateMode() *string {
	return s.CreateMode
}

func (s *UpdateAgenticDataSourceRequest) GetDataSourceUuid() *string {
	return s.DataSourceUuid
}

func (s *UpdateAgenticDataSourceRequest) GetDbType() *string {
	return s.DbType
}

func (s *UpdateAgenticDataSourceRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateAgenticDataSourceRequest) GetHost() *string {
	return s.Host
}

func (s *UpdateAgenticDataSourceRequest) GetInstanceResourceId() *string {
	return s.InstanceResourceId
}

func (s *UpdateAgenticDataSourceRequest) GetKmsSecret() *string {
	return s.KmsSecret
}

func (s *UpdateAgenticDataSourceRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateAgenticDataSourceRequest) GetPort() *int32 {
	return s.Port
}

func (s *UpdateAgenticDataSourceRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateAgenticDataSourceRequest) GetSchema() *string {
	return s.Schema
}

func (s *UpdateAgenticDataSourceRequest) GetSnowflakeRole() *string {
	return s.SnowflakeRole
}

func (s *UpdateAgenticDataSourceRequest) GetSnowflakeWarehouse() *string {
	return s.SnowflakeWarehouse
}

func (s *UpdateAgenticDataSourceRequest) GetSyncAlias() *int32 {
	return s.SyncAlias
}

func (s *UpdateAgenticDataSourceRequest) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *UpdateAgenticDataSourceRequest) SetAlias(v string) *UpdateAgenticDataSourceRequest {
	s.Alias = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetAuthAccount(v string) *UpdateAgenticDataSourceRequest {
	s.AuthAccount = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetAuthPassword(v string) *UpdateAgenticDataSourceRequest {
	s.AuthPassword = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetAuthType(v int32) *UpdateAgenticDataSourceRequest {
	s.AuthType = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetCatalogName(v string) *UpdateAgenticDataSourceRequest {
	s.CatalogName = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetCreateMode(v string) *UpdateAgenticDataSourceRequest {
	s.CreateMode = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetDataSourceUuid(v string) *UpdateAgenticDataSourceRequest {
	s.DataSourceUuid = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetDbType(v string) *UpdateAgenticDataSourceRequest {
	s.DbType = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetEnvType(v string) *UpdateAgenticDataSourceRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetHost(v string) *UpdateAgenticDataSourceRequest {
	s.Host = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetInstanceResourceId(v string) *UpdateAgenticDataSourceRequest {
	s.InstanceResourceId = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetKmsSecret(v string) *UpdateAgenticDataSourceRequest {
	s.KmsSecret = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetPolicyId(v string) *UpdateAgenticDataSourceRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetPort(v int32) *UpdateAgenticDataSourceRequest {
	s.Port = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetRoleArn(v string) *UpdateAgenticDataSourceRequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetSchema(v string) *UpdateAgenticDataSourceRequest {
	s.Schema = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetSnowflakeRole(v string) *UpdateAgenticDataSourceRequest {
	s.SnowflakeRole = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetSnowflakeWarehouse(v string) *UpdateAgenticDataSourceRequest {
	s.SnowflakeWarehouse = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetSyncAlias(v int32) *UpdateAgenticDataSourceRequest {
	s.SyncAlias = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) SetUseSsl(v int32) *UpdateAgenticDataSourceRequest {
	s.UseSsl = &v
	return s
}

func (s *UpdateAgenticDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
