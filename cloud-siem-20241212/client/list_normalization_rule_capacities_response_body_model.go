// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRuleCapacitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNormalizationRuleCapacitiesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationRuleCapacitiesResponseBody
	GetNextToken() *string
	SetNormalizationRuleCapacities(v []*ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) *ListNormalizationRuleCapacitiesResponseBody
	GetNormalizationRuleCapacities() []*ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities
	SetPageNumber(v int32) *ListNormalizationRuleCapacitiesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNormalizationRuleCapacitiesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListNormalizationRuleCapacitiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNormalizationRuleCapacitiesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListNormalizationRuleCapacitiesResponseBody
	GetTotalPage() *int32
}

type ListNormalizationRuleCapacitiesResponseBody struct {
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken                   *string                                                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NormalizationRuleCapacities []*ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities `json:"NormalizationRuleCapacities,omitempty" xml:"NormalizationRuleCapacities,omitempty" type:"Repeated"`
	// example:
	//
	// 1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 3。
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListNormalizationRuleCapacitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRuleCapacitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNormalizationRuleCapacitiesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationRuleCapacitiesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationRuleCapacitiesResponseBody) GetNormalizationRuleCapacities() []*ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities {
	return s.NormalizationRuleCapacities
}

func (s *ListNormalizationRuleCapacitiesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNormalizationRuleCapacitiesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNormalizationRuleCapacitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNormalizationRuleCapacitiesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNormalizationRuleCapacitiesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListNormalizationRuleCapacitiesResponseBody) SetMaxResults(v int32) *ListNormalizationRuleCapacitiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBody) SetNextToken(v string) *ListNormalizationRuleCapacitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBody) SetNormalizationRuleCapacities(v []*ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) *ListNormalizationRuleCapacitiesResponseBody {
	s.NormalizationRuleCapacities = v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBody) SetPageNumber(v int32) *ListNormalizationRuleCapacitiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBody) SetPageSize(v int32) *ListNormalizationRuleCapacitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBody) SetRequestId(v string) *ListNormalizationRuleCapacitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBody) SetTotalCount(v int32) *ListNormalizationRuleCapacitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBody) SetTotalPage(v int32) *ListNormalizationRuleCapacitiesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBody) Validate() error {
	if s.NormalizationRuleCapacities != nil {
		for _, item := range s.NormalizationRuleCapacities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities struct {
	Capacities []*string `json:"Capacities,omitempty" xml:"Capacities,omitempty" type:"Repeated"`
	// example:
	//
	// detection_preset_rule。
	CapacityType *string `json:"CapacityType,omitempty" xml:"CapacityType,omitempty"`
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp。
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
}

func (s ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) GoString() string {
	return s.String()
}

func (s *ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) GetCapacities() []*string {
	return s.Capacities
}

func (s *ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) GetCapacityType() *string {
	return s.CapacityType
}

func (s *ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) SetCapacities(v []*string) *ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities {
	s.Capacities = v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) SetCapacityType(v string) *ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities {
	s.CapacityType = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) SetNormalizationRuleId(v string) *ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities {
	s.NormalizationRuleId = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponseBodyNormalizationRuleCapacities) Validate() error {
	return dara.Validate(s)
}
