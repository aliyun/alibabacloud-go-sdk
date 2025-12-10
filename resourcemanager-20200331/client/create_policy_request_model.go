// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePolicyRequest
	GetDescription() *string
	SetPolicyDocument(v string) *CreatePolicyRequest
	GetPolicyDocument() *string
	SetPolicyName(v string) *CreatePolicyRequest
	GetPolicyName() *string
}

type CreatePolicyRequest struct {
	// The description of the permission policy.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// OSS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document of the policy.
	//
	// The document must be 1 to 6,144 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "Statement": [{ "Action": ["oss:*"], "Effect": "Allow", "Resource": ["acs:oss:*:*:*"]}], "Version": "1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the permission policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyRequest) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *CreatePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyRequest) SetDescription(v string) *CreatePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyDocument(v string) *CreatePolicyRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyName(v string) *CreatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyRequest) Validate() error {
	return dara.Validate(s)
}
