// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecretPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v string) *SetSecretPolicyRequest
	GetPolicy() *string
	SetPolicyName(v string) *SetSecretPolicyRequest
	GetPolicyName() *string
	SetSecretName(v string) *SetSecretPolicyRequest
	GetSecretName() *string
}

type SetSecretPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"Version":"1","Statement": [{"Sid":"kms default secret policy","Effect":"Allow","Principal":{"RAM": ["acs:ram::119285303511****:*"]},"Action":["kms:*"],"Resource": ["*"] }] }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// default
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// secret_test
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s SetSecretPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSecretPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetSecretPolicyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *SetSecretPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *SetSecretPolicyRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *SetSecretPolicyRequest) SetPolicy(v string) *SetSecretPolicyRequest {
	s.Policy = &v
	return s
}

func (s *SetSecretPolicyRequest) SetPolicyName(v string) *SetSecretPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *SetSecretPolicyRequest) SetSecretName(v string) *SetSecretPolicyRequest {
	s.SecretName = &v
	return s
}

func (s *SetSecretPolicyRequest) Validate() error {
	return dara.Validate(s)
}
