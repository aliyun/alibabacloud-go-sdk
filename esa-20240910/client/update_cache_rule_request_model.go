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
	// Enable caching on specified ports. Value range: 8880, 2052, 2082, 2086, 2095, 2053, 2083, 2087, 2096.
	//
	// example:
	//
	// 8880
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
	// no_cache
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
	// Cache deception defense. Used to defend against web cache deception attacks; only the cache content that passes the validation will be cached. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	CacheDeceptionArmor *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	// Cache retention eligibility. Used to control whether user requests bypass the cache retention node when returning to the origin. Value range:
	//
	// - bypass_cache_reserve: Requests bypass cache retention.
	//
	// - eligible_for_cache_reserve: Eligible for cache retention.
	//
	// example:
	//
	// bypass_cache_reserve
	CacheReserveEligibility *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	// Check if the cookie exists when generating cache keys, and if it does, add the cookie name (case-insensitive) to the cache key. Supports multiple cookie names, separated by spaces.
	//
	// example:
	//
	// cookiename
	CheckPresenceCookie *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	// Check if the header exists when generating cache keys, and if it does, add the header name (case-insensitive) to the cache key. Supports multiple header names, separated by spaces.
	//
	// example:
	//
	// headername
	CheckPresenceHeader *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
	// Configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
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
	// Include the specified cookie names and their values when generating cache keys, supporting multiple values separated by spaces.
	//
	// example:
	//
	// cookiename
	IncludeCookie *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	// Include the specified header names and their values when generating cache keys, supporting multiple values separated by spaces.
	//
	// example:
	//
	// headername
	IncludeHeader *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// Query strings to be retained or excluded, supporting multiple values separated by spaces.
	//
	// example:
	//
	// example
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// The processing mode of query strings when generating cache keys. Values:
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
	// ignore_all
	QueryStringMode *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
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
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Serve stale cache. When enabled, the node can still use the expired cached files to respond to user requests even if the origin server is unavailable. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	ServeStale *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
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
	// When generating cache keys, include the client device type. Value range:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	UserDeviceType *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	// Include the client\\"s geographical location when generating the cache key. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	UserGeo *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	// Include the client\\"s language type when generating the cache key. Value range:
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
