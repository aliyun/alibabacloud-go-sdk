// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnalysisTagDetailByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAnalysisTagDetailByTaskIdResponseBody
	GetCode() *string
	SetData(v []*ListAnalysisTagDetailByTaskIdResponseBodyData) *ListAnalysisTagDetailByTaskIdResponseBody
	GetData() []*ListAnalysisTagDetailByTaskIdResponseBodyData
	SetHttpStatusCode(v int32) *ListAnalysisTagDetailByTaskIdResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListAnalysisTagDetailByTaskIdResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAnalysisTagDetailByTaskIdResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAnalysisTagDetailByTaskIdResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAnalysisTagDetailByTaskIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAnalysisTagDetailByTaskIdResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListAnalysisTagDetailByTaskIdResponseBody
	GetTotalCount() *int32
}

type ListAnalysisTagDetailByTaskIdResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListAnalysisTagDetailByTaskIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// token-xxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAnalysisTagDetailByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisTagDetailByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetData() []*ListAnalysisTagDetailByTaskIdResponseBodyData {
	return s.Data
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetCode(v string) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetData(v []*ListAnalysisTagDetailByTaskIdResponseBodyData) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.Data = v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetHttpStatusCode(v int32) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetMaxResults(v int32) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetMessage(v string) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetNextToken(v string) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetRequestId(v string) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetSuccess(v bool) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetTotalCount(v int32) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAnalysisTagDetailByTaskIdResponseBodyData struct {
	// example:
	//
	// xxx
	Content     *string                                                     `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentTags []*ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags `json:"ContentTags,omitempty" xml:"ContentTags,omitempty" type:"Repeated"`
	// example:
	//
	// 112
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// summaryAndOverview
	TagTaskType *string `json:"TagTaskType,omitempty" xml:"TagTaskType,omitempty"`
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListAnalysisTagDetailByTaskIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisTagDetailByTaskIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) GetContentTags() []*ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags {
	return s.ContentTags
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) GetTagTaskType() *string {
	return s.TagTaskType
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) SetContent(v string) *ListAnalysisTagDetailByTaskIdResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) SetContentTags(v []*ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) *ListAnalysisTagDetailByTaskIdResponseBodyData {
	s.ContentTags = v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) SetId(v int64) *ListAnalysisTagDetailByTaskIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) SetTagTaskType(v string) *ListAnalysisTagDetailByTaskIdResponseBodyData {
	s.TagTaskType = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) SetTaskId(v string) *ListAnalysisTagDetailByTaskIdResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) Validate() error {
	if s.ContentTags != nil {
		for _, item := range s.ContentTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags struct {
	// example:
	//
	// xxx
	SummaryOverview *string `json:"SummaryOverview,omitempty" xml:"SummaryOverview,omitempty"`
	// example:
	//
	// xxx
	TagName *string   `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Tags    []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) GoString() string {
	return s.String()
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) GetSummaryOverview() *string {
	return s.SummaryOverview
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) GetTagName() *string {
	return s.TagName
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) GetTags() []*string {
	return s.Tags
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) SetSummaryOverview(v string) *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags {
	s.SummaryOverview = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) SetTagName(v string) *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags {
	s.TagName = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) SetTags(v []*string) *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags {
	s.Tags = v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) Validate() error {
	return dara.Validate(s)
}
