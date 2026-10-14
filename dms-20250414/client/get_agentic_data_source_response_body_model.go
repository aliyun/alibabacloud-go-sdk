// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAgenticDataSourceResponseBodyData) *GetAgenticDataSourceResponseBody
	GetData() *GetAgenticDataSourceResponseBodyData
	SetErrorCode(v string) *GetAgenticDataSourceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAgenticDataSourceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAgenticDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgenticDataSourceResponseBody
	GetSuccess() *bool
}

type GetAgenticDataSourceResponseBody struct {
	// The data source details. This parameter is empty if the data source does not exist.
	Data *GetAgenticDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the call fails. This parameter is empty if the call is successful.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails. This parameter is empty if the call is successful.
	//
	// example:
	//
	// The parameter is invalid
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// request-1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgenticDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgenticDataSourceResponseBody) GetData() *GetAgenticDataSourceResponseBodyData {
	return s.Data
}

func (s *GetAgenticDataSourceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAgenticDataSourceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAgenticDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgenticDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgenticDataSourceResponseBody) SetData(v *GetAgenticDataSourceResponseBodyData) *GetAgenticDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *GetAgenticDataSourceResponseBody) SetErrorCode(v string) *GetAgenticDataSourceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAgenticDataSourceResponseBody) SetErrorMessage(v string) *GetAgenticDataSourceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAgenticDataSourceResponseBody) SetRequestId(v string) *GetAgenticDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgenticDataSourceResponseBody) SetSuccess(v bool) *GetAgenticDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgenticDataSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgenticDataSourceResponseBodyData struct {
	// The display name of the data source.
	//
	// example:
	//
	// demo_mysql
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The logon username for the data source. This parameter is empty if username authentication is not used.
	//
	// example:
	//
	// db_user
	AuthAccount *string `json:"AuthAccount,omitempty" xml:"AuthAccount,omitempty"`
	// The masked value of the password or token. The actual value is not returned.
	//
	// example:
	//
	// P@ss****
	AuthPassword *string `json:"AuthPassword,omitempty" xml:"AuthPassword,omitempty"`
	// The authentication method. Valid values:
	//
	// - -1: no authentication required
	//
	// - 0: username and password
	//
	// - 1: password only
	//
	// - 2: password-free
	//
	// example:
	//
	// 0
	AuthType *int32 `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The name of the metadata catalog associated with the data source.
	//
	// example:
	//
	// demo_catalog
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The UUID of the metadata catalog associated with the data source.
	//
	// example:
	//
	// mc-ab-0123456789abcdefghijk
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The method used to register the data source. Valid values:
	//
	// - INSTANCE_ID: select an existing cloud instance
	//
	// - HOST_PORT: register by endpoint and port
	//
	// example:
	//
	// HOST_PORT
	CreateMode *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// The unique identifier of the data source.
	//
	// example:
	//
	// cn-hangzhou_a1b2c3d4e5f6
	DataSourceUuid *string `json:"DataSourceUuid,omitempty" xml:"DataSourceUuid,omitempty"`
	// The type of the data source. Valid values: MYSQL, POSTGRESQL, ORACLE, SQLSERVER, MARIADB, POLARDB, POLARDB_POSTGRESQL, POLARDB_ORACLE, OCEANBASEX, DRDS, ADB_MYSQL, ADB_POSTGRESQL, CLICKHOUSE, SELECTDB, STARROCKS, HOLOGRES, MAXCOMPUTE, ES_REST, and TRINO.
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
	// The time when the data source was created. The format is yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-10-13 14:30:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the data source was last modified. The format is yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-10-13 14:35:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The endpoint of the data source. A masked value may be returned in sensitive scenarios.
	//
	// example:
	//
	// mysql.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the cloud resource instance. This parameter may be empty for data sources registered by endpoint.
	//
	// example:
	//
	// rm-bp1234567890
	InstanceResourceId *string `json:"InstanceResourceId,omitempty" xml:"InstanceResourceId,omitempty"`
	// The source of the data source. Valid values:
	//
	// - RDS: ApsaraDB instance
	//
	// - PUBLIC_OWN: self-managed database on the Internet
	//
	// - ECS_OWN: self-managed database on an ECS instance
	//
	// - GATEWAY: connected through Database Gateway
	//
	// - VPC_IDC: connected through a VPC or IDC
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The name of the KMS-managed credential. This parameter is empty if KMS is not used. This field contains sensitive information.
	//
	// example:
	//
	// datasource-secret
	KmsSecret *string `json:"KmsSecret,omitempty" xml:"KmsSecret,omitempty"`
	// The ID of the data access security policy currently in use. This parameter is empty if no policy is configured.
	//
	// example:
	//
	// sp-usr-example
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The connection port of the data source.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region where the data source is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 123456789012****
	ResourceAccountId *string `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// The ARN of the RAM role used to access the data source. This parameter is empty if role-based authorization is not used.
	//
	// example:
	//
	// acs:ram::123456789012****:role/DataSourceRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The name of the connection target. For ADB_POSTGRESQL, DRDS, POSTGRESQL, and POLARDB_POSTGRESQL, this indicates the database name. For ORACLE, this indicates the SID or service name. For POLARDB_ORACLE, this indicates the SID. For MAXCOMPUTE, this indicates the project. For other types, this parameter is usually empty.
	//
	// example:
	//
	// postgres
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The name of the Snowflake role. This parameter takes effect only when DbType is set to SNOWFLAKE.
	//
	// example:
	//
	// WAREHOUSE
	SnowflakeRole *string `json:"SnowflakeRole,omitempty" xml:"SnowflakeRole,omitempty"`
	// The name of the Snowflake compute warehouse. This parameter takes effect only when DbType is set to SNOWFLAKE.
	//
	// example:
	//
	// ROLE
	SnowflakeWarehouse *string `json:"SnowflakeWarehouse,omitempty" xml:"SnowflakeWarehouse,omitempty"`
	// The status of the data source. Valid values:
	//
	// - 0: normal
	//
	// - -1: deleted
	//
	// - -2: disabled
	//
	// example:
	//
	// 0
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// Specifies whether to use the cloud instance name as the data source name. Valid values:
	//
	// - 1: yes
	//
	// - 0: no
	//
	// example:
	//
	// 0
	SyncAlias *int32 `json:"SyncAlias,omitempty" xml:"SyncAlias,omitempty"`
	// The SSL settings. Valid values:
	//
	// - 1: enabled
	//
	// - 0: disabled
	//
	// - -1: use default settings
	//
	// example:
	//
	// 1
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	// The ID of the destination vSwitch used when accessing MaxCompute. This parameter is empty in other scenarios.
	//
	// example:
	//
	// vsw-bp1234567890
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the destination VPC used when accessing MaxCompute. This parameter is empty in other scenarios.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the destination zone used when accessing MaxCompute. This parameter is empty in other scenarios.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetAgenticDataSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgenticDataSourceResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *GetAgenticDataSourceResponseBodyData) GetAuthAccount() *string {
	return s.AuthAccount
}

