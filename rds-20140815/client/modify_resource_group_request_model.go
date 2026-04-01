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
	SetDBInstanceId(v string) *ModifyResourceGroupRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyResourceGroupRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyResourceGroupRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *ModifyResourceGroupRequest
	GetResourceType() *string
}

type ModifyResourceGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bpxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. You can call the ListResourceGroups operation to obtain the resource group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acxxxxx
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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

func (s *ModifyResourceGroupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
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

func (s *ModifyResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyResourceGroupRequest) SetClientToken(v string) *ModifyResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetDBInstanceId(v string) *ModifyResourceGroupRequest {
	s.DBInstanceId = &v
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

func (s *ModifyResourceGroupRequest) SetResourceType(v string) *ModifyResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
