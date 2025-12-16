// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAnalyzerEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListUserAnalyzerEntriesResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ListUserAnalyzerEntriesResponseBody
	GetResult() map[string]interface{}
}

type ListUserAnalyzerEntriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 516A02B7-2167-8D92-12D0-B639A2A0F3C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The entries of the custom analyzer. For more information, see [UserAnalyzerEntry](https://www.alibabacloud.com/help/en/open-search/industry-algorithm-edition/useranalyzerentry).
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListUserAnalyzerEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserAnalyzerEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzerEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserAnalyzerEntriesResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ListUserAnalyzerEntriesResponseBody) SetRequestId(v string) *ListUserAnalyzerEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserAnalyzerEntriesResponseBody) SetResult(v map[string]interface{}) *ListUserAnalyzerEntriesResponseBody {
	s.Result = v
	return s
}

func (s *ListUserAnalyzerEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
