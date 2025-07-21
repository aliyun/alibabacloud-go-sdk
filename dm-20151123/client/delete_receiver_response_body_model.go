// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReceiverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteReceiverResponseBody
	GetRequestId() *string
}

type DeleteReceiverResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReceiverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteReceiverResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReceiverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteReceiverResponseBody) SetRequestId(v string) *DeleteReceiverResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteReceiverResponseBody) Validate() error {
	return dara.Validate(s)
}
