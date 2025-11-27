// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidPhoneCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidPhoneCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidPhoneCodeResponse
	GetStatusCode() *int32
	SetBody(v *ValidPhoneCodeResponseBody) *ValidPhoneCodeResponse
	GetBody() *ValidPhoneCodeResponseBody
}

type ValidPhoneCodeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidPhoneCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidPhoneCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidPhoneCodeResponse) GoString() string {
	return s.String()
}

func (s *ValidPhoneCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidPhoneCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidPhoneCodeResponse) GetBody() *ValidPhoneCodeResponseBody {
	return s.Body
}

func (s *ValidPhoneCodeResponse) SetHeaders(v map[string]*string) *ValidPhoneCodeResponse {
	s.Headers = v
	return s
}

func (s *ValidPhoneCodeResponse) SetStatusCode(v int32) *ValidPhoneCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidPhoneCodeResponse) SetBody(v *ValidPhoneCodeResponseBody) *ValidPhoneCodeResponse {
	s.Body = v
	return s
}

func (s *ValidPhoneCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
