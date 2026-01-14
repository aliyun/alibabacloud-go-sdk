// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingPortMappingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCustomRoutingPortMappingsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingPortMappingsResponseBody
	GetPageSize() *int32
	SetPortMappings(v []*ListCustomRoutingPortMappingsResponseBodyPortMappings) *ListCustomRoutingPortMappingsResponseBody
	GetPortMappings() []*ListCustomRoutingPortMappingsResponseBodyPortMappings
	SetRequestId(v string) *ListCustomRoutingPortMappingsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomRoutingPortMappingsResponseBody
	GetTotalCount() *int32
}

type ListCustomRoutingPortMappingsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Details about the port mapping table.
	PortMappings []*ListCustomRoutingPortMappingsResponseBodyPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingPortMappingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingPortMappingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingPortMappingsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingPortMappingsResponseBody) GetPortMappings() []*ListCustomRoutingPortMappingsResponseBodyPortMappings {
	return s.PortMappings
}

func (s *ListCustomRoutingPortMappingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomRoutingPortMappingsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomRoutingPortMappingsResponseBody) SetPageNumber(v int32) *ListCustomRoutingPortMappingsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBody) SetPageSize(v int32) *ListCustomRoutingPortMappingsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBody) SetPortMappings(v []*ListCustomRoutingPortMappingsResponseBodyPortMappings) *ListCustomRoutingPortMappingsResponseBody {
	s.PortMappings = v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBody) SetRequestId(v string) *ListCustomRoutingPortMappingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBody) SetTotalCount(v int32) *ListCustomRoutingPortMappingsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBody) Validate() error {
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

type ListCustomRoutingPortMappingsResponseBodyPortMappings struct {
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
	DestinationSocketAddress *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress `json:"DestinationSocketAddress,omitempty" xml:"DestinationSocketAddress,omitempty" type:"Struct"`
	// The access policy of traffic for the backend instance. Valid values:
	//
	// 	- **allow**: allows traffic to the backend instance.
	//
	// 	- **deny**: denies traffic to the backend instance.
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
	// lsr-bp1bpn0kn908w4nbw****
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

func (s ListCustomRoutingPortMappingsResponseBodyPortMappings) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingPortMappingsResponseBodyPortMappings) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) GetAcceleratorPort() *int32 {
	return s.AcceleratorPort
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) GetDestinationSocketAddress() *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress {
	return s.DestinationSocketAddress
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) GetDestinationTrafficState() *string {
	return s.DestinationTrafficState
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) GetProtocols() []*string {
	return s.Protocols
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) GetVswitch() *string {
	return s.Vswitch
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetAcceleratorId(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetAcceleratorPort(v int32) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.AcceleratorPort = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetDestinationSocketAddress(v *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.DestinationSocketAddress = v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetDestinationTrafficState(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.DestinationTrafficState = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetEndpointGroupId(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetEndpointGroupRegion(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetEndpointId(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetListenerId(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetProtocols(v []*string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.Protocols = v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetVswitch(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.Vswitch = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) Validate() error {
	if s.DestinationSocketAddress != nil {
		if err := s.DestinationSocketAddress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress struct {
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

func (s ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) GetIpAddress() *string {
	return s.IpAddress
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) GetPort() *int32 {
	return s.Port
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) SetIpAddress(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress {
	s.IpAddress = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) SetPort(v int32) *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress {
	s.Port = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) Validate() error {
	return dara.Validate(s)
}
