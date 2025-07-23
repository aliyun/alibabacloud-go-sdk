// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewCertificateOrderForPackageRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *RenewCertificateOrderForPackageRequestResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *RenewCertificateOrderForPackageRequestResponseBody
	GetRequestId() *string
}

type RenewCertificateOrderForPackageRequestResponseBody struct {
	// The ID of the certificate application order that is renewed.
	//
	// >  You can use the ID to query the status of the certificate application order. For more information, see [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html).
	//
	// example:
	//
	// 323451222
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewCertificateOrderForPackageRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewCertificateOrderForPackageRequestResponseBody) GoString() string {
	return s.String()
}

func (s *RenewCertificateOrderForPackageRequestResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RenewCertificateOrderForPackageRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewCertificateOrderForPackageRequestResponseBody) SetOrderId(v int64) *RenewCertificateOrderForPackageRequestResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewCertificateOrderForPackageRequestResponseBody) SetRequestId(v string) *RenewCertificateOrderForPackageRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewCertificateOrderForPackageRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
