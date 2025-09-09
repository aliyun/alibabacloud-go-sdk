// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectApproveCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RejectApproveCommandResponseBody
	GetRequestId() *string
}

type RejectApproveCommandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectApproveCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RejectApproveCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RejectApproveCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RejectApproveCommandResponseBody) SetRequestId(v string) *RejectApproveCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectApproveCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
