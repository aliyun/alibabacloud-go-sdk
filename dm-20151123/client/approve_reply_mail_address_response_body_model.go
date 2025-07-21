// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveReplyMailAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApproveReplyMailAddressResponseBody
	GetRequestId() *string
}

type ApproveReplyMailAddressResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveReplyMailAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApproveReplyMailAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveReplyMailAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApproveReplyMailAddressResponseBody) SetRequestId(v string) *ApproveReplyMailAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveReplyMailAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
