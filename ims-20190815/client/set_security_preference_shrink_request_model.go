// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityPreferenceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceShrinkRequest
	GetAllowUserToChangePassword() *bool
	SetAllowUserToLoginWithPasskey(v bool) *SetSecurityPreferenceShrinkRequest
	GetAllowUserToLoginWithPasskey() *bool
	SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceShrinkRequest
	GetAllowUserToManageAccessKeys() *bool
	SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceShrinkRequest
	GetAllowUserToManageMFADevices() *bool
	SetAllowUserToManagePersonalDingTalk(v bool) *SetSecurityPreferenceShrinkRequest
	GetAllowUserToManagePersonalDingTalk() *bool
	SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceShrinkRequest
	GetEnableSaveMFATicket() *bool
	SetLoginNetworkMasks(v string) *SetSecurityPreferenceShrinkRequest
	GetLoginNetworkMasks() *string
	SetLoginSessionDuration(v int32) *SetSecurityPreferenceShrinkRequest
	GetLoginSessionDuration() *int32
	SetMFAOperationForLogin(v string) *SetSecurityPreferenceShrinkRequest
	GetMFAOperationForLogin() *string
	SetOperationForRiskLogin(v string) *SetSecurityPreferenceShrinkRequest
	GetOperationForRiskLogin() *string
	SetVerificationTypesShrink(v string) *SetSecurityPreferenceShrinkRequest
	GetVerificationTypesShrink() *string
}

type SetSecurityPreferenceShrinkRequest struct {
	// Specifies whether RAM users can change their passwords. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Specifies whether a RAM user can use a passkey for logon. Valid values:
	//
	// 	- true: A RAM user can use a passkey for logon. This is the default value.
	//
	// 	- false: A RAM user cannot use a passkey for logon.
	//
	// example:
	//
	// true
	AllowUserToLoginWithPasskey *bool `json:"AllowUserToLoginWithPasskey,omitempty" xml:"AllowUserToLoginWithPasskey,omitempty"`
	// Specifies whether RAM users can manage their AccessKey pairs. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
	// Specifies whether RAM users can manage their MFA devices. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
	// Specifies whether RAM users can manage their personal DingTalk accounts, such as binding and unbinding of the accounts. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManagePersonalDingTalk *bool `json:"AllowUserToManagePersonalDingTalk,omitempty" xml:"AllowUserToManagePersonalDingTalk,omitempty"`
	// Specifies whether RAM users can remember the MFA devices for seven days. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). This parameter does not take effect on API calls that are authenticated by using AccessKey pairs.
	//
	// 	- If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	//
	// 	- If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	//
	// If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	//
	// You can specify up to 40 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users.
	//
	// Valid values: 1 to 24. Unit: hours.
	//
	// Default value: 6.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	// Specifies whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console. This parameter is used to replace EnforceMFAForLogin. EnforceMFAForLogin is still valid. However, we recommend that you use MFAOperationForLogin. Valid values:
	//
	// 	- mandatory: MFA is required for all RAM users. If you use EnforceMFAForLogin, set the value to true.
	//
	// 	- independent (default): User-specific settings are applied. If you use EnforceMFAForLogin, set the value to false.
	//
	// 	- adaptive: MFA is required only for RAM users who initiated unusual logons.
	//
	// example:
	//
	// adaptive
	MFAOperationForLogin *string `json:"MFAOperationForLogin,omitempty" xml:"MFAOperationForLogin,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable MFA for RAM users who initiated unusual logons. Valid values:
	//
	// 	- autonomous (default): yes. MFA is prompted for RAM users who initiated unusual logons. However, the RAM users are allowed to skip MFA.
	//
	// 	- enforceVerify: MFA is prompted for RAM users who initiated unusual logons and the RAM users cannot skip MFA.
	//
	// example:
	//
	// autonomous
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
	// The MFA methods.
	VerificationTypesShrink *string `json:"VerificationTypes,omitempty" xml:"VerificationTypes,omitempty"`
}

func (s SetSecurityPreferenceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceShrinkRequest) GetAllowUserToChangePassword() *bool {
	return s.AllowUserToChangePassword
}

func (s *SetSecurityPreferenceShrinkRequest) GetAllowUserToLoginWithPasskey() *bool {
	return s.AllowUserToLoginWithPasskey
}

func (s *SetSecurityPreferenceShrinkRequest) GetAllowUserToManageAccessKeys() *bool {
	return s.AllowUserToManageAccessKeys
}

func (s *SetSecurityPreferenceShrinkRequest) GetAllowUserToManageMFADevices() *bool {
	return s.AllowUserToManageMFADevices
}

func (s *SetSecurityPreferenceShrinkRequest) GetAllowUserToManagePersonalDingTalk() *bool {
	return s.AllowUserToManagePersonalDingTalk
}

func (s *SetSecurityPreferenceShrinkRequest) GetEnableSaveMFATicket() *bool {
	return s.EnableSaveMFATicket
}

func (s *SetSecurityPreferenceShrinkRequest) GetLoginNetworkMasks() *string {
	return s.LoginNetworkMasks
}

func (s *SetSecurityPreferenceShrinkRequest) GetLoginSessionDuration() *int32 {
	return s.LoginSessionDuration
}

func (s *SetSecurityPreferenceShrinkRequest) GetMFAOperationForLogin() *string {
	return s.MFAOperationForLogin
}

func (s *SetSecurityPreferenceShrinkRequest) GetOperationForRiskLogin() *string {
	return s.OperationForRiskLogin
}

func (s *SetSecurityPreferenceShrinkRequest) GetVerificationTypesShrink() *string {
	return s.VerificationTypesShrink
}

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToLoginWithPasskey(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToLoginWithPasskey = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToManagePersonalDingTalk(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToManagePersonalDingTalk = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceShrinkRequest {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetLoginNetworkMasks(v string) *SetSecurityPreferenceShrinkRequest {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetLoginSessionDuration(v int32) *SetSecurityPreferenceShrinkRequest {
	s.LoginSessionDuration = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetMFAOperationForLogin(v string) *SetSecurityPreferenceShrinkRequest {
	s.MFAOperationForLogin = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetOperationForRiskLogin(v string) *SetSecurityPreferenceShrinkRequest {
	s.OperationForRiskLogin = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetVerificationTypesShrink(v string) *SetSecurityPreferenceShrinkRequest {
	s.VerificationTypesShrink = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
