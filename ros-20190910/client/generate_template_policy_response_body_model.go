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
	SetRequestId(v string) *GenerateTemplatePolicyResponseBody
	GetRequestId() *string
}

type GenerateTemplatePolicyResponseBody struct {
	// The information about the policy.
	Policy *GenerateTemplatePolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
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

func (s *GenerateTemplatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateTemplatePolicyResponseBody) SetPolicy(v *GenerateTemplatePolicyResponseBodyPolicy) *GenerateTemplatePolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GenerateTemplatePolicyResponseBody) SetRequestId(v string) *GenerateTemplatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Validate(s)
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
