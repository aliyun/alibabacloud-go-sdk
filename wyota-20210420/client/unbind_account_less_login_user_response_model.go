// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAccountLessLoginUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindAccountLessLoginUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindAccountLessLoginUserResponse
	GetStatusCode() *int32
	SetBody(v *UnbindAccountLessLoginUserResponseBody) *UnbindAccountLessLoginUserResponse
	GetBody() *UnbindAccountLessLoginUserResponseBody
}

type UnbindAccountLessLoginUserResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindAccountLessLoginUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindAccountLessLoginUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindAccountLessLoginUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindAccountLessLoginUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindAccountLessLoginUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindAccountLessLoginUserResponse) GetBody() *UnbindAccountLessLoginUserResponseBody {
	return s.Body
}

func (s *UnbindAccountLessLoginUserResponse) SetHeaders(v map[string]*string) *UnbindAccountLessLoginUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindAccountLessLoginUserResponse) SetStatusCode(v int32) *UnbindAccountLessLoginUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAccountLessLoginUserResponse) SetBody(v *UnbindAccountLessLoginUserResponseBody) *UnbindAccountLessLoginUserResponse {
	s.Body = v
	return s
}

func (s *UnbindAccountLessLoginUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
