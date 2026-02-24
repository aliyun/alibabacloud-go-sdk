// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceByConfigRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceResult(v *GetResourceComplianceByConfigRuleResponseBodyComplianceResult) *GetResourceComplianceByConfigRuleResponseBody
	GetComplianceResult() *GetResourceComplianceByConfigRuleResponseBodyComplianceResult
	SetRequestId(v string) *GetResourceComplianceByConfigRuleResponseBody
	GetRequestId() *string
}

type GetResourceComplianceByConfigRuleResponseBody struct {
	ComplianceResult *GetResourceComplianceByConfigRuleResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	RequestId        *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceComplianceByConfigRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceByConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByConfigRuleResponseBody) GetComplianceResult() *GetResourceComplianceByConfigRuleResponseBodyComplianceResult {
	return s.ComplianceResult
}

func (s *GetResourceComplianceByConfigRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceComplianceByConfigRuleResponseBody) SetComplianceResult(v *GetResourceComplianceByConfigRuleResponseBodyComplianceResult) *GetResourceComplianceByConfigRuleResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponseBody) SetRequestId(v string) *GetResourceComplianceByConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponseBody) Validate() error {
	if s.ComplianceResult != nil {
		if err := s.ComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceComplianceByConfigRuleResponseBodyComplianceResult struct {
	Compliances []*GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	TotalCount  *int64                                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetResourceComplianceByConfigRuleResponseBodyComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceByConfigRuleResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResult) GetCompliances() []*GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances {
	return s.Compliances
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResult) SetCompliances(v []*GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) *GetResourceComplianceByConfigRuleResponseBodyComplianceResult {
	s.Compliances = v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResult) SetTotalCount(v int64) *GetResourceComplianceByConfigRuleResponseBodyComplianceResult {
	s.TotalCount = &v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResult) Validate() error {
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

type GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) GetCount() *int32 {
	return s.Count
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) SetComplianceType(v string) *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) SetCount(v int32) *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances {
	s.Count = &v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) Validate() error {
	return dara.Validate(s)
}
