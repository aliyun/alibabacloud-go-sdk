// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskResponseBody) *UpdateTaskResponse
	GetBody() *UpdateTaskResponseBody
}

type UpdateTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskResponse) GetBody() *UpdateTaskResponseBody {
	return s.Body
}

func (s *UpdateTaskResponse) SetHeaders(v map[string]*string) *UpdateTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskResponse) SetStatusCode(v int32) *UpdateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskResponse) SetBody(v *UpdateTaskResponseBody) *UpdateTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskResponse) Validate() error {
	return dara.Validate(s)
}
