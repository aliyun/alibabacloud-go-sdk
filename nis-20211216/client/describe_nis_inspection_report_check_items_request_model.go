// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionReportCheckItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryCode(v string) *DescribeNisInspectionReportCheckItemsRequest
	GetCategoryCode() *string
	SetInspectionReportId(v string) *DescribeNisInspectionReportCheckItemsRequest
	GetInspectionReportId() *string
	SetLanguage(v string) *DescribeNisInspectionReportCheckItemsRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeNisInspectionReportCheckItemsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeNisInspectionReportCheckItemsRequest
	GetNextToken() *string
	SetResourceType(v []*string) *DescribeNisInspectionReportCheckItemsRequest
	GetResourceType() []*string
	SetRiskLevel(v []*string) *DescribeNisInspectionReportCheckItemsRequest
	GetRiskLevel() []*string
}

type DescribeNisInspectionReportCheckItemsRequest struct {
	// The category of the check item.
	//
	// example:
	//
	// stability
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// The ID of the inspection report.
	//
	// This parameter is required.
	//
	// example:
	//
	// nir-ffd1af****196d0
	InspectionReportId *string `json:"InspectionReportId,omitempty" xml:"InspectionReportId,omitempty"`
	// The language of the content. Valid values: zh-CN and en-US.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The maximum number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Set this parameter to the NextToken value returned from the previous call.
	//
	// example:
	//
	// hKrS+MVXkuOgztXnvdml16/uO3mvCyHxSjzdhx9VRUC+8umDTIV2Wg9TTOUrR7ve
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource type.
	ResourceType []*string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Repeated"`
	// A collection of risk levels.
	RiskLevel []*string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
}

func (s DescribeNisInspectionReportCheckItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportCheckItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportCheckItemsRequest) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *DescribeNisInspectionReportCheckItemsRequest) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *DescribeNisInspectionReportCheckItemsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeNisInspectionReportCheckItemsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNisInspectionReportCheckItemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNisInspectionReportCheckItemsRequest) GetResourceType() []*string {
	return s.ResourceType
}

func (s *DescribeNisInspectionReportCheckItemsRequest) GetRiskLevel() []*string {
	return s.RiskLevel
}

func (s *DescribeNisInspectionReportCheckItemsRequest) SetCategoryCode(v string) *DescribeNisInspectionReportCheckItemsRequest {
	s.CategoryCode = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsRequest) SetInspectionReportId(v string) *DescribeNisInspectionReportCheckItemsRequest {
	s.InspectionReportId = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsRequest) SetLanguage(v string) *DescribeNisInspectionReportCheckItemsRequest {
	s.Language = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsRequest) SetMaxResults(v int32) *DescribeNisInspectionReportCheckItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsRequest) SetNextToken(v string) *DescribeNisInspectionReportCheckItemsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsRequest) SetResourceType(v []*string) *DescribeNisInspectionReportCheckItemsRequest {
	s.ResourceType = v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsRequest) SetRiskLevel(v []*string) *DescribeNisInspectionReportCheckItemsRequest {
	s.RiskLevel = v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsRequest) Validate() error {
	return dara.Validate(s)
}
