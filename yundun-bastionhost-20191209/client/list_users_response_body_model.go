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
	// The list of users returned.
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
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	// The end time of the validity period of the user, in seconds (UNIX timestamp format).
	//
	// example:
	//
	// 1672502400
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The start time of the validity period of the user, in seconds (UNIX timestamp format).
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
	// The language for message notifications. This parameter is required when LanguageStatus is set to Custom. Valid values:
	//
	// - **zh-cn**: Simplified Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The language setting for message notifications. Valid values:
	//
	// - **Global**: follows the global settings
	//
	// - **Custom**: custom
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
	// The country code of the mobile phone number of the user. Valid values:
	//
	// - **CN**: the Chinese mainland (+86)
	//
	// - **HK**: Hong Kong (China) (+852)
	//
	// - **MO**: Macao (China) (+853)
	//
	// - **TW**: Taiwan (China) (+886)
	//
	// - **RU**: Russia (+7)
	//
	// - **SG**: Singapore (+65)
	//
	// - **MY**: Malaysia (+60)
	//
	// - **ID**: Indonesia (+62)
	//
	// - **DE**: Germany (+49)
	//
	// - **AU**: Australia (+61)
	//
	// - **US**: United States (+1)
	//
	// - **AE**: Dubai (+971)
	//
	// - **JP**: Japan (+81)
	//
	// - **GB**: United Kingdom (+44)
	//
	// - **IN**: India (+91)
	//
	// - **KR**: South Korea (+82)
	//
	// - **PH**: Philippines (+63)
	//
	// - **CH**: Switzerland (+41)
	//
	// - **SE**: Sweden (+46)
	//
	// example:
	//
	// CN
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	// Indicates whether the password must be reset upon next logon. Valid values:
	//
	// - **true**: The password must be reset.
	//
	// - **false**: The password does not need to be reset.
	//
	// example:
	//
	// true
	NeedResetPassword *bool `json:"NeedResetPassword,omitempty" xml:"NeedResetPassword,omitempty"`
	// The source of the user. Valid values:
	//
	// - **Local**: local user
	//
	// - **Ram**: Resource Access Management (RAM) user
	//
	// - **AD**: AD user
	//
	// - **LDAP**: LDAP user
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique identity of the user.
	//
	// > This parameter is the unique identity of the Resource Access Management (RAM) user that corresponds to the bastion host user. This parameter is returned when the user source is a RAM user (that is, **Source*	- is set to **Ram**). If the user source is a local user (that is, **Source*	- is set to **Local**), this parameter is empty.
	//
	// example:
	//
	// 122748924538****
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// The array of enabled two-factor authentication methods.
	TwoFactorMethods []*string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty" type:"Repeated"`
	// The two-factor authentication status of the user. Valid values:
	//
	// - **Global**: follows the global settings
	//
	// - **Disable**: two-factor authentication disabled
	//
	// - **Enable**: two-factor authentication enabled, follows individual user settings
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
	// The user status array.
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
