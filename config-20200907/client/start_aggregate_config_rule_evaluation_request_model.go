// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAggregateConfigRuleEvaluationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *StartAggregateConfigRuleEvaluationRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *StartAggregateConfigRuleEvaluationRequest
	GetCompliancePackId() *string
	SetConfigRuleId(v string) *StartAggregateConfigRuleEvaluationRequest
	GetConfigRuleId() *string
	SetRevertEvaluation(v bool) *StartAggregateConfigRuleEvaluationRequest
	GetRevertEvaluation() *bool
}

type StartAggregateConfigRuleEvaluationRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-3a58626622af0005****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
	//
	// > You must configure either the `CompliancePackId` or `ConfigRuleId` parameter.
	//
	// example:
	//
	// cp-ac16626622af0053****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The rule ID.
	//
	// For more information about how to query the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// >  You must configure either the `CompliancePackId` or `ConfigRuleId` parameter.
	//
	// example:
	//
	// cr-c169626622af009f****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// Specifies whether to re-evaluate the ignored non-compliant resource. Valid values:
	//
	// 	- true: re-evaluates the ignored non-compliant resource based on the rule.
	//
	// 	- false (default): does not re-evaluate the ignored non-compliant resource based on the rule.
	//
	// example:
	//
	// false
	RevertEvaluation *bool `json:"RevertEvaluation,omitempty" xml:"RevertEvaluation,omitempty"`
}

func (s StartAggregateConfigRuleEvaluationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAggregateConfigRuleEvaluationRequest) GoString() string {
	return s.String()
}

func (s *StartAggregateConfigRuleEvaluationRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *StartAggregateConfigRuleEvaluationRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *StartAggregateConfigRuleEvaluationRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *StartAggregateConfigRuleEvaluationRequest) GetRevertEvaluation() *bool {
	return s.RevertEvaluation
}

func (s *StartAggregateConfigRuleEvaluationRequest) SetAggregatorId(v string) *StartAggregateConfigRuleEvaluationRequest {
	s.AggregatorId = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationRequest) SetCompliancePackId(v string) *StartAggregateConfigRuleEvaluationRequest {
	s.CompliancePackId = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationRequest) SetConfigRuleId(v string) *StartAggregateConfigRuleEvaluationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationRequest) SetRevertEvaluation(v bool) *StartAggregateConfigRuleEvaluationRequest {
	s.RevertEvaluation = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationRequest) Validate() error {
	return dara.Validate(s)
}
