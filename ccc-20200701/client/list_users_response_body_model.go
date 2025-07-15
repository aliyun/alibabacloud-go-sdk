// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUsersResponseBody
	GetCode() *string
	SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody
	GetData() *ListUsersResponseBodyData
	SetHttpStatusCode(v int32) *ListUsersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListUsersResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListUsersResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
}

type ListUsersResponseBody struct {
	// example:
	//
	// OK
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUsersResponseBody) GetData() *ListUsersResponseBodyData {
	return s.Data
}

func (s *ListUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUsersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetHttpStatusCode(v int32) *ListUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetParams(v []*string) *ListUsersResponseBody {
	s.Params = v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyData struct {
	List []*ListUsersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) GetList() []*ListUsersResponseBodyDataList {
	return s.List
}

func (s *ListUsersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUsersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUsersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUsersResponseBodyData) SetList(v []*ListUsersResponseBodyDataList) *ListUsersResponseBodyData {
	s.List = v
	return s
}

func (s *ListUsersResponseBodyData) SetPageNumber(v int32) *ListUsersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPageSize(v int32) *ListUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBodyData) SetTotalCount(v int32) *ListUsersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyDataList struct {
	// example:
	//
	// 8033****
	DeviceExt *string `json:"DeviceExt,omitempty" xml:"DeviceExt,omitempty"`
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// OFFLINE
	DeviceState *string `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	// example:
	//
	// 1001
	DisplayId   *string `json:"DisplayId,omitempty" xml:"DisplayId,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 8031****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// agent
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// example:
	//
	// 1382114****
	Mobile                     *string                                                    `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PersonalOutboundNumberList []*ListUsersResponseBodyDataListPersonalOutboundNumberList `json:"PersonalOutboundNumberList,omitempty" xml:"PersonalOutboundNumberList,omitempty" type:"Repeated"`
	Primary                    *bool                                                      `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// example:
	//
	// false
	PrimaryAccount *bool `json:"PrimaryAccount,omitempty" xml:"PrimaryAccount,omitempty"`
	// example:
	//
	// 21234502254620****
	RamId *int64 `json:"RamId,omitempty" xml:"RamId,omitempty"`
	// example:
	//
	// Admin@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// Admin
	RoleName       *string                                        `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	SkillLevelList []*ListUsersResponseBodyDataListSkillLevelList `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty" type:"Repeated"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ListUsersResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataList) GetDeviceExt() *string {
	return s.DeviceExt
}

func (s *ListUsersResponseBodyDataList) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ListUsersResponseBodyDataList) GetDeviceState() *string {
	return s.DeviceState
}

func (s *ListUsersResponseBodyDataList) GetDisplayId() *string {
	return s.DisplayId
}

func (s *ListUsersResponseBodyDataList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyDataList) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyDataList) GetExtension() *string {
	return s.Extension
}

func (s *ListUsersResponseBodyDataList) GetLoginName() *string {
	return s.LoginName
}

func (s *ListUsersResponseBodyDataList) GetMobile() *string {
	return s.Mobile
}

func (s *ListUsersResponseBodyDataList) GetPersonalOutboundNumberList() []*ListUsersResponseBodyDataListPersonalOutboundNumberList {
	return s.PersonalOutboundNumberList
}

func (s *ListUsersResponseBodyDataList) GetPrimary() *bool {
	return s.Primary
}

func (s *ListUsersResponseBodyDataList) GetPrimaryAccount() *bool {
	return s.PrimaryAccount
}

func (s *ListUsersResponseBodyDataList) GetRamId() *int64 {
	return s.RamId
}

func (s *ListUsersResponseBodyDataList) GetRoleId() *string {
	return s.RoleId
}

func (s *ListUsersResponseBodyDataList) GetRoleName() *string {
	return s.RoleName
}

func (s *ListUsersResponseBodyDataList) GetSkillLevelList() []*ListUsersResponseBodyDataListSkillLevelList {
	return s.SkillLevelList
}

