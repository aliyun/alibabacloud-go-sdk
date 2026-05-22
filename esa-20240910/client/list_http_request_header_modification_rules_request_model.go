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
	// Configuration ID, which can be obtained by calling the [ListHttpRequestHeaderModificationRules](https://help.aliyun.com/document_detail/2867483.html) API.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule configurations. Possible values:
	//
	// - global: Query global configuration;
	//
	// - rule: Query rule configuration;
	//
	// example:
	//
	// rule
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
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the configuration, defaulting to version 0.
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
