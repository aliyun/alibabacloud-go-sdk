// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteFunctionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListSiteFunctionsRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListSiteFunctionsRequest
	GetConfigType() *string
	SetFunctionName(v string) *ListSiteFunctionsRequest
	GetFunctionName() *string
	SetPageNumber(v int32) *ListSiteFunctionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSiteFunctionsRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListSiteFunctionsRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListSiteFunctionsRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListSiteFunctionsRequest
	GetSiteVersion() *int32
}

type ListSiteFunctionsRequest struct {
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can use this parameter to query global configurations or rule configurations. Valid values:
	//
	// - global: queries global configurations.
	//
	// - rule: queries rule configurations.
	//
	// This parameter is optional. If not specified, both global and rule configurations are returned. This parameter takes effect only when the FunctionName parameter is specified.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The function name of the feature.
	//
	// example:
	//
	// CacheRules
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
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
	// The rule name. Returns the rule that matches the specified name. This parameter takes effect only when the FunctionName parameter is specified.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 340035003106221
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListSiteFunctionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListSiteFunctionsRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListSiteFunctionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSiteFunctionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSiteFunctionsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListSiteFunctionsRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListSiteFunctionsRequest) SetConfigId(v int64) *ListSiteFunctionsRequest {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsRequest) SetConfigType(v string) *ListSiteFunctionsRequest {
	s.ConfigType = &v
	return s
}

func (s *ListSiteFunctionsRequest) SetFunctionName(v string) *ListSiteFunctionsRequest {
	s.FunctionName = &v
	return s
}

func (s *ListSiteFunctionsRequest) SetPageNumber(v int32) *ListSiteFunctionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSiteFunctionsRequest) SetPageSize(v int32) *ListSiteFunctionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSiteFunctionsRequest) SetRuleName(v string) *ListSiteFunctionsRequest {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsRequest) SetSiteId(v int64) *ListSiteFunctionsRequest {
	s.SiteId = &v
	return s
}

func (s *ListSiteFunctionsRequest) SetSiteVersion(v int32) *ListSiteFunctionsRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListSiteFunctionsRequest) Validate() error {
	return dara.Validate(s)
}
