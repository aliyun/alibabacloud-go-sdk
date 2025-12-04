// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErRouteMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *GetErRouteMapRequest
	GetErId() *string
	SetErRouteMapId(v string) *GetErRouteMapRequest
	GetErRouteMapId() *string
	SetRegionId(v string) *GetErRouteMapRequest
	GetRegionId() *string
}

type GetErRouteMapRequest struct {
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// routing policy ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetErRouteMapRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErRouteMapRequest) GoString() string {
	return s.String()
}

func (s *GetErRouteMapRequest) GetErId() *string {
	return s.ErId
}

func (s *GetErRouteMapRequest) GetErRouteMapId() *string {
	return s.ErRouteMapId
}

func (s *GetErRouteMapRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErRouteMapRequest) SetErId(v string) *GetErRouteMapRequest {
	s.ErId = &v
	return s
}

func (s *GetErRouteMapRequest) SetErRouteMapId(v string) *GetErRouteMapRequest {
	s.ErRouteMapId = &v
	return s
}

func (s *GetErRouteMapRequest) SetRegionId(v string) *GetErRouteMapRequest {
	s.RegionId = &v
	return s
}

func (s *GetErRouteMapRequest) Validate() error {
	return dara.Validate(s)
}
