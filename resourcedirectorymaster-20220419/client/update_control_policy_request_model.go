// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewDescription(v string) *UpdateControlPolicyRequest
	GetNewDescription() *string
	SetNewPolicyDocument(v string) *UpdateControlPolicyRequest
	GetNewPolicyDocument() *string
	SetNewPolicyName(v string) *UpdateControlPolicyRequest
	GetNewPolicyName() *string
	SetPolicyId(v string) *UpdateControlPolicyRequest
	GetPolicyId() *string
}

type UpdateControlPolicyRequest struct {
	// The new description of the access control policy.
	//
	// The description must be 1 to 1,024 characters in length. The description can contain letters, digits, underscores (_), and hyphens (-) and must start with a letter.
	//
	// example:
	//
	// ExampleControlPolicy
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new document of the access control policy.
	//
	// The document can be a maximum of 4,096 characters in length.
	//
	// For more information about the languages of access control policies, see [Languages of access control policies](https://help.aliyun.com/document_detail/179096.html).
	//
	// For more information about the examples of access control policies, see [Examples of custom access control policies](https://help.aliyun.com/document_detail/181474.html).
	//
	// example:
	//
	// {"Version":"1","Statement":[{"Effect":"Deny","Action":["ram:UpdateRole","ram:DeleteRole","ram:AttachPolicyToRole","ram:DetachPolicyFromRole"],"Resource":"acs:ram:*:*:role/ResourceDirectoryAccountAccessRole"}]}
	NewPolicyDocument *string `json:"NewPolicyDocument,omitempty" xml:"NewPolicyDocument,omitempty"`
	// The new name of the access control policy.
	//
	// The name must be 1 to 128 characters in length. The name can contain letters, digits, and hyphens (-) and must start with a letter.
	//
	// example:
	//
	// NewControlPolicy
	NewPolicyName *string `json:"NewPolicyName,omitempty" xml:"NewPolicyName,omitempty"`
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s UpdateControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyRequest) GetNewDescription() *string {
	return s.NewDescription
}

func (s *UpdateControlPolicyRequest) GetNewPolicyDocument() *string {
	return s.NewPolicyDocument
}

func (s *UpdateControlPolicyRequest) GetNewPolicyName() *string {
	return s.NewPolicyName
}

func (s *UpdateControlPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateControlPolicyRequest) SetNewDescription(v string) *UpdateControlPolicyRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetNewPolicyDocument(v string) *UpdateControlPolicyRequest {
	s.NewPolicyDocument = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetNewPolicyName(v string) *UpdateControlPolicyRequest {
	s.NewPolicyName = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetPolicyId(v string) *UpdateControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
