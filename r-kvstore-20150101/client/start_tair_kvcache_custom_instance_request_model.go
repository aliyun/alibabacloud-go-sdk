// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTairKVCacheCustomInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartTairKVCacheCustomInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *StartTairKVCacheCustomInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartTairKVCacheCustomInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StartTairKVCacheCustomInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartTairKVCacheCustomInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *StartTairKVCacheCustomInstanceRequest
	GetSecurityToken() *string
}

type StartTairKVCacheCustomInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StartTairKVCacheCustomInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTairKVCacheCustomInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartTairKVCacheCustomInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartTairKVCacheCustomInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartTairKVCacheCustomInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartTairKVCacheCustomInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartTairKVCacheCustomInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartTairKVCacheCustomInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StartTairKVCacheCustomInstanceRequest) SetInstanceId(v string) *StartTairKVCacheCustomInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartTairKVCacheCustomInstanceRequest) SetOwnerAccount(v string) *StartTairKVCacheCustomInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartTairKVCacheCustomInstanceRequest) SetOwnerId(v int64) *StartTairKVCacheCustomInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *StartTairKVCacheCustomInstanceRequest) SetResourceOwnerAccount(v string) *StartTairKVCacheCustomInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartTairKVCacheCustomInstanceRequest) SetResourceOwnerId(v int64) *StartTairKVCacheCustomInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartTairKVCacheCustomInstanceRequest) SetSecurityToken(v string) *StartTairKVCacheCustomInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *StartTairKVCacheCustomInstanceRequest) Validate() error {
	return dara.Validate(s)
}
