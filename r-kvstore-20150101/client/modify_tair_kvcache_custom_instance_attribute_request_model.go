// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTairKVCacheCustomInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest
	GetInstanceName() *string
	SetOwnerAccount(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyTairKVCacheCustomInstanceAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyTairKVCacheCustomInstanceAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest
	GetSecurityToken() *string
	SetSourceBiz(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest
	GetSourceBiz() *string
}

type ModifyTairKVCacheCustomInstanceAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// redistest
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// SDK
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
}

func (s ModifyTairKVCacheCustomInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTairKVCacheCustomInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) SetInstanceId(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) SetInstanceName(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) SetOwnerAccount(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) SetOwnerId(v int64) *ModifyTairKVCacheCustomInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) SetResourceOwnerId(v int64) *ModifyTairKVCacheCustomInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) SetSecurityToken(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) SetSourceBiz(v string) *ModifyTairKVCacheCustomInstanceAttributeRequest {
	s.SourceBiz = &v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
