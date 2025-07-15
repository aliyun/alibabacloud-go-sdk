// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalNumbersOfUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPersonalNumbersOfUserResponseBody
	GetCode() *string
	SetData(v *ListPersonalNumbersOfUserResponseBodyData) *ListPersonalNumbersOfUserResponseBody
	GetData() *ListPersonalNumbersOfUserResponseBodyData
	SetHttpStatusCode(v int32) *ListPersonalNumbersOfUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPersonalNumbersOfUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPersonalNumbersOfUserResponseBody
	GetRequestId() *string
}

type ListPersonalNumbersOfUserResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListPersonalNumbersOfUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListPersonalNumbersOfUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalNumbersOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonalNumbersOfUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPersonalNumbersOfUserResponseBody) GetData() *ListPersonalNumbersOfUserResponseBodyData {
	return s.Data
}

func (s *ListPersonalNumbersOfUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPersonalNumbersOfUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPersonalNumbersOfUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPersonalNumbersOfUserResponseBody) SetCode(v string) *ListPersonalNumbersOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBody) SetData(v *ListPersonalNumbersOfUserResponseBodyData) *ListPersonalNumbersOfUserResponseBody {
	s.Data = v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBody) SetHttpStatusCode(v int32) *ListPersonalNumbersOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBody) SetMessage(v string) *ListPersonalNumbersOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBody) SetRequestId(v string) *ListPersonalNumbersOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPersonalNumbersOfUserResponseBodyData struct {
	List []*ListPersonalNumbersOfUserResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListPersonalNumbersOfUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalNumbersOfUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonalNumbersOfUserResponseBodyData) GetList() []*ListPersonalNumbersOfUserResponseBodyDataList {
	return s.List
}

func (s *ListPersonalNumbersOfUserResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPersonalNumbersOfUserResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPersonalNumbersOfUserResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPersonalNumbersOfUserResponseBodyData) SetList(v []*ListPersonalNumbersOfUserResponseBodyDataList) *ListPersonalNumbersOfUserResponseBodyData {
	s.List = v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyData) SetPageNumber(v int32) *ListPersonalNumbersOfUserResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyData) SetPageSize(v int32) *ListPersonalNumbersOfUserResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyData) SetTotalCount(v int32) *ListPersonalNumbersOfUserResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListPersonalNumbersOfUserResponseBodyDataList struct {
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
}

func (s ListPersonalNumbersOfUserResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalNumbersOfUserResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) GetActive() *bool {
	return s.Active
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) GetCity() *string {
	return s.City
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) GetNumber() *string {
	return s.Number
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) GetProvince() *string {
	return s.Province
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetActive(v bool) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.Active = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetCity(v string) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetContactFlowId(v string) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.ContactFlowId = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetInstanceId(v string) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetNumber(v string) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetProvince(v string) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.Province = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
