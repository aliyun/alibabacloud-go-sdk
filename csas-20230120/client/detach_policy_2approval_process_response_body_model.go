// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicy2ApprovalProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachPolicy2ApprovalProcessResponseBody
	GetRequestId() *string
}

type DetachPolicy2ApprovalProcessResponseBody struct {
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicy2ApprovalProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicy2ApprovalProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicy2ApprovalProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachPolicy2ApprovalProcessResponseBody) SetRequestId(v string) *DetachPolicy2ApprovalProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachPolicy2ApprovalProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
