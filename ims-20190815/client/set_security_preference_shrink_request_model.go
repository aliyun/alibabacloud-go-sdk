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
	SetMaxIdleDaysForAccessKeys(v int32) *SetSecurityPreferenceShrinkRequest
	GetMaxIdleDaysForAccessKeys() *int32
	SetMaxIdleDaysForUsers(v int32) *SetSecurityPreferenceShrinkRequest
	GetMaxIdleDaysForUsers() *int32
	SetOperationForRiskLogin(v string) *SetSecurityPreferenceShrinkRequest
	GetOperationForRiskLogin() *string
	SetVerificationTypesShrink(v string) *SetSecurityPreferenceShrinkRequest
	GetVerificationTypesShrink() *string
}

type SetSecurityPreferenceShrinkRequest struct {
	// Specifies whether RAM users can manage their own passwords. Valid values:
	//
	// - true (default)
	//
	// - false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Specifies whether RAM users can log on using passkeys. Valid values:
	//
	// - true (default)
	//
	// - false
	//
	// example:
	//
	// true
	AllowUserToLoginWithPasskey *bool `json:"AllowUserToLoginWithPasskey,omitempty" xml:"AllowUserToLoginWithPasskey,omitempty"`
	// Specifies whether RAM users can manage their own AccessKey pairs. Valid values:
	//
	// - true:
	//
	// - false (default)
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
	// Specifies whether RAM users can manage their own MFA devices. Valid values:
	//
	// - true (default)
	//
	// - false
	//
	// example:
	//
	// true
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
	// Specifies whether RAM users can attach or detach their personal DingTalk accounts. Valid values:
	//
	// - true (default)
	//
	// - false
	//
	// example:
	//
	// true
	AllowUserToManagePersonalDingTalk *bool `json:"AllowUserToManagePersonalDingTalk,omitempty" xml:"AllowUserToManagePersonalDingTalk,omitempty"`
	// Specifies whether to save the multi-factor authentication (MFA) status for seven days after a RAM user logs on using MFA. Valid values:
	//
	// - true
	//
	// - false (default)
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// Specifies the IP addresses or CIDR blocks from which RAM users are allowed to sign in to the Alibaba Cloud console.
	//
	// - This restriction applies to both password-based and single sign-on (SSO) logons.
	//
	// - It does not affect API calls made with an AccessKey pair.
	//
	// - If a mask is not configured, logons are allowed from any IP address.
	//
	// Separate multiple entries with a semicolon (`;`). For example: 192.168.0.0/16;10.0.0.0/8.
	//
	// The mask is limited to a maximum of 40 entries and a total length of 512 characters.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The duration of a logon session for a RAM user.
	//
	// Valid values: 1 to 24. Unit: hours.
	//
	// Default value: 6.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	// Specifies whether MFA is required for logon. This parameter replaces `EnforceMFAForLogin`. The `EnforceMFAForLogin` parameter is still valid, but using this new parameter is recommended. Valid values:
	//
	// - mandatory: Enforces MFA for all RAM users. This value corresponds to `true` for the `EnforceMFAForLogin` parameter.
	//
	// - independent (default): The MFA requirement depends on the configuration of each RAM user. This value corresponds to `false` for the `EnforceMFAForLogin` parameter.
	//
	// - adaptive: Enforces MFA only for abnormal logons.
	//
	// example:
	//
	// adaptive
	MFAOperationForLogin *string `json:"MFAOperationForLogin,omitempty" xml:"MFAOperationForLogin,omitempty"`
	// The maximum number of days that a RAM user\\"s AccessKey pair can be idle. If an AccessKey pair is not used within the specified period, it is automatically disabled the next day. Valid values:
	//
	// - 90
	//
	// - 180
	//
	// - 365
	//
	// - 730 (default)
	//
	// example:
	//
	// 365
	MaxIdleDaysForAccessKeys *int32 `json:"MaxIdleDaysForAccessKeys,omitempty" xml:"MaxIdleDaysForAccessKeys,omitempty"`
	// The maximum number of days that a RAM user can be idle. If a RAM user with console logon enabled does not log on within this period, their console logon is automatically disabled the next day. This setting does not apply to single sign-on (SSO) logons. Valid values:
	//
	// - 90
	//
	// - 180
	//
	// - 365
	//
	// - 730 (default)
	//
	// example:
	//
	// 365
	MaxIdleDaysForUsers *int32 `json:"MaxIdleDaysForUsers,omitempty" xml:"MaxIdleDaysForUsers,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
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

func (s *SetSecurityPreferenceShrinkRequest) GetMaxIdleDaysForAccessKeys() *int32 {
	return s.MaxIdleDaysForAccessKeys
}

func (s *SetSecurityPreferenceShrinkRequest) GetMaxIdleDaysForUsers() *int32 {
	return s.MaxIdleDaysForUsers
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

func (s *SetSecurityPreferenceShrinkRequest) SetMaxIdleDaysForAccessKeys(v int32) *SetSecurityPreferenceShrinkRequest {
	s.MaxIdleDaysForAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetMaxIdleDaysForUsers(v int32) *SetSecurityPreferenceShrinkRequest {
	s.MaxIdleDaysForUsers = &v
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
