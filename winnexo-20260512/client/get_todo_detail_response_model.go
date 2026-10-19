// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTodoDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTodoDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetTodoDetailResponseBody) *GetTodoDetailResponse
	GetBody() *GetTodoDetailResponseBody
}

type GetTodoDetailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTodoDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTodoDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTodoDetailResponse) GoString() string {
	return s.String()
}

func (s *GetTodoDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTodoDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTodoDetailResponse) GetBody() *GetTodoDetailResponseBody {
	return s.Body
}

func (s *GetTodoDetailResponse) SetHeaders(v map[string]*string) *GetTodoDetailResponse {
	s.Headers = v
	return s
}

func (s *GetTodoDetailResponse) SetStatusCode(v int32) *GetTodoDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTodoDetailResponse) SetBody(v *GetTodoDetailResponseBody) *GetTodoDetailResponse {
	s.Body = v
	return s
}

func (s *GetTodoDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
