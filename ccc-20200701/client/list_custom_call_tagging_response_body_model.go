// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomCallTaggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCustomCallTaggingResponseBody
	GetCode() *string
	SetData(v *ListCustomCallTaggingResponseBodyData) *ListCustomCallTaggingResponseBody
	GetData() *ListCustomCallTaggingResponseBodyData
	SetHttpStatusCode(v int32) *ListCustomCallTaggingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCustomCallTaggingResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCustomCallTaggingResponseBody
	GetRequestId() *string
}

type ListCustomCallTaggingResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCustomCallTaggingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListCustomCallTaggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomCallTaggingResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomCallTaggingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCustomCallTaggingResponseBody) GetData() *ListCustomCallTaggingResponseBodyData {
	return s.Data
}

func (s *ListCustomCallTaggingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCustomCallTaggingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCustomCallTaggingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomCallTaggingResponseBody) SetCode(v string) *ListCustomCallTaggingResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomCallTaggingResponseBody) SetData(v *ListCustomCallTaggingResponseBodyData) *ListCustomCallTaggingResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomCallTaggingResponseBody) SetHttpStatusCode(v int32) *ListCustomCallTaggingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCustomCallTaggingResponseBody) SetMessage(v string) *ListCustomCallTaggingResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomCallTaggingResponseBody) SetRequestId(v string) *ListCustomCallTaggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomCallTaggingResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCustomCallTaggingResponseBodyData struct {
	List []*ListCustomCallTaggingResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListCustomCallTaggingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomCallTaggingResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomCallTaggingResponseBodyData) GetList() []*ListCustomCallTaggingResponseBodyDataList {
	return s.List
}

func (s *ListCustomCallTaggingResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomCallTaggingResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomCallTaggingResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomCallTaggingResponseBodyData) SetList(v []*ListCustomCallTaggingResponseBodyDataList) *ListCustomCallTaggingResponseBodyData {
	s.List = v
	return s
}

func (s *ListCustomCallTaggingResponseBodyData) SetPageNumber(v int32) *ListCustomCallTaggingResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCustomCallTaggingResponseBodyData) SetPageSize(v int32) *ListCustomCallTaggingResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCustomCallTaggingResponseBodyData) SetTotalCount(v int32) *ListCustomCallTaggingResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCustomCallTaggingResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListCustomCallTaggingResponseBodyDataList struct {
	CallTagList []*ListCustomCallTaggingResponseBodyDataListCallTagList `json:"CallTagList,omitempty" xml:"CallTagList,omitempty" type:"Repeated"`
	// example:
	//
	// agent
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1312121****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 2020-07-05 00:00:00.0
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCustomCallTaggingResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListCustomCallTaggingResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCustomCallTaggingResponseBodyDataList) GetCallTagList() []*ListCustomCallTaggingResponseBodyDataListCallTagList {
	return s.CallTagList
}

func (s *ListCustomCallTaggingResponseBodyDataList) GetCreator() *string {
	return s.Creator
}

func (s *ListCustomCallTaggingResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListCustomCallTaggingResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCustomCallTaggingResponseBodyDataList) GetNumber() *string {
	return s.Number
}

func (s *ListCustomCallTaggingResponseBodyDataList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCustomCallTaggingResponseBodyDataList) SetCallTagList(v []*ListCustomCallTaggingResponseBodyDataListCallTagList) *ListCustomCallTaggingResponseBodyDataList {
	s.CallTagList = v
	return s
}

func (s *ListCustomCallTaggingResponseBodyDataList) SetCreator(v string) *ListCustomCallTaggingResponseBodyDataList {
	s.Creator = &v
	return s
}

func (s *ListCustomCallTaggingResponseBodyDataList) SetDescription(v string) *ListCustomCallTaggingResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListCustomCallTaggingResponseBodyDataList) SetInstanceId(v string) *ListCustomCallTaggingResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListCustomCallTaggingResponseBodyDataList) SetNumber(v string) *ListCustomCallTaggingResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListCustomCallTaggingResponseBodyDataList) SetUpdateTime(v string) *ListCustomCallTaggingResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListCustomCallTaggingResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ListCustomCallTaggingResponseBodyDataListCallTagList struct {
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// TagA
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListCustomCallTaggingResponseBodyDataListCallTagList) String() string {
	return dara.Prettify(s)
}

func (s ListCustomCallTaggingResponseBodyDataListCallTagList) GoString() string {
	return s.String()
}

func (s *ListCustomCallTaggingResponseBodyDataListCallTagList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCustomCallTaggingResponseBodyDataListCallTagList) GetTagName() *string {
	return s.TagName
}

func (s *ListCustomCallTaggingResponseBodyDataListCallTagList) SetInstanceId(v string) *ListCustomCallTaggingResponseBodyDataListCallTagList {
	s.InstanceId = &v
	return s
}

func (s *ListCustomCallTaggingResponseBodyDataListCallTagList) SetTagName(v string) *ListCustomCallTaggingResponseBodyDataListCallTagList {
	s.TagName = &v
	return s
}

func (s *ListCustomCallTaggingResponseBodyDataListCallTagList) Validate() error {
	return dara.Validate(s)
}
