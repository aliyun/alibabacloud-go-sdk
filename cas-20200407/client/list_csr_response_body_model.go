// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCsrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCsrList(v []*ListCsrResponseBodyCsrList) *ListCsrResponseBody
	GetCsrList() []*ListCsrResponseBodyCsrList
	SetCurrentPage(v int64) *ListCsrResponseBody
	GetCurrentPage() *int64
	SetRequestId(v string) *ListCsrResponseBody
	GetRequestId() *string
	SetShowSize(v int64) *ListCsrResponseBody
	GetShowSize() *int64
	SetTotalCount(v int64) *ListCsrResponseBody
	GetTotalCount() *int64
}

type ListCsrResponseBody struct {
	// The CSRs.
	CsrList []*ListCsrResponseBodyCsrList `json:"CsrList,omitempty" xml:"CsrList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E865F6AD-0294-4A24-A58B-DAC6BE2BDD20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries per page. Default value: 50.
	//
	// example:
	//
	// 10
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCsrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCsrResponseBody) GoString() string {
	return s.String()
}

func (s *ListCsrResponseBody) GetCsrList() []*ListCsrResponseBodyCsrList {
	return s.CsrList
}

func (s *ListCsrResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListCsrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCsrResponseBody) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListCsrResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCsrResponseBody) SetCsrList(v []*ListCsrResponseBodyCsrList) *ListCsrResponseBody {
	s.CsrList = v
	return s
}

func (s *ListCsrResponseBody) SetCurrentPage(v int64) *ListCsrResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCsrResponseBody) SetRequestId(v string) *ListCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCsrResponseBody) SetShowSize(v int64) *ListCsrResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCsrResponseBody) SetTotalCount(v int64) *ListCsrResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCsrResponseBody) Validate() error {
	if s.CsrList != nil {
		for _, item := range s.CsrList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCsrResponseBodyCsrList struct {
	// The algorithm. Valid values: RSA, SM2, and ECC.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The primary domain name, which is a common name.
	//
	// example:
	//
	// example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The name of the company.
	//
	// example:
	//
	// corp_name
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	// The code of the country or region in which the organization is located. For example, you can use CN to indicate China and use US to indicate the United States. The default value is the code of the country or region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued. For more information about country codes, see the "Country codes" section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The ID of the CSR.
	//
	// example:
	//
	// 3454
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	// The department that uses the certificate.
	//
	// example:
	//
	// IT
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// Indicates whether the certificate contains a private key.
	//
	// example:
	//
	// true
	HasPrivateKey *bool `json:"HasPrivateKey,omitempty" xml:"HasPrivateKey,omitempty"`
	// The public key that is calculated by using the SHA256 algorithm.
	//
	// example:
	//
	// 1234
	KeySha2 *string `json:"KeySha2,omitempty" xml:"KeySha2,omitempty"`
	// The key length that is used by the algorithm. The key length for RSA algorithms can be 2,048, 3,072, and 4,096 bits. The key length for ECC and SM2 algorithms can be 256 bits.
	//
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The city where the company is located.
	//
	// example:
	//
	// Beijing
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the CSR.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The province or location.
	//
	// example:
	//
	// Beijing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The secondary domain names. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// www.example.com,www.aliyundoc.com
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
}

func (s ListCsrResponseBodyCsrList) String() string {
	return dara.Prettify(s)
}

func (s ListCsrResponseBodyCsrList) GoString() string {
	return s.String()
}

func (s *ListCsrResponseBodyCsrList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListCsrResponseBodyCsrList) GetCommonName() *string {
	return s.CommonName
}

func (s *ListCsrResponseBodyCsrList) GetCorpName() *string {
	return s.CorpName
}

func (s *ListCsrResponseBodyCsrList) GetCountryCode() *string {
	return s.CountryCode
}

func (s *ListCsrResponseBodyCsrList) GetCsrId() *int64 {
	return s.CsrId
}

func (s *ListCsrResponseBodyCsrList) GetDepartment() *string {
	return s.Department
}

func (s *ListCsrResponseBodyCsrList) GetHasPrivateKey() *bool {
	return s.HasPrivateKey
}

func (s *ListCsrResponseBodyCsrList) GetKeySha2() *string {
	return s.KeySha2
}

func (s *ListCsrResponseBodyCsrList) GetKeySize() *int32 {
	return s.KeySize
}

func (s *ListCsrResponseBodyCsrList) GetLocality() *string {
	return s.Locality
}

func (s *ListCsrResponseBodyCsrList) GetName() *string {
	return s.Name
}

func (s *ListCsrResponseBodyCsrList) GetProvince() *string {
	return s.Province
}

func (s *ListCsrResponseBodyCsrList) GetSans() *string {
	return s.Sans
}

func (s *ListCsrResponseBodyCsrList) SetAlgorithm(v string) *ListCsrResponseBodyCsrList {
	s.Algorithm = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetCommonName(v string) *ListCsrResponseBodyCsrList {
	s.CommonName = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetCorpName(v string) *ListCsrResponseBodyCsrList {
	s.CorpName = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetCountryCode(v string) *ListCsrResponseBodyCsrList {
	s.CountryCode = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetCsrId(v int64) *ListCsrResponseBodyCsrList {
	s.CsrId = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetDepartment(v string) *ListCsrResponseBodyCsrList {
	s.Department = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetHasPrivateKey(v bool) *ListCsrResponseBodyCsrList {
	s.HasPrivateKey = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetKeySha2(v string) *ListCsrResponseBodyCsrList {
	s.KeySha2 = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetKeySize(v int32) *ListCsrResponseBodyCsrList {
	s.KeySize = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetLocality(v string) *ListCsrResponseBodyCsrList {
	s.Locality = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetName(v string) *ListCsrResponseBodyCsrList {
	s.Name = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetProvince(v string) *ListCsrResponseBodyCsrList {
	s.Province = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) SetSans(v string) *ListCsrResponseBodyCsrList {
	s.Sans = &v
	return s
}

func (s *ListCsrResponseBodyCsrList) Validate() error {
	return dara.Validate(s)
}
