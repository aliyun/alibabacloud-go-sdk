// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceGroupByResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceResult(v *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) *GetResourceComplianceGroupByResourceTypeResponseBody
	GetComplianceResult() *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult
	SetRequestId(v string) *GetResourceComplianceGroupByResourceTypeResponseBody
	GetRequestId() *string
}

type GetResourceComplianceGroupByResourceTypeResponseBody struct {
	// The queried evaluation results.
	ComplianceResult *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 84610B68-2DD3-5AF0-B68D-E1FA8F051F7D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceComplianceGroupByResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBody) GetComplianceResult() *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult {
	return s.ComplianceResult
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBody) SetComplianceResult(v *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) *GetResourceComplianceGroupByResourceTypeResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBody) SetRequestId(v string) *GetResourceComplianceGroupByResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBody) Validate() error {
	if s.ComplianceResult != nil {
		if err := s.ComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult struct {
	// The evaluation results grouped by resource type.
	ComplianceResultList []*GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList `json:"ComplianceResultList,omitempty" xml:"ComplianceResultList,omitempty" type:"Repeated"`
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) GetComplianceResultList() []*GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList {
	return s.ComplianceResultList
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) SetComplianceResultList(v []*GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult {
	s.ComplianceResultList = v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) Validate() error {
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

type GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList struct {
	// The queried evaluation results.
	Compliances []*GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	// The type of the evaluated resource.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) GetCompliances() []*GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances {
	return s.Compliances
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) SetCompliances(v []*GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList {
	s.Compliances = v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) SetResourceType(v string) *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList {
	s.ResourceType = &v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) Validate() error {
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

type GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances struct {
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

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) GetCount() *int64 {
	return s.Count
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) SetComplianceType(v string) *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) SetCount(v int64) *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances {
	s.Count = &v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) Validate() error {
	return dara.Validate(s)
}
