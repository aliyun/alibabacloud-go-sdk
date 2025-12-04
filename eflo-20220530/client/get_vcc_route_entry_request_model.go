// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVccRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetVccRouteEntryRequest
	GetRegionId() *string
	SetVccId(v string) *GetVccRouteEntryRequest
	GetVccId() *string
	SetVccRouteEntryId(v string) *GetVccRouteEntryRequest
	GetVccRouteEntryId() *string
}

type GetVccRouteEntryRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun Connection ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The ID of the route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-rte-31ocvdhq
	VccRouteEntryId *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
}

func (s GetVccRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVccRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *GetVccRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVccRouteEntryRequest) GetVccId() *string {
	return s.VccId
}

func (s *GetVccRouteEntryRequest) GetVccRouteEntryId() *string {
	return s.VccRouteEntryId
}

func (s *GetVccRouteEntryRequest) SetRegionId(v string) *GetVccRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *GetVccRouteEntryRequest) SetVccId(v string) *GetVccRouteEntryRequest {
	s.VccId = &v
	return s
}

func (s *GetVccRouteEntryRequest) SetVccRouteEntryId(v string) *GetVccRouteEntryRequest {
	s.VccRouteEntryId = &v
	return s
}

func (s *GetVccRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
