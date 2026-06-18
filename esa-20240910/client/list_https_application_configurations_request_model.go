// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpsApplicationConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListHttpsApplicationConfigurationsRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListHttpsApplicationConfigurationsRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListHttpsApplicationConfigurationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpsApplicationConfigurationsRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListHttpsApplicationConfigurationsRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListHttpsApplicationConfigurationsRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListHttpsApplicationConfigurationsRequest
	GetSiteVersion() *int32
}

type ListHttpsApplicationConfigurationsRequest struct {
	// The configuration ID.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. Use this parameter to query the global configuration or rule configurations. Valid values:
	//
	// - global: Queries the global configuration.
	//
	// - rule: Queries rule configurations.
	//
	// If this parameter is omitted, the query returns both global and rule configurations.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The page number to return. If this parameter is omitted, the default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. The maximum value is 500. If this parameter is omitted, the default value is 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule name. This parameter filters the results to include only the rule with the specified name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The site ID. You can obtain this ID by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site version. For sites with version management enabled, use this parameter to retrieve the configuration for a specific site version. The default value is 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpsApplicationConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsApplicationConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListHttpsApplicationConfigurationsRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpsApplicationConfigurationsRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpsApplicationConfigurationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpsApplicationConfigurationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpsApplicationConfigurationsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpsApplicationConfigurationsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListHttpsApplicationConfigurationsRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpsApplicationConfigurationsRequest) SetConfigId(v int64) *ListHttpsApplicationConfigurationsRequest {
	s.ConfigId = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsRequest) SetConfigType(v string) *ListHttpsApplicationConfigurationsRequest {
	s.ConfigType = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsRequest) SetPageNumber(v int32) *ListHttpsApplicationConfigurationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsRequest) SetPageSize(v int32) *ListHttpsApplicationConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsRequest) SetRuleName(v string) *ListHttpsApplicationConfigurationsRequest {
	s.RuleName = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsRequest) SetSiteId(v int64) *ListHttpsApplicationConfigurationsRequest {
	s.SiteId = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsRequest) SetSiteVersion(v int32) *ListHttpsApplicationConfigurationsRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
