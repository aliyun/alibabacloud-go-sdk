// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteRoutineRouteRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteRoutineRouteRequest
	GetSiteId() *int64
}

type DeleteRoutineRouteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11223***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteRoutineRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineRouteRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteRoutineRouteRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteRoutineRouteRequest) SetConfigId(v int64) *DeleteRoutineRouteRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteRoutineRouteRequest) SetSiteId(v int64) *DeleteRoutineRouteRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteRoutineRouteRequest) Validate() error {
	return dara.Validate(s)
}
