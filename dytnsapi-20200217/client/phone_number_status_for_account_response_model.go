// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PhoneNumberStatusForAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PhoneNumberStatusForAccountResponse
	GetStatusCode() *int32
	SetBody(v *PhoneNumberStatusForAccountResponseBody) *PhoneNumberStatusForAccountResponse
	GetBody() *PhoneNumberStatusForAccountResponseBody
}

type PhoneNumberStatusForAccountResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberStatusForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PhoneNumberStatusForAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForAccountResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PhoneNumberStatusForAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PhoneNumberStatusForAccountResponse) GetBody() *PhoneNumberStatusForAccountResponseBody {
	return s.Body
}

func (s *PhoneNumberStatusForAccountResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForAccountResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForAccountResponse) SetStatusCode(v int32) *PhoneNumberStatusForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponse) SetBody(v *PhoneNumberStatusForAccountResponseBody) *PhoneNumberStatusForAccountResponse {
	s.Body = v
	return s
}

func (s *PhoneNumberStatusForAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
