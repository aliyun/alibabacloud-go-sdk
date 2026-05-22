// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListOriginRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListOriginRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListOriginRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOriginRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListOriginRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListOriginRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListOriginRulesRequest
	GetSiteVersion() *int32
}

type ListOriginRulesRequest struct {
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule configurations. Value range:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule configuration.
	//
	// This parameter is optional; if not provided, it does not distinguish between global and rule configurations.
	//
	// example:
	//
	// global
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
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Version number of the site. For sites with version management enabled, this parameter can specify the version of the site for which the configuration is effective, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListOriginRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOriginRulesRequest) GoString() string {
	return s.String()
}

func (s *ListOriginRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListOriginRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListOriginRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOriginRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOriginRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListOriginRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListOriginRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListOriginRulesRequest) SetConfigId(v int64) *ListOriginRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListOriginRulesRequest) SetConfigType(v string) *ListOriginRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListOriginRulesRequest) SetPageNumber(v int32) *ListOriginRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOriginRulesRequest) SetPageSize(v int32) *ListOriginRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOriginRulesRequest) SetRuleName(v string) *ListOriginRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListOriginRulesRequest) SetSiteId(v int64) *ListOriginRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListOriginRulesRequest) SetSiteVersion(v int32) *ListOriginRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListOriginRulesRequest) Validate() error {
	return dara.Validate(s)
}
