// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv6InternetBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpv6InternetBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpv6InternetBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpv6InternetBandwidthResponseBody) *DeleteIpv6InternetBandwidthResponse
	GetBody() *DeleteIpv6InternetBandwidthResponseBody
}

type DeleteIpv6InternetBandwidthResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpv6InternetBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpv6InternetBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv6InternetBandwidthResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpv6InternetBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpv6InternetBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpv6InternetBandwidthResponse) GetBody() *DeleteIpv6InternetBandwidthResponseBody {
	return s.Body
}

func (s *DeleteIpv6InternetBandwidthResponse) SetHeaders(v map[string]*string) *DeleteIpv6InternetBandwidthResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpv6InternetBandwidthResponse) SetStatusCode(v int32) *DeleteIpv6InternetBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthResponse) SetBody(v *DeleteIpv6InternetBandwidthResponseBody) *DeleteIpv6InternetBandwidthResponse {
	s.Body = v
	return s
}

func (s *DeleteIpv6InternetBandwidthResponse) Validate() error {
	return dara.Validate(s)
}
