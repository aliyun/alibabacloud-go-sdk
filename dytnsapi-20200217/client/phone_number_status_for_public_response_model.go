// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForPublicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PhoneNumberStatusForPublicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PhoneNumberStatusForPublicResponse
	GetStatusCode() *int32
	SetBody(v *PhoneNumberStatusForPublicResponseBody) *PhoneNumberStatusForPublicResponse
	GetBody() *PhoneNumberStatusForPublicResponseBody
}

type PhoneNumberStatusForPublicResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberStatusForPublicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PhoneNumberStatusForPublicResponse) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForPublicResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForPublicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PhoneNumberStatusForPublicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PhoneNumberStatusForPublicResponse) GetBody() *PhoneNumberStatusForPublicResponseBody {
	return s.Body
}

func (s *PhoneNumberStatusForPublicResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForPublicResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForPublicResponse) SetStatusCode(v int32) *PhoneNumberStatusForPublicResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForPublicResponse) SetBody(v *PhoneNumberStatusForPublicResponseBody) *PhoneNumberStatusForPublicResponse {
	s.Body = v
	return s
}

func (s *PhoneNumberStatusForPublicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
