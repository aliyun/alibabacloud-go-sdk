// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConfigRuleEvaluationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *StartConfigRuleEvaluationRequest
	GetCompliancePackId() *string
	SetConfigRuleId(v string) *StartConfigRuleEvaluationRequest
	GetConfigRuleId() *string
	SetRevertEvaluation(v bool) *StartConfigRuleEvaluationRequest
	GetRevertEvaluation() *bool
}

type StartConfigRuleEvaluationRequest struct {
	// The ID of the compliance package.
	//
	// For more information, see [ListCompliancePacks](https://help.aliyun.com/document_detail/606968.html).
	//
	// > Specify either `CompliancePackId` or `ConfigRuleId`.
	//
	// example:
	//
	// cp-ac16626622af0053****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The ID of the rule.
	//
	// For more information, see [ListConfigRules](https://help.aliyun.com/document_detail/609222.html).
	//
	// > Specify either `CompliancePackId` or `ConfigRuleId`.
	//
	// example:
	//
	// cr-9920626622af0035****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// Specifies whether to restore evaluation results that were ignored. Valid values:
	//
	// - true: Restores ignored evaluation results.
	//
	// - false (default): Does not restore ignored evaluation results.
	//
	// example:
	//
	// false
	RevertEvaluation *bool `json:"RevertEvaluation,omitempty" xml:"RevertEvaluation,omitempty"`
}

func (s StartConfigRuleEvaluationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartConfigRuleEvaluationRequest) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *StartConfigRuleEvaluationRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *StartConfigRuleEvaluationRequest) GetRevertEvaluation() *bool {
	return s.RevertEvaluation
}

func (s *StartConfigRuleEvaluationRequest) SetCompliancePackId(v string) *StartConfigRuleEvaluationRequest {
	s.CompliancePackId = &v
	return s
}

func (s *StartConfigRuleEvaluationRequest) SetConfigRuleId(v string) *StartConfigRuleEvaluationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *StartConfigRuleEvaluationRequest) SetRevertEvaluation(v bool) *StartConfigRuleEvaluationRequest {
	s.RevertEvaluation = &v
	return s
}

func (s *StartConfigRuleEvaluationRequest) Validate() error {
	return dara.Validate(s)
}
