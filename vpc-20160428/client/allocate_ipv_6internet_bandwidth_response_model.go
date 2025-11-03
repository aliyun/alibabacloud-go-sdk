// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateIpv6InternetBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateIpv6InternetBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateIpv6InternetBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *AllocateIpv6InternetBandwidthResponseBody) *AllocateIpv6InternetBandwidthResponse
	GetBody() *AllocateIpv6InternetBandwidthResponseBody
}

type AllocateIpv6InternetBandwidthResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateIpv6InternetBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateIpv6InternetBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateIpv6InternetBandwidthResponse) GoString() string {
	return s.String()
}

func (s *AllocateIpv6InternetBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateIpv6InternetBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateIpv6InternetBandwidthResponse) GetBody() *AllocateIpv6InternetBandwidthResponseBody {
	return s.Body
}

func (s *AllocateIpv6InternetBandwidthResponse) SetHeaders(v map[string]*string) *AllocateIpv6InternetBandwidthResponse {
	s.Headers = v
	return s
}

func (s *AllocateIpv6InternetBandwidthResponse) SetStatusCode(v int32) *AllocateIpv6InternetBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthResponse) SetBody(v *AllocateIpv6InternetBandwidthResponseBody) *AllocateIpv6InternetBandwidthResponse {
	s.Body = v
	return s
}

func (s *AllocateIpv6InternetBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
