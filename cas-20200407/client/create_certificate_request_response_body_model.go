// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertificateRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *CreateCertificateRequestResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateCertificateRequestResponseBody
	GetRequestId() *string
}

type CreateCertificateRequestResponseBody struct {
	// The ID of the certificate application order.
	//
	// >  You can use the ID to query the status of the certificate application. For more information, see [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html).
	//
	// example:
	//
	// 98987582437920968
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertificateRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequestResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateCertificateRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCertificateRequestResponseBody) SetOrderId(v int64) *CreateCertificateRequestResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCertificateRequestResponseBody) SetRequestId(v string) *CreateCertificateRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCertificateRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
