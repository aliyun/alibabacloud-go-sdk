// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerCertificateWithCsrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterTime(v int64) *CreateServerCertificateWithCsrRequest
	GetAfterTime() *int64
	SetAlgorithm(v string) *CreateServerCertificateWithCsrRequest
	GetAlgorithm() *string
	SetBeforeTime(v int64) *CreateServerCertificateWithCsrRequest
	GetBeforeTime() *int64
	SetCommonName(v string) *CreateServerCertificateWithCsrRequest
	GetCommonName() *string
	SetCountry(v string) *CreateServerCertificateWithCsrRequest
	GetCountry() *string
	SetCsr(v string) *CreateServerCertificateWithCsrRequest
	GetCsr() *string
	SetCustomIdentifier(v string) *CreateServerCertificateWithCsrRequest
	GetCustomIdentifier() *string
	SetDays(v int32) *CreateServerCertificateWithCsrRequest
	GetDays() *int32
	SetDomain(v string) *CreateServerCertificateWithCsrRequest
	GetDomain() *string
	SetEnableCrl(v int64) *CreateServerCertificateWithCsrRequest
	GetEnableCrl() *int64
	SetImmediately(v int32) *CreateServerCertificateWithCsrRequest
	GetImmediately() *int32
	SetLocality(v string) *CreateServerCertificateWithCsrRequest
	GetLocality() *string
	SetMonths(v int32) *CreateServerCertificateWithCsrRequest
	GetMonths() *int32
	SetOrganization(v string) *CreateServerCertificateWithCsrRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateServerCertificateWithCsrRequest
	GetOrganizationUnit() *string
	SetParentIdentifier(v string) *CreateServerCertificateWithCsrRequest
	GetParentIdentifier() *string
	SetResourceGroupId(v string) *CreateServerCertificateWithCsrRequest
	GetResourceGroupId() *string
	SetState(v string) *CreateServerCertificateWithCsrRequest
	GetState() *string
	SetTags(v []*CreateServerCertificateWithCsrRequestTags) *CreateServerCertificateWithCsrRequest
	GetTags() []*CreateServerCertificateWithCsrRequestTags
	SetYears(v int32) *CreateServerCertificateWithCsrRequest
	GetYears() *int32
}

