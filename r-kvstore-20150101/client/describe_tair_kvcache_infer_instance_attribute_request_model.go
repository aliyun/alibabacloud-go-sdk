// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheInferInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTairKVCacheInferInstanceAttributeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeTairKVCacheInferInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTairKVCacheInferInstanceAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeTairKVCacheInferInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTairKVCacheInferInstanceAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeTairKVCacheInferInstanceAttributeRequest
	GetSecurityToken() *string
}

type DescribeTairKVCacheInferInstanceAttributeRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeTairKVCacheInferInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) SetInstanceId(v string) *DescribeTairKVCacheInferInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) SetOwnerAccount(v string) *DescribeTairKVCacheInferInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) SetOwnerId(v int64) *DescribeTairKVCacheInferInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) SetResourceOwnerAccount(v string) *DescribeTairKVCacheInferInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) SetResourceOwnerId(v int64) *DescribeTairKVCacheInferInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) SetSecurityToken(v string) *DescribeTairKVCacheInferInstanceAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
