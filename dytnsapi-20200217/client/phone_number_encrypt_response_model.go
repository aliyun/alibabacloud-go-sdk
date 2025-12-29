// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberEncryptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PhoneNumberEncryptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PhoneNumberEncryptResponse
	GetStatusCode() *int32
	SetBody(v *PhoneNumberEncryptResponseBody) *PhoneNumberEncryptResponse
	GetBody() *PhoneNumberEncryptResponseBody
}

type PhoneNumberEncryptResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PhoneNumberEncryptResponse) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberEncryptResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PhoneNumberEncryptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PhoneNumberEncryptResponse) GetBody() *PhoneNumberEncryptResponseBody {
	return s.Body
}

func (s *PhoneNumberEncryptResponse) SetHeaders(v map[string]*string) *PhoneNumberEncryptResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberEncryptResponse) SetStatusCode(v int32) *PhoneNumberEncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberEncryptResponse) SetBody(v *PhoneNumberEncryptResponseBody) *PhoneNumberEncryptResponse {
	s.Body = v
	return s
}

func (s *PhoneNumberEncryptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
