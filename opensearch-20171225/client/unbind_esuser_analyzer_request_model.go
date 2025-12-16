// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindESUserAnalyzerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v interface{}) *UnbindESUserAnalyzerRequest
	GetBody() interface{}
}

type UnbindESUserAnalyzerRequest struct {
	// The request parameters.
	//
	// example:
	//
	// {
	//
	//   "name": "kevintest-analyzer"
	//
	// }
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindESUserAnalyzerRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindESUserAnalyzerRequest) GoString() string {
	return s.String()
}

func (s *UnbindESUserAnalyzerRequest) GetBody() interface{} {
	return s.Body
}

func (s *UnbindESUserAnalyzerRequest) SetBody(v interface{}) *UnbindESUserAnalyzerRequest {
	s.Body = v
	return s
}

func (s *UnbindESUserAnalyzerRequest) Validate() error {
	return dara.Validate(s)
}
