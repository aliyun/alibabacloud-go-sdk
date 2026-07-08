// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCertificateForPackageRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *CancelCertificateForPackageRequestRequest
	GetOrderId() *int64
}

type CancelCertificateForPackageRequestRequest struct {
	// The ID of the order.
	//
	// > Call [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CancelCertificateForPackageRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCertificateForPackageRequestRequest) GoString() string {
	return s.String()
}

func (s *CancelCertificateForPackageRequestRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CancelCertificateForPackageRequestRequest) SetOrderId(v int64) *CancelCertificateForPackageRequestRequest {
	s.OrderId = &v
	return s
}

func (s *CancelCertificateForPackageRequestRequest) Validate() error {
	return dara.Validate(s)
}
