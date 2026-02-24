// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluatePreConfigRulesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEnableManagedRules(v bool) *EvaluatePreConfigRulesRequest
  GetEnableManagedRules() *bool 
  SetResourceEvaluateItems(v []*EvaluatePreConfigRulesRequestResourceEvaluateItems) *EvaluatePreConfigRulesRequest
  GetResourceEvaluateItems() []*EvaluatePreConfigRulesRequestResourceEvaluateItems 
  SetResourceTypeFormat(v string) *EvaluatePreConfigRulesRequest
  GetResourceTypeFormat() *string 
}

type EvaluatePreConfigRulesRequest struct {
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
  ResourceEvaluateItems []*EvaluatePreConfigRulesRequestResourceEvaluateItems `json:"ResourceEvaluateItems,omitempty" xml:"ResourceEvaluateItems,omitempty" type:"Repeated"`
  // The query start token
  // 
  // example:
  // 
  // ros
  ResourceTypeFormat *string `json:"ResourceTypeFormat,omitempty" xml:"ResourceTypeFormat,omitempty"`
}

func (s EvaluatePreConfigRulesRequest) String() string {
  return dara.Prettify(s)
}

func (s EvaluatePreConfigRulesRequest) GoString() string {
  return s.String()
}

func (s *EvaluatePreConfigRulesRequest) GetEnableManagedRules() *bool  {
  return s.EnableManagedRules
}

func (s *EvaluatePreConfigRulesRequest) GetResourceEvaluateItems() []*EvaluatePreConfigRulesRequestResourceEvaluateItems  {
  return s.ResourceEvaluateItems
}

func (s *EvaluatePreConfigRulesRequest) GetResourceTypeFormat() *string  {
  return s.ResourceTypeFormat
}

func (s *EvaluatePreConfigRulesRequest) SetEnableManagedRules(v bool) *EvaluatePreConfigRulesRequest {
  s.EnableManagedRules = &v
  return s
}

func (s *EvaluatePreConfigRulesRequest) SetResourceEvaluateItems(v []*EvaluatePreConfigRulesRequestResourceEvaluateItems) *EvaluatePreConfigRulesRequest {
  s.ResourceEvaluateItems = v
  return s
}

func (s *EvaluatePreConfigRulesRequest) SetResourceTypeFormat(v string) *EvaluatePreConfigRulesRequest {
  s.ResourceTypeFormat = &v
  return s
}

func (s *EvaluatePreConfigRulesRequest) Validate() error {
  if s.ResourceEvaluateItems != nil {
    for _, item := range s.ResourceEvaluateItems {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EvaluatePreConfigRulesRequestResourceEvaluateItems struct {
  // The logical ID of the resource.
  // 
  // > If this parameter is empty, it is generated based on the Base64 value of `ResourceProperties`.
  // 
  // example:
  // 
  // ResourceLogicId-test
  ResourceLogicalId *string `json:"ResourceLogicalId,omitempty" xml:"ResourceLogicalId,omitempty"`
  // The resource configuration items (properties of the resource to be created), such as the specifications, region, name, status, and port or network interface switch status of the resource.
  // 
  // > The `ResourceType`, `Identifier`, and `ResourceProperties` parameters must be specified at the same time.
  // 
  // example:
  // 
  // {"ResourceGroupId":"","Memory":8192,"InstanceChargeType":"PostPaid","Cpu":2}
  ResourceProperties *string `json:"ResourceProperties,omitempty" xml:"ResourceProperties,omitempty"`
  // The type of the resource.
  // 
  // For information about how to obtain the identifier of an evaluation rule, see [ListPreManagedRules](https://help.aliyun.com/document_detail/467810.html).
  // 
  // > The `ResourceType`, `Identifier`, and `ResourceProperties` parameters must be specified at the same time.
  // 
  // example:
  // 
  // ACS::ECS::Instance
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  // An array that contains the evaluation rules.
  Rules []*EvaluatePreConfigRulesRequestResourceEvaluateItemsRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s EvaluatePreConfigRulesRequestResourceEvaluateItems) String() string {
  return dara.Prettify(s)
}

func (s EvaluatePreConfigRulesRequestResourceEvaluateItems) GoString() string {
  return s.String()
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItems) GetResourceLogicalId() *string  {
  return s.ResourceLogicalId
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItems) GetResourceProperties() *string  {
  return s.ResourceProperties
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItems) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItems) GetRules() []*EvaluatePreConfigRulesRequestResourceEvaluateItemsRules  {
  return s.Rules
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItems) SetResourceLogicalId(v string) *EvaluatePreConfigRulesRequestResourceEvaluateItems {
  s.ResourceLogicalId = &v
  return s
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItems) SetResourceProperties(v string) *EvaluatePreConfigRulesRequestResourceEvaluateItems {
  s.ResourceProperties = &v
  return s
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItems) SetResourceType(v string) *EvaluatePreConfigRulesRequestResourceEvaluateItems {
  s.ResourceType = &v
  return s
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItems) SetRules(v []*EvaluatePreConfigRulesRequestResourceEvaluateItemsRules) *EvaluatePreConfigRulesRequestResourceEvaluateItems {
  s.Rules = v
  return s
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItems) Validate() error {
  if s.Rules != nil {
    for _, item := range s.Rules {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EvaluatePreConfigRulesRequestResourceEvaluateItemsRules struct {
  // The identifier of the rule.
  // 
  // For information about how to obtain the identifier of a rule, see [ListPreManagedRules](https://help.aliyun.com/document_detail/467810.html).
  // 
  // > The `ResourceType`, `Identifier`, and `ResourceProperties` parameters must be specified at the same time.
  // 
  // example:
  // 
  // ecs-instance-deletion-protection-enabled
  Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
  // The input parameters of the rule.
  // 
  // example:
  // 
  // {}
  InputParameters *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
}

func (s EvaluatePreConfigRulesRequestResourceEvaluateItemsRules) String() string {
  return dara.Prettify(s)
}

func (s EvaluatePreConfigRulesRequestResourceEvaluateItemsRules) GoString() string {
  return s.String()
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItemsRules) GetIdentifier() *string  {
  return s.Identifier
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItemsRules) GetInputParameters() *string  {
  return s.InputParameters
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItemsRules) SetIdentifier(v string) *EvaluatePreConfigRulesRequestResourceEvaluateItemsRules {
  s.Identifier = &v
  return s
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItemsRules) SetInputParameters(v string) *EvaluatePreConfigRulesRequestResourceEvaluateItemsRules {
  s.InputParameters = &v
  return s
}

func (s *EvaluatePreConfigRulesRequestResourceEvaluateItemsRules) Validate() error {
  return dara.Validate(s)
}

