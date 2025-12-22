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
	// The expiration time of the server certificate. This value is a UNIX timestamp. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the server certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// 	- **RSA_1024**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **RSA_2048**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **RSA_4096**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **ECC_256**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **ECC_384**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **ECC_512**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **SM2_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the server certificate must be the same as the encryption algorithm of the intermediate CA certificate. The key length can be different. For example, if the key algorithm of the intermediate CA certificate is RSA_2048, the key algorithm of the server certificate must be RSA_1024, RSA_2048, or RSA_4096.
	//
	// >  You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the key algorithm of an intermediate CA certificate.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the server certificate. This value is a UNIX timestamp. The default value is the time when you call this operation. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The name of the certificate user. The user of a server certificate is a server. We recommend that you enter the domain name or IP address of the server.
	//
	// example:
	//
	// mtcsq.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located, such as CN or US.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the CSR.
	//
	// You can generate a CSR by using the OpenSSL tool or the Keytool tool. For more information, see [How do I create a CSR file?](https://help.aliyun.com/document_detail/42218.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----   ...... -----END CERTIFICATE REQUEST-----
	Csr              *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The validity period of the server certificate. Unit: days.
	//
	// You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime*	- parameters. The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// 	- If you specify the **Days*	- parameter, you can specify both the **BeforeTime*	- and **AfterTime*	- parameters or leave them both empty.********
	//
	// 	- If you do not specify the **Days*	- parameter, you must specify both the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	// >
	//
	// 	- If you specify the **Days**, **BeforeTime**, and **AfterTime*	- parameters at the same time, the validity period of the server certificate is determined by the value of the **Days*	- parameter.
	//
	// 	- The validity period of the server certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the validity period of an intermediate CA certificate.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The additional domain names or additional IP addresses of the server certificate. After you add additional domain names and additional IP addresses to a certificate, you can apply the certificate to the domain names and IP addresses.
	//
	// You can specify multiple domain names and IP addresses. If you specify multiple domain names and IP addresses, separate them with commas (,).
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// include the CRL address.
	//
	// - 0- No
	//
	// - 1- Yes
	//
	// example:
	//
	// 1
	EnableCrl *int64 `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// Specifies whether to return the certificate. Valid values:
	//
	// 	- **0**: does not return the certificate. This is the default value.
	//
	// 	- **1**: returns the certificate.
	//
	// 	- **2**: returns the certificate and the certificate chain of the certificate.
	//
	// example:
	//
	// 1
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The name of the city in which the organization is located. The value can contain letters. The default value is the name of the city in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the server certificate. Unit: months.
	//
	// example:
	//
	// 12
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	//
	// example:
	//
	// ec server o
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	//
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate CA certificate from which the server certificate is issued.
	//
	// >  You can call the [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) operation to query the unique identifier of an intermediate CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 270oe6bb538d538c70c01f81hfd3****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string                                      `json:"State,omitempty" xml:"State,omitempty"`
	Tags  []*CreateServerCertificateWithCsrRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the server certificate. Unit: years.
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
