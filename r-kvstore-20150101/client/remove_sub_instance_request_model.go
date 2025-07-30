// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSubInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveSubInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RemoveSubInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveSubInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RemoveSubInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveSubInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *RemoveSubInstanceRequest
	GetSecurityToken() *string
}

type RemoveSubInstanceRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RemoveSubInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSubInstanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveSubInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveSubInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveSubInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveSubInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveSubInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveSubInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveSubInstanceRequest) SetInstanceId(v string) *RemoveSubInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveSubInstanceRequest) SetOwnerAccount(v string) *RemoveSubInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveSubInstanceRequest) SetOwnerId(v int64) *RemoveSubInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveSubInstanceRequest) SetResourceOwnerAccount(v string) *RemoveSubInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveSubInstanceRequest) SetResourceOwnerId(v int64) *RemoveSubInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveSubInstanceRequest) SetSecurityToken(v string) *RemoveSubInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveSubInstanceRequest) Validate() error {
	return dara.Validate(s)
}
