// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindESUserAnalyzerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindESUserAnalyzerResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *BindESUserAnalyzerResponseBody
	GetResult() map[string]interface{}
}

type BindESUserAnalyzerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3AD34CAD-9603-5251-AFF5-3916C848A1D3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The custom analyzer.
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BindESUserAnalyzerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindESUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *BindESUserAnalyzerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindESUserAnalyzerResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *BindESUserAnalyzerResponseBody) SetRequestId(v string) *BindESUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindESUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *BindESUserAnalyzerResponseBody {
	s.Result = v
	return s
}

func (s *BindESUserAnalyzerResponseBody) Validate() error {
	return dara.Validate(s)
}
