// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *GetErRouteEntryRequest
	GetErId() *string
	SetErRouteEntryId(v string) *GetErRouteEntryRequest
	GetErRouteEntryId() *string
	SetRegionId(v string) *GetErRouteEntryRequest
	GetRegionId() *string
}

type GetErRouteEntryRequest struct {
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-rte-4q0jbylz
	ErRouteEntryId *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetErRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *GetErRouteEntryRequest) GetErId() *string {
	return s.ErId
}

func (s *GetErRouteEntryRequest) GetErRouteEntryId() *string {
	return s.ErRouteEntryId
}

func (s *GetErRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErRouteEntryRequest) SetErId(v string) *GetErRouteEntryRequest {
	s.ErId = &v
	return s
}

func (s *GetErRouteEntryRequest) SetErRouteEntryId(v string) *GetErRouteEntryRequest {
	s.ErRouteEntryId = &v
	return s
}

func (s *GetErRouteEntryRequest) SetRegionId(v string) *GetErRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *GetErRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
