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
	// The details of security preferences.
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
	return dara.Validate(s)
}

type GetSecurityPreferenceResponseBodySecurityPreference struct {
	// The AccessKey pair preference.
	AccessKeyPreference *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	// The logon preference.
	LoginProfilePreference *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	// The MFA preference.
	MFAPreference *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	// The maximum idle periods. Unit: days.
	MaxIdleDays *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays `json:"MaxIdleDays,omitempty" xml:"MaxIdleDays,omitempty" type:"Struct"`
	// The personal information preference.
	PersonalInfoPreference *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference `json:"PersonalInfoPreference,omitempty" xml:"PersonalInfoPreference,omitempty" type:"Struct"`
	// The MFA method preference.
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
	return dara.Validate(s)
}

type GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference struct {
	// Indicates whether RAM users can manage their AccessKey pairs. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	// Indicates whether RAM users can change their passwords. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Indicates whether a RAM user can use a passkey for logon. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToLoginWithPasskey *bool `json:"AllowUserToLoginWithPasskey,omitempty" xml:"AllowUserToLoginWithPasskey,omitempty"`
	// Indicates whether RAM users can remember the multi-factor authentication (MFA) devices for seven days. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users. Unit: hours.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	// Indicates whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console. Valid values:
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
	// Indicates whether to enable MFA for RAM users who initiated unusual logons. Valid values:
	//
	// 	- autonomous (default): yes. MFA is prompted for RAM users who initiated unusual logons. However, the RAM users are allowed to skip MFA.
	//
	// 	- enforceVerify: MFA is prompted for RAM users who initiated unusual logons and the RAM users cannot skip MFA.
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
	// Indicates whether RAM users can manage their MFA devices. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	// The maximum number of days that the AccessKey pair of a RAM user can stay unused. If an AccessKey pair is not used in the previous specified number of days, the AccessKey pair is automatically disabled on the next day. The default value is 730. You cannot change the value.
	//
	// example:
	//
	// 730
	MaxIdleDaysForAccessKeys *int32 `json:"MaxIdleDaysForAccessKeys,omitempty" xml:"MaxIdleDaysForAccessKeys,omitempty"`
	// The maximum number of days that a RAM user can stay idle. If a RAM user for whom console logon is enabled does not log on to the console in the previous specified number of days, console logon is automatically disabled for the RAM user on the next day. Single sign-on (SSO) is not involved. The default value is 730. You cannot change the value.
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
	// Indicates whether RAM users can manage their personal DingTalk accounts, such as binding and unbinding of the accounts. Valid values:
	//
	// 	- true
	//
	// 	- false
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
