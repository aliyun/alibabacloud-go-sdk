// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyDocument(v string) *CreatePolicyVersionRequest
	GetPolicyDocument() *string
	SetPolicyName(v string) *CreatePolicyVersionRequest
	GetPolicyName() *string
	SetSetAsDefault(v bool) *CreatePolicyVersionRequest
	GetSetAsDefault() *bool
}

type CreatePolicyVersionRequest struct {
	// The document of the permission policy.
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
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphen (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// Specifies whether to set the policy version as the default version.
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// false
	SetAsDefault *bool `json:"SetAsDefault,omitempty" xml:"SetAsDefault,omitempty"`
}

func (s CreatePolicyVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionRequest) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *CreatePolicyVersionRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyVersionRequest) GetSetAsDefault() *bool {
	return s.SetAsDefault
}

func (s *CreatePolicyVersionRequest) SetPolicyDocument(v string) *CreatePolicyVersionRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreatePolicyVersionRequest) SetPolicyName(v string) *CreatePolicyVersionRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyVersionRequest) SetSetAsDefault(v bool) *CreatePolicyVersionRequest {
	s.SetAsDefault = &v
	return s
}

func (s *CreatePolicyVersionRequest) Validate() error {
	return dara.Validate(s)
}
