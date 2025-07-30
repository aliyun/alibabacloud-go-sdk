// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCharacterType(v string) *DescribeParameterTemplatesRequest
	GetCharacterType() *string
	SetEngine(v string) *DescribeParameterTemplatesRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeParameterTemplatesRequest
	GetEngineVersion() *string
	SetInstanceId(v string) *DescribeParameterTemplatesRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeParameterTemplatesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterTemplatesRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeParameterTemplatesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeParameterTemplatesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterTemplatesRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeParameterTemplatesRequest
	GetSecurityToken() *string
}

type DescribeParameterTemplatesRequest struct {
	// The architecture of the instance. For more information, see [Overview](https://help.aliyun.com/document_detail/86132.html). Valid values:
	//
	// 	- **logic**: The instance is a cluster master-replica instance or a read/write splitting instance.
	//
	// 	- **normal**: The instance is a standard master-replica instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// logic
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The database engine that is run on the instance. Set the value to **Redis**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The major version of the instance. Valid values: **4.0**, **5.0**, **6.0**, and **7.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the IDs of instances.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the instance belongs. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query the IDs of resource groups.
	//
	// >  You can also query the ID of a resource group in the Resource Management console. For more information, see [View the basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeParameterTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeParameterTemplatesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParameterTemplatesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeParameterTemplatesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterTemplatesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeParameterTemplatesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterTemplatesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterTemplatesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeParameterTemplatesRequest) SetCharacterType(v string) *DescribeParameterTemplatesRequest {
	s.CharacterType = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetEngine(v string) *DescribeParameterTemplatesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetEngineVersion(v string) *DescribeParameterTemplatesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetInstanceId(v string) *DescribeParameterTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceGroupId(v string) *DescribeParameterTemplatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetSecurityToken(v string) *DescribeParameterTemplatesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
