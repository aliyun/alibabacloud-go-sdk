// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteTableEntryAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayRouteTableEntryAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayRouteTableEntryAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayRouteTableEntryAttributeResponseBody) *UpdateGatewayRouteTableEntryAttributeResponse
	GetBody() *UpdateGatewayRouteTableEntryAttributeResponseBody
}

type UpdateGatewayRouteTableEntryAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayRouteTableEntryAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayRouteTableEntryAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteTableEntryAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteTableEntryAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayRouteTableEntryAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayRouteTableEntryAttributeResponse) GetBody() *UpdateGatewayRouteTableEntryAttributeResponseBody {
	return s.Body
}

func (s *UpdateGatewayRouteTableEntryAttributeResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteTableEntryAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeResponse) SetStatusCode(v int32) *UpdateGatewayRouteTableEntryAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeResponse) SetBody(v *UpdateGatewayRouteTableEntryAttributeResponseBody) *UpdateGatewayRouteTableEntryAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayRouteTableEntryAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
