// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportedServicesDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetails(v []*DescribeImportedServicesDetailResponseBodyDetails) *DescribeImportedServicesDetailResponseBody
	GetDetails() []*DescribeImportedServicesDetailResponseBodyDetails
	SetRequestId(v string) *DescribeImportedServicesDetailResponseBody
	GetRequestId() *string
}

type DescribeImportedServicesDetailResponseBody struct {
	// The details of the services.
	Details []*DescribeImportedServicesDetailResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E0496204-7586-5B4C-B364-2361CC0EDxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImportedServicesDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportedServicesDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImportedServicesDetailResponseBody) GetDetails() []*DescribeImportedServicesDetailResponseBodyDetails {
	return s.Details
}

func (s *DescribeImportedServicesDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImportedServicesDetailResponseBody) SetDetails(v []*DescribeImportedServicesDetailResponseBodyDetails) *DescribeImportedServicesDetailResponseBody {
	s.Details = v
	return s
}

func (s *DescribeImportedServicesDetailResponseBody) SetRequestId(v string) *DescribeImportedServicesDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBody) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImportedServicesDetailResponseBodyDetails struct {
	// The clusters on the data plane.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
	// The labels of the service.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The namespace in which the service resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ports declared for the service.
	Ports []*DescribeImportedServicesDetailResponseBodyDetailsPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The name of a service.
	//
	// example:
	//
	// productpage
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// Kubernetes
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s DescribeImportedServicesDetailResponseBodyDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportedServicesDetailResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) GetClusterIds() []*string {
	return s.ClusterIds
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) GetLabels() map[string]*string {
	return s.Labels
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) GetPorts() []*DescribeImportedServicesDetailResponseBodyDetailsPorts {
	return s.Ports
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetClusterIds(v []*string) *DescribeImportedServicesDetailResponseBodyDetails {
	s.ClusterIds = v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetLabels(v map[string]*string) *DescribeImportedServicesDetailResponseBodyDetails {
	s.Labels = v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetNamespace(v string) *DescribeImportedServicesDetailResponseBodyDetails {
	s.Namespace = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetPorts(v []*DescribeImportedServicesDetailResponseBodyDetailsPorts) *DescribeImportedServicesDetailResponseBodyDetails {
	s.Ports = v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetServiceName(v string) *DescribeImportedServicesDetailResponseBodyDetails {
	s.ServiceName = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetServiceType(v string) *DescribeImportedServicesDetailResponseBodyDetails {
	s.ServiceType = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) Validate() error {
	if s.Ports != nil {
		for _, item := range s.Ports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImportedServicesDetailResponseBodyDetailsPorts struct {
	// The name of a port.
	//
	// example:
	//
	// http-0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node port.
	//
	// example:
	//
	// 12345
	NodePort *int32 `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	// The port number.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol of the port.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The container port.
	//
	// example:
	//
	// 8080
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s DescribeImportedServicesDetailResponseBodyDetailsPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportedServicesDetailResponseBodyDetailsPorts) GoString() string {
	return s.String()
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) GetName() *string {
	return s.Name
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) GetNodePort() *int32 {
	return s.NodePort
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) GetPort() *int32 {
	return s.Port
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) SetName(v string) *DescribeImportedServicesDetailResponseBodyDetailsPorts {
	s.Name = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) SetNodePort(v int32) *DescribeImportedServicesDetailResponseBodyDetailsPorts {
	s.NodePort = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) SetPort(v int32) *DescribeImportedServicesDetailResponseBodyDetailsPorts {
	s.Port = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) SetProtocol(v string) *DescribeImportedServicesDetailResponseBodyDetailsPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) SetTargetPort(v int32) *DescribeImportedServicesDetailResponseBodyDetailsPorts {
	s.TargetPort = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) Validate() error {
	return dara.Validate(s)
}
