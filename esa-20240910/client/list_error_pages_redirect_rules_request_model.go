// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErrorPagesRedirectRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListErrorPagesRedirectRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListErrorPagesRedirectRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListErrorPagesRedirectRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListErrorPagesRedirectRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListErrorPagesRedirectRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListErrorPagesRedirectRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListErrorPagesRedirectRulesRequest
	GetSiteVersion() *int32
}

type ListErrorPagesRedirectRulesRequest struct {
	// The configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can use this parameter to query global configurations or rule configurations. Valid values:
	//
	// - global: queries global configurations.
	//
	// - rule: queries rule configurations.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 500. Default value: 500.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListErrorPagesRedirectRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListErrorPagesRedirectRulesRequest) GoString() string {
	return s.String()
}

func (s *ListErrorPagesRedirectRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListErrorPagesRedirectRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListErrorPagesRedirectRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListErrorPagesRedirectRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListErrorPagesRedirectRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListErrorPagesRedirectRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListErrorPagesRedirectRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListErrorPagesRedirectRulesRequest) SetConfigId(v int64) *ListErrorPagesRedirectRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListErrorPagesRedirectRulesRequest) SetConfigType(v string) *ListErrorPagesRedirectRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListErrorPagesRedirectRulesRequest) SetPageNumber(v int32) *ListErrorPagesRedirectRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListErrorPagesRedirectRulesRequest) SetPageSize(v int32) *ListErrorPagesRedirectRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListErrorPagesRedirectRulesRequest) SetRuleName(v string) *ListErrorPagesRedirectRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListErrorPagesRedirectRulesRequest) SetSiteId(v int64) *ListErrorPagesRedirectRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListErrorPagesRedirectRulesRequest) SetSiteVersion(v int32) *ListErrorPagesRedirectRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListErrorPagesRedirectRulesRequest) Validate() error {
	return dara.Validate(s)
}
