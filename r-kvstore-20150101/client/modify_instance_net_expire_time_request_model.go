// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetExpireTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassicExpiredDays(v int32) *ModifyInstanceNetExpireTimeRequest
	GetClassicExpiredDays() *int32
	SetConnectionString(v string) *ModifyInstanceNetExpireTimeRequest
	GetConnectionString() *string
	SetInstanceId(v string) *ModifyInstanceNetExpireTimeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceNetExpireTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceNetExpireTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceNetExpireTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceNetExpireTimeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceNetExpireTimeRequest
	GetSecurityToken() *string
}

type ModifyInstanceNetExpireTimeRequest struct {
	// The extension period to retain the classic network endpoint of the instance. Unit: days. Valid values: **14**, **30**, **60**, and **120**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14
	ClassicExpiredDays *int32 `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
	// The endpoint of the classic network.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****.redis.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
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

func (s ModifyInstanceNetExpireTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetExpireTimeRequest) GetClassicExpiredDays() *int32 {
	return s.ClassicExpiredDays
}

func (s *ModifyInstanceNetExpireTimeRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyInstanceNetExpireTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceNetExpireTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceNetExpireTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceNetExpireTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceNetExpireTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceNetExpireTimeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceNetExpireTimeRequest) SetClassicExpiredDays(v int32) *ModifyInstanceNetExpireTimeRequest {
	s.ClassicExpiredDays = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetConnectionString(v string) *ModifyInstanceNetExpireTimeRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetInstanceId(v string) *ModifyInstanceNetExpireTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetOwnerAccount(v string) *ModifyInstanceNetExpireTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetOwnerId(v int64) *ModifyInstanceNetExpireTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceNetExpireTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetResourceOwnerId(v int64) *ModifyInstanceNetExpireTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) SetSecurityToken(v string) *ModifyInstanceNetExpireTimeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeRequest) Validate() error {
	return dara.Validate(s)
}
