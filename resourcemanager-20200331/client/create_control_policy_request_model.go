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
	// The description of the access control policy.
	//
	// The description must be 1 to 1,024 characters in length. The description can contain letters, digits, underscores (_), and hyphens (-) and must start with a letter.
	//
	// example:
	//
	// ExampleControlPolicy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy.
	//
	// The value RAM indicates that the access control policy takes effect only for RAM users and RAM roles.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAM
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The document of the access control policy.
	//
	// The document can be a maximum of 4,096 characters in length.
	//
	// For more information about the languages of access control policies, see [Languages of access control policies](https://help.aliyun.com/document_detail/179096.html).
	//
	// For more information about the examples of access control policies, see [Examples of custom access control policies](https://help.aliyun.com/document_detail/181474.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Version":"1","Statement":[{"Effect":"Deny","Action":["ram:UpdateRole","ram:DeleteRole","ram:AttachPolicyToRole","ram:DetachPolicyFromRole"],"Resource":"acs:ram:*:*:role/ResourceDirectoryAccountAccessRole"}]}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the access control policy.
	//
	// The name must be 1 to 128 characters in length. The name can contain letters, digits, and hyphens (-) and must start with a letter.
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
