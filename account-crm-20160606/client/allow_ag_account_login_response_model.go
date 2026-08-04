// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllowAgAccountLoginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllowAgAccountLoginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllowAgAccountLoginResponse
	GetStatusCode() *int32
	SetBody(v *AllowAgAccountLoginResponseBody) *AllowAgAccountLoginResponse
	GetBody() *AllowAgAccountLoginResponseBody
}

type AllowAgAccountLoginResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllowAgAccountLoginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllowAgAccountLoginResponse) String() string {
	return dara.Prettify(s)
}

func (s AllowAgAccountLoginResponse) GoString() string {
	return s.String()
}

func (s *AllowAgAccountLoginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllowAgAccountLoginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllowAgAccountLoginResponse) GetBody() *AllowAgAccountLoginResponseBody {
	return s.Body
}

func (s *AllowAgAccountLoginResponse) SetHeaders(v map[string]*string) *AllowAgAccountLoginResponse {
	s.Headers = v
	return s
}

func (s *AllowAgAccountLoginResponse) SetStatusCode(v int32) *AllowAgAccountLoginResponse {
	s.StatusCode = &v
	return s
}

func (s *AllowAgAccountLoginResponse) SetBody(v *AllowAgAccountLoginResponseBody) *AllowAgAccountLoginResponse {
	s.Body = v
	return s
}

func (s *AllowAgAccountLoginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
