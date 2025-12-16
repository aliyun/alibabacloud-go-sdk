// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindESUserAnalyzerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v interface{}) *BindESUserAnalyzerRequest
	GetBody() interface{}
}

type BindESUserAnalyzerRequest struct {
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

func (s BindESUserAnalyzerRequest) String() string {
	return dara.Prettify(s)
}

func (s BindESUserAnalyzerRequest) GoString() string {
	return s.String()
}

func (s *BindESUserAnalyzerRequest) GetBody() interface{} {
	return s.Body
}

func (s *BindESUserAnalyzerRequest) SetBody(v interface{}) *BindESUserAnalyzerRequest {
	s.Body = v
	return s
}

func (s *BindESUserAnalyzerRequest) Validate() error {
	return dara.Validate(s)
}
