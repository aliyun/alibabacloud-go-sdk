// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMinorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveTime(v string) *ModifyInstanceMinorVersionRequest
	GetEffectiveTime() *string
	SetInstanceId(v string) *ModifyInstanceMinorVersionRequest
	GetInstanceId() *string
	SetMinorversion(v string) *ModifyInstanceMinorVersionRequest
	GetMinorversion() *string
	SetOwnerAccount(v string) *ModifyInstanceMinorVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceMinorVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceMinorVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceMinorVersionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceMinorVersionRequest
	GetSecurityToken() *string
}

type ModifyInstanceMinorVersionRequest struct {
	// The time when you want to update the minor version. Valid values:
	//
	// 	- **Immediately*	- (default): immediately updates the minor version.
	//
	// 	- **MaintainTime**: updates the minor version during the maintenance window.
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
	// The minor version to which you want to update. Default value: **latest_version**.
	//
	// example:
	//
	// latest_version
	Minorversion         *string `json:"Minorversion,omitempty" xml:"Minorversion,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyInstanceMinorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMinorVersionRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyInstanceMinorVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceMinorVersionRequest) GetMinorversion() *string {
	return s.Minorversion
}

func (s *ModifyInstanceMinorVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceMinorVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceMinorVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceMinorVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceMinorVersionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceMinorVersionRequest) SetEffectiveTime(v string) *ModifyInstanceMinorVersionRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetInstanceId(v string) *ModifyInstanceMinorVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetMinorversion(v string) *ModifyInstanceMinorVersionRequest {
	s.Minorversion = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetOwnerAccount(v string) *ModifyInstanceMinorVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetOwnerId(v int64) *ModifyInstanceMinorVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetResourceOwnerAccount(v string) *ModifyInstanceMinorVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetResourceOwnerId(v int64) *ModifyInstanceMinorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) SetSecurityToken(v string) *ModifyInstanceMinorVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceMinorVersionRequest) Validate() error {
	return dara.Validate(s)
}