func (s *ListUsersResponseBodyDataList) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersResponseBodyDataList) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ListUsersResponseBodyDataList) SetDeviceExt(v string) *ListUsersResponseBodyDataList {
	s.DeviceExt = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetDeviceId(v string) *ListUsersResponseBodyDataList {
	s.DeviceId = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetDeviceState(v string) *ListUsersResponseBodyDataList {
	s.DeviceState = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetDisplayId(v string) *ListUsersResponseBodyDataList {
	s.DisplayId = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetDisplayName(v string) *ListUsersResponseBodyDataList {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetEmail(v string) *ListUsersResponseBodyDataList {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetExtension(v string) *ListUsersResponseBodyDataList {
	s.Extension = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetLoginName(v string) *ListUsersResponseBodyDataList {
	s.LoginName = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetMobile(v string) *ListUsersResponseBodyDataList {
	s.Mobile = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetPersonalOutboundNumberList(v []*ListUsersResponseBodyDataListPersonalOutboundNumberList) *ListUsersResponseBodyDataList {
	s.PersonalOutboundNumberList = v
	return s
}

func (s *ListUsersResponseBodyDataList) SetPrimary(v bool) *ListUsersResponseBodyDataList {
	s.Primary = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetPrimaryAccount(v bool) *ListUsersResponseBodyDataList {
	s.PrimaryAccount = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetRamId(v int64) *ListUsersResponseBodyDataList {
	s.RamId = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetRoleId(v string) *ListUsersResponseBodyDataList {
	s.RoleId = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetRoleName(v string) *ListUsersResponseBodyDataList {
	s.RoleName = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetSkillLevelList(v []*ListUsersResponseBodyDataListSkillLevelList) *ListUsersResponseBodyDataList {
	s.SkillLevelList = v
	return s
}

func (s *ListUsersResponseBodyDataList) SetUserId(v string) *ListUsersResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetWorkMode(v string) *ListUsersResponseBodyDataList {
	s.WorkMode = &v
	return s
}

func (s *ListUsersResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyDataListPersonalOutboundNumberList struct {
	// example:
	//
	// true
	Active *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	City   *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 0830011****
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// Bidirection
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListUsersResponseBodyDataListPersonalOutboundNumberList) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyDataListPersonalOutboundNumberList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) GetActive() *bool {
	return s.Active
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) GetCity() *string {
	return s.City
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) GetNumber() *string {
	return s.Number
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) GetProvince() *string {
	return s.Province
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) GetUsage() *string {
	return s.Usage
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) SetActive(v bool) *ListUsersResponseBodyDataListPersonalOutboundNumberList {
	s.Active = &v
	return s
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) SetCity(v string) *ListUsersResponseBodyDataListPersonalOutboundNumberList {
	s.City = &v
	return s
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) SetNumber(v string) *ListUsersResponseBodyDataListPersonalOutboundNumberList {
	s.Number = &v
	return s
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) SetProvince(v string) *ListUsersResponseBodyDataListPersonalOutboundNumberList {
	s.Province = &v
	return s
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) SetUsage(v string) *ListUsersResponseBodyDataListPersonalOutboundNumberList {
	s.Usage = &v
	return s
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyDataListSkillLevelList struct {
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// skillgroup
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	// example:
	//
	// 5
	SkillLevel *int32 `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
}

func (s ListUsersResponseBodyDataListSkillLevelList) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyDataListSkillLevelList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataListSkillLevelList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListUsersResponseBodyDataListSkillLevelList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListUsersResponseBodyDataListSkillLevelList) GetSkillLevel() *int32 {
	return s.SkillLevel
}

func (s *ListUsersResponseBodyDataListSkillLevelList) SetSkillGroupId(v string) *ListUsersResponseBodyDataListSkillLevelList {
	s.SkillGroupId = &v
	return s
}

func (s *ListUsersResponseBodyDataListSkillLevelList) SetSkillGroupName(v string) *ListUsersResponseBodyDataListSkillLevelList {
	s.SkillGroupName = &v
	return s
}

func (s *ListUsersResponseBodyDataListSkillLevelList) SetSkillLevel(v int32) *ListUsersResponseBodyDataListSkillLevelList {
	s.SkillLevel = &v
	return s
}

func (s *ListUsersResponseBodyDataListSkillLevelList) Validate() error {
	return dara.Validate(s)
}
