// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVccRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *DeleteVccRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetRegionId(v string) *DeleteVccRouteEntryRequest
	GetRegionId() *string
	SetVccId(v string) *DeleteVccRouteEntryRequest
	GetVccId() *string
	SetVccRouteEntryId(v string) *DeleteVccRouteEntryRequest
	GetVccRouteEntryId() *string
}

type DeleteVccRouteEntryRequest struct {
	// Destination CIDR block
	//
	// example:
	//
	// 172.16.199.128/25
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// vcc-rte-5cey1sap
	VccRouteEntryId *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
}

func (s DeleteVccRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVccRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteVccRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DeleteVccRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVccRouteEntryRequest) GetVccId() *string {
	return s.VccId
}

func (s *DeleteVccRouteEntryRequest) GetVccRouteEntryId() *string {
	return s.VccRouteEntryId
}

func (s *DeleteVccRouteEntryRequest) SetDestinationCidrBlock(v string) *DeleteVccRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DeleteVccRouteEntryRequest) SetRegionId(v string) *DeleteVccRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVccRouteEntryRequest) SetVccId(v string) *DeleteVccRouteEntryRequest {
	s.VccId = &v
	return s
}

func (s *DeleteVccRouteEntryRequest) SetVccRouteEntryId(v string) *DeleteVccRouteEntryRequest {
	s.VccRouteEntryId = &v
	return s
}

func (s *DeleteVccRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
