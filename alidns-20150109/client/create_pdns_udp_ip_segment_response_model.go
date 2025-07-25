// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdnsUdpIpSegmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePdnsUdpIpSegmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePdnsUdpIpSegmentResponse
	GetStatusCode() *int32
	SetBody(v *CreatePdnsUdpIpSegmentResponseBody) *CreatePdnsUdpIpSegmentResponse
	GetBody() *CreatePdnsUdpIpSegmentResponseBody
}

type CreatePdnsUdpIpSegmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePdnsUdpIpSegmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdnsUdpIpSegmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePdnsUdpIpSegmentResponse) GoString() string {
	return s.String()
}

func (s *CreatePdnsUdpIpSegmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePdnsUdpIpSegmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePdnsUdpIpSegmentResponse) GetBody() *CreatePdnsUdpIpSegmentResponseBody {
	return s.Body
}

func (s *CreatePdnsUdpIpSegmentResponse) SetHeaders(v map[string]*string) *CreatePdnsUdpIpSegmentResponse {
	s.Headers = v
	return s
}

func (s *CreatePdnsUdpIpSegmentResponse) SetStatusCode(v int32) *CreatePdnsUdpIpSegmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePdnsUdpIpSegmentResponse) SetBody(v *CreatePdnsUdpIpSegmentResponseBody) *CreatePdnsUdpIpSegmentResponse {
	s.Body = v
	return s
}

func (s *CreatePdnsUdpIpSegmentResponse) Validate() error {
	return dara.Validate(s)
}
