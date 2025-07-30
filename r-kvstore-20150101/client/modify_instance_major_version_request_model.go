// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMajorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveTime(v string) *ModifyInstanceMajorVersionRequest
	GetEffectiveTime() *string
	SetInstanceId(v string) *ModifyInstanceMajorVersionRequest
	GetInstanceId() *string
	SetMajorVersion(v string) *ModifyInstanceMajorVersionRequest
	GetMajorVersion() *string
	SetOwnerAccount(v string) *ModifyInstanceMajorVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceMajorVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceMajorVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceMajorVersionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceMajorVersionRequest
	GetSecurityToken() *string
}

type ModifyInstanceMajorVersionRequest struct {
	// The time when you want to upgrade the major version. Valid values:
	//
	// 	- **Immediately*	- (default): immediately upgrades the major version.
	//
	// 	- **MaintainTime**: upgrades the major version within the maintenance window.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The major version to which you want to upgrade the instance. Valid values: **4.0*	- and **5.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.0
	MajorVersion         *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyInstanceMajorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMajorVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMajorVersionRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyInstanceMajorVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceMajorVersionRequest) GetMajorVersion() *string {
	return s.MajorVersion
}

func (s *ModifyInstanceMajorVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceMajorVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceMajorVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceMajorVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceMajorVersionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceMajorVersionRequest) SetEffectiveTime(v string) *ModifyInstanceMajorVersionRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetInstanceId(v string) *ModifyInstanceMajorVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetMajorVersion(v string) *ModifyInstanceMajorVersionRequest {
	s.MajorVersion = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetOwnerAccount(v string) *ModifyInstanceMajorVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetOwnerId(v int64) *ModifyInstanceMajorVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetResourceOwnerAccount(v string) *ModifyInstanceMajorVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetResourceOwnerId(v int64) *ModifyInstanceMajorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) SetSecurityToken(v string) *ModifyInstanceMajorVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceMajorVersionRequest) Validate() error {
	return dara.Validate(s)
}
