// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceResponseBody
	GetCode() *string
	SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody
	GetData() *GetInstanceResponseBodyData
	SetHttpStatusCode(v int32) *GetInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
}

type GetInstanceResponseBody struct {
	// example:
	//
	// OK
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2778FA12-EDD6-42AA-9B15-AF855072E5E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceResponseBody) GetData() *GetInstanceResponseBodyData {
	return s.Data
}

func (s *GetInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v int32) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyData struct {
	AdminList []*GetInstanceResponseBodyDataAdminList `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
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
	Id         *string                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	NumberList []*GetInstanceResponseBodyDataNumberList `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyData) GetAdminList() []*GetInstanceResponseBodyDataAdminList {
	return s.AdminList
}

func (s *GetInstanceResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *GetInstanceResponseBodyData) GetConsoleUrl() *string {
	return s.ConsoleUrl
}

func (s *GetInstanceResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetInstanceResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *GetInstanceResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetInstanceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetInstanceResponseBodyData) GetNumberList() []*GetInstanceResponseBodyDataNumberList {
	return s.NumberList
}

func (s *GetInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyData) SetAdminList(v []*GetInstanceResponseBodyDataAdminList) *GetInstanceResponseBodyData {
	s.AdminList = v
	return s
}

func (s *GetInstanceResponseBodyData) SetAliyunUid(v string) *GetInstanceResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetConsoleUrl(v string) *GetInstanceResponseBodyData {
	s.ConsoleUrl = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetDescription(v string) *GetInstanceResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetDomainName(v string) *GetInstanceResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetId(v string) *GetInstanceResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetName(v string) *GetInstanceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetNumberList(v []*GetInstanceResponseBodyDataNumberList) *GetInstanceResponseBodyData {
	s.NumberList = v
	return s
}

func (s *GetInstanceResponseBodyData) SetStatus(v string) *GetInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataAdminList struct {
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

func (s GetInstanceResponseBodyDataAdminList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataAdminList) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataAdminList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetInstanceResponseBodyDataAdminList) GetEmail() *string {
	return s.Email
}

func (s *GetInstanceResponseBodyDataAdminList) GetExtension() *string {
	return s.Extension
}

func (s *GetInstanceResponseBodyDataAdminList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyDataAdminList) GetLoginName() *string {
	return s.LoginName
}

func (s *GetInstanceResponseBodyDataAdminList) GetMobile() *string {
	return s.Mobile
}

func (s *GetInstanceResponseBodyDataAdminList) GetRoleId() *string {
	return s.RoleId
}

func (s *GetInstanceResponseBodyDataAdminList) GetRoleName() *string {
	return s.RoleName
}

func (s *GetInstanceResponseBodyDataAdminList) GetUserId() *string {
	return s.UserId
}

func (s *GetInstanceResponseBodyDataAdminList) GetWorkMode() *string {
	return s.WorkMode
}

func (s *GetInstanceResponseBodyDataAdminList) SetDisplayName(v string) *GetInstanceResponseBodyDataAdminList {
	s.DisplayName = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetEmail(v string) *GetInstanceResponseBodyDataAdminList {
	s.Email = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetExtension(v string) *GetInstanceResponseBodyDataAdminList {
	s.Extension = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetInstanceId(v string) *GetInstanceResponseBodyDataAdminList {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetLoginName(v string) *GetInstanceResponseBodyDataAdminList {
	s.LoginName = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetMobile(v string) *GetInstanceResponseBodyDataAdminList {
	s.Mobile = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetRoleId(v string) *GetInstanceResponseBodyDataAdminList {
	s.RoleId = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetRoleName(v string) *GetInstanceResponseBodyDataAdminList {
	s.RoleName = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetUserId(v string) *GetInstanceResponseBodyDataAdminList {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetWorkMode(v string) *GetInstanceResponseBodyDataAdminList {
	s.WorkMode = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataNumberList struct {
	// example:
	//
	// true
	Active *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	City   *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 2ec7a58f-3243-4815-bb21-97b480b95f5e
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0830011****
	Number      *string                                             `json:"Number,omitempty" xml:"Number,omitempty"`
	Province    *string                                             `json:"Province,omitempty" xml:"Province,omitempty"`
	SkillGroups []*GetInstanceResponseBodyDataNumberListSkillGroups `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty" type:"Repeated"`
	// example:
	//
	// Bidirection
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetInstanceResponseBodyDataNumberList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataNumberList) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNumberList) GetActive() *bool {
	return s.Active
}

func (s *GetInstanceResponseBodyDataNumberList) GetCity() *string {
	return s.City
}

func (s *GetInstanceResponseBodyDataNumberList) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *GetInstanceResponseBodyDataNumberList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyDataNumberList) GetNumber() *string {
	return s.Number
}

func (s *GetInstanceResponseBodyDataNumberList) GetProvince() *string {
	return s.Province
}

func (s *GetInstanceResponseBodyDataNumberList) GetSkillGroups() []*GetInstanceResponseBodyDataNumberListSkillGroups {
	return s.SkillGroups
}

func (s *GetInstanceResponseBodyDataNumberList) GetUsage() *string {
	return s.Usage
}

func (s *GetInstanceResponseBodyDataNumberList) GetUserId() *string {
	return s.UserId
}

func (s *GetInstanceResponseBodyDataNumberList) SetActive(v bool) *GetInstanceResponseBodyDataNumberList {
	s.Active = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetCity(v string) *GetInstanceResponseBodyDataNumberList {
	s.City = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetContactFlowId(v string) *GetInstanceResponseBodyDataNumberList {
	s.ContactFlowId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetInstanceId(v string) *GetInstanceResponseBodyDataNumberList {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetNumber(v string) *GetInstanceResponseBodyDataNumberList {
	s.Number = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetProvince(v string) *GetInstanceResponseBodyDataNumberList {
	s.Province = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetSkillGroups(v []*GetInstanceResponseBodyDataNumberListSkillGroups) *GetInstanceResponseBodyDataNumberList {
	s.SkillGroups = v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetUsage(v string) *GetInstanceResponseBodyDataNumberList {
	s.Usage = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetUserId(v string) *GetInstanceResponseBodyDataNumberList {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataNumberListSkillGroups struct {
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
	// 3
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s GetInstanceResponseBodyDataNumberListSkillGroups) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataNumberListSkillGroups) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) GetDescription() *string {
	return s.Description
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) GetName() *string {
	return s.Name
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) GetPhoneNumberCount() *int32 {
	return s.PhoneNumberCount
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) GetUserCount() *int32 {
	return s.UserCount
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetDescription(v string) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetDisplayName(v string) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.DisplayName = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetInstanceId(v string) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetName(v string) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.Name = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetPhoneNumberCount(v int32) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.PhoneNumberCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetSkillGroupId(v string) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.SkillGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetUserCount(v int32) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.UserCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) Validate() error {
	return dara.Validate(s)
}
