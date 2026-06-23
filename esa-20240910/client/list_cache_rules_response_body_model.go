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
	// The configuration list in the response body.
	Configs []*ListCacheRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The current page number, which is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
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
	// - Enables caching on specified ports.
	//
	// - Valid values: 8880, 2052, 2082, 2086, 2095, 2053, 2083, 2087, and 2096.
	//
	// - Multiple ports are separated by commas (,).
	//
	// example:
	//
	// 8880,2052,2086
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	// The browser cache mode. Valid values:
	//
	// - no_cache: does not cache.
	//
	// - follow_origin: follows the origin cache policy.
	//
	// - override_origin: overrides the origin cache policy.
	//
	// example:
	//
	// no_cache
	BrowserCacheMode *string `json:"BrowserCacheMode,omitempty" xml:"BrowserCacheMode,omitempty"`
	// The browser cache expiration time, in seconds.
	//
	// example:
	//
	// 300
	BrowserCacheTtl *string `json:"BrowserCacheTtl,omitempty" xml:"BrowserCacheTtl,omitempty"`
	// The bypass cache mode. Valid values:
	//
	// - cache_all: caches all requests.
	//
	// - bypass_all: bypasses cache for all requests.
	//
	// example:
	//
	// cache_all
	BypassCache *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	// The cache deception armor. Protects against web cache deception attacks by caching only content that passes validation. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	CacheDeceptionArmor *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	// The cache reserve eligibility. Controls whether user requests bypass the cache reserve node during back-to-origin. Valid values:
	//
	// - bypass_cache_reserve: requests bypass cache reserve.
	//
	// - eligible_for_cache_reserve: requests are eligible for cache reserve.
	//
	// example:
	//
	// bypass_cache_reserve
	CacheReserveEligibility *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	// When generating cache keys, checks whether the specified cookies exist. If a cookie exists, its name (case-insensitive) is added to the cache key. Multiple cookie names are separated by spaces. Cookie names support the following character types:
	//
	// - Symbols: ! # $ % & \\" 	- + - . ^ _ ` | ~
	//
	// - Digits: 0-9
	//
	// - Letters: lowercase a-z.
	//
	// example:
	//
	// cookiename1 cookiename2
	CheckPresenceCookie *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	// When generating cache keys, checks whether the specified headers exist. If a header exists, its name (case-insensitive) is added to the cache key. Multiple header names are separated by spaces. Header names support the following character types:
	//
	// - Symbols: ! # $ % & \\" 	- + - . ^ _ ` | ~
	//
	// - Digits: 0-9
	//
	// - Letters: lowercase a-z.
	//
	// example:
	//
	// headername1 headername2
	CheckPresenceHeader *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can use this parameter to query global or rule configurations. Valid values:
	//
	// - global: global configuration.
	//
	// - rule: rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The edge cache mode. Valid values:
	//
	// - follow_origin: follows the origin cache policy if present. Otherwise, uses the default cache policy.
	//
	// - no_cache: does not cache.
	//
	// - override_origin: overrides the origin cache policy.
	//
	// - follow_origin_bypass: follows the origin cache policy if present. Otherwise, does not cache.
	//
	// - follow_origin_override: follows the origin cache policy if present. Otherwise, uses a custom cache TTL.
	//
	// example:
	//
	// follow_origin
	EdgeCacheMode *string `json:"EdgeCacheMode,omitempty" xml:"EdgeCacheMode,omitempty"`
	// The edge cache expiration time, in seconds.
	//
	// example:
	//
	// 300
	EdgeCacheTtl *string `json:"EdgeCacheTtl,omitempty" xml:"EdgeCacheTtl,omitempty"`
	// The status code cache expiration time, in seconds.
	//
	// - You can set the cache expiration time for specific status codes. For example, 404=10 indicates that the 404 status code is cached for 10 seconds.
	//
	// - You can set the cache expiration time for 4xx or 5xx series status codes. For example, 4xx=10 indicates that all 4xx status codes are cached for 10 seconds.
	//
	// - You can set the cache expiration time for multiple status codes. Separate multiple status codes with commas (,).
	//
	// example:
	//
	// 5xx=0,404=10
	EdgeStatusCodeCacheTtl *string `json:"EdgeStatusCodeCacheTtl,omitempty" xml:"EdgeStatusCodeCacheTtl,omitempty"`
	// The specified cookie names (case-insensitive) and their values to include when generating cache keys. Multiple values are separated by spaces. Cookie names support the following character types:
	//
	// - Symbols: ! # $ % & \\" 	- + - . ^ _ ` | ~
	//
	// - Digits: 0-9
	//
	// - Letters: lowercase a-z.
	//
	// example:
	//
	// cookiename1 cookiename2
	IncludeCookie *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	// The specified header names (case-insensitive) and their values to include when generating cache keys. Multiple values are separated by spaces. Header names support the following character types:
	//
	// - Symbols: ! # $ % & \\" 	- + - . ^ _ ` | ~
	//
	// - Digits: 0-9
	//
	// - Letters: lowercase a-z.
	//
	// example:
	//
	// headername1 headername2
	IncludeHeader *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// The cache key handling mode for POST caching. The following two modes are supported:
	//
	// - md5: calculates the MD5 hash of the body content and adds the MD5 value to the cache key.
	//
	// - ignore: ignores the body content in the cache key.
	//
	// example:
	//
	// ignore
	PostBodyCacheKey *string `json:"PostBodyCacheKey,omitempty" xml:"PostBodyCacheKey,omitempty"`
	// The body size limit for POST caching. The value is a number in KB. Valid values: 1 to 8. If this parameter is left empty, the default value of 8 KB takes effect.
	//
	// example:
	//
	// 1
	PostBodySizeLimit *string `json:"PostBodySizeLimit,omitempty" xml:"PostBodySizeLimit,omitempty"`
	// Specifies whether to enable POST caching.
	//
	// example:
	//
	// on
	PostCache *string `json:"PostCache,omitempty" xml:"PostCache,omitempty"`
	// The query strings to retain or remove when generating cache keys. Multiple values are separated by spaces.
	//
	// example:
	//
	// example
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// The query string handling mode when generating cache keys. Valid values:
	//
	// - ignore_all: ignores all query strings.
	//
	// - exclude_query_string: removes specified query strings.
	//
	// - reserve_all: retains all query strings. This is the default value.
	//
	// - include_query_string: retains specified query strings.
	//
	// example:
	//
	// ignore_all
	QueryStringMode *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	// The rule content, which uses conditional expressions to match user requests. You do not need to set this parameter when you add a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: set the value to true.
	//
	// - Match specified requests: set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to set this parameter when you add a global configuration. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when you add a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution order. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to serve stale cache. When enabled, edge nodes can respond to user requests with cached expired files when the origin server is unavailable. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	ServeStale *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Specifies whether to sort query strings. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	// Specifies whether to include the type of the client when generating cache keys. Valid values:
	//
	// - on: enabled.
	//
	// - off: shutdown.
	//
	// example:
	//
	// on
	UserDeviceType *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	// Specifies whether to include the client geographic location when generating cache keys. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	UserGeo *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	// Specifies whether to include the client language type when generating cache keys. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
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
