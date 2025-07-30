// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlushExpireKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveTime(v string) *FlushExpireKeysRequest
	GetEffectiveTime() *string
	SetInstanceId(v string) *FlushExpireKeysRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *FlushExpireKeysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *FlushExpireKeysRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *FlushExpireKeysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FlushExpireKeysRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *FlushExpireKeysRequest
	GetSecurityToken() *string
}

type FlushExpireKeysRequest struct {
	// The time when you want to delete the expired keys. Default value: Immediately. Valid values:
	//
	// 	- **Immediately**: deletes the keys immediately.
	//
	// 	- **MaintainTime**: deletes the keys during the maintenance window.
	//
	// >  You can call the [ModifyInstanceMaintainTime](https://help.aliyun.com/document_detail/473775.html) operation to modify the maintenance window of an instance.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
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

func (s FlushExpireKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s FlushExpireKeysRequest) GoString() string {
	return s.String()
}

func (s *FlushExpireKeysRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *FlushExpireKeysRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FlushExpireKeysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *FlushExpireKeysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FlushExpireKeysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FlushExpireKeysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FlushExpireKeysRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *FlushExpireKeysRequest) SetEffectiveTime(v string) *FlushExpireKeysRequest {
	s.EffectiveTime = &v
	return s
}

func (s *FlushExpireKeysRequest) SetInstanceId(v string) *FlushExpireKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *FlushExpireKeysRequest) SetOwnerAccount(v string) *FlushExpireKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FlushExpireKeysRequest) SetOwnerId(v int64) *FlushExpireKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *FlushExpireKeysRequest) SetResourceOwnerAccount(v string) *FlushExpireKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlushExpireKeysRequest) SetResourceOwnerId(v int64) *FlushExpireKeysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlushExpireKeysRequest) SetSecurityToken(v string) *FlushExpireKeysRequest {
	s.SecurityToken = &v
	return s
}

func (s *FlushExpireKeysRequest) Validate() error {
	return dara.Validate(s)
}
