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
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	BrowserCacheMode         *string `json:"BrowserCacheMode,omitempty" xml:"BrowserCacheMode,omitempty"`
	BrowserCacheTtl          *string `json:"BrowserCacheTtl,omitempty" xml:"BrowserCacheTtl,omitempty"`
	BypassCache              *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	CacheDeceptionArmor      *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	CacheReserveEligibility  *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	CheckPresenceCookie      *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	CheckPresenceHeader      *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
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
	SiteVersion             *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	UserDeviceType          *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	UserGeo                 *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	UserLanguage            *string `json:"UserLanguage,omitempty" xml:"UserLanguage,omitempty"`
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
