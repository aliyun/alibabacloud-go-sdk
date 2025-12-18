// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluatePreConfigRulesShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEnableManagedRules(v bool) *EvaluatePreConfigRulesShrinkRequest
  GetEnableManagedRules() *bool 
  SetResourceEvaluateItemsShrink(v string) *EvaluatePreConfigRulesShrinkRequest
  GetResourceEvaluateItemsShrink() *string 
  SetResourceTypeFormat(v string) *EvaluatePreConfigRulesShrinkRequest
  GetResourceTypeFormat() *string 
}

type EvaluatePreConfigRulesShrinkRequest struct {
  // Specifies whether to enable the managed rule. Valid values:
  // 
  // 	- true: enables the managed rule.
  // 
  // 	- false: does not enable the managed rule. This is the default value.
  // 
  // >  After you create an evaluation rule, a managed rule that has the same settings as the evaluation rule is created. After you create a resource, the managed rule can be used to continuously check the compliance of the resource.
  // 
  // example:
  // 
  // false
  EnableManagedRules *bool `json:"EnableManagedRules,omitempty" xml:"EnableManagedRules,omitempty"`
  // The resources that you want to evaluate.
  // 
  // This parameter is required.
  ResourceEvaluateItemsShrink *string `json:"ResourceEvaluateItems,omitempty" xml:"ResourceEvaluateItems,omitempty"`
  // 下一个查询开始Token
  // 
  // example:
  // 
  // ros
  ResourceTypeFormat *string `json:"ResourceTypeFormat,omitempty" xml:"ResourceTypeFormat,omitempty"`
}

func (s EvaluatePreConfigRulesShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EvaluatePreConfigRulesShrinkRequest) GoString() string {
  return s.String()
}

func (s *EvaluatePreConfigRulesShrinkRequest) GetEnableManagedRules() *bool  {
  return s.EnableManagedRules
}

func (s *EvaluatePreConfigRulesShrinkRequest) GetResourceEvaluateItemsShrink() *string  {
  return s.ResourceEvaluateItemsShrink
}

func (s *EvaluatePreConfigRulesShrinkRequest) GetResourceTypeFormat() *string  {
  return s.ResourceTypeFormat
}

func (s *EvaluatePreConfigRulesShrinkRequest) SetEnableManagedRules(v bool) *EvaluatePreConfigRulesShrinkRequest {
  s.EnableManagedRules = &v
  return s
}

func (s *EvaluatePreConfigRulesShrinkRequest) SetResourceEvaluateItemsShrink(v string) *EvaluatePreConfigRulesShrinkRequest {
  s.ResourceEvaluateItemsShrink = &v
  return s
}

func (s *EvaluatePreConfigRulesShrinkRequest) SetResourceTypeFormat(v string) *EvaluatePreConfigRulesShrinkRequest {
  s.ResourceTypeFormat = &v
  return s
}

func (s *EvaluatePreConfigRulesShrinkRequest) Validate() error {
  return dara.Validate(s)
}

