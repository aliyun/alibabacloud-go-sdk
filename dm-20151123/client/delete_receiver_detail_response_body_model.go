// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReceiverDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteReceiverDetailResponseBody
	GetRequestId() *string
}

type DeleteReceiverDetailResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReceiverDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteReceiverDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReceiverDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteReceiverDetailResponseBody) SetRequestId(v string) *DeleteReceiverDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteReceiverDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
