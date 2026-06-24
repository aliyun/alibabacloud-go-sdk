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
	// The page number. The default value is **1**.
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
	// The language of the request and response. The default value is **zh_cn**. Valid values:
	//
	// - **zh_cn**: Simplified Chinese.
	//
	// - **en_us**: US English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of template rules to return on each page. The default value is **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The risk level of the template rule. The value ranges from **1*	- to **11**. The default value is **null**. Valid values:
	//
	// - **1**: No risk.
	//
	// - **2**: S1.
	//
	// - **3**: S2.
	//
	// - **4**: S3.
	//
	// - **5**: S4.
	//
	// - **6**: S5.
	//
	// - **7**: S6.
	//
	// - **8**: S7.
	//
	// - **9**: S8.
	//
	// - **10**: S9.
	//
	// - **11**: S10.
	//
	// - **null**: All risk levels, including No risk, S1, S2, S3, S4, S5, S6, S7, S8, S9, and S10.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The status of the template rule. The default value is **null**. Valid values:
	//
	// - **0**: The rule is disabled.
	//
	// - **1**: The rule is enabled.
	//
	// - **null**: All rules are returned, regardless of their status.
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
