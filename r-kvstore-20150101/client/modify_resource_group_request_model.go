// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyResourceGroupRequest
	GetClientToken() *string
	SetInstanceId(v string) *ModifyResourceGroupRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyResourceGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyResourceGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyResourceGroupRequest
	GetResourceOwnerId() *int64
}

type ModifyResourceGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the generated token is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which you want to move the instance.
	//
	// > 	- You can query resource group IDs by using the Tair (Redis OSS-compatible) console or by calling the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation. For more information, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// > 	- Before you modify the resource group to which an instance belongs, you can call the [ListResources](https://help.aliyun.com/document_detail/158866.html) operation to view the resource group of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyResourceGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyResourceGroupRequest) SetClientToken(v string) *ModifyResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetInstanceId(v string) *ModifyResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetOwnerAccount(v string) *ModifyResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetOwnerId(v int64) *ModifyResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetRegionId(v string) *ModifyResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetResourceGroupId(v string) *ModifyResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
