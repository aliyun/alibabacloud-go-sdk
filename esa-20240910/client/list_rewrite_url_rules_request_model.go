// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRewriteUrlRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListRewriteUrlRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListRewriteUrlRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListRewriteUrlRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRewriteUrlRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListRewriteUrlRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListRewriteUrlRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListRewriteUrlRulesRequest
	GetSiteVersion() *int32
}

type ListRewriteUrlRulesRequest struct {
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
	// This parameter is optional. If you do not specify this parameter, both global and rule configurations are returned. This parameter takes effect only when the functionName parameter is specified.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The page number for a paged query. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Valid values: 1 to 500.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule name. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. The default value is version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListRewriteUrlRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRewriteUrlRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRewriteUrlRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListRewriteUrlRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListRewriteUrlRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRewriteUrlRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRewriteUrlRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRewriteUrlRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListRewriteUrlRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListRewriteUrlRulesRequest) SetConfigId(v int64) *ListRewriteUrlRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetConfigType(v string) *ListRewriteUrlRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetPageNumber(v int32) *ListRewriteUrlRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetPageSize(v int32) *ListRewriteUrlRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetRuleName(v string) *ListRewriteUrlRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetSiteId(v int64) *ListRewriteUrlRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetSiteVersion(v int32) *ListRewriteUrlRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) Validate() error {
	return dara.Validate(s)
}
