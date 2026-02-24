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
	// The compliance results for the rules in the compliance package.
	ConfigRuleComplianceResult *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult `json:"ConfigRuleComplianceResult,omitempty" xml:"ConfigRuleComplianceResult,omitempty" type:"Struct"`
	// The request ID.
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
	// The number of compliant rules.
	//
	// example:
	//
	// 0
	CompliantCount *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	// The list of rules in the compliance package and their compliance results.
	ConfigRuleCompliances []*GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances `json:"ConfigRuleCompliances,omitempty" xml:"ConfigRuleCompliances,omitempty" type:"Repeated"`
	// The number of ignored rules.
	//
	// example:
	//
	// 1
	IgnoredCount *int32 `json:"IgnoredCount,omitempty" xml:"IgnoredCount,omitempty"`
	// The total number of rules within the compliance package whose evaluation results are "No Data" when assessing resources.
	//
	// example:
	//
	// 1
	InsufficientDataCount *int32 `json:"InsufficientDataCount,omitempty" xml:"InsufficientDataCount,omitempty"`
	// The number of non-compliant rules.
	//
	// example:
	//
	// 0
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	// The number of rules that are not applicable.
	//
	// example:
	//
	// 1
	NotApplicableCount *int32 `json:"NotApplicableCount,omitempty" xml:"NotApplicableCount,omitempty"`
	// The total number of rules in the compliance package.
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
	// The compliance evaluation result of the rule. Valid values:
	//
	// - COMPLIANT: The rule is compliant.
	//
	// - NON_COMPLIANT: The rule is non-compliant.
	//
	// - NOT_APPLICABLE: The rule is not applicable.
	//
	// - INSUFFICIENT_DATA: No data is available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The ID of the rule in the compliance package.
	//
	// example:
	//
	// cr-fdc8626622af00f9****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule in the compliance package.
	//
	// example:
	//
	// The bandwidth of the Elastic IP instance meets the minimum requirements.
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
