// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingPortMappingsByDestinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBody
	GetPageSize() *int32
	SetPortMappings(v []*ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) *ListCustomRoutingPortMappingsByDestinationResponseBody
	GetPortMappings() []*ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings
	SetRequestId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBody
	GetTotalCount() *int32
}

type ListCustomRoutingPortMappingsByDestinationResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Details about the port mapping table.
	PortMappings []*ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// String	04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) GetPortMappings() []*ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	return s.PortMappings
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) SetPageNumber(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) SetPageSize(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) SetPortMappings(v []*ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) *ListCustomRoutingPortMappingsByDestinationResponseBody {
	s.PortMappings = v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) SetRequestId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) SetTotalCount(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) Validate() error {
	if s.PortMappings != nil {
		for _, item := range s.PortMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The acceleration port.
	//
	// example:
	//
	// 3000
	AcceleratorPort *int32 `json:"AcceleratorPort,omitempty" xml:"AcceleratorPort,omitempty"`
	// The service IP address and port of the backend instance.
	DestinationSocketAddress *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress `json:"DestinationSocketAddress,omitempty" xml:"DestinationSocketAddress,omitempty" type:"Struct"`
	// The access policy of traffic for the backend instance.
	//
	// 	- **allow**: allows traffic to the backend instance.
	//
	// 	- **deny**: denies all traffic to the backend instance.
	//
	// example:
	//
	// allow
	DestinationTrafficState *string `json:"DestinationTrafficState,omitempty" xml:"DestinationTrafficState,omitempty"`
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the region in which the endpoint group resides.
	//
	// example:
	//
	// us-west-1
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the listener.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The protocol of the backend service.
	//
	// 	- **tcp**: TCP
	//
	// 	- **udp**: UDP
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The name of the endpoint (vSwitch).
	//
	// example:
	//
	// vsw-test01
	Vswitch *string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GetAcceleratorPort() *int32 {
	return s.AcceleratorPort
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GetDestinationSocketAddress() *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress {
	return s.DestinationSocketAddress
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GetDestinationTrafficState() *string {
	return s.DestinationTrafficState
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GetProtocols() []*string {
	return s.Protocols
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GetVswitch() *string {
	return s.Vswitch
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetAcceleratorId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetAcceleratorPort(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.AcceleratorPort = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetDestinationSocketAddress(v *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.DestinationSocketAddress = v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetDestinationTrafficState(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.DestinationTrafficState = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetEndpointGroupId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetEndpointGroupRegion(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetEndpointId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetListenerId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetProtocols(v []*string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.Protocols = v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetVswitch(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.Vswitch = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) Validate() error {
	if s.DestinationSocketAddress != nil {
		if err := s.DestinationSocketAddress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress struct {
	// The service IP address of the backend instance.
	//
	// example:
	//
	// 10.0.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The service port of the backend instance.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) GetIpAddress() *string {
	return s.IpAddress
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) GetPort() *int32 {
	return s.Port
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) SetIpAddress(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress {
	s.IpAddress = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) SetPort(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress {
	s.Port = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) Validate() error {
	return dara.Validate(s)
}