type CreateServerCertificateWithCsrRequest struct {
	// Expiration time of the server-side certificate, in UNIX timestamp format. Unit: seconds.
	//
	// > The **BeforeTime*	- and **AfterTime*	- parameters must both be empty or both configured.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// Key algorithm for the server-side certificate. Use the format `<encryption algorithm>_<key length>`. Valid values:
	//
	// - **RSA_1024**: Signature algorithm is Sha256WithRSA.
	//
	// - **RSA_2048**: Signature algorithm is Sha256WithRSA.
	//
	// - **RSA_4096**: Signature algorithm is Sha256WithRSA.
	//
	// - **ECC_256**: Signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_384**: Signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_512**: Signature algorithm is Sha256WithECDSA.
	//
	// - **SM2_256**: Signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the server-side certificate must match that of the sub-CA certificate. The key length can differ. For example, if the sub-CA certificate uses RSA_2048, the server-side certificate must use RSA_1024, RSA_2048, or RSA_4096.
	//
	// > Call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) to check the key algorithm of the sub-CA certificate.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// Issue time of the server-side certificate, in UNIX timestamp format. Default: current time when you call this API. Unit: seconds.
	//
	// > The **BeforeTime*	- and **AfterTime*	- parameters must both be empty or both configured.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// Set the common name for the certificate. Supports Chinese, English, and other characters.
	//
	// > If you set the **Csr*	- parameter, the value of **CommonName*	- comes from the corresponding field in the **Csr*	- parameter.
	//
	// example:
	//
	// mtcsq.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country code. For example, CN or US.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// You can generate a CSR using OpenSSL or Keytool. For more information, see [How to create a CSR file](https://help.aliyun.com/document_detail/42218.html).
	//
	// <props="china">
	//
	// You can also create a CSR in the SSL Certificate console. For more information, see [Create a CSR](https://help.aliyun.com/document_detail/313297.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----   ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// A custom identifier. This is a unique key.
	//
	// example:
	//
	// ***e6bb538d538c70c01f81hfd3****
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The **Days**, **BeforeTime**, and **AfterTime*	- parameters cannot all be empty. The **BeforeTime*	- and **AfterTime*	- parameters must both be empty or both set. Follow these rules:
	//
	// - If you set **Days**, you can optionally set **BeforeTime*	- and **AfterTime**.
	//
	// - If you do not set **Days**, you must set both **BeforeTime*	- and **AfterTime**.
	//
	// > 	- If you set **Days**, **BeforeTime**, and **AfterTime*	- together, the validity period uses the value of **Days**.
	//
	// - The server-side certificate validity period cannot exceed that of the sub-CA certificate. Call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) to check the sub-CA certificate validity period.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// Additional domain names or IP addresses for the server-side certificate. Adding this information lets you apply the certificate to multiple domains or IP addresses.
	//
	// You can enter multiple domain names and IP addresses. Separate them with commas (,).
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to include the certificate revocation list (CRL) address.
	//
	// 0 - No
	//
	// 1 - Yes
	//
	// example:
	//
	// 1
	EnableCrl *int64 `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// Specifies whether to return the digital certificate immediately.
	//
	// - **0**: Do not return. Default.
	//
	// - **1**: Return the certificate.
	//
	// - **2**: Return the certificate and its certificate chain.
	//
	// example:
	//
	// 1
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The city where the organization for the certificate is located. The name can contain both Chinese and English characters. By default, this parameter is set to the city of the organization for the issuing subordinate Certificate Authority (CA).
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the certificate, in months.
	//
	// example:
	//
	// 12
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	//
	// example:
	//
	// 阿里云
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Alibaba Cloud CDN.
	//
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// Unique identifier of the sub-CA certificate that issues this certificate.
	//
	// > Call [DescribeCACertificateList](https://help.aliyun.com/document_detail/465957.html) to query the unique identifier of the sub-CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 270oe6bb538d538c70c01f81hfd3****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// <props="china">Set the name of the province, municipality, or autonomous region where the organization is located. Supports Chinese, English, and other characters. Defaults to the province, municipality, or autonomous region of the issuing sub-CA certificate\\"s organization.
	//
	// <props="intl">Set the name of the state or province where the organization is located. Supports Chinese, English, and other characters. Defaults to the state or province of the issuing sub-CA certificate\\"s organization.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// A list of tags.
	Tags []*CreateServerCertificateWithCsrRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the certificate, in years.
	//
	// example:
	//
	// 1
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateServerCertificateWithCsrRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServerCertificateWithCsrRequest) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateWithCsrRequest) GetAfterTime() *int64 {
	return s.AfterTime
}

func (s *CreateServerCertificateWithCsrRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateServerCertificateWithCsrRequest) GetBeforeTime() *int64 {
	return s.BeforeTime
}

func (s *CreateServerCertificateWithCsrRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateServerCertificateWithCsrRequest) GetCountry() *string {
	return s.Country
}

func (s *CreateServerCertificateWithCsrRequest) GetCsr() *string {
	return s.Csr
}

func (s *CreateServerCertificateWithCsrRequest) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *CreateServerCertificateWithCsrRequest) GetDays() *int32 {
	return s.Days
}

func (s *CreateServerCertificateWithCsrRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateServerCertificateWithCsrRequest) GetEnableCrl() *int64 {
	return s.EnableCrl
}

func (s *CreateServerCertificateWithCsrRequest) GetImmediately() *int32 {
	return s.Immediately
}

func (s *CreateServerCertificateWithCsrRequest) GetLocality() *string {
	return s.Locality
}

func (s *CreateServerCertificateWithCsrRequest) GetMonths() *int32 {
	return s.Months
}

func (s *CreateServerCertificateWithCsrRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateServerCertificateWithCsrRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateServerCertificateWithCsrRequest) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateServerCertificateWithCsrRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServerCertificateWithCsrRequest) GetState() *string {
	return s.State
}

func (s *CreateServerCertificateWithCsrRequest) GetTags() []*CreateServerCertificateWithCsrRequestTags {
	return s.Tags
}

func (s *CreateServerCertificateWithCsrRequest) GetYears() *int32 {
	return s.Years
}

func (s *CreateServerCertificateWithCsrRequest) SetAfterTime(v int64) *CreateServerCertificateWithCsrRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetAlgorithm(v string) *CreateServerCertificateWithCsrRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetBeforeTime(v int64) *CreateServerCertificateWithCsrRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetCommonName(v string) *CreateServerCertificateWithCsrRequest {
	s.CommonName = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetCountry(v string) *CreateServerCertificateWithCsrRequest {
	s.Country = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetCsr(v string) *CreateServerCertificateWithCsrRequest {
	s.Csr = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetCustomIdentifier(v string) *CreateServerCertificateWithCsrRequest {
	s.CustomIdentifier = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetDays(v int32) *CreateServerCertificateWithCsrRequest {
	s.Days = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetDomain(v string) *CreateServerCertificateWithCsrRequest {
	s.Domain = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetEnableCrl(v int64) *CreateServerCertificateWithCsrRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetImmediately(v int32) *CreateServerCertificateWithCsrRequest {
	s.Immediately = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetLocality(v string) *CreateServerCertificateWithCsrRequest {
	s.Locality = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetMonths(v int32) *CreateServerCertificateWithCsrRequest {
	s.Months = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetOrganization(v string) *CreateServerCertificateWithCsrRequest {
	s.Organization = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetOrganizationUnit(v string) *CreateServerCertificateWithCsrRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetParentIdentifier(v string) *CreateServerCertificateWithCsrRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetResourceGroupId(v string) *CreateServerCertificateWithCsrRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetState(v string) *CreateServerCertificateWithCsrRequest {
	s.State = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetTags(v []*CreateServerCertificateWithCsrRequestTags) *CreateServerCertificateWithCsrRequest {
	s.Tags = v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetYears(v int32) *CreateServerCertificateWithCsrRequest {
	s.Years = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateServerCertificateWithCsrRequestTags struct {
	// Tag key.
	//
	// example:
	//
	// account
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServerCertificateWithCsrRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateServerCertificateWithCsrRequestTags) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateWithCsrRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateServerCertificateWithCsrRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateServerCertificateWithCsrRequestTags) SetKey(v string) *CreateServerCertificateWithCsrRequestTags {
	s.Key = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequestTags) SetValue(v string) *CreateServerCertificateWithCsrRequestTags {
	s.Value = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequestTags) Validate() error {
	return dara.Validate(s)
}
