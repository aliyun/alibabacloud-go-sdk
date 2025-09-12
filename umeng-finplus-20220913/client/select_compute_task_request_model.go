// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectComputeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v int64) *SelectComputeTaskRequest
	GetBody() *int64
}

type SelectComputeTaskRequest struct {
	Body *int64 `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectComputeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTaskRequest) GoString() string {
	return s.String()
}

func (s *SelectComputeTaskRequest) GetBody() *int64 {
	return s.Body
}

func (s *SelectComputeTaskRequest) SetBody(v int64) *SelectComputeTaskRequest {
	s.Body = &v
	return s
}

func (s *SelectComputeTaskRequest) Validate() error {
	return dara.Validate(s)
}
