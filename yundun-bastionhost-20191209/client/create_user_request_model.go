// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateUserRequest
	GetComment() *string
	SetDisplayName(v string) *CreateUserRequest
	GetDisplayName() *string
	SetEffectiveEndTime(v int64) *CreateUserRequest
	GetEffectiveEndTime() *int64
	SetEffectiveStartTime(v int64) *CreateUserRequest
	GetEffectiveStartTime() *int64
	SetEmail(v string) *CreateUserRequest
	GetEmail() *string
	SetInstanceId(v string) *CreateUserRequest
	GetInstanceId() *string
	SetLanguage(v string) *CreateUserRequest
	GetLanguage() *string
	SetLanguageStatus(v string) *CreateUserRequest
	GetLanguageStatus() *string
	SetMobile(v string) *CreateUserRequest
	GetMobile() *string
	SetMobileCountryCode(v string) *CreateUserRequest
	GetMobileCountryCode() *string
	SetNeedResetPassword(v bool) *CreateUserRequest
	GetNeedResetPassword() *bool
	SetPassword(v string) *CreateUserRequest
	GetPassword() *string
	SetRegionId(v string) *CreateUserRequest
	GetRegionId() *string
	SetSource(v string) *CreateUserRequest
	GetSource() *string
	SetSourceUserId(v string) *CreateUserRequest
	GetSourceUserId() *string
	SetTwoFactorMethods(v string) *CreateUserRequest
	GetTwoFactorMethods() *string
	SetTwoFactorStatus(v string) *CreateUserRequest
	GetTwoFactorStatus() *string
	SetUserName(v string) *CreateUserRequest
	GetUserName() *string
}

type CreateUserRequest struct {
	// The remarks of the user that you want to add. The remarks can be up to 500 characters in length.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The display name of the user that you want to add. The display name can be up to 128 characters in length.
	//
	// >  If you leave this parameter empty, the logon name is used as the display name.
	//
	// example:
	//
	// Bob
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end time of the validity period of the user. Specify a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1672502400
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The start time of the validity period of the user. Specify a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1669630029
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The email address of the user that you want to add.
	//
	// >
	//
	// 	- This parameter is required if TwoFactorStatus is set to Enable and TwoFactorMethods is set to email, or if TwoFactorStatus is set to Global and TwoFactorMethods is set to email in the global two-factor authentication settings.
	//
	// 	- You can call the [GetInstanceTwoFactor](https://help.aliyun.com/document_detail/462968.html) operation to query the global two-factor authentication settings.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the bastion host to which you want to add a user.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required if LanguageStatus is set to Custom. Valid values:
	//
	// 	- **zh-cn**: simplified Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Specifies whether to send notifications in the language specified in the global settings or a custom language.
	//
	// 	- **Global**
	//
	// 	- **Custom**
	//
	// >  If you leave this parameter empty, the default value Global is used.
	//
	// example:
	//
	// Custom
	LanguageStatus *string `json:"LanguageStatus,omitempty" xml:"LanguageStatus,omitempty"`
	// The mobile phone number of the user that you want to add.
	//
	// >
	//
	// 	- This parameter is required if TwoFactorStatus is set to Enable and TwoFactorMethods is set to sms or dingtalk, or if TwoFactorStatus is set to Global and TwoFactorMethods is set to sms or dingtalk in the global two-factor authentication settings.
	//
	// 	- You can call the [GetInstanceTwoFactor](https://help.aliyun.com/document_detail/462968.html) operation to query the global two-factor authentication settings.
	//
	// example:
	//
	// 1359999****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The location where the mobile phone number of the user is registered. Default value: CN. Valid values:
	//
	// 	- **CN**: the Chinese mainland, whose international dialing code is +86.
	//
	// 	- **HK**: Hong Kong (China), whose international dialing code is +852.
	//
	// 	- **MO**: Macao (China), whose international dialing code is +853.
	//
	// 	- **TW**: Taiwan (China), whose international dialing code is +886.
	//
	// 	- **RU**: Russia, whose international dialing code is +7.
	//
	// 	- **SG**: Singapore, whose international dialing code is +65.
	//
	// 	- **MY**: Malaysia, whose international dialing code is +60.
	//
	// 	- **ID**: Indonesia, whose international dialing code is +62.
	//
	// 	- **DE**: Germany, whose international dialing code is +49.
	//
	// 	- **AU**: Australia, whose international dialing code is +61.
	//
	// 	- **US**: US, whose international dialing code is +1.
	//
	// 	- **AE**: United Arab Emirates, whose international dialing code is +971.
	//
	// 	- **JP**: Japan, whose international dialing code is +81.
	//
	// 	- **GB**: UK, whose international dialing code is +44.
	//
	// 	- **IN**: India, whose international dialing code is +91.
	//
	// 	- **KR**: Republic of Korea, whose international dialing code is +82.
	//
	// 	- **PH**: Philippines, whose international dialing code is +63.
	//
	// 	- **CH**: Switzerland, whose international dialing code is +41.
	//
	// 	- **SE:*	- Sweden, whose international dialing code is +46.
	//
	// 	- **SA:*	- Saudi Arabia, whose international dialing code is +966.
	//
	// example:
	//
	// CN
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	// Specifies whether password reset is required upon the next logon. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  If you leave this parameter empty, the default value false is used.
	//
	// example:
	//
	// true
	NeedResetPassword *bool `json:"NeedResetPassword,omitempty" xml:"NeedResetPassword,omitempty"`
	// The logon password of the user that you want to add. The logon password must be 8 to 128 characters in length. It must contain uppercase letters, lowercase letters, digits, and special characters.
	//
	// > This parameter is required if Source is set to Local.
	//
	// example:
	//
	// 213****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the bastion host to which you want to add a user.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the user that you want to add. Valid values:
	//
	// 	- **Local**: a local user.
	//
	// 	- **Ram**: a RAM user.
	//
	// 	- **AD**: an AD-authenticated user.
	//
	// 	- **LDAP**: an LDAP-authenticated user.
	//
	// This parameter is required.
	//
	// example:
	//
	// local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique ID of the user that you want to add.
	//
	// >
	//
	// 	- This parameter uniquely identifies a RAM user of the bastion host. This parameter is required if Source is set to Ram. You can call the [ListUsers](https://help.aliyun.com/document_detail/28684.html) operation in RAM to obtain the unique ID of the user from the UserId response parameter.
	//
	// 	- This parameter is required if Source is set to AD or LDAP. Specify the distinguished name (DN) of the Active Directory (AD)-authenticated user or Lightweight Directory Access Protocol (LDAP)-authenticated user that you want to add.
	//
	// example:
	//
	// 122748924538****
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// The two-factor authentication method. You can select only one method. Valid values:
	//
	// 	- **sms**: text message-based two-factor authentication.
	//
	// 	- **email**: email-based two-factor authentication.
	//
	// 	- **dingtalk**: DingTalk-based two-factor authentication.
	//
	// 	- **totp OTP**: one-time password (OTP) token-based two-factor authentication.
	//
	// >  If TwoFactorStatus is set to Enable, you must select one of the preceding values for TwoFactorMethods.
	//
	// example:
	//
	// ["sms"]
	TwoFactorMethods *string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty"`
	// Specifies whether two-factor authentication is enabled for the user. Valid values:
	//
	// 	- **Global**: The global settings apply.
	//
	// 	- **Disable**: Two-factor authentication is disabled.
	//
	// 	- **Enable**: Two-factor authentication is enabled and user-specific settings apply.
	//
	// >  If you leave this parameter empty, the default value Global is used.
	//
	// example:
	//
	// Enable
	TwoFactorStatus *string `json:"TwoFactorStatus,omitempty" xml:"TwoFactorStatus,omitempty"`
	// The logon name of the user that you want to add. The logon name must be 1 to 128 characters in length and can contain only letters, digits, and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// abc_def
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserRequest) GetEffectiveEndTime() *int64 {
	return s.EffectiveEndTime
}

