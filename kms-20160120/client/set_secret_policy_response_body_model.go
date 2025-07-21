// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecretPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSecretPolicyResponseBody
	GetRequestId() *string
}

type SetSecretPolicyResponseBody struct {
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSecretPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSecretPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetSecretPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSecretPolicyResponseBody) SetRequestId(v string) *SetSecretPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSecretPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
