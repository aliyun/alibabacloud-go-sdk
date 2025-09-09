// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ModifyUserRequest
	GetComment() *string
	SetDisplayName(v string) *ModifyUserRequest
	GetDisplayName() *string
	SetEffectiveEndTime(v int64) *ModifyUserRequest
	GetEffectiveEndTime() *int64
	SetEffectiveStartTime(v int64) *ModifyUserRequest
	GetEffectiveStartTime() *int64
	SetEmail(v string) *ModifyUserRequest
	GetEmail() *string
	SetInstanceId(v string) *ModifyUserRequest
	GetInstanceId() *string
	SetLanguage(v string) *ModifyUserRequest
	GetLanguage() *string
	SetLanguageStatus(v string) *ModifyUserRequest
	GetLanguageStatus() *string
	SetMobile(v string) *ModifyUserRequest
	GetMobile() *string
	SetMobileCountryCode(v string) *ModifyUserRequest
	GetMobileCountryCode() *string
	SetNeedResetPassword(v bool) *ModifyUserRequest
	GetNeedResetPassword() *bool
	SetPassword(v string) *ModifyUserRequest
	GetPassword() *string
	SetRegionId(v string) *ModifyUserRequest
	GetRegionId() *string
	SetTwoFactorMethods(v string) *ModifyUserRequest
	GetTwoFactorMethods() *string
	SetTwoFactorStatus(v string) *ModifyUserRequest
	GetTwoFactorStatus() *string
	SetUserId(v string) *ModifyUserRequest
	GetUserId() *string
}

type ModifyUserRequest struct {
	// The new remarks of the user. The remarks can be up to 500 characters in length.
	//
	// >  Leave this parameter empty if you do not want to change the remarks of the user.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The new display name of the user. The display name can be up to 128 characters in length.
	//
	// >  Leave this parameter empty if you do not want to change the display name of the user.
	//
	// example:
	//
	// Bob
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end time of the validity period of the user. Specify a UNIX timestamp. Unit: seconds.
	//
	// >  Leave this parameter empty if you do not want to change the end time of the validity period.
	//
	// example:
	//
	// 1672502400
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The start time of the validity period of the user. Specify a UNIX timestamp. Unit: seconds.
	//
	// >  Leave this parameter empty if you do not want to change the start time of the validity period.
	//
	// example:
	//
	// 1669630029
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The new email address of the user.
	//
	// >
	//
	// 	- This parameter is required if TwoFactorStatus is set to Enable and TwoFactorMethods is set to email, or if TwoFactorStatus is set to Global and TwoFactorMethods is set to email in the global two-factor authentication settings.
	//
	// 	- You can call the [GetInstanceTwoFactor](https://help.aliyun.com/document_detail/462968.html) operation to query the global two-factor authentication settings.
	//
	// 	- Leave this parameter empty if you do not want to change the email address of the user.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the bastion host on which you want to modify the information about the user.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required if LanguageStatus is set to Custom.
	//
	// - **zh-cn**: simplified Chinese
	//
	// - **en**: English
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
	// >  Leave this parameter empty if you do not want to change the natural language used to notify the user.
	//
	// example:
	//
	// Custom
	LanguageStatus *string `json:"LanguageStatus,omitempty" xml:"LanguageStatus,omitempty"`
	// The new mobile phone number of the user.
	//
	// >
	//
	// 	- This parameter is required if TwoFactorStatus is set to Enable and TwoFactorMethods is set to sms or dingtalk, or if TwoFactorStatus is set to Global and TwoFactorMethods is set to sms or dingtalk in the global two-factor authentication settings.
	//
	// 	- You can call the [GetInstanceTwoFactor](https://help.aliyun.com/document_detail/462968.html) operation to query the global two-factor authentication settings.
	//
	// 	- Leave this parameter empty if you do not want to change the mobile phone number of the user.
	//
	// example:
	//
	// 1358888****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The country where the new mobile number of the user is registered. Valid values:
	//
	// 	- **CN:*	- the Chinese mainland, whose country calling code is +86
	//
	// 	- **HK:*	- Hong Kong (China), whose country calling code is +852
	//
	// 	- **MO:*	- Macao (China), whose country calling code is +853
	//
	// 	- **TW:*	- Taiwan (China), whose country calling code is +886
	//
	// 	- **RU:*	- Russia, whose country calling code is +7
	//
	// 	- **SG:*	- Singapore, whose country calling code is +65
	//
	// 	- **MY:*	- Malaysia, whose country calling code is +60
	//
	// 	- **ID:*	- Indonesia, whose country calling code is +62
	//
	// 	- **DE:*	- Germany, whose country calling code is +49
	//
	// 	- **AU:*	- Australia, whose country calling code is +61
	//
	// 	- **US:*	- US, whose country calling code is +1
	//
	// 	- **AE:*	- United Arab Emirates, whose country calling code is +971
	//
	// 	- **JP:*	- Japan, whose country calling code is +81
	//
	// 	- **GB:*	- UK, whose country calling code is +44
	//
	// 	- **IN:*	- India, whose country calling code is +91
	//
	// 	- **KR:*	- Republic of Korea, whose country calling code is +82
	//
	// 	- **PH:*	- Philippines, whose country calling code is +63
	//
	// 	- **CH:*	- Switzerland, whose country calling code is +41
	//
	// 	- **SE:*	- Sweden, whose country calling code is +46
	//
	// 	- **SA:*	- Saudi Arabia, whose country calling code is +966
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
	// >  Leave this parameter empty if you do not want to change the password reset settings for the user.
	//
	// example:
	//
	// true
	NeedResetPassword *bool `json:"NeedResetPassword,omitempty" xml:"NeedResetPassword,omitempty"`
	// The new password of the user. The password must be 8 to 128 characters in length. It must contain uppercase letters, lowercase letters, digits, and special characters.
	//
	// > Leave this parameter empty if you do not want to change the password of the user.
	//
	// example:
	//
	// 321****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the bastion host on which you want to modify the information about the user.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The two-factor authentication method. You can select only one method. Valid values:
	//
	// 	- **sms**: text message-based two-factor authentication.
	//
	// 	- **email**: email-based two-factor authentication.
	//
	// 	- **dingtalk**: DingTalk-based two-factor authentication.
	//
	// 	- **totp OTP:*	- one-time password (OTP) token-based two-factor authentication.
	//
	// >  If TwoFactorStatus is set to Enable, you must specify one of the valid values as TwoFactorMethods.
	//
	// example:
	//
	// sms
	TwoFactorMethods *string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty"`
	// Specifies whether two-factor authentication is enabled for the user. Valid values:
	//
	// 	- **Global**: The global settings apply.
	//
	// 	- **Disable**: Two-factor authentication is disabled.
	//
	// 	- **Enable**: Two-factor authentication is enabled and user-specific settings apply.
	//
	// >  Leave this parameter empty if you do not want to change the two-factory authentication settings for the user.
	//
	// example:
	//
	// Enable
	TwoFactorStatus *string `json:"TwoFactorStatus,omitempty" xml:"TwoFactorStatus,omitempty"`
	// The ID of the user whose information you want to modify.
	//
	// >  You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ModifyUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ModifyUserRequest) GetEffectiveEndTime() *int64 {
	return s.EffectiveEndTime
}

