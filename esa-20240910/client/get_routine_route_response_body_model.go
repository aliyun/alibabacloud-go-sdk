// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBypass(v string) *GetRoutineRouteResponseBody
	GetBypass() *string
	SetConfigId(v int64) *GetRoutineRouteResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetRoutineRouteResponseBody
	GetConfigType() *string
	SetFallback(v string) *GetRoutineRouteResponseBody
	GetFallback() *string
	SetMode(v string) *GetRoutineRouteResponseBody
	GetMode() *string
	SetRequestId(v string) *GetRoutineRouteResponseBody
	GetRequestId() *string
	SetRouteEnable(v string) *GetRoutineRouteResponseBody
	GetRouteEnable() *string
	SetRouteName(v string) *GetRoutineRouteResponseBody
	GetRouteName() *string
	SetRoutineName(v string) *GetRoutineRouteResponseBody
	GetRoutineName() *string
	SetRule(v string) *GetRoutineRouteResponseBody
	GetRule() *string
	SetSequence(v int32) *GetRoutineRouteResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetRoutineRouteResponseBody
	GetSiteVersion() *int32
}

type GetRoutineRouteResponseBody struct {
	// example:
	//
	// on
	Bypass *string `json:"Bypass,omitempty" xml:"Bypass,omitempty"`
	// example:
	//
	// 352816******
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	Fallback   *string `json:"Fallback,omitempty" xml:"Fallback,omitempty"`
	// example:
	//
	// simple
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetRoutineRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoutineRouteResponseBody) GetBypass() *string {
	return s.Bypass
}

func (s *GetRoutineRouteResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetRoutineRouteResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetRoutineRouteResponseBody) GetFallback() *string {
	return s.Fallback
}

func (s *GetRoutineRouteResponseBody) GetMode() *string {
	return s.Mode
}

func (s *GetRoutineRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoutineRouteResponseBody) GetRouteEnable() *string {
	return s.RouteEnable
}

func (s *GetRoutineRouteResponseBody) GetRouteName() *string {
	return s.RouteName
}

func (s *GetRoutineRouteResponseBody) GetRoutineName() *string {
	return s.RoutineName
}

func (s *GetRoutineRouteResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetRoutineRouteResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetRoutineRouteResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetRoutineRouteResponseBody) SetBypass(v string) *GetRoutineRouteResponseBody {
	s.Bypass = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetConfigId(v int64) *GetRoutineRouteResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetConfigType(v string) *GetRoutineRouteResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetFallback(v string) *GetRoutineRouteResponseBody {
	s.Fallback = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetMode(v string) *GetRoutineRouteResponseBody {
	s.Mode = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetRequestId(v string) *GetRoutineRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetRouteEnable(v string) *GetRoutineRouteResponseBody {
	s.RouteEnable = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetRouteName(v string) *GetRoutineRouteResponseBody {
	s.RouteName = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetRoutineName(v string) *GetRoutineRouteResponseBody {
	s.RoutineName = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetRule(v string) *GetRoutineRouteResponseBody {
	s.Rule = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetSequence(v int32) *GetRoutineRouteResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetRoutineRouteResponseBody) SetSiteVersion(v int32) *GetRoutineRouteResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetRoutineRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
