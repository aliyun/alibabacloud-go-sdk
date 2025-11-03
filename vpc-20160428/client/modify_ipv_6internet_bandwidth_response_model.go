// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpv6InternetBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIpv6InternetBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIpv6InternetBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIpv6InternetBandwidthResponseBody) *ModifyIpv6InternetBandwidthResponse
	GetBody() *ModifyIpv6InternetBandwidthResponseBody
}

type ModifyIpv6InternetBandwidthResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpv6InternetBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIpv6InternetBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpv6InternetBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpv6InternetBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIpv6InternetBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIpv6InternetBandwidthResponse) GetBody() *ModifyIpv6InternetBandwidthResponseBody {
	return s.Body
}

func (s *ModifyIpv6InternetBandwidthResponse) SetHeaders(v map[string]*string) *ModifyIpv6InternetBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpv6InternetBandwidthResponse) SetStatusCode(v int32) *ModifyIpv6InternetBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthResponse) SetBody(v *ModifyIpv6InternetBandwidthResponseBody) *ModifyIpv6InternetBandwidthResponse {
	s.Body = v
	return s
}

func (s *ModifyIpv6InternetBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
