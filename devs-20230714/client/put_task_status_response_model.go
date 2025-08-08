// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *Task) *PutTaskStatusResponse
	GetBody() *Task
}

type PutTaskStatusResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s PutTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *PutTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutTaskStatusResponse) GetBody() *Task {
	return s.Body
}

func (s *PutTaskStatusResponse) SetHeaders(v map[string]*string) *PutTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *PutTaskStatusResponse) SetStatusCode(v int32) *PutTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PutTaskStatusResponse) SetBody(v *Task) *PutTaskStatusResponse {
	s.Body = v
	return s
}

func (s *PutTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
