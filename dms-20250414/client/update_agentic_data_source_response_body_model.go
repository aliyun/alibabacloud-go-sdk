// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgenticDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateAgenticDataSourceResponseBodyData) *UpdateAgenticDataSourceResponseBody
	GetData() *UpdateAgenticDataSourceResponseBodyData
	SetErrorCode(v string) *UpdateAgenticDataSourceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateAgenticDataSourceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateAgenticDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAgenticDataSourceResponseBody
	GetSuccess() *bool
}

type UpdateAgenticDataSourceResponseBody struct {
	// The updated data source details.
	Data *UpdateAgenticDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the call fails. This parameter is empty if the call is successful.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call fails. This parameter is empty if the call is successful.
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

func (s UpdateAgenticDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgenticDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgenticDataSourceResponseBody) GetData() *UpdateAgenticDataSourceResponseBodyData {
	return s.Data
}

func (s *UpdateAgenticDataSourceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAgenticDataSourceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateAgenticDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgenticDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAgenticDataSourceResponseBody) SetData(v *UpdateAgenticDataSourceResponseBodyData) *UpdateAgenticDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAgenticDataSourceResponseBody) SetErrorCode(v string) *UpdateAgenticDataSourceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBody) SetErrorMessage(v string) *UpdateAgenticDataSourceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBody) SetRequestId(v string) *UpdateAgenticDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBody) SetSuccess(v bool) *UpdateAgenticDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgenticDataSourceResponseBodyData struct {
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
	// The time when the data source was created. The format is yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-10-13 14:30:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the data source was last updated. The format is yyyy-MM-dd HH:mm:ss.
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
	// - -1: default settings
	//
	// example:
	//
	// 1
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
}

func (s UpdateAgenticDataSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgenticDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetAuthAccount() *string {
	return s.AuthAccount
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetAuthPassword() *string {
	return s.AuthPassword
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetAuthType() *int32 {
	return s.AuthType
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetDataSourceUuid() *string {
	return s.DataSourceUuid
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetDbType() *string {
	return s.DbType
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetInstanceResourceId() *string {
	return s.InstanceResourceId
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetKmsSecret() *string {
	return s.KmsSecret
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetResourceAccountId() *string {
	return s.ResourceAccountId
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetState() *int32 {
	return s.State
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetSyncAlias() *int32 {
	return s.SyncAlias
}

func (s *UpdateAgenticDataSourceResponseBodyData) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetAlias(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.Alias = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetAuthAccount(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.AuthAccount = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetAuthPassword(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.AuthPassword = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetAuthType(v int32) *UpdateAgenticDataSourceResponseBodyData {
	s.AuthType = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetCatalogName(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.CatalogName = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetCreateMode(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetDataSourceUuid(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.DataSourceUuid = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetDbType(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.DbType = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetEnvType(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetGmtCreate(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetGmtModified(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetHost(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.Host = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetInstanceResourceId(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.InstanceResourceId = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetInstanceSource(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.InstanceSource = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetKmsSecret(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.KmsSecret = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetPolicyId(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetPort(v int32) *UpdateAgenticDataSourceResponseBodyData {
	s.Port = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetRegionId(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetResourceAccountId(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.ResourceAccountId = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetRoleArn(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.RoleArn = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetSchema(v string) *UpdateAgenticDataSourceResponseBodyData {
	s.Schema = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetState(v int32) *UpdateAgenticDataSourceResponseBodyData {
	s.State = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetSyncAlias(v int32) *UpdateAgenticDataSourceResponseBodyData {
	s.SyncAlias = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) SetUseSsl(v int32) *UpdateAgenticDataSourceResponseBodyData {
	s.UseSsl = &v
	return s
}

func (s *UpdateAgenticDataSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
