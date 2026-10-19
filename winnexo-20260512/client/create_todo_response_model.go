// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTodoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTodoResponse
	GetStatusCode() *int32
	SetBody(v *CreateTodoResponseBody) *CreateTodoResponse
	GetBody() *CreateTodoResponseBody
}

type CreateTodoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTodoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTodoResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoResponse) GoString() string {
	return s.String()
}

func (s *CreateTodoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTodoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTodoResponse) GetBody() *CreateTodoResponseBody {
	return s.Body
}

func (s *CreateTodoResponse) SetHeaders(v map[string]*string) *CreateTodoResponse {
	s.Headers = v
	return s
}

func (s *CreateTodoResponse) SetStatusCode(v int32) *CreateTodoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTodoResponse) SetBody(v *CreateTodoResponseBody) *CreateTodoResponse {
	s.Body = v
	return s
}

func (s *CreateTodoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
