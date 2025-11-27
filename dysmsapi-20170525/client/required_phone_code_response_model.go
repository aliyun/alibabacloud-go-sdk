// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequiredPhoneCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RequiredPhoneCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RequiredPhoneCodeResponse
	GetStatusCode() *int32
	SetBody(v *RequiredPhoneCodeResponseBody) *RequiredPhoneCodeResponse
	GetBody() *RequiredPhoneCodeResponseBody
}

type RequiredPhoneCodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequiredPhoneCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequiredPhoneCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RequiredPhoneCodeResponse) GoString() string {
	return s.String()
}

func (s *RequiredPhoneCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RequiredPhoneCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RequiredPhoneCodeResponse) GetBody() *RequiredPhoneCodeResponseBody {
	return s.Body
}

func (s *RequiredPhoneCodeResponse) SetHeaders(v map[string]*string) *RequiredPhoneCodeResponse {
	s.Headers = v
	return s
}

func (s *RequiredPhoneCodeResponse) SetStatusCode(v int32) *RequiredPhoneCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RequiredPhoneCodeResponse) SetBody(v *RequiredPhoneCodeResponseBody) *RequiredPhoneCodeResponse {
	s.Body = v
	return s
}

func (s *RequiredPhoneCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
