// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientCertificateWithCsrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterTime(v int64) *CreateClientCertificateWithCsrRequest
	GetAfterTime() *int64
	SetAlgorithm(v string) *CreateClientCertificateWithCsrRequest
	GetAlgorithm() *string
	SetBeforeTime(v int64) *CreateClientCertificateWithCsrRequest
	GetBeforeTime() *int64
	SetCommonName(v string) *CreateClientCertificateWithCsrRequest
	GetCommonName() *string
	SetCountry(v string) *CreateClientCertificateWithCsrRequest
	GetCountry() *string
	SetCsr(v string) *CreateClientCertificateWithCsrRequest
	GetCsr() *string
	SetCustomIdentifier(v string) *CreateClientCertificateWithCsrRequest
	GetCustomIdentifier() *string
	SetDays(v int32) *CreateClientCertificateWithCsrRequest
	GetDays() *int32
	SetEnableCrl(v int64) *CreateClientCertificateWithCsrRequest
	GetEnableCrl() *int64
	SetImmediately(v int32) *CreateClientCertificateWithCsrRequest
	GetImmediately() *int32
	SetLocality(v string) *CreateClientCertificateWithCsrRequest
	GetLocality() *string
	SetMonths(v int32) *CreateClientCertificateWithCsrRequest
	GetMonths() *int32
	SetOrganization(v string) *CreateClientCertificateWithCsrRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateClientCertificateWithCsrRequest
	GetOrganizationUnit() *string
	SetParentIdentifier(v string) *CreateClientCertificateWithCsrRequest
	GetParentIdentifier() *string
	SetResourceGroupId(v string) *CreateClientCertificateWithCsrRequest
	GetResourceGroupId() *string
	SetSanType(v int32) *CreateClientCertificateWithCsrRequest
	GetSanType() *int32
	SetSanValue(v string) *CreateClientCertificateWithCsrRequest
	GetSanValue() *string
	SetState(v string) *CreateClientCertificateWithCsrRequest
	GetState() *string
	SetTags(v []*CreateClientCertificateWithCsrRequestTags) *CreateClientCertificateWithCsrRequest
	GetTags() []*CreateClientCertificateWithCsrRequestTags
	SetYears(v int32) *CreateClientCertificateWithCsrRequest
	GetYears() *int32
}

