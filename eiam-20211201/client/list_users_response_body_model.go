// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUsersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListUsersResponseBody
	GetTotalCount() *int64
	SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody
	GetUsers() []*ListUsersResponseBodyUsers
}

type ListUsersResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of users.
	Users []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUsersResponseBody) GetUsers() []*ListUsersResponseBodyUsers {
	return s.Users
}

func (s *ListUsersResponseBody) SetMaxResults(v int32) *ListUsersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUsersResponseBody) SetNextToken(v string) *ListUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
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
	// The account expiration time. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1652085686179
	AccountExpireTime *int64 `json:"AccountExpireTime,omitempty" xml:"AccountExpireTime,omitempty"`
	// The creation time. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user description.
	//
	// example:
	//
	// xxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the email address is verified. `true` means the user has verified the email address or an administrator has marked it as verified. `false` means the email address is not verified.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"EmailVerified,omitempty" xml:"EmailVerified,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account lock expiration time. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1652085686179
	LockExpireTime *int64 `json:"LockExpireTime,omitempty" xml:"LockExpireTime,omitempty"`
	// The password expiration time. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1652085686179
	PasswordExpireTime *int64 `json:"PasswordExpireTime,omitempty" xml:"PasswordExpireTime,omitempty"`
	// Indicates whether a password is set.
	//
	// example:
	//
	// false
	PasswordSet *bool `json:"PasswordSet,omitempty" xml:"PasswordSet,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Indicates whether the phone number is verified. `true` means the user has verified the phone number or an administrator has marked it as verified. `false` means the phone number is not verified.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"PhoneNumberVerified,omitempty" xml:"PhoneNumberVerified,omitempty"`
	// The country calling code. For example, specify `86` for Chinese mainland. Do not include `00` or a plus sign (+).
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The registration time. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1652085686179
	RegisterTime *int64 `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	// The status. Valid values:
	//
	// - `enabled`: The user is enabled.
	//
	// - `disabled`: The user is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last update time. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The external user ID. This ID maps data from an external system to a user in IDaaS. It defaults to the user ID.
	//
	// Note: The external user ID must be unique for the same source type and source ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The user source ID.
	//
	// If the user is built-in, this is the instance ID. For users from other sources, this is the enterprise ID from the source, such as the `corpId` for a DingTalk organization.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"UserSourceId,omitempty" xml:"UserSourceId,omitempty"`
	// The user source type. Valid values:
	//
	// - `build_in`: The user is a built-in user.
	//
	// - `ding_talk`: The user is imported from DingTalk.
	//
	// - `ad`: The user is imported from AD.
	//
	// - `ldap`: The user is imported from LDAP.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"UserSourceType,omitempty" xml:"UserSourceType,omitempty"`
	// The user name.
	//
	// example:
	//
	// name001
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) GetAccountExpireTime() *int64 {
	return s.AccountExpireTime
}

func (s *ListUsersResponseBodyUsers) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListUsersResponseBodyUsers) GetDescription() *string {
	return s.Description
}

func (s *ListUsersResponseBodyUsers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyUsers) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyUsers) GetEmailVerified() *bool {
	return s.EmailVerified
}

func (s *ListUsersResponseBodyUsers) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersResponseBodyUsers) GetLockExpireTime() *int64 {
	return s.LockExpireTime
}

func (s *ListUsersResponseBodyUsers) GetPasswordExpireTime() *int64 {
	return s.PasswordExpireTime
}

func (s *ListUsersResponseBodyUsers) GetPasswordSet() *bool {
	return s.PasswordSet
}

func (s *ListUsersResponseBodyUsers) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListUsersResponseBodyUsers) GetPhoneNumberVerified() *bool {
	return s.PhoneNumberVerified
}

func (s *ListUsersResponseBodyUsers) GetPhoneRegion() *string {
	return s.PhoneRegion
}

func (s *ListUsersResponseBodyUsers) GetRegisterTime() *int64 {
	return s.RegisterTime
}

func (s *ListUsersResponseBodyUsers) GetStatus() *string {
	return s.Status
}

func (s *ListUsersResponseBodyUsers) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListUsersResponseBodyUsers) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *ListUsersResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersResponseBodyUsers) GetUserSourceId() *string {
	return s.UserSourceId
}

func (s *ListUsersResponseBodyUsers) GetUserSourceType() *string {
	return s.UserSourceType
}

func (s *ListUsersResponseBodyUsers) GetUsername() *string {
	return s.Username
}

func (s *ListUsersResponseBodyUsers) SetAccountExpireTime(v int64) *ListUsersResponseBodyUsers {
	s.AccountExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetCreateTime(v int64) *ListUsersResponseBodyUsers {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetDescription(v string) *ListUsersResponseBodyUsers {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetDisplayName(v string) *ListUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmail(v string) *ListUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmailVerified(v bool) *ListUsersResponseBodyUsers {
	s.EmailVerified = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetInstanceId(v string) *ListUsersResponseBodyUsers {
	s.InstanceId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetLockExpireTime(v int64) *ListUsersResponseBodyUsers {
	s.LockExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPasswordExpireTime(v int64) *ListUsersResponseBodyUsers {
	s.PasswordExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPasswordSet(v bool) *ListUsersResponseBodyUsers {
	s.PasswordSet = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPhoneNumber(v string) *ListUsersResponseBodyUsers {
	s.PhoneNumber = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPhoneNumberVerified(v bool) *ListUsersResponseBodyUsers {
	s.PhoneNumberVerified = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPhoneRegion(v string) *ListUsersResponseBodyUsers {
	s.PhoneRegion = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetRegisterTime(v int64) *ListUsersResponseBodyUsers {
	s.RegisterTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetStatus(v string) *ListUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUpdateTime(v int64) *ListUsersResponseBodyUsers {
	s.UpdateTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserExternalId(v string) *ListUsersResponseBodyUsers {
	s.UserExternalId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserId(v string) *ListUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserSourceId(v string) *ListUsersResponseBodyUsers {
	s.UserSourceId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserSourceType(v string) *ListUsersResponseBodyUsers {
	s.UserSourceType = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUsername(v string) *ListUsersResponseBodyUsers {
	s.Username = &v
	return s
}

func (s *ListUsersResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
