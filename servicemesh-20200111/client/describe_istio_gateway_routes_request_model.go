// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIstioGatewayRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIstioGatewayName(v string) *DescribeIstioGatewayRoutesRequest
	GetIstioGatewayName() *string
	SetServiceMeshId(v string) *DescribeIstioGatewayRoutesRequest
	GetServiceMeshId() *string
}

type DescribeIstioGatewayRoutesRequest struct {
	// The name of the ASM gateway.
	//
	// example:
	//
	// ingressgateway
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeIstioGatewayRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRoutesRequest) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRoutesRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *DescribeIstioGatewayRoutesRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeIstioGatewayRoutesRequest) SetIstioGatewayName(v string) *DescribeIstioGatewayRoutesRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DescribeIstioGatewayRoutesRequest) SetServiceMeshId(v string) *DescribeIstioGatewayRoutesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeIstioGatewayRoutesRequest) Validate() error {
	return dara.Validate(s)
}
