// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleComplianceByPackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleComplianceResult(v *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) *GetConfigRuleComplianceByPackResponseBody
	GetConfigRuleComplianceResult() *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult
	SetRequestId(v string) *GetConfigRuleComplianceByPackResponseBody
	GetRequestId() *string
}

type GetConfigRuleComplianceByPackResponseBody struct {
	// The information about the compliance evaluation results returned.
	ConfigRuleComplianceResult *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult `json:"ConfigRuleComplianceResult,omitempty" xml:"ConfigRuleComplianceResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigRuleComplianceByPackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleComplianceByPackResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigRuleComplianceByPackResponseBody) GetConfigRuleComplianceResult() *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	return s.ConfigRuleComplianceResult
}

func (s *GetConfigRuleComplianceByPackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigRuleComplianceByPackResponseBody) SetConfigRuleComplianceResult(v *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) *GetConfigRuleComplianceByPackResponseBody {
	s.ConfigRuleComplianceResult = v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBody) SetRequestId(v string) *GetConfigRuleComplianceByPackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBody) Validate() error {
	if s.ConfigRuleComplianceResult != nil {
		if err := s.ConfigRuleComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult struct {
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-541e626622af0087****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliantCount   *int32  `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	// The rule enabled in the compliance package and the compliance evaluation result returned by the rule.
	ConfigRuleCompliances []*GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances `json:"ConfigRuleCompliances,omitempty" xml:"ConfigRuleCompliances,omitempty" type:"Repeated"`
	IgnoredCount          *int32                                                                                      `json:"IgnoredCount,omitempty" xml:"IgnoredCount,omitempty"`
	InsufficientDataCount *int32                                                                                      `json:"InsufficientDataCount,omitempty" xml:"InsufficientDataCount,omitempty"`
	// The number of rules against which specific resources are evaluated as non-compliant.
	//
	// example:
	//
	// 0
	NonCompliantCount  *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	NotApplicableCount *int32 `json:"NotApplicableCount,omitempty" xml:"NotApplicableCount,omitempty"`
	// The total number of rules enabled in the compliance package.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GoString() string {
	return s.String()
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetCompliantCount() *int32 {
	return s.CompliantCount
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetConfigRuleCompliances() []*GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	return s.ConfigRuleCompliances
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetIgnoredCount() *int32 {
	return s.IgnoredCount
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetInsufficientDataCount() *int32 {
	return s.InsufficientDataCount
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetNotApplicableCount() *int32 {
	return s.NotApplicableCount
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetCompliancePackId(v string) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.CompliancePackId = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetCompliantCount(v int32) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.CompliantCount = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetConfigRuleCompliances(v []*GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.ConfigRuleCompliances = v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetIgnoredCount(v int32) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.IgnoredCount = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetInsufficientDataCount(v int32) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.InsufficientDataCount = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetNonCompliantCount(v int32) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.NonCompliantCount = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetNotApplicableCount(v int32) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.NotApplicableCount = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetTotalCount(v int32) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.TotalCount = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) Validate() error {
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

type GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances struct {
	// The compliance evaluation result. Valid values:
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
	// The ID of the rule enabled in the compliance package.
	//
	// example:
	//
	// cr-fdc8626622af00f9****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule enabled in the compliance package.
	//
	// example:
	//
	// test-rule-name
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
}

func (s GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) GoString() string {
	return s.String()
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetComplianceType(v string) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetConfigRuleId(v string) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ConfigRuleId = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetConfigRuleName(v string) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ConfigRuleName = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) Validate() error {
	return dara.Validate(s)
}
