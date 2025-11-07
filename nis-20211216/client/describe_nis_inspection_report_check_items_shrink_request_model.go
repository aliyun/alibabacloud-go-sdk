// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionReportCheckItemsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryCode(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest
	GetCategoryCode() *string
	SetInspectionReportId(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest
	GetInspectionReportId() *string
	SetLanguage(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeNisInspectionReportCheckItemsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest
	GetNextToken() *string
	SetResourceTypeShrink(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest
	GetResourceTypeShrink() *string
	SetRiskLevelShrink(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest
	GetRiskLevelShrink() *string
}

type DescribeNisInspectionReportCheckItemsShrinkRequest struct {
	// example:
	//
	// stability
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// nir-ffd1af****196d0
	InspectionReportId *string `json:"InspectionReportId,omitempty" xml:"InspectionReportId,omitempty"`
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// hKrS+MVXkuOgztXnvdml16/uO3mvCyHxSjzdhx9VRUC+8umDTIV2Wg9TTOUrR7ve
	NextToken          *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceTypeShrink *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RiskLevelShrink    *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeNisInspectionReportCheckItemsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportCheckItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) GetResourceTypeShrink() *string {
	return s.ResourceTypeShrink
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) GetRiskLevelShrink() *string {
	return s.RiskLevelShrink
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) SetCategoryCode(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest {
	s.CategoryCode = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) SetInspectionReportId(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest {
	s.InspectionReportId = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) SetLanguage(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest {
	s.Language = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) SetMaxResults(v int32) *DescribeNisInspectionReportCheckItemsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) SetNextToken(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) SetResourceTypeShrink(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest {
	s.ResourceTypeShrink = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) SetRiskLevelShrink(v string) *DescribeNisInspectionReportCheckItemsShrinkRequest {
	s.RiskLevelShrink = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
