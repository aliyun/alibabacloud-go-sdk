// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairSkvDdbTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *CreateTairSkvDdbTableRequest
	GetBackupId() *string
	SetClientToken(v string) *CreateTairSkvDdbTableRequest
	GetClientToken() *string
	SetInstanceType(v string) *CreateTairSkvDdbTableRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *CreateTairSkvDdbTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTairSkvDdbTableRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTairSkvDdbTableRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTairSkvDdbTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTairSkvDdbTableRequest
	GetResourceOwnerId() *int64
	SetSchema(v string) *CreateTairSkvDdbTableRequest
	GetSchema() *string
	SetSecurityToken(v string) *CreateTairSkvDdbTableRequest
	GetSecurityToken() *string
	SetSrcDBInstanceId(v string) *CreateTairSkvDdbTableRequest
	GetSrcDBInstanceId() *string
	SetTableName(v string) *CreateTairSkvDdbTableRequest
	GetTableName() *string
	SetTtlSpec(v string) *CreateTairSkvDdbTableRequest
	GetTtlSpec() *string
	SetWorkspaceId(v string) *CreateTairSkvDdbTableRequest
	GetWorkspaceId() *string
}

type CreateTairSkvDdbTableRequest struct {
	// The cluster backup set ID. Some new cluster architectures support cluster backup set IDs. You can call [DescribeClusterBackupList](https://www.alibabacloud.com/help/en/redis/developer-reference/api-r-kvstore-2015-01-01-describeclusterbackuplist-redis) to obtain the ID.
	//
	// example:
	//
	// cb-hyxdof5x9kqb**
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value. Make sure that the value is unique among different requests. The token is case-sensitive and can contain up to 64 ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz**
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance type. Set the value to tair_skv_ddb_table.
	//
	// This parameter is required.
	//
	// example:
	//
	// tair_skv_ddb_table
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/61012.htm) to query available regions. Use this parameter to specify the region in which to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The table schema configuration in JSON format.
	//
	// example:
	//
	// {"AttributeDefinitions":[{"AttributeType":"S","AttributeName":"pk"},{"AttributeType":"S","AttributeName":"sk"}],"KeySchema":[{"KeyType":"HASH","AttributeName":"pk"},{"KeyType":"RANGE","AttributeName":"sk"}]}
	Schema        *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// To create an instance from a backup set of an existing instance, specify the ID of the source instance in this parameter.
	//
	// > This parameter must be used together with BackupId.
	//
	// example:
	//
	// r-bp1zxszhcgatnx**
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The table name. The name must be 2 to 128 characters in length and must start with an uppercase letter, a lowercase letter, or a Chinese character. The name cannot contain the following characters: @/:="<>{}[] or spaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// apitest
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The parameter settings switch in JSON format.
	//
	// example:
	//
	// {"attributeName":"Expiretime","enabled":true}
	TtlSpec *string `json:"TtlSpec,omitempty" xml:"TtlSpec,omitempty"`
	// The ID of the workspace instance. You can call [DescribeInstances](https://www.alibabacloud.com/help/en/redis/developer-reference/api-r-kvstore-2015-01-01-describeinstances-redis) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tt-bp1zxszhcgatnx**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTairSkvDdbTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTairSkvDdbTableRequest) GoString() string {
	return s.String()
}

func (s *CreateTairSkvDdbTableRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateTairSkvDdbTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTairSkvDdbTableRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateTairSkvDdbTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTairSkvDdbTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTairSkvDdbTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTairSkvDdbTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTairSkvDdbTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTairSkvDdbTableRequest) GetSchema() *string {
	return s.Schema
}

func (s *CreateTairSkvDdbTableRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateTairSkvDdbTableRequest) GetSrcDBInstanceId() *string {
	return s.SrcDBInstanceId
}

func (s *CreateTairSkvDdbTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateTairSkvDdbTableRequest) GetTtlSpec() *string {
	return s.TtlSpec
}

func (s *CreateTairSkvDdbTableRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateTairSkvDdbTableRequest) SetBackupId(v string) *CreateTairSkvDdbTableRequest {
	s.BackupId = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetClientToken(v string) *CreateTairSkvDdbTableRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetInstanceType(v string) *CreateTairSkvDdbTableRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetOwnerAccount(v string) *CreateTairSkvDdbTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetOwnerId(v int64) *CreateTairSkvDdbTableRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetRegionId(v string) *CreateTairSkvDdbTableRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetResourceOwnerAccount(v string) *CreateTairSkvDdbTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetResourceOwnerId(v int64) *CreateTairSkvDdbTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetSchema(v string) *CreateTairSkvDdbTableRequest {
	s.Schema = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetSecurityToken(v string) *CreateTairSkvDdbTableRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetSrcDBInstanceId(v string) *CreateTairSkvDdbTableRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetTableName(v string) *CreateTairSkvDdbTableRequest {
	s.TableName = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetTtlSpec(v string) *CreateTairSkvDdbTableRequest {
	s.TtlSpec = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) SetWorkspaceId(v string) *CreateTairSkvDdbTableRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateTairSkvDdbTableRequest) Validate() error {
	return dara.Validate(s)
}
