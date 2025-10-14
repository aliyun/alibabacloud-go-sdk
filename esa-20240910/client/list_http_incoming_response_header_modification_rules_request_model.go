// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpIncomingResponseHeaderModificationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListHttpIncomingResponseHeaderModificationRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListHttpIncomingResponseHeaderModificationRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListHttpIncomingResponseHeaderModificationRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpIncomingResponseHeaderModificationRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListHttpIncomingResponseHeaderModificationRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListHttpIncomingResponseHeaderModificationRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListHttpIncomingResponseHeaderModificationRulesRequest
	GetSiteVersion() *int32
}

type ListHttpIncomingResponseHeaderModificationRulesRequest struct {
	// The configuration ID. You can call the ListHttpIncomingResponseHeaderModificationRules operation to query the ID.
	//
	// example:
	//
	// 432915173664768
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type to query. Valid values:
	//
	// 	- global: global configurations.
	//
	// 	- rule: rule configurations.
	//
	// example:
	//
	// rule
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 624516866852544
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the website configurations. You can use this parameter to specify a version of your website to apply the feature settings. By default, version 0 is used.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpIncomingResponseHeaderModificationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpIncomingResponseHeaderModificationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) SetConfigId(v int64) *ListHttpIncomingResponseHeaderModificationRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) SetConfigType(v string) *ListHttpIncomingResponseHeaderModificationRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) SetPageNumber(v int32) *ListHttpIncomingResponseHeaderModificationRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) SetPageSize(v int32) *ListHttpIncomingResponseHeaderModificationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) SetRuleName(v string) *ListHttpIncomingResponseHeaderModificationRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) SetSiteId(v int64) *ListHttpIncomingResponseHeaderModificationRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) SetSiteVersion(v int32) *ListHttpIncomingResponseHeaderModificationRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesRequest) Validate() error {
	return dara.Validate(s)
}
