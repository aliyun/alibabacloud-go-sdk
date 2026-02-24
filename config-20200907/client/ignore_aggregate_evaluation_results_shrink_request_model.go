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
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	IgnoreDate   *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
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
