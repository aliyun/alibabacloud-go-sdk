// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnhanhcedNatGatewayAvailableZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnhanhcedNatGatewayAvailableZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnhanhcedNatGatewayAvailableZonesResponse
	GetStatusCode() *int32
	SetBody(v *ListEnhanhcedNatGatewayAvailableZonesResponseBody) *ListEnhanhcedNatGatewayAvailableZonesResponse
	GetBody() *ListEnhanhcedNatGatewayAvailableZonesResponseBody
}

type ListEnhanhcedNatGatewayAvailableZonesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnhanhcedNatGatewayAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnhanhcedNatGatewayAvailableZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnhanhcedNatGatewayAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponse) GetBody() *ListEnhanhcedNatGatewayAvailableZonesResponseBody {
	return s.Body
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponse) SetHeaders(v map[string]*string) *ListEnhanhcedNatGatewayAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponse) SetStatusCode(v int32) *ListEnhanhcedNatGatewayAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponse) SetBody(v *ListEnhanhcedNatGatewayAvailableZonesResponseBody) *ListEnhanhcedNatGatewayAvailableZonesResponse {
	s.Body = v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
