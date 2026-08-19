// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificatePackageCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNoticeCountDetail(v string) *GetCertificatePackageCountResponseBody
	GetNoticeCountDetail() *string
	SetProductCountList(v string) *GetCertificatePackageCountResponseBody
	GetProductCountList() *string
	SetProxyCountDetail(v string) *GetCertificatePackageCountResponseBody
	GetProxyCountDetail() *string
	SetRequestId(v string) *GetCertificatePackageCountResponseBody
	GetRequestId() *string
	SetTotalCountDetail(v string) *GetCertificatePackageCountResponseBody
	GetTotalCountDetail() *string
	SetTrusteeCountDetail(v string) *GetCertificatePackageCountResponseBody
	GetTrusteeCountDetail() *string
}

type GetCertificatePackageCountResponseBody struct {
	// example:
	//
	// {TotalCount=189, RemainCount=94, UsedCount=95}
	NoticeCountDetail *string `json:"NoticeCountDetail,omitempty" xml:"NoticeCountDetail,omitempty"`
	// example:
	//
	// [
	//
	//   {
	//
	//     "BrandName": "CFCA",
	//
	//     "TotalCount": 14,
	//
	//     "DomainType": "ONE",
	//
	//     "RemainCount": 14,
	//
	//     "ProductCode": "cfca-ev-1-advanced",
	//
	//     "CertType": "EV",
	//
	//     "ProductId": 8,
	//
	//     "UsedCount": 0
	//
	//   }
	//
	// ]
	ProductCountList *string `json:"ProductCountList,omitempty" xml:"ProductCountList,omitempty"`
	// example:
	//
	// {TotalCount=116900, RemainCount=90448, AutoPay=0, AutoPayCount=5000, UsedCount=26452}
	ProxyCountDetail *string `json:"ProxyCountDetail,omitempty" xml:"ProxyCountDetail,omitempty"`
	// example:
	//
	// 08F45EA0-66A7-4504-9B31-3589F5CE308D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {TotalCount=573, RemainCount=243, FreeQuota=10, UsedCount=330}
	TotalCountDetail *string `json:"TotalCountDetail,omitempty" xml:"TotalCountDetail,omitempty"`
	// example:
	//
	// {TotalCount=177, ValidCount=6, RemainCount=129, UsedCount=48}
	TrusteeCountDetail *string `json:"TrusteeCountDetail,omitempty" xml:"TrusteeCountDetail,omitempty"`
}

func (s GetCertificatePackageCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCertificatePackageCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertificatePackageCountResponseBody) GetNoticeCountDetail() *string {
	return s.NoticeCountDetail
}

func (s *GetCertificatePackageCountResponseBody) GetProductCountList() *string {
	return s.ProductCountList
}

func (s *GetCertificatePackageCountResponseBody) GetProxyCountDetail() *string {
	return s.ProxyCountDetail
}

func (s *GetCertificatePackageCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCertificatePackageCountResponseBody) GetTotalCountDetail() *string {
	return s.TotalCountDetail
}

func (s *GetCertificatePackageCountResponseBody) GetTrusteeCountDetail() *string {
	return s.TrusteeCountDetail
}

func (s *GetCertificatePackageCountResponseBody) SetNoticeCountDetail(v string) *GetCertificatePackageCountResponseBody {
	s.NoticeCountDetail = &v
	return s
}

func (s *GetCertificatePackageCountResponseBody) SetProductCountList(v string) *GetCertificatePackageCountResponseBody {
	s.ProductCountList = &v
	return s
}

func (s *GetCertificatePackageCountResponseBody) SetProxyCountDetail(v string) *GetCertificatePackageCountResponseBody {
	s.ProxyCountDetail = &v
	return s
}

func (s *GetCertificatePackageCountResponseBody) SetRequestId(v string) *GetCertificatePackageCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertificatePackageCountResponseBody) SetTotalCountDetail(v string) *GetCertificatePackageCountResponseBody {
	s.TotalCountDetail = &v
	return s
}

func (s *GetCertificatePackageCountResponseBody) SetTrusteeCountDetail(v string) *GetCertificatePackageCountResponseBody {
	s.TrusteeCountDetail = &v
	return s
}

func (s *GetCertificatePackageCountResponseBody) Validate() error {
	return dara.Validate(s)
}
