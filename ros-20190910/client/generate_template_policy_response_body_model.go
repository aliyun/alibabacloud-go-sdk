// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTemplatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *GenerateTemplatePolicyResponseBodyPolicy) *GenerateTemplatePolicyResponseBody
	GetPolicy() *GenerateTemplatePolicyResponseBodyPolicy
	SetPolicyFunctions(v []*GenerateTemplatePolicyResponseBodyPolicyFunctions) *GenerateTemplatePolicyResponseBody
	GetPolicyFunctions() []*GenerateTemplatePolicyResponseBodyPolicyFunctions
	SetRequestId(v string) *GenerateTemplatePolicyResponseBody
	GetRequestId() *string
}

type GenerateTemplatePolicyResponseBody struct {
	// The information about the policy.
	Policy          *GenerateTemplatePolicyResponseBodyPolicy            `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PolicyFunctions []*GenerateTemplatePolicyResponseBodyPolicyFunctions `json:"PolicyFunctions,omitempty" xml:"PolicyFunctions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateTemplatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBody) GetPolicy() *GenerateTemplatePolicyResponseBodyPolicy {
	return s.Policy
}

func (s *GenerateTemplatePolicyResponseBody) GetPolicyFunctions() []*GenerateTemplatePolicyResponseBodyPolicyFunctions {
	return s.PolicyFunctions
}

func (s *GenerateTemplatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateTemplatePolicyResponseBody) SetPolicy(v *GenerateTemplatePolicyResponseBodyPolicy) *GenerateTemplatePolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GenerateTemplatePolicyResponseBody) SetPolicyFunctions(v []*GenerateTemplatePolicyResponseBodyPolicyFunctions) *GenerateTemplatePolicyResponseBody {
	s.PolicyFunctions = v
	return s
}

func (s *GenerateTemplatePolicyResponseBody) SetRequestId(v string) *GenerateTemplatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	if s.PolicyFunctions != nil {
		for _, item := range s.PolicyFunctions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateTemplatePolicyResponseBodyPolicy struct {
	// The statements that are contained in the policy.
	Statement []*GenerateTemplatePolicyResponseBodyPolicyStatement `json:"Statement,omitempty" xml:"Statement,omitempty" type:"Repeated"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GenerateTemplatePolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBodyPolicy) GetStatement() []*GenerateTemplatePolicyResponseBodyPolicyStatement {
	return s.Statement
}

func (s *GenerateTemplatePolicyResponseBodyPolicy) GetVersion() *string {
	return s.Version
}

func (s *GenerateTemplatePolicyResponseBodyPolicy) SetStatement(v []*GenerateTemplatePolicyResponseBodyPolicyStatement) *GenerateTemplatePolicyResponseBodyPolicy {
	s.Statement = v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicy) SetVersion(v string) *GenerateTemplatePolicyResponseBodyPolicy {
	s.Version = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicy) Validate() error {
	if s.Statement != nil {
		for _, item := range s.Statement {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateTemplatePolicyResponseBodyPolicyStatement struct {
	// The operations that are performed on the specified resource.
	Action []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
	// The condition that is required for the policy to take effect.
	//
	// example:
	//
	// {
	//
	//     "StringEquals": {
	//
	//         "acs:Service": "fc.aliyuncs.com"
	//
	//     }
	//
	// }
	Condition map[string]interface{} `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The effect of the statement. Valid values:
	//
	// 	- Allow
	//
	// 	- Deny
	//
	// example:
	//
	// Allow
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The objects that the statement covers. An asterisk (\\*) indicates all resources.
	//
	// example:
	//
	// *
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s GenerateTemplatePolicyResponseBodyPolicyStatement) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBodyPolicyStatement) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) GetAction() []*string {
	return s.Action
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) GetCondition() map[string]interface{} {
	return s.Condition
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) GetEffect() *string {
	return s.Effect
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) GetResource() *string {
	return s.Resource
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetAction(v []*string) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Action = v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetCondition(v map[string]interface{}) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Condition = v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetEffect(v string) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Effect = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetResource(v string) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Resource = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) Validate() error {
	return dara.Validate(s)
}

type GenerateTemplatePolicyResponseBodyPolicyFunctions struct {
	Action                *string                                                                   `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionPolicyFunctions []*GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions `json:"ActionPolicyFunctions,omitempty" xml:"ActionPolicyFunctions,omitempty" type:"Repeated"`
	RequirementLevel      *string                                                                   `json:"RequirementLevel,omitempty" xml:"RequirementLevel,omitempty"`
}

func (s GenerateTemplatePolicyResponseBodyPolicyFunctions) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBodyPolicyFunctions) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctions) GetAction() *string {
	return s.Action
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctions) GetActionPolicyFunctions() []*GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions {
	return s.ActionPolicyFunctions
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctions) GetRequirementLevel() *string {
	return s.RequirementLevel
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctions) SetAction(v string) *GenerateTemplatePolicyResponseBodyPolicyFunctions {
	s.Action = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctions) SetActionPolicyFunctions(v []*GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions) *GenerateTemplatePolicyResponseBodyPolicyFunctions {
	s.ActionPolicyFunctions = v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctions) SetRequirementLevel(v string) *GenerateTemplatePolicyResponseBodyPolicyFunctions {
	s.RequirementLevel = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctions) Validate() error {
	if s.ActionPolicyFunctions != nil {
		for _, item := range s.ActionPolicyFunctions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions struct {
	Functions         []*GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
	LogicalResourceId *string                                                                            `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	ResourceType      *string                                                                            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions) GetFunctions() []*GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions {
	return s.Functions
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions) GetResourceType() *string {
	return s.ResourceType
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions) SetFunctions(v []*GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions {
	s.Functions = v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions) SetLogicalResourceId(v string) *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions {
	s.LogicalResourceId = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions) SetResourceType(v string) *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions {
	s.ResourceType = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctions) Validate() error {
	if s.Functions != nil {
		for _, item := range s.Functions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions struct {
	Function          *string   `json:"Function,omitempty" xml:"Function,omitempty"`
	OperationType     *string   `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	RelatedProperties []*string `json:"RelatedProperties,omitempty" xml:"RelatedProperties,omitempty" type:"Repeated"`
	RequirementLevel  *string   `json:"RequirementLevel,omitempty" xml:"RequirementLevel,omitempty"`
}

func (s GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) GetFunction() *string {
	return s.Function
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) GetOperationType() *string {
	return s.OperationType
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) GetRelatedProperties() []*string {
	return s.RelatedProperties
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) GetRequirementLevel() *string {
	return s.RequirementLevel
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) SetFunction(v string) *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions {
	s.Function = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) SetOperationType(v string) *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions {
	s.OperationType = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) SetRelatedProperties(v []*string) *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions {
	s.RelatedProperties = v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) SetRequirementLevel(v string) *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions {
	s.RequirementLevel = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyFunctionsActionPolicyFunctionsFunctions) Validate() error {
	return dara.Validate(s)
}
