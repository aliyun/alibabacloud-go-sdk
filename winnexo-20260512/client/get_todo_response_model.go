// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTodoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTodoResponse
	GetStatusCode() *int32
	SetBody(v *GetTodoResponseBody) *GetTodoResponse
	GetBody() *GetTodoResponseBody
}

type GetTodoResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTodoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTodoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTodoResponse) GoString() string {
	return s.String()
}

func (s *GetTodoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTodoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTodoResponse) GetBody() *GetTodoResponseBody {
	return s.Body
}

func (s *GetTodoResponse) SetHeaders(v map[string]*string) *GetTodoResponse {
	s.Headers = v
	return s
}

func (s *GetTodoResponse) SetStatusCode(v int32) *GetTodoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTodoResponse) SetBody(v *GetTodoResponseBody) *GetTodoResponse {
	s.Body = v
	return s
}

func (s *GetTodoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
