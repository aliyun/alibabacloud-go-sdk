// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAINodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *AddAINodeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *AddAINodeResponseBody
	GetRequestId() *string
}

type AddAINodeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 111111111111
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAINodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAINodeResponseBody) GoString() string {
	return s.String()
}

func (s *AddAINodeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *AddAINodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAINodeResponseBody) SetOrderId(v string) *AddAINodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *AddAINodeResponseBody) SetRequestId(v string) *AddAINodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAINodeResponseBody) Validate() error {
	return dara.Validate(s)
}
