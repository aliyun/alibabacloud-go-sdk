// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoNotCallNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDoNotCallNumbersResponseBody
	GetCode() *string
	SetData(v *ListDoNotCallNumbersResponseBodyData) *ListDoNotCallNumbersResponseBody
	GetData() *ListDoNotCallNumbersResponseBodyData
	SetHttpStatusCode(v int32) *ListDoNotCallNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDoNotCallNumbersResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListDoNotCallNumbersResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListDoNotCallNumbersResponseBody
	GetRequestId() *string
}

type ListDoNotCallNumbersResponseBody struct {
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListDoNotCallNumbersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDoNotCallNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoNotCallNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoNotCallNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDoNotCallNumbersResponseBody) GetData() *ListDoNotCallNumbersResponseBodyData {
	return s.Data
}

func (s *ListDoNotCallNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDoNotCallNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDoNotCallNumbersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListDoNotCallNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoNotCallNumbersResponseBody) SetCode(v string) *ListDoNotCallNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBody) SetData(v *ListDoNotCallNumbersResponseBodyData) *ListDoNotCallNumbersResponseBody {
	s.Data = v
	return s
}

func (s *ListDoNotCallNumbersResponseBody) SetHttpStatusCode(v int32) *ListDoNotCallNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBody) SetMessage(v string) *ListDoNotCallNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBody) SetParams(v []*string) *ListDoNotCallNumbersResponseBody {
	s.Params = v
	return s
}

func (s *ListDoNotCallNumbersResponseBody) SetRequestId(v string) *ListDoNotCallNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDoNotCallNumbersResponseBodyData struct {
	List []*ListDoNotCallNumbersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListDoNotCallNumbersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoNotCallNumbersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoNotCallNumbersResponseBodyData) GetList() []*ListDoNotCallNumbersResponseBodyDataList {
	return s.List
}

func (s *ListDoNotCallNumbersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDoNotCallNumbersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDoNotCallNumbersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoNotCallNumbersResponseBodyData) SetList(v []*ListDoNotCallNumbersResponseBodyDataList) *ListDoNotCallNumbersResponseBodyData {
	s.List = v
	return s
}

func (s *ListDoNotCallNumbersResponseBodyData) SetPageNumber(v int32) *ListDoNotCallNumbersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBodyData) SetPageSize(v int32) *ListDoNotCallNumbersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBodyData) SetTotalCount(v int32) *ListDoNotCallNumbersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDoNotCallNumbersResponseBodyDataList struct {
	// example:
	//
	// 1626962425000
	CreateTime  *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// agent
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 1900000****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// INSTANCE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListDoNotCallNumbersResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListDoNotCallNumbersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListDoNotCallNumbersResponseBodyDataList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDoNotCallNumbersResponseBodyDataList) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListDoNotCallNumbersResponseBodyDataList) GetCreator() *string {
	return s.Creator
}

func (s *ListDoNotCallNumbersResponseBodyDataList) GetNumber() *string {
	return s.Number
}

func (s *ListDoNotCallNumbersResponseBodyDataList) GetRemark() *string {
	return s.Remark
}

func (s *ListDoNotCallNumbersResponseBodyDataList) GetScope() *string {
	return s.Scope
}

func (s *ListDoNotCallNumbersResponseBodyDataList) SetCreateTime(v int64) *ListDoNotCallNumbersResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBodyDataList) SetCreatedTime(v int64) *ListDoNotCallNumbersResponseBodyDataList {
	s.CreatedTime = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBodyDataList) SetCreator(v string) *ListDoNotCallNumbersResponseBodyDataList {
	s.Creator = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBodyDataList) SetNumber(v string) *ListDoNotCallNumbersResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBodyDataList) SetRemark(v string) *ListDoNotCallNumbersResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBodyDataList) SetScope(v string) *ListDoNotCallNumbersResponseBodyDataList {
	s.Scope = &v
	return s
}

func (s *ListDoNotCallNumbersResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
