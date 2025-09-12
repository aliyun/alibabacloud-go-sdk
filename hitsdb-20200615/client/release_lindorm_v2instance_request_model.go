// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseLindormV2InstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImmediately(v bool) *ReleaseLindormV2InstanceRequest
	GetImmediately() *bool
	SetInstanceId(v string) *ReleaseLindormV2InstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ReleaseLindormV2InstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseLindormV2InstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReleaseLindormV2InstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseLindormV2InstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ReleaseLindormV2InstanceRequest
	GetSecurityToken() *string
}

type ReleaseLindormV2InstanceRequest struct {
	// example:
	//
	// true
	Immediately *bool `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ReleaseLindormV2InstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseLindormV2InstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseLindormV2InstanceRequest) GetImmediately() *bool {
	return s.Immediately
}

func (s *ReleaseLindormV2InstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseLindormV2InstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseLindormV2InstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseLindormV2InstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseLindormV2InstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseLindormV2InstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ReleaseLindormV2InstanceRequest) SetImmediately(v bool) *ReleaseLindormV2InstanceRequest {
	s.Immediately = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetInstanceId(v string) *ReleaseLindormV2InstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetOwnerAccount(v string) *ReleaseLindormV2InstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetOwnerId(v int64) *ReleaseLindormV2InstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetResourceOwnerAccount(v string) *ReleaseLindormV2InstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetResourceOwnerId(v int64) *ReleaseLindormV2InstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetSecurityToken(v string) *ReleaseLindormV2InstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) Validate() error {
	return dara.Validate(s)
}
