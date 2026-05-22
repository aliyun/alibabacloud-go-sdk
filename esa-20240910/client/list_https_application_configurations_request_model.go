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
	// Configuration ID.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule-based configurations. Possible values:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule-based configuration.
	//
	// This parameter is optional. If not provided, it will not distinguish between global and rule-based configurations.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Page number, default is 1 if not provided.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page, maximum is 500, default is 500 if not provided.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Rule name, which can be used to find the rule with the specified name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Site version number. For sites with version management enabled, this parameter can specify the site version for which the configuration is effective, default is version 0.
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