func (s *CreateUserRequest) GetEffectiveStartTime() *int64 {
	return s.EffectiveStartTime
}

func (s *CreateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateUserRequest) GetLanguageStatus() *string {
	return s.LanguageStatus
}

func (s *CreateUserRequest) GetMobile() *string {
	return s.Mobile
}

func (s *CreateUserRequest) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *CreateUserRequest) GetNeedResetPassword() *bool {
	return s.NeedResetPassword
}

func (s *CreateUserRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateUserRequest) GetSource() *string {
	return s.Source
}

func (s *CreateUserRequest) GetSourceUserId() *string {
	return s.SourceUserId
}

func (s *CreateUserRequest) GetTwoFactorMethods() *string {
	return s.TwoFactorMethods
}

func (s *CreateUserRequest) GetTwoFactorStatus() *string {
	return s.TwoFactorStatus
}

func (s *CreateUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserRequest) SetComment(v string) *CreateUserRequest {
	s.Comment = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEffectiveEndTime(v int64) *CreateUserRequest {
	s.EffectiveEndTime = &v
	return s
}

func (s *CreateUserRequest) SetEffectiveStartTime(v int64) *CreateUserRequest {
	s.EffectiveStartTime = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetInstanceId(v string) *CreateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserRequest) SetLanguage(v string) *CreateUserRequest {
	s.Language = &v
	return s
}

func (s *CreateUserRequest) SetLanguageStatus(v string) *CreateUserRequest {
	s.LanguageStatus = &v
	return s
}

func (s *CreateUserRequest) SetMobile(v string) *CreateUserRequest {
	s.Mobile = &v
	return s
}

func (s *CreateUserRequest) SetMobileCountryCode(v string) *CreateUserRequest {
	s.MobileCountryCode = &v
	return s
}

func (s *CreateUserRequest) SetNeedResetPassword(v bool) *CreateUserRequest {
	s.NeedResetPassword = &v
	return s
}

func (s *CreateUserRequest) SetPassword(v string) *CreateUserRequest {
	s.Password = &v
	return s
}

func (s *CreateUserRequest) SetRegionId(v string) *CreateUserRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUserRequest) SetSource(v string) *CreateUserRequest {
	s.Source = &v
	return s
}

func (s *CreateUserRequest) SetSourceUserId(v string) *CreateUserRequest {
	s.SourceUserId = &v
	return s
}

func (s *CreateUserRequest) SetTwoFactorMethods(v string) *CreateUserRequest {
	s.TwoFactorMethods = &v
	return s
}

func (s *CreateUserRequest) SetTwoFactorStatus(v string) *CreateUserRequest {
	s.TwoFactorStatus = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	return dara.Validate(s)
}
