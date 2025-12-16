// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserAnalyzerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserAnalyzerResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CreateUserAnalyzerResponseBody
	GetResult() map[string]interface{}
}

type CreateUserAnalyzerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 98724351-D6B2-5D8A-B089-7FFD1821A7E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The custom analyzer.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateUserAnalyzerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserAnalyzerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserAnalyzerResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CreateUserAnalyzerResponseBody) SetRequestId(v string) *CreateUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *CreateUserAnalyzerResponseBody {
	s.Result = v
	return s
}

func (s *CreateUserAnalyzerResponseBody) Validate() error {
	return dara.Validate(s)
}
