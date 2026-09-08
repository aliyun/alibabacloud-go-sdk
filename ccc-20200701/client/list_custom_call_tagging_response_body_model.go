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
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *ListCustomCallTaggingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomCallTaggingResponseBodyData struct {
	// The list of inbound number marks.
	List []*ListCustomCallTaggingResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number, ranging from 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size, ranging from 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total count.
	//
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

type ListCustomCallTaggingResponseBodyDataList struct {
	// List of number tags.
	CallTagList []*ListCustomCallTaggingResponseBodyDataListCallTagList `json:"CallTagList,omitempty" xml:"CallTagList,omitempty" type:"Repeated"`
	// Creator.
	//
	// example:
	//
	// agent
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The description of the inbound number mark.
	//
	// example:
	//
	// 王先生
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the inbound number mark.
	//
	// example:
	//
	// 1312121****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// Last update time.
	//
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
	if s.CallTagList != nil {
		for _, item := range s.CallTagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomCallTaggingResponseBodyDataListCallTagList struct {
	// Instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Number tag name.
	//
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
