// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpRequestHeaderModificationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListHttpRequestHeaderModificationRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListHttpRequestHeaderModificationRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListHttpRequestHeaderModificationRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpRequestHeaderModificationRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListHttpRequestHeaderModificationRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListHttpRequestHeaderModificationRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListHttpRequestHeaderModificationRulesRequest
	GetSiteVersion() *int32
}

type ListHttpRequestHeaderModificationRulesRequest struct {
	// The configuration ID. You can call the [ListHttpRequestHeaderModificationRules](https://help.aliyun.com/document_detail/2867483.html) operation to obtain the configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can use this parameter to query global or rule configurations. Valid values:
	//
	// - global: queries global configurations.
	//
	// - rule: queries rule configurations.
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
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. The default value is version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpRequestHeaderModificationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpRequestHeaderModificationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetConfigId(v int64) *ListHttpRequestHeaderModificationRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetConfigType(v string) *ListHttpRequestHeaderModificationRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetPageNumber(v int32) *ListHttpRequestHeaderModificationRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetPageSize(v int32) *ListHttpRequestHeaderModificationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetRuleName(v string) *ListHttpRequestHeaderModificationRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetSiteId(v int64) *ListHttpRequestHeaderModificationRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetSiteVersion(v int32) *ListHttpRequestHeaderModificationRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) Validate() error {
	return dara.Validate(s)
}
