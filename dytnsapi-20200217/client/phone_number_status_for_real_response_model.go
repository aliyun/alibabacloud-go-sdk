// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForRealResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PhoneNumberStatusForRealResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PhoneNumberStatusForRealResponse
	GetStatusCode() *int32
	SetBody(v *PhoneNumberStatusForRealResponseBody) *PhoneNumberStatusForRealResponse
	GetBody() *PhoneNumberStatusForRealResponseBody
}

type PhoneNumberStatusForRealResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PhoneNumberStatusForRealResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PhoneNumberStatusForRealResponse) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForRealResponse) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PhoneNumberStatusForRealResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PhoneNumberStatusForRealResponse) GetBody() *PhoneNumberStatusForRealResponseBody {
	return s.Body
}

func (s *PhoneNumberStatusForRealResponse) SetHeaders(v map[string]*string) *PhoneNumberStatusForRealResponse {
	s.Headers = v
	return s
}

func (s *PhoneNumberStatusForRealResponse) SetStatusCode(v int32) *PhoneNumberStatusForRealResponse {
	s.StatusCode = &v
	return s
}

func (s *PhoneNumberStatusForRealResponse) SetBody(v *PhoneNumberStatusForRealResponseBody) *PhoneNumberStatusForRealResponse {
	s.Body = v
	return s
}

func (s *PhoneNumberStatusForRealResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
