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
	SetRotateStrategy(v string) *CreatePolicyVersionRequest
	GetRotateStrategy() *string
	SetSetAsDefault(v bool) *CreatePolicyVersionRequest
	GetSetAsDefault() *bool
}

type CreatePolicyVersionRequest struct {
	// The document of the policy. The document can be up to 6,144 bytes in length.
	//
	// example:
	//
	// {"Statement":[{"Action":["oss:*"],"Effect":"Allow","Resource":["acs:oss:*:*:*"]}],"Version":"1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The rotation strategy of the policy. The rotation strategy can be used to delete an early policy version.
	//
	// Valid values:
	//
	// 	- `None`: disables the rotation strategy.
	//
	// 	- `DeleteOldestNonDefaultVersionWhenLimitExceeded`: deletes the earliest non-active version if the number of versions exceeds the limit.
	//
	// Default value: `None`.
	//
	// example:
	//
	// None
	RotateStrategy *string `json:"RotateStrategy,omitempty" xml:"RotateStrategy,omitempty"`
	// Specifies whether to set this policy as the default policy. Default value: `false`.
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

func (s *CreatePolicyVersionRequest) GetRotateStrategy() *string {
	return s.RotateStrategy
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

func (s *CreatePolicyVersionRequest) SetRotateStrategy(v string) *CreatePolicyVersionRequest {
	s.RotateStrategy = &v
	return s
}

func (s *CreatePolicyVersionRequest) SetSetAsDefault(v bool) *CreatePolicyVersionRequest {
	s.SetAsDefault = &v
	return s
}

func (s *CreatePolicyVersionRequest) Validate() error {
	return dara.Validate(s)
}
