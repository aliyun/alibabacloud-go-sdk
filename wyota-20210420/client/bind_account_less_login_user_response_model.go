// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAccountLessLoginUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAccountLessLoginUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAccountLessLoginUserResponse
	GetStatusCode() *int32
	SetBody(v *BindAccountLessLoginUserResponseBody) *BindAccountLessLoginUserResponse
	GetBody() *BindAccountLessLoginUserResponseBody
}

type BindAccountLessLoginUserResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAccountLessLoginUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAccountLessLoginUserResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAccountLessLoginUserResponse) GoString() string {
	return s.String()
}

func (s *BindAccountLessLoginUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAccountLessLoginUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAccountLessLoginUserResponse) GetBody() *BindAccountLessLoginUserResponseBody {
	return s.Body
}

func (s *BindAccountLessLoginUserResponse) SetHeaders(v map[string]*string) *BindAccountLessLoginUserResponse {
	s.Headers = v
	return s
}

func (s *BindAccountLessLoginUserResponse) SetStatusCode(v int32) *BindAccountLessLoginUserResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAccountLessLoginUserResponse) SetBody(v *BindAccountLessLoginUserResponseBody) *BindAccountLessLoginUserResponse {
	s.Body = v
	return s
}

func (s *BindAccountLessLoginUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
