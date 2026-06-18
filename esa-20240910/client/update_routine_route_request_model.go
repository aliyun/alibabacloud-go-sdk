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
	SetTimeout(v string) *UpdateRoutineRouteRequest
	GetTimeout() *string
}

type UpdateRoutineRouteRequest struct {
	// Specifies whether to enable bypass mode. Valid values:
	//
	// - on: Enabled
	//
	// - off: Disabled
	//
	// example:
	//
	// on
	Bypass *string `json:"Bypass,omitempty" xml:"Bypass,omitempty"`
	// The configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable fallback. If enabled, requests fall back to the origin server if the function encounters an exception, such as exceeding the CPU usage limit. Valid values:
	//
	// - on: Enabled
	//
	// - off: Disabled
	//
	// example:
	//
	// on
	Fallback *string `json:"Fallback,omitempty" xml:"Fallback,omitempty"`
	// Specifies whether to enable the route. Valid values:
	//
	// - on: Enabled
	//
	// - off: Disabled
	//
	// example:
	//
	// on
	RouteEnable *string `json:"RouteEnable,omitempty" xml:"RouteEnable,omitempty"`
	// The route name.
	//
	// example:
	//
	// test_route
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The name of the Routine.
	//
	// example:
	//
	// test-routine1
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
	// The content of the rule.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The execution sequence of the rule.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The ID of the site. You can obtain this ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456******
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 5
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
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

func (s *UpdateRoutineRouteRequest) GetTimeout() *string {
	return s.Timeout
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

func (s *UpdateRoutineRouteRequest) SetTimeout(v string) *UpdateRoutineRouteRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateRoutineRouteRequest) Validate() error {
	return dara.Validate(s)
}
