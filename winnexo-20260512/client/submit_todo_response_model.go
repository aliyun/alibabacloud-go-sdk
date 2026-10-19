// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTodoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTodoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTodoResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTodoResponseBody) *SubmitTodoResponse
	GetBody() *SubmitTodoResponseBody
}

type SubmitTodoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTodoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTodoResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTodoResponse) GoString() string {
	return s.String()
}

func (s *SubmitTodoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTodoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTodoResponse) GetBody() *SubmitTodoResponseBody {
	return s.Body
}

func (s *SubmitTodoResponse) SetHeaders(v map[string]*string) *SubmitTodoResponse {
	s.Headers = v
	return s
}

func (s *SubmitTodoResponse) SetStatusCode(v int32) *SubmitTodoResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTodoResponse) SetBody(v *SubmitTodoResponseBody) *SubmitTodoResponse {
	s.Body = v
	return s
}

func (s *SubmitTodoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
