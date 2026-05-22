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
	// This parameter is optional. If not provided, it does not distinguish between global and rule configurations.
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
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
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
