// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyToRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachPolicyToRoleResponseBody
	GetRequestId() *string
}

type AttachPolicyToRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 697852FB-50D7-44D9-9774-530C31EAC572
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPolicyToRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyToRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPolicyToRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachPolicyToRoleResponseBody) SetRequestId(v string) *AttachPolicyToRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachPolicyToRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
