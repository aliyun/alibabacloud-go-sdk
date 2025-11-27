// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilterUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *FilterUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *FilterUsersResponseBody
	GetRequestId() *string
	SetUsers(v []*FilterUsersResponseBodyUsers) *FilterUsersResponseBody
	GetUsers() []*FilterUsersResponseBodyUsers
}

type FilterUsersResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results. If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the returned NextToken value to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DBD276B5-00FF-5E04-8EF7-5CBA09BF112A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the convenience accounts.
	Users []*FilterUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s FilterUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersResponseBody) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *FilterUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FilterUsersResponseBody) GetUsers() []*FilterUsersResponseBodyUsers {
	return s.Users
}

func (s *FilterUsersResponseBody) SetNextToken(v string) *FilterUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *FilterUsersResponseBody) SetRequestId(v string) *FilterUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *FilterUsersResponseBody) SetUsers(v []*FilterUsersResponseBodyUsers) *FilterUsersResponseBody {
	s.Users = v
	return s
}

func (s *FilterUsersResponseBody) Validate() error {
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

type FilterUsersResponseBodyUsers struct {
	// The date when a convenience account is automatically locked.
	//
	// example:
	//
	// 2023-03-03
	AutoLockTime *string `json:"AutoLockTime,omitempty" xml:"AutoLockTime,omitempty"`
	// The number of cloud desktops that are assigned to the convenience user.
	//
	// example:
	//
	// 1
	DesktopCount *int64 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The number of cloud desktop pools that are assigned to the convenience user. This value is returned if you set `IncludeDesktopGroupCount` to `true`.
	//
	// example:
	//
	// 2
	DesktopGroupCount *int64 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// The email address of the convenience user.
	//
	// example:
	//
	// testName@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the convenience user is a local administrator.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// The username of the convenience user.
	//
	// example:
	//
	// testName
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The additional information about the convenience user.
	ExternalInfo *FilterUsersResponseBodyUsersExternalInfo `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty" type:"Struct"`
	Groups       []*FilterUsersResponseBodyUsersGroups     `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the convenience user.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the convenience user is a tenant administrator.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	IsTenantManager *bool `json:"IsTenantManager,omitempty" xml:"IsTenantManager,omitempty"`
	// The organizations to which the user belongs.
	OrgList []*FilterUsersResponseBodyUsersOrgList `json:"OrgList,omitempty" xml:"OrgList,omitempty" type:"Repeated"`
	// The type of the account ownership.
	//
	// Valid values:
	//
	// 	- CreateFromManager: administrator-activated
	//
	// 	- Normal: user-activated
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// By default, user account passwords do not expire. However, you can set a validity period between 30 and 365 days. Once the period expires, end users must change their password before they can log on to terminals.
	//
	// >  The feature is in invitational preview. If you want to use this feature, submit a ticket.
	//
	// example:
	//
	// 30
	PasswordExpireDays *int32 `json:"PasswordExpireDays,omitempty" xml:"PasswordExpireDays,omitempty"`
	// The number of days remaining until the account password expires.
	//
	// example:
	//
	// 10
	PasswordExpireRestDays *int32 `json:"PasswordExpireRestDays,omitempty" xml:"PasswordExpireRestDays,omitempty"`
	// The mobile number of the convenience user.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The nickname of the convenience user.
	//
	// example:
	//
	// Oliver
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The remarks on the convenience user.
	//
	// example:
	//
	// 1
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The remarks on the convenience account.
	//
	// Valid values:
	//
	// 	- 0: The convenience account is normal.
	//
	// 	- 9: The convenience account is locked.
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The supported identity provider logon methods.
	SupportLoginIdps []*FilterUsersResponseBodyUsersSupportLoginIdps `json:"SupportLoginIdps,omitempty" xml:"SupportLoginIdps,omitempty" type:"Repeated"`
	// The information about the properties.
	UserSetPropertiesModels []*FilterUsersResponseBodyUsersUserSetPropertiesModels `json:"UserSetPropertiesModels,omitempty" xml:"UserSetPropertiesModels,omitempty" type:"Repeated"`
}

func (s FilterUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsers) GetAutoLockTime() *string {
	return s.AutoLockTime
}

func (s *FilterUsersResponseBodyUsers) GetDesktopCount() *int64 {
	return s.DesktopCount
}

func (s *FilterUsersResponseBodyUsers) GetDesktopGroupCount() *int64 {
	return s.DesktopGroupCount
}

func (s *FilterUsersResponseBodyUsers) GetEmail() *string {
	return s.Email
}

func (s *FilterUsersResponseBodyUsers) GetEnableAdminAccess() *bool {
	return s.EnableAdminAccess
}

func (s *FilterUsersResponseBodyUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *FilterUsersResponseBodyUsers) GetExternalInfo() *FilterUsersResponseBodyUsersExternalInfo {
	return s.ExternalInfo
}

func (s *FilterUsersResponseBodyUsers) GetGroups() []*FilterUsersResponseBodyUsersGroups {
	return s.Groups
}

func (s *FilterUsersResponseBodyUsers) GetId() *int64 {
	return s.Id
}

func (s *FilterUsersResponseBodyUsers) GetIsTenantManager() *bool {
	return s.IsTenantManager
}

func (s *FilterUsersResponseBodyUsers) GetOrgList() []*FilterUsersResponseBodyUsersOrgList {
	return s.OrgList
}

func (s *FilterUsersResponseBodyUsers) GetOwnerType() *string {
	return s.OwnerType
}

func (s *FilterUsersResponseBodyUsers) GetPasswordExpireDays() *int32 {
	return s.PasswordExpireDays
}

func (s *FilterUsersResponseBodyUsers) GetPasswordExpireRestDays() *int32 {
	return s.PasswordExpireRestDays
}

func (s *FilterUsersResponseBodyUsers) GetPhone() *string {
	return s.Phone
}

func (s *FilterUsersResponseBodyUsers) GetRealNickName() *string {
	return s.RealNickName
}

func (s *FilterUsersResponseBodyUsers) GetRemark() *string {
	return s.Remark
}

func (s *FilterUsersResponseBodyUsers) GetStatus() *int64 {
	return s.Status
}

func (s *FilterUsersResponseBodyUsers) GetSupportLoginIdps() []*FilterUsersResponseBodyUsersSupportLoginIdps {
	return s.SupportLoginIdps
}

func (s *FilterUsersResponseBodyUsers) GetUserSetPropertiesModels() []*FilterUsersResponseBodyUsersUserSetPropertiesModels {
	return s.UserSetPropertiesModels
}

func (s *FilterUsersResponseBodyUsers) SetAutoLockTime(v string) *FilterUsersResponseBodyUsers {
	s.AutoLockTime = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetDesktopCount(v int64) *FilterUsersResponseBodyUsers {
	s.DesktopCount = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetDesktopGroupCount(v int64) *FilterUsersResponseBodyUsers {
	s.DesktopGroupCount = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetEmail(v string) *FilterUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetEnableAdminAccess(v bool) *FilterUsersResponseBodyUsers {
	s.EnableAdminAccess = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetEndUserId(v string) *FilterUsersResponseBodyUsers {
	s.EndUserId = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetExternalInfo(v *FilterUsersResponseBodyUsersExternalInfo) *FilterUsersResponseBodyUsers {
	s.ExternalInfo = v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetGroups(v []*FilterUsersResponseBodyUsersGroups) *FilterUsersResponseBodyUsers {
	s.Groups = v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetId(v int64) *FilterUsersResponseBodyUsers {
	s.Id = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetIsTenantManager(v bool) *FilterUsersResponseBodyUsers {
	s.IsTenantManager = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetOrgList(v []*FilterUsersResponseBodyUsersOrgList) *FilterUsersResponseBodyUsers {
	s.OrgList = v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetOwnerType(v string) *FilterUsersResponseBodyUsers {
	s.OwnerType = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetPasswordExpireDays(v int32) *FilterUsersResponseBodyUsers {
	s.PasswordExpireDays = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetPasswordExpireRestDays(v int32) *FilterUsersResponseBodyUsers {
	s.PasswordExpireRestDays = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetPhone(v string) *FilterUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetRealNickName(v string) *FilterUsersResponseBodyUsers {
	s.RealNickName = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetRemark(v string) *FilterUsersResponseBodyUsers {
	s.Remark = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetStatus(v int64) *FilterUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetSupportLoginIdps(v []*FilterUsersResponseBodyUsersSupportLoginIdps) *FilterUsersResponseBodyUsers {
	s.SupportLoginIdps = v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetUserSetPropertiesModels(v []*FilterUsersResponseBodyUsersUserSetPropertiesModels) *FilterUsersResponseBodyUsers {
	s.UserSetPropertiesModels = v
	return s
}

func (s *FilterUsersResponseBodyUsers) Validate() error {
	if s.ExternalInfo != nil {
		if err := s.ExternalInfo.Validate(); err != nil {
			return err
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
	if s.OrgList != nil {
		for _, item := range s.OrgList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SupportLoginIdps != nil {
		for _, item := range s.SupportLoginIdps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserSetPropertiesModels != nil {
		for _, item := range s.UserSetPropertiesModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FilterUsersResponseBodyUsersExternalInfo struct {
	// The account that is associated with the convenience user.
	//
	// example:
	//
	// test
	ExternalName *string `json:"ExternalName,omitempty" xml:"ExternalName,omitempty"`
	// The account, student ID, or employee ID that is associated with the convenience user.
	//
	// example:
	//
	// 030801
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
}

func (s FilterUsersResponseBodyUsersExternalInfo) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersResponseBodyUsersExternalInfo) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersExternalInfo) GetExternalName() *string {
	return s.ExternalName
}

func (s *FilterUsersResponseBodyUsersExternalInfo) GetJobNumber() *string {
	return s.JobNumber
}

func (s *FilterUsersResponseBodyUsersExternalInfo) SetExternalName(v string) *FilterUsersResponseBodyUsersExternalInfo {
	s.ExternalName = &v
	return s
}

func (s *FilterUsersResponseBodyUsersExternalInfo) SetJobNumber(v string) *FilterUsersResponseBodyUsersExternalInfo {
	s.JobNumber = &v
	return s
}

func (s *FilterUsersResponseBodyUsersExternalInfo) Validate() error {
	return dara.Validate(s)
}

type FilterUsersResponseBodyUsersGroups struct {
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s FilterUsersResponseBodyUsersGroups) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersResponseBodyUsersGroups) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *FilterUsersResponseBodyUsersGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *FilterUsersResponseBodyUsersGroups) SetGroupId(v string) *FilterUsersResponseBodyUsersGroups {
	s.GroupId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersGroups) SetGroupName(v string) *FilterUsersResponseBodyUsersGroups {
	s.GroupName = &v
	return s
}

func (s *FilterUsersResponseBodyUsersGroups) Validate() error {
	return dara.Validate(s)
}

type FilterUsersResponseBodyUsersOrgList struct {
	// The organization ID.
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The organization name.
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The organization name path.
	OrgNamePath *string `json:"OrgNamePath,omitempty" xml:"OrgNamePath,omitempty"`
}

func (s FilterUsersResponseBodyUsersOrgList) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersResponseBodyUsersOrgList) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersOrgList) GetOrgId() *string {
	return s.OrgId
}

func (s *FilterUsersResponseBodyUsersOrgList) GetOrgName() *string {
	return s.OrgName
}

func (s *FilterUsersResponseBodyUsersOrgList) GetOrgNamePath() *string {
	return s.OrgNamePath
}

func (s *FilterUsersResponseBodyUsersOrgList) SetOrgId(v string) *FilterUsersResponseBodyUsersOrgList {
	s.OrgId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersOrgList) SetOrgName(v string) *FilterUsersResponseBodyUsersOrgList {
	s.OrgName = &v
	return s
}

func (s *FilterUsersResponseBodyUsersOrgList) SetOrgNamePath(v string) *FilterUsersResponseBodyUsersOrgList {
	s.OrgNamePath = &v
	return s
}

func (s *FilterUsersResponseBodyUsersOrgList) Validate() error {
	return dara.Validate(s)
}

type FilterUsersResponseBodyUsersSupportLoginIdps struct {
	// The enterprise identity provider ID.
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The enterprise identity provider name.
	IdpName *string `json:"IdpName,omitempty" xml:"IdpName,omitempty"`
}

func (s FilterUsersResponseBodyUsersSupportLoginIdps) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersResponseBodyUsersSupportLoginIdps) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersSupportLoginIdps) GetIdpId() *string {
	return s.IdpId
}

func (s *FilterUsersResponseBodyUsersSupportLoginIdps) GetIdpName() *string {
	return s.IdpName
}

func (s *FilterUsersResponseBodyUsersSupportLoginIdps) SetIdpId(v string) *FilterUsersResponseBodyUsersSupportLoginIdps {
	s.IdpId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersSupportLoginIdps) SetIdpName(v string) *FilterUsersResponseBodyUsersSupportLoginIdps {
	s.IdpName = &v
	return s
}

func (s *FilterUsersResponseBodyUsersSupportLoginIdps) Validate() error {
	return dara.Validate(s)
}

type FilterUsersResponseBodyUsersUserSetPropertiesModels struct {
	// The property ID.
	//
	// example:
	//
	// 12
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The property name.
	//
	// example:
	//
	// department
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The property type.
	//
	// example:
	//
	// 2
	PropertyType *int32 `json:"PropertyType,omitempty" xml:"PropertyType,omitempty"`
	// The property values.
	PropertyValues []*FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
	// The ID of the convenience user that is bound to the property.
	//
	// example:
	//
	// 12345
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the convenience user that is bound to the property.
	//
	// example:
	//
	// testName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModels) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModels) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) GetPropertyType() *int32 {
	return s.PropertyType
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) GetPropertyValues() []*FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues {
	return s.PropertyValues
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) GetUserId() *int64 {
	return s.UserId
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) GetUserName() *string {
	return s.UserName
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyId(v int64) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyKey(v string) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyKey = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyType(v int32) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyType = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyValues(v []*FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyValues = v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetUserId(v int64) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.UserId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetUserName(v string) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.UserName = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) Validate() error {
	if s.PropertyValues != nil {
		for _, item := range s.PropertyValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues struct {
	// The property value.
	//
	// example:
	//
	// A
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 42
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) GetPropertyValueId() *int64 {
	return s.PropertyValueId
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) SetPropertyValue(v string) *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) SetPropertyValueId(v int64) *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues {
	s.PropertyValueId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) Validate() error {
	return dara.Validate(s)
}
