// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCacheRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalCacheablePorts(v string) *CreateCacheRuleRequest
	GetAdditionalCacheablePorts() *string
	SetBrowserCacheMode(v string) *CreateCacheRuleRequest
	GetBrowserCacheMode() *string
	SetBrowserCacheTtl(v string) *CreateCacheRuleRequest
	GetBrowserCacheTtl() *string
	SetBypassCache(v string) *CreateCacheRuleRequest
	GetBypassCache() *string
	SetCacheDeceptionArmor(v string) *CreateCacheRuleRequest
	GetCacheDeceptionArmor() *string
	SetCacheReserveEligibility(v string) *CreateCacheRuleRequest
	GetCacheReserveEligibility() *string
	SetCheckPresenceCookie(v string) *CreateCacheRuleRequest
	GetCheckPresenceCookie() *string
	SetCheckPresenceHeader(v string) *CreateCacheRuleRequest
	GetCheckPresenceHeader() *string
	SetEdgeCacheMode(v string) *CreateCacheRuleRequest
	GetEdgeCacheMode() *string
	SetEdgeCacheTtl(v string) *CreateCacheRuleRequest
	GetEdgeCacheTtl() *string
	SetEdgeStatusCodeCacheTtl(v string) *CreateCacheRuleRequest
	GetEdgeStatusCodeCacheTtl() *string
	SetIncludeCookie(v string) *CreateCacheRuleRequest
	GetIncludeCookie() *string
	SetIncludeHeader(v string) *CreateCacheRuleRequest
	GetIncludeHeader() *string
	SetPostBodyCacheKey(v string) *CreateCacheRuleRequest
	GetPostBodyCacheKey() *string
	SetPostBodySizeLimit(v string) *CreateCacheRuleRequest
	GetPostBodySizeLimit() *string
	SetPostCache(v string) *CreateCacheRuleRequest
	GetPostCache() *string
	SetQueryString(v string) *CreateCacheRuleRequest
	GetQueryString() *string
	SetQueryStringMode(v string) *CreateCacheRuleRequest
	GetQueryStringMode() *string
	SetRule(v string) *CreateCacheRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateCacheRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateCacheRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateCacheRuleRequest
	GetSequence() *int32
	SetServeStale(v string) *CreateCacheRuleRequest
	GetServeStale() *string
	SetSiteId(v int64) *CreateCacheRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateCacheRuleRequest
	GetSiteVersion() *int32
	SetSortQueryStringForCache(v string) *CreateCacheRuleRequest
	GetSortQueryStringForCache() *string
	SetUserDeviceType(v string) *CreateCacheRuleRequest
	GetUserDeviceType() *string
	SetUserGeo(v string) *CreateCacheRuleRequest
	GetUserGeo() *string
	SetUserLanguage(v string) *CreateCacheRuleRequest
	GetUserLanguage() *string
}

