// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreAggregateEvaluationResultsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *IgnoreAggregateEvaluationResultsShrinkRequest
	GetAggregatorId() *string
	SetConfigRuleId(v string) *IgnoreAggregateEvaluationResultsShrinkRequest
	GetConfigRuleId() *string
	SetIgnoreDate(v string) *IgnoreAggregateEvaluationResultsShrinkRequest
	GetIgnoreDate() *string
	SetReason(v string) *IgnoreAggregateEvaluationResultsShrinkRequest
	GetReason() *string
	SetResourcesShrink(v string) *IgnoreAggregateEvaluationResultsShrinkRequest
	GetResourcesShrink() *string
}

type IgnoreAggregateEvaluationResultsShrinkRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-5b6c626622af008f****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the rule.
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-7e72626622af0051****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The date from which the system automatically re-evaluates the ignored incompliant resources.
	//
	// >  If you leave this parameter empty, the system does not automatically re-evaluate the ignored incompliant resources. You must manually re-evaluate the ignored incompliant resources.
	//
	// example:
	//
	// 2022-06-01
	IgnoreDate *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	// The reason why you ignore the resource.
	//
	// example:
	//
	// The reason why you ignore the resource.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The resources to be ignored.
	//
	// This parameter is required.
	ResourcesShrink *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s IgnoreAggregateEvaluationResultsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) GetIgnoreDate() *string {
	return s.IgnoreDate
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) GetResourcesShrink() *string {
	return s.ResourcesShrink
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) SetAggregatorId(v string) *IgnoreAggregateEvaluationResultsShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) SetConfigRuleId(v string) *IgnoreAggregateEvaluationResultsShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) SetIgnoreDate(v string) *IgnoreAggregateEvaluationResultsShrinkRequest {
	s.IgnoreDate = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) SetReason(v string) *IgnoreAggregateEvaluationResultsShrinkRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) SetResourcesShrink(v string) *IgnoreAggregateEvaluationResultsShrinkRequest {
	s.ResourcesShrink = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
