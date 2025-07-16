// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIntranetSupportedZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayIntranetSupportedZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayIntranetSupportedZoneResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayIntranetSupportedZoneResponseBody) *ListGatewayIntranetSupportedZoneResponse
	GetBody() *ListGatewayIntranetSupportedZoneResponseBody
}

type ListGatewayIntranetSupportedZoneResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayIntranetSupportedZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayIntranetSupportedZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetSupportedZoneResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetSupportedZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayIntranetSupportedZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayIntranetSupportedZoneResponse) GetBody() *ListGatewayIntranetSupportedZoneResponseBody {
	return s.Body
}

func (s *ListGatewayIntranetSupportedZoneResponse) SetHeaders(v map[string]*string) *ListGatewayIntranetSupportedZoneResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayIntranetSupportedZoneResponse) SetStatusCode(v int32) *ListGatewayIntranetSupportedZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayIntranetSupportedZoneResponse) SetBody(v *ListGatewayIntranetSupportedZoneResponseBody) *ListGatewayIntranetSupportedZoneResponse {
	s.Body = v
	return s
}

func (s *ListGatewayIntranetSupportedZoneResponse) Validate() error {
	return dara.Validate(s)
}
