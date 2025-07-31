// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestroyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DestroyInstanceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DestroyInstanceRequest
	GetDBInstanceId() *string
	SetInstanceId(v string) *DestroyInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DestroyInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DestroyInstanceRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DestroyInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DestroyInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DestroyInstanceRequest
	GetResourceOwnerId() *int64
}

type DestroyInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// > **InstanceId*	- and **DBInstanceId*	- serve the same function. You need only to specify one of them.
	//
	// example:
	//
	// dds-bp147acd4783****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance ID.
	//
	// > **InstanceId*	- and **DBInstanceId*	- serve the same function. You need only to specify one of them.
	//
	// example:
	//
	// dds-bp147acd4783****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DestroyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DestroyInstanceRequest) GoString() string {
	return s.String()
}

func (s *DestroyInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DestroyInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DestroyInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DestroyInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DestroyInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DestroyInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DestroyInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DestroyInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DestroyInstanceRequest) SetClientToken(v string) *DestroyInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DestroyInstanceRequest) SetDBInstanceId(v string) *DestroyInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DestroyInstanceRequest) SetInstanceId(v string) *DestroyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DestroyInstanceRequest) SetOwnerAccount(v string) *DestroyInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DestroyInstanceRequest) SetOwnerId(v int64) *DestroyInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DestroyInstanceRequest) SetResourceGroupId(v string) *DestroyInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DestroyInstanceRequest) SetResourceOwnerAccount(v string) *DestroyInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DestroyInstanceRequest) SetResourceOwnerId(v int64) *DestroyInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DestroyInstanceRequest) Validate() error {
	return dara.Validate(s)
}
