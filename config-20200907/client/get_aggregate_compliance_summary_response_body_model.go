// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateComplianceSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceSummary(v *GetAggregateComplianceSummaryResponseBodyComplianceSummary) *GetAggregateComplianceSummaryResponseBody
	GetComplianceSummary() *GetAggregateComplianceSummaryResponseBodyComplianceSummary
	SetRequestId(v string) *GetAggregateComplianceSummaryResponseBody
	GetRequestId() *string
}

type GetAggregateComplianceSummaryResponseBody struct {
	// The compliance statistics.
	ComplianceSummary *GetAggregateComplianceSummaryResponseBodyComplianceSummary `json:"ComplianceSummary,omitempty" xml:"ComplianceSummary,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 929B8360-BD57-54FF-96DB-AD1D9B476769
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateComplianceSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateComplianceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateComplianceSummaryResponseBody) GetComplianceSummary() *GetAggregateComplianceSummaryResponseBodyComplianceSummary {
	return s.ComplianceSummary
}

func (s *GetAggregateComplianceSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateComplianceSummaryResponseBody) SetComplianceSummary(v *GetAggregateComplianceSummaryResponseBodyComplianceSummary) *GetAggregateComplianceSummaryResponseBody {
	s.ComplianceSummary = v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBody) SetRequestId(v string) *GetAggregateComplianceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBody) Validate() error {
	if s.ComplianceSummary != nil {
		if err := s.ComplianceSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateComplianceSummaryResponseBodyComplianceSummary struct {
	// The summary of compliance statistics from the rule dimension.
	ComplianceSummaryByConfigRule *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule `json:"ComplianceSummaryByConfigRule,omitempty" xml:"ComplianceSummaryByConfigRule,omitempty" type:"Struct"`
	// The summary of compliance statistics from the resource dimension.
	ComplianceSummaryByResource *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource `json:"ComplianceSummaryByResource,omitempty" xml:"ComplianceSummaryByResource,omitempty" type:"Struct"`
}

func (s GetAggregateComplianceSummaryResponseBodyComplianceSummary) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateComplianceSummaryResponseBodyComplianceSummary) GoString() string {
	return s.String()
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummary) GetComplianceSummaryByConfigRule() *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	return s.ComplianceSummaryByConfigRule
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummary) GetComplianceSummaryByResource() *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	return s.ComplianceSummaryByResource
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummary) SetComplianceSummaryByConfigRule(v *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) *GetAggregateComplianceSummaryResponseBodyComplianceSummary {
	s.ComplianceSummaryByConfigRule = v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummary) SetComplianceSummaryByResource(v *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) *GetAggregateComplianceSummaryResponseBodyComplianceSummary {
	s.ComplianceSummaryByResource = v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummary) Validate() error {
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

type GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule struct {
	// The time when the compliance statistics were collected. Unit: milliseconds.
	//
	// example:
	//
	// 1589853822103
	ComplianceSummaryTimestamp *int64 `json:"ComplianceSummaryTimestamp,omitempty" xml:"ComplianceSummaryTimestamp,omitempty"`
	// The number of compliant rules.
	//
	// example:
	//
	// 4
	CompliantCount *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	// The number of non-compliant rules.
	//
	// example:
	//
	// 5
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 9
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GoString() string {
	return s.String()
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GetComplianceSummaryTimestamp() *int64 {
	return s.ComplianceSummaryTimestamp
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GetCompliantCount() *int32 {
	return s.CompliantCount
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetComplianceSummaryTimestamp(v int64) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.ComplianceSummaryTimestamp = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetCompliantCount(v int32) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.CompliantCount = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetNonCompliantCount(v int32) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.NonCompliantCount = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetTotalCount(v int64) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.TotalCount = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) Validate() error {
	return dara.Validate(s)
}

type GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource struct {
	// The time when the compliance statistics were collected. Unit: milliseconds.
	//
	// example:
	//
	// 1589853822103
	ComplianceSummaryTimestamp *int64 `json:"ComplianceSummaryTimestamp,omitempty" xml:"ComplianceSummaryTimestamp,omitempty"`
	// The number of compliant resources.
	//
	// example:
	//
	// 20
	CompliantCount                          *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	HighRiskRuleNonCompliantResourceCount   *int32 `json:"HighRiskRuleNonCompliantResourceCount,omitempty" xml:"HighRiskRuleNonCompliantResourceCount,omitempty"`
	LowRiskRuleNonCompliantResourceCount    *int32 `json:"LowRiskRuleNonCompliantResourceCount,omitempty" xml:"LowRiskRuleNonCompliantResourceCount,omitempty"`
	MediumRiskRuleNonCompliantResourceCount *int32 `json:"MediumRiskRuleNonCompliantResourceCount,omitempty" xml:"MediumRiskRuleNonCompliantResourceCount,omitempty"`
	// The number of non-compliant resources.
	//
	// example:
	//
	// 11
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	// The total number of resources.
	//
	// example:
	//
	// 31
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GoString() string {
	return s.String()
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetComplianceSummaryTimestamp() *int64 {
	return s.ComplianceSummaryTimestamp
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetCompliantCount() *int32 {
	return s.CompliantCount
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetHighRiskRuleNonCompliantResourceCount() *int32 {
	return s.HighRiskRuleNonCompliantResourceCount
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetLowRiskRuleNonCompliantResourceCount() *int32 {
	return s.LowRiskRuleNonCompliantResourceCount
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetMediumRiskRuleNonCompliantResourceCount() *int32 {
	return s.MediumRiskRuleNonCompliantResourceCount
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetComplianceSummaryTimestamp(v int64) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.ComplianceSummaryTimestamp = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetCompliantCount(v int32) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.CompliantCount = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetHighRiskRuleNonCompliantResourceCount(v int32) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.HighRiskRuleNonCompliantResourceCount = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetLowRiskRuleNonCompliantResourceCount(v int32) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.LowRiskRuleNonCompliantResourceCount = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetMediumRiskRuleNonCompliantResourceCount(v int32) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.MediumRiskRuleNonCompliantResourceCount = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetNonCompliantCount(v int32) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.NonCompliantCount = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetTotalCount(v int64) *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.TotalCount = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) Validate() error {
	return dara.Validate(s)
}
