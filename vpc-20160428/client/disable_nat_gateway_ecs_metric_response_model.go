// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNatGatewayEcsMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableNatGatewayEcsMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableNatGatewayEcsMetricResponse
	GetStatusCode() *int32
	SetBody(v *DisableNatGatewayEcsMetricResponseBody) *DisableNatGatewayEcsMetricResponse
	GetBody() *DisableNatGatewayEcsMetricResponseBody
}

type DisableNatGatewayEcsMetricResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableNatGatewayEcsMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableNatGatewayEcsMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableNatGatewayEcsMetricResponse) GoString() string {
	return s.String()
}

func (s *DisableNatGatewayEcsMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableNatGatewayEcsMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableNatGatewayEcsMetricResponse) GetBody() *DisableNatGatewayEcsMetricResponseBody {
	return s.Body
}

func (s *DisableNatGatewayEcsMetricResponse) SetHeaders(v map[string]*string) *DisableNatGatewayEcsMetricResponse {
	s.Headers = v
	return s
}

func (s *DisableNatGatewayEcsMetricResponse) SetStatusCode(v int32) *DisableNatGatewayEcsMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableNatGatewayEcsMetricResponse) SetBody(v *DisableNatGatewayEcsMetricResponseBody) *DisableNatGatewayEcsMetricResponse {
	s.Body = v
	return s
}

func (s *DisableNatGatewayEcsMetricResponse) Validate() error {
	return dara.Validate(s)
}
