// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v int64) *CancelTaskRequest
	GetBody() *int64
}

type CancelTaskRequest struct {
	Body *int64 `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelTaskRequest) GetBody() *int64 {
	return s.Body
}

func (s *CancelTaskRequest) SetBody(v int64) *CancelTaskRequest {
	s.Body = &v
	return s
}

func (s *CancelTaskRequest) Validate() error {
	return dara.Validate(s)
}
