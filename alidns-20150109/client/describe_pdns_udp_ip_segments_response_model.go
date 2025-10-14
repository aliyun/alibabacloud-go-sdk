// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsUdpIpSegmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsUdpIpSegmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsUdpIpSegmentsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsUdpIpSegmentsResponseBody) *DescribePdnsUdpIpSegmentsResponse
	GetBody() *DescribePdnsUdpIpSegmentsResponseBody
}

type DescribePdnsUdpIpSegmentsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsUdpIpSegmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsUdpIpSegmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsUdpIpSegmentsResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsUdpIpSegmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsUdpIpSegmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsUdpIpSegmentsResponse) GetBody() *DescribePdnsUdpIpSegmentsResponseBody {
	return s.Body
}

func (s *DescribePdnsUdpIpSegmentsResponse) SetHeaders(v map[string]*string) *DescribePdnsUdpIpSegmentsResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponse) SetStatusCode(v int32) *DescribePdnsUdpIpSegmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponse) SetBody(v *DescribePdnsUdpIpSegmentsResponseBody) *DescribePdnsUdpIpSegmentsResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
