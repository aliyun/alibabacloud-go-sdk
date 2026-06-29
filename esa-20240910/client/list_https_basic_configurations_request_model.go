// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpsBasicConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListHttpsBasicConfigurationsRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListHttpsBasicConfigurationsRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListHttpsBasicConfigurationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpsBasicConfigurationsRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListHttpsBasicConfigurationsRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListHttpsBasicConfigurationsRequest
	GetSiteId() *int64
}

type ListHttpsBasicConfigurationsRequest struct {
	// The configuration ID.
	//
	// example:
	//
	// 35281609698****
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
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListHttpsBasicConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsBasicConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListHttpsBasicConfigurationsRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpsBasicConfigurationsRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpsBasicConfigurationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpsBasicConfigurationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpsBasicConfigurationsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpsBasicConfigurationsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListHttpsBasicConfigurationsRequest) SetConfigId(v int64) *ListHttpsBasicConfigurationsRequest {
	s.ConfigId = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) SetConfigType(v string) *ListHttpsBasicConfigurationsRequest {
	s.ConfigType = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) SetPageNumber(v int32) *ListHttpsBasicConfigurationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) SetPageSize(v int32) *ListHttpsBasicConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) SetRuleName(v string) *ListHttpsBasicConfigurationsRequest {
	s.RuleName = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) SetSiteId(v int64) *ListHttpsBasicConfigurationsRequest {
	s.SiteId = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
