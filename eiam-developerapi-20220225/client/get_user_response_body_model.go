// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountExpireTime(v int64) *GetUserResponseBody
	GetAccountExpireTime() *int64
	SetCreateTime(v int64) *GetUserResponseBody
	GetCreateTime() *int64
	SetCustomFields(v []*GetUserResponseBodyCustomFields) *GetUserResponseBody
	GetCustomFields() []*GetUserResponseBodyCustomFields
	SetDescription(v string) *GetUserResponseBody
	GetDescription() *string
	SetDisplayName(v string) *GetUserResponseBody
	GetDisplayName() *string
	SetEmail(v string) *GetUserResponseBody
	GetEmail() *string
	SetEmailVerified(v bool) *GetUserResponseBody
	GetEmailVerified() *bool
	SetGroups(v []*GetUserResponseBodyGroups) *GetUserResponseBody
	GetGroups() []*GetUserResponseBodyGroups
	SetInstanceId(v string) *GetUserResponseBody
	GetInstanceId() *string
	SetLockExpireTime(v int64) *GetUserResponseBody
	GetLockExpireTime() *int64
	SetOrganizationalUnits(v []*GetUserResponseBodyOrganizationalUnits) *GetUserResponseBody
	GetOrganizationalUnits() []*GetUserResponseBodyOrganizationalUnits
	SetPasswordSet(v bool) *GetUserResponseBody
	GetPasswordSet() *bool
	SetPhoneNumber(v string) *GetUserResponseBody
	GetPhoneNumber() *string
	SetPhoneNumberVerified(v bool) *GetUserResponseBody
	GetPhoneNumberVerified() *bool
	SetPhoneRegion(v string) *GetUserResponseBody
	GetPhoneRegion() *string
	SetPrimaryOrganizationalUnitId(v string) *GetUserResponseBody
	GetPrimaryOrganizationalUnitId() *string
	SetRegisterTime(v int64) *GetUserResponseBody
	GetRegisterTime() *int64
	SetStatus(v string) *GetUserResponseBody
	GetStatus() *string
	SetUpdateTime(v int64) *GetUserResponseBody
	GetUpdateTime() *int64
	SetUserExternalId(v string) *GetUserResponseBody
	GetUserExternalId() *string
	SetUserId(v string) *GetUserResponseBody
	GetUserId() *string
	SetUserSourceId(v string) *GetUserResponseBody
	GetUserSourceId() *string
	SetUserSourceType(v string) *GetUserResponseBody
	GetUserSourceType() *string
	SetUsername(v string) *GetUserResponseBody
	GetUsername() *string
}

type GetUserResponseBody struct {
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
	// The extended fields of the account.
	CustomFields []*GetUserResponseBodyCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	// The description of the account.
	//
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
	// The email address of the user.
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
	// The groups to which the account belongs.
	Groups []*GetUserResponseBodyGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
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
	// The organizational units to which the account belongs.
	OrganizationalUnits []*GetUserResponseBodyOrganizationalUnits `json:"organizationalUnits,omitempty" xml:"organizationalUnits,omitempty" type:"Repeated"`
	// Indicates whether the password is set.
	//
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
	// The ID of the primary organizational unit of the account.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	PrimaryOrganizationalUnitId *string `json:"primaryOrganizationalUnitId,omitempty" xml:"primaryOrganizationalUnitId,omitempty"`
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

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetAccountExpireTime() *int64 {
	return s.AccountExpireTime
}

func (s *GetUserResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetUserResponseBody) GetCustomFields() []*GetUserResponseBodyCustomFields {
	return s.CustomFields
}

func (s *GetUserResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetUserResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserResponseBody) GetEmail() *string {
	return s.Email
}

func (s *GetUserResponseBody) GetEmailVerified() *bool {
	return s.EmailVerified
}

func (s *GetUserResponseBody) GetGroups() []*GetUserResponseBodyGroups {
	return s.Groups
}

func (s *GetUserResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserResponseBody) GetLockExpireTime() *int64 {
	return s.LockExpireTime
}

func (s *GetUserResponseBody) GetOrganizationalUnits() []*GetUserResponseBodyOrganizationalUnits {
	return s.OrganizationalUnits
}

func (s *GetUserResponseBody) GetPasswordSet() *bool {
	return s.PasswordSet
}

func (s *GetUserResponseBody) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetUserResponseBody) GetPhoneNumberVerified() *bool {
	return s.PhoneNumberVerified
}

func (s *GetUserResponseBody) GetPhoneRegion() *string {
	return s.PhoneRegion
}

func (s *GetUserResponseBody) GetPrimaryOrganizationalUnitId() *string {
	return s.PrimaryOrganizationalUnitId
}

func (s *GetUserResponseBody) GetRegisterTime() *int64 {
	return s.RegisterTime
}

func (s *GetUserResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetUserResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetUserResponseBody) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *GetUserResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetUserResponseBody) GetUserSourceId() *string {
	return s.UserSourceId
}

func (s *GetUserResponseBody) GetUserSourceType() *string {
	return s.UserSourceType
}

func (s *GetUserResponseBody) GetUsername() *string {
	return s.Username
}

func (s *GetUserResponseBody) SetAccountExpireTime(v int64) *GetUserResponseBody {
	s.AccountExpireTime = &v
	return s
}

func (s *GetUserResponseBody) SetCreateTime(v int64) *GetUserResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBody) SetCustomFields(v []*GetUserResponseBodyCustomFields) *GetUserResponseBody {
	s.CustomFields = v
	return s
}

