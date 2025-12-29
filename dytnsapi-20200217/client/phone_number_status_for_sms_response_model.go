// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForSmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PhoneNumberStatusForSmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PhoneNumberStatusForSmsResponse
	GetStatusCode() *int32
	SetBody(v *PhoneNumberStatusForSmsResponseBody) *PhoneNumberStatusForSmsResponse
	GetBody() *PhoneNumberStatusForSmsResponseBody
}

type PhoneNumberStatusForSmsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberStatusForSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PhoneNumberStatusForSmsResponse) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForSmsResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PhoneNumberStatusForSmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PhoneNumberStatusForSmsResponse) GetBody() *PhoneNumberStatusForSmsResponseBody {
	return s.Body
}

func (s *PhoneNumberStatusForSmsResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForSmsResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForSmsResponse) SetStatusCode(v int32) *PhoneNumberStatusForSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponse) SetBody(v *PhoneNumberStatusForSmsResponseBody) *PhoneNumberStatusForSmsResponse {
	s.Body = v
	return s
}

func (s *PhoneNumberStatusForSmsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
