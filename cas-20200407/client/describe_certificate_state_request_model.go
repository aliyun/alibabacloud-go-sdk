// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificateStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *DescribeCertificateStateRequest
	GetOrderId() *int64
}

type DescribeCertificateStateRequest struct {
	// The ID of the certificate application order that you want to query.
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

func (s DescribeCertificateStateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateStateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificateStateRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeCertificateStateRequest) SetOrderId(v int64) *DescribeCertificateStateRequest {
	s.OrderId = &v
	return s
}

func (s *DescribeCertificateStateRequest) Validate() error {
	return dara.Validate(s)
}
