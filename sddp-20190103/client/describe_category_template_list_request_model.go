// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCategoryTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeCategoryTemplateListRequest
	GetCurrentPage() *int32
	SetFeatureType(v int32) *DescribeCategoryTemplateListRequest
	GetFeatureType() *int32
	SetLang(v string) *DescribeCategoryTemplateListRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeCategoryTemplateListRequest
	GetPageSize() *int32
	SetUsageScenario(v int32) *DescribeCategoryTemplateListRequest
	GetUsageScenario() *int32
}

type DescribeCategoryTemplateListRequest struct {
	// Page number for paginated queries. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// Set the language type for the request and response messages, default is **zh_cn**.
	//
	// Values:
	//
	// - **zh_cn**: Chinese (Simplified)
	//
	// - **en_us**: English (United States)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// When performing a paginated query, set the number of items per page. Default value is **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// API call scenario, default is **null**.
	//
	// Values:
	//
	// - **null**: Old version
	//
	// - **0**: Old version
	//
	// - **1**: New version
	//
	// example:
	//
	// 1
	UsageScenario *int32 `json:"UsageScenario,omitempty" xml:"UsageScenario,omitempty"`
}

func (s DescribeCategoryTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryTemplateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCategoryTemplateListRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeCategoryTemplateListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCategoryTemplateListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCategoryTemplateListRequest) GetUsageScenario() *int32 {
	return s.UsageScenario
}

func (s *DescribeCategoryTemplateListRequest) SetCurrentPage(v int32) *DescribeCategoryTemplateListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCategoryTemplateListRequest) SetFeatureType(v int32) *DescribeCategoryTemplateListRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeCategoryTemplateListRequest) SetLang(v string) *DescribeCategoryTemplateListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCategoryTemplateListRequest) SetPageSize(v int32) *DescribeCategoryTemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCategoryTemplateListRequest) SetUsageScenario(v int32) *DescribeCategoryTemplateListRequest {
	s.UsageScenario = &v
	return s
}

func (s *DescribeCategoryTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
