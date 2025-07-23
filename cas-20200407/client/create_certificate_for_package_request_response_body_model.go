// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertificateForPackageRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *CreateCertificateForPackageRequestResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateCertificateForPackageRequestResponseBody
	GetRequestId() *string
}

type CreateCertificateForPackageRequestResponseBody struct {
	// The ID of the certificate application order.
	//
	// >  You can use the ID to query the status of the certificate application order. For more information, see [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html).
	//
	// example:
	//
	// 2021010
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5890029B-938A-589E-98B9-3DEC7BA7C400
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertificateForPackageRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateForPackageRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateForPackageRequestResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateCertificateForPackageRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCertificateForPackageRequestResponseBody) SetOrderId(v int64) *CreateCertificateForPackageRequestResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCertificateForPackageRequestResponseBody) SetRequestId(v string) *CreateCertificateForPackageRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCertificateForPackageRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
