// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionReportCheckItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckItemList(v []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) *DescribeNisInspectionReportCheckItemsResponseBody
	GetCheckItemList() []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList
	SetInspectionReportId(v string) *DescribeNisInspectionReportCheckItemsResponseBody
	GetInspectionReportId() *string
	SetMaxResults(v int32) *DescribeNisInspectionReportCheckItemsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeNisInspectionReportCheckItemsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeNisInspectionReportCheckItemsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNisInspectionReportCheckItemsResponseBody
	GetTotalCount() *int32
}

type DescribeNisInspectionReportCheckItemsResponseBody struct {
	CheckItemList []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList `json:"CheckItemList,omitempty" xml:"CheckItemList,omitempty" type:"Repeated"`
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
	// hKrS+MVXkuOgztXnvdml194Cz/lMNdmr+DEh0th6dVlNEo/F148UPCh2itDku7Qj
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNisInspectionReportCheckItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportCheckItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) GetCheckItemList() []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList {
	return s.CheckItemList
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) SetCheckItemList(v []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) *DescribeNisInspectionReportCheckItemsResponseBody {
	s.CheckItemList = v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) SetInspectionReportId(v string) *DescribeNisInspectionReportCheckItemsResponseBody {
	s.InspectionReportId = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) SetMaxResults(v int32) *DescribeNisInspectionReportCheckItemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) SetNextToken(v string) *DescribeNisInspectionReportCheckItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) SetRequestId(v string) *DescribeNisInspectionReportCheckItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) SetTotalCount(v int32) *DescribeNisInspectionReportCheckItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBody) Validate() error {
	if s.CheckItemList != nil {
		for _, item := range s.CheckItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList struct {
	// example:
	//
	// stability
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// example:
	//
	// item_nat_water_level_check
	CheckItemCode   *string                                                                          `json:"CheckItemCode,omitempty" xml:"CheckItemCode,omitempty"`
	CheckItemName   *string                                                                          `json:"CheckItemName,omitempty" xml:"CheckItemName,omitempty"`
	CheckResultList []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList `json:"CheckResultList,omitempty" xml:"CheckResultList,omitempty" type:"Repeated"`
	Description     *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// item_ep_high_availability_check_description
	DescriptionCode    *string                                                                             `json:"DescriptionCode,omitempty" xml:"DescriptionCode,omitempty"`
	RecommendationList []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList `json:"RecommendationList,omitempty" xml:"RecommendationList,omitempty" type:"Repeated"`
	// example:
	//
	// NAT
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) GetCheckItemCode() *string {
	return s.CheckItemCode
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) GetCheckItemName() *string {
	return s.CheckItemName
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) GetCheckResultList() []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList {
	return s.CheckResultList
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) GetDescription() *string {
	return s.Description
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) GetDescriptionCode() *string {
	return s.DescriptionCode
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) GetRecommendationList() []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList {
	return s.RecommendationList
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) SetCategoryCode(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList {
	s.CategoryCode = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) SetCheckItemCode(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList {
	s.CheckItemCode = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) SetCheckItemName(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList {
	s.CheckItemName = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) SetCheckResultList(v []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList {
	s.CheckResultList = v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) SetDescription(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList {
	s.Description = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) SetDescriptionCode(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList {
	s.DescriptionCode = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) SetRecommendationList(v []*DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList {
	s.RecommendationList = v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) SetResourceType(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList {
	s.ResourceType = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemList) Validate() error {
	if s.CheckResultList != nil {
		for _, item := range s.CheckResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RecommendationList != nil {
		for _, item := range s.RecommendationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList struct {
	// example:
	//
	// 0
	Count     *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList) SetCount(v int32) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList {
	s.Count = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList) SetRiskLevel(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListCheckResultList) Validate() error {
	return dara.Validate(s)
}

type DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList struct {
	Abnormality *string `json:"Abnormality,omitempty" xml:"Abnormality,omitempty"`
	Metadata    *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ReasonCode  *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// nat_snat_cross_az_warn
	RecommendationCode *string `json:"RecommendationCode,omitempty" xml:"RecommendationCode,omitempty"`
	RiskLevel          *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Suggestion         *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	SuggestionCode     *string `json:"SuggestionCode,omitempty" xml:"SuggestionCode,omitempty"`
}

func (s DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) GetAbnormality() *string {
	return s.Abnormality
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) GetMetadata() *string {
	return s.Metadata
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) GetReason() *string {
	return s.Reason
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) GetRecommendationCode() *string {
	return s.RecommendationCode
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) GetSuggestion() *string {
	return s.Suggestion
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) GetSuggestionCode() *string {
	return s.SuggestionCode
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) SetAbnormality(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList {
	s.Abnormality = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) SetMetadata(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList {
	s.Metadata = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) SetReason(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList {
	s.Reason = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) SetReasonCode(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList {
	s.ReasonCode = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) SetRecommendationCode(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList {
	s.RecommendationCode = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) SetRiskLevel(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) SetSuggestion(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList {
	s.Suggestion = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) SetSuggestionCode(v string) *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList {
	s.SuggestionCode = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponseBodyCheckItemListRecommendationList) Validate() error {
	return dara.Validate(s)
}
