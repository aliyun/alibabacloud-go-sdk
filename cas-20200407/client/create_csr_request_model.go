// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCsrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateCsrRequest
	GetAlgorithm() *string
	SetCommonName(v string) *CreateCsrRequest
	GetCommonName() *string
	SetCorpName(v string) *CreateCsrRequest
	GetCorpName() *string
	SetCountryCode(v string) *CreateCsrRequest
	GetCountryCode() *string
	SetDepartment(v string) *CreateCsrRequest
	GetDepartment() *string
	SetKeySize(v int32) *CreateCsrRequest
	GetKeySize() *int32
	SetLocality(v string) *CreateCsrRequest
	GetLocality() *string
	SetName(v string) *CreateCsrRequest
	GetName() *string
	SetProvince(v string) *CreateCsrRequest
	GetProvince() *string
	SetSans(v string) *CreateCsrRequest
	GetSans() *string
}

type CreateCsrRequest struct {
	// The algorithm. Valid values: RSA, SM2, and ECC. For more information about algorithms, see [Select an SSL certificate](https://help.aliyun.com/document_detail/197871.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The primary domain name, which is a common name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The name of the company.
	//
	// example:
	//
	// aly
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	// The code of the country or region in which the organization is located. For example, you can use CN to indicate China and use US to indicate the United States.
	//
	// This parameter is required.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The department that uses the certificate.
	//
	// example:
	//
	// IT
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The key length that is used by the algorithm.
	//
	// 	- The key length for RSA algorithms can be 2,048, 3,072, and 4,096 bits.
	//
	// 	- The key length for ECC and SM2 algorithms can be 256 bits.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The city where the company is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// Beijing
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the CSR. The name can be up to 50 characters in length and can contain letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// example:
	//
	// csr-123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The province or location where the company is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// Beijing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The secondary domain names. Separate multiple domain names with commas (,). You can use the CSR to apply for a certificate for both the primary and secondary domain names.
	//
	// example:
	//
	// www.example.com,www.aliyundoc.com
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
}

func (s CreateCsrRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCsrRequest) GoString() string {
	return s.String()
}

func (s *CreateCsrRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateCsrRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateCsrRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *CreateCsrRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *CreateCsrRequest) GetDepartment() *string {
	return s.Department
}

func (s *CreateCsrRequest) GetKeySize() *int32 {
	return s.KeySize
}

func (s *CreateCsrRequest) GetLocality() *string {
	return s.Locality
}

func (s *CreateCsrRequest) GetName() *string {
	return s.Name
}

func (s *CreateCsrRequest) GetProvince() *string {
	return s.Province
}

func (s *CreateCsrRequest) GetSans() *string {
	return s.Sans
}

func (s *CreateCsrRequest) SetAlgorithm(v string) *CreateCsrRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateCsrRequest) SetCommonName(v string) *CreateCsrRequest {
	s.CommonName = &v
	return s
}

func (s *CreateCsrRequest) SetCorpName(v string) *CreateCsrRequest {
	s.CorpName = &v
	return s
}

func (s *CreateCsrRequest) SetCountryCode(v string) *CreateCsrRequest {
	s.CountryCode = &v
	return s
}

func (s *CreateCsrRequest) SetDepartment(v string) *CreateCsrRequest {
	s.Department = &v
	return s
}

func (s *CreateCsrRequest) SetKeySize(v int32) *CreateCsrRequest {
	s.KeySize = &v
	return s
}

func (s *CreateCsrRequest) SetLocality(v string) *CreateCsrRequest {
	s.Locality = &v
	return s
}

func (s *CreateCsrRequest) SetName(v string) *CreateCsrRequest {
	s.Name = &v
	return s
}

func (s *CreateCsrRequest) SetProvince(v string) *CreateCsrRequest {
	s.Province = &v
	return s
}

func (s *CreateCsrRequest) SetSans(v string) *CreateCsrRequest {
	s.Sans = &v
	return s
}

func (s *CreateCsrRequest) Validate() error {
	return dara.Validate(s)
}
