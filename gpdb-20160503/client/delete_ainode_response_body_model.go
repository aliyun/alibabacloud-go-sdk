// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAINodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *DeleteAINodeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *DeleteAINodeResponseBody
	GetRequestId() *string
}

type DeleteAINodeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// *********
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAINodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAINodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAINodeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DeleteAINodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAINodeResponseBody) SetOrderId(v string) *DeleteAINodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *DeleteAINodeResponseBody) SetRequestId(v string) *DeleteAINodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAINodeResponseBody) Validate() error {
	return dara.Validate(s)
}
