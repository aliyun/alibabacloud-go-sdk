// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v int64) *CreateProductResponseBody
	GetProductId() *int64
	SetRequestId(v string) *CreateProductResponseBody
	GetRequestId() *string
}

type CreateProductResponseBody struct {
	// example:
	//
	// 123456
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBody) GetProductId() *int64 {
	return s.ProductId
}

func (s *CreateProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProductResponseBody) SetProductId(v int64) *CreateProductResponseBody {
	s.ProductId = &v
	return s
}

func (s *CreateProductResponseBody) SetRequestId(v string) *CreateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductResponseBody) Validate() error {
	return dara.Validate(s)
}
