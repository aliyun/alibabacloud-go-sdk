// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgenticDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddAgenticDataSourceResponseBodyData) *AddAgenticDataSourceResponseBody
	GetData() *AddAgenticDataSourceResponseBodyData
	SetErrorCode(v string) *AddAgenticDataSourceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddAgenticDataSourceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddAgenticDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAgenticDataSourceResponseBody
	GetSuccess() *bool
}

type AddAgenticDataSourceResponseBody struct {
	// The details of the created data source.
	Data *AddAgenticDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The request ID, which is used to trace the call.
	//
	// example:
	//
	// request-1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAgenticDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAgenticDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddAgenticDataSourceResponseBody) GetData() *AddAgenticDataSourceResponseBodyData {
	return s.Data
}

func (s *AddAgenticDataSourceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddAgenticDataSourceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddAgenticDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAgenticDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAgenticDataSourceResponseBody) SetData(v *AddAgenticDataSourceResponseBodyData) *AddAgenticDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *AddAgenticDataSourceResponseBody) SetErrorCode(v string) *AddAgenticDataSourceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddAgenticDataSourceResponseBody) SetErrorMessage(v string) *AddAgenticDataSourceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddAgenticDataSourceResponseBody) SetRequestId(v string) *AddAgenticDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAgenticDataSourceResponseBody) SetSuccess(v bool) *AddAgenticDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *AddAgenticDataSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAgenticDataSourceResponseBodyData struct {
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
	// The registration method of the data source. Valid values:
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
	// The time when the data source was created. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2026-10-13 14:30:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the data source was last modified. The time is in the yyyy-MM-dd HH:mm:ss format.
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
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The name of the KMS-managed credential. This parameter is empty if KMS is not used. This is sensitive information.
	//
	// example:
	//
	// datasource-secret
	KmsSecret *string `json:"KmsSecret,omitempty" xml:"KmsSecret,omitempty"`
	// The ID of the data access security policy in use. This parameter is empty if no policy is configured.
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
	// The region where the data source resides.
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
	// The ARN of the RAM role used to access the data source. This parameter is empty if role authorization is not used.
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
	// Indicates whether the cloud instance name is used as the data source name. Valid values:
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
	// - -1: default settings
	//
	// example:
	//
	// 1
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
}

func (s AddAgenticDataSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddAgenticDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddAgenticDataSourceResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *AddAgenticDataSourceResponseBodyData) GetAuthAccount() *string {
	return s.AuthAccount
}

func (s *AddAgenticDataSourceResponseBodyData) GetAuthPassword() *string {
	return s.AuthPassword
}

func (s *AddAgenticDataSourceResponseBodyData) GetAuthType() *int32 {
	return s.AuthType
}

func (s *AddAgenticDataSourceResponseBodyData) GetCatalogName() *string {
	return s.CatalogName
}

func (s *AddAgenticDataSourceResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *AddAgenticDataSourceResponseBodyData) GetDataSourceUuid() *string {
	return s.DataSourceUuid
}

func (s *AddAgenticDataSourceResponseBodyData) GetDbType() *string {
	return s.DbType
}

func (s *AddAgenticDataSourceResponseBodyData) GetEnvType() *string {
	return s.EnvType
}

func (s *AddAgenticDataSourceResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AddAgenticDataSourceResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AddAgenticDataSourceResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *AddAgenticDataSourceResponseBodyData) GetInstanceResourceId() *string {
	return s.InstanceResourceId
}

func (s *AddAgenticDataSourceResponseBodyData) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *AddAgenticDataSourceResponseBodyData) GetKmsSecret() *string {
	return s.KmsSecret
}

func (s *AddAgenticDataSourceResponseBodyData) GetPolicyId() *string {
	return s.PolicyId
}

func (s *AddAgenticDataSourceResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *AddAgenticDataSourceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *AddAgenticDataSourceResponseBodyData) GetResourceAccountId() *string {
	return s.ResourceAccountId
}

func (s *AddAgenticDataSourceResponseBodyData) GetRoleArn() *string {
	return s.RoleArn
}

func (s *AddAgenticDataSourceResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *AddAgenticDataSourceResponseBodyData) GetState() *int32 {
	return s.State
}

func (s *AddAgenticDataSourceResponseBodyData) GetSyncAlias() *int32 {
	return s.SyncAlias
}

func (s *AddAgenticDataSourceResponseBodyData) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *AddAgenticDataSourceResponseBodyData) SetAlias(v string) *AddAgenticDataSourceResponseBodyData {
	s.Alias = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetAuthAccount(v string) *AddAgenticDataSourceResponseBodyData {
	s.AuthAccount = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetAuthPassword(v string) *AddAgenticDataSourceResponseBodyData {
	s.AuthPassword = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetAuthType(v int32) *AddAgenticDataSourceResponseBodyData {
	s.AuthType = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetCatalogName(v string) *AddAgenticDataSourceResponseBodyData {
	s.CatalogName = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetCreateMode(v string) *AddAgenticDataSourceResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetDataSourceUuid(v string) *AddAgenticDataSourceResponseBodyData {
	s.DataSourceUuid = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetDbType(v string) *AddAgenticDataSourceResponseBodyData {
	s.DbType = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetEnvType(v string) *AddAgenticDataSourceResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetGmtCreate(v string) *AddAgenticDataSourceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetGmtModified(v string) *AddAgenticDataSourceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetHost(v string) *AddAgenticDataSourceResponseBodyData {
	s.Host = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetInstanceResourceId(v string) *AddAgenticDataSourceResponseBodyData {
	s.InstanceResourceId = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetInstanceSource(v string) *AddAgenticDataSourceResponseBodyData {
	s.InstanceSource = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetKmsSecret(v string) *AddAgenticDataSourceResponseBodyData {
	s.KmsSecret = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetPolicyId(v string) *AddAgenticDataSourceResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetPort(v int32) *AddAgenticDataSourceResponseBodyData {
	s.Port = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetRegionId(v string) *AddAgenticDataSourceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetResourceAccountId(v string) *AddAgenticDataSourceResponseBodyData {
	s.ResourceAccountId = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetRoleArn(v string) *AddAgenticDataSourceResponseBodyData {
	s.RoleArn = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetSchema(v string) *AddAgenticDataSourceResponseBodyData {
	s.Schema = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetState(v int32) *AddAgenticDataSourceResponseBodyData {
	s.State = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetSyncAlias(v int32) *AddAgenticDataSourceResponseBodyData {
	s.SyncAlias = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) SetUseSsl(v int32) *AddAgenticDataSourceResponseBodyData {
	s.UseSsl = &v
	return s
}

func (s *AddAgenticDataSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
