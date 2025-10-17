// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotTopicSummariesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHotTopicSummariesResponseBody
	GetCode() *string
	SetData(v []*ListHotTopicSummariesResponseBodyData) *ListHotTopicSummariesResponseBody
	GetData() []*ListHotTopicSummariesResponseBodyData
	SetHttpStatusCode(v int32) *ListHotTopicSummariesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListHotTopicSummariesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListHotTopicSummariesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListHotTopicSummariesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListHotTopicSummariesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHotTopicSummariesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListHotTopicSummariesResponseBody
	GetTotalCount() *int32
}

type ListHotTopicSummariesResponseBody struct {
	// example:
	//
	// xx
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListHotTopicSummariesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// JlroP3CjgQh5PQDlH3ArzADkBTPZgVqo+64jhZRglNq0mEYoV5SlGb/Juvo8CdfYE9rlwEr2pIJQwdaYotak9g==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListHotTopicSummariesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicSummariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHotTopicSummariesResponseBody) GetData() []*ListHotTopicSummariesResponseBodyData {
	return s.Data
}

func (s *ListHotTopicSummariesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHotTopicSummariesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotTopicSummariesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotTopicSummariesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotTopicSummariesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotTopicSummariesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHotTopicSummariesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHotTopicSummariesResponseBody) SetCode(v string) *ListHotTopicSummariesResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetData(v []*ListHotTopicSummariesResponseBodyData) *ListHotTopicSummariesResponseBody {
	s.Data = v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetHttpStatusCode(v int32) *ListHotTopicSummariesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetMaxResults(v int32) *ListHotTopicSummariesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetMessage(v string) *ListHotTopicSummariesResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetNextToken(v string) *ListHotTopicSummariesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetRequestId(v string) *ListHotTopicSummariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetSuccess(v bool) *ListHotTopicSummariesResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) SetTotalCount(v int32) *ListHotTopicSummariesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHotTopicSummariesResponseBody) Validate() error {
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

type ListHotTopicSummariesResponseBodyData struct {
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// xx
	HotTopic *string `json:"hotTopic,omitempty" xml:"hotTopic,omitempty"`
	// example:
	//
	// 2024-09-13_12
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// example:
	//
	// 1000000
	HotValue *float64 `json:"hotValue,omitempty" xml:"hotValue,omitempty"`
	// example:
	//
	// db5dc5b3d8954a30b65ba700c9dda3bb
	Id      *string                                       `json:"id,omitempty" xml:"id,omitempty"`
	News    []*ListHotTopicSummariesResponseBodyDataNews  `json:"news,omitempty" xml:"news,omitempty" type:"Repeated"`
	Summary *ListHotTopicSummariesResponseBodyDataSummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
	// example:
	//
	// xx
	TextSummary *string `json:"textSummary,omitempty" xml:"textSummary,omitempty"`
}

func (s ListHotTopicSummariesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicSummariesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *ListHotTopicSummariesResponseBodyData) GetHotTopic() *string {
	return s.HotTopic
}

func (s *ListHotTopicSummariesResponseBodyData) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *ListHotTopicSummariesResponseBodyData) GetHotValue() *float64 {
	return s.HotValue
}

func (s *ListHotTopicSummariesResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListHotTopicSummariesResponseBodyData) GetNews() []*ListHotTopicSummariesResponseBodyDataNews {
	return s.News
}

func (s *ListHotTopicSummariesResponseBodyData) GetSummary() *ListHotTopicSummariesResponseBodyDataSummary {
	return s.Summary
}

func (s *ListHotTopicSummariesResponseBodyData) GetTextSummary() *string {
	return s.TextSummary
}

