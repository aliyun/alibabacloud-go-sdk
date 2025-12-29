// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberConvertServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PhoneNumberConvertServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PhoneNumberConvertServiceResponse
	GetStatusCode() *int32
	SetBody(v *PhoneNumberConvertServiceResponseBody) *PhoneNumberConvertServiceResponse
	GetBody() *PhoneNumberConvertServiceResponseBody
}

type PhoneNumberConvertServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberConvertServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PhoneNumberConvertServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberConvertServiceResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberConvertServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PhoneNumberConvertServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PhoneNumberConvertServiceResponse) GetBody() *PhoneNumberConvertServiceResponseBody {
	return s.Body
}

func (s *PhoneNumberConvertServiceResponse) SetHeaders(v map[string]*string) *PhoneNumberConvertServiceResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberConvertServiceResponse) SetStatusCode(v int32) *PhoneNumberConvertServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberConvertServiceResponse) SetBody(v *PhoneNumberConvertServiceResponseBody) *PhoneNumberConvertServiceResponse {
	s.Body = v
	return s
}

func (s *PhoneNumberConvertServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
