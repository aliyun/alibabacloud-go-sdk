// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCallTagsResponseBody
	GetCode() *string
	SetData(v *ListCallTagsResponseBodyData) *ListCallTagsResponseBody
	GetData() *ListCallTagsResponseBodyData
	SetHttpStatusCode(v int32) *ListCallTagsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCallTagsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCallTagsResponseBody
	GetRequestId() *string
}

type ListCallTagsResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCallTagsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListCallTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCallTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallTagsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCallTagsResponseBody) GetData() *ListCallTagsResponseBodyData {
	return s.Data
}

func (s *ListCallTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCallTagsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCallTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCallTagsResponseBody) SetCode(v string) *ListCallTagsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallTagsResponseBody) SetData(v *ListCallTagsResponseBodyData) *ListCallTagsResponseBody {
	s.Data = v
	return s
}

func (s *ListCallTagsResponseBody) SetHttpStatusCode(v int32) *ListCallTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCallTagsResponseBody) SetMessage(v string) *ListCallTagsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCallTagsResponseBody) SetRequestId(v string) *ListCallTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCallTagsResponseBodyData struct {
	List []*ListCallTagsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListCallTagsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCallTagsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallTagsResponseBodyData) GetList() []*ListCallTagsResponseBodyDataList {
	return s.List
}

func (s *ListCallTagsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallTagsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallTagsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCallTagsResponseBodyData) SetList(v []*ListCallTagsResponseBodyDataList) *ListCallTagsResponseBodyData {
	s.List = v
	return s
}

func (s *ListCallTagsResponseBodyData) SetPageNumber(v int32) *ListCallTagsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCallTagsResponseBodyData) SetPageSize(v int32) *ListCallTagsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCallTagsResponseBodyData) SetTotalCount(v int32) *ListCallTagsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCallTagsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListCallTagsResponseBodyDataList struct {
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// TagA
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListCallTagsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListCallTagsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCallTagsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallTagsResponseBodyDataList) GetTagName() *string {
	return s.TagName
}

func (s *ListCallTagsResponseBodyDataList) SetInstanceId(v string) *ListCallTagsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListCallTagsResponseBodyDataList) SetTagName(v string) *ListCallTagsResponseBodyDataList {
	s.TagName = &v
	return s
}

func (s *ListCallTagsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
