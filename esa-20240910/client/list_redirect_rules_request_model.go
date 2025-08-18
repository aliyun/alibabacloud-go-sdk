// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRedirectRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListRedirectRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListRedirectRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListRedirectRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRedirectRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListRedirectRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListRedirectRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListRedirectRulesRequest
	GetSiteVersion() *int32
}

type ListRedirectRulesRequest struct {
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type. Possible values:
	//
	// - global: Global configuration.
	//
	// - rule: Rule configuration.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, default is **500**, and the value range is **1~500**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Rule name, which can be used to find the rule with the specified name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456******
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the site configuration, with the default being version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListRedirectRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRedirectRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRedirectRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListRedirectRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListRedirectRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRedirectRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRedirectRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRedirectRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListRedirectRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListRedirectRulesRequest) SetConfigId(v int64) *ListRedirectRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListRedirectRulesRequest) SetConfigType(v string) *ListRedirectRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListRedirectRulesRequest) SetPageNumber(v int32) *ListRedirectRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRedirectRulesRequest) SetPageSize(v int32) *ListRedirectRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRedirectRulesRequest) SetRuleName(v string) *ListRedirectRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListRedirectRulesRequest) SetSiteId(v int64) *ListRedirectRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListRedirectRulesRequest) SetSiteVersion(v int32) *ListRedirectRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListRedirectRulesRequest) Validate() error {
	return dara.Validate(s)
}
