// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceGroupByResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceResult(v *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) *GetAggregateResourceComplianceGroupByResourceTypeResponseBody
	GetComplianceResult() *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult
	SetRequestId(v string) *GetAggregateResourceComplianceGroupByResourceTypeResponseBody
	GetRequestId() *string
}

type GetAggregateResourceComplianceGroupByResourceTypeResponseBody struct {
	// The queried evaluation results.
	ComplianceResult *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0CE6AEE1-B3D8-530A-9302-6606B20503BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBody) GetComplianceResult() *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult {
	return s.ComplianceResult
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBody) SetComplianceResult(v *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) *GetAggregateResourceComplianceGroupByResourceTypeResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBody) SetRequestId(v string) *GetAggregateResourceComplianceGroupByResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBody) Validate() error {
	if s.ComplianceResult != nil {
		if err := s.ComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult struct {
	// The evaluation results grouped by resource type.
	ComplianceResultList []*GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList `json:"ComplianceResultList,omitempty" xml:"ComplianceResultList,omitempty" type:"Repeated"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) GetComplianceResultList() []*GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList {
	return s.ComplianceResultList
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) SetComplianceResultList(v []*GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult {
	s.ComplianceResultList = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) Validate() error {
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

type GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList struct {
	// The queried evaluation results.
	Compliances []*GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	// The type of the evaluated resource.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) GetCompliances() []*GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances {
	return s.Compliances
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) SetCompliances(v []*GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList {
	s.Compliances = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) SetResourceType(v string) *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) Validate() error {
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

type GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances struct {
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
	// The number of statistical results.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) GetCount() *int64 {
	return s.Count
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) SetComplianceType(v string) *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) SetCount(v int64) *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances {
	s.Count = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) Validate() error {
	return dara.Validate(s)
}
