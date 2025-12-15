// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestroyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DestroyInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DestroyInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DestroyInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DestroyInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DestroyInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DestroyInstanceRequest
	GetSecurityToken() *string
}

type DestroyInstanceRequest struct {
	// The ID of the instance in the recycle bin.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-8vb2rhccnvd82f****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DestroyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DestroyInstanceRequest) GoString() string {
	return s.String()
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

func (s *DestroyInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DestroyInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DestroyInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
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

func (s *DestroyInstanceRequest) SetResourceOwnerAccount(v string) *DestroyInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DestroyInstanceRequest) SetResourceOwnerId(v int64) *DestroyInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DestroyInstanceRequest) SetSecurityToken(v string) *DestroyInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DestroyInstanceRequest) Validate() error {
	return dara.Validate(s)
}
