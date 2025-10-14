// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidatePdnsUdpIpSegmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidatePdnsUdpIpSegmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidatePdnsUdpIpSegmentResponse
	GetStatusCode() *int32
	SetBody(v *ValidatePdnsUdpIpSegmentResponseBody) *ValidatePdnsUdpIpSegmentResponse
	GetBody() *ValidatePdnsUdpIpSegmentResponseBody
}

type ValidatePdnsUdpIpSegmentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidatePdnsUdpIpSegmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidatePdnsUdpIpSegmentResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidatePdnsUdpIpSegmentResponse) GoString() string {
	return s.String()
}

func (s *ValidatePdnsUdpIpSegmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidatePdnsUdpIpSegmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidatePdnsUdpIpSegmentResponse) GetBody() *ValidatePdnsUdpIpSegmentResponseBody {
	return s.Body
}

func (s *ValidatePdnsUdpIpSegmentResponse) SetHeaders(v map[string]*string) *ValidatePdnsUdpIpSegmentResponse {
	s.Headers = v
	return s
}

func (s *ValidatePdnsUdpIpSegmentResponse) SetStatusCode(v int32) *ValidatePdnsUdpIpSegmentResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidatePdnsUdpIpSegmentResponse) SetBody(v *ValidatePdnsUdpIpSegmentResponseBody) *ValidatePdnsUdpIpSegmentResponse {
	s.Body = v
	return s
}

func (s *ValidatePdnsUdpIpSegmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
