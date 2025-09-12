// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDataSetTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v int64) *SubmitDataSetTaskRequest
	GetBody() *int64
}

type SubmitDataSetTaskRequest struct {
	Body *int64 `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDataSetTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDataSetTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitDataSetTaskRequest) GetBody() *int64 {
	return s.Body
}

func (s *SubmitDataSetTaskRequest) SetBody(v int64) *SubmitDataSetTaskRequest {
	s.Body = &v
	return s
}

func (s *SubmitDataSetTaskRequest) Validate() error {
	return dara.Validate(s)
}
