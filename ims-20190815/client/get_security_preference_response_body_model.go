// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSecurityPreferenceResponseBody
	GetRequestId() *string
	SetSecurityPreference(v *GetSecurityPreferenceResponseBodySecurityPreference) *GetSecurityPreferenceResponseBody
	GetSecurityPreference() *GetSecurityPreferenceResponseBodySecurityPreference
}

type GetSecurityPreferenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30C9068D-FBAA-4998-9986-8A562FED0BC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about security preferences.
	SecurityPreference *GetSecurityPreferenceResponseBodySecurityPreference `json:"SecurityPreference,omitempty" xml:"SecurityPreference,omitempty" type:"Struct"`
}

func (s GetSecurityPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityPreferenceResponseBody) GetSecurityPreference() *GetSecurityPreferenceResponseBodySecurityPreference {
	return s.SecurityPreference
}

func (s *GetSecurityPreferenceResponseBody) SetRequestId(v string) *GetSecurityPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityPreferenceResponseBody) SetSecurityPreference(v *GetSecurityPreferenceResponseBodySecurityPreference) *GetSecurityPreferenceResponseBody {
	s.SecurityPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBody) Validate() error {
	if s.SecurityPreference != nil {
		if err := s.SecurityPreference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityPreferenceResponseBodySecurityPreference struct {
	// The AccessKey preference.
	AccessKeyPreference *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	// The logon preferences.
	LoginProfilePreference *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	// The MFA preferences.
	MFAPreference *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	// The configuration of the maximum idle period.
	MaxIdleDays *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays `json:"MaxIdleDays,omitempty" xml:"MaxIdleDays,omitempty" type:"Struct"`
	// The personal information preferences.
	PersonalInfoPreference *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference `json:"PersonalInfoPreference,omitempty" xml:"PersonalInfoPreference,omitempty" type:"Struct"`
	// The preferences for MFA methods.
	VerificationPreference *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference `json:"VerificationPreference,omitempty" xml:"VerificationPreference,omitempty" type:"Struct"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) GetAccessKeyPreference() *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	return s.AccessKeyPreference
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) GetLoginProfilePreference() *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	return s.LoginProfilePreference
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) GetMFAPreference() *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	return s.MFAPreference
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) GetMaxIdleDays() *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays {
	return s.MaxIdleDays
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) GetPersonalInfoPreference() *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference {
	return s.PersonalInfoPreference
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) GetVerificationPreference() *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference {
	return s.VerificationPreference
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetAccessKeyPreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.AccessKeyPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetLoginProfilePreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.LoginProfilePreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetMFAPreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.MFAPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetMaxIdleDays(v *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.MaxIdleDays = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetPersonalInfoPreference(v *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.PersonalInfoPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetVerificationPreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.VerificationPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) Validate() error {
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

type GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference struct {
	// Indicates whether RAM users can manage their own AccessKey pairs. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GetAllowUserToManageAccessKeys() *bool {
	return s.AllowUserToManageAccessKeys
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) SetAllowUserToManageAccessKeys(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) Validate() error {
	return dara.Validate(s)
}

type GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference struct {
	// Indicates whether RAM users can manage their own passwords. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Indicates whether RAM users can log on using passkeys. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	AllowUserToLoginWithPasskey *bool `json:"AllowUserToLoginWithPasskey,omitempty" xml:"AllowUserToLoginWithPasskey,omitempty"`
	// Indicates whether to save the multi-factor authentication (MFA) status for seven days after a RAM user logs on using MFA. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The the IP addresses or CIDR blocks from which RAM users are allowed to sign in to the Alibaba Cloud console.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The duration of a logon session for a RAM user. Unit: hours.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	// Indicates whether MFA is required for logon. This parameter replaces `EnforceMFAForLogin`. The `EnforceMFAForLogin` parameter is still valid, but using this new parameter is recommended. Valid values:
	//
	// - mandatory: MFA is required for all RAM users. This value corresponds to `true` for the `EnforceMFAForLogin` parameter.
	//
	// - independent (default): The MFA configuration of each RAM user is used. This value corresponds to `false` for the `EnforceMFAForLogin` parameter.
	//
	// - adaptive: MFA is required only for abnormal logons.
	//
	// example:
	//
	// adaptive
	MFAOperationForLogin *string `json:"MFAOperationForLogin,omitempty" xml:"MFAOperationForLogin,omitempty"`
	// Deprecated
	//
	// Indicates whether to use MFA for secondary authentication during an abnormal logon. Valid values:
	//
	// - autonomous (default): The secondary authentication can be skipped. The attachment of an MFA device is not required.
	//
	// - enforceVerify: The secondary authentication is required.
	//
	// example:
	//
	// autonomous
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetAllowUserToChangePassword() *bool {
	return s.AllowUserToChangePassword
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetAllowUserToLoginWithPasskey() *bool {
	return s.AllowUserToLoginWithPasskey
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetEnableSaveMFATicket() *bool {
	return s.EnableSaveMFATicket
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetLoginNetworkMasks() *string {
	return s.LoginNetworkMasks
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetLoginSessionDuration() *int32 {
	return s.LoginSessionDuration
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetMFAOperationForLogin() *string {
	return s.MFAOperationForLogin
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GetOperationForRiskLogin() *string {
	return s.OperationForRiskLogin
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToChangePassword(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToLoginWithPasskey(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToLoginWithPasskey = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetEnableSaveMFATicket(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginNetworkMasks(v string) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginSessionDuration(v int32) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginSessionDuration = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetMFAOperationForLogin(v string) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.MFAOperationForLogin = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetOperationForRiskLogin(v string) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.OperationForRiskLogin = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) Validate() error {
	return dara.Validate(s)
}

type GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference struct {
	// Indicates whether RAM users can manage their own MFA devices. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GetAllowUserToManageMFADevices() *bool {
	return s.AllowUserToManageMFADevices
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) SetAllowUserToManageMFADevices(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) Validate() error {
	return dara.Validate(s)
}

type GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays struct {
	// The maximum idle period for the AccessKey pair of a RAM user. If an AccessKey pair remains unused for this period, it is automatically disabled the next day.
	//
	// Default value: 730.
	//
	// Unit: days.
	//
	// example:
	//
	// 730
	MaxIdleDaysForAccessKeys *int32 `json:"MaxIdleDaysForAccessKeys,omitempty" xml:"MaxIdleDaysForAccessKeys,omitempty"`
	// The maximum idle period for RAM users. If a RAM user with console logon enabled remains inactive for this period, their console logon is automatically disabled the next day. This does not apply to single sign-on (SSO) logons.
	//
	// Default value: 730.
	//
	// Unit: days.
	//
	// example:
	//
	// 730
	MaxIdleDaysForUsers *int32 `json:"MaxIdleDaysForUsers,omitempty" xml:"MaxIdleDaysForUsers,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) GetMaxIdleDaysForAccessKeys() *int32 {
	return s.MaxIdleDaysForAccessKeys
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) GetMaxIdleDaysForUsers() *int32 {
	return s.MaxIdleDaysForUsers
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) SetMaxIdleDaysForAccessKeys(v int32) *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays {
	s.MaxIdleDaysForAccessKeys = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) SetMaxIdleDaysForUsers(v int32) *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays {
	s.MaxIdleDaysForUsers = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) Validate() error {
	return dara.Validate(s)
}

type GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference struct {
	// Indicates whether RAM users can attach or detach their personal DingTalk accounts. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	AllowUserToManagePersonalDingTalk *bool `json:"AllowUserToManagePersonalDingTalk,omitempty" xml:"AllowUserToManagePersonalDingTalk,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) GetAllowUserToManagePersonalDingTalk() *bool {
	return s.AllowUserToManagePersonalDingTalk
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) SetAllowUserToManagePersonalDingTalk(v bool) *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference {
	s.AllowUserToManagePersonalDingTalk = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) Validate() error {
	return dara.Validate(s)
}

type GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference struct {
	// The MFA methods.
	VerificationTypes []*string `json:"VerificationTypes,omitempty" xml:"VerificationTypes,omitempty" type:"Repeated"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) GetVerificationTypes() []*string {
	return s.VerificationTypes
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) SetVerificationTypes(v []*string) *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference {
	s.VerificationTypes = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) Validate() error {
	return dara.Validate(s)
}
