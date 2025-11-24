// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegisteredServiceEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndPointSlice(v *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) *GetRegisteredServiceEndpointsResponseBody
	GetEndPointSlice() *GetRegisteredServiceEndpointsResponseBodyEndPointSlice
	SetRequestId(v string) *GetRegisteredServiceEndpointsResponseBody
	GetRequestId() *string
	SetServiceEndpoints(v []*GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) *GetRegisteredServiceEndpointsResponseBody
	GetServiceEndpoints() []*GetRegisteredServiceEndpointsResponseBodyServiceEndpoints
}

type GetRegisteredServiceEndpointsResponseBody struct {
	// The name of the registered service.
	EndPointSlice *GetRegisteredServiceEndpointsResponseBodyEndPointSlice `json:"EndPointSlice,omitempty" xml:"EndPointSlice,omitempty" type:"Struct"`
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IP address of the registered service.
	ServiceEndpoints []*GetRegisteredServiceEndpointsResponseBodyServiceEndpoints `json:"ServiceEndpoints,omitempty" xml:"ServiceEndpoints,omitempty" type:"Repeated"`
}

func (s GetRegisteredServiceEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegisteredServiceEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsResponseBody) GetEndPointSlice() *GetRegisteredServiceEndpointsResponseBodyEndPointSlice {
	return s.EndPointSlice
}

func (s *GetRegisteredServiceEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRegisteredServiceEndpointsResponseBody) GetServiceEndpoints() []*GetRegisteredServiceEndpointsResponseBodyServiceEndpoints {
	return s.ServiceEndpoints
}

func (s *GetRegisteredServiceEndpointsResponseBody) SetEndPointSlice(v *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) *GetRegisteredServiceEndpointsResponseBody {
	s.EndPointSlice = v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBody) SetRequestId(v string) *GetRegisteredServiceEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBody) SetServiceEndpoints(v []*GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) *GetRegisteredServiceEndpointsResponseBody {
	s.ServiceEndpoints = v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBody) Validate() error {
	if s.EndPointSlice != nil {
		if err := s.EndPointSlice.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceEndpoints != nil {
		for _, item := range s.ServiceEndpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRegisteredServiceEndpointsResponseBodyEndPointSlice struct {
	// The name of the pod.
	EndpointsDetails []*GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails `json:"EndpointsDetails,omitempty" xml:"EndpointsDetails,omitempty" type:"Repeated"`
	// The details of the endpoint of the registered service.
	//
	// example:
	//
	// MESH_INTERNAL
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The location of the registered service. Valid values:
	//
	// 	- `MESH_INTERNAL`: The service is deployed inside the ASM instance.
	//
	// 	- `MESH_EXTERNAL`: The service is deployed outside the ASM instance.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// reviews
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetRegisteredServiceEndpointsResponseBodyEndPointSlice) String() string {
	return dara.Prettify(s)
}

func (s GetRegisteredServiceEndpointsResponseBodyEndPointSlice) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) GetEndpointsDetails() []*GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	return s.EndpointsDetails
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) GetLocation() *string {
	return s.Location
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) GetNamespace() *string {
	return s.Namespace
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) SetEndpointsDetails(v []*GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) *GetRegisteredServiceEndpointsResponseBodyEndPointSlice {
	s.EndpointsDetails = v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) SetLocation(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSlice {
	s.Location = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) SetNamespace(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSlice {
	s.Namespace = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) SetServiceName(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSlice {
	s.ServiceName = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) Validate() error {
	if s.EndpointsDetails != nil {
		for _, item := range s.EndpointsDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails struct {
	// The port of the registered service.
	//
	// example:
	//
	// ``127.2.**.**``
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the region in which the registered service resides.
	//
	// example:
	//
	// www.demo.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The IP address of the registered service.
	//
	// example:
	//
	// provider-v1-8c86b6898-h***
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// The host name of the registered service.
	Ports []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// Indicates whether sidecar proxies are injected. Valid values:
	//
	// 	- `true`: yes
	//
	// 	- `false`: no
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// false
	SidecarInjected *bool `json:"SidecarInjected,omitempty" xml:"SidecarInjected,omitempty"`
}

func (s GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) String() string {
	return dara.Prettify(s)
}

func (s GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) GetAddress() *string {
	return s.Address
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) GetHostname() *string {
	return s.Hostname
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) GetPodName() *string {
	return s.PodName
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) GetPorts() []*int32 {
	return s.Ports
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) GetRegion() *string {
	return s.Region
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) GetSidecarInjected() *bool {
	return s.SidecarInjected
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetAddress(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.Address = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetHostname(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.Hostname = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetPodName(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.PodName = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetPorts(v []*int32) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.Ports = v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetRegion(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.Region = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetSidecarInjected(v bool) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.SidecarInjected = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) Validate() error {
	return dara.Validate(s)
}

type GetRegisteredServiceEndpointsResponseBodyServiceEndpoints struct {
	// The ID of the cluster on the data plane.
	//
	// example:
	//
	// 192.168.25.153
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The details of the endpoints of the registered service.
	//
	// example:
	//
	// c80f45444b3da447da60a911390c2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) String() string {
	return dara.Prettify(s)
}

func (s GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) GetAddress() *string {
	return s.Address
}

func (s *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) SetAddress(v string) *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints {
	s.Address = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) SetClusterId(v string) *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints {
	s.ClusterId = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) Validate() error {
	return dara.Validate(s)
}
