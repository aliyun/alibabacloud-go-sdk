// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnassignedNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUnassignedNumbersResponseBody
	GetCode() *string
	SetData(v *ListUnassignedNumbersResponseBodyData) *ListUnassignedNumbersResponseBody
	GetData() *ListUnassignedNumbersResponseBodyData
	SetHttpStatusCode(v int32) *ListUnassignedNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListUnassignedNumbersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListUnassignedNumbersResponseBody
	GetRequestId() *string
}

type ListUnassignedNumbersResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListUnassignedNumbersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListUnassignedNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUnassignedNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUnassignedNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUnassignedNumbersResponseBody) GetData() *ListUnassignedNumbersResponseBodyData {
	return s.Data
}

func (s *ListUnassignedNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUnassignedNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUnassignedNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUnassignedNumbersResponseBody) SetCode(v string) *ListUnassignedNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUnassignedNumbersResponseBody) SetData(v *ListUnassignedNumbersResponseBodyData) *ListUnassignedNumbersResponseBody {
	s.Data = v
	return s
}

func (s *ListUnassignedNumbersResponseBody) SetHttpStatusCode(v int32) *ListUnassignedNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUnassignedNumbersResponseBody) SetMessage(v string) *ListUnassignedNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUnassignedNumbersResponseBody) SetRequestId(v string) *ListUnassignedNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUnassignedNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUnassignedNumbersResponseBodyData struct {
	List []*ListUnassignedNumbersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUnassignedNumbersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUnassignedNumbersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUnassignedNumbersResponseBodyData) GetList() []*ListUnassignedNumbersResponseBodyDataList {
	return s.List
}

func (s *ListUnassignedNumbersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUnassignedNumbersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnassignedNumbersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUnassignedNumbersResponseBodyData) SetList(v []*ListUnassignedNumbersResponseBodyDataList) *ListUnassignedNumbersResponseBodyData {
	s.List = v
	return s
}

func (s *ListUnassignedNumbersResponseBodyData) SetPageNumber(v int32) *ListUnassignedNumbersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyData) SetPageSize(v int32) *ListUnassignedNumbersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyData) SetTotalCount(v int32) *ListUnassignedNumbersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUnassignedNumbersResponseBodyDataList struct {
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 08330011****
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListUnassignedNumbersResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListUnassignedNumbersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListUnassignedNumbersResponseBodyDataList) GetCity() *string {
	return s.City
}

func (s *ListUnassignedNumbersResponseBodyDataList) GetNumber() *string {
	return s.Number
}

func (s *ListUnassignedNumbersResponseBodyDataList) GetProvince() *string {
	return s.Province
}

func (s *ListUnassignedNumbersResponseBodyDataList) SetCity(v string) *ListUnassignedNumbersResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyDataList) SetNumber(v string) *ListUnassignedNumbersResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyDataList) SetProvince(v string) *ListUnassignedNumbersResponseBodyDataList {
	s.Province = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
