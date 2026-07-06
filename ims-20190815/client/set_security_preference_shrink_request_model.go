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
	SetAllowUserToManageServiceCredentials(v bool) *SetSecurityPreferenceShrinkRequest
	GetAllowUserToManageServiceCredentials() *bool
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
	// Specifies whether RAM users can change their own passwords. Valid values:
	//
	// - true (default): Allowed.
	//
	// - false: Not allowed.
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Specifies whether RAM users can use passkeys to log on to the console. Valid values:
	//
	// - true (default): Allowed.
	//
	// - false: Not allowed.
	//
	// example:
	//
	// true
	AllowUserToLoginWithPasskey *bool `json:"AllowUserToLoginWithPasskey,omitempty" xml:"AllowUserToLoginWithPasskey,omitempty"`
	// Specifies whether RAM users can manage their own AccessKeys. Valid values:
	//
	// - true: Allowed.
	//
	// - false (default): Not allowed.
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
	// Specifies whether RAM users can manage their own MFA devices. Valid values:
	//
	// - true (default): Allowed.
	//
	// - false: Not allowed.
	//
	// example:
	//
	// true
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
	// Specifies whether RAM users can link or unlink their personal DingTalk accounts. Valid values:
	//
	// - true (default): Allowed.
	//
	// - false: Not allowed.
	//
	// example:
	//
	// true
	AllowUserToManagePersonalDingTalk *bool `json:"AllowUserToManagePersonalDingTalk,omitempty" xml:"AllowUserToManagePersonalDingTalk,omitempty"`
	// Specifies whether RAM users can manage their own API keys. Valid values:
	//
	// - true: Allowed.
	//
	// - false: Not allowed.
	//
	// example:
	//
	// false
	AllowUserToManageServiceCredentials *bool `json:"AllowUserToManageServiceCredentials,omitempty" xml:"AllowUserToManageServiceCredentials,omitempty"`
	// Specifies whether a RAM user who logs on with multi-factor authentication (MFA) can skip MFA for the next seven days. Valid values:
	//
	// - true: Allowed.
	//
	// - false (default): Not allowed.
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The IP address mask that is used to log on to the console. This mask applies to password-based logons and single sign-on (SSO) logons, but does not affect API calls that are initiated by using an AccessKey pair.
	//
	// - If you specify a mask, RAM users can log on to the console only from the specified IP addresses.
	//
	// - If you do not specify a mask, RAM users can log on to the console from all IP addresses.
	//
	// If you need to specify multiple masks, separate them with semicolons (`;`). Example: `192.168.0.0/16;10.0.0.0/8`.
	//
	// You can specify up to 40 masks. The total length cannot exceed 512 characters.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The session duration of a RAM user who logs on to the console. Unit: hours.
	//
	// Valid values: 1 to 24.
	//
	// Default value: 6.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	// Specifies the MFA policy for user logon. This parameter replaces `EnforceMFAForLogin`. We recommend that you use this parameter. `EnforceMFAForLogin` is still valid. Valid values:
	//
	// - `mandatory`: enforces MFA for all RAM users. This is equivalent to setting `EnforceMFAForLogin` to `true`.
	//
	// - `independent` (default): The MFA settings for each RAM user are not affected. This is equivalent to setting `EnforceMFAForLogin` to `false`.
	//
	// - `adaptive`: enforces MFA only for unusual logons.
	//
	// example:
	//
	// adaptive
	MFAOperationForLogin *string `json:"MFAOperationForLogin,omitempty" xml:"MFAOperationForLogin,omitempty"`
	// The maximum idle period of the AccessKey pairs of RAM users. An AccessKey pair that is not used for the specified period of time is automatically disabled on the next day. You can set the value to one of the following numbers:
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
	// The maximum idle period of RAM users. If a RAM user who can log on to the console does not log on to the console for the specified period of time (SSO logons are not included), the console logon feature of the RAM user is disabled on the next day. You can set the value to one of the following numbers:
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

func (s *SetSecurityPreferenceShrinkRequest) GetAllowUserToManageServiceCredentials() *bool {
	return s.AllowUserToManageServiceCredentials
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

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToManageServiceCredentials(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToManageServiceCredentials = &v
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
