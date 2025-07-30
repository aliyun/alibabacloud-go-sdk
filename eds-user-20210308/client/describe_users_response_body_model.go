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
	// The token that determines the start point of the next query. If this parameter is left empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The convenience accounts.
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
	return dara.Validate(s)
}

type DescribeUsersResponseBodyUsers struct {
	// The work address of the convenience user.
	//
	// example:
	//
	// Hangzhou \\*\\*\\*
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The profile picture of the convenience user.
	//
	// example:
	//
	// https://cdn.*****
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// The email address of the convenience user.
	//
	// example:
	//
	// username@example.com
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EnableAdminAccess *bool   `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// The username of the convenience user.
	//
	// example:
	//
	// alice
	EndUserId    *string                               `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ExternalName *string                               `json:"ExternalName,omitempty" xml:"ExternalName,omitempty"`
	Extras       *DescribeUsersResponseBodyUsersExtras `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Struct"`
	// The user groups to which the convenience user belongs.
	Groups []*DescribeUsersResponseBodyUsersGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the convenience user.
	//
	// example:
	//
	// 4205**
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the convenience user is an administrator. If the convenience user is of the administrator-activated type, you must specify a user administrator. Notifications such as password reset on a client are sent to the email address or mobile number of the user administrator. For more information, see [Create a convenience user](https://help.aliyun.com/document_detail/214472.html).
	//
	// example:
	//
	// true
	IsTenantManager *bool `json:"IsTenantManager,omitempty" xml:"IsTenantManager,omitempty"`
	// The employee number of the convenience user.
	//
	// example:
	//
	// A10000**
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
	// The nickname of the convenience user.
	//
	// example:
	//
	// Lee
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The ID of the organization to which the convenience user belongs.
	//
	// >  This parameter will be deprecated in the future.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The organizations to which the convenience user belongs.
	Orgs []*DescribeUsersResponseBodyUsersOrgs `json:"Orgs,omitempty" xml:"Orgs,omitempty" type:"Repeated"`
	// The type of the convenience account.
	//
	// 	- Administrator-activated type: The administrator specifies the username and password of the convenience account. User notifications such as password reset notifications are sent to the email address or mobile number of the administrator.
	//
	// 	- User-activated type: The administrator specifies the username and the email address or mobile number of a convenience user. Notifications such as activation notifications that contain the default password are sent to the email address or mobile number of the convenience user.
	//
	// Valid values:
	//
	// 	- CreateFromManager
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     administrator-activated
	//
	//     <!-- -->
	//
	// 	- Normal
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     user-activated
	//
	//     <!-- -->
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The mobile number of the convenience user. If you leave this parameter empty, the value of this parameter is not returned.
	//
	// example:
	//
	// 1381111****
	Phone        *string                                     `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Properties   []*DescribeUsersResponseBodyUsersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	RealNickName *string                                     `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The remarks on the convenience user.
	//
	// example:
	//
	// TestUser
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the convenience user.
	//
	// Valid values:
	//
	// 	- 0: The convenience user is normal.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- 9: The convenience user is locked.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The globally unique ID of the convenience user.
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
	return dara.Validate(s)
}

type DescribeUsersResponseBodyUsersExtras struct {
	AssignedResourceCount map[string]interface{} `json:"AssignedResourceCount,omitempty" xml:"AssignedResourceCount,omitempty"`
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

func (s *DescribeUsersResponseBodyUsersExtras) SetAssignedResourceCount(v map[string]interface{}) *DescribeUsersResponseBodyUsersExtras {
	s.AssignedResourceCount = v
	return s
}

func (s *DescribeUsersResponseBodyUsersExtras) Validate() error {
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
	// User Group 1
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
	// The organization ID.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The organization name.
	//
	// example:
	//
	// Organization 1
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
