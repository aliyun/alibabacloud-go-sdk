// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListRoutineRoutesResponseBodyConfigs) *ListRoutineRoutesResponseBody
	GetConfigs() []*ListRoutineRoutesResponseBodyConfigs
	SetPageNumber(v int32) *ListRoutineRoutesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRoutineRoutesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRoutineRoutesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListRoutineRoutesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListRoutineRoutesResponseBody
	GetTotalPage() *int32
}

type ListRoutineRoutesResponseBody struct {
	// The returned configurations.
	Configs []*ListRoutineRoutesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 83
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListRoutineRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoutineRoutesResponseBody) GetConfigs() []*ListRoutineRoutesResponseBodyConfigs {
	return s.Configs
}

func (s *ListRoutineRoutesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRoutineRoutesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRoutineRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoutineRoutesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRoutineRoutesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListRoutineRoutesResponseBody) SetConfigs(v []*ListRoutineRoutesResponseBodyConfigs) *ListRoutineRoutesResponseBody {
	s.Configs = v
	return s
}

func (s *ListRoutineRoutesResponseBody) SetPageNumber(v int32) *ListRoutineRoutesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRoutineRoutesResponseBody) SetPageSize(v int32) *ListRoutineRoutesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRoutineRoutesResponseBody) SetRequestId(v string) *ListRoutineRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoutineRoutesResponseBody) SetTotalCount(v int32) *ListRoutineRoutesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRoutineRoutesResponseBody) SetTotalPage(v int32) *ListRoutineRoutesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListRoutineRoutesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRoutineRoutesResponseBodyConfigs struct {
	// Bypass mode. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	Bypass *string `json:"Bypass,omitempty" xml:"Bypass,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type to query. Valid values:
	//
	// 	- global: global configurations.
	//
	// 	- rule: queries rule configurations.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The exception origin switch. After you turn on this switch, if a function exception occurs, such as CPU usage exceeding the upper limit, requests are sent back to the origin. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	Fallback *string `json:"Fallback,omitempty" xml:"Fallback,omitempty"`
	// The configuration mode. Valid values:
	//
	// 	- simple: Simple mode.
	//
	// 	- custom: Custom mode.
	//
	// example:
	//
	// simple
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The route switch. Valid values:
	//
	// 	- on
	//
	// 	- off
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
	// The routine name.
	//
	// example:
	//
	// test-routine1
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The order in which the rule is executed.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 554889455535696
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The website name.
	//
	// example:
	//
	// test.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The version number of the website configurations.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListRoutineRoutesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineRoutesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetBypass() *string {
	return s.Bypass
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetFallback() *string {
	return s.Fallback
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetMode() *string {
	return s.Mode
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetRouteEnable() *string {
	return s.RouteEnable
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetRouteName() *string {
	return s.RouteName
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetRoutineName() *string {
	return s.RoutineName
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetSiteName() *string {
	return s.SiteName
}

func (s *ListRoutineRoutesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetBypass(v string) *ListRoutineRoutesResponseBodyConfigs {
	s.Bypass = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetConfigId(v int64) *ListRoutineRoutesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetConfigType(v string) *ListRoutineRoutesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetFallback(v string) *ListRoutineRoutesResponseBodyConfigs {
	s.Fallback = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetMode(v string) *ListRoutineRoutesResponseBodyConfigs {
	s.Mode = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetRouteEnable(v string) *ListRoutineRoutesResponseBodyConfigs {
	s.RouteEnable = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetRouteName(v string) *ListRoutineRoutesResponseBodyConfigs {
	s.RouteName = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetRoutineName(v string) *ListRoutineRoutesResponseBodyConfigs {
	s.RoutineName = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetRule(v string) *ListRoutineRoutesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetSequence(v int32) *ListRoutineRoutesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetSiteId(v int64) *ListRoutineRoutesResponseBodyConfigs {
	s.SiteId = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetSiteName(v string) *ListRoutineRoutesResponseBodyConfigs {
	s.SiteName = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) SetSiteVersion(v int32) *ListRoutineRoutesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
