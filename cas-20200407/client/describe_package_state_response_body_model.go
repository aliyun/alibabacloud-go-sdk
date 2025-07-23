// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackageStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIssuedCount(v int64) *DescribePackageStateResponseBody
	GetIssuedCount() *int64
	SetProductCode(v string) *DescribePackageStateResponseBody
	GetProductCode() *string
	SetRequestId(v string) *DescribePackageStateResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePackageStateResponseBody
	GetTotalCount() *int64
	SetUsedCount(v int64) *DescribePackageStateResponseBody
	GetUsedCount() *int64
}

type DescribePackageStateResponseBody struct {
	// The number of issued certificates of the specified specifications.
	//
	// example:
	//
	// 1
	IssuedCount *int64 `json:"IssuedCount,omitempty" xml:"IssuedCount,omitempty"`
	// The specifications of the certificate resource plan. Valid values:
	//
	// 	- **digicert-free-1-free**: DigiCert single-domain DV certificate in a three-month free trial, available only on the China site (aliyun.com).
	//
	// 	- **symantec-free-1-free**: DigiCert single-domain DV certificate in a one-year free trial, available only on the China site (aliyun.com).
	//
	// 	- **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	//
	// 	- **symantec-ov-1-personal**: DigiCert single-domain OV certificate.
	//
	// 	- **symantec-ov-w-personal**: DigiCert wildcard OV certificate.
	//
	// 	- **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	//
	// 	- **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	//
	// 	- **geotrust-ov-1-personal**: GeoTrust single-domain OV certificate.
	//
	// 	- **geotrust-ov-w-personal**: GeoTrust wildcard OV certificate.
	//
	// 	- **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	//
	// 	- **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	//
	// 	- **globalsign-ov-1-personal**: GlobalSign single-domain OV certificate.
	//
	// 	- **globalsign-ov-w-advanced**: GlobalSign wildcard OV certificate.
	//
	// 	- **cfca-ov-1-personal**: CFCA single-domain OV certificate, available only on the China site (aliyun.com).
	//
	// 	- **cfca-ev-w-advanced**: CFCA wildcard OV certificate, available only on the China site (aliyun.com).
	//
	// example:
	//
	// symantec-free-1-free
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 10CFA380-1C58-45C7-8075-06215F3DB681
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of purchased certificate resource plans of the specified specifications.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of certificate applications that you submitted for certificates of the specified specifications.
	//
	// > : A successful call of the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/204087.html), [CreateCertificateRequest](https://help.aliyun.com/document_detail/164105.html), or [CreateCertificateWithCsrRequest](https://help.aliyun.com/document_detail/178732.html) operation is counted as one a certificate application, regardless of whether the certificate is issued.
	//
	// example:
	//
	// 2
	UsedCount *int64 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribePackageStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageStateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackageStateResponseBody) GetIssuedCount() *int64 {
	return s.IssuedCount
}

func (s *DescribePackageStateResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribePackageStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePackageStateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePackageStateResponseBody) GetUsedCount() *int64 {
	return s.UsedCount
}

func (s *DescribePackageStateResponseBody) SetIssuedCount(v int64) *DescribePackageStateResponseBody {
	s.IssuedCount = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetProductCode(v string) *DescribePackageStateResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetRequestId(v string) *DescribePackageStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetTotalCount(v int64) *DescribePackageStateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetUsedCount(v int64) *DescribePackageStateResponseBody {
	s.UsedCount = &v
	return s
}

func (s *DescribePackageStateResponseBody) Validate() error {
	return dara.Validate(s)
}