func (s *GetAgenticDataSourceResponseBodyData) GetAuthPassword() *string {
	return s.AuthPassword
}

func (s *GetAgenticDataSourceResponseBodyData) GetAuthType() *int32 {
	return s.AuthType
}

func (s *GetAgenticDataSourceResponseBodyData) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetAgenticDataSourceResponseBodyData) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *GetAgenticDataSourceResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *GetAgenticDataSourceResponseBodyData) GetDataSourceUuid() *string {
	return s.DataSourceUuid
}

func (s *GetAgenticDataSourceResponseBodyData) GetDbType() *string {
	return s.DbType
}

func (s *GetAgenticDataSourceResponseBodyData) GetEnvType() *string {
	return s.EnvType
}

func (s *GetAgenticDataSourceResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAgenticDataSourceResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAgenticDataSourceResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetAgenticDataSourceResponseBodyData) GetInstanceResourceId() *string {
	return s.InstanceResourceId
}

func (s *GetAgenticDataSourceResponseBodyData) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *GetAgenticDataSourceResponseBodyData) GetKmsSecret() *string {
	return s.KmsSecret
}

func (s *GetAgenticDataSourceResponseBodyData) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetAgenticDataSourceResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *GetAgenticDataSourceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAgenticDataSourceResponseBodyData) GetResourceAccountId() *string {
	return s.ResourceAccountId
}

