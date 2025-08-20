// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetStackPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetStackPolicyResponseBody
	GetRequestId() *string
}

type SetStackPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetStackPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetStackPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetStackPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetStackPolicyResponseBody) SetRequestId(v string) *SetStackPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetStackPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
