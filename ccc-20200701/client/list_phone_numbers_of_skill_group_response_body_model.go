// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhoneNumbersOfSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPhoneNumbersOfSkillGroupResponseBody
	GetCode() *string
	SetData(v *ListPhoneNumbersOfSkillGroupResponseBodyData) *ListPhoneNumbersOfSkillGroupResponseBody
	GetData() *ListPhoneNumbersOfSkillGroupResponseBodyData
	SetHttpStatusCode(v int32) *ListPhoneNumbersOfSkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPhoneNumbersOfSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPhoneNumbersOfSkillGroupResponseBody
	GetRequestId() *string
}

type ListPhoneNumbersOfSkillGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListPhoneNumbersOfSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPhoneNumbersOfSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersOfSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) GetData() *ListPhoneNumbersOfSkillGroupResponseBodyData {
	return s.Data
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) SetCode(v string) *ListPhoneNumbersOfSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) SetData(v *ListPhoneNumbersOfSkillGroupResponseBodyData) *ListPhoneNumbersOfSkillGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) SetHttpStatusCode(v int32) *ListPhoneNumbersOfSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) SetMessage(v string) *ListPhoneNumbersOfSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) SetRequestId(v string) *ListPhoneNumbersOfSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPhoneNumbersOfSkillGroupResponseBodyData struct {
	List []*ListPhoneNumbersOfSkillGroupResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListPhoneNumbersOfSkillGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersOfSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) GetList() []*ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	return s.List
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) SetList(v []*ListPhoneNumbersOfSkillGroupResponseBodyDataList) *ListPhoneNumbersOfSkillGroupResponseBodyData {
	s.List = v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) SetPageNumber(v int32) *ListPhoneNumbersOfSkillGroupResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) SetPageSize(v int32) *ListPhoneNumbersOfSkillGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) SetTotalCount(v int32) *ListPhoneNumbersOfSkillGroupResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListPhoneNumbersOfSkillGroupResponseBodyDataList struct {
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
	// 08330011****
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// Bidirection
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListPhoneNumbersOfSkillGroupResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersOfSkillGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) GetActive() *bool {
	return s.Active
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) GetCity() *string {
	return s.City
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) GetNumber() *string {
	return s.Number
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) GetProvince() *string {
	return s.Province
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) GetUsage() *string {
	return s.Usage
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetActive(v bool) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.Active = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetCity(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetContactFlowId(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.ContactFlowId = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetInstanceId(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetNumber(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetProvince(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.Province = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetUsage(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.Usage = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