func (s *GetUserResponseBody) SetDescription(v string) *GetUserResponseBody {
	s.Description = &v
	return s
}

func (s *GetUserResponseBody) SetDisplayName(v string) *GetUserResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBody) SetEmail(v string) *GetUserResponseBody {
	s.Email = &v
	return s
}

func (s *GetUserResponseBody) SetEmailVerified(v bool) *GetUserResponseBody {
	s.EmailVerified = &v
	return s
}

func (s *GetUserResponseBody) SetGroups(v []*GetUserResponseBodyGroups) *GetUserResponseBody {
	s.Groups = v
	return s
}

func (s *GetUserResponseBody) SetInstanceId(v string) *GetUserResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetUserResponseBody) SetLockExpireTime(v int64) *GetUserResponseBody {
	s.LockExpireTime = &v
	return s
}

func (s *GetUserResponseBody) SetOrganizationalUnits(v []*GetUserResponseBodyOrganizationalUnits) *GetUserResponseBody {
	s.OrganizationalUnits = v
	return s
}

func (s *GetUserResponseBody) SetPasswordSet(v bool) *GetUserResponseBody {
	s.PasswordSet = &v
	return s
}

func (s *GetUserResponseBody) SetPhoneNumber(v string) *GetUserResponseBody {
	s.PhoneNumber = &v
	return s
}

func (s *GetUserResponseBody) SetPhoneNumberVerified(v bool) *GetUserResponseBody {
	s.PhoneNumberVerified = &v
	return s
}

func (s *GetUserResponseBody) SetPhoneRegion(v string) *GetUserResponseBody {
	s.PhoneRegion = &v
	return s
}

func (s *GetUserResponseBody) SetPrimaryOrganizationalUnitId(v string) *GetUserResponseBody {
	s.PrimaryOrganizationalUnitId = &v
	return s
}

func (s *GetUserResponseBody) SetRegisterTime(v int64) *GetUserResponseBody {
	s.RegisterTime = &v
	return s
}

func (s *GetUserResponseBody) SetStatus(v string) *GetUserResponseBody {
	s.Status = &v
	return s
}

func (s *GetUserResponseBody) SetUpdateTime(v int64) *GetUserResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetUserResponseBody) SetUserExternalId(v string) *GetUserResponseBody {
	s.UserExternalId = &v
	return s
}

func (s *GetUserResponseBody) SetUserId(v string) *GetUserResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBody) SetUserSourceId(v string) *GetUserResponseBody {
	s.UserSourceId = &v
	return s
}

func (s *GetUserResponseBody) SetUserSourceType(v string) *GetUserResponseBody {
	s.UserSourceType = &v
	return s
}

func (s *GetUserResponseBody) SetUsername(v string) *GetUserResponseBody {
	s.Username = &v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	if s.CustomFields != nil {
		for _, item := range s.CustomFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OrganizationalUnits != nil {
		for _, item := range s.OrganizationalUnits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserResponseBodyCustomFields struct {
	// The name of the extended field.
	//
	// example:
	//
	// xxxx
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// The value of the extended field. Field values are returned as strings regardless of the data types of the fields. For example, true or false is returned for a Boolean field.
	//
	// example:
	//
	// 字段数据值
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s GetUserResponseBodyCustomFields) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyCustomFields) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyCustomFields) GetFieldName() *string {
	return s.FieldName
}

func (s *GetUserResponseBodyCustomFields) GetFieldValue() *string {
	return s.FieldValue
}

func (s *GetUserResponseBodyCustomFields) SetFieldName(v string) *GetUserResponseBodyCustomFields {
	s.FieldName = &v
	return s
}

func (s *GetUserResponseBodyCustomFields) SetFieldValue(v string) *GetUserResponseBodyCustomFields {
	s.FieldValue = &v
	return s
}

func (s *GetUserResponseBodyCustomFields) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyGroups struct {
	// The group description.
	//
	// example:
	//
	// description_demo
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The group ID.
	//
	// example:
	//
	// group_ufdsasn35ea5lmthk267xxxxx
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The group name.
	//
	// example:
	//
	// name_test
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s GetUserResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyGroups) GetDescription() *string {
	return s.Description
}

func (s *GetUserResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *GetUserResponseBodyGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *GetUserResponseBodyGroups) SetDescription(v string) *GetUserResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *GetUserResponseBodyGroups) SetGroupId(v string) *GetUserResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *GetUserResponseBodyGroups) SetGroupName(v string) *GetUserResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *GetUserResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyOrganizationalUnits struct {
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	// Indicates whether the organizational unit is the primary organizational unit.
	//
	// example:
	//
	// true
	Primary *bool `json:"primary,omitempty" xml:"primary,omitempty"`
}

func (s GetUserResponseBodyOrganizationalUnits) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyOrganizationalUnits) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyOrganizationalUnits) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *GetUserResponseBodyOrganizationalUnits) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *GetUserResponseBodyOrganizationalUnits) GetPrimary() *bool {
	return s.Primary
}

func (s *GetUserResponseBodyOrganizationalUnits) SetOrganizationalUnitId(v string) *GetUserResponseBodyOrganizationalUnits {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetUserResponseBodyOrganizationalUnits) SetOrganizationalUnitName(v string) *GetUserResponseBodyOrganizationalUnits {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetUserResponseBodyOrganizationalUnits) SetPrimary(v bool) *GetUserResponseBodyOrganizationalUnits {
	s.Primary = &v
	return s
}

func (s *GetUserResponseBodyOrganizationalUnits) Validate() error {
	return dara.Validate(s)
}
