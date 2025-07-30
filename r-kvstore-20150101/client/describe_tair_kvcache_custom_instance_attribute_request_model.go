// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheCustomInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTairKVCacheCustomInstanceAttributeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeTairKVCacheCustomInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTairKVCacheCustomInstanceAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeTairKVCacheCustomInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTairKVCacheCustomInstanceAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeTairKVCacheCustomInstanceAttributeRequest
	GetSecurityToken() *string
}

type DescribeTairKVCacheCustomInstanceAttributeRequest struct {
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

func (s DescribeTairKVCacheCustomInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) SetInstanceId(v string) *DescribeTairKVCacheCustomInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) SetOwnerAccount(v string) *DescribeTairKVCacheCustomInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) SetOwnerId(v int64) *DescribeTairKVCacheCustomInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) SetResourceOwnerAccount(v string) *DescribeTairKVCacheCustomInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) SetResourceOwnerId(v int64) *DescribeTairKVCacheCustomInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) SetSecurityToken(v string) *DescribeTairKVCacheCustomInstanceAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
