// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseDirectConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleaseDirectConnectionRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ReleaseDirectConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseDirectConnectionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReleaseDirectConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseDirectConnectionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ReleaseDirectConnectionRequest
	GetSecurityToken() *string
}

type ReleaseDirectConnectionRequest struct {
	// The ID of the instance.
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

func (s ReleaseDirectConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseDirectConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseDirectConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseDirectConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseDirectConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseDirectConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseDirectConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseDirectConnectionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ReleaseDirectConnectionRequest) SetInstanceId(v string) *ReleaseDirectConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) SetOwnerAccount(v string) *ReleaseDirectConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) SetOwnerId(v int64) *ReleaseDirectConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) SetResourceOwnerAccount(v string) *ReleaseDirectConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) SetResourceOwnerId(v int64) *ReleaseDirectConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) SetSecurityToken(v string) *ReleaseDirectConnectionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReleaseDirectConnectionRequest) Validate() error {
	return dara.Validate(s)
}
