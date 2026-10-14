// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAgenticDataSourcesResponseBodyData) *ListAgenticDataSourcesResponseBody
	GetData() []*ListAgenticDataSourcesResponseBodyData
	SetErrorCode(v string) *ListAgenticDataSourcesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAgenticDataSourcesResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListAgenticDataSourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgenticDataSourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgenticDataSourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgenticDataSourcesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListAgenticDataSourcesResponseBody
	GetTotalCount() *int64
}

type ListAgenticDataSourcesResponseBody struct {
	// The data sources on the current page.
	Data []*ListAgenticDataSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// The actual number of entries returned per page in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The cursor for the next page. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// page-2-token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of data sources that meet the filter conditions. This parameter may be empty in some scenarios.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgenticDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgenticDataSourcesResponseBody) GetData() []*ListAgenticDataSourcesResponseBodyData {
	return s.Data
}

func (s *ListAgenticDataSourcesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAgenticDataSourcesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAgenticDataSourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgenticDataSourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgenticDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgenticDataSourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgenticDataSourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAgenticDataSourcesResponseBody) SetData(v []*ListAgenticDataSourcesResponseBodyData) *ListAgenticDataSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListAgenticDataSourcesResponseBody) SetErrorCode(v string) *ListAgenticDataSourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBody) SetErrorMessage(v string) *ListAgenticDataSourcesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBody) SetMaxResults(v int32) *ListAgenticDataSourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBody) SetNextToken(v string) *ListAgenticDataSourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBody) SetRequestId(v string) *ListAgenticDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBody) SetSuccess(v bool) *ListAgenticDataSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBody) SetTotalCount(v int64) *ListAgenticDataSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgenticDataSourcesResponseBodyData struct {
	// The display name of the data source.
	//
	// example:
	//
	// demo_mysql
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The logon username of the data source. This parameter is empty if username authentication is not used.
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
	// The unique identifier of the data source.
	//
	// example:
	//
	// cn-hangzhou_a1b2c3d4e5f6
	DataSourceUuid *string `json:"DataSourceUuid,omitempty" xml:"DataSourceUuid,omitempty"`
	// The data source type. Valid values: MYSQL, POSTGRESQL, ORACLE, SQLSERVER, MARIADB, POLARDB, POLARDB_POSTGRESQL, POLARDB_ORACLE, OCEANBASEX, DRDS, ADB_MYSQL, ADB_POSTGRESQL, CLICKHOUSE, SELECTDB, STARROCKS, HOLOGRES, MAXCOMPUTE, ES_REST, and TRINO.
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
	// - RDS: ApsaraDB instance
	//
	// - PUBLIC_OWN: self-managed database on the public network
	//
	// - ECS_OWN: self-managed database on an ECS instance
	//
	// - GATEWAY: connected through Database Gateway
	//
	// - VPC_IDC: connected over a VPC or IDC
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The name of the KMS managed credential. This parameter is empty if KMS is not used. This is sensitive information.
	//
	// example:
	//
	// datasource-secret
	KmsSecret *string `json:"KmsSecret,omitempty" xml:"KmsSecret,omitempty"`
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
	// The connection target name. For ADB_POSTGRESQL, DRDS, POSTGRESQL, and POLARDB_POSTGRESQL, this parameter indicates the database name. For ORACLE, this parameter indicates the SID or service name. For POLARDB_ORACLE, this parameter indicates the SID. For MAXCOMPUTE, this parameter indicates the project. For other types, this parameter is usually empty.
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

func (s ListAgenticDataSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticDataSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgenticDataSourcesResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *ListAgenticDataSourcesResponseBodyData) GetAuthAccount() *string {
	return s.AuthAccount
}

func (s *ListAgenticDataSourcesResponseBodyData) GetAuthPassword() *string {
	return s.AuthPassword
}

func (s *ListAgenticDataSourcesResponseBodyData) GetAuthType() *int32 {
	return s.AuthType
}

func (s *ListAgenticDataSourcesResponseBodyData) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListAgenticDataSourcesResponseBodyData) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *ListAgenticDataSourcesResponseBodyData) GetDataSourceUuid() *string {
	return s.DataSourceUuid
}

