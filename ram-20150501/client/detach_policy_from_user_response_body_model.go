// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyFromUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachPolicyFromUserResponseBody
	GetRequestId() *string
}

type DetachPolicyFromUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicyFromUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachPolicyFromUserResponseBody) SetRequestId(v string) *DetachPolicyFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachPolicyFromUserResponseBody) Validate() error {
	return dara.Validate(s)
}
