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
  ResourceEvaluateItems []*EvaluatePreConfigRulesRequestResourceEvaluateItems `json:"ResourceEvaluateItems,omitempty" xml:"ResourceEvaluateItems,omitempty" type:"Repeated"`
  // 下一个查询开始Token
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
  // example:
  // 
  // ResourceLogicId-test
  ResourceLogicalId *string `json:"ResourceLogicalId,omitempty" xml:"ResourceLogicalId,omitempty"`
  // The properties of the resource.
  // 
  // example:
  // 
  // {
  // 
  //     "ImageId": "ubuntu_18_04_64_20G_alibase_20190624.vhd",
  // 
  //     "SecurityGroupId": "sg-bp15ed6xe1yxeycg****",
  // 
  //     "HostName": "LocalHostName",
  // 
  //     "RegionId": "cn-hangzhou"
  // 
  // }
  ResourceProperties *string `json:"ResourceProperties,omitempty" xml:"ResourceProperties,omitempty"`
  // The type of the resource.
  // 
  // example:
  // 
  // ACS::ECS::Instance
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  // The evaluation rules.
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
  // The identifier of the evaluation rule.
  // 
  // For more information about how to obtain the identifier of an evaluation rule, see [ListManagedRules](https://help.aliyun.com/document_detail/467810.html).
  // 
  // example:
  // 
  // ecs-instance-deletion-protection-enabled
  Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
  // The input parameters of the evaluation rule.
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

