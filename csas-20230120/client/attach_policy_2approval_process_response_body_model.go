// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicy2ApprovalProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachPolicy2ApprovalProcessResponseBody
	GetRequestId() *string
}

type AttachPolicy2ApprovalProcessResponseBody struct {
	// example:
	//
	// C51D9340-4604-5331-AE62-407F3B408F86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPolicy2ApprovalProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicy2ApprovalProcessResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPolicy2ApprovalProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachPolicy2ApprovalProcessResponseBody) SetRequestId(v string) *AttachPolicy2ApprovalProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachPolicy2ApprovalProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
