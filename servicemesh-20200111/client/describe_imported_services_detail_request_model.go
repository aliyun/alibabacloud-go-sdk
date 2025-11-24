// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportedServicesDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetASMGatewayName(v string) *DescribeImportedServicesDetailRequest
	GetASMGatewayName() *string
	SetServiceMeshId(v string) *DescribeImportedServicesDetailRequest
	GetServiceMeshId() *string
	SetServiceNamespace(v string) *DescribeImportedServicesDetailRequest
	GetServiceNamespace() *string
}

type DescribeImportedServicesDetailRequest struct {
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
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The namespace in which the services reside.
	//
	// example:
	//
	// default
	ServiceNamespace *string `json:"ServiceNamespace,omitempty" xml:"ServiceNamespace,omitempty"`
}

func (s DescribeImportedServicesDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportedServicesDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeImportedServicesDetailRequest) GetASMGatewayName() *string {
	return s.ASMGatewayName
}

func (s *DescribeImportedServicesDetailRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeImportedServicesDetailRequest) GetServiceNamespace() *string {
	return s.ServiceNamespace
}

func (s *DescribeImportedServicesDetailRequest) SetASMGatewayName(v string) *DescribeImportedServicesDetailRequest {
	s.ASMGatewayName = &v
	return s
}

func (s *DescribeImportedServicesDetailRequest) SetServiceMeshId(v string) *DescribeImportedServicesDetailRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeImportedServicesDetailRequest) SetServiceNamespace(v string) *DescribeImportedServicesDetailRequest {
	s.ServiceNamespace = &v
	return s
}

func (s *DescribeImportedServicesDetailRequest) Validate() error {
	return dara.Validate(s)
}
