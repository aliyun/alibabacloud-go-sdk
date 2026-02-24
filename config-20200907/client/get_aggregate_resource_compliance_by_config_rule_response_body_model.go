// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceByConfigRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceResult(v *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) *GetAggregateResourceComplianceByConfigRuleResponseBody
	GetComplianceResult() *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult
	SetRequestId(v string) *GetAggregateResourceComplianceByConfigRuleResponseBody
	GetRequestId() *string
}

type GetAggregateResourceComplianceByConfigRuleResponseBody struct {
	ComplianceResult *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	RequestId        *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBody) GetComplianceResult() *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult {
	return s.ComplianceResult
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBody) SetComplianceResult(v *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) *GetAggregateResourceComplianceByConfigRuleResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBody) SetRequestId(v string) *GetAggregateResourceComplianceByConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBody) Validate() error {
	if s.ComplianceResult != nil {
		if err := s.ComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult struct {
	Compliances []*GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	TotalCount  *int64                                                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) GetCompliances() []*GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances {
	return s.Compliances
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) SetCompliances(v []*GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult {
	s.Compliances = v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) SetTotalCount(v int64) *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult {
	s.TotalCount = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) Validate() error {
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

type GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) GetCount() *int32 {
	return s.Count
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) SetComplianceType(v string) *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) SetCount(v int32) *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances {
	s.Count = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) Validate() error {
	return dara.Validate(s)
}
