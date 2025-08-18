// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoutineRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBypass(v string) *UpdateRoutineRouteRequest
	GetBypass() *string
	SetConfigId(v int64) *UpdateRoutineRouteRequest
	GetConfigId() *int64
	SetFallback(v string) *UpdateRoutineRouteRequest
	GetFallback() *string
	SetRouteEnable(v string) *UpdateRoutineRouteRequest
	GetRouteEnable() *string
	SetRouteName(v string) *UpdateRoutineRouteRequest
	GetRouteName() *string
	SetRoutineName(v string) *UpdateRoutineRouteRequest
	GetRoutineName() *string
	SetRule(v string) *UpdateRoutineRouteRequest
	GetRule() *string
	SetSequence(v int32) *UpdateRoutineRouteRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateRoutineRouteRequest
	GetSiteId() *int64
}

type UpdateRoutineRouteRequest struct {
	// example:
	//
	// on
	Bypass *string `json:"Bypass,omitempty" xml:"Bypass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Fallback *string `json:"Fallback,omitempty" xml:"Fallback,omitempty"`
	// example:
	//
	// on
	RouteEnable *string `json:"RouteEnable,omitempty" xml:"RouteEnable,omitempty"`
	// example:
	//
	// test_route
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// example:
	//
	// test-routine1
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456******
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateRoutineRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoutineRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoutineRouteRequest) GetBypass() *string {
	return s.Bypass
}

func (s *UpdateRoutineRouteRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateRoutineRouteRequest) GetFallback() *string {
	return s.Fallback
}

func (s *UpdateRoutineRouteRequest) GetRouteEnable() *string {
	return s.RouteEnable
}

func (s *UpdateRoutineRouteRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *UpdateRoutineRouteRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *UpdateRoutineRouteRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateRoutineRouteRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateRoutineRouteRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateRoutineRouteRequest) SetBypass(v string) *UpdateRoutineRouteRequest {
	s.Bypass = &v
	return s
}

func (s *UpdateRoutineRouteRequest) SetConfigId(v int64) *UpdateRoutineRouteRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateRoutineRouteRequest) SetFallback(v string) *UpdateRoutineRouteRequest {
	s.Fallback = &v
	return s
}

func (s *UpdateRoutineRouteRequest) SetRouteEnable(v string) *UpdateRoutineRouteRequest {
	s.RouteEnable = &v
	return s
}

func (s *UpdateRoutineRouteRequest) SetRouteName(v string) *UpdateRoutineRouteRequest {
	s.RouteName = &v
	return s
}

func (s *UpdateRoutineRouteRequest) SetRoutineName(v string) *UpdateRoutineRouteRequest {
	s.RoutineName = &v
	return s
}

func (s *UpdateRoutineRouteRequest) SetRule(v string) *UpdateRoutineRouteRequest {
	s.Rule = &v
	return s
}

func (s *UpdateRoutineRouteRequest) SetSequence(v int32) *UpdateRoutineRouteRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateRoutineRouteRequest) SetSiteId(v int64) *UpdateRoutineRouteRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateRoutineRouteRequest) Validate() error {
	return dara.Validate(s)
}
