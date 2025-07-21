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