type CreateCacheRuleRequest struct {
	// - Specifies additional ports on which caching is enabled.
	//
	// - Valid values: 8880, 2052, 2082, 2086, 2095, 2053, 2083, 2087, and 2096.
	//
	// - You can specify multiple ports, separated by commas (,).
	//
	// example:
	//
	// 8880,2052,2086
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	// The browser cache mode. Valid values:
	//
	// - `no_cache`: Disables browser caching.
	//
	// - `follow_origin`: Follows the origin server cache policy.
	//
	// - `override_origin`: Overrides the origin server cache policy.
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
	// The bypass cache mode. Valid values:
	//
	// - `cache_all`: Caches all requests.
	//
	// - `bypass_all`: Bypasses the cache for all requests.
	//
	// example:
	//
	// cache_all
	BypassCache *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	// Specifies whether to enable cache deception defense. This feature helps defend against web cache deception attacks. When this feature is enabled, only content that passes validation is cached. Valid values:
	//
	// - `on`: Enables the defense.
	//
	// - `off`: Disables the defense.
	//
	// example:
	//
	// on
	CacheDeceptionArmor *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	// Specifies whether requests bypass cache reservation nodes during an origin fetch. Valid values:
	//
	// - `bypass_cache_reserve`: The request bypasses cache reservation.
	//
	// - `eligible_for_cache_reserve`: The request is eligible for cache reservation.
	//
	// example:
	//
	// bypass_cache_reserve
	CacheReserveEligibility *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	// Specifies the cookies to check for presence when generating a cache key. If a specified cookie is present in the request, its name (case-insensitive) is included in the cache key. To specify multiple cookies, separate their names with spaces. The cookie names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Letters: a-z (lowercase)
	//
	// example:
	//
	// cookiename1 cookiename2
	CheckPresenceCookie *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	// Specifies the headers to check for presence when generating a cache key. If a specified header is present in the request, its name (case-insensitive) is included in the cache key. To specify multiple headers, separate their names with spaces. The header names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Letters: a-z (lowercase)
	//
	// example:
	//
	// headername1 headername2
	CheckPresenceHeader *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
	// The edge cache mode. Valid values:
	//
	// - `follow_origin`: Follows the origin server cache policy if one exists; otherwise, uses the default cache policy.
	//
	// - `no_cache`: Disables caching on the edge node.
	//
	// - `override_origin`: Overrides the origin server cache policy.
	//
	// - `follow_origin_bypass`: Follows the origin server cache policy if one exists; otherwise, the content is not cached.
	//
	// - `follow_origin_override`: Follows the origin server cache policy if one exists; otherwise, uses a custom edge cache TTL.
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
	// - You can set the cache TTL for a series of status codes, such as 4xx or 5xx. For example, `4xx=10` caches all responses that have a status code in the 4xx series for 10 seconds.
	//
	// - You can specify multiple status code TTLs, separated by commas (,).
	//
	// example:
	//
	// 5xx=0,404=10
	EdgeStatusCodeCacheTtl *string `json:"EdgeStatusCodeCacheTtl,omitempty" xml:"EdgeStatusCodeCacheTtl,omitempty"`
	// The cookies to include in the cache key. Both the cookie names (case-insensitive) and their values are included. Separate multiple cookie names with spaces. The cookie names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Letters: a-z (lowercase)
	//
	// example:
	//
	// cookiename1 cookiename2
	IncludeCookie *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	// The headers to include in the cache key. Both the header names (case-insensitive) and their values are included. Separate multiple header names with spaces. The header names can contain the following characters:
	//
	// - Symbols: ! # $ % & \\" \\	- + - . ^ _ | \\~
	//
	// - Digits: 0-9
	//
	// - Letters: a-z (lowercase)
	//
	// example:
	//
	// headername1 headername2
	IncludeHeader *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// Specifies how to process the request body when generating a cache key for POST requests. The following modes are supported:
	//
	// - `md5`: Calculates the MD5 hash of the request body and adds the hash value to the cache key.
	//
	// - `ignore`: Ignores the request body when the cache key is generated.
	//
	// example:
	//
	// ignore
	PostBodyCacheKey *string `json:"PostBodyCacheKey,omitempty" xml:"PostBodyCacheKey,omitempty"`
	// The size limit for the request body when using POST request caching, in KB. Supported values range from 1 to 8. If unspecified, the default is 8 KB.
	//
	// example:
	//
	// 1
	PostBodySizeLimit *string `json:"PostBodySizeLimit,omitempty" xml:"PostBodySizeLimit,omitempty"`
	// Specifies whether to enable POST request caching.
	//
	// example:
	//
	// on
	PostCache *string `json:"PostCache,omitempty" xml:"PostCache,omitempty"`
	// The query strings to include in or exclude from the cache key. Separate multiple query strings with spaces.
	//
	// example:
	//
	// example
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// The mode for processing query strings when generating a cache key. Valid values:
	//
	// - `ignore_all`: Ignores all query strings.
	//
	// - `exclude_query_string`: Excludes specified query strings.
	//
	// - `reserve_all`: Includes all query strings (the default).
	//
	// - `include_query_string`: Includes only specified query strings.
	//
	// example:
	//
	// reserve_all
	QueryStringMode *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	// The content of the rule, which is a conditional expression used to match user requests. This parameter is not required for a global configuration.
	//
	// - To match all requests, set the value to `true`.
	//
	// - To match specific requests, set the value to a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. This parameter is not required for a global configuration. Valid values:
	//
	// - `on`: Enables the rule.
	//
	// - `off`: Disables the rule.
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
	// The execution order of the rule. A smaller number indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to serve stale content. If enabled, an edge node can serve stale (expired) content when the origin server is unavailable. Valid values:
	//
	// - `on`: Enables serving stale content.
	//
	// - `off`: Disables serving stale content.
	//
	// example:
	//
	// on
	ServeStale *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 340035003106221
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site configuration version. For sites with version management enabled, this parameter specifies the site version to which the configuration applies.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Specifies whether to sort query strings. This feature is disabled by default. Valid values:
	//
	// - `on`: Enables sorting.
	//
	// - `off`: Disables sorting.
	//
	// example:
	//
	// on
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	// Specifies whether to include the client device type in the cache key. Valid values:
	//
	// - `on`: Includes the client device type.
	//
	// - `off`: Does not include the client device type.
	//
	// example:
	//
	// on
	UserDeviceType *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	// Specifies whether to include the client\\"s geographic location in the cache key. Valid values:
	//
	// - `on`: Includes the geographic location.
	//
	// - `off`: Does not include the geographic location.
	//
	// example:
	//
	// on
	UserGeo *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	// Specifies whether to include the client\\"s language in the cache key. Valid values:
	//
	// - `on`: Includes the language.
	//
	// - `off`: Does not include the language.
	//
	// example:
	//
	// on
	UserLanguage *string `json:"UserLanguage,omitempty" xml:"UserLanguage,omitempty"`
}

