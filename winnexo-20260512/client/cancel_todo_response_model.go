// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTodoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelTodoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelTodoResponse
	GetStatusCode() *int32
	SetBody(v *CancelTodoResponseBody) *CancelTodoResponse
	GetBody() *CancelTodoResponseBody
}

type CancelTodoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelTodoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelTodoResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelTodoResponse) GoString() string {
	return s.String()
}

func (s *CancelTodoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelTodoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelTodoResponse) GetBody() *CancelTodoResponseBody {
	return s.Body
}

func (s *CancelTodoResponse) SetHeaders(v map[string]*string) *CancelTodoResponse {
	s.Headers = v
	return s
}

func (s *CancelTodoResponse) SetStatusCode(v int32) *CancelTodoResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelTodoResponse) SetBody(v *CancelTodoResponseBody) *CancelTodoResponse {
	s.Body = v
	return s
}

func (s *CancelTodoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
