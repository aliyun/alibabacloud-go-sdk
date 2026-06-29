// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompressionRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListCompressionRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListCompressionRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListCompressionRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCompressionRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListCompressionRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListCompressionRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListCompressionRulesRequest
	GetSiteVersion() *int32
}

type ListCompressionRulesRequest struct {
	// The configuration ID. You can call the [ListRedirectRules](https://help.aliyun.com/document_detail/2867474.html) operation to obtain the configuration ID.
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
	// 34003500310****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListCompressionRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCompressionRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCompressionRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListCompressionRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListCompressionRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCompressionRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCompressionRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCompressionRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListCompressionRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListCompressionRulesRequest) SetConfigId(v int64) *ListCompressionRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListCompressionRulesRequest) SetConfigType(v string) *ListCompressionRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListCompressionRulesRequest) SetPageNumber(v int32) *ListCompressionRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCompressionRulesRequest) SetPageSize(v int32) *ListCompressionRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCompressionRulesRequest) SetRuleName(v string) *ListCompressionRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListCompressionRulesRequest) SetSiteId(v int64) *ListCompressionRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListCompressionRulesRequest) SetSiteVersion(v int32) *ListCompressionRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListCompressionRulesRequest) Validate() error {
	return dara.Validate(s)
}
