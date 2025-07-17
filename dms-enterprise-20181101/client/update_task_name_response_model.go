// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskNameResponseBody) *UpdateTaskNameResponse
	GetBody() *UpdateTaskNameResponseBody
}

type UpdateTaskNameResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskNameResponse) GetBody() *UpdateTaskNameResponseBody {
	return s.Body
}

func (s *UpdateTaskNameResponse) SetHeaders(v map[string]*string) *UpdateTaskNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskNameResponse) SetStatusCode(v int32) *UpdateTaskNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskNameResponse) SetBody(v *UpdateTaskNameResponseBody) *UpdateTaskNameResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskNameResponse) Validate() error {
	return dara.Validate(s)
}
