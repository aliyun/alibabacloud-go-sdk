// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessKeyPolicyForAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyPolicy(v string) *SetAccessKeyPolicyForAccountResponseBody
	GetAccessKeyPolicy() *string
	SetRequestId(v string) *SetAccessKeyPolicyForAccountResponseBody
	GetRequestId() *string
}

type SetAccessKeyPolicyForAccountResponseBody struct {
	// The network access restriction policy. The value is a JSON-formatted string. For more information, see the AccessKeyPolicy structure description in the SetAccessKeyPolicy documentation.
	//
	// example:
	//
	// {"Status":"Inactive","Statements":[{"Value":"AllowAllVPC","Type":"VPCWhiteList","IPList":["::/0","0.0.0.0/0"]}]}
	AccessKeyPolicy *string `json:"AccessKeyPolicy,omitempty" xml:"AccessKeyPolicy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6616F09B-2768-4C11-8866-A8EE4C4A583E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccessKeyPolicyForAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAccessKeyPolicyForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccessKeyPolicyForAccountResponseBody) GetAccessKeyPolicy() *string {
	return s.AccessKeyPolicy
}

func (s *SetAccessKeyPolicyForAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAccessKeyPolicyForAccountResponseBody) SetAccessKeyPolicy(v string) *SetAccessKeyPolicyForAccountResponseBody {
	s.AccessKeyPolicy = &v
	return s
}

func (s *SetAccessKeyPolicyForAccountResponseBody) SetRequestId(v string) *SetAccessKeyPolicyForAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAccessKeyPolicyForAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
