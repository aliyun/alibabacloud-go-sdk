// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetTairKVCacheCustomInstancePasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ResetTairKVCacheCustomInstancePasswordRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ResetTairKVCacheCustomInstancePasswordRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResetTairKVCacheCustomInstancePasswordRequest
	GetOwnerId() *int64
	SetPassword(v string) *ResetTairKVCacheCustomInstancePasswordRequest
	GetPassword() *string
	SetResourceOwnerAccount(v string) *ResetTairKVCacheCustomInstancePasswordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResetTairKVCacheCustomInstancePasswordRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ResetTairKVCacheCustomInstancePasswordRequest
	GetSecurityToken() *string
	SetSourceBiz(v string) *ResetTairKVCacheCustomInstancePasswordRequest
	GetSourceBiz() *string
}

type ResetTairKVCacheCustomInstancePasswordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Pass!123456
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// SDK
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
}

func (s ResetTairKVCacheCustomInstancePasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetTairKVCacheCustomInstancePasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) SetInstanceId(v string) *ResetTairKVCacheCustomInstancePasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) SetOwnerAccount(v string) *ResetTairKVCacheCustomInstancePasswordRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) SetOwnerId(v int64) *ResetTairKVCacheCustomInstancePasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) SetPassword(v string) *ResetTairKVCacheCustomInstancePasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) SetResourceOwnerAccount(v string) *ResetTairKVCacheCustomInstancePasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) SetResourceOwnerId(v int64) *ResetTairKVCacheCustomInstancePasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) SetSecurityToken(v string) *ResetTairKVCacheCustomInstancePasswordRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) SetSourceBiz(v string) *ResetTairKVCacheCustomInstancePasswordRequest {
	s.SourceBiz = &v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordRequest) Validate() error {
	return dara.Validate(s)
}
