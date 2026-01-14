// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingPortMappingsByDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationAddress(v string) *ListCustomRoutingPortMappingsByDestinationRequest
	GetDestinationAddress() *string
	SetEndpointId(v string) *ListCustomRoutingPortMappingsByDestinationRequest
	GetEndpointId() *string
	SetPageNumber(v int32) *ListCustomRoutingPortMappingsByDestinationRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingPortMappingsByDestinationRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCustomRoutingPortMappingsByDestinationRequest
	GetRegionId() *string
}

type ListCustomRoutingPortMappingsByDestinationRequest struct {
	// The service IP address of the backend instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX
	DestinationAddress *string `json:"DestinationAddress,omitempty" xml:"DestinationAddress,omitempty"`
	// The ID of the endpoint to which the backend instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-bp16jdc00bhe97sr5****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
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
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomRoutingPortMappingsByDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingPortMappingsByDestinationRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) GetDestinationAddress() *string {
	return s.DestinationAddress
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) SetDestinationAddress(v string) *ListCustomRoutingPortMappingsByDestinationRequest {
	s.DestinationAddress = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) SetEndpointId(v string) *ListCustomRoutingPortMappingsByDestinationRequest {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) SetPageNumber(v int32) *ListCustomRoutingPortMappingsByDestinationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) SetPageSize(v int32) *ListCustomRoutingPortMappingsByDestinationRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) SetRegionId(v string) *ListCustomRoutingPortMappingsByDestinationRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) Validate() error {
	return dara.Validate(s)
}
