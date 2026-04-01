// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDBNodesShrinkRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *CreateDBNodesShrinkRequest
	GetDBInstanceId() *string
	SetDBNodeShrink(v string) *CreateDBNodesShrinkRequest
	GetDBNodeShrink() *string
	SetOwnerAccount(v string) *CreateDBNodesShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBNodesShrinkRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateDBNodesShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDBNodesShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBNodesShrinkRequest
	GetResourceOwnerId() *int64
}

type CreateDBNodesShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze450g4ctg6t****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The details of the node.
	//
	// This parameter is required.
	DBNodeShrink *string `json:"DBNode,omitempty" xml:"DBNode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDBNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBNodesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBNodesShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBNodesShrinkRequest) GetDBNodeShrink() *string {
	return s.DBNodeShrink
}

func (s *CreateDBNodesShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBNodesShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBNodesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBNodesShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBNodesShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBNodesShrinkRequest) SetClientToken(v string) *CreateDBNodesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBNodesShrinkRequest) SetDBInstanceId(v string) *CreateDBNodesShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBNodesShrinkRequest) SetDBNodeShrink(v string) *CreateDBNodesShrinkRequest {
	s.DBNodeShrink = &v
	return s
}

func (s *CreateDBNodesShrinkRequest) SetOwnerAccount(v string) *CreateDBNodesShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBNodesShrinkRequest) SetOwnerId(v int64) *CreateDBNodesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBNodesShrinkRequest) SetResourceGroupId(v string) *CreateDBNodesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBNodesShrinkRequest) SetResourceOwnerAccount(v string) *CreateDBNodesShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBNodesShrinkRequest) SetResourceOwnerId(v int64) *CreateDBNodesShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
