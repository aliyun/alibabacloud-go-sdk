// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesResponseBody
	GetCode() *string
	SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody
	GetData() *ListInstancesResponseBodyData
	SetHttpStatusCode(v int32) *ListInstancesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
}

type ListInstancesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 26A34338-5CD9-4C95-A7A6-5BDCE76C6B94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesResponseBody) GetData() *ListInstancesResponseBodyData {
	return s.Data
}

func (s *ListInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyData struct {
	// List.
	List []*ListInstancesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number, ranging from 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, ranging from 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) GetList() []*ListInstancesResponseBodyDataList {
	return s.List
}

func (s *ListInstancesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBodyData) SetList(v []*ListInstancesResponseBodyDataList) *ListInstancesResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageNumber(v int32) *ListInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageSize(v int32) *ListInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetTotalCount(v int32) *ListInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyDataList struct {
	// Administrator list.
	AdminList []*ListInstancesResponseBodyDataListAdminList `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	// The Alibaba Cloud account ID to which the instance belongs.
	//
	// example:
	//
	// 157123456789****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The URL of the Cloud Contact Center instance, used to access the homepage of the Cloud Call Center instance. It is composed of the specific Cloud Call Center URL followed by the instance ID.
	//
	// example:
	//
	// https://ccc-v2.aliyun.com/#/workbench/ccc-test
	ConsoleUrl *string `json:"ConsoleUrl,omitempty" xml:"ConsoleUrl,omitempty"`
	// The creation time of the instance.
	//
	// example:
	//
	// 1624679747000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// 云联络中心的测试实例。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name of the instance, which is globally unique.
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
	NumberList []*ListInstancesResponseBodyDataListNumberList `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
	// Instance status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataList) GetAdminList() []*ListInstancesResponseBodyDataListAdminList {
	return s.AdminList
}

func (s *ListInstancesResponseBodyDataList) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ListInstancesResponseBodyDataList) GetConsoleUrl() *string {
	return s.ConsoleUrl
}

func (s *ListInstancesResponseBodyDataList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListInstancesResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesResponseBodyDataList) GetDomainName() *string {
	return s.DomainName
}

func (s *ListInstancesResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *ListInstancesResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListInstancesResponseBodyDataList) GetNumberList() []*ListInstancesResponseBodyDataListNumberList {
	return s.NumberList
}

func (s *ListInstancesResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyDataList) SetAdminList(v []*ListInstancesResponseBodyDataListAdminList) *ListInstancesResponseBodyDataList {
	s.AdminList = v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetAliyunUid(v string) *ListInstancesResponseBodyDataList {
	s.AliyunUid = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetConsoleUrl(v string) *ListInstancesResponseBodyDataList {
	s.ConsoleUrl = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetCreateTime(v int64) *ListInstancesResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetDescription(v string) *ListInstancesResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetDomainName(v string) *ListInstancesResponseBodyDataList {
	s.DomainName = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetId(v string) *ListInstancesResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetName(v string) *ListInstancesResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetNumberList(v []*ListInstancesResponseBodyDataListNumberList) *ListInstancesResponseBodyDataList {
	s.NumberList = v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetStatus(v string) *ListInstancesResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) Validate() error {
	if s.AdminList != nil {
		for _, item := range s.AdminList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type ListInstancesResponseBodyDataListAdminList struct {
	// The administrator\\"s name.
	//
	// example:
	//
	// 测试坐席
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Mailbox.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Agent extension number.
	//
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Agent logon name.
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
	// The role ID, in the format: role\\@instance ID.
	//
	// example:
	//
	// Admin@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// Role name.
	//
	// example:
	//
	// Admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Agent ID.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ListInstancesResponseBodyDataListAdminList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataListAdminList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataListAdminList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListInstancesResponseBodyDataListAdminList) GetEmail() *string {
	return s.Email
}

func (s *ListInstancesResponseBodyDataListAdminList) GetExtension() *string {
	return s.Extension
}

func (s *ListInstancesResponseBodyDataListAdminList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyDataListAdminList) GetLoginName() *string {
	return s.LoginName
}

func (s *ListInstancesResponseBodyDataListAdminList) GetMobile() *string {
	return s.Mobile
}

func (s *ListInstancesResponseBodyDataListAdminList) GetRoleId() *string {
	return s.RoleId
}

func (s *ListInstancesResponseBodyDataListAdminList) GetRoleName() *string {
	return s.RoleName
}

func (s *ListInstancesResponseBodyDataListAdminList) GetUserId() *string {
	return s.UserId
}

func (s *ListInstancesResponseBodyDataListAdminList) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ListInstancesResponseBodyDataListAdminList) SetDisplayName(v string) *ListInstancesResponseBodyDataListAdminList {
	s.DisplayName = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetEmail(v string) *ListInstancesResponseBodyDataListAdminList {
	s.Email = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetExtension(v string) *ListInstancesResponseBodyDataListAdminList {
	s.Extension = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetInstanceId(v string) *ListInstancesResponseBodyDataListAdminList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetLoginName(v string) *ListInstancesResponseBodyDataListAdminList {
	s.LoginName = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetMobile(v string) *ListInstancesResponseBodyDataListAdminList {
	s.Mobile = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetRoleId(v string) *ListInstancesResponseBodyDataListAdminList {
	s.RoleId = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetRoleName(v string) *ListInstancesResponseBodyDataListAdminList {
	s.RoleName = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetUserId(v string) *ListInstancesResponseBodyDataListAdminList {
	s.UserId = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetWorkMode(v string) *ListInstancesResponseBodyDataListAdminList {
	s.WorkMode = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyDataListNumberList struct {
	// The number.
	//
	// example:
	//
	// 0830011****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s ListInstancesResponseBodyDataListNumberList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataListNumberList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataListNumberList) GetNumber() *string {
	return s.Number
}

func (s *ListInstancesResponseBodyDataListNumberList) SetNumber(v string) *ListInstancesResponseBodyDataListNumberList {
	s.Number = &v
	return s
}

func (s *ListInstancesResponseBodyDataListNumberList) Validate() error {
	return dara.Validate(s)
}
