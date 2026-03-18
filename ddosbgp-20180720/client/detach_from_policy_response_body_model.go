// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachFromPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachFromPolicyResponseBody
	GetRequestId() *string
}

type DetachFromPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1B0F7EC6-51D7-4D70-B0EC-CD8A9E99****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachFromPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachFromPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachFromPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachFromPolicyResponseBody) SetRequestId(v string) *DetachFromPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachFromPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
