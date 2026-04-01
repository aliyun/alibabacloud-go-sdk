// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDBNodesShrinkRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DeleteDBNodesShrinkRequest
	GetDBInstanceId() *string
	SetDBNodeIdShrink(v string) *DeleteDBNodesShrinkRequest
	GetDBNodeIdShrink() *string
	SetOwnerAccount(v string) *DeleteDBNodesShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBNodesShrinkRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DeleteDBNodesShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteDBNodesShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBNodesShrinkRequest
	GetResourceOwnerId() *int64
}

type DeleteDBNodesShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests.
	//
	// The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The node IDs.
	//
	// This parameter is required.
	DBNodeIdShrink *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group. You can call the DescribeDBInstanceAttribute operation to obtain the ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBNodesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDBNodesShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBNodesShrinkRequest) GetDBNodeIdShrink() *string {
	return s.DBNodeIdShrink
}

func (s *DeleteDBNodesShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBNodesShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBNodesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteDBNodesShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBNodesShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBNodesShrinkRequest) SetClientToken(v string) *DeleteDBNodesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBNodesShrinkRequest) SetDBInstanceId(v string) *DeleteDBNodesShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBNodesShrinkRequest) SetDBNodeIdShrink(v string) *DeleteDBNodesShrinkRequest {
	s.DBNodeIdShrink = &v
	return s
}

func (s *DeleteDBNodesShrinkRequest) SetOwnerAccount(v string) *DeleteDBNodesShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBNodesShrinkRequest) SetOwnerId(v int64) *DeleteDBNodesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBNodesShrinkRequest) SetResourceGroupId(v string) *DeleteDBNodesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteDBNodesShrinkRequest) SetResourceOwnerAccount(v string) *DeleteDBNodesShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBNodesShrinkRequest) SetResourceOwnerId(v int64) *DeleteDBNodesShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
