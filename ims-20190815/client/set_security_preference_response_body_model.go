// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSecurityPreferenceResponseBody
	GetRequestId() *string
	SetSecurityPreference(v *SetSecurityPreferenceResponseBodySecurityPreference) *SetSecurityPreferenceResponseBody
	GetSecurityPreference() *SetSecurityPreferenceResponseBodySecurityPreference
}

type SetSecurityPreferenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 17494710-B4BA-4185-BBBB-C1A6ABDE1639
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security preferences.
	SecurityPreference *SetSecurityPreferenceResponseBodySecurityPreference `json:"SecurityPreference,omitempty" xml:"SecurityPreference,omitempty" type:"Struct"`
}

func (s SetSecurityPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSecurityPreferenceResponseBody) GetSecurityPreference() *SetSecurityPreferenceResponseBodySecurityPreference {
	return s.SecurityPreference
}

func (s *SetSecurityPreferenceResponseBody) SetRequestId(v string) *SetSecurityPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSecurityPreferenceResponseBody) SetSecurityPreference(v *SetSecurityPreferenceResponseBodySecurityPreference) *SetSecurityPreferenceResponseBody {
	s.SecurityPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBody) Validate() error {
	if s.SecurityPreference != nil {
		if err := s.SecurityPreference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetSecurityPreferenceResponseBodySecurityPreference struct {
	// The AccessKey preferences.
	AccessKeyPreference *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	// The login preferences.
	LoginProfilePreference *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	// The MFA (multi-factor authentication) preferences.
	MFAPreference *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	// Settings for the maximum idle period in days.
	MaxIdleDays *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays `json:"MaxIdleDays,omitempty" xml:"MaxIdleDays,omitempty" type:"Struct"`
	// The personal information preferences.
	PersonalInfoPreference *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference `json:"PersonalInfoPreference,omitempty" xml:"PersonalInfoPreference,omitempty" type:"Struct"`
	// The preferences for MFA methods.
	VerificationPreference *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference `json:"VerificationPreference,omitempty" xml:"VerificationPreference,omitempty" type:"Struct"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) GetAccessKeyPreference() *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	return s.AccessKeyPreference
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) GetLoginProfilePreference() *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	return s.LoginProfilePreference
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) GetMFAPreference() *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	return s.MFAPreference
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) GetMaxIdleDays() *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays {
	return s.MaxIdleDays
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) GetPersonalInfoPreference() *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference {
	return s.PersonalInfoPreference
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) GetVerificationPreference() *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference {
	return s.VerificationPreference
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetAccessKeyPreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.AccessKeyPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetLoginProfilePreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.LoginProfilePreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetMFAPreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.MFAPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetMaxIdleDays(v *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.MaxIdleDays = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetPersonalInfoPreference(v *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.PersonalInfoPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetVerificationPreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.VerificationPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) Validate() error {
	if s.AccessKeyPreference != nil {
		if err := s.AccessKeyPreference.Validate(); err != nil {
			return err
		}
	}
	if s.LoginProfilePreference != nil {
		if err := s.LoginProfilePreference.Validate(); err != nil {
			return err
		}
	}
	if s.MFAPreference != nil {
		if err := s.MFAPreference.Validate(); err != nil {
			return err
		}
	}
	if s.MaxIdleDays != nil {
		if err := s.MaxIdleDays.Validate(); err != nil {
			return err
		}
	}
	if s.PersonalInfoPreference != nil {
		if err := s.PersonalInfoPreference.Validate(); err != nil {
			return err
		}
	}
	if s.VerificationPreference != nil {
		if err := s.VerificationPreference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference struct {
	// Specifies whether RAM users can manage their own AccessKeys.
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
	// Specifies whether RAM users can manage their own API keys. Valid values:
	//
	// - true: RAM users can manage their own API keys.
	//
	// - false: RAM users cannot manage their own API keys.
	//
	// example:
	//
	// false
	AllowUserToManageServiceCredentials *bool `json:"AllowUserToManageServiceCredentials,omitempty" xml:"AllowUserToManageServiceCredentials,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GetAllowUserToManageAccessKeys() *bool {
	return s.AllowUserToManageAccessKeys
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GetAllowUserToManageServiceCredentials() *bool {
	return s.AllowUserToManageServiceCredentials
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) SetAllowUserToManageServiceCredentials(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageServiceCredentials = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) Validate() error {
	return dara.Validate(s)
}

type SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference struct {
	// Specifies whether RAM users can manage their own passwords.
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Specifies whether RAM users can sign in with a passkey.
	//
	// example:
	//
	// false
	AllowUserToLoginWithPasskey *bool `json:"AllowUserToLoginWithPasskey,omitempty" xml:"AllowUserToLoginWithPasskey,omitempty"`
	// Specifies whether to save the verification status for seven days after a RAM user completes multi-factor authentication (MFA) during sign-in.
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The login network mask.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The duration of the login session for a RAM user, in hours.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	// The MFA policy for user sign-in. This parameter is the recommended replacement for `EnforceMFAForLogin`, which is still supported.
	//
	// example:
	//
	// adaptive
	MFAOperationForLogin *string `json:"MFAOperationForLogin,omitempty" xml:"MFAOperationForLogin,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// example:
	//
	// autonomous
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetAllowUserToChangePassword() *bool {
	return s.AllowUserToChangePassword
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetAllowUserToLoginWithPasskey() *bool {
	return s.AllowUserToLoginWithPasskey
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetEnableSaveMFATicket() *bool {
	return s.EnableSaveMFATicket
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetLoginNetworkMasks() *string {
	return s.LoginNetworkMasks
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetLoginSessionDuration() *int32 {
	return s.LoginSessionDuration
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetMFAOperationForLogin() *string {
	return s.MFAOperationForLogin
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetOperationForRiskLogin() *string {
	return s.OperationForRiskLogin
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToLoginWithPasskey(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToLoginWithPasskey = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginNetworkMasks(v string) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginSessionDuration(v int32) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginSessionDuration = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetMFAOperationForLogin(v string) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.MFAOperationForLogin = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetOperationForRiskLogin(v string) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.OperationForRiskLogin = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) Validate() error {
	return dara.Validate(s)
}

type SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference struct {
	// Specifies whether RAM users can manage their own MFA devices.
	//
	// example:
	//
	// false
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GetAllowUserToManageMFADevices() *bool {
	return s.AllowUserToManageMFADevices
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) Validate() error {
	return dara.Validate(s)
}

type SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays struct {
	// The maximum idle period, in days, for an AccessKey of a RAM user. If an AccessKey is not used within the specified period, the system automatically disables it the next day.
	//
	// Default value: 730.
	//
	// example:
	//
	// 730
	MaxIdleDaysForAccessKeys *int32 `json:"MaxIdleDaysForAccessKeys,omitempty" xml:"MaxIdleDaysForAccessKeys,omitempty"`
	// The maximum idle period, in days, for a RAM user. If a RAM user with console sign-in enabled does not sign in within this period, the system automatically disables their console sign-in the next day. This setting does not apply to single sign-on (SSO).
	//
	// Default value: 730.
	//
	// example:
	//
	// 730
	MaxIdleDaysForUsers *int32 `json:"MaxIdleDaysForUsers,omitempty" xml:"MaxIdleDaysForUsers,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) GetMaxIdleDaysForAccessKeys() *int32 {
	return s.MaxIdleDaysForAccessKeys
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) GetMaxIdleDaysForUsers() *int32 {
	return s.MaxIdleDaysForUsers
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) SetMaxIdleDaysForAccessKeys(v int32) *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays {
	s.MaxIdleDaysForAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) SetMaxIdleDaysForUsers(v int32) *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays {
	s.MaxIdleDaysForUsers = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) Validate() error {
	return dara.Validate(s)
}

type SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference struct {
	// Specifies whether RAM users can attach or detach their personal DingTalk accounts.
	//
	// example:
	//
	// true
	AllowUserToManagePersonalDingTalk *bool `json:"AllowUserToManagePersonalDingTalk,omitempty" xml:"AllowUserToManagePersonalDingTalk,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) GetAllowUserToManagePersonalDingTalk() *bool {
	return s.AllowUserToManagePersonalDingTalk
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) SetAllowUserToManagePersonalDingTalk(v bool) *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference {
	s.AllowUserToManagePersonalDingTalk = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) Validate() error {
	return dara.Validate(s)
}

type SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference struct {
	// The allowed MFA methods.
	VerificationTypes []*string `json:"VerificationTypes,omitempty" xml:"VerificationTypes,omitempty" type:"Repeated"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) GetVerificationTypes() []*string {
	return s.VerificationTypes
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) SetVerificationTypes(v []*string) *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference {
	s.VerificationTypes = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) Validate() error {
	return dara.Validate(s)
}
