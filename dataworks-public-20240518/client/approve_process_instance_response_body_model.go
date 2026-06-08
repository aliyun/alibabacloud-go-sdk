// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveProcessInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApproveProcessInstanceResponseBody
	GetRequestId() *string
}

type ApproveProcessInstanceResponseBody struct {
	// example:
	//
	// 0bc5df3a17***903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveProcessInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApproveProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveProcessInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApproveProcessInstanceResponseBody) SetRequestId(v string) *ApproveProcessInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveProcessInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
