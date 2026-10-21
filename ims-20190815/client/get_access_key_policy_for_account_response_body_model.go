// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyPolicyForAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyPolicy(v string) *GetAccessKeyPolicyForAccountResponseBody
	GetAccessKeyPolicy() *string
	SetRequestId(v string) *GetAccessKeyPolicyForAccountResponseBody
	GetRequestId() *string
}

type GetAccessKeyPolicyForAccountResponseBody struct {
	// The network access restriction policy. The value is a JSON-formatted string. For more information about the AccessKeyPolicy structure, see the SetAccessKeyPolicy documentation.
	//
	// example:
	//
	// {"Status":"Inactive","Statements":[{"Value":"AllowAllVPC","Type":"VPCWhiteList","IPList":["::/0","0.0.0.0/0"]}]}
	AccessKeyPolicy *string `json:"AccessKeyPolicy,omitempty" xml:"AccessKeyPolicy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyPolicyForAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyPolicyForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyPolicyForAccountResponseBody) GetAccessKeyPolicy() *string {
	return s.AccessKeyPolicy
}

func (s *GetAccessKeyPolicyForAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessKeyPolicyForAccountResponseBody) SetAccessKeyPolicy(v string) *GetAccessKeyPolicyForAccountResponseBody {
	s.AccessKeyPolicy = &v
	return s
}

func (s *GetAccessKeyPolicyForAccountResponseBody) SetRequestId(v string) *GetAccessKeyPolicyForAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyPolicyForAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
