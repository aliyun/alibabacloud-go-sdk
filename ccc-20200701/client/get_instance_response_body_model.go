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
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *GetInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyData struct {
	// The list of administrators.
	AdminList []*GetInstanceResponseBodyDataAdminList `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	AgentType *string                                 `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The ID of the Alibaba Cloud account to which the instance belongs.
	//
	// example:
	//
	// 157123456789****
	AliyunUid           *string                                         `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	ChatbotBusinessUnit *GetInstanceResponseBodyDataChatbotBusinessUnit `json:"ChatbotBusinessUnit,omitempty" xml:"ChatbotBusinessUnit,omitempty" type:"Struct"`
	// The URL of the Cloud Contact Center instance homepage. This URL is formed by combining the base URL of Cloud Contact Center and the instance ID.
	//
	// example:
	//
	// https://ccc-v2.aliyun.com/#/workbench/ccc-test
	ConsoleUrl *string `json:"ConsoleUrl,omitempty" xml:"ConsoleUrl,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// 云联络中心的测试实例。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name of the instance. It is globally unique.
	//
	// example:
	//
	// ccc-test
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ccc-test
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance name.
	//
	// example:
	//
	// 测试实例
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of numbers.
	NumberList []*GetInstanceResponseBodyDataNumberList `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
	// The instance status.
	//
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

func (s *GetInstanceResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *GetInstanceResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *GetInstanceResponseBodyData) GetChatbotBusinessUnit() *GetInstanceResponseBodyDataChatbotBusinessUnit {
	return s.ChatbotBusinessUnit
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

func (s *GetInstanceResponseBodyData) SetAgentType(v string) *GetInstanceResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetAliyunUid(v string) *GetInstanceResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetChatbotBusinessUnit(v *GetInstanceResponseBodyDataChatbotBusinessUnit) *GetInstanceResponseBodyData {
	s.ChatbotBusinessUnit = v
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
	if s.AdminList != nil {
		for _, item := range s.AdminList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ChatbotBusinessUnit != nil {
		if err := s.ChatbotBusinessUnit.Validate(); err != nil {
			return err
		}
	}
	if s.NumberList != nil {
		for _, item := range s.NumberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceResponseBodyDataAdminList struct {
	// The name of the administrator.
	//
	// example:
	//
	// 管理员
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The mailbox.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The agent\\"s extension number.
	//
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The agent\\"s logon name.
	//
	// example:
	//
	// agent
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// The agent\\"s personal phone number.
	//
	// example:
	//
	// 1382114****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The role ID. The format is: Role\\@Instance ID.
	//
	// example:
	//
	// Admin@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The role name.
	//
	// example:
	//
	// Admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The agent ID.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The work mode.
	//
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

type GetInstanceResponseBodyDataChatbotBusinessUnit struct {
	UnitId  *int64  `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	UnitKey *string `json:"UnitKey,omitempty" xml:"UnitKey,omitempty"`
}

func (s GetInstanceResponseBodyDataChatbotBusinessUnit) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataChatbotBusinessUnit) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataChatbotBusinessUnit) GetUnitId() *int64 {
	return s.UnitId
}

func (s *GetInstanceResponseBodyDataChatbotBusinessUnit) GetUnitKey() *string {
	return s.UnitKey
}

func (s *GetInstanceResponseBodyDataChatbotBusinessUnit) SetUnitId(v int64) *GetInstanceResponseBodyDataChatbotBusinessUnit {
	s.UnitId = &v
	return s
}

func (s *GetInstanceResponseBodyDataChatbotBusinessUnit) SetUnitKey(v string) *GetInstanceResponseBodyDataChatbotBusinessUnit {
	s.UnitKey = &v
	return s
}

func (s *GetInstanceResponseBodyDataChatbotBusinessUnit) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataNumberList struct {
	// Indicates whether the number is active.
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The city where the number is registered.
	//
	// example:
	//
	// 乐山
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The ID of the contact flow (IVR) associated with the phone number.
	//
	// example:
	//
	// 2ec7a58f-3243-4815-bb21-97b480b95f5e
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 0830011****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The province where the number is registered.
	//
	// example:
	//
	// 四川
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The list of skill groups associated with the number.
	SkillGroups []*GetInstanceResponseBodyDataNumberListSkillGroups `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty" type:"Repeated"`
	// The purpose of the number.
	//
	// example:
	//
	// Bidirection
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The agent ID. If this parameter is not empty, the number is a personal outbound number for the agent.
	//
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
	if s.SkillGroups != nil {
		for _, item := range s.SkillGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceResponseBodyDataNumberListSkillGroups struct {
	// The description of the skill group.
	//
	// example:
	//
	// 云联络中心的测试技能组。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the skill group.
	//
	// example:
	//
	// 测试技能组
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the skill group.
	//
	// example:
	//
	// skillgroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of phone numbers associated with the skill group.
	//
	// example:
	//
	// 1
	PhoneNumberCount *int32 `json:"PhoneNumberCount,omitempty" xml:"PhoneNumberCount,omitempty"`
	// The skill group ID.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The number of agents associated with the skill group.
	//
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
