// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListCustomRoutingEndpointGroupsRequest
	GetAcceleratorId() *string
	SetAccessLogSwitch(v string) *ListCustomRoutingEndpointGroupsRequest
	GetAccessLogSwitch() *string
	SetEndpointGroupId(v string) *ListCustomRoutingEndpointGroupsRequest
	GetEndpointGroupId() *string
	SetListenerId(v string) *ListCustomRoutingEndpointGroupsRequest
	GetListenerId() *string
	SetPageNumber(v int32) *ListCustomRoutingEndpointGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingEndpointGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCustomRoutingEndpointGroupsRequest
	GetRegionId() *string
}

type ListCustomRoutingEndpointGroupsRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Specifies whether the access log feature is enabled.
	//
	// You can set the parameter to **on**.
	//
	// If you leave the parameter empty, all endpoint groups are returned.
	//
	// example:
	//
	// on
	AccessLogSwitch *string `json:"AccessLogSwitch,omitempty" xml:"AccessLogSwitch,omitempty"`
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the custom routing listener.
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
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomRoutingEndpointGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingEndpointGroupsRequest) GetAccessLogSwitch() *string {
	return s.AccessLogSwitch
}

func (s *ListCustomRoutingEndpointGroupsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingEndpointGroupsRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingEndpointGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingEndpointGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingEndpointGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetAcceleratorId(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetAccessLogSwitch(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.AccessLogSwitch = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetEndpointGroupId(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetListenerId(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetPageNumber(v int32) *ListCustomRoutingEndpointGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetPageSize(v int32) *ListCustomRoutingEndpointGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetRegionId(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) Validate() error {
	return dara.Validate(s)
}
