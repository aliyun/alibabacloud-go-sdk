// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBenchmarkTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateBenchmarkTaskRequest
	GetBody() *string
}

type UpdateBenchmarkTaskRequest struct {
	// The request body. The body includes the parameters that are set to create a stress testing task. For more information, see **Table 1. Fields in the base parameter**.
	//
	// example:
	//
	// {
	//
	//     "base":  {
	//
	//          "qps": 200
	//
	//     }
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBenchmarkTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateBenchmarkTaskRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateBenchmarkTaskRequest) SetBody(v string) *UpdateBenchmarkTaskRequest {
	s.Body = &v
	return s
}

func (s *UpdateBenchmarkTaskRequest) Validate() error {
	return dara.Validate(s)
}
