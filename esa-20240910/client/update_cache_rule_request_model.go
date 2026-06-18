// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalCacheablePorts(v string) *UpdateCacheRuleRequest
	GetAdditionalCacheablePorts() *string
	SetBrowserCacheMode(v string) *UpdateCacheRuleRequest
	GetBrowserCacheMode() *string
	SetBrowserCacheTtl(v string) *UpdateCacheRuleRequest
	GetBrowserCacheTtl() *string
	SetBypassCache(v string) *UpdateCacheRuleRequest
	GetBypassCache() *string
	SetCacheDeceptionArmor(v string) *UpdateCacheRuleRequest
	GetCacheDeceptionArmor() *string
	SetCacheReserveEligibility(v string) *UpdateCacheRuleRequest
	GetCacheReserveEligibility() *string
	SetCheckPresenceCookie(v string) *UpdateCacheRuleRequest
	GetCheckPresenceCookie() *string
	SetCheckPresenceHeader(v string) *UpdateCacheRuleRequest
	GetCheckPresenceHeader() *string
	SetConfigId(v int64) *UpdateCacheRuleRequest
	GetConfigId() *int64
	SetEdgeCacheMode(v string) *UpdateCacheRuleRequest
	GetEdgeCacheMode() *string
	SetEdgeCacheTtl(v string) *UpdateCacheRuleRequest
	GetEdgeCacheTtl() *string
	SetEdgeStatusCodeCacheTtl(v string) *UpdateCacheRuleRequest
	GetEdgeStatusCodeCacheTtl() *string
	SetIncludeCookie(v string) *UpdateCacheRuleRequest
	GetIncludeCookie() *string
	SetIncludeHeader(v string) *UpdateCacheRuleRequest
	GetIncludeHeader() *string
	SetPostBodyCacheKey(v string) *UpdateCacheRuleRequest
	GetPostBodyCacheKey() *string
	SetPostBodySizeLimit(v string) *UpdateCacheRuleRequest
	GetPostBodySizeLimit() *string
	SetPostCache(v string) *UpdateCacheRuleRequest
	GetPostCache() *string
	SetQueryString(v string) *UpdateCacheRuleRequest
	GetQueryString() *string
	SetQueryStringMode(v string) *UpdateCacheRuleRequest
	GetQueryStringMode() *string
	SetRule(v string) *UpdateCacheRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateCacheRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateCacheRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateCacheRuleRequest
	GetSequence() *int32
	SetServeStale(v string) *UpdateCacheRuleRequest
	GetServeStale() *string
	SetSiteId(v int64) *UpdateCacheRuleRequest
	GetSiteId() *int64
	SetSortQueryStringForCache(v string) *UpdateCacheRuleRequest
	GetSortQueryStringForCache() *string
	SetUserDeviceType(v string) *UpdateCacheRuleRequest
	GetUserDeviceType() *string
	SetUserGeo(v string) *UpdateCacheRuleRequest
	GetUserGeo() *string
	SetUserLanguage(v string) *UpdateCacheRuleRequest
	GetUserLanguage() *string
}

