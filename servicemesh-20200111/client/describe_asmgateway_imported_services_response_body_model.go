// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeASMGatewayImportedServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportedServices(v []*DescribeASMGatewayImportedServicesResponseBodyImportedServices) *DescribeASMGatewayImportedServicesResponseBody
	GetImportedServices() []*DescribeASMGatewayImportedServicesResponseBodyImportedServices
	SetRequestId(v string) *DescribeASMGatewayImportedServicesResponseBody
	GetRequestId() *string
}

type DescribeASMGatewayImportedServicesResponseBody struct {
	// The list of the imported services.
	ImportedServices []*DescribeASMGatewayImportedServicesResponseBodyImportedServices `json:"ImportedServices,omitempty" xml:"ImportedServices,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 11fd0027-c27e-41bb-a565-75583054****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeASMGatewayImportedServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeASMGatewayImportedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeASMGatewayImportedServicesResponseBody) GetImportedServices() []*DescribeASMGatewayImportedServicesResponseBodyImportedServices {
	return s.ImportedServices
}

func (s *DescribeASMGatewayImportedServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeASMGatewayImportedServicesResponseBody) SetImportedServices(v []*DescribeASMGatewayImportedServicesResponseBodyImportedServices) *DescribeASMGatewayImportedServicesResponseBody {
	s.ImportedServices = v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponseBody) SetRequestId(v string) *DescribeASMGatewayImportedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponseBody) Validate() error {
	if s.ImportedServices != nil {
		for _, item := range s.ImportedServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeASMGatewayImportedServicesResponseBodyImportedServices struct {
	// The name of a service.
	//
	// example:
	//
	// productpage
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The namespace in which the service resides.
	//
	// example:
	//
	// default
	ServiceNamespace *string `json:"ServiceNamespace,omitempty" xml:"ServiceNamespace,omitempty"`
}

func (s DescribeASMGatewayImportedServicesResponseBodyImportedServices) String() string {
	return dara.Prettify(s)
}

func (s DescribeASMGatewayImportedServicesResponseBodyImportedServices) GoString() string {
	return s.String()
}

func (s *DescribeASMGatewayImportedServicesResponseBodyImportedServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeASMGatewayImportedServicesResponseBodyImportedServices) GetServiceNamespace() *string {
	return s.ServiceNamespace
}

func (s *DescribeASMGatewayImportedServicesResponseBodyImportedServices) SetServiceName(v string) *DescribeASMGatewayImportedServicesResponseBodyImportedServices {
	s.ServiceName = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponseBodyImportedServices) SetServiceNamespace(v string) *DescribeASMGatewayImportedServicesResponseBodyImportedServices {
	s.ServiceNamespace = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponseBodyImportedServices) Validate() error {
	return dara.Validate(s)
}
