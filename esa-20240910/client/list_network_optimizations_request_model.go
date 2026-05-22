// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkOptimizationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListNetworkOptimizationsRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListNetworkOptimizationsRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListNetworkOptimizationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNetworkOptimizationsRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListNetworkOptimizationsRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListNetworkOptimizationsRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListNetworkOptimizationsRequest
	GetSiteVersion() *int32
}

type ListNetworkOptimizationsRequest struct {
	// Configuration ID.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule-based configurations. The value range is as follows:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule-based configuration.
	//
	// This parameter is optional; if not provided, it does not distinguish between global and rule-based configurations.
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
	// Rule name, which can be used to find the rule with the specified name.
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
	// 1231231221****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Site version number. For sites with version management enabled, this parameter can specify the site version for which the configuration takes effect, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListNetworkOptimizationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkOptimizationsRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkOptimizationsRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListNetworkOptimizationsRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListNetworkOptimizationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNetworkOptimizationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNetworkOptimizationsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListNetworkOptimizationsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListNetworkOptimizationsRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListNetworkOptimizationsRequest) SetConfigId(v int64) *ListNetworkOptimizationsRequest {
	s.ConfigId = &v
	return s
}

func (s *ListNetworkOptimizationsRequest) SetConfigType(v string) *ListNetworkOptimizationsRequest {
	s.ConfigType = &v
	return s
}

func (s *ListNetworkOptimizationsRequest) SetPageNumber(v int32) *ListNetworkOptimizationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNetworkOptimizationsRequest) SetPageSize(v int32) *ListNetworkOptimizationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListNetworkOptimizationsRequest) SetRuleName(v string) *ListNetworkOptimizationsRequest {
	s.RuleName = &v
	return s
}

func (s *ListNetworkOptimizationsRequest) SetSiteId(v int64) *ListNetworkOptimizationsRequest {
	s.SiteId = &v
	return s
}

func (s *ListNetworkOptimizationsRequest) SetSiteVersion(v int32) *ListNetworkOptimizationsRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListNetworkOptimizationsRequest) Validate() error {
	return dara.Validate(s)
}