func (s *ModifyUserRequest) GetEffectiveStartTime() *int64 {
	return s.EffectiveStartTime
}

func (s *ModifyUserRequest) GetEmail() *string {
	return s.Email
}

func (s *ModifyUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserRequest) GetLanguage() *string {
	return s.Language
}

func (s *ModifyUserRequest) GetLanguageStatus() *string {
	return s.LanguageStatus
}

func (s *ModifyUserRequest) GetMobile() *string {
	return s.Mobile
}

func (s *ModifyUserRequest) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *ModifyUserRequest) GetNeedResetPassword() *bool {
	return s.NeedResetPassword
}

func (s *ModifyUserRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserRequest) GetTwoFactorMethods() *string {
	return s.TwoFactorMethods
}

func (s *ModifyUserRequest) GetTwoFactorStatus() *string {
	return s.TwoFactorStatus
}

func (s *ModifyUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ModifyUserRequest) SetComment(v string) *ModifyUserRequest {
	s.Comment = &v
	return s
}

func (s *ModifyUserRequest) SetDisplayName(v string) *ModifyUserRequest {
	s.DisplayName = &v
	return s
}

func (s *ModifyUserRequest) SetEffectiveEndTime(v int64) *ModifyUserRequest {
	s.EffectiveEndTime = &v
	return s
}

func (s *ModifyUserRequest) SetEffectiveStartTime(v int64) *ModifyUserRequest {
	s.EffectiveStartTime = &v
	return s
}

func (s *ModifyUserRequest) SetEmail(v string) *ModifyUserRequest {
	s.Email = &v
	return s
}

func (s *ModifyUserRequest) SetInstanceId(v string) *ModifyUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserRequest) SetLanguage(v string) *ModifyUserRequest {
	s.Language = &v
	return s
}

func (s *ModifyUserRequest) SetLanguageStatus(v string) *ModifyUserRequest {
	s.LanguageStatus = &v
	return s
}

func (s *ModifyUserRequest) SetMobile(v string) *ModifyUserRequest {
	s.Mobile = &v
	return s
}

func (s *ModifyUserRequest) SetMobileCountryCode(v string) *ModifyUserRequest {
	s.MobileCountryCode = &v
	return s
}

func (s *ModifyUserRequest) SetNeedResetPassword(v bool) *ModifyUserRequest {
	s.NeedResetPassword = &v
	return s
}

func (s *ModifyUserRequest) SetPassword(v string) *ModifyUserRequest {
	s.Password = &v
	return s
}

func (s *ModifyUserRequest) SetRegionId(v string) *ModifyUserRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserRequest) SetTwoFactorMethods(v string) *ModifyUserRequest {
	s.TwoFactorMethods = &v
	return s
}

func (s *ModifyUserRequest) SetTwoFactorStatus(v string) *ModifyUserRequest {
	s.TwoFactorStatus = &v
	return s
}

func (s *ModifyUserRequest) SetUserId(v string) *ModifyUserRequest {
	s.UserId = &v
	return s
}

func (s *ModifyUserRequest) Validate() error {
	return dara.Validate(s)
}
