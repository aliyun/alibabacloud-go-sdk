// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalCacheablePorts(v string) *GetCacheRuleResponseBody
	GetAdditionalCacheablePorts() *string
	SetBrowserCacheMode(v string) *GetCacheRuleResponseBody
	GetBrowserCacheMode() *string
	SetBrowserCacheTtl(v string) *GetCacheRuleResponseBody
	GetBrowserCacheTtl() *string
	SetBypassCache(v string) *GetCacheRuleResponseBody
	GetBypassCache() *string
	SetCacheDeceptionArmor(v string) *GetCacheRuleResponseBody
	GetCacheDeceptionArmor() *string
	SetCacheReserveEligibility(v string) *GetCacheRuleResponseBody
	GetCacheReserveEligibility() *string
	SetCheckPresenceCookie(v string) *GetCacheRuleResponseBody
	GetCheckPresenceCookie() *string
	SetCheckPresenceHeader(v string) *GetCacheRuleResponseBody
	GetCheckPresenceHeader() *string
	SetConfigId(v int64) *GetCacheRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetCacheRuleResponseBody
	GetConfigType() *string
	SetEdgeCacheMode(v string) *GetCacheRuleResponseBody
	GetEdgeCacheMode() *string
	SetEdgeCacheTtl(v string) *GetCacheRuleResponseBody
	GetEdgeCacheTtl() *string
	SetEdgeStatusCodeCacheTtl(v string) *GetCacheRuleResponseBody
	GetEdgeStatusCodeCacheTtl() *string
	SetIncludeCookie(v string) *GetCacheRuleResponseBody
	GetIncludeCookie() *string
	SetIncludeHeader(v string) *GetCacheRuleResponseBody
	GetIncludeHeader() *string
	SetPostBodyCacheKey(v string) *GetCacheRuleResponseBody
	GetPostBodyCacheKey() *string
	SetPostBodySizeLimit(v string) *GetCacheRuleResponseBody
	GetPostBodySizeLimit() *string
	SetPostCache(v string) *GetCacheRuleResponseBody
	GetPostCache() *string
	SetQueryString(v string) *GetCacheRuleResponseBody
	GetQueryString() *string
	SetQueryStringMode(v string) *GetCacheRuleResponseBody
	GetQueryStringMode() *string
	SetRequestId(v string) *GetCacheRuleResponseBody
	GetRequestId() *string
	SetRule(v string) *GetCacheRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetCacheRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetCacheRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetCacheRuleResponseBody
	GetSequence() *int32
	SetServeStale(v string) *GetCacheRuleResponseBody
	GetServeStale() *string
	SetSiteVersion(v int32) *GetCacheRuleResponseBody
	GetSiteVersion() *int32
	SetSortQueryStringForCache(v string) *GetCacheRuleResponseBody
	GetSortQueryStringForCache() *string
	SetUserDeviceType(v string) *GetCacheRuleResponseBody
	GetUserDeviceType() *string
	SetUserGeo(v string) *GetCacheRuleResponseBody
	GetUserGeo() *string
	SetUserLanguage(v string) *GetCacheRuleResponseBody
	GetUserLanguage() *string
}

type GetCacheRuleResponseBody struct {
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
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s GetCacheRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCacheRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetCacheRuleResponseBody) GetAdditionalCacheablePorts() *string {
	return s.AdditionalCacheablePorts
}

func (s *GetCacheRuleResponseBody) GetBrowserCacheMode() *string {
	return s.BrowserCacheMode
}

func (s *GetCacheRuleResponseBody) GetBrowserCacheTtl() *string {
	return s.BrowserCacheTtl
}

func (s *GetCacheRuleResponseBody) GetBypassCache() *string {
	return s.BypassCache
}

func (s *GetCacheRuleResponseBody) GetCacheDeceptionArmor() *string {
	return s.CacheDeceptionArmor
}

func (s *GetCacheRuleResponseBody) GetCacheReserveEligibility() *string {
	return s.CacheReserveEligibility
}

func (s *GetCacheRuleResponseBody) GetCheckPresenceCookie() *string {
	return s.CheckPresenceCookie
}

func (s *GetCacheRuleResponseBody) GetCheckPresenceHeader() *string {
	return s.CheckPresenceHeader
}

func (s *GetCacheRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetCacheRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetCacheRuleResponseBody) GetEdgeCacheMode() *string {
	return s.EdgeCacheMode
}

func (s *GetCacheRuleResponseBody) GetEdgeCacheTtl() *string {
	return s.EdgeCacheTtl
}

