// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointGroupDestinationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest
	GetAcceleratorId() *string
	SetEndpointGroupId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest
	GetEndpointGroupId() *string
	SetFromPort(v int32) *ListCustomRoutingEndpointGroupDestinationsRequest
	GetFromPort() *int32
	SetListenerId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest
	GetListenerId() *string
	SetPageNumber(v int32) *ListCustomRoutingEndpointGroupDestinationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingEndpointGroupDestinationsRequest
	GetPageSize() *int32
	SetProtocols(v []*string) *ListCustomRoutingEndpointGroupDestinationsRequest
	GetProtocols() []*string
	SetRegionId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest
	GetRegionId() *string
	SetToPort(v int32) *ListCustomRoutingEndpointGroupDestinationsRequest
	GetToPort() *int32
}

type ListCustomRoutingEndpointGroupDestinationsRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp16jdc00bhe97sr5****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The start port of the backend service port range of the endpoint group.
	//
	// Valid values: **1*	- to **65499**. The **FromPort*	- value must be smaller than or equal to the **ToPort*	- value.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The backend service protocols of the endpoint group. Valid values:
	//
	// - **TCP**: TCP.
	//
	// - **UDP**: UDP.
	//
	// - **TCP,UDP**: TCP and UDP.
	//
	// If this parameter is empty, all types of protocols are queried.
	//
	// You can specify up to 10 protocols.
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The end port of the backend service port range of the endpoint group.
	//
	// Valid values: **1*	- to **65499**. The **FromPort*	- value must be smaller than or equal to the **ToPort*	- value.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s ListCustomRoutingEndpointGroupDestinationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupDestinationsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) GetFromPort() *int32 {
	return s.FromPort
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) GetProtocols() []*string {
	return s.Protocols
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) GetToPort() *int32 {
	return s.ToPort
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetAcceleratorId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetEndpointGroupId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetFromPort(v int32) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.FromPort = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetListenerId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetPageNumber(v int32) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetPageSize(v int32) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetProtocols(v []*string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.Protocols = v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetRegionId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetToPort(v int32) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.ToPort = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) Validate() error {
	return dara.Validate(s)
}
