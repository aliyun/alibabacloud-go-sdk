// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingPortMappingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListCustomRoutingPortMappingsRequest
	GetAcceleratorId() *string
	SetEndpointGroupId(v string) *ListCustomRoutingPortMappingsRequest
	GetEndpointGroupId() *string
	SetListenerId(v string) *ListCustomRoutingPortMappingsRequest
	GetListenerId() *string
	SetPageNumber(v int32) *ListCustomRoutingPortMappingsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingPortMappingsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCustomRoutingPortMappingsRequest
	GetRegionId() *string
}

type ListCustomRoutingPortMappingsRequest struct {
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
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomRoutingPortMappingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingPortMappingsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingPortMappingsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingPortMappingsRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingPortMappingsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingPortMappingsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingPortMappingsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCustomRoutingPortMappingsRequest) SetAcceleratorId(v string) *ListCustomRoutingPortMappingsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) SetEndpointGroupId(v string) *ListCustomRoutingPortMappingsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) SetListenerId(v string) *ListCustomRoutingPortMappingsRequest {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) SetPageNumber(v int32) *ListCustomRoutingPortMappingsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) SetPageSize(v int32) *ListCustomRoutingPortMappingsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) SetRegionId(v string) *ListCustomRoutingPortMappingsRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) Validate() error {
	return dara.Validate(s)
}
