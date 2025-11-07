// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionRecommendationResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionReportId(v string) *DescribeNisInspectionRecommendationResourcesResponseBody
	GetInspectionReportId() *string
	SetMaxResults(v int32) *DescribeNisInspectionRecommendationResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeNisInspectionRecommendationResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeNisInspectionRecommendationResourcesResponseBody
	GetRequestId() *string
	SetResourceList(v []*DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) *DescribeNisInspectionRecommendationResourcesResponseBody
	GetResourceList() []*DescribeNisInspectionRecommendationResourcesResponseBodyResourceList
	SetTotalCount(v int32) *DescribeNisInspectionRecommendationResourcesResponseBody
	GetTotalCount() *int32
}

type DescribeNisInspectionRecommendationResourcesResponseBody struct {
	// example:
	//
	// nir-ffd1af****196d0
	InspectionReportId *string `json:"InspectionReportId,omitempty" xml:"InspectionReportId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// qt0DqY2lXxwBt9/ROQoS/7J9p90D1vF2vFbwzb/1oSWr3AxcM6/KpObZ7Z1PZdcV
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId    *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceList []*DescribeNisInspectionRecommendationResourcesResponseBodyResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// example:
	//
	// 192
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNisInspectionRecommendationResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionRecommendationResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) GetResourceList() []*DescribeNisInspectionRecommendationResourcesResponseBodyResourceList {
	return s.ResourceList
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) SetInspectionReportId(v string) *DescribeNisInspectionRecommendationResourcesResponseBody {
	s.InspectionReportId = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) SetMaxResults(v int32) *DescribeNisInspectionRecommendationResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) SetNextToken(v string) *DescribeNisInspectionRecommendationResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) SetRequestId(v string) *DescribeNisInspectionRecommendationResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) SetResourceList(v []*DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) *DescribeNisInspectionRecommendationResourcesResponseBody {
	s.ResourceList = v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) SetTotalCount(v int32) *DescribeNisInspectionRecommendationResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBody) Validate() error {
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNisInspectionRecommendationResourcesResponseBodyResourceList struct {
	// example:
	//
	// {ResourceId: "ngw-p0wn04hi4****q2us6q7q"}
	AnalysisData *string `json:"AnalysisData,omitempty" xml:"AnalysisData,omitempty"`
	// example:
	//
	// ngw-p0wn04hi4****q2us6q7q
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) GetAnalysisData() *string {
	return s.AnalysisData
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) SetAnalysisData(v string) *DescribeNisInspectionRecommendationResourcesResponseBodyResourceList {
	s.AnalysisData = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) SetResourceId(v string) *DescribeNisInspectionRecommendationResourcesResponseBodyResourceList {
	s.ResourceId = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) SetResourceName(v string) *DescribeNisInspectionRecommendationResourcesResponseBodyResourceList {
	s.ResourceName = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponseBodyResourceList) Validate() error {
	return dara.Validate(s)
}
