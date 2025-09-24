// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuyProductRequestId(v string) *CreateProductOrdersResponseBody
	GetBuyProductRequestId() *string
	SetMessage(v string) *CreateProductOrdersResponseBody
	GetMessage() *string
	SetOrderId(v string) *CreateProductOrdersResponseBody
	GetOrderId() *string
	SetProductIds(v []*string) *CreateProductOrdersResponseBody
	GetProductIds() []*string
	SetRequestId(v string) *CreateProductOrdersResponseBody
	GetRequestId() *string
}

type CreateProductOrdersResponseBody struct {
	// The ID of the product purchase request.
	//
	// example:
	//
	// 3ed6a882-0d85-4dd8-ad36-cd8d74ab9fdb
	BuyProductRequestId *string `json:"BuyProductRequestId,omitempty" xml:"BuyProductRequestId,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The purchase order ID.
	//
	// example:
	//
	// 210292536260646
	OrderId    *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ProductIds []*string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ksdjf-jksd-*****slkdjf
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProductOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProductOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductOrdersResponseBody) GetBuyProductRequestId() *string {
	return s.BuyProductRequestId
}

func (s *CreateProductOrdersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateProductOrdersResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateProductOrdersResponseBody) GetProductIds() []*string {
	return s.ProductIds
}

func (s *CreateProductOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProductOrdersResponseBody) SetBuyProductRequestId(v string) *CreateProductOrdersResponseBody {
	s.BuyProductRequestId = &v
	return s
}

func (s *CreateProductOrdersResponseBody) SetMessage(v string) *CreateProductOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProductOrdersResponseBody) SetOrderId(v string) *CreateProductOrdersResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateProductOrdersResponseBody) SetProductIds(v []*string) *CreateProductOrdersResponseBody {
	s.ProductIds = v
	return s
}

func (s *CreateProductOrdersResponseBody) SetRequestId(v string) *CreateProductOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductOrdersResponseBody) Validate() error {
	return dara.Validate(s)
}
