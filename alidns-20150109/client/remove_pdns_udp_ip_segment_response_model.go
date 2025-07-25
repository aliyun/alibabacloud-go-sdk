// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePdnsUdpIpSegmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePdnsUdpIpSegmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePdnsUdpIpSegmentResponse
	GetStatusCode() *int32
	SetBody(v *RemovePdnsUdpIpSegmentResponseBody) *RemovePdnsUdpIpSegmentResponse
	GetBody() *RemovePdnsUdpIpSegmentResponseBody
}

type RemovePdnsUdpIpSegmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePdnsUdpIpSegmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePdnsUdpIpSegmentResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePdnsUdpIpSegmentResponse) GoString() string {
	return s.String()
}

func (s *RemovePdnsUdpIpSegmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePdnsUdpIpSegmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePdnsUdpIpSegmentResponse) GetBody() *RemovePdnsUdpIpSegmentResponseBody {
	return s.Body
}

func (s *RemovePdnsUdpIpSegmentResponse) SetHeaders(v map[string]*string) *RemovePdnsUdpIpSegmentResponse {
	s.Headers = v
	return s
}

func (s *RemovePdnsUdpIpSegmentResponse) SetStatusCode(v int32) *RemovePdnsUdpIpSegmentResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePdnsUdpIpSegmentResponse) SetBody(v *RemovePdnsUdpIpSegmentResponseBody) *RemovePdnsUdpIpSegmentResponse {
	s.Body = v
	return s
}

func (s *RemovePdnsUdpIpSegmentResponse) Validate() error {
	return dara.Validate(s)
}
