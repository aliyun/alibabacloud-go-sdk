// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PhoneNumberStatusForVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PhoneNumberStatusForVoiceResponse
	GetStatusCode() *int32
	SetBody(v *PhoneNumberStatusForVoiceResponseBody) *PhoneNumberStatusForVoiceResponse
	GetBody() *PhoneNumberStatusForVoiceResponseBody
}

type PhoneNumberStatusForVoiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberStatusForVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PhoneNumberStatusForVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForVoiceResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PhoneNumberStatusForVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PhoneNumberStatusForVoiceResponse) GetBody() *PhoneNumberStatusForVoiceResponseBody {
	return s.Body
}

func (s *PhoneNumberStatusForVoiceResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForVoiceResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForVoiceResponse) SetStatusCode(v int32) *PhoneNumberStatusForVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponse) SetBody(v *PhoneNumberStatusForVoiceResponseBody) *PhoneNumberStatusForVoiceResponse {
	s.Body = v
	return s
}

func (s *PhoneNumberStatusForVoiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
