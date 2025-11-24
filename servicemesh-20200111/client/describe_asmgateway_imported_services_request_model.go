// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeASMGatewayImportedServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetASMGatewayName(v string) *DescribeASMGatewayImportedServicesRequest
	GetASMGatewayName() *string
	SetServiceMeshId(v string) *DescribeASMGatewayImportedServicesRequest
	GetServiceMeshId() *string
	SetServiceNamespace(v string) *DescribeASMGatewayImportedServicesRequest
	GetServiceNamespace() *string
}

type DescribeASMGatewayImportedServicesRequest struct {
	// The name of the ASM gateway.
	//
	// example:
	//
	// ingressgateway
	ASMGatewayName *string `json:"ASMGatewayName,omitempty" xml:"ASMGatewayName,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The namespace in which the services reside.
	//
	// example:
	//
	// default
	ServiceNamespace *string `json:"ServiceNamespace,omitempty" xml:"ServiceNamespace,omitempty"`
}

func (s DescribeASMGatewayImportedServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeASMGatewayImportedServicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeASMGatewayImportedServicesRequest) GetASMGatewayName() *string {
	return s.ASMGatewayName
}

func (s *DescribeASMGatewayImportedServicesRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeASMGatewayImportedServicesRequest) GetServiceNamespace() *string {
	return s.ServiceNamespace
}

func (s *DescribeASMGatewayImportedServicesRequest) SetASMGatewayName(v string) *DescribeASMGatewayImportedServicesRequest {
	s.ASMGatewayName = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesRequest) SetServiceMeshId(v string) *DescribeASMGatewayImportedServicesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesRequest) SetServiceNamespace(v string) *DescribeASMGatewayImportedServicesRequest {
	s.ServiceNamespace = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesRequest) Validate() error {
	return dara.Validate(s)
}
