// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListUsersResponseBodyData) *ListUsersResponseBody
	GetData() []*ListUsersResponseBodyData
	SetTotalCount(v int64) *ListUsersResponseBody
	GetTotalCount() *int64
}

type ListUsersResponseBody struct {
	// The queried EIAM accounts.
	Data []*ListUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetData() []*ListUsersResponseBodyData {
	return s.Data
}

func (s *ListUsersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUsersResponseBody) SetData(v []*ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersResponseBodyData struct {
	// The time when the account expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	AccountExpireTime *int64 `json:"accountExpireTime,omitempty" xml:"accountExpireTime,omitempty"`
	// The time when the account was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// xxxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address of the user who owns the account.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Indicates whether the email address has been verified. A value of true indicates that the email address has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the email address has not been verified.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The time when the account lock expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	LockExpireTime *int64 `json:"lockExpireTime,omitempty" xml:"lockExpireTime,omitempty"`
	// example:
	//
	// true
	PasswordSet *bool `json:"passwordSet,omitempty" xml:"passwordSet,omitempty"`
	// The mobile number of the user who owns the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// Indicates whether the mobile number has been verified. A value of true indicates that the mobile number has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the mobile number has not been verified.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	// The time when the account was registered. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	RegisterTime *int64 `json:"registerTime,omitempty" xml:"registerTime,omitempty"`
	// The status of the account. Valid values: enabled disabled
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the account was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The external ID of the account. The external ID can be used to map external data to the data of the account in EIAM of Identity as a Service (IDaaS). By default, the external ID is the account ID.
	//
	// Note: For accounts with the same source type and source ID, each account has a unique external ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The source ID of the account.
	//
	// If the account was created in IDaaS, its source ID is the ID of the IDaaS instance. If the account was imported, its source ID is the enterprise ID in the source. For example, if the account was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"userSourceId,omitempty" xml:"userSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// 	- build_in: The account was created in IDaaS.
	//
	// 	- ding_talk: The account was imported from DingTalk.
	//
	// 	- ad: The account was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The account was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"userSourceType,omitempty" xml:"userSourceType,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// name001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) GetAccountExpireTime() *int64 {
	return s.AccountExpireTime
}

func (s *ListUsersResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListUsersResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListUsersResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyData) GetEmailVerified() *bool {
	return s.EmailVerified
}

func (s *ListUsersResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersResponseBodyData) GetLockExpireTime() *int64 {
	return s.LockExpireTime
}

func (s *ListUsersResponseBodyData) GetPasswordSet() *bool {
	return s.PasswordSet
}

func (s *ListUsersResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListUsersResponseBodyData) GetPhoneNumberVerified() *bool {
	return s.PhoneNumberVerified
}

func (s *ListUsersResponseBodyData) GetPhoneRegion() *string {
	return s.PhoneRegion
}

func (s *ListUsersResponseBodyData) GetRegisterTime() *int64 {
	return s.RegisterTime
}

func (s *ListUsersResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListUsersResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListUsersResponseBodyData) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *ListUsersResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersResponseBodyData) GetUserSourceId() *string {
	return s.UserSourceId
}

func (s *ListUsersResponseBodyData) GetUserSourceType() *string {
	return s.UserSourceType
}

func (s *ListUsersResponseBodyData) GetUsername() *string {
	return s.Username
}

func (s *ListUsersResponseBodyData) SetAccountExpireTime(v int64) *ListUsersResponseBodyData {
	s.AccountExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetCreateTime(v int64) *ListUsersResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetDescription(v string) *ListUsersResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyData) SetDisplayName(v string) *ListUsersResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyData) SetEmail(v string) *ListUsersResponseBodyData {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyData) SetEmailVerified(v bool) *ListUsersResponseBodyData {
	s.EmailVerified = &v
	return s
}

func (s *ListUsersResponseBodyData) SetInstanceId(v string) *ListUsersResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetLockExpireTime(v int64) *ListUsersResponseBodyData {
	s.LockExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPasswordSet(v bool) *ListUsersResponseBodyData {
	s.PasswordSet = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPhoneNumber(v string) *ListUsersResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPhoneNumberVerified(v bool) *ListUsersResponseBodyData {
	s.PhoneNumberVerified = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPhoneRegion(v string) *ListUsersResponseBodyData {
	s.PhoneRegion = &v
	return s
}

func (s *ListUsersResponseBodyData) SetRegisterTime(v int64) *ListUsersResponseBodyData {
	s.RegisterTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetStatus(v string) *ListUsersResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUpdateTime(v int64) *ListUsersResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserExternalId(v string) *ListUsersResponseBodyData {
	s.UserExternalId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserId(v string) *ListUsersResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserSourceId(v string) *ListUsersResponseBodyData {
	s.UserSourceId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserSourceType(v string) *ListUsersResponseBodyData {
	s.UserSourceType = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUsername(v string) *ListUsersResponseBodyData {
	s.Username = &v
	return s
}

func (s *ListUsersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
