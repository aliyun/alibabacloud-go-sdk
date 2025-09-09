// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUsersResponseBody
	GetTotalCount() *int32
	SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody
	GetUsers() []*ListUsersResponseBodyUsers
}

type ListUsersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of users returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The users returned.
	Users []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUsersResponseBody) GetUsers() []*ListUsersResponseBodyUsers {
	return s.Users
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int32) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyUsers struct {
	// The remarks of the user.
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
	// The end time of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1672502400
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The start time of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
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
	// Indicates whether notifications are sent in the language specified in the global settings or a custom language.
	//
	// 	- **Global**
	//
	// 	- **Custom**
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
	// The location where the mobile phone number of the user is registered. Valid values:
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
	// 	- **JP:*	- Japan, whose international dialing code is +81.
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
	// 	- **SE**: Sweden, whose international dialing code is +46.
	//
	// example:
	//
	// CN
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	// Indicates whether password reset is required upon the next logon. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedResetPassword *bool `json:"NeedResetPassword,omitempty" xml:"NeedResetPassword,omitempty"`
	// The type of the user. Valid values:
	//
	// 	- **Local**: a local user.
	//
	// 	- **Ram**: a RAM user.
	//
	// 	- **AD**: an AD-authenticated user.
	//
	// 	- **LDAP**: an LDAP-authenticated user.
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique ID of the user.
	//
	// >  This parameter uniquely identifies a RAM user of the bastion host. A value is returned for this parameter if **Source*	- is set to **Ram**. No value is returned for this parameter if **Source*	- is set to **Local**.
	//
	// example:
	//
	// 122748924538****
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// An array of the enabled two-factor authentication methods.
	TwoFactorMethods []*string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty" type:"Repeated"`
	// Indicates whether two-factor authentication is enabled for the user. Valid values:
	//
	// 	- **Global**: The global setting applies.
	//
	// 	- **Disable**: Two-factor authentication is disabled.
	//
	// 	- **Enable**: Two-factor authentication is enabled. The user-specific setting for the authentication method applies.
	//
	// example:
	//
	// Enable
	TwoFactorStatus *string `json:"TwoFactorStatus,omitempty" xml:"TwoFactorStatus,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the user.
	//
	// example:
	//
	// abc_def
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// An array that lists the states of users.
	UserState []*string `json:"UserState,omitempty" xml:"UserState,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) GetComment() *string {
	return s.Comment
}

func (s *ListUsersResponseBodyUsers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyUsers) GetEffectiveEndTime() *int64 {
	return s.EffectiveEndTime
}

func (s *ListUsersResponseBodyUsers) GetEffectiveStartTime() *int64 {
	return s.EffectiveStartTime
}

func (s *ListUsersResponseBodyUsers) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyUsers) GetLanguage() *string {
	return s.Language
}

func (s *ListUsersResponseBodyUsers) GetLanguageStatus() *string {
	return s.LanguageStatus
}

func (s *ListUsersResponseBodyUsers) GetMobile() *string {
	return s.Mobile
}

func (s *ListUsersResponseBodyUsers) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *ListUsersResponseBodyUsers) GetNeedResetPassword() *bool {
	return s.NeedResetPassword
}

func (s *ListUsersResponseBodyUsers) GetSource() *string {
	return s.Source
}

func (s *ListUsersResponseBodyUsers) GetSourceUserId() *string {
	return s.SourceUserId
}

func (s *ListUsersResponseBodyUsers) GetTwoFactorMethods() []*string {
	return s.TwoFactorMethods
}

func (s *ListUsersResponseBodyUsers) GetTwoFactorStatus() *string {
	return s.TwoFactorStatus
}

func (s *ListUsersResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersResponseBodyUsers) GetUserName() *string {
	return s.UserName
}

func (s *ListUsersResponseBodyUsers) GetUserState() []*string {
	return s.UserState
}

func (s *ListUsersResponseBodyUsers) SetComment(v string) *ListUsersResponseBodyUsers {
	s.Comment = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetDisplayName(v string) *ListUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEffectiveEndTime(v int64) *ListUsersResponseBodyUsers {
	s.EffectiveEndTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEffectiveStartTime(v int64) *ListUsersResponseBodyUsers {
	s.EffectiveStartTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmail(v string) *ListUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetLanguage(v string) *ListUsersResponseBodyUsers {
	s.Language = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetLanguageStatus(v string) *ListUsersResponseBodyUsers {
	s.LanguageStatus = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetMobile(v string) *ListUsersResponseBodyUsers {
	s.Mobile = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetMobileCountryCode(v string) *ListUsersResponseBodyUsers {
	s.MobileCountryCode = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetNeedResetPassword(v bool) *ListUsersResponseBodyUsers {
	s.NeedResetPassword = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetSource(v string) *ListUsersResponseBodyUsers {
	s.Source = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetSourceUserId(v string) *ListUsersResponseBodyUsers {
	s.SourceUserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetTwoFactorMethods(v []*string) *ListUsersResponseBodyUsers {
	s.TwoFactorMethods = v
	return s
}

func (s *ListUsersResponseBodyUsers) SetTwoFactorStatus(v string) *ListUsersResponseBodyUsers {
	s.TwoFactorStatus = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserId(v string) *ListUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserName(v string) *ListUsersResponseBodyUsers {
	s.UserName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserState(v []*string) *ListUsersResponseBodyUsers {
	s.UserState = v
	return s
}

func (s *ListUsersResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
