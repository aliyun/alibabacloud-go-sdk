// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyApprovalConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPolicyApprovalConfigResponseBody
	GetRequestId() *string
}

type SetPolicyApprovalConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPolicyApprovalConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyApprovalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetPolicyApprovalConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPolicyApprovalConfigResponseBody) SetRequestId(v string) *SetPolicyApprovalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPolicyApprovalConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
