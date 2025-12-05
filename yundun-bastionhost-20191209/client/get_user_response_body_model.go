// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
	SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody
	GetUser() *GetUserResponseBodyUser
}

type GetUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed information about the queried user.
	User *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) GetUser() *GetUserResponseBodyUser {
	return s.User
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserResponseBodyUser struct {
	// The description of the user.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The display name of the user.
	//
	// example:
	//
	// Bob
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1672502400
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The beginning of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1669630029
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// 1099**@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required if LanguageStatus is set to Custom.
	//
	// - **zh-cn**: simplified Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Indicates whether notifications are sent in the language specified in the global settings or a custom language.
	//
	// 	- **Global**: Global
	//
	// 	- **Custom**: Custom
	//
	// example:
	//
	// Custom
	LanguageStatus *string `json:"LanguageStatus,omitempty" xml:"LanguageStatus,omitempty"`
	// The mobile phone number of the user.
	//
	// example:
	//
	// 1359999****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The location in which the mobile number of the user is registered. Valid values:
	//
	// 	- **CN**: the Chinese mainland, whose country calling code is +86
	//
	// 	- **HK**: Hong Kong (China), whose country calling code is +852
	//
	// 	- **MO**: Macao (China), whose country calling code is +853
	//
	// 	- **TW**: Taiwan (China), whose country calling code is +886
	//
	// 	- **RU**: Russia, whose country calling code is +7
	//
	// 	- **SG**: Singapore, whose country calling code is +65
	//
	// 	- **MY**: Malaysia, whose country calling code is +60
	//
	// 	- **ID**: Indonesia, whose country calling code is +62
	//
	// 	- **DE**: Germany, whose country calling code is +49
	//
	// 	- **AU**: Australia, whose country calling code is +61
	//
	// 	- **US**: US, whose country calling code is +1
	//
	// 	- **AE**: United Arab Emirates, whose country calling code is +971
	//
	// 	- **JP:*	- Japan, whose country calling code is +81
	//
	// 	- **GB**: UK, whose country calling code is +44
	//
	// 	- **IN**: India, whose country calling code is +91
	//
	// 	- **KR**: Republic of Korea, whose country calling code is +82
	//
	// 	- **PH**: Philippines, whose country calling code is +63
	//
	// 	- **CH**: Switzerland, whose country calling code is +41
	//
	// 	- **SE**: Sweden, whose country calling code is +46
	//
	// example:
	//
	// CN
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	// Specifies whether password reset is required upon the next logon. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	NeedResetPassword *bool `json:"NeedResetPassword,omitempty" xml:"NeedResetPassword,omitempty"`
	// The source of the user. Valid values:
	//
	// 	- **Local**: a local user
	//
	// 	- **Ram**: a RAM user
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique ID of the user.
	//
	// > This parameter uniquely identifies a RAM user of the bastion host. A value is returned for this parameter if the **Source*	- parameter is set to **Ram**. No value is returned for this parameter if the **Source*	- parameter is set to **Local**.
	//
	// example:
	//
	// 122748924538****
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// An array that consists of the details of the two-factor authentication method.
	TwoFactorMethods []*string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty" type:"Repeated"`
	// The two-factor authentication status of the user. Valid values:
	//
	// 	- **Global**: The global settings are used.
	//
	// 	- **Disable**: The two-factor authentication is disabled.
	//
	// 	- **Enable**: The two-factor authentication is enabled and the user-specific setting is used.
	//
	// example:
	//
	// Enable
	TwoFactorStatus *string `json:"TwoFactorStatus,omitempty" xml:"TwoFactorStatus,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the user.
	//
	// example:
	//
	// abcabc_def
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// An array that consists of the details of the user status.
	UserState []*string `json:"UserState,omitempty" xml:"UserState,omitempty" type:"Repeated"`
}

func (s GetUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) GetComment() *string {
	return s.Comment
}

func (s *GetUserResponseBodyUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserResponseBodyUser) GetEffectiveEndTime() *int64 {
	return s.EffectiveEndTime
}

func (s *GetUserResponseBodyUser) GetEffectiveStartTime() *int64 {
	return s.EffectiveStartTime
}

func (s *GetUserResponseBodyUser) GetEmail() *string {
	return s.Email
}

func (s *GetUserResponseBodyUser) GetLanguage() *string {
	return s.Language
}

func (s *GetUserResponseBodyUser) GetLanguageStatus() *string {
	return s.LanguageStatus
}

func (s *GetUserResponseBodyUser) GetMobile() *string {
	return s.Mobile
}

func (s *GetUserResponseBodyUser) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *GetUserResponseBodyUser) GetNeedResetPassword() *bool {
	return s.NeedResetPassword
}

func (s *GetUserResponseBodyUser) GetSource() *string {
	return s.Source
}

func (s *GetUserResponseBodyUser) GetSourceUserId() *string {
	return s.SourceUserId
}

func (s *GetUserResponseBodyUser) GetTwoFactorMethods() []*string {
	return s.TwoFactorMethods
}

func (s *GetUserResponseBodyUser) GetTwoFactorStatus() *string {
	return s.TwoFactorStatus
}

func (s *GetUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *GetUserResponseBodyUser) GetUserName() *string {
	return s.UserName
}

func (s *GetUserResponseBodyUser) GetUserState() []*string {
	return s.UserState
}

func (s *GetUserResponseBodyUser) SetComment(v string) *GetUserResponseBodyUser {
	s.Comment = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDisplayName(v string) *GetUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEffectiveEndTime(v int64) *GetUserResponseBodyUser {
	s.EffectiveEndTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEffectiveStartTime(v int64) *GetUserResponseBodyUser {
	s.EffectiveStartTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmail(v string) *GetUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLanguage(v string) *GetUserResponseBodyUser {
	s.Language = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLanguageStatus(v string) *GetUserResponseBodyUser {
	s.LanguageStatus = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMobile(v string) *GetUserResponseBodyUser {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMobileCountryCode(v string) *GetUserResponseBodyUser {
	s.MobileCountryCode = &v
	return s
}

func (s *GetUserResponseBodyUser) SetNeedResetPassword(v bool) *GetUserResponseBodyUser {
	s.NeedResetPassword = &v
	return s
}

func (s *GetUserResponseBodyUser) SetSource(v string) *GetUserResponseBodyUser {
	s.Source = &v
	return s
}

func (s *GetUserResponseBodyUser) SetSourceUserId(v string) *GetUserResponseBodyUser {
	s.SourceUserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetTwoFactorMethods(v []*string) *GetUserResponseBodyUser {
	s.TwoFactorMethods = v
	return s
}

func (s *GetUserResponseBodyUser) SetTwoFactorStatus(v string) *GetUserResponseBodyUser {
	s.TwoFactorStatus = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserName(v string) *GetUserResponseBodyUser {
	s.UserName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserState(v []*string) *GetUserResponseBodyUser {
	s.UserState = v
	return s
}

func (s *GetUserResponseBodyUser) Validate() error {
	return dara.Validate(s)
}
