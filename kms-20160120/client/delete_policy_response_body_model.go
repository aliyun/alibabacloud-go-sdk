// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePolicyResponseBody
	GetRequestId() *string
}

type DeletePolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 00a26a33-d992-42f3-9012-5fd12764430f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolicyResponseBody) SetRequestId(v string) *DeletePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
