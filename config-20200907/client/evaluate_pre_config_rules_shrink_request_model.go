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
  // Specifies whether to enable rule templates. Valid values:
  // 
  // - true: enables rule templates.
  // 
  // - false (default): does not enable rule templates.
  // 
  // example:
  // 
  // false
  EnableManagedRules *bool `json:"EnableManagedRules,omitempty" xml:"EnableManagedRules,omitempty"`
  // An array that contains the resources that you want to evaluate.
  // 
  // This parameter is required.
  ResourceEvaluateItemsShrink *string `json:"ResourceEvaluateItems,omitempty" xml:"ResourceEvaluateItems,omitempty"`
  // The query start token
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

