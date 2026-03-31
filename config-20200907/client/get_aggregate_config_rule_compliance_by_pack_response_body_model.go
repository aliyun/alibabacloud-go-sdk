// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRuleComplianceByPackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleComplianceResult(v *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) *GetAggregateConfigRuleComplianceByPackResponseBody
	GetConfigRuleComplianceResult() *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult
	SetRequestId(v string) *GetAggregateConfigRuleComplianceByPackResponseBody
	GetRequestId() *string
}

type GetAggregateConfigRuleComplianceByPackResponseBody struct {
	// The compliance evaluation results that are returned by rules in the compliance package.
	ConfigRuleComplianceResult *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult `json:"ConfigRuleComplianceResult,omitempty" xml:"ConfigRuleComplianceResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C6B0C0A8-3245-48F1-AEAB-BC1A446E99D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateConfigRuleComplianceByPackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleComplianceByPackResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBody) GetConfigRuleComplianceResult() *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	return s.ConfigRuleComplianceResult
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBody) SetConfigRuleComplianceResult(v *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) *GetAggregateConfigRuleComplianceByPackResponseBody {
	s.ConfigRuleComplianceResult = v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBody) SetRequestId(v string) *GetAggregateConfigRuleComplianceByPackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBody) Validate() error {
	if s.ConfigRuleComplianceResult != nil {
		if err := s.ConfigRuleComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult struct {
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-541e626622af0087****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliantCount   *int32  `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	// The information about rules in the compliance package.
	ConfigRuleCompliances []*GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances `json:"ConfigRuleCompliances,omitempty" xml:"ConfigRuleCompliances,omitempty" type:"Repeated"`
	IgnoredCount          *int32                                                                                               `json:"IgnoredCount,omitempty" xml:"IgnoredCount,omitempty"`
	InsufficientDataCount *int32                                                                                               `json:"InsufficientDataCount,omitempty" xml:"InsufficientDataCount,omitempty"`
	// The number of rules against which specific resources are evaluated as non-compliant.
	//
	// example:
	//
	// 0
	NonCompliantCount  *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	NotApplicableCount *int32 `json:"NotApplicableCount,omitempty" xml:"NotApplicableCount,omitempty"`
	// The total number of rules in the compliance package.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetCompliantCount() *int32 {
	return s.CompliantCount
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetConfigRuleCompliances() []*GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	return s.ConfigRuleCompliances
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetIgnoredCount() *int32 {
	return s.IgnoredCount
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetInsufficientDataCount() *int32 {
	return s.InsufficientDataCount
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetNotApplicableCount() *int32 {
	return s.NotApplicableCount
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetCompliancePackId(v string) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetCompliantCount(v int32) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.CompliantCount = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetConfigRuleCompliances(v []*GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.ConfigRuleCompliances = v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetIgnoredCount(v int32) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.IgnoredCount = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetInsufficientDataCount(v int32) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.InsufficientDataCount = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetNonCompliantCount(v int32) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.NonCompliantCount = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetNotApplicableCount(v int32) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.NotApplicableCount = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetTotalCount(v int32) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.TotalCount = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) Validate() error {
	if s.ConfigRuleCompliances != nil {
		for _, item := range s.ConfigRuleCompliances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances struct {
	// The compliance evaluation result returned by the rule. Valid values:
	//
	// 	- COMPLIANT: The relevant resources are evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The relevant resources are evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to your resources.
	//
	// 	- INSUFFICIENT_DATA: No resource data is available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// cr-fdc8626622af00f9****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test-rule-name
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
}

func (s GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetComplianceType(v string) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetConfigRuleId(v string) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ConfigRuleId = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetConfigRuleName(v string) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ConfigRuleName = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) Validate() error {
	return dara.Validate(s)
}
