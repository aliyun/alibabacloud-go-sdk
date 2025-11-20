// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskExecutorStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTodoTaskExecutorStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTodoTaskExecutorStatusResponseBody) *UpdateTodoTaskExecutorStatusResponse
	GetBody() *UpdateTodoTaskExecutorStatusResponseBody
}

type UpdateTodoTaskExecutorStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTodoTaskExecutorStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTodoTaskExecutorStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTodoTaskExecutorStatusResponse) GetBody() *UpdateTodoTaskExecutorStatusResponseBody {
	return s.Body
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetStatusCode(v int32) *UpdateTodoTaskExecutorStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetBody(v *UpdateTodoTaskExecutorStatusResponseBody) *UpdateTodoTaskExecutorStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