func (s *ListAgenticDataSourcesResponseBodyData) GetDbType() *string {
	return s.DbType
}

func (s *ListAgenticDataSourcesResponseBodyData) GetEnvType() *string {
	return s.EnvType
}

func (s *ListAgenticDataSourcesResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAgenticDataSourcesResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAgenticDataSourcesResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *ListAgenticDataSourcesResponseBodyData) GetInstanceResourceId() *string {
	return s.InstanceResourceId
}

func (s *ListAgenticDataSourcesResponseBodyData) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *ListAgenticDataSourcesResponseBodyData) GetKmsSecret() *string {
	return s.KmsSecret
}

func (s *ListAgenticDataSourcesResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *ListAgenticDataSourcesResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAgenticDataSourcesResponseBodyData) GetResourceAccountId() *string {
	return s.ResourceAccountId
}

func (s *ListAgenticDataSourcesResponseBodyData) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListAgenticDataSourcesResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *ListAgenticDataSourcesResponseBodyData) GetState() *int32 {
	return s.State
}

func (s *ListAgenticDataSourcesResponseBodyData) GetSyncAlias() *int32 {
	return s.SyncAlias
}

func (s *ListAgenticDataSourcesResponseBodyData) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *ListAgenticDataSourcesResponseBodyData) SetAlias(v string) *ListAgenticDataSourcesResponseBodyData {
	s.Alias = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetAuthAccount(v string) *ListAgenticDataSourcesResponseBodyData {
	s.AuthAccount = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetAuthPassword(v string) *ListAgenticDataSourcesResponseBodyData {
	s.AuthPassword = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetAuthType(v int32) *ListAgenticDataSourcesResponseBodyData {
	s.AuthType = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetCatalogName(v string) *ListAgenticDataSourcesResponseBodyData {
	s.CatalogName = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetCatalogUuid(v string) *ListAgenticDataSourcesResponseBodyData {
	s.CatalogUuid = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetDataSourceUuid(v string) *ListAgenticDataSourcesResponseBodyData {
	s.DataSourceUuid = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetDbType(v string) *ListAgenticDataSourcesResponseBodyData {
	s.DbType = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetEnvType(v string) *ListAgenticDataSourcesResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetGmtCreate(v string) *ListAgenticDataSourcesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetGmtModified(v string) *ListAgenticDataSourcesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetHost(v string) *ListAgenticDataSourcesResponseBodyData {
	s.Host = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetInstanceResourceId(v string) *ListAgenticDataSourcesResponseBodyData {
	s.InstanceResourceId = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetInstanceSource(v string) *ListAgenticDataSourcesResponseBodyData {
	s.InstanceSource = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetKmsSecret(v string) *ListAgenticDataSourcesResponseBodyData {
	s.KmsSecret = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetPort(v int32) *ListAgenticDataSourcesResponseBodyData {
	s.Port = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetRegionId(v string) *ListAgenticDataSourcesResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetResourceAccountId(v string) *ListAgenticDataSourcesResponseBodyData {
	s.ResourceAccountId = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetRoleArn(v string) *ListAgenticDataSourcesResponseBodyData {
	s.RoleArn = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetSchema(v string) *ListAgenticDataSourcesResponseBodyData {
	s.Schema = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetState(v int32) *ListAgenticDataSourcesResponseBodyData {
	s.State = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetSyncAlias(v int32) *ListAgenticDataSourcesResponseBodyData {
	s.SyncAlias = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) SetUseSsl(v int32) *ListAgenticDataSourcesResponseBodyData {
	s.UseSsl = &v
	return s
}

func (s *ListAgenticDataSourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
