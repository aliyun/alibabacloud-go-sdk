// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpResponseHeaderModificationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListHttpResponseHeaderModificationRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListHttpResponseHeaderModificationRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListHttpResponseHeaderModificationRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpResponseHeaderModificationRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListHttpResponseHeaderModificationRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListHttpResponseHeaderModificationRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListHttpResponseHeaderModificationRulesRequest
	GetSiteVersion() *int32
}

type ListHttpResponseHeaderModificationRulesRequest struct {
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule configurations. Possible values:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule configuration.
	//
	// This parameter is optional. If not provided, it does not distinguish between global and rule configurations.
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
	// Page size, default is 500. The value range is any integer from 1 to 500.
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
	// Version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the configuration, with the default being version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpResponseHeaderModificationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpResponseHeaderModificationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetConfigId(v int64) *ListHttpResponseHeaderModificationRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetConfigType(v string) *ListHttpResponseHeaderModificationRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetPageNumber(v int32) *ListHttpResponseHeaderModificationRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetPageSize(v int32) *ListHttpResponseHeaderModificationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetRuleName(v string) *ListHttpResponseHeaderModificationRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetSiteId(v int64) *ListHttpResponseHeaderModificationRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetSiteVersion(v int32) *ListHttpResponseHeaderModificationRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) Validate() error {
	return dara.Validate(s)
}