type UpdateCacheRuleRequest struct {
	// - Enables caching on the specified ports.
	//
	// - Valid values: `8880`, `2052`, `2082`, `2086`, `2095`, `2053`, `2083`, `2087`, `2096`
	//
	// - Separate multiple ports with commas.
	//
	// example:
	//
	// 8880,2052,2086
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	// The browser cache mode. Valid values:
	//
	// - `no_cache`: Does not cache content in the browser.
	//
	// - `follow_origin`: Follows the caching policy of the origin server.
	//
	// - `override_origin`: Overrides the caching policy of the origin server.
	//
	// example:
	//
	// no_cache
	BrowserCacheMode *string `json:"BrowserCacheMode,omitempty" xml:"BrowserCacheMode,omitempty"`
	// The browser cache TTL (Time to Live), in seconds.
	//
	// example:
	//
	// 300
	BrowserCacheTtl *string `json:"BrowserCacheTtl,omitempty" xml:"BrowserCacheTtl,omitempty"`
	// The cache bypass mode. Valid values:
	//
	// - `cache_all`: Caches all requests.
	//
	// - `bypass_all`: Bypasses the cache for all requests.
	//
	// example:
	//
	// cache_all
	BypassCache *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	// Defends against Web Cache Deception attacks by caching only validated content. Valid values:
	//
	// - `on`: Enables the feature.
	//
	// - `off`: Disables the feature.
	//
	// example:
	//
	// on
	CacheDeceptionArmor *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	// Controls whether requests bypass the cache reserve node during an origin-pull. Valid values:
	//
	// - `bypass_cache_reserve`: The request bypasses the cache reserve.
	//
	// - `eligible_for_cache_reserve`: The request is eligible for cache reserve.
	//
	// example:
	//
	// bypass_cache_reserve
	CacheReserveEligibility *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	// The cookies to check for. If a specified cookie is present in the request, its name (case-insensitive) is added to the cache key. Separate multiple cookies with spaces. Cookie names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Lowercase letters: a-z
	//
	// example:
	//
	// cookiename1 cookiename2
	CheckPresenceCookie *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	// The headers to check for. If a specified header is present in the request, its name (case-insensitive) is added to the cache key. Separate multiple headers with spaces. Header names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Lowercase letters: a-z
	//
	// example:
	//
	// headername1 headername2
	CheckPresenceHeader *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
	// The configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The cache mode for the edge node. Valid values:
	//
	// - `follow_origin`: Follows the origin server\\"s caching policy. If the origin server has no policy, the default policy is used.
	//
	// - `no_cache`: Does not cache content.
	//
	// - `override_origin`: Overrides the caching policy of the origin server.
	//
	// - `follow_origin_bypass`: Follows the caching policy of the origin server, if one exists. Otherwise, content is not cached.
	//
	// - `follow_origin_override`: Follows the caching policy of the origin server, if one exists. Otherwise, a custom cache TTL is used.
	//
	// example:
	//
	// follow_origin
	EdgeCacheMode *string `json:"EdgeCacheMode,omitempty" xml:"EdgeCacheMode,omitempty"`
	// The edge node cache TTL (Time to Live), in seconds.
	//
	// example:
	//
	// 300
	EdgeCacheTtl *string `json:"EdgeCacheTtl,omitempty" xml:"EdgeCacheTtl,omitempty"`
	// The cache TTL for specific status codes, in seconds.
	//
	// - You can set the cache TTL for a specific status code. For example, `404=10` caches responses with a 404 status code for 10 seconds.
	//
	// - You can set the cache TTL for `4xx` and `5xx` status code ranges. For example, `4xx=10` caches all responses with a `4xx` status code for 10 seconds.
	//
	// - Separate multiple status code settings with commas.
	//
	// example:
	//
	// 5xx=0,404=10
	EdgeStatusCodeCacheTtl *string `json:"EdgeStatusCodeCacheTtl,omitempty" xml:"EdgeStatusCodeCacheTtl,omitempty"`
	// The cookies to include in the cache key. Both the cookie names (case-insensitive) and their values are used. Separate multiple cookies with spaces. Cookie names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Lowercase letters: a-z
	//
	// example:
	//
	// cookiename1 cookiename2
	IncludeCookie *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	// The headers to include in the cache key. Both the header names (case-insensitive) and their values are used. Separate multiple headers with spaces. Header names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Lowercase letters: a-z
	//
	// example:
	//
	// headername1 headername2
	IncludeHeader *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// Controls how the request body is used to generate the cache key for POST requests. Valid values:
	//
	// - `md5`: Calculates the MD5 hash of the request body and includes the hash in the cache key.
	//
	// - `ignore`: Ignores the request body when generating the cache key.
	//
	// example:
	//
	// ignore
	PostBodyCacheKey *string `json:"PostBodyCacheKey,omitempty" xml:"PostBodyCacheKey,omitempty"`
	// The maximum size of a request body for POST caching, in KB. The value must be an integer from 1 to 8. If you leave this parameter empty, the default value of 8 KB is used.
	//
	// example:
	//
	// 1
	PostBodySizeLimit *string `json:"PostBodySizeLimit,omitempty" xml:"PostBodySizeLimit,omitempty"`
	// Controls whether to cache responses to POST requests.
	//
	// example:
	//
	// on
	PostCache *string `json:"PostCache,omitempty" xml:"PostCache,omitempty"`
	// The query string parameters to include in or exclude from the cache key. Separate multiple parameters with spaces.
	//
	// example:
	//
	// example1 example2
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// Controls how query strings are used to generate a cache key. Valid values:
	//
	// - `ignore_all`: Ignores all query strings.
	//
	// - `exclude_query_string`: Removes specified query strings.
	//
	// - `reserve_all`: Retains all query strings. This is the default value.
	//
	// - `include_query_string`: Retains only specified query strings.
	//
	// example:
	//
	// ignore_all
	QueryStringMode *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	// A conditional expression that matches user requests. This parameter is optional for a global configuration. Two scenarios are supported:
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, set the value to a custom expression, for example, `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Controls whether the rule is enabled. This parameter is optional for a global configuration. Valid values:
	//
	// - `on`: Enables the rule.
	//
	// - `off`: Disables the rule.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. This parameter is optional for a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution priority of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Controls whether to serve stale content. If enabled, an edge node can serve expired content from its cache if the origin server is unavailable. Valid values:
	//
	// - `on`: Enables this feature.
	//
	// - `off`: Disables this feature.
	//
	// example:
	//
	// on
	ServeStale *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	// The ID of the site. To get this ID, call the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Controls whether to sort query string parameters when generating a cache key. Valid values:
	//
	// - `on`: Enables sorting.
	//
	// - `off`: Disables sorting.
	//
	// example:
	//
	// on
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	// Controls whether to include the client device type in the cache key. Valid values:
	//
	// - `on`: Enables this feature.
	//
	// - `off`: Disables this feature.
	//
	// example:
	//
	// on
	UserDeviceType *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	// Controls whether to include the client\\"s geographic location in the cache key. Valid values:
	//
	// - `on`: Enables this feature.
	//
	// - `off`: Disables this feature.
	//
	// example:
	//
	// on
	UserGeo *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	// Controls whether to include the client\\"s language in the cache key. Valid values:
	//
	// - `on`: Enables this feature.
	//
	// - `off`: Disables this feature.
	//
	// example:
	//
	// on
	UserLanguage *string `json:"UserLanguage,omitempty" xml:"UserLanguage,omitempty"`
}

func (s UpdateCacheRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateCacheRuleRequest) GetAdditionalCacheablePorts() *string {
	return s.AdditionalCacheablePorts
}

func (s *UpdateCacheRuleRequest) GetBrowserCacheMode() *string {
	return s.BrowserCacheMode
}

func (s *UpdateCacheRuleRequest) GetBrowserCacheTtl() *string {
	return s.BrowserCacheTtl
}

func (s *UpdateCacheRuleRequest) GetBypassCache() *string {
	return s.BypassCache
}

func (s *UpdateCacheRuleRequest) GetCacheDeceptionArmor() *string {
	return s.CacheDeceptionArmor
}

func (s *UpdateCacheRuleRequest) GetCacheReserveEligibility() *string {
	return s.CacheReserveEligibility
}

func (s *UpdateCacheRuleRequest) GetCheckPresenceCookie() *string {
	return s.CheckPresenceCookie
}

func (s *UpdateCacheRuleRequest) GetCheckPresenceHeader() *string {
	return s.CheckPresenceHeader
}

func (s *UpdateCacheRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateCacheRuleRequest) GetEdgeCacheMode() *string {
	return s.EdgeCacheMode
}

func (s *UpdateCacheRuleRequest) GetEdgeCacheTtl() *string {
	return s.EdgeCacheTtl
}

func (s *UpdateCacheRuleRequest) GetEdgeStatusCodeCacheTtl() *string {
	return s.EdgeStatusCodeCacheTtl
}

func (s *UpdateCacheRuleRequest) GetIncludeCookie() *string {
	return s.IncludeCookie
}

func (s *UpdateCacheRuleRequest) GetIncludeHeader() *string {
	return s.IncludeHeader
}

func (s *UpdateCacheRuleRequest) GetPostBodyCacheKey() *string {
	return s.PostBodyCacheKey
}

func (s *UpdateCacheRuleRequest) GetPostBodySizeLimit() *string {
	return s.PostBodySizeLimit
}

func (s *UpdateCacheRuleRequest) GetPostCache() *string {
	return s.PostCache
}

func (s *UpdateCacheRuleRequest) GetQueryString() *string {
	return s.QueryString
}

func (s *UpdateCacheRuleRequest) GetQueryStringMode() *string {
	return s.QueryStringMode
}

func (s *UpdateCacheRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateCacheRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateCacheRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateCacheRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateCacheRuleRequest) GetServeStale() *string {
	return s.ServeStale
}

func (s *UpdateCacheRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateCacheRuleRequest) GetSortQueryStringForCache() *string {
	return s.SortQueryStringForCache
}

func (s *UpdateCacheRuleRequest) GetUserDeviceType() *string {
	return s.UserDeviceType
}

func (s *UpdateCacheRuleRequest) GetUserGeo() *string {
	return s.UserGeo
}

func (s *UpdateCacheRuleRequest) GetUserLanguage() *string {
	return s.UserLanguage
}

func (s *UpdateCacheRuleRequest) SetAdditionalCacheablePorts(v string) *UpdateCacheRuleRequest {
	s.AdditionalCacheablePorts = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetBrowserCacheMode(v string) *UpdateCacheRuleRequest {
	s.BrowserCacheMode = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetBrowserCacheTtl(v string) *UpdateCacheRuleRequest {
	s.BrowserCacheTtl = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetBypassCache(v string) *UpdateCacheRuleRequest {
	s.BypassCache = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetCacheDeceptionArmor(v string) *UpdateCacheRuleRequest {
	s.CacheDeceptionArmor = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetCacheReserveEligibility(v string) *UpdateCacheRuleRequest {
	s.CacheReserveEligibility = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetCheckPresenceCookie(v string) *UpdateCacheRuleRequest {
	s.CheckPresenceCookie = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetCheckPresenceHeader(v string) *UpdateCacheRuleRequest {
	s.CheckPresenceHeader = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetConfigId(v int64) *UpdateCacheRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetEdgeCacheMode(v string) *UpdateCacheRuleRequest {
	s.EdgeCacheMode = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetEdgeCacheTtl(v string) *UpdateCacheRuleRequest {
	s.EdgeCacheTtl = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetEdgeStatusCodeCacheTtl(v string) *UpdateCacheRuleRequest {
	s.EdgeStatusCodeCacheTtl = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetIncludeCookie(v string) *UpdateCacheRuleRequest {
	s.IncludeCookie = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetIncludeHeader(v string) *UpdateCacheRuleRequest {
	s.IncludeHeader = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetPostBodyCacheKey(v string) *UpdateCacheRuleRequest {
	s.PostBodyCacheKey = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetPostBodySizeLimit(v string) *UpdateCacheRuleRequest {
	s.PostBodySizeLimit = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetPostCache(v string) *UpdateCacheRuleRequest {
	s.PostCache = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetQueryString(v string) *UpdateCacheRuleRequest {
	s.QueryString = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetQueryStringMode(v string) *UpdateCacheRuleRequest {
	s.QueryStringMode = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetRule(v string) *UpdateCacheRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetRuleEnable(v string) *UpdateCacheRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetRuleName(v string) *UpdateCacheRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetSequence(v int32) *UpdateCacheRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetServeStale(v string) *UpdateCacheRuleRequest {
	s.ServeStale = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetSiteId(v int64) *UpdateCacheRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetSortQueryStringForCache(v string) *UpdateCacheRuleRequest {
	s.SortQueryStringForCache = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetUserDeviceType(v string) *UpdateCacheRuleRequest {
	s.UserDeviceType = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetUserGeo(v string) *UpdateCacheRuleRequest {
	s.UserGeo = &v
	return s
}

func (s *UpdateCacheRuleRequest) SetUserLanguage(v string) *UpdateCacheRuleRequest {
	s.UserLanguage = &v
	return s
}

func (s *UpdateCacheRuleRequest) Validate() error {
	return dara.Validate(s)
}
