// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalDistributeCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveTime(v string) *CreateGlobalDistributeCacheRequest
	GetEffectiveTime() *string
	SetOwnerAccount(v string) *CreateGlobalDistributeCacheRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateGlobalDistributeCacheRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateGlobalDistributeCacheRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateGlobalDistributeCacheRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateGlobalDistributeCacheRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateGlobalDistributeCacheRequest
	GetSecurityToken() *string
	SetSeedSubInstanceId(v string) *CreateGlobalDistributeCacheRequest
	GetSeedSubInstanceId() *string
}

type CreateGlobalDistributeCacheRequest struct {
	// Specifies when to perform the operation. Valid values:
	//
	// - **Immediately**: Performs the operation immediately.
	//
	// - **MaintainTime**: Performs the operation during the maintenance window. This is the default value.
	//
	// > You can call the [ModifyInstanceMaintainTime](https://help.aliyun.com/document_detail/473775.html) operation to change the maintenance window of the instance.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// > This is a system parameter. You do not need to specify it.
	//
	// example:
	//
	// -
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	SeedSubInstanceId *string `json:"SeedSubInstanceId,omitempty" xml:"SeedSubInstanceId,omitempty"`
}

func (s CreateGlobalDistributeCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalDistributeCacheRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalDistributeCacheRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *CreateGlobalDistributeCacheRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateGlobalDistributeCacheRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateGlobalDistributeCacheRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateGlobalDistributeCacheRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateGlobalDistributeCacheRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateGlobalDistributeCacheRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateGlobalDistributeCacheRequest) GetSeedSubInstanceId() *string {
	return s.SeedSubInstanceId
}

func (s *CreateGlobalDistributeCacheRequest) SetEffectiveTime(v string) *CreateGlobalDistributeCacheRequest {
	s.EffectiveTime = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetOwnerAccount(v string) *CreateGlobalDistributeCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetOwnerId(v int64) *CreateGlobalDistributeCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetResourceGroupId(v string) *CreateGlobalDistributeCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetResourceOwnerAccount(v string) *CreateGlobalDistributeCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetResourceOwnerId(v int64) *CreateGlobalDistributeCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetSecurityToken(v string) *CreateGlobalDistributeCacheRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) SetSeedSubInstanceId(v string) *CreateGlobalDistributeCacheRequest {
	s.SeedSubInstanceId = &v
	return s
}

func (s *CreateGlobalDistributeCacheRequest) Validate() error {
	return dara.Validate(s)
}
