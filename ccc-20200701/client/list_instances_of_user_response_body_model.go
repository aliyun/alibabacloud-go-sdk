// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesOfUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesOfUserResponseBody
	GetCode() *string
	SetData(v *ListInstancesOfUserResponseBodyData) *ListInstancesOfUserResponseBody
	GetData() *ListInstancesOfUserResponseBodyData
	SetHttpStatusCode(v int32) *ListInstancesOfUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstancesOfUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstancesOfUserResponseBody
	GetRequestId() *string
}

type ListInstancesOfUserResponseBody struct {
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListInstancesOfUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3969FC68-CEC2-4398-B76A-60D2F7EDEBAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesOfUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesOfUserResponseBody) GetData() *ListInstancesOfUserResponseBodyData {
	return s.Data
}

func (s *ListInstancesOfUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstancesOfUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesOfUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesOfUserResponseBody) SetCode(v string) *ListInstancesOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesOfUserResponseBody) SetData(v *ListInstancesOfUserResponseBodyData) *ListInstancesOfUserResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesOfUserResponseBody) SetHttpStatusCode(v int32) *ListInstancesOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesOfUserResponseBody) SetMessage(v string) *ListInstancesOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesOfUserResponseBody) SetRequestId(v string) *ListInstancesOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesOfUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstancesOfUserResponseBodyData struct {
	List []*ListInstancesOfUserResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesOfUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesOfUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBodyData) GetList() []*ListInstancesOfUserResponseBodyDataList {
	return s.List
}

func (s *ListInstancesOfUserResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesOfUserResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesOfUserResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesOfUserResponseBodyData) SetList(v []*ListInstancesOfUserResponseBodyDataList) *ListInstancesOfUserResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstancesOfUserResponseBodyData) SetPageNumber(v int32) *ListInstancesOfUserResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyData) SetPageSize(v int32) *ListInstancesOfUserResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyData) SetTotalCount(v int32) *ListInstancesOfUserResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListInstancesOfUserResponseBodyDataList struct {
	AdminList []*ListInstancesOfUserResponseBodyDataListAdminList `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	// example:
	//
	// 157123456789****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// example:
	//
	// https://ccc-v2.aliyun.com/#/workbench/ccc-test
	ConsoleUrl  *string `json:"ConsoleUrl,omitempty" xml:"ConsoleUrl,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ccc-test
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// ccc-test
	Id         *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	NumberList []*ListInstancesOfUserResponseBodyDataListNumberList `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesOfUserResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesOfUserResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBodyDataList) GetAdminList() []*ListInstancesOfUserResponseBodyDataListAdminList {
	return s.AdminList
}

func (s *ListInstancesOfUserResponseBodyDataList) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ListInstancesOfUserResponseBodyDataList) GetConsoleUrl() *string {
	return s.ConsoleUrl
}

func (s *ListInstancesOfUserResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesOfUserResponseBodyDataList) GetDomainName() *string {
	return s.DomainName
}

func (s *ListInstancesOfUserResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *ListInstancesOfUserResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListInstancesOfUserResponseBodyDataList) GetNumberList() []*ListInstancesOfUserResponseBodyDataListNumberList {
	return s.NumberList
}

func (s *ListInstancesOfUserResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesOfUserResponseBodyDataList) SetAdminList(v []*ListInstancesOfUserResponseBodyDataListAdminList) *ListInstancesOfUserResponseBodyDataList {
	s.AdminList = v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetAliyunUid(v string) *ListInstancesOfUserResponseBodyDataList {
	s.AliyunUid = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetConsoleUrl(v string) *ListInstancesOfUserResponseBodyDataList {
	s.ConsoleUrl = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetDescription(v string) *ListInstancesOfUserResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetDomainName(v string) *ListInstancesOfUserResponseBodyDataList {
	s.DomainName = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetId(v string) *ListInstancesOfUserResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetName(v string) *ListInstancesOfUserResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetNumberList(v []*ListInstancesOfUserResponseBodyDataListNumberList) *ListInstancesOfUserResponseBodyDataList {
	s.NumberList = v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetStatus(v string) *ListInstancesOfUserResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ListInstancesOfUserResponseBodyDataListAdminList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// agent
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// example:
	//
	// 1382114****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// Admin@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// Admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ListInstancesOfUserResponseBodyDataListAdminList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesOfUserResponseBodyDataListAdminList) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) GetEmail() *string {
	return s.Email
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) GetExtension() *string {
	return s.Extension
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) GetLoginName() *string {
	return s.LoginName
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) GetMobile() *string {
	return s.Mobile
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) GetRoleId() *string {
	return s.RoleId
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) GetRoleName() *string {
	return s.RoleName
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) GetUserId() *string {
	return s.UserId
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetDisplayName(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.DisplayName = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetEmail(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.Email = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetExtension(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.Extension = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetInstanceId(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetLoginName(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.LoginName = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetMobile(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.Mobile = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetRoleId(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.RoleId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetRoleName(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.RoleName = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetUserId(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.UserId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetWorkMode(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.WorkMode = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) Validate() error {
	return dara.Validate(s)
}

type ListInstancesOfUserResponseBodyDataListNumberList struct {
	// example:
	//
	// true
	Active *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	City   *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// a3fb6c62-9b49-4942-ae5b-cf2abd4123ek
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0830011****
	Number      *string                                                         `json:"Number,omitempty" xml:"Number,omitempty"`
	Province    *string                                                         `json:"Province,omitempty" xml:"Province,omitempty"`
	SkillGroups []*ListInstancesOfUserResponseBodyDataListNumberListSkillGroups `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty" type:"Repeated"`
	// example:
	//
	// Bidirection
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListInstancesOfUserResponseBodyDataListNumberList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesOfUserResponseBodyDataListNumberList) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) GetActive() *bool {
	return s.Active
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) GetCity() *string {
	return s.City
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) GetNumber() *string {
	return s.Number
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) GetProvince() *string {
	return s.Province
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) GetSkillGroups() []*ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	return s.SkillGroups
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) GetUsage() *string {
	return s.Usage
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) GetUserId() *string {
	return s.UserId
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetActive(v bool) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.Active = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetCity(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.City = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetContactFlowId(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.ContactFlowId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetInstanceId(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetNumber(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.Number = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetProvince(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.Province = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetSkillGroups(v []*ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.SkillGroups = v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetUsage(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.Usage = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetUserId(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.UserId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) Validate() error {
	return dara.Validate(s)
}

type ListInstancesOfUserResponseBodyDataListNumberListSkillGroups struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// skillgroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PhoneNumberCount *int32 `json:"PhoneNumberCount,omitempty" xml:"PhoneNumberCount,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// 2
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) GetName() *string {
	return s.Name
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) GetPhoneNumberCount() *int32 {
	return s.PhoneNumberCount
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) GetUserCount() *int32 {
	return s.UserCount
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetDescription(v string) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.Description = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetDisplayName(v string) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.DisplayName = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetInstanceId(v string) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetName(v string) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.Name = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetPhoneNumberCount(v int32) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.PhoneNumberCount = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetSkillGroupId(v string) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.SkillGroupId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetUserCount(v int32) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.UserCount = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) Validate() error {
	return dara.Validate(s)
}
