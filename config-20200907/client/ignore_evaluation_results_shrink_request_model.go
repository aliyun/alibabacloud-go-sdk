// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreEvaluationResultsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *IgnoreEvaluationResultsShrinkRequest
	GetConfigRuleId() *string
	SetIgnoreDate(v string) *IgnoreEvaluationResultsShrinkRequest
	GetIgnoreDate() *string
	SetReason(v string) *IgnoreEvaluationResultsShrinkRequest
	GetReason() *string
	SetResourcesShrink(v string) *IgnoreEvaluationResultsShrinkRequest
	GetResourcesShrink() *string
}

type IgnoreEvaluationResultsShrinkRequest struct {
	// The rule ID.
	//
	// For more information about how to obtain the rule ID, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-7e72626622af0051****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The date on which the ignored evaluation results are automatically restored.
	//
	// > If this parameter is left empty, the ignored evaluation results are not automatically restored. You must manually restore them.
	//
	// example:
	//
	// 2022-06-01
	IgnoreDate *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	// The reason for ignoring the resources.
	//
	// example:
	//
	// 无需检测。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The list of resources to be ignored.
	//
	// This parameter is required.
	ResourcesShrink *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s IgnoreEvaluationResultsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IgnoreEvaluationResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *IgnoreEvaluationResultsShrinkRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *IgnoreEvaluationResultsShrinkRequest) GetIgnoreDate() *string {
	return s.IgnoreDate
}

func (s *IgnoreEvaluationResultsShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *IgnoreEvaluationResultsShrinkRequest) GetResourcesShrink() *string {
	return s.ResourcesShrink
}

func (s *IgnoreEvaluationResultsShrinkRequest) SetConfigRuleId(v string) *IgnoreEvaluationResultsShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *IgnoreEvaluationResultsShrinkRequest) SetIgnoreDate(v string) *IgnoreEvaluationResultsShrinkRequest {
	s.IgnoreDate = &v
	return s
}

func (s *IgnoreEvaluationResultsShrinkRequest) SetReason(v string) *IgnoreEvaluationResultsShrinkRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreEvaluationResultsShrinkRequest) SetResourcesShrink(v string) *IgnoreEvaluationResultsShrinkRequest {
	s.ResourcesShrink = &v
	return s
}

func (s *IgnoreEvaluationResultsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
