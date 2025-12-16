// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushUserAnalyzerEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PushUserAnalyzerEntriesResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *PushUserAnalyzerEntriesResponseBody
	GetResult() map[string]interface{}
}

type PushUserAnalyzerEntriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result returned.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushUserAnalyzerEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushUserAnalyzerEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *PushUserAnalyzerEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushUserAnalyzerEntriesResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *PushUserAnalyzerEntriesResponseBody) SetRequestId(v string) *PushUserAnalyzerEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushUserAnalyzerEntriesResponseBody) SetResult(v map[string]interface{}) *PushUserAnalyzerEntriesResponseBody {
	s.Result = v
	return s
}

func (s *PushUserAnalyzerEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
