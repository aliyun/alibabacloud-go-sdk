// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundPhoneNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOutboundPhoneNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOutboundPhoneNumberResponse
	GetStatusCode() *int32
	SetBody(v *ListOutboundPhoneNumberResponseBody) *ListOutboundPhoneNumberResponse
	GetBody() *ListOutboundPhoneNumberResponseBody
}

type ListOutboundPhoneNumberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOutboundPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOutboundPhoneNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOutboundPhoneNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOutboundPhoneNumberResponse) GetBody() *ListOutboundPhoneNumberResponseBody {
	return s.Body
}

func (s *ListOutboundPhoneNumberResponse) SetHeaders(v map[string]*string) *ListOutboundPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *ListOutboundPhoneNumberResponse) SetStatusCode(v int32) *ListOutboundPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOutboundPhoneNumberResponse) SetBody(v *ListOutboundPhoneNumberResponseBody) *ListOutboundPhoneNumberResponse {
	s.Body = v
	return s
}

func (s *ListOutboundPhoneNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
