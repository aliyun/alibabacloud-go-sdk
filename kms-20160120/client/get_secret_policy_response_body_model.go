// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v string) *GetSecretPolicyResponseBody
	GetPolicy() *string
	SetRequestId(v string) *GetSecretPolicyResponseBody
	GetRequestId() *string
}

type GetSecretPolicyResponseBody struct {
	// example:
	//
	// {"Version":"1","Statement": [{"Sid":"kms default secret policy","Effect":"Allow","Principal":{"RAM": ["acs:ram::119285303511****:*"]},"Action":["kms:*"],"Resource": ["*"] }] }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSecretPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecretPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretPolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GetSecretPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecretPolicyResponseBody) SetPolicy(v string) *GetSecretPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetSecretPolicyResponseBody) SetRequestId(v string) *GetSecretPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
