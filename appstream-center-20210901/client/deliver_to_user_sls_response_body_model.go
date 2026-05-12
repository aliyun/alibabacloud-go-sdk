// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeliverToUserSlsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeliverToUserSlsResponseBody
	GetRequestId() *string
}

type DeliverToUserSlsResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeliverToUserSlsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeliverToUserSlsResponseBody) GoString() string {
	return s.String()
}

func (s *DeliverToUserSlsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeliverToUserSlsResponseBody) SetRequestId(v string) *DeliverToUserSlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeliverToUserSlsResponseBody) Validate() error {
	return dara.Validate(s)
}
