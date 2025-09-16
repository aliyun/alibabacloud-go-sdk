// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateTaskResponseBody) *CreateTaskResponse
	GetBody() *CreateTaskResponseBody
}

type CreateTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTaskResponse) GetBody() *CreateTaskResponseBody {
	return s.Body
}

func (s *CreateTaskResponse) SetHeaders(v map[string]*string) *CreateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskResponse) SetStatusCode(v int32) *CreateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskResponse) SetBody(v *CreateTaskResponseBody) *CreateTaskResponse {
	s.Body = v
	return s
}

func (s *CreateTaskResponse) Validate() error {
	return dara.Validate(s)
}
