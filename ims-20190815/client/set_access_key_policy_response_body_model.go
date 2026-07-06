// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessKeyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyId(v string) *SetAccessKeyPolicyResponseBody
	GetAccessKeyId() *string
	SetAccessKeyPolicy(v string) *SetAccessKeyPolicyResponseBody
	GetAccessKeyPolicy() *string
	SetRequestId(v string) *SetAccessKeyPolicyResponseBody
	GetRequestId() *string
}

type SetAccessKeyPolicyResponseBody struct {
	// The AccessKey ID.
	//
	// example:
	//
	// LTAI*******************
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The network access restriction policy.
	//
	// A JSON-formatted string. For more information, see the AccessKeyPolicy structure description.
	//
	// example:
	//
	// {"Status":"Inactive","Statements":[{"Value":"AllowAllVPC","Type":"VPCWhiteList","IPList":["::/0","0.0.0.0/0"]}]}
	AccessKeyPolicy *string `json:"AccessKeyPolicy,omitempty" xml:"AccessKeyPolicy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30C9068D-FBAA-4998-9986-8A562FED0BC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccessKeyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAccessKeyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccessKeyPolicyResponseBody) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *SetAccessKeyPolicyResponseBody) GetAccessKeyPolicy() *string {
	return s.AccessKeyPolicy
}

func (s *SetAccessKeyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAccessKeyPolicyResponseBody) SetAccessKeyId(v string) *SetAccessKeyPolicyResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *SetAccessKeyPolicyResponseBody) SetAccessKeyPolicy(v string) *SetAccessKeyPolicyResponseBody {
	s.AccessKeyPolicy = &v
	return s
}

func (s *SetAccessKeyPolicyResponseBody) SetRequestId(v string) *SetAccessKeyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAccessKeyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
