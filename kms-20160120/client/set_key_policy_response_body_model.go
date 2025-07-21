// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetKeyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetKeyPolicyResponseBody
	GetRequestId() *string
}

type SetKeyPolicyResponseBody struct {
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetKeyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetKeyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetKeyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetKeyPolicyResponseBody) SetRequestId(v string) *SetKeyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetKeyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
