// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachPolicyResponseBody
	GetRequestId() *string
}

type DetachPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachPolicyResponseBody) SetRequestId(v string) *DetachPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