type CreateClientCertificateWithCsrRequest struct {
	// The expiration time of the client certificate. This is a UNIX timestamp in seconds.
	//
	// > Specify the **BeforeTime*	- and **AfterTime*	- parameters together, or omit both.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the client certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// - **RSA_1024**: The corresponding signature algorithm is Sha256WithRSA.
	//
	// - **RSA_2048**: The corresponding signature algorithm is Sha256WithRSA.
	//
	// - **RSA_4096**: The corresponding signature algorithm is Sha256WithRSA.
	//
	// - **ECC_256**: The corresponding signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_384**: The corresponding signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_512**: The corresponding signature algorithm is Sha256WithECDSA.
	//
	// - **SM2_256**: The corresponding signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the client certificate must be the same as that of the subordinate CA certificate, but the key length can be different. For example, if the key algorithm of the subordinate CA certificate is RSA_2048, the key algorithm of the client certificate must be one of RSA_1024, RSA_2048, and RSA_4096.
	//
	// > Call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) to query the key algorithm of the subordinate CA certificate.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the client certificate. This is a UNIX timestamp in seconds. The default value is the time of the API call.
	//
	// > The **BeforeTime*	- and **AfterTime*	- parameters must be specified together or left empty.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The common name of the certificate. Chinese and English characters are supported.
	//
	// > If you specify the **Csr*	- parameter, the value of this parameter is determined by the information in the **Csr*	- parameter.
	//
	// example:
	//
	// aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country code, for example, **CN*	- or **US**.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the CSR. Use OpenSSL or Keytool to generate a CSR. For more information, see [Create a CSR file](https://help.aliyun.com/document_detail/42218.html).
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----   ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// A custom identifier. This is a unique key.
	//
	// example:
	//
	// ***e6bb538d538c70c01f81fg3****
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The validity period of the client certificate, in days. You must specify the validity period using one of the following methods:
	//
	// - Specify the **Days*	- parameter.
	//
	// - Specify both the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	// > 	- If you specify **Days**, **BeforeTime**, and **AfterTime*	- at the same time, the value of **Days*	- is used.
	//
	// - The validity period of the client certificate cannot exceed that of the subordinate CA certificate. Call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) to view the validity period of the subordinate CA certificate.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// Specifies whether to include the Certificate Revocation List (CRL) address.
	//
	// 0: No
	//
	// 1: Yes
	//
	// example:
	//
	// 1
	EnableCrl *int64 `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// Specifies whether to return the digital certificate.
	//
	// - **0**: Do not return the certificate. This is the default value.
	//
	// - **1**: Return the certificate.
	//
	// - **2**: Return the certificate and its certificate chain.
	//
	// example:
	//
	// 1
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The name of the city where the organization is located. Chinese and English characters are supported. By default, this parameter uses the city name of the organization that is associated with the issuing subordinate CA certificate.
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
	// Alibaba Inc
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Alibaba Cloud CDN.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the subordinate CA certificate that issues the client certificate.
	//
	// > Call [DescribeCACertificateList](https://help.aliyun.com/document_detail/465957.html) to query the unique identifiers of subordinate CA certificates.
	//
	// example:
	//
	// 270ae6bb538d538c70c01f81fg3****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The ID of the resource group to which the certificate belongs.
	//
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the Subject Alternative Name (SAN) extension for the client certificate. Valid values:
	//
	// - **1**: Email address.
	//
	// - **6**: Uniform Resource Identifier (URI).
	//
	// example:
	//
	// 2
	SanType *int32 `json:"SanType,omitempty" xml:"SanType,omitempty"`
	// The extension for the client certificate. To specify multiple extensions, separate them with a comma.
	//
	// example:
	//
	// somebody@example.com
	SanValue *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	// Specify the name of the province or state where the certificate organization is located. The value can contain letters. The default value is the name of the province or state of the intermediate CA\\"s organization.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// A list of tags.
	Tags []*CreateClientCertificateWithCsrRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the certificate, in years.
	//
	// example:
	//
	// 1
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateClientCertificateWithCsrRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateWithCsrRequest) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateWithCsrRequest) GetAfterTime() *int64 {
	return s.AfterTime
}

func (s *CreateClientCertificateWithCsrRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateClientCertificateWithCsrRequest) GetBeforeTime() *int64 {
	return s.BeforeTime
}

func (s *CreateClientCertificateWithCsrRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateClientCertificateWithCsrRequest) GetCountry() *string {
	return s.Country
}

func (s *CreateClientCertificateWithCsrRequest) GetCsr() *string {
	return s.Csr
}

func (s *CreateClientCertificateWithCsrRequest) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *CreateClientCertificateWithCsrRequest) GetDays() *int32 {
	return s.Days
}

func (s *CreateClientCertificateWithCsrRequest) GetEnableCrl() *int64 {
	return s.EnableCrl
}

func (s *CreateClientCertificateWithCsrRequest) GetImmediately() *int32 {
	return s.Immediately
}

func (s *CreateClientCertificateWithCsrRequest) GetLocality() *string {
	return s.Locality
}

func (s *CreateClientCertificateWithCsrRequest) GetMonths() *int32 {
	return s.Months
}

func (s *CreateClientCertificateWithCsrRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateClientCertificateWithCsrRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateClientCertificateWithCsrRequest) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateClientCertificateWithCsrRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClientCertificateWithCsrRequest) GetSanType() *int32 {
	return s.SanType
}

func (s *CreateClientCertificateWithCsrRequest) GetSanValue() *string {
	return s.SanValue
}

func (s *CreateClientCertificateWithCsrRequest) GetState() *string {
	return s.State
}

func (s *CreateClientCertificateWithCsrRequest) GetTags() []*CreateClientCertificateWithCsrRequestTags {
	return s.Tags
}

func (s *CreateClientCertificateWithCsrRequest) GetYears() *int32 {
	return s.Years
}

func (s *CreateClientCertificateWithCsrRequest) SetAfterTime(v int64) *CreateClientCertificateWithCsrRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetAlgorithm(v string) *CreateClientCertificateWithCsrRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetBeforeTime(v int64) *CreateClientCertificateWithCsrRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetCommonName(v string) *CreateClientCertificateWithCsrRequest {
	s.CommonName = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetCountry(v string) *CreateClientCertificateWithCsrRequest {
	s.Country = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetCsr(v string) *CreateClientCertificateWithCsrRequest {
	s.Csr = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetCustomIdentifier(v string) *CreateClientCertificateWithCsrRequest {
	s.CustomIdentifier = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetDays(v int32) *CreateClientCertificateWithCsrRequest {
	s.Days = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetEnableCrl(v int64) *CreateClientCertificateWithCsrRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetImmediately(v int32) *CreateClientCertificateWithCsrRequest {
	s.Immediately = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetLocality(v string) *CreateClientCertificateWithCsrRequest {
	s.Locality = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetMonths(v int32) *CreateClientCertificateWithCsrRequest {
	s.Months = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetOrganization(v string) *CreateClientCertificateWithCsrRequest {
	s.Organization = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetOrganizationUnit(v string) *CreateClientCertificateWithCsrRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetParentIdentifier(v string) *CreateClientCertificateWithCsrRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetResourceGroupId(v string) *CreateClientCertificateWithCsrRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetSanType(v int32) *CreateClientCertificateWithCsrRequest {
	s.SanType = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetSanValue(v string) *CreateClientCertificateWithCsrRequest {
	s.SanValue = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetState(v string) *CreateClientCertificateWithCsrRequest {
	s.State = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetTags(v []*CreateClientCertificateWithCsrRequestTags) *CreateClientCertificateWithCsrRequest {
	s.Tags = v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetYears(v int32) *CreateClientCertificateWithCsrRequest {
	s.Years = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) Validate() error {
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

type CreateClientCertificateWithCsrRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// database
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClientCertificateWithCsrRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateWithCsrRequestTags) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateWithCsrRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateClientCertificateWithCsrRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateClientCertificateWithCsrRequestTags) SetKey(v string) *CreateClientCertificateWithCsrRequestTags {
	s.Key = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequestTags) SetValue(v string) *CreateClientCertificateWithCsrRequestTags {
	s.Value = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequestTags) Validate() error {
	return dara.Validate(s)
}
