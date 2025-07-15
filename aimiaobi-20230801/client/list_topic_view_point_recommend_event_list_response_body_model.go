// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicViewPointRecommendEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTopicViewPointRecommendEventListResponseBody
	GetCode() *string
	SetData(v []*string) *ListTopicViewPointRecommendEventListResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *ListTopicViewPointRecommendEventListResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListTopicViewPointRecommendEventListResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListTopicViewPointRecommendEventListResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTopicViewPointRecommendEventListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTopicViewPointRecommendEventListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTopicViewPointRecommendEventListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListTopicViewPointRecommendEventListResponseBody
	GetTotalCount() *int32
}

type ListTopicViewPointRecommendEventListResponseBody struct {
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
	// 8
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
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
	// 32
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTopicViewPointRecommendEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTopicViewPointRecommendEventListResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicViewPointRecommendEventListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTopicViewPointRecommendEventListResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListTopicViewPointRecommendEventListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTopicViewPointRecommendEventListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTopicViewPointRecommendEventListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTopicViewPointRecommendEventListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTopicViewPointRecommendEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTopicViewPointRecommendEventListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTopicViewPointRecommendEventListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetCode(v string) *ListTopicViewPointRecommendEventListResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetData(v []*string) *ListTopicViewPointRecommendEventListResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetHttpStatusCode(v int32) *ListTopicViewPointRecommendEventListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetMaxResults(v int32) *ListTopicViewPointRecommendEventListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetMessage(v string) *ListTopicViewPointRecommendEventListResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetNextToken(v string) *ListTopicViewPointRecommendEventListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetRequestId(v string) *ListTopicViewPointRecommendEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetSuccess(v bool) *ListTopicViewPointRecommendEventListResponseBody {
	s.Success = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetTotalCount(v int32) *ListTopicViewPointRecommendEventListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) Validate() error {
	return dara.Validate(s)
}
