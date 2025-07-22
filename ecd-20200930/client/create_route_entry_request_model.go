// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRouteEntryRequest
	GetDescription() *string
	SetDestinationCidrBlock(v string) *CreateRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetNextHopId(v string) *CreateRouteEntryRequest
	GetNextHopId() *string
	SetNextHopType(v string) *CreateRouteEntryRequest
	GetNextHopType() *string
	SetRegionId(v string) *CreateRouteEntryRequest
	GetRegionId() *string
	SetRouteEntryName(v string) *CreateRouteEntryRequest
	GetRouteEntryName() *string
	SetRouteTableId(v string) *CreateRouteEntryRequest
	GetRouteTableId() *string
}

type CreateRouteEntryRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// This parameter is required.
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// This parameter is required.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s CreateRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteEntryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateRouteEntryRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *CreateRouteEntryRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *CreateRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRouteEntryRequest) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *CreateRouteEntryRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateRouteEntryRequest) SetDescription(v string) *CreateRouteEntryRequest {
	s.Description = &v
	return s
}

func (s *CreateRouteEntryRequest) SetDestinationCidrBlock(v string) *CreateRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateRouteEntryRequest) SetNextHopId(v string) *CreateRouteEntryRequest {
	s.NextHopId = &v
	return s
}

func (s *CreateRouteEntryRequest) SetNextHopType(v string) *CreateRouteEntryRequest {
	s.NextHopType = &v
	return s
}

func (s *CreateRouteEntryRequest) SetRegionId(v string) *CreateRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRouteEntryRequest) SetRouteEntryName(v string) *CreateRouteEntryRequest {
	s.RouteEntryName = &v
	return s
}

func (s *CreateRouteEntryRequest) SetRouteTableId(v string) *CreateRouteEntryRequest {
	s.RouteTableId = &v
	return s
}

func (s *CreateRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
