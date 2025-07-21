// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetKeyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *SetKeyPolicyRequest
	GetKeyId() *string
	SetPolicy(v string) *SetKeyPolicyRequest
	GetPolicy() *string
	SetPolicyName(v string) *SetKeyPolicyRequest
	GetPolicyName() *string
}

type SetKeyPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Statement":[{"Action":["kms:*"],"Effect":"Allow","Principal":{"RAM":["acs:ram::119285303511****:*"]},"Resource":["*"],"Sid":"kms default key policy"},{"Action":["kms:List*","kms:Describe*","kms:Create*","kms:Enable*","kms:Disable*","kms:Get*","kms:Set*","kms:Update*","kms:Delete*","kms:Cancel*","kms:TagResource","kms:UntagResource","kms:ImportKeyMaterial","kms:ScheduleKeyDeletion"],"Effect":"Allow","Principal":{"RAM":["acs:ram::119285303511****:user/for_test_policy"]},"Resource":["*"]}],"Version":"1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// default
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s SetKeyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetKeyPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetKeyPolicyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *SetKeyPolicyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *SetKeyPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *SetKeyPolicyRequest) SetKeyId(v string) *SetKeyPolicyRequest {
	s.KeyId = &v
	return s
}

func (s *SetKeyPolicyRequest) SetPolicy(v string) *SetKeyPolicyRequest {
	s.Policy = &v
	return s
}

func (s *SetKeyPolicyRequest) SetPolicyName(v string) *SetKeyPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *SetKeyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
