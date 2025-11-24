// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewaySecretDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIstioGatewayName(v string) *DescribeGatewaySecretDetailsRequest
	GetIstioGatewayName() *string
	SetServiceMeshId(v string) *DescribeGatewaySecretDetailsRequest
	GetServiceMeshId() *string
}

type DescribeGatewaySecretDetailsRequest struct {
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

func (s DescribeGatewaySecretDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewaySecretDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySecretDetailsRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *DescribeGatewaySecretDetailsRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeGatewaySecretDetailsRequest) SetIstioGatewayName(v string) *DescribeGatewaySecretDetailsRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DescribeGatewaySecretDetailsRequest) SetServiceMeshId(v string) *DescribeGatewaySecretDetailsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeGatewaySecretDetailsRequest) Validate() error {
	return dara.Validate(s)
}
