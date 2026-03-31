// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComplianceSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceSummary(v *GetComplianceSummaryResponseBodyComplianceSummary) *GetComplianceSummaryResponseBody
	GetComplianceSummary() *GetComplianceSummaryResponseBodyComplianceSummary
	SetRequestId(v string) *GetComplianceSummaryResponseBody
	GetRequestId() *string
}

type GetComplianceSummaryResponseBody struct {
	// The summary of compliance statistics.
	ComplianceSummary *GetComplianceSummaryResponseBodyComplianceSummary `json:"ComplianceSummary,omitempty" xml:"ComplianceSummary,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CAEE6F34-DEDC-4BAA-AA8C-946D5D008737
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetComplianceSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComplianceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetComplianceSummaryResponseBody) GetComplianceSummary() *GetComplianceSummaryResponseBodyComplianceSummary {
	return s.ComplianceSummary
}

func (s *GetComplianceSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComplianceSummaryResponseBody) SetComplianceSummary(v *GetComplianceSummaryResponseBodyComplianceSummary) *GetComplianceSummaryResponseBody {
	s.ComplianceSummary = v
	return s
}

func (s *GetComplianceSummaryResponseBody) SetRequestId(v string) *GetComplianceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComplianceSummaryResponseBody) Validate() error {
	if s.ComplianceSummary != nil {
		if err := s.ComplianceSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComplianceSummaryResponseBodyComplianceSummary struct {
	// The summary of compliance statistics from the rule dimension.
	ComplianceSummaryByConfigRule *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule `json:"ComplianceSummaryByConfigRule,omitempty" xml:"ComplianceSummaryByConfigRule,omitempty" type:"Struct"`
	// The summary of compliance statistics from the resource dimension.
	ComplianceSummaryByResource *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource `json:"ComplianceSummaryByResource,omitempty" xml:"ComplianceSummaryByResource,omitempty" type:"Struct"`
}

func (s GetComplianceSummaryResponseBodyComplianceSummary) String() string {
	return dara.Prettify(s)
}

func (s GetComplianceSummaryResponseBodyComplianceSummary) GoString() string {
	return s.String()
}

func (s *GetComplianceSummaryResponseBodyComplianceSummary) GetComplianceSummaryByConfigRule() *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	return s.ComplianceSummaryByConfigRule
}

func (s *GetComplianceSummaryResponseBodyComplianceSummary) GetComplianceSummaryByResource() *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	return s.ComplianceSummaryByResource
}

func (s *GetComplianceSummaryResponseBodyComplianceSummary) SetComplianceSummaryByConfigRule(v *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) *GetComplianceSummaryResponseBodyComplianceSummary {
	s.ComplianceSummaryByConfigRule = v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummary) SetComplianceSummaryByResource(v *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) *GetComplianceSummaryResponseBodyComplianceSummary {
	s.ComplianceSummaryByResource = v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummary) Validate() error {
	if s.ComplianceSummaryByConfigRule != nil {
		if err := s.ComplianceSummaryByConfigRule.Validate(); err != nil {
			return err
		}
	}
	if s.ComplianceSummaryByResource != nil {
		if err := s.ComplianceSummaryByResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule struct {
	// The time when the compliance summary was generated. Unit: milliseconds.
	//
	// example:
	//
	// 1589853712165
	ComplianceSummaryTimestamp *int64 `json:"ComplianceSummaryTimestamp,omitempty" xml:"ComplianceSummaryTimestamp,omitempty"`
	// The number of rules evaluated as compliant.
	//
	// example:
	//
	// 5
	CompliantCount *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	// The number of rules evaluated as non-compliant.
	//
	// example:
	//
	// 11
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 16
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) String() string {
	return dara.Prettify(s)
}

func (s GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GoString() string {
	return s.String()
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GetComplianceSummaryTimestamp() *int64 {
	return s.ComplianceSummaryTimestamp
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GetCompliantCount() *int32 {
	return s.CompliantCount
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetComplianceSummaryTimestamp(v int64) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.ComplianceSummaryTimestamp = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetCompliantCount(v int32) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.CompliantCount = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetNonCompliantCount(v int32) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.NonCompliantCount = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetTotalCount(v int64) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.TotalCount = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) Validate() error {
	return dara.Validate(s)
}

type GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource struct {
	// The time when the compliance summary was generated. Unit: milliseconds.
	//
	// example:
	//
	// 1589853712165
	ComplianceSummaryTimestamp *int64 `json:"ComplianceSummaryTimestamp,omitempty" xml:"ComplianceSummaryTimestamp,omitempty"`
	// The number of resources evaluated as compliant.
	//
	// example:
	//
	// 1
	CompliantCount                          *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	HighRiskRuleNonCompliantResourceCount   *int32 `json:"HighRiskRuleNonCompliantResourceCount,omitempty" xml:"HighRiskRuleNonCompliantResourceCount,omitempty"`
	LowRiskRuleNonCompliantResourceCount    *int32 `json:"LowRiskRuleNonCompliantResourceCount,omitempty" xml:"LowRiskRuleNonCompliantResourceCount,omitempty"`
	MediumRiskRuleNonCompliantResourceCount *int32 `json:"MediumRiskRuleNonCompliantResourceCount,omitempty" xml:"MediumRiskRuleNonCompliantResourceCount,omitempty"`
	// The number of resources evaluated as non-compliant.
	//
	// example:
	//
	// 12
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	// The total number of resources.
	//
	// example:
	//
	// 13
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) String() string {
	return dara.Prettify(s)
}

func (s GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GoString() string {
	return s.String()
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetComplianceSummaryTimestamp() *int64 {
	return s.ComplianceSummaryTimestamp
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetCompliantCount() *int32 {
	return s.CompliantCount
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetHighRiskRuleNonCompliantResourceCount() *int32 {
	return s.HighRiskRuleNonCompliantResourceCount
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetLowRiskRuleNonCompliantResourceCount() *int32 {
	return s.LowRiskRuleNonCompliantResourceCount
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetMediumRiskRuleNonCompliantResourceCount() *int32 {
	return s.MediumRiskRuleNonCompliantResourceCount
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetComplianceSummaryTimestamp(v int64) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.ComplianceSummaryTimestamp = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetCompliantCount(v int32) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.CompliantCount = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetHighRiskRuleNonCompliantResourceCount(v int32) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.HighRiskRuleNonCompliantResourceCount = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetLowRiskRuleNonCompliantResourceCount(v int32) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.LowRiskRuleNonCompliantResourceCount = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetMediumRiskRuleNonCompliantResourceCount(v int32) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.MediumRiskRuleNonCompliantResourceCount = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetNonCompliantCount(v int32) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.NonCompliantCount = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetTotalCount(v int64) *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.TotalCount = &v
	return s
}

func (s *GetComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) Validate() error {
	return dara.Validate(s)
}
