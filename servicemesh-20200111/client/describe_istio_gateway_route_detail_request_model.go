// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIstioGatewayRouteDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIstioGatewayName(v string) *DescribeIstioGatewayRouteDetailRequest
	GetIstioGatewayName() *string
	SetRouteName(v string) *DescribeIstioGatewayRouteDetailRequest
	GetRouteName() *string
	SetServiceMeshId(v string) *DescribeIstioGatewayRouteDetailRequest
	GetServiceMeshId() *string
}

type DescribeIstioGatewayRouteDetailRequest struct {
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
	// demo-route
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

func (s DescribeIstioGatewayRouteDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *DescribeIstioGatewayRouteDetailRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *DescribeIstioGatewayRouteDetailRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeIstioGatewayRouteDetailRequest) SetIstioGatewayName(v string) *DescribeIstioGatewayRouteDetailRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailRequest) SetRouteName(v string) *DescribeIstioGatewayRouteDetailRequest {
	s.RouteName = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailRequest) SetServiceMeshId(v string) *DescribeIstioGatewayRouteDetailRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailRequest) Validate() error {
	return dara.Validate(s)
}
