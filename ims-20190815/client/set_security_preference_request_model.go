// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityPreferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceRequest
	GetAllowUserToChangePassword() *bool
	SetAllowUserToLoginWithPasskey(v bool) *SetSecurityPreferenceRequest
	GetAllowUserToLoginWithPasskey() *bool
	SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceRequest
	GetAllowUserToManageAccessKeys() *bool
	SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceRequest
	GetAllowUserToManageMFADevices() *bool
	SetAllowUserToManagePersonalDingTalk(v bool) *SetSecurityPreferenceRequest
	GetAllowUserToManagePersonalDingTalk() *bool
	SetAllowUserToManageServiceCredentials(v bool) *SetSecurityPreferenceRequest
	GetAllowUserToManageServiceCredentials() *bool
	SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceRequest
	GetEnableSaveMFATicket() *bool
	SetLoginNetworkMasks(v string) *SetSecurityPreferenceRequest
	GetLoginNetworkMasks() *string
	SetLoginSessionDuration(v int32) *SetSecurityPreferenceRequest
	GetLoginSessionDuration() *int32
	SetMFAOperationForLogin(v string) *SetSecurityPreferenceRequest
	GetMFAOperationForLogin() *string
	SetMaxIdleDaysForAccessKeys(v int32) *SetSecurityPreferenceRequest
	GetMaxIdleDaysForAccessKeys() *int32
	SetMaxIdleDaysForUsers(v int32) *SetSecurityPreferenceRequest
	GetMaxIdleDaysForUsers() *int32
	SetOperationForRiskLogin(v string) *SetSecurityPreferenceRequest
	GetOperationForRiskLogin() *string
	SetVerificationTypes(v []*string) *SetSecurityPreferenceRequest
	GetVerificationTypes() []*string
}

type SetSecurityPreferenceRequest struct {
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
	VerificationTypes []*string `json:"VerificationTypes,omitempty" xml:"VerificationTypes,omitempty" type:"Repeated"`
}

func (s SetSecurityPreferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceRequest) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceRequest) GetAllowUserToChangePassword() *bool {
	return s.AllowUserToChangePassword
}

func (s *SetSecurityPreferenceRequest) GetAllowUserToLoginWithPasskey() *bool {
	return s.AllowUserToLoginWithPasskey
}

func (s *SetSecurityPreferenceRequest) GetAllowUserToManageAccessKeys() *bool {
	return s.AllowUserToManageAccessKeys
}

func (s *SetSecurityPreferenceRequest) GetAllowUserToManageMFADevices() *bool {
	return s.AllowUserToManageMFADevices
}

func (s *SetSecurityPreferenceRequest) GetAllowUserToManagePersonalDingTalk() *bool {
	return s.AllowUserToManagePersonalDingTalk
}

func (s *SetSecurityPreferenceRequest) GetAllowUserToManageServiceCredentials() *bool {
	return s.AllowUserToManageServiceCredentials
}

func (s *SetSecurityPreferenceRequest) GetEnableSaveMFATicket() *bool {
	return s.EnableSaveMFATicket
}

func (s *SetSecurityPreferenceRequest) GetLoginNetworkMasks() *string {
	return s.LoginNetworkMasks
}

func (s *SetSecurityPreferenceRequest) GetLoginSessionDuration() *int32 {
	return s.LoginSessionDuration
}

func (s *SetSecurityPreferenceRequest) GetMFAOperationForLogin() *string {
	return s.MFAOperationForLogin
}

func (s *SetSecurityPreferenceRequest) GetMaxIdleDaysForAccessKeys() *int32 {
	return s.MaxIdleDaysForAccessKeys
}

func (s *SetSecurityPreferenceRequest) GetMaxIdleDaysForUsers() *int32 {
	return s.MaxIdleDaysForUsers
}

func (s *SetSecurityPreferenceRequest) GetOperationForRiskLogin() *string {
	return s.OperationForRiskLogin
}

func (s *SetSecurityPreferenceRequest) GetVerificationTypes() []*string {
	return s.VerificationTypes
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToLoginWithPasskey(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToLoginWithPasskey = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManagePersonalDingTalk(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManagePersonalDingTalk = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageServiceCredentials(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageServiceCredentials = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceRequest {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetLoginNetworkMasks(v string) *SetSecurityPreferenceRequest {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetLoginSessionDuration(v int32) *SetSecurityPreferenceRequest {
	s.LoginSessionDuration = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetMFAOperationForLogin(v string) *SetSecurityPreferenceRequest {
	s.MFAOperationForLogin = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetMaxIdleDaysForAccessKeys(v int32) *SetSecurityPreferenceRequest {
	s.MaxIdleDaysForAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetMaxIdleDaysForUsers(v int32) *SetSecurityPreferenceRequest {
	s.MaxIdleDaysForUsers = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetOperationForRiskLogin(v string) *SetSecurityPreferenceRequest {
	s.OperationForRiskLogin = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetVerificationTypes(v []*string) *SetSecurityPreferenceRequest {
	s.VerificationTypes = v
	return s
}

func (s *SetSecurityPreferenceRequest) Validate() error {
	return dara.Validate(s)
}
