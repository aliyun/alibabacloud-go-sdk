// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeUsersResponseBody
	GetRequestId() *string
	SetUsers(v []*DescribeUsersResponseBodyUsers) *DescribeUsersResponseBody
	GetUsers() []*DescribeUsersResponseBodyUsers
}

type DescribeUsersResponseBody struct {
	// The token to start the next query. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the convenience accounts.
	Users []*DescribeUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUsersResponseBody) GetUsers() []*DescribeUsersResponseBodyUsers {
	return s.Users
}

func (s *DescribeUsersResponseBody) SetNextToken(v string) *DescribeUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersResponseBody) SetRequestId(v string) *DescribeUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersResponseBody) SetUsers(v []*DescribeUsersResponseBodyUsers) *DescribeUsersResponseBody {
	s.Users = v
	return s
}

func (s *DescribeUsersResponseBody) Validate() error {
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

type DescribeUsersResponseBodyUsers struct {
	// The work address of the user.
	//
	// example:
	//
	// 杭州市***
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The URL of the user\\"s avatar.
	//
	// example:
	//
	// https://cdn.*****
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// The email address.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether administrator access is enabled.
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// The end user ID.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The name of the user imported from an external system.
	//
	// > This parameter is for internal use only.
	//
	// example:
	//
	// 马**
	ExternalName *string `json:"ExternalName,omitempty" xml:"ExternalName,omitempty"`
	// The extended properties of the user.
	Extras *DescribeUsersResponseBodyUsersExtras `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Struct"`
	// The user groups to which the convenience account belongs.
	Groups []*DescribeUsersResponseBodyUsersGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the convenience account.
	//
	// example:
	//
	// 4205**
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the user is a tenant manager. When you create a convenience account of the `CreateFromManager` type, you must specify a tenant manager. Notifications, such as password resets initiated by an end user from a client, are sent to the tenant manager\\"s email or mobile phone. For more information, see [Create a convenience account](https://help.aliyun.com/document_detail/214472.html).
	//
	// example:
	//
	// true
	IsTenantManager *bool `json:"IsTenantManager,omitempty" xml:"IsTenantManager,omitempty"`
	// The employee ID.
	//
	// example:
	//
	// A10000**
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
	// The nickname of the user.<br>
	//
	// The value is determined from the following parameters, in order of priority:<br>
	//
	// - `RealNickName`
	//
	// - `Remark`
	//
	// - `EndUserId`
	//
	// example:
	//
	// 李**
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The ID of the organization to which the convenience account belongs.
	//
	// > This parameter is deprecated and may be removed in a future release.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The organizations to which the convenience account belongs.
	Orgs []*DescribeUsersResponseBodyUsersOrgs `json:"Orgs,omitempty" xml:"Orgs,omitempty" type:"Repeated"`
	// The type of the convenience account. The account can be activated in one of the following ways:
	//
	// - Tenant manager-activated: The tenant manager sets the username and password. Notifications such as password resets are sent to the tenant manager\\"s email address or mobile phone.
	//
	// - End user-activated: The tenant manager sets the username and the end user\\"s email address or mobile phone. Notifications for the end user, such as the initial password for the cloud desktop, are sent to the end user\\"s email address or mobile phone.
	//
	// example:
	//
	// Normal
	OwnerType              *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	PasswordExpireDays     *int32  `json:"PasswordExpireDays,omitempty" xml:"PasswordExpireDays,omitempty"`
	PasswordExpireRestDays *int32  `json:"PasswordExpireRestDays,omitempty" xml:"PasswordExpireRestDays,omitempty"`
	// The phone number. This parameter is returned only if a phone number is set.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// A list of custom properties for the user.
	Properties []*DescribeUsersResponseBodyUsersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// The display name of the user.
	//
	// example:
	//
	// 李**
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The note about the convenience account.
	//
	// example:
	//
	// Test user.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the convenience account.
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The globally unique ID of the convenience account.
	//
	// example:
	//
	// 41fd1254d8f7****
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s DescribeUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsers) GetAddress() *string {
	return s.Address
}

func (s *DescribeUsersResponseBodyUsers) GetAvatar() *string {
	return s.Avatar
}

func (s *DescribeUsersResponseBodyUsers) GetEmail() *string {
	return s.Email
}

func (s *DescribeUsersResponseBodyUsers) GetEnableAdminAccess() *bool {
	return s.EnableAdminAccess
}

func (s *DescribeUsersResponseBodyUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeUsersResponseBodyUsers) GetExternalName() *string {
	return s.ExternalName
}

func (s *DescribeUsersResponseBodyUsers) GetExtras() *DescribeUsersResponseBodyUsersExtras {
	return s.Extras
}

func (s *DescribeUsersResponseBodyUsers) GetGroups() []*DescribeUsersResponseBodyUsersGroups {
	return s.Groups
}

func (s *DescribeUsersResponseBodyUsers) GetId() *int64 {
	return s.Id
}

func (s *DescribeUsersResponseBodyUsers) GetIsTenantManager() *bool {
	return s.IsTenantManager
}

func (s *DescribeUsersResponseBodyUsers) GetJobNumber() *string {
	return s.JobNumber
}

func (s *DescribeUsersResponseBodyUsers) GetNickName() *string {
	return s.NickName
}

func (s *DescribeUsersResponseBodyUsers) GetOrgId() *string {
	return s.OrgId
}

func (s *DescribeUsersResponseBodyUsers) GetOrgs() []*DescribeUsersResponseBodyUsersOrgs {
	return s.Orgs
}

func (s *DescribeUsersResponseBodyUsers) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DescribeUsersResponseBodyUsers) GetPasswordExpireDays() *int32 {
	return s.PasswordExpireDays
}

