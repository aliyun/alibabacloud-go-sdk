// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnblockSendingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnblockSendingResponseBody
	GetRequestId() *string
}

type UnblockSendingResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnblockSendingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnblockSendingResponseBody) GoString() string {
	return s.String()
}

func (s *UnblockSendingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnblockSendingResponseBody) SetRequestId(v string) *UnblockSendingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnblockSendingResponseBody) Validate() error {
	return dara.Validate(s)
}
