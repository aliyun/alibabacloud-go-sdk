// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlushInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *FlushInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *FlushInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *FlushInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *FlushInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FlushInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *FlushInstanceRequest
	GetSecurityToken() *string
}

type FlushInstanceRequest struct {
	// The ID of the request.
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

func (s FlushInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s FlushInstanceRequest) GoString() string {
	return s.String()
}

func (s *FlushInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FlushInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *FlushInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FlushInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FlushInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FlushInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *FlushInstanceRequest) SetInstanceId(v string) *FlushInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *FlushInstanceRequest) SetOwnerAccount(v string) *FlushInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FlushInstanceRequest) SetOwnerId(v int64) *FlushInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *FlushInstanceRequest) SetResourceOwnerAccount(v string) *FlushInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlushInstanceRequest) SetResourceOwnerId(v int64) *FlushInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlushInstanceRequest) SetSecurityToken(v string) *FlushInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *FlushInstanceRequest) Validate() error {
	return dara.Validate(s)
}