func (s *DescribeUsersResponseBodyUsers) GetPasswordExpireRestDays() *int32 {
	return s.PasswordExpireRestDays
}

func (s *DescribeUsersResponseBodyUsers) GetPhone() *string {
	return s.Phone
}

func (s *DescribeUsersResponseBodyUsers) GetProperties() []*DescribeUsersResponseBodyUsersProperties {
	return s.Properties
}

func (s *DescribeUsersResponseBodyUsers) GetRealNickName() *string {
	return s.RealNickName
}

func (s *DescribeUsersResponseBodyUsers) GetRemark() *string {
	return s.Remark
}

func (s *DescribeUsersResponseBodyUsers) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeUsersResponseBodyUsers) GetWyId() *string {
	return s.WyId
}

func (s *DescribeUsersResponseBodyUsers) SetAddress(v string) *DescribeUsersResponseBodyUsers {
	s.Address = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetAvatar(v string) *DescribeUsersResponseBodyUsers {
	s.Avatar = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetEmail(v string) *DescribeUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetEnableAdminAccess(v bool) *DescribeUsersResponseBodyUsers {
	s.EnableAdminAccess = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetEndUserId(v string) *DescribeUsersResponseBodyUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetExternalName(v string) *DescribeUsersResponseBodyUsers {
	s.ExternalName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetExtras(v *DescribeUsersResponseBodyUsersExtras) *DescribeUsersResponseBodyUsers {
	s.Extras = v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetGroups(v []*DescribeUsersResponseBodyUsersGroups) *DescribeUsersResponseBodyUsers {
	s.Groups = v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetId(v int64) *DescribeUsersResponseBodyUsers {
	s.Id = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetIsTenantManager(v bool) *DescribeUsersResponseBodyUsers {
	s.IsTenantManager = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetJobNumber(v string) *DescribeUsersResponseBodyUsers {
	s.JobNumber = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetNickName(v string) *DescribeUsersResponseBodyUsers {
	s.NickName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetOrgId(v string) *DescribeUsersResponseBodyUsers {
	s.OrgId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetOrgs(v []*DescribeUsersResponseBodyUsersOrgs) *DescribeUsersResponseBodyUsers {
	s.Orgs = v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetOwnerType(v string) *DescribeUsersResponseBodyUsers {
	s.OwnerType = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetPasswordExpireDays(v int32) *DescribeUsersResponseBodyUsers {
	s.PasswordExpireDays = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetPasswordExpireRestDays(v int32) *DescribeUsersResponseBodyUsers {
	s.PasswordExpireRestDays = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetPhone(v string) *DescribeUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetProperties(v []*DescribeUsersResponseBodyUsersProperties) *DescribeUsersResponseBodyUsers {
	s.Properties = v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetRealNickName(v string) *DescribeUsersResponseBodyUsers {
	s.RealNickName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetRemark(v string) *DescribeUsersResponseBodyUsers {
	s.Remark = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetStatus(v int64) *DescribeUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetWyId(v string) *DescribeUsersResponseBodyUsers {
	s.WyId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) Validate() error {
	if s.Extras != nil {
		if err := s.Extras.Validate(); err != nil {
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
	if s.Orgs != nil {
		for _, item := range s.Orgs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUsersResponseBodyUsersExtras struct {
	// The number of assigned cloud resources.
	AssignedResourceCount map[string]interface{}                                    `json:"AssignedResourceCount,omitempty" xml:"AssignedResourceCount,omitempty"`
	ResourcePolicyList    []*DescribeUsersResponseBodyUsersExtrasResourcePolicyList `json:"ResourcePolicyList,omitempty" xml:"ResourcePolicyList,omitempty" type:"Repeated"`
}

func (s DescribeUsersResponseBodyUsersExtras) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersExtras) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersExtras) GetAssignedResourceCount() map[string]interface{} {
	return s.AssignedResourceCount
}

func (s *DescribeUsersResponseBodyUsersExtras) GetResourcePolicyList() []*DescribeUsersResponseBodyUsersExtrasResourcePolicyList {
	return s.ResourcePolicyList
}

func (s *DescribeUsersResponseBodyUsersExtras) SetAssignedResourceCount(v map[string]interface{}) *DescribeUsersResponseBodyUsersExtras {
	s.AssignedResourceCount = v
	return s
}

func (s *DescribeUsersResponseBodyUsersExtras) SetResourcePolicyList(v []*DescribeUsersResponseBodyUsersExtrasResourcePolicyList) *DescribeUsersResponseBodyUsersExtras {
	s.ResourcePolicyList = v
	return s
}

func (s *DescribeUsersResponseBodyUsersExtras) Validate() error {
	if s.ResourcePolicyList != nil {
		for _, item := range s.ResourcePolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUsersResponseBodyUsersExtrasResourcePolicyList struct {
	PolicyId   *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DescribeUsersResponseBodyUsersExtrasResourcePolicyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersExtrasResourcePolicyList) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersExtrasResourcePolicyList) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeUsersResponseBodyUsersExtrasResourcePolicyList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeUsersResponseBodyUsersExtrasResourcePolicyList) SetPolicyId(v string) *DescribeUsersResponseBodyUsersExtrasResourcePolicyList {
	s.PolicyId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersExtrasResourcePolicyList) SetPolicyName(v string) *DescribeUsersResponseBodyUsersExtrasResourcePolicyList {
	s.PolicyName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersExtrasResourcePolicyList) Validate() error {
	return dara.Validate(s)
}

type DescribeUsersResponseBodyUsersGroups struct {
	// The ID of the user group.
	//
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// 用户组1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeUsersResponseBodyUsersGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersGroups) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeUsersResponseBodyUsersGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeUsersResponseBodyUsersGroups) SetGroupId(v string) *DescribeUsersResponseBodyUsersGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersGroups) SetGroupName(v string) *DescribeUsersResponseBodyUsersGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeUsersResponseBodyUsersOrgs struct {
	// The ID of the organization.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The name of the organization.
	//
	// example:
	//
	// 部门1
	OrgName     *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	OrgNamePath *string `json:"OrgNamePath,omitempty" xml:"OrgNamePath,omitempty"`
}

func (s DescribeUsersResponseBodyUsersOrgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersOrgs) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersOrgs) GetOrgId() *string {
	return s.OrgId
}

func (s *DescribeUsersResponseBodyUsersOrgs) GetOrgName() *string {
	return s.OrgName
}

func (s *DescribeUsersResponseBodyUsersOrgs) GetOrgNamePath() *string {
	return s.OrgNamePath
}

func (s *DescribeUsersResponseBodyUsersOrgs) SetOrgId(v string) *DescribeUsersResponseBodyUsersOrgs {
	s.OrgId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersOrgs) SetOrgName(v string) *DescribeUsersResponseBodyUsersOrgs {
	s.OrgName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersOrgs) SetOrgNamePath(v string) *DescribeUsersResponseBodyUsersOrgs {
	s.OrgNamePath = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersOrgs) Validate() error {
	return dara.Validate(s)
}

type DescribeUsersResponseBodyUsersProperties struct {
	// The property key.
	//
	// example:
	//
	// Role
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The property value.
	//
	// example:
	//
	// Student
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeUsersResponseBodyUsersProperties) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersProperties) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersProperties) GetKey() *string {
	return s.Key
}

func (s *DescribeUsersResponseBodyUsersProperties) GetValue() *string {
	return s.Value
}

func (s *DescribeUsersResponseBodyUsersProperties) SetKey(v string) *DescribeUsersResponseBodyUsersProperties {
	s.Key = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersProperties) SetValue(v string) *DescribeUsersResponseBodyUsersProperties {
	s.Value = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersProperties) Validate() error {
	return dara.Validate(s)
}
