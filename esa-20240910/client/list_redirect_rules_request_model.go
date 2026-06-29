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
	// The configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. Valid values:
	//
	// - global: global configuration.
	//
	// - rule: rule configuration.
	//
	// example:
	//
	// rule
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule name. You can use this parameter to query the rule whose name matches the specified value.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456******
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
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
