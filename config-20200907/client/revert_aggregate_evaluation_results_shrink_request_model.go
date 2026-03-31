// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertAggregateEvaluationResultsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *RevertAggregateEvaluationResultsShrinkRequest
	GetAggregatorId() *string
	SetConfigRuleId(v string) *RevertAggregateEvaluationResultsShrinkRequest
	GetConfigRuleId() *string
	SetResourcesShrink(v string) *RevertAggregateEvaluationResultsShrinkRequest
	GetResourcesShrink() *string
}

type RevertAggregateEvaluationResultsShrinkRequest struct {
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
	// The ID of the rule in the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-7e72626622af0051****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The resources that you want to re-evaluate.
	//
	// This parameter is required.
	ResourcesShrink *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s RevertAggregateEvaluationResultsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevertAggregateEvaluationResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevertAggregateEvaluationResultsShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *RevertAggregateEvaluationResultsShrinkRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *RevertAggregateEvaluationResultsShrinkRequest) GetResourcesShrink() *string {
	return s.ResourcesShrink
}

func (s *RevertAggregateEvaluationResultsShrinkRequest) SetAggregatorId(v string) *RevertAggregateEvaluationResultsShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsShrinkRequest) SetConfigRuleId(v string) *RevertAggregateEvaluationResultsShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsShrinkRequest) SetResourcesShrink(v string) *RevertAggregateEvaluationResultsShrinkRequest {
	s.ResourcesShrink = &v
	return s
}

func (s *RevertAggregateEvaluationResultsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
