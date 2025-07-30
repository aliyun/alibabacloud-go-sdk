// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificateCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCACertificateCountResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCACertificateCountResponseBody
	GetTotalCount() *int32
}

type DescribeCACertificateCountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of created CA certificates, which includes root CA certificates and intermediate CA certificates.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCACertificateCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCACertificateCountResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCACertificateCountResponseBody) SetRequestId(v string) *DescribeCACertificateCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCACertificateCountResponseBody) SetTotalCount(v int32) *DescribeCACertificateCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCACertificateCountResponseBody) Validate() error {
	return dara.Validate(s)
}
