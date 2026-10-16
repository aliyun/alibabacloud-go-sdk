// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateControlPolicyRequest
	GetDescription() *string
	SetEffectScope(v string) *CreateControlPolicyRequest
	GetEffectScope() *string
	SetPolicyDocument(v string) *CreateControlPolicyRequest
	GetPolicyDocument() *string
	SetPolicyName(v string) *CreateControlPolicyRequest
	GetPolicyName() *string
}

type CreateControlPolicyRequest struct {
	// The description of the control policy.
	//
	// The description must be 1 to 1024 characters in length, must start with a letter, and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// ExampleControlPolicy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the control policy.
	//
	// Valid values: RAM, which indicates that the control policy takes effect only for Resource Access Management (RAM) users or RAM roles.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAM
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The content of the control policy.
	//
	// The maximum length is 4096 characters.
	//
	// For information about the control policy language, see [Control policy language](https://help.aliyun.com/document_detail/179096.html).
	//
	// For examples of control policies, see [Examples of custom control policies](https://help.aliyun.com/document_detail/181474.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Version":"1","Statement":[{"Effect":"Deny","Action":["ram:UpdateRole","ram:DeleteRole","ram:AttachPolicyToRole","ram:DetachPolicyFromRole"],"Resource":"acs:ram:*:*:role/ResourceDirectoryAccountAccessRole"}]}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// Policy Name of the control policy.
	//
	// Policy Name must be 1 to 128 characters in length, must start with a letter, and can contain letters, digits, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// ExampleControlPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s CreateControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateControlPolicyRequest) GetEffectScope() *string {
	return s.EffectScope
}

func (s *CreateControlPolicyRequest) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *CreateControlPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateControlPolicyRequest) SetDescription(v string) *CreateControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateControlPolicyRequest) SetEffectScope(v string) *CreateControlPolicyRequest {
	s.EffectScope = &v
	return s
}

func (s *CreateControlPolicyRequest) SetPolicyDocument(v string) *CreateControlPolicyRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreateControlPolicyRequest) SetPolicyName(v string) *CreateControlPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
