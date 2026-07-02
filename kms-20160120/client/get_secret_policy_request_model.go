// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *GetSecretPolicyRequest
	GetPolicyName() *string
	SetSecretName(v string) *GetSecretPolicyRequest
	GetSecretName() *string
}

type GetSecretPolicyRequest struct {
	// The name of the credential policy. Only the static field default is supported.
	//
	// example:
	//
	// default
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The name or Alibaba Cloud Resource Name (ARN) of the credential.
	//
	// > If you access a credential that belongs to another Alibaba Cloud account, you must specify the ARN of the credential. The ARN of a credential must be in the `acs:kms:${region}:${account}:secret/${secret-name}` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret_test
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s GetSecretPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecretPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetSecretPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetSecretPolicyRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *GetSecretPolicyRequest) SetPolicyName(v string) *GetSecretPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *GetSecretPolicyRequest) SetSecretName(v string) *GetSecretPolicyRequest {
	s.SecretName = &v
	return s
}

func (s *GetSecretPolicyRequest) Validate() error {
	return dara.Validate(s)
}
