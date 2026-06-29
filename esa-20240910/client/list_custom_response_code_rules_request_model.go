// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomResponseCodeRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListCustomResponseCodeRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListCustomResponseCodeRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListCustomResponseCodeRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomResponseCodeRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListCustomResponseCodeRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListCustomResponseCodeRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListCustomResponseCodeRulesRequest
	GetSiteVersion() *int32
}

type ListCustomResponseCodeRulesRequest struct {
	// Configuration ID.
	//
	// example:
	//
	// 436490043994112
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type. You can use this parameter to query global or rule configurations. Valid values:
	//
	// - global: queries global configurations.
	//
	// - rule: queries rule configurations.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries per page. Maximum value: 500. Default value: 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Rule name.
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
	// 340035003106221
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListCustomResponseCodeRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomResponseCodeRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomResponseCodeRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListCustomResponseCodeRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListCustomResponseCodeRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomResponseCodeRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomResponseCodeRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCustomResponseCodeRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListCustomResponseCodeRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListCustomResponseCodeRulesRequest) SetConfigId(v int64) *ListCustomResponseCodeRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListCustomResponseCodeRulesRequest) SetConfigType(v string) *ListCustomResponseCodeRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListCustomResponseCodeRulesRequest) SetPageNumber(v int32) *ListCustomResponseCodeRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomResponseCodeRulesRequest) SetPageSize(v int32) *ListCustomResponseCodeRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomResponseCodeRulesRequest) SetRuleName(v string) *ListCustomResponseCodeRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListCustomResponseCodeRulesRequest) SetSiteId(v int64) *ListCustomResponseCodeRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListCustomResponseCodeRulesRequest) SetSiteVersion(v int32) *ListCustomResponseCodeRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListCustomResponseCodeRulesRequest) Validate() error {
	return dara.Validate(s)
}