func (s CreateCacheRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateCacheRuleRequest) GetAdditionalCacheablePorts() *string {
	return s.AdditionalCacheablePorts
}

func (s *CreateCacheRuleRequest) GetBrowserCacheMode() *string {
	return s.BrowserCacheMode
}

func (s *CreateCacheRuleRequest) GetBrowserCacheTtl() *string {
	return s.BrowserCacheTtl
}

func (s *CreateCacheRuleRequest) GetBypassCache() *string {
	return s.BypassCache
}

func (s *CreateCacheRuleRequest) GetCacheDeceptionArmor() *string {
	return s.CacheDeceptionArmor
}

func (s *CreateCacheRuleRequest) GetCacheReserveEligibility() *string {
	return s.CacheReserveEligibility
}

func (s *CreateCacheRuleRequest) GetCheckPresenceCookie() *string {
	return s.CheckPresenceCookie
}

func (s *CreateCacheRuleRequest) GetCheckPresenceHeader() *string {
	return s.CheckPresenceHeader
}

func (s *CreateCacheRuleRequest) GetEdgeCacheMode() *string {
	return s.EdgeCacheMode
}

func (s *CreateCacheRuleRequest) GetEdgeCacheTtl() *string {
	return s.EdgeCacheTtl
}

func (s *CreateCacheRuleRequest) GetEdgeStatusCodeCacheTtl() *string {
	return s.EdgeStatusCodeCacheTtl
}

func (s *CreateCacheRuleRequest) GetIncludeCookie() *string {
	return s.IncludeCookie
}

func (s *CreateCacheRuleRequest) GetIncludeHeader() *string {
	return s.IncludeHeader
}

func (s *CreateCacheRuleRequest) GetPostBodyCacheKey() *string {
	return s.PostBodyCacheKey
}

func (s *CreateCacheRuleRequest) GetPostBodySizeLimit() *string {
	return s.PostBodySizeLimit
}

func (s *CreateCacheRuleRequest) GetPostCache() *string {
	return s.PostCache
}

func (s *CreateCacheRuleRequest) GetQueryString() *string {
	return s.QueryString
}

func (s *CreateCacheRuleRequest) GetQueryStringMode() *string {
	return s.QueryStringMode
}

func (s *CreateCacheRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateCacheRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateCacheRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateCacheRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateCacheRuleRequest) GetServeStale() *string {
	return s.ServeStale
}

func (s *CreateCacheRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateCacheRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateCacheRuleRequest) GetSortQueryStringForCache() *string {
	return s.SortQueryStringForCache
}

func (s *CreateCacheRuleRequest) GetUserDeviceType() *string {
	return s.UserDeviceType
}

func (s *CreateCacheRuleRequest) GetUserGeo() *string {
	return s.UserGeo
}

func (s *CreateCacheRuleRequest) GetUserLanguage() *string {
	return s.UserLanguage
}

func (s *CreateCacheRuleRequest) SetAdditionalCacheablePorts(v string) *CreateCacheRuleRequest {
	s.AdditionalCacheablePorts = &v
	return s
}

