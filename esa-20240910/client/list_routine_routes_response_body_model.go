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
	Configs    []*ListRoutineRoutesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                  `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRoutineRoutesResponseBodyConfigs struct {
	Bypass      *string `json:"Bypass,omitempty" xml:"Bypass,omitempty"`
	ConfigId    *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType  *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	Fallback    *string `json:"Fallback,omitempty" xml:"Fallback,omitempty"`
	Mode        *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RouteEnable *string `json:"RouteEnable,omitempty" xml:"RouteEnable,omitempty"`
	RouteName   *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
	Rule        *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	Sequence    *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteId      *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName    *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	SiteVersion *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// example:
	//
	// 5
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
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

func (s *ListRoutineRoutesResponseBodyConfigs) GetTimeout() *string {
	return s.Timeout
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

func (s *ListRoutineRoutesResponseBodyConfigs) SetTimeout(v string) *ListRoutineRoutesResponseBodyConfigs {
	s.Timeout = &v
	return s
}

func (s *ListRoutineRoutesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
