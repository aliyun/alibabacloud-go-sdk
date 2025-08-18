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
	// Response body configuration.
	Configs []*ListCacheRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	return dara.Validate(s)
}

type ListCacheRulesResponseBodyConfigs struct {
	// Enable caching on specified ports. Value range: 8880, 2052, 2082, 2086, 2095, 2053, 2083, 2087, 2096.
	//
	// example:
	//
	// 2082
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	// Browser cache mode. Possible values:
	//
	// - no_cache: Do not cache.
	//
	// - follow_origin: Follow origin cache policy.
	//
	// - override_origin: Override origin cache policy.
	//
	// example:
	//
	// no_cache
	BrowserCacheMode *string `json:"BrowserCacheMode,omitempty" xml:"BrowserCacheMode,omitempty"`
	// Browser cache expiration time, in seconds.
	//
	// example:
	//
	// 300
	BrowserCacheTtl *string `json:"BrowserCacheTtl,omitempty" xml:"BrowserCacheTtl,omitempty"`
	// Set bypass cache mode. Possible values:
	//
	// - cache_all: Cache all requests.
	//
	// - bypass_all: Bypass cache for all requests.
	//
	// example:
	//
	// cache_all
	BypassCache *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	// Cache deception defense. Used to defend against web cache deception attacks; only verified cache content will be cached. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	CacheDeceptionArmor *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	// Cache reserve eligibility. This is used to control whether user requests bypass the cache reserve node when returning to the origin. The value range is as follows:
	//
	// - bypass_cache_reserve: Requests bypass the cache reserve.
	//
	// - eligible_for_cache_reserve: Eligible for cache reserve.
	//
	// example:
	//
	// bypass_cache_reserve
	CacheReserveEligibility *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	// When generating the cache key, check if the cookie exists. If it does, add the cookie name (case-insensitive) to the cache key. Multiple cookie names are supported, separated by spaces.
	//
	// example:
	//
	// cookiename
	CheckPresenceCookie *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	// When generating the cache key, check if the header exists. If it does, add the header name (case-insensitive) to the cache key. Multiple header names are supported, separated by spaces.
	//
	// example:
	//
	// headername
	CheckPresenceHeader *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
	// Configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule-based configurations. Possible values:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule-based configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Edge cache mode. The value range is as follows:
	//
	// - follow_origin: Follow the origin server\\"s cache policy (if it exists), otherwise use the default cache policy.
	//
	// - no_cache: Do not cache.
	//
	// - override_origin: Override the origin server\\"s cache policy.
	//
	// - follow_origin_bypass: Follow the origin server\\"s cache policy (if it exists), otherwise do not cache.
	//
	// example:
	//
	// follow_origin
	EdgeCacheMode *string `json:"EdgeCacheMode,omitempty" xml:"EdgeCacheMode,omitempty"`
	// Edge cache expiration time, in seconds.
	//
	// example:
	//
	// 300
	EdgeCacheTtl *string `json:"EdgeCacheTtl,omitempty" xml:"EdgeCacheTtl,omitempty"`
	// Edge cache expiration time, in seconds.
	//
	// example:
	//
	// 300
	EdgeStatusCodeCacheTtl *string `json:"EdgeStatusCodeCacheTtl,omitempty" xml:"EdgeStatusCodeCacheTtl,omitempty"`
	// Include the specified cookie names and their values when generating the cache key. Multiple values are supported, separated by spaces.
	//
	// example:
	//
	// cookie_exapmle
	IncludeCookie *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	// Include the specified header names and their values when generating the cache key. Multiple values are supported, separated by spaces.
	//
	// example:
	//
	// example
	IncludeHeader *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// The query strings to be reserved or excluded. Multiple values are supported, separated by spaces.
	//
	// example:
	//
	// example
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// The processing mode for query strings when generating the cache key. The value range is as follows:
	//
	// - ignore_all: Ignore all query strings.
	//
	// - exclude_query_string: Exclude specified query strings.
	//
	// - reserve_all: Default, reserve all query strings.
	//
	// - include_query_string: Include specified query strings.
	//
	// example:
	//
	// ignore_all
	QueryStringMode *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Serve stale cache. When enabled, the node can still respond to user requests with expired cached files even when the origin server is unavailable. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	ServeStale *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	// Site configuration version number. For sites with version management enabled, this parameter can specify the site version for which the configuration takes effect, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Query string sorting. The value range is as follows:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	// Include the client device type when generating the cache key. The value range is as follows:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	UserDeviceType *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	// Include the client\\"s geographic location when generating the cache key. The value range is as follows:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	UserGeo *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	// Include the client\\"s language type when generating the cache key. The value range is as follows:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	UserLanguage *string `json:"UserLanguage,omitempty" xml:"UserLanguage,omitempty"`
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