func (s *ListHotTopicSummariesResponseBodyData) SetCategory(v string) *ListHotTopicSummariesResponseBodyData {
	s.Category = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetHotTopic(v string) *ListHotTopicSummariesResponseBodyData {
	s.HotTopic = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetHotTopicVersion(v string) *ListHotTopicSummariesResponseBodyData {
	s.HotTopicVersion = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetHotValue(v float64) *ListHotTopicSummariesResponseBodyData {
	s.HotValue = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetId(v string) *ListHotTopicSummariesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetNews(v []*ListHotTopicSummariesResponseBodyDataNews) *ListHotTopicSummariesResponseBodyData {
	s.News = v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetSummary(v *ListHotTopicSummariesResponseBodyDataSummary) *ListHotTopicSummariesResponseBodyData {
	s.Summary = v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) SetTextSummary(v string) *ListHotTopicSummariesResponseBodyData {
	s.TextSummary = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyData) Validate() error {
	if s.News != nil {
		for _, item := range s.News {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotTopicSummariesResponseBodyDataNews struct {
	Comments []*ListHotTopicSummariesResponseBodyDataNewsComments `json:"comments,omitempty" xml:"comments,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2024-09-10 15:32:00
	PubTime *string `json:"pubTime,omitempty" xml:"pubTime,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// url
	//
	// example:
	//
	// http://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListHotTopicSummariesResponseBodyDataNews) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicSummariesResponseBodyDataNews) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBodyDataNews) GetComments() []*ListHotTopicSummariesResponseBodyDataNewsComments {
	return s.Comments
}

func (s *ListHotTopicSummariesResponseBodyDataNews) GetContent() *string {
	return s.Content
}

func (s *ListHotTopicSummariesResponseBodyDataNews) GetPubTime() *string {
	return s.PubTime
}

func (s *ListHotTopicSummariesResponseBodyDataNews) GetTitle() *string {
	return s.Title
}

func (s *ListHotTopicSummariesResponseBodyDataNews) GetUrl() *string {
	return s.Url
}

func (s *ListHotTopicSummariesResponseBodyDataNews) SetComments(v []*ListHotTopicSummariesResponseBodyDataNewsComments) *ListHotTopicSummariesResponseBodyDataNews {
	s.Comments = v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataNews) SetContent(v string) *ListHotTopicSummariesResponseBodyDataNews {
	s.Content = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataNews) SetPubTime(v string) *ListHotTopicSummariesResponseBodyDataNews {
	s.PubTime = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataNews) SetTitle(v string) *ListHotTopicSummariesResponseBodyDataNews {
	s.Title = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataNews) SetUrl(v string) *ListHotTopicSummariesResponseBodyDataNews {
	s.Url = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataNews) Validate() error {
	if s.Comments != nil {
		for _, item := range s.Comments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotTopicSummariesResponseBodyDataNewsComments struct {
	// example:
	//
	// xx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ListHotTopicSummariesResponseBodyDataNewsComments) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicSummariesResponseBodyDataNewsComments) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBodyDataNewsComments) GetText() *string {
	return s.Text
}

func (s *ListHotTopicSummariesResponseBodyDataNewsComments) SetText(v string) *ListHotTopicSummariesResponseBodyDataNewsComments {
	s.Text = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataNewsComments) Validate() error {
	return dara.Validate(s)
}

type ListHotTopicSummariesResponseBodyDataSummary struct {
	Summaries []*ListHotTopicSummariesResponseBodyDataSummarySummaries `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
}

func (s ListHotTopicSummariesResponseBodyDataSummary) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicSummariesResponseBodyDataSummary) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBodyDataSummary) GetSummaries() []*ListHotTopicSummariesResponseBodyDataSummarySummaries {
	return s.Summaries
}

func (s *ListHotTopicSummariesResponseBodyDataSummary) SetSummaries(v []*ListHotTopicSummariesResponseBodyDataSummarySummaries) *ListHotTopicSummariesResponseBodyDataSummary {
	s.Summaries = v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataSummary) Validate() error {
	if s.Summaries != nil {
		for _, item := range s.Summaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotTopicSummariesResponseBodyDataSummarySummaries struct {
	// example:
	//
	// xx
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListHotTopicSummariesResponseBodyDataSummarySummaries) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicSummariesResponseBodyDataSummarySummaries) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponseBodyDataSummarySummaries) GetSummary() *string {
	return s.Summary
}

func (s *ListHotTopicSummariesResponseBodyDataSummarySummaries) GetTitle() *string {
	return s.Title
}

func (s *ListHotTopicSummariesResponseBodyDataSummarySummaries) SetSummary(v string) *ListHotTopicSummariesResponseBodyDataSummarySummaries {
	s.Summary = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataSummarySummaries) SetTitle(v string) *ListHotTopicSummariesResponseBodyDataSummarySummaries {
	s.Title = &v
	return s
}

func (s *ListHotTopicSummariesResponseBodyDataSummarySummaries) Validate() error {
	return dara.Validate(s)
}
