// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointTrafficPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest
	GetAcceleratorId() *string
	SetAddress(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest
	GetAddress() *string
	SetEndpointGroupId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest
	GetEndpointGroupId() *string
	SetEndpointId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest
	GetEndpointId() *string
	SetListenerId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest
	GetListenerId() *string
	SetPageNumber(v int32) *ListCustomRoutingEndpointTrafficPoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingEndpointTrafficPoliciesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest
	GetRegionId() *string
}

type ListCustomRoutingEndpointTrafficPoliciesRequest struct {
	// The ID of the GA instance to which the traffic policies belong.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The IP address of the traffic destination.
	//
	// example:
	//
	// 10.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the endpoint group to which the traffic policies belong.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the endpoint to which the traffic policies belong.
	//
	// example:
	//
	// ep-bp16jdc00bhe97sr5****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the listener to which the traffic policies belong.
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
	// The number of entries per page. Valid values: **1 to 50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomRoutingEndpointTrafficPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) GetAddress() *string {
	return s.Address
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetAcceleratorId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetAddress(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.Address = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetEndpointGroupId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetEndpointId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetListenerId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetPageNumber(v int32) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetPageSize(v int32) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetRegionId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
