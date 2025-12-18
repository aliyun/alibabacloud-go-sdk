// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluatePreConfigRulesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EvaluatePreConfigRulesResponseBody
  GetRequestId() *string 
  SetResourceEvaluations(v []*EvaluatePreConfigRulesResponseBodyResourceEvaluations) *EvaluatePreConfigRulesResponseBody
  GetResourceEvaluations() []*EvaluatePreConfigRulesResponseBodyResourceEvaluations 
}

type EvaluatePreConfigRulesResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 129ECF1C-7897-1131-BD0F-4B588AC05400
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The details of the compliance evaluation result.
  ResourceEvaluations []*EvaluatePreConfigRulesResponseBodyResourceEvaluations `json:"ResourceEvaluations,omitempty" xml:"ResourceEvaluations,omitempty" type:"Repeated"`
}

func (s EvaluatePreConfigRulesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EvaluatePreConfigRulesResponseBody) GoString() string {
  return s.String()
}

func (s *EvaluatePreConfigRulesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EvaluatePreConfigRulesResponseBody) GetResourceEvaluations() []*EvaluatePreConfigRulesResponseBodyResourceEvaluations  {
  return s.ResourceEvaluations
}

func (s *EvaluatePreConfigRulesResponseBody) SetRequestId(v string) *EvaluatePreConfigRulesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EvaluatePreConfigRulesResponseBody) SetResourceEvaluations(v []*EvaluatePreConfigRulesResponseBodyResourceEvaluations) *EvaluatePreConfigRulesResponseBody {
  s.ResourceEvaluations = v
  return s
}

func (s *EvaluatePreConfigRulesResponseBody) Validate() error {
  if s.ResourceEvaluations != nil {
    for _, item := range s.ResourceEvaluations {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EvaluatePreConfigRulesResponseBodyResourceEvaluations struct {
  // The logical ID of the resource.
  // 
  // >  If the ResourceLogicalId request parameter is left empty, the value of the ResourceLogicalId response parameter is generated based on the value of the `ResourceProperties` parameter.
  // 
  // example:
  // 
  // ResourceLogicId-test
  ResourceLogicalId *string `json:"ResourceLogicalId,omitempty" xml:"ResourceLogicalId,omitempty"`
  // The type of the resource.
  // 
  // example:
  // 
  // ACS::ECS::Instance
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  // The evaluation rules.
  Rules []*EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s EvaluatePreConfigRulesResponseBodyResourceEvaluations) String() string {
  return dara.Prettify(s)
}

func (s EvaluatePreConfigRulesResponseBodyResourceEvaluations) GoString() string {
  return s.String()
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluations) GetResourceLogicalId() *string  {
  return s.ResourceLogicalId
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluations) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluations) GetRules() []*EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules  {
  return s.Rules
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluations) SetResourceLogicalId(v string) *EvaluatePreConfigRulesResponseBodyResourceEvaluations {
  s.ResourceLogicalId = &v
  return s
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluations) SetResourceType(v string) *EvaluatePreConfigRulesResponseBodyResourceEvaluations {
  s.ResourceType = &v
  return s
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluations) SetRules(v []*EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) *EvaluatePreConfigRulesResponseBodyResourceEvaluations {
  s.Rules = v
  return s
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluations) Validate() error {
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

type EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules struct {
  // The reason why the resource was evaluated as incompliant.
  // 
  // example:
  // 
  // {\\"configuration\\":\\"false\\",\\"desiredValue\\":\\"True\\",\\"operator\\":\\"StringEquals\\",\\"property\\":\\"$.DeletionProtection\\"}
  Annotation *string `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
  // The compliance type of the resource that was evaluated by using the evaluation rule. Valid values:
  // 
  // 	- COMPLIANT: The resource was evaluated as compliant.
  // 
  // 	- NON_COMPLIANT: The resource was evaluated as incompliant.
  // 
  // 	- NOT_APPLICABLE: The evaluation rule does not apply to the resource.
  // 
  // example:
  // 
  // NON_COMPLIANT
  ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
  // The URL of the topic that describes how the managed rule remediates the incompliant configurations.
  // 
  // example:
  // 
  // https://example.aliyundoc.com
  HelpUrl *string `json:"HelpUrl,omitempty" xml:"HelpUrl,omitempty"`
  // The identifier of the evaluation rule.
  // 
  // example:
  // 
  // ecs-instance-deletion-protection-enabled
  Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) String() string {
  return dara.Prettify(s)
}

func (s EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) GoString() string {
  return s.String()
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) GetAnnotation() *string  {
  return s.Annotation
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) GetComplianceType() *string  {
  return s.ComplianceType
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) GetHelpUrl() *string  {
  return s.HelpUrl
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) GetIdentifier() *string  {
  return s.Identifier
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) SetAnnotation(v string) *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules {
  s.Annotation = &v
  return s
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) SetComplianceType(v string) *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules {
  s.ComplianceType = &v
  return s
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) SetHelpUrl(v string) *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules {
  s.HelpUrl = &v
  return s
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) SetIdentifier(v string) *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules {
  s.Identifier = &v
  return s
}

func (s *EvaluatePreConfigRulesResponseBodyResourceEvaluationsRules) Validate() error {
  return dara.Validate(s)
}

