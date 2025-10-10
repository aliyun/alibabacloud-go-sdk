// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteProductResponseBody
	GetRequestId() *string
}

type DeleteProductResponseBody struct {
	// example:
	//
	// PRODUCT_NOT_ALONE
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProductResponseBody) SetMessage(v string) *DeleteProductResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProductResponseBody) SetRequestId(v string) *DeleteProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProductResponseBody) Validate() error {
	return dara.Validate(s)
}
