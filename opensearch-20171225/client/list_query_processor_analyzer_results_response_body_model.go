// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryProcessorAnalyzerResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListQueryProcessorAnalyzerResultsResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ListQueryProcessorAnalyzerResultsResponseBody
	GetResult() map[string]interface{}
}

type ListQueryProcessorAnalyzerResultsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 98724351-D6B2-5D8A-B089-7FFD1821A7E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListQueryProcessorAnalyzerResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorAnalyzerResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorAnalyzerResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueryProcessorAnalyzerResultsResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ListQueryProcessorAnalyzerResultsResponseBody) SetRequestId(v string) *ListQueryProcessorAnalyzerResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsResponseBody) SetResult(v map[string]interface{}) *ListQueryProcessorAnalyzerResultsResponseBody {
	s.Result = v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsResponseBody) Validate() error {
	return dara.Validate(s)
}
