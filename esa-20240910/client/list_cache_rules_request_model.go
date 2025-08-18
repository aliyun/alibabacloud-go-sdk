// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCacheRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListCacheRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListCacheRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListCacheRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCacheRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListCacheRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListCacheRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListCacheRulesRequest
	GetSiteVersion() *int32
}

type ListCacheRulesRequest struct {
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule-based configurations. Possible values:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule-based configuration.
	//
	// This parameter is optional; if not provided, it will not distinguish between global and rule-based configurations.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Page number, defaulting to 1 if not provided.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page, with a maximum of 500. Defaults to 500 if not provided.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Site version number. For sites with version management enabled, this parameter can specify the site version for which the configuration takes effect, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListCacheRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCacheRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCacheRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListCacheRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListCacheRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCacheRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCacheRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCacheRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListCacheRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListCacheRulesRequest) SetConfigId(v int64) *ListCacheRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListCacheRulesRequest) SetConfigType(v string) *ListCacheRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListCacheRulesRequest) SetPageNumber(v int32) *ListCacheRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCacheRulesRequest) SetPageSize(v int32) *ListCacheRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCacheRulesRequest) SetRuleName(v string) *ListCacheRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListCacheRulesRequest) SetSiteId(v int64) *ListCacheRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListCacheRulesRequest) SetSiteVersion(v int32) *ListCacheRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListCacheRulesRequest) Validate() error {
	return dara.Validate(s)
}