func (s *GetCacheRuleResponseBody) GetEdgeStatusCodeCacheTtl() *string {
	return s.EdgeStatusCodeCacheTtl
}

func (s *GetCacheRuleResponseBody) GetIncludeCookie() *string {
	return s.IncludeCookie
}

func (s *GetCacheRuleResponseBody) GetIncludeHeader() *string {
	return s.IncludeHeader
}

func (s *GetCacheRuleResponseBody) GetPostBodyCacheKey() *string {
	return s.PostBodyCacheKey
}

func (s *GetCacheRuleResponseBody) GetPostBodySizeLimit() *string {
	return s.PostBodySizeLimit
}

func (s *GetCacheRuleResponseBody) GetPostCache() *string {
	return s.PostCache
}

func (s *GetCacheRuleResponseBody) GetQueryString() *string {
	return s.QueryString
}

func (s *GetCacheRuleResponseBody) GetQueryStringMode() *string {
	return s.QueryStringMode
}

func (s *GetCacheRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCacheRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetCacheRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetCacheRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetCacheRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetCacheRuleResponseBody) GetServeStale() *string {
	return s.ServeStale
}

func (s *GetCacheRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetCacheRuleResponseBody) GetSortQueryStringForCache() *string {
	return s.SortQueryStringForCache
}

func (s *GetCacheRuleResponseBody) GetUserDeviceType() *string {
	return s.UserDeviceType
}

func (s *GetCacheRuleResponseBody) GetUserGeo() *string {
	return s.UserGeo
}

func (s *GetCacheRuleResponseBody) GetUserLanguage() *string {
	return s.UserLanguage
}

func (s *GetCacheRuleResponseBody) SetAdditionalCacheablePorts(v string) *GetCacheRuleResponseBody {
	s.AdditionalCacheablePorts = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetBrowserCacheMode(v string) *GetCacheRuleResponseBody {
	s.BrowserCacheMode = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetBrowserCacheTtl(v string) *GetCacheRuleResponseBody {
	s.BrowserCacheTtl = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetBypassCache(v string) *GetCacheRuleResponseBody {
	s.BypassCache = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetCacheDeceptionArmor(v string) *GetCacheRuleResponseBody {
	s.CacheDeceptionArmor = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetCacheReserveEligibility(v string) *GetCacheRuleResponseBody {
	s.CacheReserveEligibility = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetCheckPresenceCookie(v string) *GetCacheRuleResponseBody {
	s.CheckPresenceCookie = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetCheckPresenceHeader(v string) *GetCacheRuleResponseBody {
	s.CheckPresenceHeader = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetConfigId(v int64) *GetCacheRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetConfigType(v string) *GetCacheRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetEdgeCacheMode(v string) *GetCacheRuleResponseBody {
	s.EdgeCacheMode = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetEdgeCacheTtl(v string) *GetCacheRuleResponseBody {
	s.EdgeCacheTtl = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetEdgeStatusCodeCacheTtl(v string) *GetCacheRuleResponseBody {
	s.EdgeStatusCodeCacheTtl = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetIncludeCookie(v string) *GetCacheRuleResponseBody {
	s.IncludeCookie = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetIncludeHeader(v string) *GetCacheRuleResponseBody {
	s.IncludeHeader = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetPostBodyCacheKey(v string) *GetCacheRuleResponseBody {
	s.PostBodyCacheKey = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetPostBodySizeLimit(v string) *GetCacheRuleResponseBody {
	s.PostBodySizeLimit = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetPostCache(v string) *GetCacheRuleResponseBody {
	s.PostCache = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetQueryString(v string) *GetCacheRuleResponseBody {
	s.QueryString = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetQueryStringMode(v string) *GetCacheRuleResponseBody {
	s.QueryStringMode = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetRequestId(v string) *GetCacheRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetRule(v string) *GetCacheRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetRuleEnable(v string) *GetCacheRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetRuleName(v string) *GetCacheRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetSequence(v int32) *GetCacheRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetServeStale(v string) *GetCacheRuleResponseBody {
	s.ServeStale = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetSiteVersion(v int32) *GetCacheRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetSortQueryStringForCache(v string) *GetCacheRuleResponseBody {
	s.SortQueryStringForCache = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetUserDeviceType(v string) *GetCacheRuleResponseBody {
	s.UserDeviceType = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetUserGeo(v string) *GetCacheRuleResponseBody {
	s.UserGeo = &v
	return s
}

func (s *GetCacheRuleResponseBody) SetUserLanguage(v string) *GetCacheRuleResponseBody {
	s.UserLanguage = &v
	return s
}

func (s *GetCacheRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
