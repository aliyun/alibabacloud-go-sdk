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
	// - Additional ports on which caching is enabled.
	//
	// - Valid values: `8880`, `2052`, `2082`, `2086`, `2095`, `2053`, `2083`, `2087`, `2096`.
	//
	// - Separate multiple ports with commas.
	//
	// example:
	//
	// 8880,2052,2086
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	// The browser cache mode. Valid values:
	//
	// - `no_cache`: Does not cache content.
	//
	// - `follow_origin`: Follows the origin cache policy.
	//
	// - `override_origin`: Overrides the origin cache policy.
	//
	// example:
	//
	// follow_origin
	BrowserCacheMode *string `json:"BrowserCacheMode,omitempty" xml:"BrowserCacheMode,omitempty"`
	// The browser cache TTL, in seconds.
	//
	// example:
	//
	// 300
	BrowserCacheTtl *string `json:"BrowserCacheTtl,omitempty" xml:"BrowserCacheTtl,omitempty"`
	// Specifies whether to cache requests or bypass the cache. Valid values:
	//
	// - `cache_all`: Caches all requests.
	//
	// - `bypass_all`: Bypasses the cache for all requests.
	//
	// example:
	//
	// cache_all
	BypassCache *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	// Specifies whether to enable Cache Deception Armor. This feature helps mitigate web cache deception attacks by ensuring that only validated content is cached. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	CacheDeceptionArmor *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	// The eligibility for cache reserve. This controls whether a request bypasses the cache reserve node during an origin fetch. Valid values:
	//
	// - `bypass_cache_reserve`: The request bypasses the cache reserve.
	//
	// - `eligible_for_cache_reserve`: The request is eligible for cache reserve.
	//
	// example:
	//
	// bypass_cache_reserve
	CacheReserveEligibility *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	// Specifies cookies whose presence is checked when generating a cache key. If a specified cookie exists in the request, its name (case-insensitive) is added to the cache key. Separate multiple cookie names with spaces. Cookie names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Letters: lowercase English letters a-z
	//
	// example:
	//
	// cookiename1 cookiename2
	CheckPresenceCookie *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	// Specifies headers whose presence is checked when generating a cache key. If a specified header exists in the request, its name (case-insensitive) is added to the cache key. Separate multiple header names with spaces. Header names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Letters: lowercase English letters a-z
	//
	// example:
	//
	// headername1 headername2
	CheckPresenceHeader *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Indicates whether the response contains a global or a rule configuration. Valid values:
	//
	// - `global`: A global configuration.
	//
	// - `rule`: A rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The edge cache mode. Valid values:
	//
	// - `follow_origin`: Uses the origin server\\"s cache policy. If none is provided, the default policy applies.
	//
	// - `no_cache`: Does not cache content.
	//
	// - `override_origin`: Overrides the origin cache policy.
	//
	// - `follow_origin_bypass`: Uses the origin server\\"s cache policy. If none is provided, content is not cached.
	//
	// - `follow_origin_override`: Uses the origin server\\"s cache policy. If none is provided, a custom cache TTL applies.
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
	// - Set the cache TTL for a specific status code. For example, `404=10` specifies that responses with a 404 status code are cached for 10 seconds.
	//
	// - Set the cache TTL for status code classes, such as 4xx and 5xx. For example, `4xx=10` specifies that all responses with a 4xx status code are cached for 10 seconds.
	//
	// - Separate multiple entries with commas.
	//
	// example:
	//
	// 5xx=0,404=10
	EdgeStatusCodeCacheTtl *string `json:"EdgeStatusCodeCacheTtl,omitempty" xml:"EdgeStatusCodeCacheTtl,omitempty"`
	// Specifies the cookies to include in the cache key. Both the cookie names (case-insensitive) and their values are added to the key. Separate multiple cookie names with spaces. Cookie names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Letters: lowercase English letters a-z
	//
	// example:
	//
	// cookiename1 cookiename2
	IncludeCookie *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	// Specifies the headers to include in the cache key. Both the header names (case-insensitive) and their values are added to the key. Separate multiple header names with spaces. Header names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Letters: lowercase English letters a-z
	//
	// example:
	//
	// headername1 headername2
	IncludeHeader *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// The mode for handling the body content when generating a cache key for POST requests. Valid values:
	//
	// - `md5`: Calculates the MD5 hash of the body content and adds the hash to the cache key.
	//
	// - `ignore`: Ignores the body content in the cache key.
	//
	// example:
	//
	// ignore
	PostBodyCacheKey *string `json:"PostBodyCacheKey,omitempty" xml:"PostBodyCacheKey,omitempty"`
	// The size limit (in KB) of the body content for POST caching. The value is an integer from 1 to 8. A null or empty value defaults to 8 KB.
	//
	// example:
	//
	// 1
	PostBodySizeLimit *string `json:"PostBodySizeLimit,omitempty" xml:"PostBodySizeLimit,omitempty"`
	// Specifies whether to enable the POST cache feature.
	//
	// example:
	//
	// on
	PostCache *string `json:"PostCache,omitempty" xml:"PostCache,omitempty"`
	// Specifies the query strings to include in or exclude from the cache key. Separate multiple query strings with spaces.
	//
	// example:
	//
	// example
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// The mode for handling query strings when generating a cache key. Valid values:
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
	// reserve_all
	QueryStringMode *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule content, which is a conditional expression used to match user requests. This parameter applies only to rule configurations.
	//
	// - To match all incoming requests, use `true`.
	//
	// - To match specific requests, use a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. This parameter applies only to rule configurations. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. This parameter applies only to rule configurations.
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
	// Specifies whether to serve stale content. If enabled, edge nodes serve stale (expired) content from the cache when the origin server is unavailable. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	ServeStale *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, this indicates the configuration version. The default is 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Specifies whether to sort query strings before generating the cache key. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	// Specifies whether to include the client\\"s device type in the cache key. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	UserDeviceType *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	// Specifies whether to include the client\\"s geographic location in the cache key. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	UserGeo *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	// Specifies whether to include the client\\"s language in the cache key. Valid values:
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
