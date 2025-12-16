// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindESUserAnalyzerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindESUserAnalyzerResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *UnbindESUserAnalyzerResponseBody
	GetResult() map[string]interface{}
}

type UnbindESUserAnalyzerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The custom analyzer.
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UnbindESUserAnalyzerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindESUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindESUserAnalyzerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindESUserAnalyzerResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *UnbindESUserAnalyzerResponseBody) SetRequestId(v string) *UnbindESUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindESUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *UnbindESUserAnalyzerResponseBody {
	s.Result = v
	return s
}

func (s *UnbindESUserAnalyzerResponseBody) Validate() error {
	return dara.Validate(s)
}
