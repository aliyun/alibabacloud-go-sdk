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
	// The list of configurations.
	Configs []*ListCacheRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
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
	// The total count of records.
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
	// - Enables caching on the specified ports.
	//
	// - Valid values: `8880`, `2052`, `2082`, `2086`, `2095`, `2053`, `2083`, `2087`, and `2096`.
	//
	// - You can specify multiple ports, separated by commas (`,`).
	//
	// example:
	//
	// 8880,2052,2086
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	// The browser cache mode. Valid values:
	//
	// - `no_cache`: Disables browser caching.
	//
	// - `follow_origin`: Follows the origin server\\"s cache policy.
	//
	// - `override_origin`: Overrides the origin server\\"s cache policy.
	//
	// example:
	//
	// no_cache
	BrowserCacheMode *string `json:"BrowserCacheMode,omitempty" xml:"BrowserCacheMode,omitempty"`
	// The browser cache TTL, in seconds.
	//
	// example:
	//
	// 300
	BrowserCacheTtl *string `json:"BrowserCacheTtl,omitempty" xml:"BrowserCacheTtl,omitempty"`
	// Specifies the bypass cache mode. Valid values:
	//
	// - `cache_all`: Caches all requests.
	//
	// - `bypass_all`: Bypasses all requests.
	//
	// example:
	//
	// cache_all
	BypassCache *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	// The cache deception protection. This feature defends against web cache deception attacks by caching only validated content. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	CacheDeceptionArmor *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	// The cache reserve eligibility. This setting controls whether a user request bypasses the cache reserve node when it is forwarded to the origin server. Valid values:
	//
	// - `bypass_cache_reserve`: The request bypasses the cache reserve.
	//
	// - `eligible_for_cache_reserve`: The request is eligible for the cache reserve.
	//
	// example:
	//
	// bypass_cache_reserve
	CacheReserveEligibility *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	// Checks for the presence of specified cookies when generating the cache key. If a cookie exists, its name (case-insensitive) is included in the cache key. Separate multiple cookie names with spaces. Cookie names can contain the following characters:
	//
	// - Symbols: ``! # $ % & \\" 	- + - . ^ _ ` | ~``
	//
	// - Digits: `0-9`
	//
	// - Letters: lowercase English letters `a-z`
	//
	// example:
	//
	// cookiename1 cookiename2
	CheckPresenceCookie *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	// Checks for the presence of specified headers when generating the cache key. If a header exists, its name (case-insensitive) is included in the cache key. Separate multiple header names with spaces. Header names can contain the following characters:
	//
	// - Symbols: ``! # $ % & \\" 	- + - . ^ _ ` | ~``
	//
	// - Digits: `0-9`
	//
	// - Letters: lowercase English letters `a-z`
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
	// The configuration type, which indicates whether the configuration is global or rule-specific. Valid values:
	//
	// - `global`
	//
	// - `rule`
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The edge cache mode. Valid values:
	//
	// - `follow_origin`: Follows the origin server\\"s cache policy. If no policy exists, the default policy is used.
	//
	// - `no_cache`: Disables caching on edge nodes.
	//
	// - `override_origin`: Overrides the origin server\\"s cache policy.
	//
	// - `follow_origin_bypass`: Follows the origin server\\"s cache policy. If no policy exists, requests bypass the cache.
	//
	// - `follow_origin_override`: Follows the cache policy of the origin server. If no policy exists, a custom cache TTL is used.
	//
	// example:
	//
	// follow_origin
	EdgeCacheMode *string `json:"EdgeCacheMode,omitempty" xml:"EdgeCacheMode,omitempty"`
	// The edge cache TTL, in seconds.
	//
	// example:
	//
	// 300
	EdgeCacheTtl *string `json:"EdgeCacheTtl,omitempty" xml:"EdgeCacheTtl,omitempty"`
	// The status code cache TTL, in seconds.
	//
	// - You can set the cache TTL for a specific status code. For example, `404=10` caches responses with a 404 status code for 10 seconds.
	//
	// - You can set the cache TTL for a series of status codes, such as 4xx and 5xx. For example, `4xx=10` caches all responses with a 4xx status code for 10 seconds.
	//
	// - Separate multiple settings with commas (`,`).
	//
	// example:
	//
	// 5xx=0,404=10
	EdgeStatusCodeCacheTtl *string `json:"EdgeStatusCodeCacheTtl,omitempty" xml:"EdgeStatusCodeCacheTtl,omitempty"`
	// The cookie names whose values are included in the cache key. Names are case-insensitive. Separate multiple names with spaces. Cookie names can contain the following characters:
	//
	// - Symbols: ``! # $ % & \\" 	- + - . ^ _ ` | ~``
	//
	// - Digits: `0-9`
	//
	// - Letters: lowercase English letters `a-z`
	//
	// example:
	//
	// cookiename1 cookiename2
	IncludeCookie *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	// The header names whose values are included in the cache key. Names are case-insensitive. Separate multiple names with spaces. Header names can contain the following characters:
	//
	// - Symbols: ``! # $ % & \\" 	- + - . ^ _ ` | ~``
	//
	// - Digits: `0-9`
	//
	// - Letters: lowercase English letters `a-z`
	//
	// example:
	//
	// headername1 headername2
	IncludeHeader *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// The handling mode for the request body when generating the cache key for a POST request.
	//
	// - `md5`: Calculates the MD5 hash of the body content and includes the hash in the cache key.
	//
	// - `ignore`: Ignores the body content in the cache key.
	//
	// example:
	//
	// ignore
	PostBodyCacheKey *string `json:"PostBodyCacheKey,omitempty" xml:"PostBodyCacheKey,omitempty"`
	// The maximum size of a POST request body that can be cached, in KB. The value must be an integer from 1 to 8. The default is 8 KB.
	//
	// example:
	//
	// 1
	PostBodySizeLimit *string `json:"PostBodySizeLimit,omitempty" xml:"PostBodySizeLimit,omitempty"`
	// Specifies whether to enable caching for POST requests.
	//
	// example:
	//
	// on
	PostCache *string `json:"PostCache,omitempty" xml:"PostCache,omitempty"`
	// The query strings to include in or exclude from the cache key. Separate multiple values with spaces.
	//
	// example:
	//
	// example
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// Specifies how to handle query strings when generating a cache key. Valid values:
	//
	// - `ignore_all`: Ignores all query strings.
	//
	// - `exclude_query_string`: Excludes specified query strings.
	//
	// - `reserve_all`: Retains all query strings. This is the default value.
	//
	// - `include_query_string`: Includes specified query strings.
	//
	// example:
	//
	// ignore_all
	QueryStringMode *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. This parameter is not required for a global configuration.
	//
	// - To match all incoming requests, set this to `true`.
	//
	// - To match specific requests, set this to a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule status. This parameter is not required for a global configuration. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. This parameter is not required for a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution sequence. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to serve stale content. If enabled, edge nodes serve expired cached files when the origin server is unavailable. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	ServeStale *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	// The site version. If version management is enabled for the site, this specifies the version to which the configuration applies. The default is 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Specifies whether to enable query string sorting. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	// Specifies whether to include the client device type in the cache key. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	UserDeviceType *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	// Specifies whether to include the client\\"s geographical location in the cache key. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	UserGeo *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	// Specifies whether to include the client language in the cache key. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
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
