// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVccRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *CreateVccRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetRegionId(v string) *CreateVccRouteEntryRequest
	GetRegionId() *string
	SetVccId(v string) *CreateVccRouteEntryRequest
	GetVccId() *string
}

type CreateVccRouteEntryRequest struct {
	// Destination CIDR block
	//
	// example:
	//
	// 192.168.98.112/28
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
}

func (s CreateVccRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVccRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateVccRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateVccRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVccRouteEntryRequest) GetVccId() *string {
	return s.VccId
}

func (s *CreateVccRouteEntryRequest) SetDestinationCidrBlock(v string) *CreateVccRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateVccRouteEntryRequest) SetRegionId(v string) *CreateVccRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVccRouteEntryRequest) SetVccId(v string) *CreateVccRouteEntryRequest {
	s.VccId = &v
	return s
}

func (s *CreateVccRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
