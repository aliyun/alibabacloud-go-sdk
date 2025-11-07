// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionRecommendationResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionReportId(v string) *DescribeNisInspectionRecommendationResourcesRequest
	GetInspectionReportId() *string
	SetLanguage(v string) *DescribeNisInspectionRecommendationResourcesRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeNisInspectionRecommendationResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeNisInspectionRecommendationResourcesRequest
	GetNextToken() *string
	SetRecommendationCode(v string) *DescribeNisInspectionRecommendationResourcesRequest
	GetRecommendationCode() *string
}

type DescribeNisInspectionRecommendationResourcesRequest struct {
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
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// nat_snat_cross_az_warn
	RecommendationCode *string `json:"RecommendationCode,omitempty" xml:"RecommendationCode,omitempty"`
}

func (s DescribeNisInspectionRecommendationResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionRecommendationResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) GetRecommendationCode() *string {
	return s.RecommendationCode
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) SetInspectionReportId(v string) *DescribeNisInspectionRecommendationResourcesRequest {
	s.InspectionReportId = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) SetLanguage(v string) *DescribeNisInspectionRecommendationResourcesRequest {
	s.Language = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) SetMaxResults(v int32) *DescribeNisInspectionRecommendationResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) SetNextToken(v string) *DescribeNisInspectionRecommendationResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) SetRecommendationCode(v string) *DescribeNisInspectionRecommendationResourcesRequest {
	s.RecommendationCode = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesRequest) Validate() error {
	return dara.Validate(s)
}
