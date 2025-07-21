// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v string) *GetKeyPolicyResponseBody
	GetPolicy() *string
	SetRequestId(v string) *GetKeyPolicyResponseBody
	GetRequestId() *string
}

type GetKeyPolicyResponseBody struct {
	// example:
	//
	// {"Statement": [{"Action": ["kms:*"],"Effect": "Allow","Principal": {"RAM": ["acs:ram::190325303126****:*","acs:ram::119285303511****:*"]},"Resource": ["*"],"Sid": "kms default key policy"}],"Version": "1" }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3B84B523C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetKeyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKeyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeyPolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GetKeyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKeyPolicyResponseBody) SetPolicy(v string) *GetKeyPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetKeyPolicyResponseBody) SetRequestId(v string) *GetKeyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKeyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
