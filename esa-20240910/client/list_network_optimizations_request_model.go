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
	// The configuration ID.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can use this parameter to query global or rule configurations. Valid values:
	//
	// - global: queries global configurations.
	//
	// - rule: queries rule configurations.
	//
	// This parameter is optional. If you do not specify this parameter, both global and rule configurations are returned.
	//
	// example:
	//
	// global
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
	// The rule name. You can use this parameter to query the rule that matches the specified name.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1231231221****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
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
