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
	// The configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can use this parameter to query global or rule configurations. Valid values:
	//
	// - global: Query global configurations.
	//
	// - rule: Query rule configurations.
	//
	// This parameter is optional. If not specified, both global and rule configurations are returned without distinction.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The page number for paginated queries. The value must be greater than or equal to 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paginated queries. Valid values: 1 to 500. Default value: 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule name. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The site ID. You can obtain the site ID by calling the [ListSites](~~ListSites~~) API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
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
