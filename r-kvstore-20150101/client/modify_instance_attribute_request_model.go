// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceAttributeRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyInstanceAttributeRequest
	GetInstanceName() *string
	SetInstanceReleaseProtection(v bool) *ModifyInstanceAttributeRequest
	GetInstanceReleaseProtection() *bool
	SetNewPassword(v string) *ModifyInstanceAttributeRequest
	GetNewPassword() *string
	SetOwnerAccount(v string) *ModifyInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceAttributeRequest
	GetSecurityToken() *string
}

type ModifyInstanceAttributeRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new name of the instance. The name must be 2 to 80 characters in length. The name must start with a letter and cannot contain spaces and the following special characters: `@ / : = " < > { [ ] }`
	//
	// example:
	//
	// newinstancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// [Specifies whether to enable release protection for the instance.](https://help.aliyun.com/document_detail/165005.html) Valid values:
	//
	// 	- **true**: enables release protection.
	//
	// 	- **false**: disables release protection.
	//
	// >  This parameter is available only for pay-as-you-go instances.
	//
	// example:
	//
	// true
	InstanceReleaseProtection *bool `json:"InstanceReleaseProtection,omitempty" xml:"InstanceReleaseProtection,omitempty"`
	// The new password for the default account. The default account is named after the instance ID. Example: r-bp10noxlhcoim2\\*\\*\\*\\*.
	//
	// > The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. These special characters include `! @ # $ % ^ & 	- ( ) _ + - =`
	//
	// example:
	//
	// uW8+nsrp
	NewPassword          *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAttributeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceAttributeRequest) GetInstanceReleaseProtection() *bool {
	return s.InstanceReleaseProtection
}

func (s *ModifyInstanceAttributeRequest) GetNewPassword() *string {
	return s.NewPassword
}

func (s *ModifyInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceName(v string) *ModifyInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceReleaseProtection(v bool) *ModifyInstanceAttributeRequest {
	s.InstanceReleaseProtection = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetNewPassword(v string) *ModifyInstanceAttributeRequest {
	s.NewPassword = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetOwnerAccount(v string) *ModifyInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetOwnerId(v int64) *ModifyInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetResourceOwnerId(v int64) *ModifyInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetSecurityToken(v string) *ModifyInstanceAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
