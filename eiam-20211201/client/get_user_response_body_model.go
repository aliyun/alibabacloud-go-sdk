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
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data object of the account.
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
	return dara.Validate(s)
}

type GetUserResponseBodyUser struct {
	// The time when the account expires. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	AccountExpireTime *int64 `json:"AccountExpireTime,omitempty" xml:"AccountExpireTime,omitempty"`
	// The time when the account was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The list of custom fields that describe the account.
	CustomFields []*GetUserResponseBodyUserCustomFields `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	// The description of the account.
	//
	// example:
	//
	// Test account
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user who owns the account.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the email address has been verified. A value of true indicates that the email address has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the email address has not been verified.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"EmailVerified,omitempty" xml:"EmailVerified,omitempty"`
	// The organizational units to which the account belongs.
	Groups []*GetUserResponseBodyUserGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the instance
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the account lock expires. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	LockExpireTime *int64 `json:"LockExpireTime,omitempty" xml:"LockExpireTime,omitempty"`
	// The organizational units to which the account belongs.
	OrganizationalUnits []*GetUserResponseBodyUserOrganizationalUnits `json:"OrganizationalUnits,omitempty" xml:"OrganizationalUnits,omitempty" type:"Repeated"`
	// The time when the password of the account expires. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// 	- If the value -1 is returned, the password does not expire.
	//
	// 	- If no value is returned, the password does not expire.
	//
	// 	- If a UNIX timestamp is returned, the password expires at the indicated point of time.
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
	// The mobile number of the user who owns the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Indicates whether the mobile number has been verified. A value of true indicates that the mobile number has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the mobile number has not been verified.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"PhoneNumberVerified,omitempty" xml:"PhoneNumberVerified,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// Preferred language
	//
	// example:
	//
	// en-US
	PreferredLanguage *string `json:"PreferredLanguage,omitempty" xml:"PreferredLanguage,omitempty"`
	// The ID of the primary organizational unit to which the account belongs.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	PrimaryOrganizationalUnitId *string `json:"PrimaryOrganizationalUnitId,omitempty" xml:"PrimaryOrganizationalUnitId,omitempty"`
	// The time when the account was registered. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	RegisterTime *int64 `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	// The status of the account. Valid values:
	//
	// 	- enabled: The account is enabled.
	//
	// 	- disabled: The account is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the account was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The external ID of the account. The external ID can be used by external data to map the data of the account in IDaaS EIAM. By default, the external ID is the account ID.
	//
	// For accounts with the same source type and source ID, each account has a unique external ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The ID of the account.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The source ID of the account.
	//
	// If the account was created in IDaaS, its source ID is the ID of the IDaaS instance. If the account was imported, its source ID is the enterprise ID in the source. For example, if the account was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"UserSourceId,omitempty" xml:"UserSourceId,omitempty"`
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
	UserSourceType *string `json:"UserSourceType,omitempty" xml:"UserSourceType,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// name001
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) GetAccountExpireTime() *int64 {
	return s.AccountExpireTime
}

func (s *GetUserResponseBodyUser) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetUserResponseBodyUser) GetCustomFields() []*GetUserResponseBodyUserCustomFields {
	return s.CustomFields
}

func (s *GetUserResponseBodyUser) GetDescription() *string {
	return s.Description
}

func (s *GetUserResponseBodyUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserResponseBodyUser) GetEmail() *string {
	return s.Email
}

func (s *GetUserResponseBodyUser) GetEmailVerified() *bool {
	return s.EmailVerified
}

func (s *GetUserResponseBodyUser) GetGroups() []*GetUserResponseBodyUserGroups {
	return s.Groups
}

func (s *GetUserResponseBodyUser) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserResponseBodyUser) GetLockExpireTime() *int64 {
	return s.LockExpireTime
}

func (s *GetUserResponseBodyUser) GetOrganizationalUnits() []*GetUserResponseBodyUserOrganizationalUnits {
	return s.OrganizationalUnits
}

func (s *GetUserResponseBodyUser) GetPasswordExpireTime() *int64 {
	return s.PasswordExpireTime
}

func (s *GetUserResponseBodyUser) GetPasswordSet() *bool {
	return s.PasswordSet
}

func (s *GetUserResponseBodyUser) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetUserResponseBodyUser) GetPhoneNumberVerified() *bool {
	return s.PhoneNumberVerified
}

func (s *GetUserResponseBodyUser) GetPhoneRegion() *string {
	return s.PhoneRegion
}

func (s *GetUserResponseBodyUser) GetPreferredLanguage() *string {
	return s.PreferredLanguage
}

func (s *GetUserResponseBodyUser) GetPrimaryOrganizationalUnitId() *string {
	return s.PrimaryOrganizationalUnitId
}

func (s *GetUserResponseBodyUser) GetRegisterTime() *int64 {
	return s.RegisterTime
}

func (s *GetUserResponseBodyUser) GetStatus() *string {
	return s.Status
}

func (s *GetUserResponseBodyUser) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetUserResponseBodyUser) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *GetUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *GetUserResponseBodyUser) GetUserSourceId() *string {
	return s.UserSourceId
}

func (s *GetUserResponseBodyUser) GetUserSourceType() *string {
	return s.UserSourceType
}

func (s *GetUserResponseBodyUser) GetUsername() *string {
	return s.Username
}

func (s *GetUserResponseBodyUser) SetAccountExpireTime(v int64) *GetUserResponseBodyUser {
	s.AccountExpireTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetCreateTime(v int64) *GetUserResponseBodyUser {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetCustomFields(v []*GetUserResponseBodyUserCustomFields) *GetUserResponseBodyUser {
	s.CustomFields = v
	return s
}

func (s *GetUserResponseBodyUser) SetDescription(v string) *GetUserResponseBodyUser {
	s.Description = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDisplayName(v string) *GetUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmail(v string) *GetUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmailVerified(v bool) *GetUserResponseBodyUser {
	s.EmailVerified = &v
	return s
}

func (s *GetUserResponseBodyUser) SetGroups(v []*GetUserResponseBodyUserGroups) *GetUserResponseBodyUser {
	s.Groups = v
	return s
}

func (s *GetUserResponseBodyUser) SetInstanceId(v string) *GetUserResponseBodyUser {
	s.InstanceId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLockExpireTime(v int64) *GetUserResponseBodyUser {
	s.LockExpireTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetOrganizationalUnits(v []*GetUserResponseBodyUserOrganizationalUnits) *GetUserResponseBodyUser {
	s.OrganizationalUnits = v
	return s
}

func (s *GetUserResponseBodyUser) SetPasswordExpireTime(v int64) *GetUserResponseBodyUser {
	s.PasswordExpireTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPasswordSet(v bool) *GetUserResponseBodyUser {
	s.PasswordSet = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPhoneNumber(v string) *GetUserResponseBodyUser {
	s.PhoneNumber = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPhoneNumberVerified(v bool) *GetUserResponseBodyUser {
	s.PhoneNumberVerified = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPhoneRegion(v string) *GetUserResponseBodyUser {
	s.PhoneRegion = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPreferredLanguage(v string) *GetUserResponseBodyUser {
	s.PreferredLanguage = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPrimaryOrganizationalUnitId(v string) *GetUserResponseBodyUser {
	s.PrimaryOrganizationalUnitId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetRegisterTime(v int64) *GetUserResponseBodyUser {
	s.RegisterTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetStatus(v string) *GetUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUpdateTime(v int64) *GetUserResponseBodyUser {
	s.UpdateTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserExternalId(v string) *GetUserResponseBodyUser {
	s.UserExternalId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserSourceId(v string) *GetUserResponseBodyUser {
	s.UserSourceId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserSourceType(v string) *GetUserResponseBodyUser {
	s.UserSourceType = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUsername(v string) *GetUserResponseBodyUser {
	s.Username = &v
	return s
}

func (s *GetUserResponseBodyUser) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyUserCustomFields struct {
	// The identifier of the custom field.
	//
	// example:
	//
	// age
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The value of the custom field.
	//
	// example:
	//
	// 10
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
}

func (s GetUserResponseBodyUserCustomFields) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUserCustomFields) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserCustomFields) GetFieldName() *string {
	return s.FieldName
}

func (s *GetUserResponseBodyUserCustomFields) GetFieldValue() *string {
	return s.FieldValue
}

func (s *GetUserResponseBodyUserCustomFields) SetFieldName(v string) *GetUserResponseBodyUserCustomFields {
	s.FieldName = &v
	return s
}

func (s *GetUserResponseBodyUserCustomFields) SetFieldValue(v string) *GetUserResponseBodyUserCustomFields {
	s.FieldValue = &v
	return s
}

func (s *GetUserResponseBodyUserCustomFields) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyUserGroups struct {
	// The description of the organizational unit.
	//
	// example:
	//
	// this is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// group_test_name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetUserResponseBodyUserGroups) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserGroups) GetDescription() *string {
	return s.Description
}

func (s *GetUserResponseBodyUserGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *GetUserResponseBodyUserGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *GetUserResponseBodyUserGroups) SetDescription(v string) *GetUserResponseBodyUserGroups {
	s.Description = &v
	return s
}

func (s *GetUserResponseBodyUserGroups) SetGroupId(v string) *GetUserResponseBodyUserGroups {
	s.GroupId = &v
	return s
}

func (s *GetUserResponseBodyUserGroups) SetGroupName(v string) *GetUserResponseBodyUserGroups {
	s.GroupName = &v
	return s
}

func (s *GetUserResponseBodyUserGroups) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyUserOrganizationalUnits struct {
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// test_ou_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// Indicates whether the organization is the primary organization.
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
}

func (s GetUserResponseBodyUserOrganizationalUnits) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUserOrganizationalUnits) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserOrganizationalUnits) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *GetUserResponseBodyUserOrganizationalUnits) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *GetUserResponseBodyUserOrganizationalUnits) GetPrimary() *bool {
	return s.Primary
}

func (s *GetUserResponseBodyUserOrganizationalUnits) SetOrganizationalUnitId(v string) *GetUserResponseBodyUserOrganizationalUnits {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetUserResponseBodyUserOrganizationalUnits) SetOrganizationalUnitName(v string) *GetUserResponseBodyUserOrganizationalUnits {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetUserResponseBodyUserOrganizationalUnits) SetPrimary(v bool) *GetUserResponseBodyUserOrganizationalUnits {
	s.Primary = &v
	return s
}

func (s *GetUserResponseBodyUserOrganizationalUnits) Validate() error {
	return dara.Validate(s)
}
