// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceGroupByRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceResult(v *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) *GetAggregateResourceComplianceGroupByRegionResponseBody
	GetComplianceResult() *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult
	SetRequestId(v string) *GetAggregateResourceComplianceGroupByRegionResponseBody
	GetRequestId() *string
}

type GetAggregateResourceComplianceGroupByRegionResponseBody struct {
	// The queried evaluation results.
	ComplianceResult *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5E3A847A-5D40-54A1-A2CE-77A87823ED07
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBody) GetComplianceResult() *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult {
	return s.ComplianceResult
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBody) SetComplianceResult(v *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) *GetAggregateResourceComplianceGroupByRegionResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBody) SetRequestId(v string) *GetAggregateResourceComplianceGroupByRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBody) Validate() error {
	if s.ComplianceResult != nil {
		if err := s.ComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult struct {
	// The evaluation results grouped by region.
	ComplianceResultList []*GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList `json:"ComplianceResultList,omitempty" xml:"ComplianceResultList,omitempty" type:"Repeated"`
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) GetComplianceResultList() []*GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList {
	return s.ComplianceResultList
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) SetComplianceResultList(v []*GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult {
	s.ComplianceResultList = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) Validate() error {
	if s.ComplianceResultList != nil {
		for _, item := range s.ComplianceResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList struct {
	// The queried evaluation results.
	Compliances []*GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	// The region ID of the evaluated resource.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) GetCompliances() []*GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances {
	return s.Compliances
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) SetCompliances(v []*GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList {
	s.Compliances = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) SetRegionId(v string) *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList {
	s.RegionId = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) Validate() error {
	if s.Compliances != nil {
		for _, item := range s.Compliances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances struct {
	// The evaluation result. Valid values:
	//
	// 	- COMPLIANT: The resource is evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resource is evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to the resource.
	//
	// 	- INSUFFICIENT_DATA: No data is available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The total number of evaluation results.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) GetCount() *int64 {
	return s.Count
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) SetComplianceType(v string) *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) SetCount(v int64) *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances {
	s.Count = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) Validate() error {
	return dara.Validate(s)
}
