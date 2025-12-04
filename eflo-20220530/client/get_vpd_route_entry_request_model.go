// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpdRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetVpdRouteEntryRequest
	GetRegionId() *string
	SetVpdId(v string) *GetVpdRouteEntryRequest
	GetVpdId() *string
	SetVpdRouteEntryId(v string) *GetVpdRouteEntryRequest
	GetVpdRouteEntryId() *string
}

type GetVpdRouteEntryRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-iv2zm1qf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The ID of the route entry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-rte-toekyqel
	VpdRouteEntryId *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
}

func (s GetVpdRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpdRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *GetVpdRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpdRouteEntryRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *GetVpdRouteEntryRequest) GetVpdRouteEntryId() *string {
	return s.VpdRouteEntryId
}

func (s *GetVpdRouteEntryRequest) SetRegionId(v string) *GetVpdRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpdRouteEntryRequest) SetVpdId(v string) *GetVpdRouteEntryRequest {
	s.VpdId = &v
	return s
}

func (s *GetVpdRouteEntryRequest) SetVpdRouteEntryId(v string) *GetVpdRouteEntryRequest {
	s.VpdRouteEntryId = &v
	return s
}

func (s *GetVpdRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
