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
	// The token for the next query. If a query does not return all results, the returned NextToken is not empty. You can pass the returned NextToken in a subsequent query to continue the query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The collection of convenience user information.
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
	// The date when the account is automatically locked, in the YYYY-MM-DD format.
	//
	// example:
	//
	// 2023-03-03
	AutoLockTime *string `json:"AutoLockTime,omitempty" xml:"AutoLockTime,omitempty"`
	// The number of cloud computers assigned to the user.
	//
	// example:
	//
	// 1
	DesktopCount *int64 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The number of cloud computer pools authorized for the user. This value is returned when `IncludeDesktopGroupCount` is set to `true`.
	//
	// example:
	//
	// 2
	DesktopGroupCount *int64 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// The email address.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the user is a local administrator.
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// The username.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The extended user information.
	ExternalInfo *FilterUsersResponseBodyUsersExternalInfo `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty" type:"Struct"`
	// The list of user groups to which the convenience user belongs.
	Groups []*FilterUsersResponseBodyUsersGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 4205**
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the user is a tenant administrator.
	//
	// example:
	//
	// true
	IsTenantManager *bool `json:"IsTenantManager,omitempty" xml:"IsTenantManager,omitempty"`
	// The list of organizations to which the user belongs.
	OrgList []*FilterUsersResponseBodyUsersOrgList `json:"OrgList,omitempty" xml:"OrgList,omitempty" type:"Repeated"`
	// The ownership type of the account.
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The password of the user account is valid permanently by default. You can use this parameter to set a password validity period of 30 to 365 days. After the password expires, the end user must change the password before logging on again.
	//
	// > This feature is currently in invitational preview. To use this feature, submit a ticket.
	//
	// example:
	//
	// 30
	PasswordExpireDays *int32 `json:"PasswordExpireDays,omitempty" xml:"PasswordExpireDays,omitempty"`
	// The remaining days of the password validity period for the account.
	//
	// example:
	//
	// 10
	PasswordExpireRestDays *int32 `json:"PasswordExpireRestDays,omitempty" xml:"PasswordExpireRestDays,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The user nickname.
	//
	// example:
	//
	// Xiaoming
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The user remarks.
	//
	// example:
	//
	// For testing only
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The list of resource policies associated with the convenience user.
	ResourcePolicyList []*FilterUsersResponseBodyUsersResourcePolicyList `json:"ResourcePolicyList,omitempty" xml:"ResourcePolicyList,omitempty" type:"Repeated"`
	// The status of the convenience user.
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The supported logon identity sources.
	SupportLoginIdps []*FilterUsersResponseBodyUsersSupportLoginIdps `json:"SupportLoginIdps,omitempty" xml:"SupportLoginIdps,omitempty" type:"Repeated"`
	// The collection of user properties.
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

func (s *FilterUsersResponseBodyUsers) GetResourcePolicyList() []*FilterUsersResponseBodyUsersResourcePolicyList {
	return s.ResourcePolicyList
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

func (s *FilterUsersResponseBodyUsers) SetResourcePolicyList(v []*FilterUsersResponseBodyUsersResourcePolicyList) *FilterUsersResponseBodyUsers {
	s.ResourcePolicyList = v
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
	if s.ResourcePolicyList != nil {
		for _, item := range s.ResourcePolicyList {
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
	// The name of the external account associated with the user.
	//
	// example:
	//
	// account
	ExternalName *string `json:"ExternalName,omitempty" xml:"ExternalName,omitempty"`
	// The external account, student ID, or employee ID associated with the user.
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
	// The user group ID.
	//
	// example:
	//
	// ug-2412ojkwtybd****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The user group name.
	//
	// example:
	//
	// TestGroup
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
	// The organization structure ID.
	//
	// example:
	//
	// org-1mox****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The organization structure name.
	//
	// example:
	//
	// Design Department
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The organization structure name path.
	//
	// example:
	//
	// Product Department/Design Department
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

type FilterUsersResponseBodyUsersResourcePolicyList struct {
	// The resource policy ID.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The resource policy name.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s FilterUsersResponseBodyUsersResourcePolicyList) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersResponseBodyUsersResourcePolicyList) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersResourcePolicyList) GetPolicyId() *string {
	return s.PolicyId
}

func (s *FilterUsersResponseBodyUsersResourcePolicyList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *FilterUsersResponseBodyUsersResourcePolicyList) SetPolicyId(v string) *FilterUsersResponseBodyUsersResourcePolicyList {
	s.PolicyId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersResourcePolicyList) SetPolicyName(v string) *FilterUsersResponseBodyUsersResourcePolicyList {
	s.PolicyName = &v
	return s
}

func (s *FilterUsersResponseBodyUsersResourcePolicyList) Validate() error {
	return dara.Validate(s)
}

type FilterUsersResponseBodyUsersSupportLoginIdps struct {
	// The enterprise identity source ID.
	//
	// example:
	//
	// idp-hlyexfvwert9m8****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The enterprise identity source name.
	//
	// example:
	//
	// DingTalk
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
	// The user property ID.
	//
	// example:
	//
	// 12
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The user property name.
	//
	// example:
	//
	// department
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The user property number.
	//
	// example:
	//
	// 2
	PropertyType *int32 `json:"PropertyType,omitempty" xml:"PropertyType,omitempty"`
	// The user property value.
	PropertyValues []*FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
	// The ID of the user to which the property is bound.
	//
	// example:
	//
	// 4205**
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the user to which the property is bound.
	//
	// example:
	//
	// alice
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
	// The user property value.
	//
	// example:
	//
	// dev
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The user property value ID.
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
