// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDingtalkPersonalTodoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDingtalkPersonalTodoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDingtalkPersonalTodoTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDingtalkPersonalTodoTaskResponseBody) *CreateDingtalkPersonalTodoTaskResponse
	GetBody() *CreateDingtalkPersonalTodoTaskResponseBody
}

type CreateDingtalkPersonalTodoTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDingtalkPersonalTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDingtalkPersonalTodoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDingtalkPersonalTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDingtalkPersonalTodoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDingtalkPersonalTodoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDingtalkPersonalTodoTaskResponse) GetBody() *CreateDingtalkPersonalTodoTaskResponseBody {
	return s.Body
}

func (s *CreateDingtalkPersonalTodoTaskResponse) SetHeaders(v map[string]*string) *CreateDingtalkPersonalTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskResponse) SetStatusCode(v int32) *CreateDingtalkPersonalTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskResponse) SetBody(v *CreateDingtalkPersonalTodoTaskResponseBody) *CreateDingtalkPersonalTodoTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
