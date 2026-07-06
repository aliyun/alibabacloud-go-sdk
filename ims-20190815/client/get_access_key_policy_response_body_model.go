// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyId(v string) *GetAccessKeyPolicyResponseBody
	GetAccessKeyId() *string
	SetAccessKeyPolicy(v string) *GetAccessKeyPolicyResponseBody
	GetAccessKeyPolicy() *string
	SetRequestId(v string) *GetAccessKeyPolicyResponseBody
	GetRequestId() *string
}

type GetAccessKeyPolicyResponseBody struct {
	// The access key ID.
	//
	// example:
	//
	// LTAI*******************
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The network access restriction policy. The value is a JSON string. For more information, see the AccessKeyPolicy structure description in the SetAccessKeyPolicy documentation.
	//
	// example:
	//
	// {"Status":"Inactive","Statements":[{"Value":"AllowAllVPC","Type":"VPCWhiteList","IPList":["::/0","0.0.0.0/0"]}]}
	AccessKeyPolicy *string `json:"AccessKeyPolicy,omitempty" xml:"AccessKeyPolicy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyPolicyResponseBody) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetAccessKeyPolicyResponseBody) GetAccessKeyPolicy() *string {
	return s.AccessKeyPolicy
}

func (s *GetAccessKeyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessKeyPolicyResponseBody) SetAccessKeyId(v string) *GetAccessKeyPolicyResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetAccessKeyPolicyResponseBody) SetAccessKeyPolicy(v string) *GetAccessKeyPolicyResponseBody {
	s.AccessKeyPolicy = &v
	return s
}

func (s *GetAccessKeyPolicyResponseBody) SetRequestId(v string) *GetAccessKeyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
