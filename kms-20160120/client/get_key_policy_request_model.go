// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *GetKeyPolicyRequest
	GetKeyId() *string
	SetPolicyName(v string) *GetKeyPolicyRequest
	GetPolicyName() *string
}

type GetKeyPolicyRequest struct {
	// The ID or Alibaba Cloud Resource Name (ARN) of the key.
	//
	// > When you access a key in another Alibaba Cloud account, you must enter the ARN of the key. The ARN of a key is in the `acs:kms:${region}:${account}:key/${keyid}` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the key policy. Only the static value default is supported.
	//
	// example:
	//
	// default
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s GetKeyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKeyPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetKeyPolicyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GetKeyPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetKeyPolicyRequest) SetKeyId(v string) *GetKeyPolicyRequest {
	s.KeyId = &v
	return s
}

func (s *GetKeyPolicyRequest) SetPolicyName(v string) *GetKeyPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *GetKeyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
