// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptApproveCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AcceptApproveCommandResponseBody
	GetRequestId() *string
}

type AcceptApproveCommandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptApproveCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcceptApproveCommandResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptApproveCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcceptApproveCommandResponseBody) SetRequestId(v string) *AcceptApproveCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptApproveCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
