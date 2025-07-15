// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicRecommendEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTopicRecommendEventListResponseBody
	GetCode() *string
	SetData(v []*string) *ListTopicRecommendEventListResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *ListTopicRecommendEventListResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListTopicRecommendEventListResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListTopicRecommendEventListResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTopicRecommendEventListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTopicRecommendEventListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTopicRecommendEventListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListTopicRecommendEventListResponseBody
	GetTotalCount() *int32
}

type ListTopicRecommendEventListResponseBody struct {
	// example:
	//
	// NoData
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 71
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// x\\"x\\"x
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
	// 60
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTopicRecommendEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTopicRecommendEventListResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicRecommendEventListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTopicRecommendEventListResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListTopicRecommendEventListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTopicRecommendEventListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTopicRecommendEventListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTopicRecommendEventListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTopicRecommendEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTopicRecommendEventListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTopicRecommendEventListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTopicRecommendEventListResponseBody) SetCode(v string) *ListTopicRecommendEventListResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetData(v []*string) *ListTopicRecommendEventListResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetHttpStatusCode(v int32) *ListTopicRecommendEventListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetMaxResults(v int32) *ListTopicRecommendEventListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetMessage(v string) *ListTopicRecommendEventListResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetNextToken(v string) *ListTopicRecommendEventListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetRequestId(v string) *ListTopicRecommendEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetSuccess(v bool) *ListTopicRecommendEventListResponseBody {
	s.Success = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetTotalCount(v int32) *ListTopicRecommendEventListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) Validate() error {
	return dara.Validate(s)
}
