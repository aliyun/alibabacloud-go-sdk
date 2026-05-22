// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCacheRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListCacheRulesResponseBodyConfigs) *ListCacheRulesResponseBody
	GetConfigs() []*ListCacheRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListCacheRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCacheRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCacheRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCacheRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListCacheRulesResponseBody
	GetTotalPage() *int32
}

type ListCacheRulesResponseBody struct {
	Configs    []*ListCacheRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListCacheRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCacheRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCacheRulesResponseBody) GetConfigs() []*ListCacheRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListCacheRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCacheRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCacheRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCacheRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCacheRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListCacheRulesResponseBody) SetConfigs(v []*ListCacheRulesResponseBodyConfigs) *ListCacheRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListCacheRulesResponseBody) SetPageNumber(v int32) *ListCacheRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCacheRulesResponseBody) SetPageSize(v int32) *ListCacheRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCacheRulesResponseBody) SetRequestId(v string) *ListCacheRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCacheRulesResponseBody) SetTotalCount(v int32) *ListCacheRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCacheRulesResponseBody) SetTotalPage(v int32) *ListCacheRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListCacheRulesResponseBody) Validate() error {
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

type ListCacheRulesResponseBodyConfigs struct {
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	BrowserCacheMode         *string `json:"BrowserCacheMode,omitempty" xml:"BrowserCacheMode,omitempty"`
	BrowserCacheTtl          *string `json:"BrowserCacheTtl,omitempty" xml:"BrowserCacheTtl,omitempty"`
	BypassCache              *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	CacheDeceptionArmor      *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	CacheReserveEligibility  *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	CheckPresenceCookie      *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	CheckPresenceHeader      *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
	ConfigId                 *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType               *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	EdgeCacheMode            *string `json:"EdgeCacheMode,omitempty" xml:"EdgeCacheMode,omitempty"`
	EdgeCacheTtl             *string `json:"EdgeCacheTtl,omitempty" xml:"EdgeCacheTtl,omitempty"`
	EdgeStatusCodeCacheTtl   *string `json:"EdgeStatusCodeCacheTtl,omitempty" xml:"EdgeStatusCodeCacheTtl,omitempty"`
	IncludeCookie            *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	IncludeHeader            *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// example:
	//
	// ignore
	PostBodyCacheKey  *string `json:"PostBodyCacheKey,omitempty" xml:"PostBodyCacheKey,omitempty"`
	PostBodySizeLimit *string `json:"PostBodySizeLimit,omitempty" xml:"PostBodySizeLimit,omitempty"`
	// example:
	//
	// on
	PostCache               *string `json:"PostCache,omitempty" xml:"PostCache,omitempty"`
	QueryString             *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	QueryStringMode         *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	Rule                    *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable              *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	ServeStale              *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	SiteVersion             *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	UserDeviceType          *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	UserGeo                 *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	UserLanguage            *string `json:"UserLanguage,omitempty" xml:"UserLanguage,omitempty"`
}

func (s ListCacheRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListCacheRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListCacheRulesResponseBodyConfigs) GetAdditionalCacheablePorts() *string {
	return s.AdditionalCacheablePorts
}

func (s *ListCacheRulesResponseBodyConfigs) GetBrowserCacheMode() *string {
	return s.BrowserCacheMode
}

func (s *ListCacheRulesResponseBodyConfigs) GetBrowserCacheTtl() *string {
	return s.BrowserCacheTtl
}

func (s *ListCacheRulesResponseBodyConfigs) GetBypassCache() *string {
	return s.BypassCache
}

func (s *ListCacheRulesResponseBodyConfigs) GetCacheDeceptionArmor() *string {
	return s.CacheDeceptionArmor
}

func (s *ListCacheRulesResponseBodyConfigs) GetCacheReserveEligibility() *string {
	return s.CacheReserveEligibility
}

func (s *ListCacheRulesResponseBodyConfigs) GetCheckPresenceCookie() *string {
	return s.CheckPresenceCookie
}

func (s *ListCacheRulesResponseBodyConfigs) GetCheckPresenceHeader() *string {
	return s.CheckPresenceHeader
}

func (s *ListCacheRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListCacheRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListCacheRulesResponseBodyConfigs) GetEdgeCacheMode() *string {
	return s.EdgeCacheMode
}

func (s *ListCacheRulesResponseBodyConfigs) GetEdgeCacheTtl() *string {
	return s.EdgeCacheTtl
}

func (s *ListCacheRulesResponseBodyConfigs) GetEdgeStatusCodeCacheTtl() *string {
	return s.EdgeStatusCodeCacheTtl
}

func (s *ListCacheRulesResponseBodyConfigs) GetIncludeCookie() *string {
	return s.IncludeCookie
}

func (s *ListCacheRulesResponseBodyConfigs) GetIncludeHeader() *string {
	return s.IncludeHeader
}

func (s *ListCacheRulesResponseBodyConfigs) GetPostBodyCacheKey() *string {
	return s.PostBodyCacheKey
}

func (s *ListCacheRulesResponseBodyConfigs) GetPostBodySizeLimit() *string {
	return s.PostBodySizeLimit
}

func (s *ListCacheRulesResponseBodyConfigs) GetPostCache() *string {
	return s.PostCache
}

func (s *ListCacheRulesResponseBodyConfigs) GetQueryString() *string {
	return s.QueryString
}

func (s *ListCacheRulesResponseBodyConfigs) GetQueryStringMode() *string {
	return s.QueryStringMode
}

func (s *ListCacheRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListCacheRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListCacheRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCacheRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListCacheRulesResponseBodyConfigs) GetServeStale() *string {
	return s.ServeStale
}

func (s *ListCacheRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListCacheRulesResponseBodyConfigs) GetSortQueryStringForCache() *string {
	return s.SortQueryStringForCache
}

func (s *ListCacheRulesResponseBodyConfigs) GetUserDeviceType() *string {
	return s.UserDeviceType
}

func (s *ListCacheRulesResponseBodyConfigs) GetUserGeo() *string {
	return s.UserGeo
}

func (s *ListCacheRulesResponseBodyConfigs) GetUserLanguage() *string {
	return s.UserLanguage
}

func (s *ListCacheRulesResponseBodyConfigs) SetAdditionalCacheablePorts(v string) *ListCacheRulesResponseBodyConfigs {
	s.AdditionalCacheablePorts = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetBrowserCacheMode(v string) *ListCacheRulesResponseBodyConfigs {
	s.BrowserCacheMode = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetBrowserCacheTtl(v string) *ListCacheRulesResponseBodyConfigs {
	s.BrowserCacheTtl = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetBypassCache(v string) *ListCacheRulesResponseBodyConfigs {
	s.BypassCache = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetCacheDeceptionArmor(v string) *ListCacheRulesResponseBodyConfigs {
	s.CacheDeceptionArmor = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetCacheReserveEligibility(v string) *ListCacheRulesResponseBodyConfigs {
	s.CacheReserveEligibility = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetCheckPresenceCookie(v string) *ListCacheRulesResponseBodyConfigs {
	s.CheckPresenceCookie = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetCheckPresenceHeader(v string) *ListCacheRulesResponseBodyConfigs {
	s.CheckPresenceHeader = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetConfigId(v int64) *ListCacheRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetConfigType(v string) *ListCacheRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetEdgeCacheMode(v string) *ListCacheRulesResponseBodyConfigs {
	s.EdgeCacheMode = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetEdgeCacheTtl(v string) *ListCacheRulesResponseBodyConfigs {
	s.EdgeCacheTtl = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetEdgeStatusCodeCacheTtl(v string) *ListCacheRulesResponseBodyConfigs {
	s.EdgeStatusCodeCacheTtl = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetIncludeCookie(v string) *ListCacheRulesResponseBodyConfigs {
	s.IncludeCookie = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetIncludeHeader(v string) *ListCacheRulesResponseBodyConfigs {
	s.IncludeHeader = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetPostBodyCacheKey(v string) *ListCacheRulesResponseBodyConfigs {
	s.PostBodyCacheKey = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetPostBodySizeLimit(v string) *ListCacheRulesResponseBodyConfigs {
	s.PostBodySizeLimit = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetPostCache(v string) *ListCacheRulesResponseBodyConfigs {
	s.PostCache = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetQueryString(v string) *ListCacheRulesResponseBodyConfigs {
	s.QueryString = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetQueryStringMode(v string) *ListCacheRulesResponseBodyConfigs {
	s.QueryStringMode = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetRule(v string) *ListCacheRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetRuleEnable(v string) *ListCacheRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetRuleName(v string) *ListCacheRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetSequence(v int32) *ListCacheRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetServeStale(v string) *ListCacheRulesResponseBodyConfigs {
	s.ServeStale = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListCacheRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetSortQueryStringForCache(v string) *ListCacheRulesResponseBodyConfigs {
	s.SortQueryStringForCache = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetUserDeviceType(v string) *ListCacheRulesResponseBodyConfigs {
	s.UserDeviceType = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetUserGeo(v string) *ListCacheRulesResponseBodyConfigs {
	s.UserGeo = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) SetUserLanguage(v string) *ListCacheRulesResponseBodyConfigs {
	s.UserLanguage = &v
	return s
}

func (s *ListCacheRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
