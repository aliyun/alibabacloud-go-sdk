// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhoneNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPhoneNumbersResponseBody
	GetCode() *string
	SetData(v *ListPhoneNumbersResponseBodyData) *ListPhoneNumbersResponseBody
	GetData() *ListPhoneNumbersResponseBodyData
	SetHttpStatusCode(v int32) *ListPhoneNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPhoneNumbersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPhoneNumbersResponseBody
	GetRequestId() *string
}

type ListPhoneNumbersResponseBody struct {
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListPhoneNumbersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPhoneNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPhoneNumbersResponseBody) GetData() *ListPhoneNumbersResponseBodyData {
	return s.Data
}

func (s *ListPhoneNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPhoneNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPhoneNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPhoneNumbersResponseBody) SetCode(v string) *ListPhoneNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ListPhoneNumbersResponseBody) SetData(v *ListPhoneNumbersResponseBodyData) *ListPhoneNumbersResponseBody {
	s.Data = v
	return s
}

func (s *ListPhoneNumbersResponseBody) SetHttpStatusCode(v int32) *ListPhoneNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPhoneNumbersResponseBody) SetMessage(v string) *ListPhoneNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ListPhoneNumbersResponseBody) SetRequestId(v string) *ListPhoneNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhoneNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPhoneNumbersResponseBodyData struct {
	List []*ListPhoneNumbersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPhoneNumbersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersResponseBodyData) GetList() []*ListPhoneNumbersResponseBodyDataList {
	return s.List
}

func (s *ListPhoneNumbersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPhoneNumbersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPhoneNumbersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPhoneNumbersResponseBodyData) SetList(v []*ListPhoneNumbersResponseBodyDataList) *ListPhoneNumbersResponseBodyData {
	s.List = v
	return s
}

func (s *ListPhoneNumbersResponseBodyData) SetPageNumber(v int32) *ListPhoneNumbersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyData) SetPageSize(v int32) *ListPhoneNumbersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyData) SetTotalCount(v int32) *ListPhoneNumbersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListPhoneNumbersResponseBodyDataList struct {
	// example:
	//
	// true
	Active *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	City   *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// a3fb6c62-9b49-4942-ae5b-cf2abd4123ek
	ContactFlowId   *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	ContactFlowName *string `json:"ContactFlowName,omitempty" xml:"ContactFlowName,omitempty"`
	// example:
	//
	// 1617958538000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 08330011****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// alicom
	Provider    *string                                            `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Province    *string                                            `json:"Province,omitempty" xml:"Province,omitempty"`
	SkillGroups []*ListPhoneNumbersResponseBodyDataListSkillGroups `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty" type:"Repeated"`
	// example:
	//
	// M1
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// Bidirection
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListPhoneNumbersResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersResponseBodyDataList) GetActive() *bool {
	return s.Active
}

func (s *ListPhoneNumbersResponseBodyDataList) GetCity() *string {
	return s.City
}

func (s *ListPhoneNumbersResponseBodyDataList) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *ListPhoneNumbersResponseBodyDataList) GetContactFlowName() *string {
	return s.ContactFlowName
}

func (s *ListPhoneNumbersResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPhoneNumbersResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPhoneNumbersResponseBodyDataList) GetNumber() *string {
	return s.Number
}

func (s *ListPhoneNumbersResponseBodyDataList) GetProvider() *string {
	return s.Provider
}

func (s *ListPhoneNumbersResponseBodyDataList) GetProvince() *string {
	return s.Province
}

func (s *ListPhoneNumbersResponseBodyDataList) GetSkillGroups() []*ListPhoneNumbersResponseBodyDataListSkillGroups {
	return s.SkillGroups
}

func (s *ListPhoneNumbersResponseBodyDataList) GetTags() *string {
	return s.Tags
}

func (s *ListPhoneNumbersResponseBodyDataList) GetUsage() *string {
	return s.Usage
}

func (s *ListPhoneNumbersResponseBodyDataList) GetUserId() *string {
	return s.UserId
}

func (s *ListPhoneNumbersResponseBodyDataList) SetActive(v bool) *ListPhoneNumbersResponseBodyDataList {
	s.Active = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetCity(v string) *ListPhoneNumbersResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetContactFlowId(v string) *ListPhoneNumbersResponseBodyDataList {
	s.ContactFlowId = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetContactFlowName(v string) *ListPhoneNumbersResponseBodyDataList {
	s.ContactFlowName = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetCreateTime(v string) *ListPhoneNumbersResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetInstanceId(v string) *ListPhoneNumbersResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetNumber(v string) *ListPhoneNumbersResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetProvider(v string) *ListPhoneNumbersResponseBodyDataList {
	s.Provider = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetProvince(v string) *ListPhoneNumbersResponseBodyDataList {
	s.Province = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetSkillGroups(v []*ListPhoneNumbersResponseBodyDataListSkillGroups) *ListPhoneNumbersResponseBodyDataList {
	s.SkillGroups = v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetTags(v string) *ListPhoneNumbersResponseBodyDataList {
	s.Tags = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetUsage(v string) *ListPhoneNumbersResponseBodyDataList {
	s.Usage = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetUserId(v string) *ListPhoneNumbersResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ListPhoneNumbersResponseBodyDataListSkillGroups struct {
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
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListPhoneNumbersResponseBodyDataListSkillGroups) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersResponseBodyDataListSkillGroups) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) GetName() *string {
	return s.Name
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) SetDisplayName(v string) *ListPhoneNumbersResponseBodyDataListSkillGroups {
	s.DisplayName = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) SetInstanceId(v string) *ListPhoneNumbersResponseBodyDataListSkillGroups {
	s.InstanceId = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) SetName(v string) *ListPhoneNumbersResponseBodyDataListSkillGroups {
	s.Name = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) SetSkillGroupId(v string) *ListPhoneNumbersResponseBodyDataListSkillGroups {
	s.SkillGroupId = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) Validate() error {
	return dara.Validate(s)
}
