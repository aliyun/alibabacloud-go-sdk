// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertEvaluationResultsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *RevertEvaluationResultsShrinkRequest
	GetConfigRuleId() *string
	SetResourcesShrink(v string) *RevertEvaluationResultsShrinkRequest
	GetResourcesShrink() *string
}

type RevertEvaluationResultsShrinkRequest struct {
	// The rule ID.
	//
	// For more information about how to obtain the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-7e72626622af0051****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The resources that are to be re-evaluated.
	//
	// This parameter is required.
	ResourcesShrink *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s RevertEvaluationResultsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevertEvaluationResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevertEvaluationResultsShrinkRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *RevertEvaluationResultsShrinkRequest) GetResourcesShrink() *string {
	return s.ResourcesShrink
}

func (s *RevertEvaluationResultsShrinkRequest) SetConfigRuleId(v string) *RevertEvaluationResultsShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *RevertEvaluationResultsShrinkRequest) SetResourcesShrink(v string) *RevertEvaluationResultsShrinkRequest {
	s.ResourcesShrink = &v
	return s
}

func (s *RevertEvaluationResultsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
