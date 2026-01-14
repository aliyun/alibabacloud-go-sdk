// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListCustomRoutingEndpointsRequest
	GetAcceleratorId() *string
	SetEndpointGroupId(v string) *ListCustomRoutingEndpointsRequest
	GetEndpointGroupId() *string
	SetListenerId(v string) *ListCustomRoutingEndpointsRequest
	GetListenerId() *string
	SetPageNumber(v int32) *ListCustomRoutingEndpointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingEndpointsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCustomRoutingEndpointsRequest
	GetRegionId() *string
}

type ListCustomRoutingEndpointsRequest struct {
	// The GA instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The endpoint group ID.
	//
	// example:
	//
	// epg-bp16jdc00bhe97sr5****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The listener ID.
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

func (s ListCustomRoutingEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingEndpointsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingEndpointsRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingEndpointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingEndpointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCustomRoutingEndpointsRequest) SetAcceleratorId(v string) *ListCustomRoutingEndpointsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) SetEndpointGroupId(v string) *ListCustomRoutingEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) SetListenerId(v string) *ListCustomRoutingEndpointsRequest {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) SetPageNumber(v int32) *ListCustomRoutingEndpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) SetPageSize(v int32) *ListCustomRoutingEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) SetRegionId(v string) *ListCustomRoutingEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
