// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCalculateNatGatewayCapacityUnitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CalculateNatGatewayCapacityUnitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CalculateNatGatewayCapacityUnitsResponse
	GetStatusCode() *int32
	SetBody(v *CalculateNatGatewayCapacityUnitsResponseBody) *CalculateNatGatewayCapacityUnitsResponse
	GetBody() *CalculateNatGatewayCapacityUnitsResponseBody
}

type CalculateNatGatewayCapacityUnitsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CalculateNatGatewayCapacityUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CalculateNatGatewayCapacityUnitsResponse) String() string {
	return dara.Prettify(s)
}

func (s CalculateNatGatewayCapacityUnitsResponse) GoString() string {
	return s.String()
}

func (s *CalculateNatGatewayCapacityUnitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CalculateNatGatewayCapacityUnitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CalculateNatGatewayCapacityUnitsResponse) GetBody() *CalculateNatGatewayCapacityUnitsResponseBody {
	return s.Body
}

func (s *CalculateNatGatewayCapacityUnitsResponse) SetHeaders(v map[string]*string) *CalculateNatGatewayCapacityUnitsResponse {
	s.Headers = v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponse) SetStatusCode(v int32) *CalculateNatGatewayCapacityUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponse) SetBody(v *CalculateNatGatewayCapacityUnitsResponseBody) *CalculateNatGatewayCapacityUnitsResponse {
	s.Body = v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