func (s *CreateCacheRuleRequest) SetBrowserCacheMode(v string) *CreateCacheRuleRequest {
	s.BrowserCacheMode = &v
	return s
}

func (s *CreateCacheRuleRequest) SetBrowserCacheTtl(v string) *CreateCacheRuleRequest {
	s.BrowserCacheTtl = &v
	return s
}

func (s *CreateCacheRuleRequest) SetBypassCache(v string) *CreateCacheRuleRequest {
	s.BypassCache = &v
	return s
}

func (s *CreateCacheRuleRequest) SetCacheDeceptionArmor(v string) *CreateCacheRuleRequest {
	s.CacheDeceptionArmor = &v
	return s
}

func (s *CreateCacheRuleRequest) SetCacheReserveEligibility(v string) *CreateCacheRuleRequest {
	s.CacheReserveEligibility = &v
	return s
}

func (s *CreateCacheRuleRequest) SetCheckPresenceCookie(v string) *CreateCacheRuleRequest {
	s.CheckPresenceCookie = &v
	return s
}

func (s *CreateCacheRuleRequest) SetCheckPresenceHeader(v string) *CreateCacheRuleRequest {
	s.CheckPresenceHeader = &v
	return s
}

func (s *CreateCacheRuleRequest) SetEdgeCacheMode(v string) *CreateCacheRuleRequest {
	s.EdgeCacheMode = &v
	return s
}

func (s *CreateCacheRuleRequest) SetEdgeCacheTtl(v string) *CreateCacheRuleRequest {
	s.EdgeCacheTtl = &v
	return s
}

func (s *CreateCacheRuleRequest) SetEdgeStatusCodeCacheTtl(v string) *CreateCacheRuleRequest {
	s.EdgeStatusCodeCacheTtl = &v
	return s
}

func (s *CreateCacheRuleRequest) SetIncludeCookie(v string) *CreateCacheRuleRequest {
	s.IncludeCookie = &v
	return s
}

func (s *CreateCacheRuleRequest) SetIncludeHeader(v string) *CreateCacheRuleRequest {
	s.IncludeHeader = &v
	return s
}

func (s *CreateCacheRuleRequest) SetPostBodyCacheKey(v string) *CreateCacheRuleRequest {
	s.PostBodyCacheKey = &v
	return s
}

func (s *CreateCacheRuleRequest) SetPostBodySizeLimit(v string) *CreateCacheRuleRequest {
	s.PostBodySizeLimit = &v
	return s
}

func (s *CreateCacheRuleRequest) SetPostCache(v string) *CreateCacheRuleRequest {
	s.PostCache = &v
	return s
}

func (s *CreateCacheRuleRequest) SetQueryString(v string) *CreateCacheRuleRequest {
	s.QueryString = &v
	return s
}

func (s *CreateCacheRuleRequest) SetQueryStringMode(v string) *CreateCacheRuleRequest {
	s.QueryStringMode = &v
	return s
}

func (s *CreateCacheRuleRequest) SetRule(v string) *CreateCacheRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateCacheRuleRequest) SetRuleEnable(v string) *CreateCacheRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateCacheRuleRequest) SetRuleName(v string) *CreateCacheRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateCacheRuleRequest) SetSequence(v int32) *CreateCacheRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateCacheRuleRequest) SetServeStale(v string) *CreateCacheRuleRequest {
	s.ServeStale = &v
	return s
}

func (s *CreateCacheRuleRequest) SetSiteId(v int64) *CreateCacheRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateCacheRuleRequest) SetSiteVersion(v int32) *CreateCacheRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateCacheRuleRequest) SetSortQueryStringForCache(v string) *CreateCacheRuleRequest {
	s.SortQueryStringForCache = &v
	return s
}

func (s *CreateCacheRuleRequest) SetUserDeviceType(v string) *CreateCacheRuleRequest {
	s.UserDeviceType = &v
	return s
}

func (s *CreateCacheRuleRequest) SetUserGeo(v string) *CreateCacheRuleRequest {
	s.UserGeo = &v
	return s
}

func (s *CreateCacheRuleRequest) SetUserLanguage(v string) *CreateCacheRuleRequest {
	s.UserLanguage = &v
	return s
}

func (s *CreateCacheRuleRequest) Validate() error {
	return dara.Validate(s)
}
