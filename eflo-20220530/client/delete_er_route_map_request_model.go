// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErRouteMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *DeleteErRouteMapRequest
	GetErId() *string
	SetErRouteMapId(v string) *DeleteErRouteMapRequest
	GetErRouteMapId() *string
	SetErRouteMapIds(v []*string) *DeleteErRouteMapRequest
	GetErRouteMapIds() []*string
	SetRegionId(v string) *DeleteErRouteMapRequest
	GetRegionId() *string
}

type DeleteErRouteMapRequest struct {
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId         *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// routing policy Instance ID List
	ErRouteMapIds []*string `json:"ErRouteMapIds,omitempty" xml:"ErRouteMapIds,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteErRouteMapRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteErRouteMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteErRouteMapRequest) GetErId() *string {
	return s.ErId
}

func (s *DeleteErRouteMapRequest) GetErRouteMapId() *string {
	return s.ErRouteMapId
}

func (s *DeleteErRouteMapRequest) GetErRouteMapIds() []*string {
	return s.ErRouteMapIds
}

func (s *DeleteErRouteMapRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteErRouteMapRequest) SetErId(v string) *DeleteErRouteMapRequest {
	s.ErId = &v
	return s
}

func (s *DeleteErRouteMapRequest) SetErRouteMapId(v string) *DeleteErRouteMapRequest {
	s.ErRouteMapId = &v
	return s
}

func (s *DeleteErRouteMapRequest) SetErRouteMapIds(v []*string) *DeleteErRouteMapRequest {
	s.ErRouteMapIds = v
	return s
}

func (s *DeleteErRouteMapRequest) SetRegionId(v string) *DeleteErRouteMapRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteErRouteMapRequest) Validate() error {
	return dara.Validate(s)
}
