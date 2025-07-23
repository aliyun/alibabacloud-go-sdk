// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *CancelOrderRequestRequest
	GetOrderId() *int64
}

type CancelOrderRequestRequest struct {
	// The order ID.
	//
	// >  You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123451222
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CancelOrderRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderRequestRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRequestRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CancelOrderRequestRequest) SetOrderId(v int64) *CancelOrderRequestRequest {
	s.OrderId = &v
	return s
}

func (s *CancelOrderRequestRequest) Validate() error {
	return dara.Validate(s)
}
