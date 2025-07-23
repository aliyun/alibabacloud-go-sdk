// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCertificateRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *DeleteCertificateRequestRequest
	GetOrderId() *int64
}

type DeleteCertificateRequestRequest struct {
	// The ID of the certificate application order that you want to delete.
	//
	// >  After you call the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html), [CreateCertificateRequest](https://help.aliyun.com/document_detail/455292.html), or [CreateCertificateWithCsrRequest](https://help.aliyun.com/document_detail/455801.html) operation to submit a certificate application, you can obtain the ID of the certificate application order from the **OrderId*	- response parameter. You can also call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the order ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123451222
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DeleteCertificateRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCertificateRequestRequest) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequestRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DeleteCertificateRequestRequest) SetOrderId(v int64) *DeleteCertificateRequestRequest {
	s.OrderId = &v
	return s
}

func (s *DeleteCertificateRequestRequest) Validate() error {
	return dara.Validate(s)
}
