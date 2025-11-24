// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateASMGatewayImportedServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetASMGatewayName(v string) *UpdateASMGatewayImportedServicesRequest
	GetASMGatewayName() *string
	SetServiceMeshId(v string) *UpdateASMGatewayImportedServicesRequest
	GetServiceMeshId() *string
	SetServiceNames(v string) *UpdateASMGatewayImportedServicesRequest
	GetServiceNames() *string
	SetServiceNamespace(v string) *UpdateASMGatewayImportedServicesRequest
	GetServiceNamespace() *string
}

type UpdateASMGatewayImportedServicesRequest struct {
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
	// The names of the services. Separate multiple service names with commas (,). Example: reviews,sleep.
	//
	// example:
	//
	// reviews,sleep
	ServiceNames *string `json:"ServiceNames,omitempty" xml:"ServiceNames,omitempty"`
	// The namespace in which the service resides.
	//
	// example:
	//
	// default
	ServiceNamespace *string `json:"ServiceNamespace,omitempty" xml:"ServiceNamespace,omitempty"`
}

func (s UpdateASMGatewayImportedServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateASMGatewayImportedServicesRequest) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayImportedServicesRequest) GetASMGatewayName() *string {
	return s.ASMGatewayName
}

func (s *UpdateASMGatewayImportedServicesRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateASMGatewayImportedServicesRequest) GetServiceNames() *string {
	return s.ServiceNames
}

func (s *UpdateASMGatewayImportedServicesRequest) GetServiceNamespace() *string {
	return s.ServiceNamespace
}

func (s *UpdateASMGatewayImportedServicesRequest) SetASMGatewayName(v string) *UpdateASMGatewayImportedServicesRequest {
	s.ASMGatewayName = &v
	return s
}

func (s *UpdateASMGatewayImportedServicesRequest) SetServiceMeshId(v string) *UpdateASMGatewayImportedServicesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateASMGatewayImportedServicesRequest) SetServiceNames(v string) *UpdateASMGatewayImportedServicesRequest {
	s.ServiceNames = &v
	return s
}

func (s *UpdateASMGatewayImportedServicesRequest) SetServiceNamespace(v string) *UpdateASMGatewayImportedServicesRequest {
	s.ServiceNamespace = &v
	return s
}

func (s *UpdateASMGatewayImportedServicesRequest) Validate() error {
	return dara.Validate(s)
}
