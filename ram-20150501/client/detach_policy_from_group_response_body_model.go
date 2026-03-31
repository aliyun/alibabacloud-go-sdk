// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyFromGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachPolicyFromGroupResponseBody
	GetRequestId() *string
}

type DetachPolicyFromGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicyFromGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachPolicyFromGroupResponseBody) SetRequestId(v string) *DetachPolicyFromGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachPolicyFromGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
