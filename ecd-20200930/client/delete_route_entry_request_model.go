// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *DeleteRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetNextHopId(v string) *DeleteRouteEntryRequest
	GetNextHopId() *string
	SetRegionId(v string) *DeleteRouteEntryRequest
	GetRegionId() *string
	SetRouteEntryId(v string) *DeleteRouteEntryRequest
	GetRouteEntryId() *string
	SetRouteTableId(v string) *DeleteRouteEntryRequest
	GetRouteTableId() *string
}

type DeleteRouteEntryRequest struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// This parameter is required.
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// This parameter is required.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DeleteRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DeleteRouteEntryRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DeleteRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouteEntryRequest) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DeleteRouteEntryRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DeleteRouteEntryRequest) SetDestinationCidrBlock(v string) *DeleteRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetNextHopId(v string) *DeleteRouteEntryRequest {
	s.NextHopId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetRegionId(v string) *DeleteRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetRouteEntryId(v string) *DeleteRouteEntryRequest {
	s.RouteEntryId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetRouteTableId(v string) *DeleteRouteEntryRequest {
	s.RouteTableId = &v
	return s
}

func (s *DeleteRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
