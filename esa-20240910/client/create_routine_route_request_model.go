// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBypass(v string) *CreateRoutineRouteRequest
	GetBypass() *string
	SetFallback(v string) *CreateRoutineRouteRequest
	GetFallback() *string
	SetRouteEnable(v string) *CreateRoutineRouteRequest
	GetRouteEnable() *string
	SetRouteName(v string) *CreateRoutineRouteRequest
	GetRouteName() *string
	SetRoutineName(v string) *CreateRoutineRouteRequest
	GetRoutineName() *string
	SetRule(v string) *CreateRoutineRouteRequest
	GetRule() *string
	SetSequence(v int32) *CreateRoutineRouteRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateRoutineRouteRequest
	GetSiteId() *int64
}

type CreateRoutineRouteRequest struct {
	// example:
	//
	// on
	Bypass   *string `json:"Bypass,omitempty" xml:"Bypass,omitempty"`
	Fallback *string `json:"Fallback,omitempty" xml:"Fallback,omitempty"`
	// example:
	//
	// on
	RouteEnable *string `json:"RouteEnable,omitempty" xml:"RouteEnable,omitempty"`
	// example:
	//
	// test_route
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// This parameter is required.
	//
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
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateRoutineRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineRouteRequest) GetBypass() *string {
	return s.Bypass
}

func (s *CreateRoutineRouteRequest) GetFallback() *string {
	return s.Fallback
}

func (s *CreateRoutineRouteRequest) GetRouteEnable() *string {
	return s.RouteEnable
}

func (s *CreateRoutineRouteRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *CreateRoutineRouteRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *CreateRoutineRouteRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateRoutineRouteRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateRoutineRouteRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateRoutineRouteRequest) SetBypass(v string) *CreateRoutineRouteRequest {
	s.Bypass = &v
	return s
}

func (s *CreateRoutineRouteRequest) SetFallback(v string) *CreateRoutineRouteRequest {
	s.Fallback = &v
	return s
}

func (s *CreateRoutineRouteRequest) SetRouteEnable(v string) *CreateRoutineRouteRequest {
	s.RouteEnable = &v
	return s
}

func (s *CreateRoutineRouteRequest) SetRouteName(v string) *CreateRoutineRouteRequest {
	s.RouteName = &v
	return s
}

func (s *CreateRoutineRouteRequest) SetRoutineName(v string) *CreateRoutineRouteRequest {
	s.RoutineName = &v
	return s
}

func (s *CreateRoutineRouteRequest) SetRule(v string) *CreateRoutineRouteRequest {
	s.Rule = &v
	return s
}

func (s *CreateRoutineRouteRequest) SetSequence(v int32) *CreateRoutineRouteRequest {
	s.Sequence = &v
	return s
}

func (s *CreateRoutineRouteRequest) SetSiteId(v int64) *CreateRoutineRouteRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRoutineRouteRequest) Validate() error {
	return dara.Validate(s)
}
