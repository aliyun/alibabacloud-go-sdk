// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertList(v []*ListCertResponseBodyCertList) *ListCertResponseBody
	GetCertList() []*ListCertResponseBodyCertList
	SetCurrentPage(v int64) *ListCertResponseBody
	GetCurrentPage() *int64
	SetRequestId(v string) *ListCertResponseBody
	GetRequestId() *string
	SetShowSize(v int64) *ListCertResponseBody
	GetShowSize() *int64
	SetTotalCount(v int64) *ListCertResponseBody
	GetTotalCount() *int64
}

type ListCertResponseBody struct {
	// An array that consists of the certificates.
	CertList []*ListCertResponseBodyCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	// The page number of the returned page. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned per page. Default value: 50.
	//
	// example:
	//
	// 50
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCertResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertResponseBody) GetCertList() []*ListCertResponseBodyCertList {
	return s.CertList
}

func (s *ListCertResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCertResponseBody) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListCertResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCertResponseBody) SetCertList(v []*ListCertResponseBodyCertList) *ListCertResponseBody {
	s.CertList = v
	return s
}

func (s *ListCertResponseBody) SetCurrentPage(v int64) *ListCertResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCertResponseBody) SetRequestId(v string) *ListCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertResponseBody) SetShowSize(v int64) *ListCertResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCertResponseBody) SetTotalCount(v int64) *ListCertResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCertResponseBody) Validate() error {
	if s.CertList != nil {
		for _, item := range s.CertList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCertResponseBodyCertList struct {
	// The expiration time of the certificate. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1634283958000
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The issuance time of the certificate. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// 证书的类型 。取值：
	//
	// - **CA**：表示CA证书。
	//
	// - **CERT**：表示签发的证书。
	//
	// example:
	//
	// CERT
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The domain name.
	//
	// example:
	//
	// aliyun.alibaba.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// Indicates whether the certificate contains a private key. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ExistPrivateKey *bool `json:"ExistPrivateKey,omitempty" xml:"ExistPrivateKey,omitempty"`
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 14dcc8afc7578e
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The issuer of the certificate.
	//
	// example:
	//
	// mySSL
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The domain names that are bound to the certificate. Multiple domain names are separated by commas.
	//
	// example:
	//
	// *.alibaba.com,aliyun.alibaba.com
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The source of the certificate. Valid values:
	//
	// 	- **upload**: uploaded certificate
	//
	// 	- **aliyun**: Alibaba Cloud certificate
	//
	// example:
	//
	// aliyun
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- **ISSUE**: issued
	//
	// 	- **REVOKE**: revoked
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the certificate repository.
	//
	// example:
	//
	// 2
	WhId *int64 `json:"WhId,omitempty" xml:"WhId,omitempty"`
	// The instance ID of the certificate repository.
	//
	// example:
	//
	// test_whInstanceId
	WhInstanceId *string `json:"WhInstanceId,omitempty" xml:"WhInstanceId,omitempty"`
}

func (s ListCertResponseBodyCertList) String() string {
	return dara.Prettify(s)
}

func (s ListCertResponseBodyCertList) GoString() string {
	return s.String()
}

func (s *ListCertResponseBodyCertList) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *ListCertResponseBodyCertList) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *ListCertResponseBodyCertList) GetCertType() *string {
	return s.CertType
}

func (s *ListCertResponseBodyCertList) GetCommonName() *string {
	return s.CommonName
}

func (s *ListCertResponseBodyCertList) GetExistPrivateKey() *bool {
	return s.ExistPrivateKey
}

func (s *ListCertResponseBodyCertList) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListCertResponseBodyCertList) GetIssuer() *string {
	return s.Issuer
}

func (s *ListCertResponseBodyCertList) GetSans() *string {
	return s.Sans
}

func (s *ListCertResponseBodyCertList) GetSourceType() *string {
	return s.SourceType
}

func (s *ListCertResponseBodyCertList) GetStatus() *string {
	return s.Status
}

func (s *ListCertResponseBodyCertList) GetWhId() *int64 {
	return s.WhId
}

func (s *ListCertResponseBodyCertList) GetWhInstanceId() *string {
	return s.WhInstanceId
}

func (s *ListCertResponseBodyCertList) SetAfterDate(v int64) *ListCertResponseBodyCertList {
	s.AfterDate = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetBeforeDate(v int64) *ListCertResponseBodyCertList {
	s.BeforeDate = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetCertType(v string) *ListCertResponseBodyCertList {
	s.CertType = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetCommonName(v string) *ListCertResponseBodyCertList {
	s.CommonName = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetExistPrivateKey(v bool) *ListCertResponseBodyCertList {
	s.ExistPrivateKey = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetIdentifier(v string) *ListCertResponseBodyCertList {
	s.Identifier = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetIssuer(v string) *ListCertResponseBodyCertList {
	s.Issuer = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetSans(v string) *ListCertResponseBodyCertList {
	s.Sans = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetSourceType(v string) *ListCertResponseBodyCertList {
	s.SourceType = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetStatus(v string) *ListCertResponseBodyCertList {
	s.Status = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetWhId(v int64) *ListCertResponseBodyCertList {
	s.WhId = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetWhInstanceId(v string) *ListCertResponseBodyCertList {
	s.WhInstanceId = &v
	return s
}

func (s *ListCertResponseBodyCertList) Validate() error {
	return dara.Validate(s)
}