func (s *GetAgenticDataSourceResponseBodyData) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GetAgenticDataSourceResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *GetAgenticDataSourceResponseBodyData) GetSnowflakeRole() *string {
	return s.SnowflakeRole
}

func (s *GetAgenticDataSourceResponseBodyData) GetSnowflakeWarehouse() *string {
	return s.SnowflakeWarehouse
}

func (s *GetAgenticDataSourceResponseBodyData) GetState() *int32 {
	return s.State
}

func (s *GetAgenticDataSourceResponseBodyData) GetSyncAlias() *int32 {
	return s.SyncAlias
}

func (s *GetAgenticDataSourceResponseBodyData) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *GetAgenticDataSourceResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetAgenticDataSourceResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetAgenticDataSourceResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetAgenticDataSourceResponseBodyData) SetAlias(v string) *GetAgenticDataSourceResponseBodyData {
	s.Alias = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetAuthAccount(v string) *GetAgenticDataSourceResponseBodyData {
	s.AuthAccount = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetAuthPassword(v string) *GetAgenticDataSourceResponseBodyData {
	s.AuthPassword = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetAuthType(v int32) *GetAgenticDataSourceResponseBodyData {
	s.AuthType = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetCatalogName(v string) *GetAgenticDataSourceResponseBodyData {
	s.CatalogName = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetCatalogUuid(v string) *GetAgenticDataSourceResponseBodyData {
	s.CatalogUuid = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetCreateMode(v string) *GetAgenticDataSourceResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetDataSourceUuid(v string) *GetAgenticDataSourceResponseBodyData {
	s.DataSourceUuid = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetDbType(v string) *GetAgenticDataSourceResponseBodyData {
	s.DbType = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetEnvType(v string) *GetAgenticDataSourceResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetGmtCreate(v string) *GetAgenticDataSourceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetGmtModified(v string) *GetAgenticDataSourceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetHost(v string) *GetAgenticDataSourceResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetInstanceResourceId(v string) *GetAgenticDataSourceResponseBodyData {
	s.InstanceResourceId = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetInstanceSource(v string) *GetAgenticDataSourceResponseBodyData {
	s.InstanceSource = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetKmsSecret(v string) *GetAgenticDataSourceResponseBodyData {
	s.KmsSecret = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetPolicyId(v string) *GetAgenticDataSourceResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetPort(v int32) *GetAgenticDataSourceResponseBodyData {
	s.Port = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetRegionId(v string) *GetAgenticDataSourceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetResourceAccountId(v string) *GetAgenticDataSourceResponseBodyData {
	s.ResourceAccountId = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetRoleArn(v string) *GetAgenticDataSourceResponseBodyData {
	s.RoleArn = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetSchema(v string) *GetAgenticDataSourceResponseBodyData {
	s.Schema = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetSnowflakeRole(v string) *GetAgenticDataSourceResponseBodyData {
	s.SnowflakeRole = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetSnowflakeWarehouse(v string) *GetAgenticDataSourceResponseBodyData {
	s.SnowflakeWarehouse = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetState(v int32) *GetAgenticDataSourceResponseBodyData {
	s.State = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetSyncAlias(v int32) *GetAgenticDataSourceResponseBodyData {
	s.SyncAlias = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetUseSsl(v int32) *GetAgenticDataSourceResponseBodyData {
	s.UseSsl = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetVSwitchId(v string) *GetAgenticDataSourceResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetVpcId(v string) *GetAgenticDataSourceResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) SetZoneId(v string) *GetAgenticDataSourceResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *GetAgenticDataSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
