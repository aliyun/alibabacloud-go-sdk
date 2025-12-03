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
	// Enable caching on the specified ports. Value range: 8880, 2052, 2082, 2086, 2095, 2053, 2083, 2087, 2096.
	//
	// example:
	//
	// 2095
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	// Browser cache mode. Value range:
	//
	// - no_cache: Do not cache.
	//
	// - follow_origin: Follow origin cache policy.
	//
	// - override_origin: Override origin cache policy.
	//
	// example:
	//
	// follow_origin
	BrowserCacheMode *string `json:"BrowserCacheMode,omitempty" xml:"BrowserCacheMode,omitempty"`
	// Browser cache expiration time, in seconds.
	//
	// example:
	//
	// 300
	BrowserCacheTtl *string `json:"BrowserCacheTtl,omitempty" xml:"BrowserCacheTtl,omitempty"`
	// Set bypass cache mode. Value range:
	//
	// - cache_all: Cache all requests.
	//
	// - bypass_all: Bypass cache for all requests.
	//
	// example:
	//
	// cache_all
	BypassCache *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	// Cache deception defense. Used to defend against web cache deception attacks. Only the verified cache content will be cached. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	CacheDeceptionArmor *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	// Cache reserve eligibility. Used to control whether user requests bypass the cache reserve node when returning to the origin. Value range:
	//
	// - bypass_cache_reserve: Requests bypass the cache reserve.
	//
	// - eligible_for_cache_reserve: Eligible for cache reserve.
	//
	// example:
	//
	// bypass_cache_reserve
	CacheReserveEligibility *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	// When generating the cache key, check if the cookie exists. If it does, add the cookie name (cookie names are case-insensitive) to the cache key. Supports multiple cookie names, separated by spaces.
	//
	// example:
	//
	// cookiename
	CheckPresenceCookie *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	// When generating the cache key, check if the header exists. If it does, add the header name (header names are case-insensitive) to the cache key. Supports multiple header names, separated by spaces.
	//
	// example:
	//
	// headername
	CheckPresenceHeader *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
	// Configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule configurations. Value range:
	//
	// - global: Query global configuration;
	//
	// - rule: Query rule configuration;
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Edge cache mode. Value range:
	//
	// - follow_origin: Follow origin cache policy (if exists), otherwise use the default cache policy.
	//
	// - no_cache: Do not cache.
	//
	// - override_origin: Override origin cache policy.
	//
	// - follow_origin_bypass: Follow origin cache policy (if exists), otherwise do not cache.
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
	// Status code cache expiration time, in seconds.
	//
	// example:
	//
	// 300
	EdgeStatusCodeCacheTtl *string `json:"EdgeStatusCodeCacheTtl,omitempty" xml:"EdgeStatusCodeCacheTtl,omitempty"`
	// When generating the cache key, include the specified cookie names and their values. Supports multiple values, separated by spaces.
	//
	// example:
	//
	// cookie_exapmle
	IncludeCookie *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	// When generating the cache key, include the specified header names and their values. Supports multiple values, separated by spaces.
	//
	// example:
	//
	// example
	IncludeHeader *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// example:
	//
	// ignore
	PostBodyCacheKey  *string `json:"PostBodyCacheKey,omitempty" xml:"PostBodyCacheKey,omitempty"`
	PostBodySizeLimit *string `json:"PostBodySizeLimit,omitempty" xml:"PostBodySizeLimit,omitempty"`
	// example:
	//
	// on
	PostCache *string `json:"PostCache,omitempty" xml:"PostCache,omitempty"`
	// The query strings to be retained or deleted, supporting multiple values separated by spaces.
	//
	// example:
	//
	// example
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// The processing mode for query strings when generating cache keys. Value range:
	//
	// - ignore_all: Ignore all.
	//
	// - exclude_query_string: Exclude specified query strings.
	//
	// - reserve_all: Default, reserve all.
	//
	// - include_query_string: Include specified query strings.
	//
	// example:
	//
	// reserve_all
	QueryStringMode *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter does not need to be set when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter does not need to be set when adding a global configuration. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter does not need to be set when adding a global configuration.
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
	// Serve stale cache. When enabled, the node can still use the cached expired files to respond to user requests even if the origin server is unavailable. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	ServeStale *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	// Version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the configuration, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Query string sorting. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	// When generating the cache key, include the client device type. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	UserDeviceType *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	// When generating the cache key, include the client\\"s geographic location. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	UserGeo *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	// When generating the cache key, include the client\\"s language type. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
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
