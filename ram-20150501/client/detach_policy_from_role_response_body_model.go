// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyFromRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachPolicyFromRoleResponseBody
	GetRequestId() *string
}

type DetachPolicyFromRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicyFromRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyFromRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachPolicyFromRoleResponseBody) SetRequestId(v string) *DetachPolicyFromRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachPolicyFromRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
