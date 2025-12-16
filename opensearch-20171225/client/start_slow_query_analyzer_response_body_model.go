// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSlowQueryAnalyzerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartSlowQueryAnalyzerResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *StartSlowQueryAnalyzerResponseBody
	GetResult() map[string]interface{}
}

type StartSlowQueryAnalyzerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StartSlowQueryAnalyzerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSlowQueryAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *StartSlowQueryAnalyzerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSlowQueryAnalyzerResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *StartSlowQueryAnalyzerResponseBody) SetRequestId(v string) *StartSlowQueryAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSlowQueryAnalyzerResponseBody) SetResult(v map[string]interface{}) *StartSlowQueryAnalyzerResponseBody {
	s.Result = v
	return s
}

func (s *StartSlowQueryAnalyzerResponseBody) Validate() error {
	return dara.Validate(s)
}
