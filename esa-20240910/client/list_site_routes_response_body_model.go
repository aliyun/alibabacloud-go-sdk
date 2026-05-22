// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListSiteRoutesResponseBodyConfigs) *ListSiteRoutesResponseBody
	GetConfigs() []*ListSiteRoutesResponseBodyConfigs
	SetPageNumber(v int32) *ListSiteRoutesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSiteRoutesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSiteRoutesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListSiteRoutesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListSiteRoutesResponseBody
	GetTotalPage() *int32
}

type ListSiteRoutesResponseBody struct {
	Configs    []*ListSiteRoutesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListSiteRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSiteRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSiteRoutesResponseBody) GetConfigs() []*ListSiteRoutesResponseBodyConfigs {
	return s.Configs
}

func (s *ListSiteRoutesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSiteRoutesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSiteRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSiteRoutesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSiteRoutesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListSiteRoutesResponseBody) SetConfigs(v []*ListSiteRoutesResponseBodyConfigs) *ListSiteRoutesResponseBody {
	s.Configs = v
	return s
}

func (s *ListSiteRoutesResponseBody) SetPageNumber(v int32) *ListSiteRoutesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSiteRoutesResponseBody) SetPageSize(v int32) *ListSiteRoutesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSiteRoutesResponseBody) SetRequestId(v string) *ListSiteRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSiteRoutesResponseBody) SetTotalCount(v int32) *ListSiteRoutesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSiteRoutesResponseBody) SetTotalPage(v int32) *ListSiteRoutesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListSiteRoutesResponseBody) Validate() error {
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

type ListSiteRoutesResponseBodyConfigs struct {
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
	SiteVersion *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	Timeout     *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListSiteRoutesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListSiteRoutesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListSiteRoutesResponseBodyConfigs) GetBypass() *string {
	return s.Bypass
}

func (s *ListSiteRoutesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteRoutesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListSiteRoutesResponseBodyConfigs) GetFallback() *string {
	return s.Fallback
}

func (s *ListSiteRoutesResponseBodyConfigs) GetMode() *string {
	return s.Mode
}

func (s *ListSiteRoutesResponseBodyConfigs) GetRouteEnable() *string {
	return s.RouteEnable
}

func (s *ListSiteRoutesResponseBodyConfigs) GetRouteName() *string {
	return s.RouteName
}

func (s *ListSiteRoutesResponseBodyConfigs) GetRoutineName() *string {
	return s.RoutineName
}

func (s *ListSiteRoutesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListSiteRoutesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListSiteRoutesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListSiteRoutesResponseBodyConfigs) GetTimeout() *string {
	return s.Timeout
}

func (s *ListSiteRoutesResponseBodyConfigs) SetBypass(v string) *ListSiteRoutesResponseBodyConfigs {
	s.Bypass = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetConfigId(v int64) *ListSiteRoutesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetConfigType(v string) *ListSiteRoutesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetFallback(v string) *ListSiteRoutesResponseBodyConfigs {
	s.Fallback = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetMode(v string) *ListSiteRoutesResponseBodyConfigs {
	s.Mode = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetRouteEnable(v string) *ListSiteRoutesResponseBodyConfigs {
	s.RouteEnable = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetRouteName(v string) *ListSiteRoutesResponseBodyConfigs {
	s.RouteName = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetRoutineName(v string) *ListSiteRoutesResponseBodyConfigs {
	s.RoutineName = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetRule(v string) *ListSiteRoutesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetSequence(v int32) *ListSiteRoutesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetSiteVersion(v int32) *ListSiteRoutesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) SetTimeout(v string) *ListSiteRoutesResponseBodyConfigs {
	s.Timeout = &v
	return s
}

func (s *ListSiteRoutesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
