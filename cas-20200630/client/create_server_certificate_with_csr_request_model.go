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
	SetAsynchronousFlag(v bool) *CreateServerCertificateWithCsrRequest
	GetAsynchronousFlag() *bool
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
	// The expiration time of the server certificate in UNIX timestamp format. Unit: seconds.
	//
	// >The **BeforeTime*	- and **AfterTime*	- parameters must both be empty or both be specified.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the server certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// - **RSA_1024**: The signature algorithm is Sha256WithRSA.
	//
	// - **RSA_2048**: The signature algorithm is Sha256WithRSA.
	//
	// - **RSA_4096**: The signature algorithm is Sha256WithRSA.
	//
	// - **ECC_256**: The signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_384**: The signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_512**: The signature algorithm is Sha256WithECDSA.
	//
	// - **SM2_256**: The signature algorithm is SM3WithSM2.
	//
	//
	// The encryption algorithm of the server certificate must be the same as that of the subordinate CA certificate, but the key length can be different. For example, if the key algorithm of the subordinate CA certificate is RSA_2048, the key algorithm of the server certificate must be RSA_1024, RSA_2048, or RSA_4096.
	//
	// >You can call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) to query the key algorithm of the subordinate CA certificate.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// false
	AsynchronousFlag *bool `json:"AsynchronousFlag,omitempty" xml:"AsynchronousFlag,omitempty"`
	// The issuance time of the server certificate in UNIX timestamp format. The default value is the time when you call this operation. Unit: seconds.
	//
	// >The **BeforeTime*	- and **AfterTime*	- parameters must both be empty or both be specified.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The common name of the certificate. Chinese characters, English characters, and other characters are supported.
	//
	// >If you set the **Csr*	- parameter, the value of the **CommonName*	- parameter is determined by the corresponding information in the **Csr*	- parameter.
	//
	// example:
	//
	// mtcsq.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country code, such as **CN**.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The CSR content.
	//
	// You can use OpenSSL or Keytool to generate a CSR. For more information, see [How do I create a CSR file?](https://help.aliyun.com/document_detail/42218.html).
	//
	// <props="china">You can also create a CSR in the SSL Certificates Service console. For more information, see [Create a CSR](https://help.aliyun.com/document_detail/313297.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----   ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The user-defined identifier, which serves as a unique key.
	//
	// example:
	//
	// ***e6bb538d538c70c01f81hfd3****
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The validity period of the server certificate. Unit: days.
	//
	// The **Days**, **BeforeTime**, and **AfterTime*	- parameters cannot all be empty. The **BeforeTime*	- and **AfterTime*	- parameters must both be empty or both be specified. The following rules apply:
	//
	// - If you set the **Days*	- parameter, you can choose to set or not set the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	//
	// - If you do not set the **Days*	- parameter, you must set the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	// >- If you set the **Days**, **BeforeTime**, and **AfterTime*	- parameters at the same time, the validity period of the server certificate is determined by the value of the **Days*	- parameter.
	//
	// - The validity period of the server certificate cannot exceed the validity period of the subordinate CA certificate. You can call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) to query the validity period of the subordinate CA certificate.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The extended domain name or extended IP address of the server certificate. After you add extended information to the certificate, you can apply the certificate to multiple domain names or IP addresses.
	//
	// You can enter multiple domain names and IP addresses at the same time. Separate multiple values with commas (,).
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to include the CRL address. Valid values:
	//
	// - 0: No.
	//
	// - 1: Yes.
	//
	// example:
	//
	// 1
	EnableCrl *int64 `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// Specifies whether to immediately return the digital certificate. Valid values:
	//
	// - **0**: Does not return the certificate. This is the default value.
	//
	// - **1**: Returns the certificate.
	//
	// - **2**: Returns the certificate and its certificate chain.
	//
	// example:
	//
	// 1
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The name of the city where the certificate organization is located. Chinese characters, English characters, and other characters are supported.
	//
	// The default value is the name of the city where the organization of the subordinate CA certificate that issues this certificate is located.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The certificate validity period. Unit: months.
	//
	// example:
	//
	// 12
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The organization name. Default value: Alibaba Inc.
	//
	// example:
	//
	// Alibaba Cloud
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The department name. Default value: Aliyun CDN.
	//
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the subordinate CA certificate that issues this certificate.
	//
	// >You can call [DescribeCACertificateList](https://help.aliyun.com/document_detail/465957.html) to query the unique identifier of the subordinate CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 270oe6bb538d538c70c01f81hfd3****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// <props="china">The name of the province, municipality, or autonomous region where the certificate organization is located. Chinese characters, English characters, and other characters are supported. The default value is the name of the province, municipality, or autonomous region where the organization of the subordinate CA certificate that issues this certificate is located.
	//
	// <props="intl">The name of the province or state where the certificate organization is located. Chinese characters, English characters, and other characters are supported. The default value is the name of the province or state where the organization of the subordinate CA certificate that issues this certificate is located.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tag list.
	Tags []*CreateServerCertificateWithCsrRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The certificate validity period. Unit: years.
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

func (s *CreateServerCertificateWithCsrRequest) GetAsynchronousFlag() *bool {
	return s.AsynchronousFlag
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

func (s *CreateServerCertificateWithCsrRequest) SetAsynchronousFlag(v bool) *CreateServerCertificateWithCsrRequest {
	s.AsynchronousFlag = &v
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
	// The tag key.
	//
	// example:
	//
	// account
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
