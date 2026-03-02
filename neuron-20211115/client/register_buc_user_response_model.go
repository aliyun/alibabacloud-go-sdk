// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterBucUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterBucUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterBucUserResponse
	GetStatusCode() *int32
	SetBody(v *RegisterBucUserResponseBody) *RegisterBucUserResponse
	GetBody() *RegisterBucUserResponseBody
}

type RegisterBucUserResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterBucUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterBucUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterBucUserResponse) GoString() string {
	return s.String()
}

func (s *RegisterBucUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterBucUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterBucUserResponse) GetBody() *RegisterBucUserResponseBody {
	return s.Body
}

func (s *RegisterBucUserResponse) SetHeaders(v map[string]*string) *RegisterBucUserResponse {
	s.Headers = v
	return s
}

func (s *RegisterBucUserResponse) SetStatusCode(v int32) *RegisterBucUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterBucUserResponse) SetBody(v *RegisterBucUserResponseBody) *RegisterBucUserResponse {
	s.Body = v
	return s
}

func (s *RegisterBucUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
