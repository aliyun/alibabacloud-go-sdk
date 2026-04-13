// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteChatResponseBody
	GetRequestId() *string
}

type DeleteChatResponseBody struct {
	// example:
	//
	// 44553E9A-******-37ADC33FE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChatResponseBody) SetRequestId(v string) *DeleteChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChatResponseBody) Validate() error {
	return dara.Validate(s)
}
