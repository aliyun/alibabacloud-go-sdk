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
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	BrowserCacheMode         *string `json:"BrowserCacheMode,omitempty" xml:"BrowserCacheMode,omitempty"`
	BrowserCacheTtl          *string `json:"BrowserCacheTtl,omitempty" xml:"BrowserCacheTtl,omitempty"`
	BypassCache              *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	CacheDeceptionArmor      *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	CacheReserveEligibility  *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	CheckPresenceCookie      *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	CheckPresenceHeader      *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
	// This parameter is required.
	ConfigId               *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	EdgeCacheMode          *string `json:"EdgeCacheMode,omitempty" xml:"EdgeCacheMode,omitempty"`
	EdgeCacheTtl           *string `json:"EdgeCacheTtl,omitempty" xml:"EdgeCacheTtl,omitempty"`
	EdgeStatusCodeCacheTtl *string `json:"EdgeStatusCodeCacheTtl,omitempty" xml:"EdgeStatusCodeCacheTtl,omitempty"`
	IncludeCookie          *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	IncludeHeader          *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// example:
	//
	// ignore
	PostBodyCacheKey  *string `json:"PostBodyCacheKey,omitempty" xml:"PostBodyCacheKey,omitempty"`
	PostBodySizeLimit *string `json:"PostBodySizeLimit,omitempty" xml:"PostBodySizeLimit,omitempty"`
	// example:
	//
	// on
	PostCache       *string `json:"PostCache,omitempty" xml:"PostCache,omitempty"`
	QueryString     *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	QueryStringMode *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	Rule            *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable      *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence        *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	ServeStale      *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	// This parameter is required.
	SiteId                  *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	UserDeviceType          *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	UserGeo                 *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	UserLanguage            *string `json:"UserLanguage,omitempty" xml:"UserLanguage,omitempty"`
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
