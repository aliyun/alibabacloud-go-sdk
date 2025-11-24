// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIstioGatewayName(v string) *DeleteGatewayRouteRequest
	GetIstioGatewayName() *string
	SetRouteName(v string) *DeleteGatewayRouteRequest
	GetRouteName() *string
	SetServiceMeshId(v string) *DeleteGatewayRouteRequest
	GetServiceMeshId() *string
}

type DeleteGatewayRouteRequest struct {
	// The name of the ASM gateway.
	//
	// example:
	//
	// ingressgateway
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The name of the routing rule.
	//
	// example:
	//
	// http-route
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteGatewayRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *DeleteGatewayRouteRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *DeleteGatewayRouteRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DeleteGatewayRouteRequest) SetIstioGatewayName(v string) *DeleteGatewayRouteRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DeleteGatewayRouteRequest) SetRouteName(v string) *DeleteGatewayRouteRequest {
	s.RouteName = &v
	return s
}

func (s *DeleteGatewayRouteRequest) SetServiceMeshId(v string) *DeleteGatewayRouteRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DeleteGatewayRouteRequest) Validate() error {
	return dara.Validate(s)
}
