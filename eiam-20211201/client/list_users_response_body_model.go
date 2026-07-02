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
	// The number of entries per page for paging.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
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
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of account data.
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
	// The account expiration time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	AccountExpireTime *int64 `json:"AccountExpireTime,omitempty" xml:"AccountExpireTime,omitempty"`
	// The account creation time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the account.
	//
	// example:
	//
	// xxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the account.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the email address has been verified. A value of true indicates that the email address has been verified by the user or set as verified by the administrator. A value of false indicates that the email address has not been verified.
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
	// The account lock expiration time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	LockExpireTime *int64 `json:"LockExpireTime,omitempty" xml:"LockExpireTime,omitempty"`
	// The password expiration time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	PasswordExpireTime *int64 `json:"PasswordExpireTime,omitempty" xml:"PasswordExpireTime,omitempty"`
	// Indicates whether a password has been set.
	//
	// example:
	//
	// false
	PasswordSet *bool `json:"PasswordSet,omitempty" xml:"PasswordSet,omitempty"`
	// The phone number of the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Indicates whether the phone number has been verified. A value of true indicates that the phone number has been verified by the user or set as verified by the administrator. A value of false indicates that the phone number has not been verified.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"PhoneNumberVerified,omitempty" xml:"PhoneNumberVerified,omitempty"`
	// The phone region code. Example: The region code for the Chinese mainland is 86, without the 00 or + prefix.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The account registration time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	RegisterTime *int64 `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	// The account status. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the account was last updated, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The external ID of the account, which is used to associate external data with IDaaS accounts. The default value is the IDaaS account ID.
	//
	// Note: The external ID must be unique within the same source type and source ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The source ID of the account.
	//
	// For self-built accounts, the default value is the instance ID. For other types, the value corresponds to the enterprise ID of the respective source. For example, for a DingTalk source, the value corresponds to the corpId of the DingTalk enterprise.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"UserSourceId,omitempty" xml:"UserSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// - build_in: self-built.
	//
	// - ding_talk: imported from DingTalk.
	//
	// - ad: imported from AD.
	//
	// - ldap: imported from LDAP.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"UserSourceType,omitempty" xml:"UserSourceType,omitempty"`
	// The username.
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
