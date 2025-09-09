// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCategoryTemplateRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeCategoryTemplateRuleListRequest
	GetCurrentPage() *int32
	SetFeatureType(v int32) *DescribeCategoryTemplateRuleListRequest
	GetFeatureType() *int32
	SetLang(v string) *DescribeCategoryTemplateRuleListRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeCategoryTemplateRuleListRequest
	GetPageSize() *int32
	SetRiskLevelId(v int64) *DescribeCategoryTemplateRuleListRequest
	GetRiskLevelId() *int64
	SetStatus(v int32) *DescribeCategoryTemplateRuleListRequest
	GetStatus() *int32
}

type DescribeCategoryTemplateRuleListRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sensitivity level of the data that is not compliant with the rule. Valid values: **1*	- to **11**. Default value: **null**.
	//
	// 	- **1**: No sensitive data is detected.
	//
	// 	- **2**: specifies the S1 sensitivity level.
	//
	// 	- **3**: specifies the S2 sensitivity level.
	//
	// 	- **4**: specifies the S3 sensitivity level.
	//
	// 	- **5**: specifies the S4 sensitivity level.
	//
	// 	- **6**: specifies the S5 sensitivity level.
	//
	// 	- **7**: specifies the S6 sensitivity level.
	//
	// 	- **8**: specifies the S7 sensitivity level.
	//
	// 	- **9**: specifies the S8 sensitivity level.
	//
	// 	- **10**: specifies the S9 sensitivity level.
	//
	// 	- **11**: specifies the S10 sensitivity level.
	//
	// 	- **null**: specifies all preceding sensitivity levels.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The status of the rule. Default value: **null**. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// 	- **null**: all states
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCategoryTemplateRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryTemplateRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateRuleListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCategoryTemplateRuleListRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeCategoryTemplateRuleListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCategoryTemplateRuleListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCategoryTemplateRuleListRequest) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeCategoryTemplateRuleListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCategoryTemplateRuleListRequest) SetCurrentPage(v int32) *DescribeCategoryTemplateRuleListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) SetFeatureType(v int32) *DescribeCategoryTemplateRuleListRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) SetLang(v string) *DescribeCategoryTemplateRuleListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) SetPageSize(v int32) *DescribeCategoryTemplateRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) SetRiskLevelId(v int64) *DescribeCategoryTemplateRuleListRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) SetStatus(v int32) *DescribeCategoryTemplateRuleListRequest {
	s.Status = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) Validate() error {
	return dara.Validate(s)
}
