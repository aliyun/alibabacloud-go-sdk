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
	// This parameter is required.
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleId     *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	RevertEvaluation *bool   `json:"RevertEvaluation,omitempty" xml:"RevertEvaluation,omitempty"`
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
