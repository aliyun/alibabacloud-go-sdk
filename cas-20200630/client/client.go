// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateClientCertificateRequest struct {
	// The expiration time of the client certificate. This value is a UNIX timestamp. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the client certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
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
	// The encryption algorithm of the client certificate must be the same with the encryption algorithm of the intermediate certificate authority (CA) certificate. The key length can be different. For example, if the key algorithm of the intermediate CA certificate is RSA_2048, the key algorithm of the client certificate must be RSA_1024, RSA_2048, or RSA_4096.
	//
	// > You can call the [DescribeCACertificate] operation to query the key algorithm of an intermediate CA certificate.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the client certificate. This value is a UNIX timestamp. The default value is the time when you call this operation. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The name of the client certificate user. In most cases, the user of a client certificate is an individual, a company, an organization, or an application. We recommend that you enter the common name of a user. Examples: Bob, Alibaba, Alibaba Cloud password platform, and Tmall Genie.
	//
	// example:
	//
	// aliyun
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country in which the organization is located. Default value: CN.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The validity period of the client certificate. Unit: day. You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime*	- parameters. The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// 	- If you specify the **Days*	- parameter, you can specify both the **BeforeTime*	- and **AfterTime*	- parameters or leave them both empty.
	//
	// 	- If you do not specify the **Days*	- parameter, you must specify both the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	// >
	//
	// 	- If you specify the **Days**, **BeforeTime**, and **AfterTime*	- parameters at the same time, the validity period of the client certificate is determined by the value of the **Days*	- parameter.
	//
	// 	- The validity period of the client certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) operation to query the validity period of an intermediate CA certificate.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
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
	// The validity period of the client certificate. Unit: months.
	//
	// example:
	//
	// 1
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	//
	// example:
	//
	// Alibaba Cloud
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	//
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate CA certificate from which the server certificate is issued.
	//
	// > You can call the [DescribeCACertificateList] operation to query the unique identifier of an intermediate CA certificate.
	//
	// example:
	//
	// 273ae6bb538d538c70c01f81jh2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The type of the Subject Alternative Name (SAN) extension that is supported by the client certificate. Valid values:
	//
	// 	- **1**: an email address
	//
	// 	- **6**: a Uniform Resource Identifier (URI)
	//
	// example:
	//
	// 1
	SanType *int32 `json:"SanType,omitempty" xml:"SanType,omitempty"`
	// The content of the extension. You can specify multiple SAN extensions. If you want to specify multiple SAN extensions, separate them with commas (,).
	//
	// example:
	//
	// somebody@example.com
	SanValue *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the client certificate. Unit: years.
	//
	// example:
	//
	// 5
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateClientCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateRequest) SetAfterTime(v int64) *CreateClientCertificateRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateClientCertificateRequest) SetAlgorithm(v string) *CreateClientCertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateClientCertificateRequest) SetBeforeTime(v int64) *CreateClientCertificateRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateClientCertificateRequest) SetCommonName(v string) *CreateClientCertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateClientCertificateRequest) SetCountry(v string) *CreateClientCertificateRequest {
	s.Country = &v
	return s
}

func (s *CreateClientCertificateRequest) SetDays(v int32) *CreateClientCertificateRequest {
	s.Days = &v
	return s
}

func (s *CreateClientCertificateRequest) SetEnableCrl(v int64) *CreateClientCertificateRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateClientCertificateRequest) SetImmediately(v int32) *CreateClientCertificateRequest {
	s.Immediately = &v
	return s
}

func (s *CreateClientCertificateRequest) SetLocality(v string) *CreateClientCertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateClientCertificateRequest) SetMonths(v int32) *CreateClientCertificateRequest {
	s.Months = &v
	return s
}

func (s *CreateClientCertificateRequest) SetOrganization(v string) *CreateClientCertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateClientCertificateRequest) SetOrganizationUnit(v string) *CreateClientCertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateClientCertificateRequest) SetParentIdentifier(v string) *CreateClientCertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateClientCertificateRequest) SetSanType(v int32) *CreateClientCertificateRequest {
	s.SanType = &v
	return s
}

func (s *CreateClientCertificateRequest) SetSanValue(v string) *CreateClientCertificateRequest {
	s.SanValue = &v
	return s
}

func (s *CreateClientCertificateRequest) SetState(v string) *CreateClientCertificateRequest {
	s.State = &v
	return s
}

func (s *CreateClientCertificateRequest) SetYears(v int32) *CreateClientCertificateRequest {
	s.Years = &v
	return s
}

type CreateClientCertificateResponseBody struct {
	// The certificate chain of the client certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the client certificate.
	//
	// example:
	//
	// 190ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 0f29522da2dae7a1c4b6ab7132ad3c06
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the client certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s CreateClientCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateResponseBody) SetCertificateChain(v string) *CreateClientCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetIdentifier(v string) *CreateClientCertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetRequestId(v string) *CreateClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetSerialNumber(v string) *CreateClientCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetX509Certificate(v string) *CreateClientCertificateResponseBody {
	s.X509Certificate = &v
	return s
}

type CreateClientCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateResponse) SetHeaders(v map[string]*string) *CreateClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateClientCertificateResponse) SetStatusCode(v int32) *CreateClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientCertificateResponse) SetBody(v *CreateClientCertificateResponseBody) *CreateClientCertificateResponse {
	s.Body = v
	return s
}

type CreateClientCertificateWithCsrRequest struct {
	// The expiration time of the client certificate. This value is a UNIX timestamp. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the client certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
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
	// The encryption algorithm of the client certificate must be the same with the encryption algorithm of the intermediate CA certificate. The key length can be different. For example, if the key algorithm of the intermediate CA certificate is RSA_2048, the key algorithm of the client certificate must be RSA_1024, RSA_2048, or RSA_4096.
	//
	// >  You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the key algorithm of an intermediate CA certificate.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the client certificate. This value is a UNIX timestamp. The default value is the time when you call this operation. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The common name of the certificate. The value can contain letters.
	//
	// >  If you specify the **CsrPemString*	- parameter, the value of the **CommonName*	- parameter is determined by the **CsrPemString*	- parameter.
	//
	// example:
	//
	// aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located, such as **CN*	- and **US**.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the CSR file. You can generate a CSR file by using the OpenSSL tool or Keytool. For more information, see [How do I create a CSR file?](https://help.aliyun.com/document_detail/42218.html) You can also create a CSR file in the Certificate Management Service console. For more information, see [Create a CSR](https://help.aliyun.com/document_detail/313297.html).
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----   ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The validity period of the client certificate. Unit: days. You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime*	- parameters. The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// 	- If you specify the **Days*	- parameter, you can specify both the **BeforeTime*	- and **AfterTime*	- parameters or leave them both empty.********
	//
	// 	- If you do not specify the **Days*	- parameter, you must specify both the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	// >
	//
	// 	- If you specify the **Days**, **BeforeTime**, and **AfterTime*	- parameters together, the validity period of the client certificate is determined by the value of the **Days*	- parameter.
	//
	// 	- The validity period of the client certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the validity period of an intermediate CA certificate.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
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
	// The validity period of the client certificate. Unit: months.
	//
	// example:
	//
	// 12
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	//
	// example:
	//
	// Alibaba Cloud Computing Co., Ltd.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate CA certificate from which the client certificate is issued.
	//
	// >  You can call the [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) operation to query the unique identifier of an intermediate CA certificate.
	//
	// example:
	//
	// 270ae6bb538d538c70c01f81fg3****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The type of the Subject Alternative Name (SAN) extension that is supported by the client certificate. Valid values:
	//
	// 	- **1**: an email address
	//
	// 	- **6**: a Uniform Resource Identifier (URI)
	//
	// example:
	//
	// 1
	SanType *int32 `json:"SanType,omitempty" xml:"SanType,omitempty"`
	// The content of the extension. You can specify multiple SAN extensions. If you want to specify multiple SAN extensions, separate them with commas (,).
	//
	// example:
	//
	// somebody@example.com
	SanValue *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the client certificate. Unit: years.
	//
	// example:
	//
	// 1
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateClientCertificateWithCsrRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClientCertificateWithCsrRequest) GoString() string {
	return s.String()
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

func (s *CreateClientCertificateWithCsrRequest) SetYears(v int32) *CreateClientCertificateWithCsrRequest {
	s.Years = &v
	return s
}

type CreateClientCertificateWithCsrResponseBody struct {
	// CertKmcRep1.
	//
	// example:
	//
	// userSeal=MHkCIEu94PQAahFWuFk%
	//
	// ***
	//
	// EtFw%2FkMMBjw8i5bFfSkV%2FIUrcOJD
	CertKmcRep1 *string `json:"CertKmcRep1,omitempty" xml:"CertKmcRep1,omitempty"`
	// Cert Sign Buf Kmc.
	//
	// example:
	//
	// userSeal=MHkCIEu94PQAahFWuFk%
	//
	// ***
	//
	// EtFw%2FkMMBjw8i5bFfSkV%2FIUrcOJD
	CertSignBufKmc *string `json:"CertSignBufKmc,omitempty" xml:"CertSignBufKmc,omitempty"`
	// The certificate chain of the client certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the client certificate.
	//
	// example:
	//
	// 200ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 31C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the server certificate.
	//
	// example:
	//
	// 0f29522da2dae7a1c4b6ab7132ad3c06
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the client certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s CreateClientCertificateWithCsrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClientCertificateWithCsrResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateWithCsrResponseBody) SetCertKmcRep1(v string) *CreateClientCertificateWithCsrResponseBody {
	s.CertKmcRep1 = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetCertSignBufKmc(v string) *CreateClientCertificateWithCsrResponseBody {
	s.CertSignBufKmc = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetCertificateChain(v string) *CreateClientCertificateWithCsrResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetIdentifier(v string) *CreateClientCertificateWithCsrResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetRequestId(v string) *CreateClientCertificateWithCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetSerialNumber(v string) *CreateClientCertificateWithCsrResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetX509Certificate(v string) *CreateClientCertificateWithCsrResponseBody {
	s.X509Certificate = &v
	return s
}

type CreateClientCertificateWithCsrResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientCertificateWithCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientCertificateWithCsrResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClientCertificateWithCsrResponse) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateWithCsrResponse) SetHeaders(v map[string]*string) *CreateClientCertificateWithCsrResponse {
	s.Headers = v
	return s
}

func (s *CreateClientCertificateWithCsrResponse) SetStatusCode(v int32) *CreateClientCertificateWithCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponse) SetBody(v *CreateClientCertificateWithCsrResponseBody) *CreateClientCertificateWithCsrResponse {
	s.Body = v
	return s
}

type CreateCustomCertificateRequest struct {
	// The passthrough parameters.
	ApiPassthrough *CreateCustomCertificateRequestApiPassthrough `json:"ApiPassthrough,omitempty" xml:"ApiPassthrough,omitempty" type:"Struct"`
	// The content of the CSR. You can generate a CSR by using the OpenSSL tool or the Keytool tool. For more information, see [How do I create a CSR file?](https://help.aliyun.com/document_detail/42218.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----
	//
	// MIIBczCCARgCAQAwgYoxFDASBgNVBAMMC2FsaXl1bi50ZXN0MQ0wCwYDVQQ
	//
	// ...
	//
	// ...
	//
	// ...
	//
	// vbIgMQIhAKHDWD6/WAMbtezAt4bysJ/BZIDz1jPWuUR5GV4TJ/mS
	//
	// -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
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
	// Specifies whether to immediately issue the certificate. Valid values:
	//
	// 	- 0: asynchronously issues the certificate.
	//
	// 	- 1: immediately issues the certificate.
	//
	// 	- 2: immediately issues the certificate and returns the certificate chain.
	//
	// example:
	//
	// 0
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The identifier of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ed4068c-6f1b-6deb-8e32-3f8439a851cb
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The validity period of the certificate. The value cannot exceed the validity period of the certificate instance. Relative time and absolute time are supported.
	//
	// Units of relative time: year, month, and day.
	//
	// 	- Use y to specify years.
	//
	// 	- Use m to specify months.
	//
	// 	- Use d to specify days.
	//
	// Absolute time: Use Greenwich Mean Time (GMT). Format: `yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\"`
	//
	// 	- Format of the end time: $NotAfter
	//
	// 	- Format of the start time and end time: $NotBefore/$NotAfter
	//
	// This parameter is required.
	//
	// example:
	//
	// Relative time:
	//
	//  ● 1y
	//
	//  ● 3m
	//
	//  ● 7d
	//
	// Absolute time:
	//
	// ● 2006-01-02T15:04:05Z
	//
	// ● 2006-01-02T15:04:05Z/2023-03-09T17:48:13Z
	Validity *string `json:"Validity,omitempty" xml:"Validity,omitempty"`
}

func (s CreateCustomCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequest) SetApiPassthrough(v *CreateCustomCertificateRequestApiPassthrough) *CreateCustomCertificateRequest {
	s.ApiPassthrough = v
	return s
}

func (s *CreateCustomCertificateRequest) SetCsr(v string) *CreateCustomCertificateRequest {
	s.Csr = &v
	return s
}

func (s *CreateCustomCertificateRequest) SetEnableCrl(v int64) *CreateCustomCertificateRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateCustomCertificateRequest) SetImmediately(v int32) *CreateCustomCertificateRequest {
	s.Immediately = &v
	return s
}

func (s *CreateCustomCertificateRequest) SetParentIdentifier(v string) *CreateCustomCertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateCustomCertificateRequest) SetValidity(v string) *CreateCustomCertificateRequest {
	s.Validity = &v
	return s
}

type CreateCustomCertificateRequestApiPassthrough struct {
	// The extensions of the certificate.
	Extensions *CreateCustomCertificateRequestApiPassthroughExtensions `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Struct"`
	// The serial number MUST be a positive integer assigned by the CA to each certificate.
	//
	// example:
	//
	// 16889526086333
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The name of the entity that uses the certificate.
	Subject *CreateCustomCertificateRequestApiPassthroughSubject `json:"Subject,omitempty" xml:"Subject,omitempty" type:"Struct"`
}

func (s CreateCustomCertificateRequestApiPassthrough) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthrough) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthrough) SetExtensions(v *CreateCustomCertificateRequestApiPassthroughExtensions) *CreateCustomCertificateRequestApiPassthrough {
	s.Extensions = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthrough) SetSerialNumber(v string) *CreateCustomCertificateRequestApiPassthrough {
	s.SerialNumber = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthrough) SetSubject(v *CreateCustomCertificateRequestApiPassthroughSubject) *CreateCustomCertificateRequestApiPassthrough {
	s.Subject = v
	return s
}

type CreateCustomCertificateRequestApiPassthroughExtensions struct {
	// If it is a necessary parameter, the critical list contains the parameter name.
	Criticals []*string `json:"Criticals,omitempty" xml:"Criticals,omitempty" type:"Repeated"`
	// The extended key usage.
	ExtendedKeyUsages []*string `json:"ExtendedKeyUsages,omitempty" xml:"ExtendedKeyUsages,omitempty" type:"Repeated"`
	// The key usage.
	KeyUsage *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty" type:"Struct"`
	// The aliases of the entities.
	SubjectAlternativeNames []*CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
}

func (s CreateCustomCertificateRequestApiPassthroughExtensions) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughExtensions) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) SetCriticals(v []*string) *CreateCustomCertificateRequestApiPassthroughExtensions {
	s.Criticals = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) SetExtendedKeyUsages(v []*string) *CreateCustomCertificateRequestApiPassthroughExtensions {
	s.ExtendedKeyUsages = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) SetKeyUsage(v *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) *CreateCustomCertificateRequestApiPassthroughExtensions {
	s.KeyUsage = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) SetSubjectAlternativeNames(v []*CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) *CreateCustomCertificateRequestApiPassthroughExtensions {
	s.SubjectAlternativeNames = v
	return s
}

type CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage struct {
	// The original name of the parameter is NonRepudiation.
	//
	// example:
	//
	// false
	ContentCommitment *bool `json:"ContentCommitment,omitempty" xml:"ContentCommitment,omitempty"`
	// Specifies whether the key can be used for data encryption.
	//
	// example:
	//
	// false
	DataEncipherment *bool `json:"DataEncipherment,omitempty" xml:"DataEncipherment,omitempty"`
	// Specifies whether the key can be used only for data decryption.
	//
	// example:
	//
	// false
	DecipherOnly *bool `json:"DecipherOnly,omitempty" xml:"DecipherOnly,omitempty"`
	// Specifies whether the key can be used for digital signing. If you set this parameter to true, the private key of the certificate can be used to generate digital signatures, and the public key of the certificate can be used to verify digital signatures.
	//
	// example:
	//
	// true
	DigitalSignature *bool `json:"DigitalSignature,omitempty" xml:"DigitalSignature,omitempty"`
	// Specifies whether the key can be used only for data encryption.
	//
	// example:
	//
	// false
	EncipherOnly *bool `json:"EncipherOnly,omitempty" xml:"EncipherOnly,omitempty"`
	// Specifies whether the key can be used for key agreement.
	//
	// example:
	//
	// false
	KeyAgreement *bool `json:"KeyAgreement,omitempty" xml:"KeyAgreement,omitempty"`
	// Specifies whether the key can be used for data encipherment.
	//
	// example:
	//
	// false
	KeyEncipherment *bool `json:"KeyEncipherment,omitempty" xml:"KeyEncipherment,omitempty"`
	// Specifies whether the key can be used for non-repudiation. This parameter is renamed ContentCommitment in the X.509 standard.
	//
	// example:
	//
	// false
	NonRepudiation *bool `json:"NonRepudiation,omitempty" xml:"NonRepudiation,omitempty"`
}

func (s CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetContentCommitment(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.ContentCommitment = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetDataEncipherment(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.DataEncipherment = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetDecipherOnly(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.DecipherOnly = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetDigitalSignature(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.DigitalSignature = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetEncipherOnly(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.EncipherOnly = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetKeyAgreement(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.KeyAgreement = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetKeyEncipherment(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.KeyEncipherment = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetNonRepudiation(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.NonRepudiation = &v
	return s
}

type CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames struct {
	// The type of the alias. Valid values:
	//
	// 	- rfc822Name: email address
	//
	// 	- dNSName: domain name
	//
	// 	- uniformResourceIdentifier: URI
	//
	// 	- iPAddress: IP address
	//
	// This parameter is required.
	//
	// example:
	//
	// dNSName
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The alias that meets the requirement of a specified type.
	//
	// example:
	//
	// rfc822Name:
	//
	// exmaple@certqa.cn
	//
	// dNSName:
	//
	// www.certqa.cn
	//
	// uniformResourceIdentifier:
	//
	// acs:ecs:regionid:15619224785*****:instance/i-bp1bzvz55uz27hf*****
	//
	// iPAddress:
	//
	// 127.0.0.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) SetType(v string) *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames {
	s.Type = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) SetValue(v string) *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames {
	s.Value = &v
	return s
}

type CreateCustomCertificateRequestApiPassthroughSubject struct {
	// The common name of the certificate user.
	//
	// example:
	//
	// Bob
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country. The value is an alpha-2 country code that complies with the ISO 3166-1 standard. For more information about country codes, visit <https://www.iso.org/obp/ui/#search/code/>.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// Customize the Subject attributes of the certificate.
	CustomAttributes []*CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty" type:"Repeated"`
	// The name of the city in which the organization is located. The value can contain letters.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the organization.
	//
	// example:
	//
	// XXX company
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization.
	//
	// example:
	//
	// XXX department
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The name of the province or state in which the organization associated with the certificate is located.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateCustomCertificateRequestApiPassthroughSubject) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughSubject) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetCommonName(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.CommonName = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetCountry(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.Country = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetCustomAttributes(v []*CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.CustomAttributes = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetLocality(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.Locality = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetOrganization(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.Organization = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetOrganizationUnit(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetState(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.State = &v
	return s
}

type CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes struct {
	// Custom attribute type as:
	//
	// - 2.5.4.6 : country
	//
	// - 2.5.4.10 : organization
	//
	// - 2.5.4.11 : organizational unit
	//
	// - 2.5.4.12 : title
	//
	// - 2.5.4.3 : common name
	//
	// - 2.5.4.9 : street
	//
	// - 2.5.4.5 : serial number
	//
	// - 2.5.4.7 : locality
	//
	// - 2.5.4.8 : state
	//
	// - 1.3.6.1.4.1.37244.1.1 : Matter Operational Certificate - Node ID
	//
	// - 1.3.6.1.4.1.37244.1.5 : Matter Operational Certificate - Fabric ID
	//
	// - 1.3.6.1.4.1.37244.2.1 : Matter Device Attestation Certificate Vender ID (VID)
	//
	// - 1.3.6.1.4.1.37244.2.2 : Matter Device Attestation Certificate Product ID (PID).
	//
	// example:
	//
	// 2.5.4.3
	ObjectIdentifier *string `json:"ObjectIdentifier,omitempty" xml:"ObjectIdentifier,omitempty"`
	// Custom attribute value.
	//
	// example:
	//
	// Aliyun
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) SetObjectIdentifier(v string) *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes {
	s.ObjectIdentifier = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) SetValue(v string) *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes {
	s.Value = &v
	return s
}

type CreateCustomCertificateResponseBody struct {
	// The content of the certificate. This parameter is returned only if Immediately is set to 1 or 2.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/
	//
	// ...
	//
	// ...
	//
	// ...
	//
	// KOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==
	//
	// -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain of the certificate. This parameter is returned only if Immediately is set to 2.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIBfzCCATGgAwIBAgIUfI5kSdcO2S0+LkpdL3b2VUJG10YwBQYDK2VwMDUxCzAJ
	//
	// ...
	//
	// ...
	//
	// ...
	//
	// ZYYG
	//
	// -----END CERTIFICATE-----
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIBczCCARgCAQAwgYoxFDASBgNVBAMMC2FsaXl1bi50ZXN0MQ0wCwYDVQQ
	//
	// ...
	//
	// ...
	//
	// ...
	//
	// KL5cUmF
	//
	// -----END CERTIFICATE-----
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate. This parameter is returned only if Immediately is set to 1 or 2.
	//
	// example:
	//
	// 084bde9cd233f0ddae33adc438cfbbbd****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s CreateCustomCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateResponseBody) SetCertificate(v string) *CreateCustomCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *CreateCustomCertificateResponseBody) SetCertificateChain(v string) *CreateCustomCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateCustomCertificateResponseBody) SetIdentifier(v string) *CreateCustomCertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateCustomCertificateResponseBody) SetRequestId(v string) *CreateCustomCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomCertificateResponseBody) SetSerialNumber(v string) *CreateCustomCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

type CreateCustomCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateResponse) SetHeaders(v map[string]*string) *CreateCustomCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomCertificateResponse) SetStatusCode(v int32) *CreateCustomCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomCertificateResponse) SetBody(v *CreateCustomCertificateResponseBody) *CreateCustomCertificateResponse {
	s.Body = v
	return s
}

type CreateRevokeClientCertificateRequest struct {
	// The unique identifier of the client certificate or server certificate that you want to revoke.
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers of all client certificates and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s CreateRevokeClientCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRevokeClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateRevokeClientCertificateRequest) SetIdentifier(v string) *CreateRevokeClientCertificateRequest {
	s.Identifier = &v
	return s
}

type CreateRevokeClientCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRevokeClientCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRevokeClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRevokeClientCertificateResponseBody) SetRequestId(v string) *CreateRevokeClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

type CreateRevokeClientCertificateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRevokeClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRevokeClientCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRevokeClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateRevokeClientCertificateResponse) SetHeaders(v map[string]*string) *CreateRevokeClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateRevokeClientCertificateResponse) SetStatusCode(v int32) *CreateRevokeClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRevokeClientCertificateResponse) SetBody(v *CreateRevokeClientCertificateResponseBody) *CreateRevokeClientCertificateResponse {
	s.Body = v
	return s
}

type CreateRootCACertificateRequest struct {
	// The key algorithm of the root CA certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
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
	// The encryption algorithm of the root CA certificate must be consistent with the **encryption algorithm*	- of the private root CA instance that you purchase. For example, if the **encryption algorithm*	- of the private root CA instance that you purchase is **RSA**, the key algorithm of the root CA certificate must be **RSA_1024**, **RSA_2048**, or **RSA_4096**.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The common name or abbreviation of the organization. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country or region in which the organization is located. You can enter an alpha-2 code. For example, you can use **CN*	- to indicate China and use **US*	- to indicate the United States.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The name of the city in which the organization is located. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the organization that is associated with the root CA certificate. You can enter the name of your enterprise or company. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the root CA certificate. Unit: years.
	//
	// >  We recommend that you set this parameter to a value from 5 to 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateRootCACertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRootCACertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateRootCACertificateRequest) SetAlgorithm(v string) *CreateRootCACertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetCommonName(v string) *CreateRootCACertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetCountryCode(v string) *CreateRootCACertificateRequest {
	s.CountryCode = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetLocality(v string) *CreateRootCACertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetOrganization(v string) *CreateRootCACertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetOrganizationUnit(v string) *CreateRootCACertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetState(v string) *CreateRootCACertificateRequest {
	s.State = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetYears(v int32) *CreateRootCACertificateRequest {
	s.Years = &v
	return s
}

type CreateRootCACertificateResponseBody struct {
	// The root CA certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain of the root CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the root CA certificate.
	//
	// example:
	//
	// 1a83bcbb89e562885e40aa0108f5****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6D9B4C5F-7140-5B41-924C-329181DC00C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRootCACertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRootCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRootCACertificateResponseBody) SetCertificate(v string) *CreateRootCACertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *CreateRootCACertificateResponseBody) SetCertificateChain(v string) *CreateRootCACertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateRootCACertificateResponseBody) SetIdentifier(v string) *CreateRootCACertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateRootCACertificateResponseBody) SetRequestId(v string) *CreateRootCACertificateResponseBody {
	s.RequestId = &v
	return s
}

type CreateRootCACertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRootCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRootCACertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRootCACertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateRootCACertificateResponse) SetHeaders(v map[string]*string) *CreateRootCACertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateRootCACertificateResponse) SetStatusCode(v int32) *CreateRootCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRootCACertificateResponse) SetBody(v *CreateRootCACertificateResponseBody) *CreateRootCACertificateResponse {
	s.Body = v
	return s
}

type CreateServerCertificateRequest struct {
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
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located, such as CN or US.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The validity period of the server certificate. Unit: days. You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime*	- parameters. The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// 	- If you specify the **Days*	- parameter, you can specify both the **BeforeTime*	- and **AfterTime*	- parameters or leave them both empty.
	//
	// 	- If you do not specify the **Days*	- parameter, you must specify both the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	// >
	//
	// 	- If you specify the **Days**, **BeforeTime**, and **AfterTime*	- parameters together, the validity period of the server certificate is determined by the value of the **Days*	- parameter.
	//
	// 	- The validity period of the server certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the validity period of an intermediate CA certificate.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The additional domain names and additional IP addresses of the server certificate. After you add additional domain names and additional IP addresses to a certificate, you can apply the certificate to the domain names and IP addresses.
	//
	// Separate multiple domain names and multiple IP addresses with commas (,).
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
	// Alibaba Cloud
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
	// 271ae6bb538d538c70c01f81dg3****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the server certificate. Unit: years.
	//
	// example:
	//
	// 1
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateServerCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServerCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateRequest) SetAfterTime(v int64) *CreateServerCertificateRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateServerCertificateRequest) SetAlgorithm(v string) *CreateServerCertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateServerCertificateRequest) SetBeforeTime(v int64) *CreateServerCertificateRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateServerCertificateRequest) SetCommonName(v string) *CreateServerCertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateServerCertificateRequest) SetCountry(v string) *CreateServerCertificateRequest {
	s.Country = &v
	return s
}

func (s *CreateServerCertificateRequest) SetDays(v int32) *CreateServerCertificateRequest {
	s.Days = &v
	return s
}

func (s *CreateServerCertificateRequest) SetDomain(v string) *CreateServerCertificateRequest {
	s.Domain = &v
	return s
}

func (s *CreateServerCertificateRequest) SetEnableCrl(v int64) *CreateServerCertificateRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateServerCertificateRequest) SetImmediately(v int32) *CreateServerCertificateRequest {
	s.Immediately = &v
	return s
}

func (s *CreateServerCertificateRequest) SetLocality(v string) *CreateServerCertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateServerCertificateRequest) SetMonths(v int32) *CreateServerCertificateRequest {
	s.Months = &v
	return s
}

func (s *CreateServerCertificateRequest) SetOrganization(v string) *CreateServerCertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateServerCertificateRequest) SetOrganizationUnit(v string) *CreateServerCertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateServerCertificateRequest) SetParentIdentifier(v string) *CreateServerCertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateServerCertificateRequest) SetState(v string) *CreateServerCertificateRequest {
	s.State = &v
	return s
}

func (s *CreateServerCertificateRequest) SetYears(v int32) *CreateServerCertificateRequest {
	s.Years = &v
	return s
}

type CreateServerCertificateResponseBody struct {
	// The certificate chain of the server certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the server certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the server certificate.
	//
	// example:
	//
	// 0f29522da2dae7a1c4b6ab7132ad3c06
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the server certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s CreateServerCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServerCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateResponseBody) SetCertificateChain(v string) *CreateServerCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetIdentifier(v string) *CreateServerCertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetRequestId(v string) *CreateServerCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetSerialNumber(v string) *CreateServerCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetX509Certificate(v string) *CreateServerCertificateResponseBody {
	s.X509Certificate = &v
	return s
}

type CreateServerCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServerCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServerCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServerCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateResponse) SetHeaders(v map[string]*string) *CreateServerCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateServerCertificateResponse) SetStatusCode(v int32) *CreateServerCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServerCertificateResponse) SetBody(v *CreateServerCertificateResponseBody) *CreateServerCertificateResponse {
	s.Body = v
	return s
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
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
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
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the server certificate. Unit: years.
	//
	// example:
	//
	// 1
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateServerCertificateWithCsrRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServerCertificateWithCsrRequest) GoString() string {
	return s.String()
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

func (s *CreateServerCertificateWithCsrRequest) SetState(v string) *CreateServerCertificateWithCsrRequest {
	s.State = &v
	return s
}

func (s *CreateServerCertificateWithCsrRequest) SetYears(v int32) *CreateServerCertificateWithCsrRequest {
	s.Years = &v
	return s
}

type CreateServerCertificateWithCsrResponseBody struct {
	// The certificate chain of the server certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the server certificate.
	//
	// example:
	//
	// 180ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 55C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the server certificate.
	//
	// example:
	//
	// 084bde9cd233f0ddae33adc438cfbbbd****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the server certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s CreateServerCertificateWithCsrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServerCertificateWithCsrResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateWithCsrResponseBody) SetCertificateChain(v string) *CreateServerCertificateWithCsrResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetIdentifier(v string) *CreateServerCertificateWithCsrResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetRequestId(v string) *CreateServerCertificateWithCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetSerialNumber(v string) *CreateServerCertificateWithCsrResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetX509Certificate(v string) *CreateServerCertificateWithCsrResponseBody {
	s.X509Certificate = &v
	return s
}

type CreateServerCertificateWithCsrResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServerCertificateWithCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServerCertificateWithCsrResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServerCertificateWithCsrResponse) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateWithCsrResponse) SetHeaders(v map[string]*string) *CreateServerCertificateWithCsrResponse {
	s.Headers = v
	return s
}

func (s *CreateServerCertificateWithCsrResponse) SetStatusCode(v int32) *CreateServerCertificateWithCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponse) SetBody(v *CreateServerCertificateWithCsrResponseBody) *CreateServerCertificateWithCsrResponse {
	s.Body = v
	return s
}

type CreateSubCACertificateRequest struct {
	// The type of the key algorithm of the intermediate CA. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// 	- **RSA_1024**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **RSA_2048**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **RSA_4096**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **ECC_256**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **SM2_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of an intermediate CA certificate must be consistent with the encryption algorithm of a root CA certificate. The length of the keys can be different. For example, if the key algorithm of the root CA certificate is **RSA_2048**, the key algorithm of the intermediate CA certificate must be **RSA_1024**, **RSA_2048**, or **RSA_4096**.
	//
	// > You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) operation to query the key algorithm of a root CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The common name or abbreviation of the organization. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country or region in which the organization is located. You can enter an alpha-2 or alpha-3 code. For example, you can use **CN*	- to indicate China and use **US*	- to indicate the United States.
	//
	// For more information about country codes, see the **"Country codes"*	- section in [Manage company profiles](https://help.aliyun.com/document_detail/198289.html).
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// CRL validity period: 1-365 days
	//
	// example:
	//
	// 30
	CrlDay *int32 `json:"CrlDay,omitempty" xml:"CrlDay,omitempty"`
	// Enable Crl Service.
	//
	// - 0- No
	//
	// - 1- Yes
	//
	// example:
	//
	// 1
	EnableCrl *bool `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// The extended key usages of the certificate.
	ExtendedKeyUsages []*string `json:"ExtendedKeyUsages,omitempty" xml:"ExtendedKeyUsages,omitempty" type:"Repeated"`
	// The name of the city in which the organization is located. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the organization that is associated with the intermediate CA certificate. You can enter the name of your enterprise or company. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Maizi Technology
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the root CA certificate.
	//
	// > You can call the [DescribeCACertificateList] operation to query the unique identifiers of all CA certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1a83bcbb89e562885e40aa0108f5****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The path length constraint of the certificate. Default value: 0.
	//
	// example:
	//
	// 0
	PathLenConstraint *int32 `json:"PathLenConstraint,omitempty" xml:"PathLenConstraint,omitempty"`
	// The name of the province or state in which the organization is located. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the intermediate CA certificate. Unit: years.
	//
	// We recommend that you set this parameter to 5 to 10.
	//
	// > The validity period of the intermediate CA certificate cannot exceed the validity period of the root CA certificate. You can call the [DescribeCACertificate]operation to query the validity period of a root CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateSubCACertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubCACertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateSubCACertificateRequest) SetAlgorithm(v string) *CreateSubCACertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetCommonName(v string) *CreateSubCACertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetCountryCode(v string) *CreateSubCACertificateRequest {
	s.CountryCode = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetCrlDay(v int32) *CreateSubCACertificateRequest {
	s.CrlDay = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetEnableCrl(v bool) *CreateSubCACertificateRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetExtendedKeyUsages(v []*string) *CreateSubCACertificateRequest {
	s.ExtendedKeyUsages = v
	return s
}

func (s *CreateSubCACertificateRequest) SetLocality(v string) *CreateSubCACertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetOrganization(v string) *CreateSubCACertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetOrganizationUnit(v string) *CreateSubCACertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetParentIdentifier(v string) *CreateSubCACertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetPathLenConstraint(v int32) *CreateSubCACertificateRequest {
	s.PathLenConstraint = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetState(v string) *CreateSubCACertificateRequest {
	s.State = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetYears(v int32) *CreateSubCACertificateRequest {
	s.Years = &v
	return s
}

type CreateSubCACertificateResponseBody struct {
	// The CA certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain of the CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the sub CA certificate created in this request.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of this call request is a unique identifier generated by Alibaba Cloud for the request, which can be used for troubleshooting and locating issues.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSubCACertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubCACertificateResponseBody) SetCertificate(v string) *CreateSubCACertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *CreateSubCACertificateResponseBody) SetCertificateChain(v string) *CreateSubCACertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateSubCACertificateResponseBody) SetIdentifier(v string) *CreateSubCACertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateSubCACertificateResponseBody) SetRequestId(v string) *CreateSubCACertificateResponseBody {
	s.RequestId = &v
	return s
}

type CreateSubCACertificateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubCACertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubCACertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateSubCACertificateResponse) SetHeaders(v map[string]*string) *CreateSubCACertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateSubCACertificateResponse) SetStatusCode(v int32) *CreateSubCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubCACertificateResponse) SetBody(v *CreateSubCACertificateResponseBody) *CreateSubCACertificateResponse {
	s.Body = v
	return s
}

type DeleteClientCertificateRequest struct {
	// The unique identifier of the client certificate or server certificate that you want to delete. The status of the certificate must be **REVOKE**.
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers and status of all client certificates and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DeleteClientCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientCertificateRequest) SetIdentifier(v string) *DeleteClientCertificateRequest {
	s.Identifier = &v
	return s
}

type DeleteClientCertificateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClientCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientCertificateResponseBody) SetRequestId(v string) *DeleteClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClientCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientCertificateResponse) SetHeaders(v map[string]*string) *DeleteClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientCertificateResponse) SetStatusCode(v int32) *DeleteClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientCertificateResponse) SetBody(v *DeleteClientCertificateResponseBody) *DeleteClientCertificateResponse {
	s.Body = v
	return s
}

type DescribeCACertificateRequest struct {
	// The unique identifier of the CA certificate that you want to query.
	//
	// >  You can call the [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) operation to query the unique identifiers of all CA certificates.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DescribeCACertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateRequest) SetIdentifier(v string) *DescribeCACertificateRequest {
	s.Identifier = &v
	return s
}

type DescribeCACertificateResponseBody struct {
	// The details about the CA certificate.
	Certificate *DescribeCACertificateResponseBodyCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The validity period of the CA certificate. Unit: years.
	//
	// example:
	//
	// 10
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s DescribeCACertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateResponseBody) SetCertificate(v *DescribeCACertificateResponseBodyCertificate) *DescribeCACertificateResponseBody {
	s.Certificate = v
	return s
}

func (s *DescribeCACertificateResponseBody) SetRequestId(v string) *DescribeCACertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCACertificateResponseBody) SetYears(v int32) *DescribeCACertificateResponseBody {
	s.Years = &v
	return s
}

type DescribeCACertificateResponseBodyCertificate struct {
	// The expiration date of the CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The encryption algorithm of the CA certificate. Valid values:
	//
	// 	- **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	//
	// 	- **ECC**: the elliptic curve cryptography (ECC) algorithm.
	//
	// 	- **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance date of the CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1634283958000
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// CA certificate chain.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// 用户证书
	//
	// -----END CERTIFICATE-----
	//
	// -----BEGIN CERTIFICATE-----
	//
	// 中间证书
	//
	// -----END CERTIFICATE-----
	//
	// -----BEGIN CERTIFICATE-----
	//
	// 根证书
	//
	// -----END CERTIFICATE-----
	CaCertChain *string `json:"CaCertChain,omitempty" xml:"CaCertChain,omitempty"`
	// The number of certificates issued by private CA instances.
	//
	// example:
	//
	// 10
	CertIssuedCount *int64 `json:"CertIssuedCount,omitempty" xml:"CertIssuedCount,omitempty"`
	// The remaining number of assignable certificate quotas.
	//
	// example:
	//
	// 30
	CertRemainingCount *int64 `json:"CertRemainingCount,omitempty" xml:"CertRemainingCount,omitempty"`
	// The total number of purchased certificate quotas.
	//
	// example:
	//
	// 40
	CertTotalCount *int64 `json:"CertTotalCount,omitempty" xml:"CertTotalCount,omitempty"`
	// The type of the CA certificate. Valid values:
	//
	// 	- **ROOT**: root CA certificate
	//
	// 	- **SUB_ROOT**: intermediate CA certificate
	//
	// example:
	//
	// SUB_ROOT
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name or abbreviation of the organization that is associated with the CA certificate.
	//
	// example:
	//
	// Aliyun
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// CRL validity period: 1-365 days.
	//
	// example:
	//
	// 90
	CrlDay *int32 `json:"CrlDay,omitempty" xml:"CrlDay,omitempty"`
	// The status of the certificate revocation list (CRL) feature.
	//
	// example:
	//
	// ACTIVE
	CrlStatus *string `json:"CrlStatus,omitempty" xml:"CrlStatus,omitempty"`
	// The address of the CRL.
	//
	// example:
	//
	// https://crl-cn-publish.oss-cn-hangzhou.aliyuncs.com/pca/crl/1925647866611395/1ed40789-483f-6023-b6b8-29ddd3bb0a9a.crl
	CrlUrl *string `json:"CrlUrl,omitempty" xml:"CrlUrl,omitempty"`
	// The unique identifier of the CA certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the CA certificate.
	//
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the CA certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization that is associated with the CA certificate.
	//
	// example:
	//
	// Alibaba Cloud Computing Co., Ltd.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization that is associated with the CA certificate.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the root CA certificate from which the CA certificate is issued.
	//
	// >  This parameter is returned only if the value of the **CertificateType*	- parameter is **SUB_ROOT**. The value SUB_ROOT indicates an intermediate CA certificate.
	//
	// example:
	//
	// 1a83bcbb89e562885e40aa0108f5****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the CA certificate.
	//
	// example:
	//
	// 70e3b2566d92805173767869727fb92e****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the CA certificate.
	//
	// example:
	//
	// 14dcc8afc7578e1fcec36d658f7e20de18f6957bbac42b373a66bc9de4e9****
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the CA certificate.
	//
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the CA certificate. Valid values:
	//
	// 	- **ISSUE**: The CA certificate is issued.
	//
	// 	- **REVOKE**: The CA certificate is revoked.
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user attribute of the CA certificate, which contains the following information:
	//
	// 	- **C**: the country code in which the organization is located
	//
	// 	- **O**: the name of the organization
	//
	// 	- **OU**: the name of the department or branch in the organization
	//
	// 	- **L**: the name of the city in which the organization is located
	//
	// 	- **ST**: the name of the province, municipality, or autonomous region in which the organization is located
	//
	// 	- **CN**: the common name or abbreviation of the organization
	//
	// example:
	//
	// C=CN,O=Alibaba Cloud Computing Co., Ltd.,OU=Security,L=Hangzhou,ST=Zhejiang,CN=Aliyun
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// The content of the CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- …… -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s DescribeCACertificateResponseBodyCertificate) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificateResponseBodyCertificate) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateResponseBodyCertificate) SetAfterDate(v int64) *DescribeCACertificateResponseBodyCertificate {
	s.AfterDate = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetAlgorithm(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Algorithm = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetBeforeDate(v int64) *DescribeCACertificateResponseBodyCertificate {
	s.BeforeDate = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCaCertChain(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CaCertChain = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCertIssuedCount(v int64) *DescribeCACertificateResponseBodyCertificate {
	s.CertIssuedCount = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCertRemainingCount(v int64) *DescribeCACertificateResponseBodyCertificate {
	s.CertRemainingCount = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCertTotalCount(v int64) *DescribeCACertificateResponseBodyCertificate {
	s.CertTotalCount = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCertificateType(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CertificateType = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCommonName(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CommonName = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCountryCode(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CountryCode = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCrlDay(v int32) *DescribeCACertificateResponseBodyCertificate {
	s.CrlDay = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCrlStatus(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CrlStatus = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCrlUrl(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CrlUrl = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetIdentifier(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Identifier = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetKeySize(v int32) *DescribeCACertificateResponseBodyCertificate {
	s.KeySize = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetLocality(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Locality = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetMd5(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Md5 = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetOrganization(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Organization = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetOrganizationUnit(v string) *DescribeCACertificateResponseBodyCertificate {
	s.OrganizationUnit = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetParentIdentifier(v string) *DescribeCACertificateResponseBodyCertificate {
	s.ParentIdentifier = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetSans(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Sans = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetSerialNumber(v string) *DescribeCACertificateResponseBodyCertificate {
	s.SerialNumber = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetSha2(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Sha2 = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetSignAlgorithm(v string) *DescribeCACertificateResponseBodyCertificate {
	s.SignAlgorithm = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetState(v string) *DescribeCACertificateResponseBodyCertificate {
	s.State = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetStatus(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Status = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetSubjectDN(v string) *DescribeCACertificateResponseBodyCertificate {
	s.SubjectDN = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetX509Certificate(v string) *DescribeCACertificateResponseBodyCertificate {
	s.X509Certificate = &v
	return s
}

type DescribeCACertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCACertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateResponse) SetHeaders(v map[string]*string) *DescribeCACertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCACertificateResponse) SetStatusCode(v int32) *DescribeCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCACertificateResponse) SetBody(v *DescribeCACertificateResponseBody) *DescribeCACertificateResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s DescribeCACertificateCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateCountResponseBody) SetRequestId(v string) *DescribeCACertificateCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCACertificateCountResponseBody) SetTotalCount(v int32) *DescribeCACertificateCountResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCACertificateCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCACertificateCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCACertificateCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificateCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateCountResponse) SetHeaders(v map[string]*string) *DescribeCACertificateCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeCACertificateCountResponse) SetStatusCode(v int32) *DescribeCACertificateCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCACertificateCountResponse) SetBody(v *DescribeCACertificateCountResponseBody) *DescribeCACertificateCountResponse {
	s.Body = v
	return s
}

type DescribeCACertificateListRequest struct {
	// Ca status.
	//
	// - issue: inUse.
	//
	// - forbidden: forbidden.
	//
	// - revoke: revoked.
	//
	// example:
	//
	// issue
	CaStatus *string `json:"CaStatus,omitempty" xml:"CaStatus,omitempty"`
	// The type of the certificate. Valid values:
	//
	// - root: rootCA.
	//
	// - subRoot: subCA.
	//
	// - externalCa: import.
	//
	// example:
	//
	// subRoot
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The unique identifier of the CA certificate.
	//
	// >  You can call the [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) operation to query the unique identifiers of all CA certificates.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The CA Issuer Type.
	//
	// - local: Private certificate.
	//
	// - iTrusChina: Compliance CA.
	//
	// - external: External Import.
	//
	// example:
	//
	// local
	IssuerType *string `json:"IssuerType,omitempty" xml:"IssuerType,omitempty"`
	// The number of CA certificates per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// valid time.
	//
	// - valid: means in the valid period.
	//
	// - notValid: means expired.
	//
	// example:
	//
	// valid
	ValidStatus *string `json:"ValidStatus,omitempty" xml:"ValidStatus,omitempty"`
}

func (s DescribeCACertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateListRequest) SetCaStatus(v string) *DescribeCACertificateListRequest {
	s.CaStatus = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetCertType(v string) *DescribeCACertificateListRequest {
	s.CertType = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetCurrentPage(v int32) *DescribeCACertificateListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetIdentifier(v string) *DescribeCACertificateListRequest {
	s.Identifier = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetIssuerType(v string) *DescribeCACertificateListRequest {
	s.IssuerType = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetShowSize(v int32) *DescribeCACertificateListRequest {
	s.ShowSize = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetValidStatus(v string) *DescribeCACertificateListRequest {
	s.ValidStatus = &v
	return s
}

type DescribeCACertificateListResponseBody struct {
	// The details about the CA certificates.
	CertificateList []*DescribeCACertificateListResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of returned pages.
	//
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of CA certificates returned per page.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of root CA certificates and intermediate CA certificates that are returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCACertificateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateListResponseBody) SetCertificateList(v []*DescribeCACertificateListResponseBodyCertificateList) *DescribeCACertificateListResponseBody {
	s.CertificateList = v
	return s
}

func (s *DescribeCACertificateListResponseBody) SetCurrentPage(v int32) *DescribeCACertificateListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCACertificateListResponseBody) SetPageCount(v int32) *DescribeCACertificateListResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeCACertificateListResponseBody) SetRequestId(v string) *DescribeCACertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCACertificateListResponseBody) SetShowSize(v int32) *DescribeCACertificateListResponseBody {
	s.ShowSize = &v
	return s
}

func (s *DescribeCACertificateListResponseBody) SetTotalCount(v int32) *DescribeCACertificateListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCACertificateListResponseBodyCertificateList struct {
	// The expiration date of the CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The encryption algorithm of the CA certificate. Valid values:
	//
	// 	- **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	//
	// 	- **ECC**: the elliptic curve cryptography (ECC) algorithm.
	//
	// 	- **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The alias of the CA.
	//
	// example:
	//
	// Aliyun_CA
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The issuance date of the CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1634283958000
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the CA certificate. Valid values:
	//
	// 	- **ROOT**: a root CA certificate.
	//
	// 	- **SUB_ROOT**: an intermediate CA certificate.
	//
	// example:
	//
	// SUB_ROOT
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name or abbreviation of the organization that is associated with the CA certificate.
	//
	// example:
	//
	// Aliyun
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Gift        *int32  `json:"Gift,omitempty" xml:"Gift,omitempty"`
	// The unique identifier of the CA certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the CA certificate.
	//
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the CA certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization that is associated with the CA certificate.
	//
	// example:
	//
	// Alibaba Cloud Computing Co., Ltd.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization that is associated with the CA certificate.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the root CA certificate from which the CA certificate is issued.
	//
	// >  This parameter is returned only if the value of the **CertificateType*	- parameter is **SUB_ROOT**. The value SUB_ROOT indicates an intermediate CA certificate.
	//
	// example:
	//
	// 1a83bcbb89e562885e40aa0108f5****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the CA certificate.
	//
	// example:
	//
	// 70e3b2566d92805173767869727fb92e****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the CA certificate.
	//
	// example:
	//
	// 14dcc8afc7578e1fcec36d658f7e20de18f6957bbac42b373a66bc9de4e9****
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the CA certificate.
	//
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the CA certificate. Valid values:
	//
	// 	- **ISSUE**: The CA certificate is issued.
	//
	// 	- **REVOKE**: The CA certificate is revoked.
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Distinguished Name (DN) attribute of the CA certificate, which indicates the user information of the certificate. The DN attribute contains the following information:
	//
	// 	- **C**: the code of the country in which the organization is located.
	//
	// 	- **O**: the name of the organization.
	//
	// 	- **OU**: the name of the department or branch in the organization.
	//
	// 	- **L**: the name of the city in which the organization is located.
	//
	// 	- **CN**: the common name or abbreviation of the organization.
	//
	// example:
	//
	// C=CN,O=Alibaba Cloud Computing Co., Ltd.,OU=Security,L=Hangzhou,ST=Zhejiang,CN=Aliyun
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	Trial     *int32  `json:"Trial,omitempty" xml:"Trial,omitempty"`
	// The content of the CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- …… -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
	// The validity period of the CA certificate. Unit: years.
	//
	// example:
	//
	// 3
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s DescribeCACertificateListResponseBodyCertificateList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificateListResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetAfterDate(v int64) *DescribeCACertificateListResponseBodyCertificateList {
	s.AfterDate = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetAlgorithm(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetAlias(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Alias = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetBeforeDate(v int64) *DescribeCACertificateListResponseBodyCertificateList {
	s.BeforeDate = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetCertificateType(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.CertificateType = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetCommonName(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetCountryCode(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.CountryCode = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetGift(v int32) *DescribeCACertificateListResponseBodyCertificateList {
	s.Gift = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetIdentifier(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Identifier = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetKeySize(v int32) *DescribeCACertificateListResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetLocality(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Locality = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetMd5(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Md5 = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetOrganization(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Organization = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetOrganizationUnit(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.OrganizationUnit = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetParentIdentifier(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.ParentIdentifier = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetSans(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Sans = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetSerialNumber(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.SerialNumber = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetSha2(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Sha2 = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetSignAlgorithm(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.SignAlgorithm = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetState(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.State = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetStatus(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Status = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetSubjectDN(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.SubjectDN = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetTrial(v int32) *DescribeCACertificateListResponseBodyCertificateList {
	s.Trial = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetX509Certificate(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.X509Certificate = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetYears(v int32) *DescribeCACertificateListResponseBodyCertificateList {
	s.Years = &v
	return s
}

type DescribeCACertificateListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCACertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCACertificateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateListResponse) SetHeaders(v map[string]*string) *DescribeCACertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCACertificateListResponse) SetStatusCode(v int32) *DescribeCACertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCACertificateListResponse) SetBody(v *DescribeCACertificateListResponseBody) *DescribeCACertificateListResponse {
	s.Body = v
	return s
}

type DescribeCertificatePrivateKeyRequest struct {
	// The password that is used to encrypt the private key. The password can contain letters, digits, and special characters, such as `, + - _ #`. The password can be up to 32 bytes in length.
	//
	// **Warning*	- You must remember the password that you specify. The password is required to decrypt the encrypted private key. If you forget the password, the encrypted private key that is returned cannot be decrypted. You must call this operation again.
	//
	// This parameter is required.
	//
	// example:
	//
	// !QA@WS3ed
	EncryptedCode *string `json:"EncryptedCode,omitempty" xml:"EncryptedCode,omitempty"`
	// The unique identifier of the client certificate or server certificate that you want to query.
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers of all client certificates and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// bc37133bb7ed68c7938d928fd26d****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DescribeCertificatePrivateKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificatePrivateKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificatePrivateKeyRequest) SetEncryptedCode(v string) *DescribeCertificatePrivateKeyRequest {
	s.EncryptedCode = &v
	return s
}

func (s *DescribeCertificatePrivateKeyRequest) SetIdentifier(v string) *DescribeCertificatePrivateKeyRequest {
	s.Identifier = &v
	return s
}

type DescribeCertificatePrivateKeyResponseBody struct {
	// The content of the encrypted private key.
	//
	// example:
	//
	// -----BEGIN ENCRYPTED PRIVATE KEY----- …… -----END ENCRYPTED PRIVATE KEY-----
	EncryptedData *string `json:"EncryptedData,omitempty" xml:"EncryptedData,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 09470F19-CEE8-5C63-BF2C-02B5E3F07A17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCertificatePrivateKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificatePrivateKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificatePrivateKeyResponseBody) SetEncryptedData(v string) *DescribeCertificatePrivateKeyResponseBody {
	s.EncryptedData = &v
	return s
}

func (s *DescribeCertificatePrivateKeyResponseBody) SetRequestId(v string) *DescribeCertificatePrivateKeyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCertificatePrivateKeyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertificatePrivateKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertificatePrivateKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificatePrivateKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificatePrivateKeyResponse) SetHeaders(v map[string]*string) *DescribeCertificatePrivateKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificatePrivateKeyResponse) SetStatusCode(v int32) *DescribeCertificatePrivateKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertificatePrivateKeyResponse) SetBody(v *DescribeCertificatePrivateKeyResponseBody) *DescribeCertificatePrivateKeyResponse {
	s.Body = v
	return s
}

type DescribeClientCertificateRequest struct {
	// The unique identifier of the client certificate or the server certificate that you want to query.
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers of all client certificates and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// d3b95700998e47afc4d95f886579****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DescribeClientCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateRequest) SetIdentifier(v string) *DescribeClientCertificateRequest {
	s.Identifier = &v
	return s
}

type DescribeClientCertificateResponseBody struct {
	// The details about the client certificate or the server certificate.
	Certificate *DescribeClientCertificateResponseBodyCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateResponseBody) SetCertificate(v *DescribeClientCertificateResponseBodyCertificate) *DescribeClientCertificateResponseBody {
	s.Certificate = v
	return s
}

func (s *DescribeClientCertificateResponseBody) SetRequestId(v string) *DescribeClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClientCertificateResponseBodyCertificate struct {
	// The expiration date of the certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The type of the encryption algorithm of the certificate. Valid values:
	//
	// 	- **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	//
	// 	- **ECC**: the elliptic curve cryptography (ECC) algorithm.
	//
	// 	- **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance date of the certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1634283958000
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **CLIENT**: client certificate
	//
	// 	- **SERVER**: server certificate
	//
	// example:
	//
	// SERVER
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// aliyun.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The validity period of the certificate. Unit: days.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The unique identifier of the certificate.
	//
	// example:
	//
	// d3b95700998e47afc4d95f886579****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the certificate.
	//
	// example:
	//
	// 4096
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the certificate.
	//
	// example:
	//
	// d3b95700998e47afc4d95f886579****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Alibaba Cloud Computing Co., Ltd.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department in the organization. The organization is associated with the intermediate certificate authority (CA) certificate from which the certificate is issued.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate certificate from which the client certificate is issued.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The subject alternative name (SAN) extension of the certificate. The value indicates additional information, including the additional domain names or IP addresses that are associated with the certificate.
	//
	// The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that corresponds to a SAN extension. A SAN extension struct contains the following parameters:
	//
	// 	- **Type**: the type of the extension. Data type: integer. Valid values:
	//
	//     	- **1**: an email address
	//
	//     	- **2**: a domain name
	//
	//     	- **6**: a Uniform Resource Identifier (URI)
	//
	//     	- **7**: an IP address
	//
	// 	- **Value**: the value of the extension. Data type: string.
	//
	// example:
	//
	// [ {"Type": 7, "Value": "192.0.XX.XX"}, {"Type": 2, "Value": "www.aliyundoc.com"}, ]
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 62b2b943a32d96883a6650e672ea0276****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	//
	// example:
	//
	// 14dcc8afc7578e1fcec36d658f7e20de18f6957bbac42b373a66bc9de4e9****
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the certificate.
	//
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
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
	// The distinguished name (DN) extension of the certificate, which indicates the user of the certificate. The DN extension includes the following information:
	//
	// 	- **C**: the country
	//
	// 	- **O**: the organization
	//
	// 	- **OU**: the department
	//
	// 	- **L**: the city
	//
	// 	- **ST**: the province, municipality, or autonomous region
	//
	// 	- **CN**: the common name
	//
	// example:
	//
	// C=CN,O=Alibaba Cloud Computing Co., Ltd.,OU=Security,L=Hangzhou,ST=Zhejiang,CN=Aliyun
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// The content of the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  ...... -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s DescribeClientCertificateResponseBodyCertificate) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateResponseBodyCertificate) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetAfterDate(v int64) *DescribeClientCertificateResponseBodyCertificate {
	s.AfterDate = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetAlgorithm(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Algorithm = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetBeforeDate(v int64) *DescribeClientCertificateResponseBodyCertificate {
	s.BeforeDate = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetCertificateType(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.CertificateType = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetCommonName(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.CommonName = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetCountryCode(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.CountryCode = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetDays(v int32) *DescribeClientCertificateResponseBodyCertificate {
	s.Days = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetIdentifier(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Identifier = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetKeySize(v int32) *DescribeClientCertificateResponseBodyCertificate {
	s.KeySize = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetLocality(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Locality = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetMd5(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Md5 = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetOrganization(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Organization = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetOrganizationUnit(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.OrganizationUnit = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetParentIdentifier(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.ParentIdentifier = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetSans(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Sans = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetSerialNumber(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetSha2(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Sha2 = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetSignAlgorithm(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.SignAlgorithm = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetState(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.State = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetStatus(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Status = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetSubjectDN(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.SubjectDN = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetX509Certificate(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.X509Certificate = &v
	return s
}

type DescribeClientCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateResponse) SetHeaders(v map[string]*string) *DescribeClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientCertificateResponse) SetStatusCode(v int32) *DescribeClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientCertificateResponse) SetBody(v *DescribeClientCertificateResponseBody) *DescribeClientCertificateResponse {
	s.Body = v
	return s
}

type DescribeClientCertificateStatusRequest struct {
	// The unique identifiers of the client certificates or server certificates that you want to query. Separate multiple unique identifiers with commas (,).
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers of all client certificates and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DescribeClientCertificateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusRequest) SetIdentifier(v string) *DescribeClientCertificateStatusRequest {
	s.Identifier = &v
	return s
}

type DescribeClientCertificateStatusResponseBody struct {
	// An array that consists of the status information about the certificates.
	CertificateStatus []*DescribeClientCertificateStatusResponseBodyCertificateStatus `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientCertificateStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusResponseBody) SetCertificateStatus(v []*DescribeClientCertificateStatusResponseBodyCertificateStatus) *DescribeClientCertificateStatusResponseBody {
	s.CertificateStatus = v
	return s
}

func (s *DescribeClientCertificateStatusResponseBody) SetRequestId(v string) *DescribeClientCertificateStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClientCertificateStatusResponseBodyCertificateStatus struct {
	// The date on which the certificate was revoked.
	//
	// >  This parameter is returned only when the value of the **Status*	- parameter is **revoked**. The value revoked indicates that the certificate is revoked.
	//
	// example:
	//
	// 2021-01-01T00:00Z
	RevokeTime *int64 `json:"RevokeTime,omitempty" xml:"RevokeTime,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// b67e53ebcea9b77d65b0c3236646d715****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- **good**: The certificate is not revoked.
	//
	// 	- **revoked**: The certificate is revoked.
	//
	// 	- **unknown**: The server cannot determine the status of the certificate.
	//
	// example:
	//
	// good
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClientCertificateStatusResponseBodyCertificateStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateStatusResponseBodyCertificateStatus) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusResponseBodyCertificateStatus) SetRevokeTime(v int64) *DescribeClientCertificateStatusResponseBodyCertificateStatus {
	s.RevokeTime = &v
	return s
}

func (s *DescribeClientCertificateStatusResponseBodyCertificateStatus) SetSerialNumber(v string) *DescribeClientCertificateStatusResponseBodyCertificateStatus {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientCertificateStatusResponseBodyCertificateStatus) SetStatus(v string) *DescribeClientCertificateStatusResponseBodyCertificateStatus {
	s.Status = &v
	return s
}

type DescribeClientCertificateStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientCertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientCertificateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusResponse) SetHeaders(v map[string]*string) *DescribeClientCertificateStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientCertificateStatusResponse) SetStatusCode(v int32) *DescribeClientCertificateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientCertificateStatusResponse) SetBody(v *DescribeClientCertificateStatusResponseBody) *DescribeClientCertificateStatusResponse {
	s.Body = v
	return s
}

type GetCAInstanceStatusRequest struct {
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the private CA instance.
	//
	// >  After you purchase a private CA instance by using the [SSL Certificates Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist), you can click **Details*	- for the private CA instance on the **Private Certificates*	- page to query the ID of the private CA instance.
	//
	// example:
	//
	// cas-member-0hmi****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCAInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCAInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCAInstanceStatusRequest) SetIdentifier(v string) *GetCAInstanceStatusRequest {
	s.Identifier = &v
	return s
}

func (s *GetCAInstanceStatusRequest) SetInstanceId(v string) *GetCAInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

type GetCAInstanceStatusResponseBody struct {
	// The status information of the private CA instance.
	InstanceStatusList []*GetCAInstanceStatusResponseBodyInstanceStatusList `json:"InstanceStatusList,omitempty" xml:"InstanceStatusList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25589516-2A56-5159-AB88-4A1D9824E183
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCAInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCAInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCAInstanceStatusResponseBody) SetInstanceStatusList(v []*GetCAInstanceStatusResponseBodyInstanceStatusList) *GetCAInstanceStatusResponseBody {
	s.InstanceStatusList = v
	return s
}

func (s *GetCAInstanceStatusResponseBody) SetRequestId(v string) *GetCAInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetCAInstanceStatusResponseBodyInstanceStatusList struct {
	// The expiration date of the private CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// >  This parameter is returned only when the value of the **Status*	- parameter is **USED*	- or **REVOKE**. The value USED indicates that the private CA instance is enabled, and the value REVOKE indicates that the instance is revoked.
	//
	// example:
	//
	// 1792944000000
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The issuance date of the private CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// >  This parameter is returned only when the value of the **Status*	- parameter is **USED*	- or **REVOKE**. The value USED indicates that the private CA instance is enabled, and the value REVOKE indicates that the instance is revoked.
	//
	// example:
	//
	// 1635177600000
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The number of certificates that are issued by using the private CA instance.
	//
	// example:
	//
	// 1
	CertIssuedCount *int32 `json:"CertIssuedCount,omitempty" xml:"CertIssuedCount,omitempty"`
	// The number of certificates that can be issued by using the private CA instance.
	//
	// For a private root CA instance whose **Type*	- is **ROOT**, this parameter indicates the number of intermediate CA certificates that can be issued.
	//
	// For a private intermediate CA instance whose **Type*	- is **SUB_ROOT**, this parameter indicates the total number of client certificates and server certificates that can be issued
	//
	// example:
	//
	// 10
	CertTotalCount *int32 `json:"CertTotalCount,omitempty" xml:"CertTotalCount,omitempty"`
	// The unique identifier of the private CA certificate.
	//
	// >  This parameter is returned only when the value of the **Status*	- parameter is **USED*	- or **REVOKE**. The value USED indicates that the private CA instance is enabled, and the value REVOKE indicates that the instance is revoked.
	//
	// example:
	//
	// a7bb2dd212a2112128cd5cc9b753****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the private CA instance.
	//
	// example:
	//
	// cas-member-0hmi****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the private CA instance. Valid values:
	//
	// 	- **BUY**: The private CA instance is purchased but is not enabled.
	//
	// 	- **USED**: The private CA instance is enabled.
	//
	// 	- **REFUND**: The private CA instance is refunded.
	//
	// 	- **REVOKE**: The private CA instance is revoked.
	//
	// example:
	//
	// USED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the private CA instance. Valid values:
	//
	// 	- **ROOT**: root CA instance
	//
	// 	- **SUB_ROOT**: intermediate CA instance
	//
	// example:
	//
	// ROOT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The expiration date of the private CA instance. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// >  This parameter corresponds to the duration that you select when you purchase the private CA instance. The duration indicates the subscription period of the Private Certificate Authority (PCA) service.
	//
	// example:
	//
	// 1637251200000
	UseExpireTime *int64 `json:"UseExpireTime,omitempty" xml:"UseExpireTime,omitempty"`
}

func (s GetCAInstanceStatusResponseBodyInstanceStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetCAInstanceStatusResponseBodyInstanceStatusList) GoString() string {
	return s.String()
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetAfterTime(v int64) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.AfterTime = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetBeforeTime(v int64) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.BeforeTime = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetCertIssuedCount(v int32) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.CertIssuedCount = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetCertTotalCount(v int32) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.CertTotalCount = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetIdentifier(v string) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.Identifier = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetInstanceId(v string) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.InstanceId = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetStatus(v string) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.Status = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetType(v string) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.Type = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetUseExpireTime(v int64) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.UseExpireTime = &v
	return s
}

type GetCAInstanceStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCAInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCAInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCAInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCAInstanceStatusResponse) SetHeaders(v map[string]*string) *GetCAInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCAInstanceStatusResponse) SetStatusCode(v int32) *GetCAInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCAInstanceStatusResponse) SetBody(v *GetCAInstanceStatusResponseBody) *GetCAInstanceStatusResponse {
	s.Body = v
	return s
}

type ListCertRequest struct {
	// example:
	//
	// 2024-05-13 12:59:45
	AfterDate *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// example:
	//
	// 2025-09-04
	BeforeDate *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1ef79512-569b-6a4e-9105-9b91473562f7
	InstanceUuid *string `json:"InstanceUuid,omitempty" xml:"InstanceUuid,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 50
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// CLIENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCertRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCertRequest) GoString() string {
	return s.String()
}

func (s *ListCertRequest) SetAfterDate(v string) *ListCertRequest {
	s.AfterDate = &v
	return s
}

func (s *ListCertRequest) SetBeforeDate(v string) *ListCertRequest {
	s.BeforeDate = &v
	return s
}

func (s *ListCertRequest) SetCurrentPage(v int32) *ListCertRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCertRequest) SetInstanceUuid(v string) *ListCertRequest {
	s.InstanceUuid = &v
	return s
}

func (s *ListCertRequest) SetMaxResults(v int32) *ListCertRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCertRequest) SetNextToken(v string) *ListCertRequest {
	s.NextToken = &v
	return s
}

func (s *ListCertRequest) SetShowSize(v int32) *ListCertRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCertRequest) SetStatus(v string) *ListCertRequest {
	s.Status = &v
	return s
}

func (s *ListCertRequest) SetType(v string) *ListCertRequest {
	s.Type = &v
	return s
}

type ListCertResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []*ListCertResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 50
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCertResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertResponseBody) SetCurrentPage(v int32) *ListCertResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCertResponseBody) SetList(v []*ListCertResponseBodyList) *ListCertResponseBody {
	s.List = v
	return s
}

func (s *ListCertResponseBody) SetMaxResults(v int32) *ListCertResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCertResponseBody) SetNextToken(v string) *ListCertResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCertResponseBody) SetPageCount(v int32) *ListCertResponseBody {
	s.PageCount = &v
	return s
}

func (s *ListCertResponseBody) SetRequestId(v string) *ListCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertResponseBody) SetShowSize(v int32) *ListCertResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCertResponseBody) SetTotalCount(v int64) *ListCertResponseBody {
	s.TotalCount = &v
	return s
}

type ListCertResponseBodyList struct {
	// example:
	//
	// 2024-05-13 12:59:45
	AfterDate *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// example:
	//
	// 1728921600000
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 2026-05-19
	BeforeDate *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// example:
	//
	// 1728921600000
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// example:
	//
	// Server
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// example:
	//
	// www.kfsjn.xyz
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// {\\"appId\\":\\"APP_PFHMIGUHKDUW6S3N7ZL2\\"}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// 1806958
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1ef539a8-1e1f-6b88-8c11-21cf01a203e9
	Identifier    *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	KeyExportable *bool   `json:"KeyExportable,omitempty" xml:"KeyExportable,omitempty"`
	Organization  *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// example:
	//
	// 3a3ee3c3597d675e
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// example:
	//
	// complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// SubjectDn
	SubjectDn *string   `json:"SubjectDn,omitempty" xml:"SubjectDn,omitempty"`
	Tags      []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListCertResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListCertResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListCertResponseBodyList) SetAfterDate(v string) *ListCertResponseBodyList {
	s.AfterDate = &v
	return s
}

func (s *ListCertResponseBodyList) SetAfterTime(v int64) *ListCertResponseBodyList {
	s.AfterTime = &v
	return s
}

func (s *ListCertResponseBodyList) SetAlgorithm(v string) *ListCertResponseBodyList {
	s.Algorithm = &v
	return s
}

func (s *ListCertResponseBodyList) SetAliasName(v string) *ListCertResponseBodyList {
	s.AliasName = &v
	return s
}

func (s *ListCertResponseBodyList) SetBeforeDate(v string) *ListCertResponseBodyList {
	s.BeforeDate = &v
	return s
}

func (s *ListCertResponseBodyList) SetBeforeTime(v int64) *ListCertResponseBodyList {
	s.BeforeTime = &v
	return s
}

func (s *ListCertResponseBodyList) SetCertificateType(v string) *ListCertResponseBodyList {
	s.CertificateType = &v
	return s
}

func (s *ListCertResponseBodyList) SetCommonName(v string) *ListCertResponseBodyList {
	s.CommonName = &v
	return s
}

func (s *ListCertResponseBodyList) SetExtra(v string) *ListCertResponseBodyList {
	s.Extra = &v
	return s
}

func (s *ListCertResponseBodyList) SetId(v string) *ListCertResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListCertResponseBodyList) SetIdentifier(v string) *ListCertResponseBodyList {
	s.Identifier = &v
	return s
}

func (s *ListCertResponseBodyList) SetKeyExportable(v bool) *ListCertResponseBodyList {
	s.KeyExportable = &v
	return s
}

func (s *ListCertResponseBodyList) SetOrganization(v string) *ListCertResponseBodyList {
	s.Organization = &v
	return s
}

func (s *ListCertResponseBodyList) SetOrganizationUnit(v string) *ListCertResponseBodyList {
	s.OrganizationUnit = &v
	return s
}

func (s *ListCertResponseBodyList) SetSerialNumber(v string) *ListCertResponseBodyList {
	s.SerialNumber = &v
	return s
}

func (s *ListCertResponseBodyList) SetStatus(v string) *ListCertResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListCertResponseBodyList) SetSubjectDn(v string) *ListCertResponseBodyList {
	s.SubjectDn = &v
	return s
}

func (s *ListCertResponseBodyList) SetTags(v []*string) *ListCertResponseBodyList {
	s.Tags = v
	return s
}

type ListCertResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCertResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCertResponse) GoString() string {
	return s.String()
}

func (s *ListCertResponse) SetHeaders(v map[string]*string) *ListCertResponse {
	s.Headers = v
	return s
}

func (s *ListCertResponse) SetStatusCode(v int32) *ListCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCertResponse) SetBody(v *ListCertResponseBody) *ListCertResponse {
	s.Body = v
	return s
}

type ListClientCertificateRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The unique identifier of the client certificate or the server certificate that you want to query.
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers of all client certificates and server certificates.
	//
	// example:
	//
	// 190ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The number of certificates to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListClientCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *ListClientCertificateRequest) SetCurrentPage(v int32) *ListClientCertificateRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClientCertificateRequest) SetIdentifier(v string) *ListClientCertificateRequest {
	s.Identifier = &v
	return s
}

func (s *ListClientCertificateRequest) SetShowSize(v int32) *ListClientCertificateRequest {
	s.ShowSize = &v
	return s
}

type ListClientCertificateResponseBody struct {
	// An array that consists of the details about all client certificates and server certificates.
	CertificateList []*ListClientCertificateResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// The page number of the current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of certificates that are returned per page.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The number of client certificates and server certificates that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClientCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientCertificateResponseBody) SetCertificateList(v []*ListClientCertificateResponseBodyCertificateList) *ListClientCertificateResponseBody {
	s.CertificateList = v
	return s
}

func (s *ListClientCertificateResponseBody) SetCurrentPage(v int32) *ListClientCertificateResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListClientCertificateResponseBody) SetPageCount(v int32) *ListClientCertificateResponseBody {
	s.PageCount = &v
	return s
}

func (s *ListClientCertificateResponseBody) SetRequestId(v string) *ListClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientCertificateResponseBody) SetShowSize(v int32) *ListClientCertificateResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListClientCertificateResponseBody) SetTotalCount(v int64) *ListClientCertificateResponseBody {
	s.TotalCount = &v
	return s
}

type ListClientCertificateResponseBodyCertificateList struct {
	// The expiration date of the certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The type of the encryption algorithm of the certificate. Valid values:
	//
	// 	- **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	//
	// 	- **ECC**: the elliptic curve cryptography (ECC) algorithm.
	//
	// 	- **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance date of the certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1634283958000
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **CLIENT**: client certificate
	//
	// 	- **SERVER**: server certificate
	//
	// example:
	//
	// SERVER
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The validity period of the certificate. Unit: days.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The unique identifier of the certificate.
	//
	// example:
	//
	// d3b95700998e47afc4d95f886579****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the certificate.
	//
	// example:
	//
	// 4096
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the certificate.
	//
	// example:
	//
	// d3b95700998e47afc4d95f886579****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Alibaba Cloud Computing Co., Ltd.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department in the organization. The organization is associated with the intermediate certificate authority (CA) certificate from which the certificate is issued.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate certificate from which the client certificate is issued.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The subject alternative name (SAN) extension of the certificate. The value indicates additional information, including the additional domain names or IP addresses that are associated with the certificate.
	//
	// The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that corresponds to a SAN extension. A SAN extension struct contains the following parameters:
	//
	// 	- **Type**: the type of the extension. Data type: integer. Valid values:
	//
	//     	- **1**: an email address
	//
	//     	- **2**: a domain name
	//
	//     	- **6**: a Uniform Resource Identifier (URI)
	//
	//     	- **7**: an IP address
	//
	// 	- **Value**: the value of the extension. Data type: string.
	//
	// example:
	//
	// [ {"Type": 7, "Value": "192.0.XX.XX"}, {"Type": 2, "Value": "www.aliyundoc.com"}, ]
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 62b2b943a32d96883a6650e672ea0276****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	//
	// example:
	//
	// 14dcc8afc7578e1fcec36d658f7e20de18f6957bbac42b373a66bc9de4e9****
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the certificate.
	//
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
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
	// The distinguished name (DN) extension of the certificate, which indicates the user of the certificate. The DN extension includes the following information:
	//
	// 	- **C**: the country
	//
	// 	- **O**: the organization
	//
	// 	- **OU**: the department
	//
	// 	- **L**: the city
	//
	// 	- **ST**: the province, municipality, or autonomous region
	//
	// 	- **CN**: the common name
	//
	// example:
	//
	// C=CN,O=Alibaba Cloud Computing Co., Ltd.,OU=Security,L=Hangzhou,ST=Zhejiang,CN=Aliyun
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// The content of the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  ...... -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s ListClientCertificateResponseBodyCertificateList) String() string {
	return tea.Prettify(s)
}

func (s ListClientCertificateResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *ListClientCertificateResponseBodyCertificateList) SetAfterDate(v int64) *ListClientCertificateResponseBodyCertificateList {
	s.AfterDate = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetAlgorithm(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetBeforeDate(v int64) *ListClientCertificateResponseBodyCertificateList {
	s.BeforeDate = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetCertificateType(v string) *ListClientCertificateResponseBodyCertificateList {
	s.CertificateType = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetCommonName(v string) *ListClientCertificateResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetCountryCode(v string) *ListClientCertificateResponseBodyCertificateList {
	s.CountryCode = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetDays(v int32) *ListClientCertificateResponseBodyCertificateList {
	s.Days = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetIdentifier(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Identifier = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetKeySize(v int32) *ListClientCertificateResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetLocality(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Locality = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetMd5(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Md5 = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetOrganization(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Organization = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetOrganizationUnit(v string) *ListClientCertificateResponseBodyCertificateList {
	s.OrganizationUnit = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetParentIdentifier(v string) *ListClientCertificateResponseBodyCertificateList {
	s.ParentIdentifier = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetSans(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Sans = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetSerialNumber(v string) *ListClientCertificateResponseBodyCertificateList {
	s.SerialNumber = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetSha2(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Sha2 = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetSignAlgorithm(v string) *ListClientCertificateResponseBodyCertificateList {
	s.SignAlgorithm = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetState(v string) *ListClientCertificateResponseBodyCertificateList {
	s.State = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetStatus(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Status = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetSubjectDN(v string) *ListClientCertificateResponseBodyCertificateList {
	s.SubjectDN = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetX509Certificate(v string) *ListClientCertificateResponseBodyCertificateList {
	s.X509Certificate = &v
	return s
}

type ListClientCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *ListClientCertificateResponse) SetHeaders(v map[string]*string) *ListClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *ListClientCertificateResponse) SetStatusCode(v int32) *ListClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientCertificateResponse) SetBody(v *ListClientCertificateResponseBody) *ListClientCertificateResponse {
	s.Body = v
	return s
}

type ListRevokeCertificateRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of revoked certificates to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListRevokeCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRevokeCertificateRequest) GoString() string {
	return s.String()
}

func (s *ListRevokeCertificateRequest) SetCurrentPage(v int32) *ListRevokeCertificateRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListRevokeCertificateRequest) SetShowSize(v int32) *ListRevokeCertificateRequest {
	s.ShowSize = &v
	return s
}

type ListRevokeCertificateResponseBody struct {
	// An array that consists of the details about the revoked client certificates or server certificates.
	CertificateList []*ListRevokeCertificateResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// The page number of the current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of revoked certificates that are returned per page.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of revoked client certificates and server certificates that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRevokeCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRevokeCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ListRevokeCertificateResponseBody) SetCertificateList(v []*ListRevokeCertificateResponseBodyCertificateList) *ListRevokeCertificateResponseBody {
	s.CertificateList = v
	return s
}

func (s *ListRevokeCertificateResponseBody) SetCurrentPage(v int32) *ListRevokeCertificateResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListRevokeCertificateResponseBody) SetPageCount(v int32) *ListRevokeCertificateResponseBody {
	s.PageCount = &v
	return s
}

func (s *ListRevokeCertificateResponseBody) SetRequestId(v string) *ListRevokeCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRevokeCertificateResponseBody) SetShowSize(v int32) *ListRevokeCertificateResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListRevokeCertificateResponseBody) SetTotalCount(v int64) *ListRevokeCertificateResponseBody {
	s.TotalCount = &v
	return s
}

type ListRevokeCertificateResponseBodyCertificateList struct {
	// The expiration date of the certificate. The date is in the `yyyy-MM-ddT00:00Z` format. For example, the value `2021-12-31T00:00Z` indicates December 31, 2021.
	//
	// example:
	//
	// 2021-12-31T00:00Z
	AfterDate *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The type of the encryption algorithm of the certificate. Valid values:
	//
	// 	- **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	//
	// 	- **ECC**: the elliptic curve cryptography (ECC) algorithm.
	//
	// 	- **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance date of the certificate. The date is in the `yyyy-MM-ddT00:00Z` format. For example, the value `2021-01-01T00:00Z` indicates January 1, 2021.
	//
	// example:
	//
	// 2021-01-01T00:00Z
	BeforeDate *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the certificate.
	//
	// example:
	//
	// SERVER
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 05e148d8d3ecc9976d9ecd2b2f25****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the certificate.
	//
	// example:
	//
	// 4096
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the certificate.
	//
	// example:
	//
	// 05e148d8d3ecc9976d9ecd2b2f25****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Alibaba Cloud Computing Co., Ltd.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department in the organization. The organization is associated with the intermediate certificate authority (CA) certificate from which the certificate is issued.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The identifier of the root certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The date on which the certificate was revoked. The value is in the `yyyy-MM-ddT00:00Z` format. For example, the value `2021-09-01T00:00Z` indicates September 1, 2021.
	//
	// example:
	//
	// 2021-09-01T00:00Z
	RevokeDate *string `json:"RevokeDate,omitempty" xml:"RevokeDate,omitempty"`
	// The subject alternative name (SAN) extension of the certificate.
	//
	// The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that corresponds to a SAN extension. A SAN extension struct contains the following parameters:
	//
	// 	- **Type**: the type of the extension. Data type: integer. Valid values:
	//
	//     	- **1**: an email address
	//
	//     	- **2**: a domain name
	//
	//     	- **6**: a Uniform Resource Identifier (URI)
	//
	//     	- **7**: an IP address
	//
	// 	- **Value**: the value of the extension. Data type: string.
	//
	// example:
	//
	// [ {"Type": 7, "Value": "192.0.XX.XX"}, {"Type": 2, "Value": "www.aliyundoc.com"}, ]
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 168b12c42e62339f8d2340ff530f9365****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	//
	// example:
	//
	// b60eff7e04323ff662f9ab5e6986f849f626a9c7bf2c59dcc752fa23779a****
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the certificate.
	//
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status.
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The distinguished name (DN) extension of the certificate, which indicates the user of the certificate. The DN extension includes the following information:
	//
	// 	- **C**: the country
	//
	// 	- **O**: the organization
	//
	// 	- **OU**: the department
	//
	// 	- **L**: the city
	//
	// 	- **ST**: the province, municipality, or autonomous region
	//
	// 	- **CN**: the common name
	//
	// example:
	//
	// C=CN,O=Alibaba Cloud Computing Co., Ltd.,OU=Security,L=ZheJiang,ST=HangZhou,CN=aliyundoc.com
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
}

func (s ListRevokeCertificateResponseBodyCertificateList) String() string {
	return tea.Prettify(s)
}

func (s ListRevokeCertificateResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetAfterDate(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.AfterDate = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetAlgorithm(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetBeforeDate(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.BeforeDate = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetCertificateType(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.CertificateType = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetCommonName(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetCountryCode(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.CountryCode = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetIdentifier(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Identifier = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetKeySize(v int32) *ListRevokeCertificateResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetLocality(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Locality = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetMd5(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Md5 = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetOrganization(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Organization = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetOrganizationUnit(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.OrganizationUnit = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetParentIdentifier(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.ParentIdentifier = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetRevokeDate(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.RevokeDate = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSans(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Sans = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSerialNumber(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.SerialNumber = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSha2(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Sha2 = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSignAlgorithm(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.SignAlgorithm = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetState(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.State = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetStatus(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Status = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSubjectDN(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.SubjectDN = &v
	return s
}

type ListRevokeCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRevokeCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRevokeCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRevokeCertificateResponse) GoString() string {
	return s.String()
}

func (s *ListRevokeCertificateResponse) SetHeaders(v map[string]*string) *ListRevokeCertificateResponse {
	s.Headers = v
	return s
}

func (s *ListRevokeCertificateResponse) SetStatusCode(v int32) *ListRevokeCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRevokeCertificateResponse) SetBody(v *ListRevokeCertificateResponseBody) *ListRevokeCertificateResponse {
	s.Body = v
	return s
}

type UpdateCACertificateStatusRequest struct {
	// The unique identifier of the CA certificate whose status you want to change.
	//
	// >  You can call the [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) operation to query the unique identifiers of all CA certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The state to which you want to change the CA certificate. Set to the value to **REVOKE**. After this operation is called, the status of the CA certificate is changed to **REVOKE**.
	//
	// >  You can call this operation only if the status of a CA certificate is **ISSUE**. You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the status of a CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// REVOKE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateCACertificateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCACertificateStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCACertificateStatusRequest) SetIdentifier(v string) *UpdateCACertificateStatusRequest {
	s.Identifier = &v
	return s
}

func (s *UpdateCACertificateStatusRequest) SetStatus(v string) *UpdateCACertificateStatusRequest {
	s.Status = &v
	return s
}

type UpdateCACertificateStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCACertificateStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCACertificateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCACertificateStatusResponseBody) SetRequestId(v string) *UpdateCACertificateStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCACertificateStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCACertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCACertificateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCACertificateStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCACertificateStatusResponse) SetHeaders(v map[string]*string) *UpdateCACertificateStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCACertificateStatusResponse) SetStatusCode(v int32) *UpdateCACertificateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCACertificateStatusResponse) SetBody(v *UpdateCACertificateStatusResponseBody) *UpdateCACertificateStatusResponse {
	s.Body = v
	return s
}

type UploadPcaCertToCasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 59425,59426
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s UploadPcaCertToCasRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadPcaCertToCasRequest) GoString() string {
	return s.String()
}

func (s *UploadPcaCertToCasRequest) SetIds(v string) *UploadPcaCertToCasRequest {
	s.Ids = &v
	return s
}

type UploadPcaCertToCasResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadPcaCertToCasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadPcaCertToCasResponseBody) GoString() string {
	return s.String()
}

func (s *UploadPcaCertToCasResponseBody) SetRequestId(v string) *UploadPcaCertToCasResponseBody {
	s.RequestId = &v
	return s
}

type UploadPcaCertToCasResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadPcaCertToCasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadPcaCertToCasResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadPcaCertToCasResponse) GoString() string {
	return s.String()
}

func (s *UploadPcaCertToCasResponse) SetHeaders(v map[string]*string) *UploadPcaCertToCasResponse {
	s.Headers = v
	return s
}

func (s *UploadPcaCertToCasResponse) SetStatusCode(v int32) *UploadPcaCertToCasResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadPcaCertToCasResponse) SetBody(v *UploadPcaCertToCasResponseBody) *UploadPcaCertToCasResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":                 tea.String("cas.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("cas.aliyuncs.com"),
		"ap-southeast-3":              tea.String("cas.aliyuncs.com"),
		"ap-southeast-5":              tea.String("cas.aliyuncs.com"),
		"cn-beijing":                  tea.String("cas.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cas.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cas.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cas.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cas.aliyuncs.com"),
		"cn-chengdu":                  tea.String("cas.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cas.aliyuncs.com"),
		"cn-fujian":                   tea.String("cas.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cas.aliyuncs.com"),
		"cn-hongkong":                 tea.String("cas.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cas.aliyuncs.com"),
		"cn-huhehaote":                tea.String("cas.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("cas.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("cas.aliyuncs.com"),
		"cn-qingdao":                  tea.String("cas.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cas.aliyuncs.com"),
		"cn-shanghai":                 tea.String("cas.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cas.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cas.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cas.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cas.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cas.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cas.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cas.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("cas.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cas.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("cas.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cas.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("cas.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cas.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cas.aliyuncs.com"),
		"eu-west-1":                   tea.String("cas.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cas.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cas.aliyuncs.com"),
		"us-east-1":                   tea.String("cas.aliyuncs.com"),
		"us-west-1":                   tea.String("cas.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues a client certificate by using a system-generated certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~CreateRootCACertificate~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~CreateRootCACertificate~~) operation. Only intermediate CA certificates can issue client certificates.
//
// ## QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientCertificateResponse
func (client *Client) CreateClientCertificateWithOptions(request *CreateClientCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AfterTime)) {
		query["AfterTime"] = request.AfterTime
	}

	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.BeforeTime)) {
		query["BeforeTime"] = request.BeforeTime
	}

	if !tea.BoolValue(util.IsUnset(request.CommonName)) {
		query["CommonName"] = request.CommonName
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.Days)) {
		query["Days"] = request.Days
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCrl)) {
		query["EnableCrl"] = request.EnableCrl
	}

	if !tea.BoolValue(util.IsUnset(request.Immediately)) {
		query["Immediately"] = request.Immediately
	}

	if !tea.BoolValue(util.IsUnset(request.Locality)) {
		query["Locality"] = request.Locality
	}

	if !tea.BoolValue(util.IsUnset(request.Months)) {
		query["Months"] = request.Months
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationUnit)) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ParentIdentifier)) {
		query["ParentIdentifier"] = request.ParentIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SanType)) {
		query["SanType"] = request.SanType
	}

	if !tea.BoolValue(util.IsUnset(request.SanValue)) {
		query["SanValue"] = request.SanValue
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Years)) {
		query["Years"] = request.Years
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClientCertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues a client certificate by using a system-generated certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~CreateRootCACertificate~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~CreateRootCACertificate~~) operation. Only intermediate CA certificates can issue client certificates.
//
// ## QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateClientCertificateRequest
//
// @return CreateClientCertificateResponse
func (client *Client) CreateClientCertificate(request *CreateClientCertificateRequest) (_result *CreateClientCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClientCertificateResponse{}
	_body, _err := client.CreateClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues a client certificate by using a custom certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue client certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateClientCertificateWithCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientCertificateWithCsrResponse
func (client *Client) CreateClientCertificateWithCsrWithOptions(request *CreateClientCertificateWithCsrRequest, runtime *util.RuntimeOptions) (_result *CreateClientCertificateWithCsrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AfterTime)) {
		query["AfterTime"] = request.AfterTime
	}

	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.BeforeTime)) {
		query["BeforeTime"] = request.BeforeTime
	}

	if !tea.BoolValue(util.IsUnset(request.CommonName)) {
		query["CommonName"] = request.CommonName
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.Csr)) {
		query["Csr"] = request.Csr
	}

	if !tea.BoolValue(util.IsUnset(request.Days)) {
		query["Days"] = request.Days
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCrl)) {
		query["EnableCrl"] = request.EnableCrl
	}

	if !tea.BoolValue(util.IsUnset(request.Immediately)) {
		query["Immediately"] = request.Immediately
	}

	if !tea.BoolValue(util.IsUnset(request.Locality)) {
		query["Locality"] = request.Locality
	}

	if !tea.BoolValue(util.IsUnset(request.Months)) {
		query["Months"] = request.Months
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationUnit)) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ParentIdentifier)) {
		query["ParentIdentifier"] = request.ParentIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SanType)) {
		query["SanType"] = request.SanType
	}

	if !tea.BoolValue(util.IsUnset(request.SanValue)) {
		query["SanValue"] = request.SanValue
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Years)) {
		query["Years"] = request.Years
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClientCertificateWithCsr"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClientCertificateWithCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues a client certificate by using a custom certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue client certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateClientCertificateWithCsrRequest
//
// @return CreateClientCertificateWithCsrResponse
func (client *Client) CreateClientCertificateWithCsr(request *CreateClientCertificateWithCsrRequest) (_result *CreateClientCertificateWithCsrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClientCertificateWithCsrResponse{}
	_body, _err := client.CreateClientCertificateWithCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues a certificate based on the specified key usage, extended key usage, and name and alias of the entity that uses the certificate.
//
// Description:
//
// By default, the name of the entity is obtained from the certificate signing request (CSR) of the certificate that you want to issue. If you specify a different name for the entity, the name of the entity in the CSR becomes invalid. The specified name is used to issue the certificate.
//
// You must specify the key usage and extended key usage based on the certificate type. The following list describes common certificate types:
//
//   - Server certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: serverAuth
//
//   - Client certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: clientAuth
//
//   - Mutual Transport Layer Security (TLS) authentication certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: serverAuth or clientAuth
//
//   - Email certificate
//
// Key usage: digitalSignature or contentCommitment
//
// Extended key usage: emailProtection
//
// Note: Compliant certificate authorities (CAs) are managed by third-party authorities. This operation is not supported for compliant CAs.
//
// @param request - CreateCustomCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomCertificateResponse
func (client *Client) CreateCustomCertificateWithOptions(request *CreateCustomCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateCustomCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiPassthrough)) {
		query["ApiPassthrough"] = request.ApiPassthrough
	}

	if !tea.BoolValue(util.IsUnset(request.Csr)) {
		query["Csr"] = request.Csr
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCrl)) {
		query["EnableCrl"] = request.EnableCrl
	}

	if !tea.BoolValue(util.IsUnset(request.Immediately)) {
		query["Immediately"] = request.Immediately
	}

	if !tea.BoolValue(util.IsUnset(request.ParentIdentifier)) {
		query["ParentIdentifier"] = request.ParentIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.Validity)) {
		query["Validity"] = request.Validity
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomCertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues a certificate based on the specified key usage, extended key usage, and name and alias of the entity that uses the certificate.
//
// Description:
//
// By default, the name of the entity is obtained from the certificate signing request (CSR) of the certificate that you want to issue. If you specify a different name for the entity, the name of the entity in the CSR becomes invalid. The specified name is used to issue the certificate.
//
// You must specify the key usage and extended key usage based on the certificate type. The following list describes common certificate types:
//
//   - Server certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: serverAuth
//
//   - Client certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: clientAuth
//
//   - Mutual Transport Layer Security (TLS) authentication certificate
//
// Key usage: digitalSignature or keyEncipherment
//
// Extended key usage: serverAuth or clientAuth
//
//   - Email certificate
//
// Key usage: digitalSignature or contentCommitment
//
// Extended key usage: emailProtection
//
// Note: Compliant certificate authorities (CAs) are managed by third-party authorities. This operation is not supported for compliant CAs.
//
// @param request - CreateCustomCertificateRequest
//
// @return CreateCustomCertificateResponse
func (client *Client) CreateCustomCertificate(request *CreateCustomCertificateRequest) (_result *CreateCustomCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomCertificateResponse{}
	_body, _err := client.CreateCustomCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes a client certificate or a server certificate.
//
// Description:
//
// After a client certificate or a server certificate is revoked, the client or the server on which the certificate is installed cannot establish HTTPS connections with other devices.
//
// After a client certificate or a server certificate is revoked, you can call the [DeleteClientCertificate](https://help.aliyun.com/document_detail/330880.html) operation to permanently delete the certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateRevokeClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRevokeClientCertificateResponse
func (client *Client) CreateRevokeClientCertificateWithOptions(request *CreateRevokeClientCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateRevokeClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRevokeClientCertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRevokeClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a client certificate or a server certificate.
//
// Description:
//
// After a client certificate or a server certificate is revoked, the client or the server on which the certificate is installed cannot establish HTTPS connections with other devices.
//
// After a client certificate or a server certificate is revoked, you can call the [DeleteClientCertificate](https://help.aliyun.com/document_detail/330880.html) operation to permanently delete the certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateRevokeClientCertificateRequest
//
// @return CreateRevokeClientCertificateResponse
func (client *Client) CreateRevokeClientCertificate(request *CreateRevokeClientCertificateRequest) (_result *CreateRevokeClientCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRevokeClientCertificateResponse{}
	_body, _err := client.CreateRevokeClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a root certificate authority (CA) certificate.
//
// Description:
//
// You can call the CreateRootCACertificate operation to create a self-signed root CA certificate. A root CA certificate is the trust anchor in a chain of trust for private certificates that are used within an enterprise. You must create a root CA certificate before you can use the root CA certificate to issue intermediate CA certificates. Then, you can use the intermediate CA certificates to issue client certificates and server certificates.
//
// Before you call this operation, make sure that you have purchased a private root CA instance by using the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateRootCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRootCACertificateResponse
func (client *Client) CreateRootCACertificateWithOptions(request *CreateRootCACertificateRequest, runtime *util.RuntimeOptions) (_result *CreateRootCACertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CommonName)) {
		query["CommonName"] = request.CommonName
	}

	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		query["CountryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.Locality)) {
		query["Locality"] = request.Locality
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationUnit)) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Years)) {
		query["Years"] = request.Years
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRootCACertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRootCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a root certificate authority (CA) certificate.
//
// Description:
//
// You can call the CreateRootCACertificate operation to create a self-signed root CA certificate. A root CA certificate is the trust anchor in a chain of trust for private certificates that are used within an enterprise. You must create a root CA certificate before you can use the root CA certificate to issue intermediate CA certificates. Then, you can use the intermediate CA certificates to issue client certificates and server certificates.
//
// Before you call this operation, make sure that you have purchased a private root CA instance by using the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateRootCACertificateRequest
//
// @return CreateRootCACertificateResponse
func (client *Client) CreateRootCACertificate(request *CreateRootCACertificateRequest) (_result *CreateRootCACertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRootCACertificateResponse{}
	_body, _err := client.CreateRootCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues a server certificate by using a system-generated certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue server certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateServerCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServerCertificateResponse
func (client *Client) CreateServerCertificateWithOptions(request *CreateServerCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateServerCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AfterTime)) {
		query["AfterTime"] = request.AfterTime
	}

	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.BeforeTime)) {
		query["BeforeTime"] = request.BeforeTime
	}

	if !tea.BoolValue(util.IsUnset(request.CommonName)) {
		query["CommonName"] = request.CommonName
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.Days)) {
		query["Days"] = request.Days
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCrl)) {
		query["EnableCrl"] = request.EnableCrl
	}

	if !tea.BoolValue(util.IsUnset(request.Immediately)) {
		query["Immediately"] = request.Immediately
	}

	if !tea.BoolValue(util.IsUnset(request.Locality)) {
		query["Locality"] = request.Locality
	}

	if !tea.BoolValue(util.IsUnset(request.Months)) {
		query["Months"] = request.Months
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationUnit)) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ParentIdentifier)) {
		query["ParentIdentifier"] = request.ParentIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Years)) {
		query["Years"] = request.Years
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServerCertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServerCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues a server certificate by using a system-generated certificate signing request (CSR) file.
//
// Description:
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue server certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateServerCertificateRequest
//
// @return CreateServerCertificateResponse
func (client *Client) CreateServerCertificate(request *CreateServerCertificateRequest) (_result *CreateServerCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServerCertificateResponse{}
	_body, _err := client.CreateServerCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issues a server certificate by using a custom certificate signing request (CSR) file.
//
// Description:
//
// ## Usage notes
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue server certificates.
//
// @param request - CreateServerCertificateWithCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServerCertificateWithCsrResponse
func (client *Client) CreateServerCertificateWithCsrWithOptions(request *CreateServerCertificateWithCsrRequest, runtime *util.RuntimeOptions) (_result *CreateServerCertificateWithCsrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AfterTime)) {
		query["AfterTime"] = request.AfterTime
	}

	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.BeforeTime)) {
		query["BeforeTime"] = request.BeforeTime
	}

	if !tea.BoolValue(util.IsUnset(request.CommonName)) {
		query["CommonName"] = request.CommonName
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.Csr)) {
		query["Csr"] = request.Csr
	}

	if !tea.BoolValue(util.IsUnset(request.Days)) {
		query["Days"] = request.Days
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCrl)) {
		query["EnableCrl"] = request.EnableCrl
	}

	if !tea.BoolValue(util.IsUnset(request.Immediately)) {
		query["Immediately"] = request.Immediately
	}

	if !tea.BoolValue(util.IsUnset(request.Locality)) {
		query["Locality"] = request.Locality
	}

	if !tea.BoolValue(util.IsUnset(request.Months)) {
		query["Months"] = request.Months
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationUnit)) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ParentIdentifier)) {
		query["ParentIdentifier"] = request.ParentIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Years)) {
		query["Years"] = request.Years
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServerCertificateWithCsr"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServerCertificateWithCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues a server certificate by using a custom certificate signing request (CSR) file.
//
// Description:
//
// ## Usage notes
//
// Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](https://help.aliyun.com/document_detail/328093.html) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](https://help.aliyun.com/document_detail/328094.html) operation. Only intermediate CA certificates can be used to issue server certificates.
//
// @param request - CreateServerCertificateWithCsrRequest
//
// @return CreateServerCertificateWithCsrResponse
func (client *Client) CreateServerCertificateWithCsr(request *CreateServerCertificateWithCsrRequest) (_result *CreateServerCertificateWithCsrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServerCertificateWithCsrResponse{}
	_body, _err := client.CreateServerCertificateWithCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an intermediate certificate authority (CA) certificate.
//
// Description:
//
// You can call this operation to issue an intermediate certificate authority (CA) certificate by using an existing root CA certificate. Intermediate CA certificates can be used to issue client certificates and server certificates.
//
// Before you call this operation, make sure that you have issued a root CA certificate by calling the [CreateRootCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateSubCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubCACertificateResponse
func (client *Client) CreateSubCACertificateWithOptions(request *CreateSubCACertificateRequest, runtime *util.RuntimeOptions) (_result *CreateSubCACertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CommonName)) {
		query["CommonName"] = request.CommonName
	}

	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		query["CountryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.CrlDay)) {
		query["CrlDay"] = request.CrlDay
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCrl)) {
		query["EnableCrl"] = request.EnableCrl
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendedKeyUsages)) {
		query["ExtendedKeyUsages"] = request.ExtendedKeyUsages
	}

	if !tea.BoolValue(util.IsUnset(request.Locality)) {
		query["Locality"] = request.Locality
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationUnit)) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ParentIdentifier)) {
		query["ParentIdentifier"] = request.ParentIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.PathLenConstraint)) {
		query["PathLenConstraint"] = request.PathLenConstraint
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Years)) {
		query["Years"] = request.Years
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSubCACertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSubCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an intermediate certificate authority (CA) certificate.
//
// Description:
//
// You can call this operation to issue an intermediate certificate authority (CA) certificate by using an existing root CA certificate. Intermediate CA certificates can be used to issue client certificates and server certificates.
//
// Before you call this operation, make sure that you have issued a root CA certificate by calling the [CreateRootCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateSubCACertificateRequest
//
// @return CreateSubCACertificateResponse
func (client *Client) CreateSubCACertificate(request *CreateSubCACertificateRequest) (_result *CreateSubCACertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubCACertificateResponse{}
	_body, _err := client.CreateSubCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a client certificate or a server certificate that is revoked.
//
// Description:
//
// Before you call this operation, you must call the [CreateRevokeClientCertificate](https://help.aliyun.com/document_detail/330876.html) operation to revoke a client certificate or a server certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientCertificateResponse
func (client *Client) DeleteClientCertificateWithOptions(request *DeleteClientCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteClientCertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a client certificate or a server certificate that is revoked.
//
// Description:
//
// Before you call this operation, you must call the [CreateRevokeClientCertificate](https://help.aliyun.com/document_detail/330876.html) operation to revoke a client certificate or a server certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteClientCertificateRequest
//
// @return DeleteClientCertificateResponse
func (client *Client) DeleteClientCertificate(request *DeleteClientCertificateRequest) (_result *DeleteClientCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClientCertificateResponse{}
	_body, _err := client.DeleteClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a root certificate authority (CA) certificate or an intermediate CA certificate.
//
// Description:
//
// You can call the DescribeCACertificate operation to query the details about a root CA certificate or an intermediate CA certificate by using the unique identifier of the root CA certificate or intermediate CA certificate. The details include the serial number, user information, and content of a CA certificate.
//
// Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate] operation or an intermediate CA certificate by calling the [CreateSubCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCACertificateResponse
func (client *Client) DescribeCACertificateWithOptions(request *DescribeCACertificateRequest, runtime *util.RuntimeOptions) (_result *DescribeCACertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCACertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a root certificate authority (CA) certificate or an intermediate CA certificate.
//
// Description:
//
// You can call the DescribeCACertificate operation to query the details about a root CA certificate or an intermediate CA certificate by using the unique identifier of the root CA certificate or intermediate CA certificate. The details include the serial number, user information, and content of a CA certificate.
//
// Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate] operation or an intermediate CA certificate by calling the [CreateSubCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCACertificateRequest
//
// @return DescribeCACertificateResponse
func (client *Client) DescribeCACertificate(request *DescribeCACertificateRequest) (_result *DescribeCACertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCACertificateResponse{}
	_body, _err := client.DescribeCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of certificate authority (CA) certificates that you create.
//
// Description:
//
// You can call the DescribeCACertificateCount operation to query the number of created CA certificates, which includes root CA certificates and intermediate CA certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCACertificateCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCACertificateCountResponse
func (client *Client) DescribeCACertificateCountWithOptions(runtime *util.RuntimeOptions) (_result *DescribeCACertificateCountResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeCACertificateCount"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCACertificateCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of certificate authority (CA) certificates that you create.
//
// Description:
//
// You can call the DescribeCACertificateCount operation to query the number of created CA certificates, which includes root CA certificates and intermediate CA certificates.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @return DescribeCACertificateCountResponse
func (client *Client) DescribeCACertificateCount() (_result *DescribeCACertificateCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCACertificateCountResponse{}
	_body, _err := client.DescribeCACertificateCountWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about all root certificate authority (CA) certificates and intermediate CA certificates.
//
// Description:
//
// You can call the DescribeCACertificateList operation to perform a paged query of the details about all CA certificates that you create. The details include the unique identifier, serial number, user information, and content of each root CA certificate or intermediate CA certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCACertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCACertificateListResponse
func (client *Client) DescribeCACertificateListWithOptions(request *DescribeCACertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeCACertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaStatus)) {
		query["CaStatus"] = request.CaStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		query["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.IssuerType)) {
		query["IssuerType"] = request.IssuerType
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	if !tea.BoolValue(util.IsUnset(request.ValidStatus)) {
		query["ValidStatus"] = request.ValidStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCACertificateList"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCACertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about all root certificate authority (CA) certificates and intermediate CA certificates.
//
// Description:
//
// You can call the DescribeCACertificateList operation to perform a paged query of the details about all CA certificates that you create. The details include the unique identifier, serial number, user information, and content of each root CA certificate or intermediate CA certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCACertificateListRequest
//
// @return DescribeCACertificateListResponse
func (client *Client) DescribeCACertificateList(request *DescribeCACertificateListRequest) (_result *DescribeCACertificateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCACertificateListResponse{}
	_body, _err := client.DescribeCACertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the encrypted private key of a client certificate or a server certificate.
//
// Description:
//
// ## Usage notes
//
// You can call the DescribeCertificatePrivateKey operation to obtain the encrypted private key of a client certificate or a server certificate. The certificate is issued based on a system-generated certificate signing request (CSR). Before you call this operation, make sure that you have issued a client certificate or a server certificate by calling the following operation:
//
//   - [CreateClientCertificate](https://help.aliyun.com/document_detail/330873.html)
//
//   - [CreateServerCertificate](https://help.aliyun.com/document_detail/330877.html)
//
// To ensure the security of private key transmission, the DescribeCertificatePrivateKey operation encrypts the private key by using the private key password that you specify and returns the encrypted private key. The private key password is a string that is used to encrypt the private key. After you obtain the encrypted private key of the certificate, you can use the following methods to decrypt the private key:
//
//   - If the encryption algorithm of the certificate is RSA, you must run the `openssl rsa -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
//   - If the encryption algorithm of the certificate is ECC, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
//   - If the encryption algorithm of the certificate is SM2, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
// >  You can call the [DescribeClientCertificate] operation to query the encryption algorithm type of a client certificate or a server certificate.
//
// ## Limits
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCertificatePrivateKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertificatePrivateKeyResponse
func (client *Client) DescribeCertificatePrivateKeyWithOptions(request *DescribeCertificatePrivateKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeCertificatePrivateKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptedCode)) {
		query["EncryptedCode"] = request.EncryptedCode
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCertificatePrivateKey"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCertificatePrivateKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the encrypted private key of a client certificate or a server certificate.
//
// Description:
//
// ## Usage notes
//
// You can call the DescribeCertificatePrivateKey operation to obtain the encrypted private key of a client certificate or a server certificate. The certificate is issued based on a system-generated certificate signing request (CSR). Before you call this operation, make sure that you have issued a client certificate or a server certificate by calling the following operation:
//
//   - [CreateClientCertificate](https://help.aliyun.com/document_detail/330873.html)
//
//   - [CreateServerCertificate](https://help.aliyun.com/document_detail/330877.html)
//
// To ensure the security of private key transmission, the DescribeCertificatePrivateKey operation encrypts the private key by using the private key password that you specify and returns the encrypted private key. The private key password is a string that is used to encrypt the private key. After you obtain the encrypted private key of the certificate, you can use the following methods to decrypt the private key:
//
//   - If the encryption algorithm of the certificate is RSA, you must run the `openssl rsa -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
//   - If the encryption algorithm of the certificate is ECC, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
//   - If the encryption algorithm of the certificate is SM2, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
//
// >  You can call the [DescribeClientCertificate] operation to query the encryption algorithm type of a client certificate or a server certificate.
//
// ## Limits
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCertificatePrivateKeyRequest
//
// @return DescribeCertificatePrivateKeyResponse
func (client *Client) DescribeCertificatePrivateKey(request *DescribeCertificatePrivateKeyRequest) (_result *DescribeCertificatePrivateKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCertificatePrivateKeyResponse{}
	_body, _err := client.DescribeCertificatePrivateKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a client certificate or a server certificate by using the unique identifier of the certificate.
//
// Description:
//
// You can call the DescribeClientCertificate operation to query the details about a client certificate or a server certificate by using the unique identifier of the certificate. The details include the serial number, user information, content, and status of each certificate.
//
// Before you call this operation, make sure that you have created a client certificate or a server certificate.
//
// For more information about how to call an operation to create a client certificate, see the following topics:
//
//   - [CreateClientCertificate](https://help.aliyun.com/document_detail/330873.html)
//
//   - [CreateClientCertificateWithCsr](https://help.aliyun.com/document_detail/330875.html)
//
// For more information about how to call an operation to create a server certificate, see the following topics:
//
//   - [CreateServerCertificate](https://help.aliyun.com/document_detail/330877.html)
//
//   - [CreateServerCertificateWithCsr](https://help.aliyun.com/document_detail/330878.html)
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientCertificateResponse
func (client *Client) DescribeClientCertificateWithOptions(request *DescribeClientCertificateRequest, runtime *util.RuntimeOptions) (_result *DescribeClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClientCertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a client certificate or a server certificate by using the unique identifier of the certificate.
//
// Description:
//
// You can call the DescribeClientCertificate operation to query the details about a client certificate or a server certificate by using the unique identifier of the certificate. The details include the serial number, user information, content, and status of each certificate.
//
// Before you call this operation, make sure that you have created a client certificate or a server certificate.
//
// For more information about how to call an operation to create a client certificate, see the following topics:
//
//   - [CreateClientCertificate](https://help.aliyun.com/document_detail/330873.html)
//
//   - [CreateClientCertificateWithCsr](https://help.aliyun.com/document_detail/330875.html)
//
// For more information about how to call an operation to create a server certificate, see the following topics:
//
//   - [CreateServerCertificate](https://help.aliyun.com/document_detail/330877.html)
//
//   - [CreateServerCertificateWithCsr](https://help.aliyun.com/document_detail/330878.html)
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeClientCertificateRequest
//
// @return DescribeClientCertificateResponse
func (client *Client) DescribeClientCertificate(request *DescribeClientCertificateRequest) (_result *DescribeClientCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClientCertificateResponse{}
	_body, _err := client.DescribeClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status information about client certificates and server certificates by using the unique identifiers of the certificates.
//
// Description:
//
// You can call the DescribeClientCertificateStatus operation to query the status information about multiple client certificates or server certificates at a time by using the unique identifiers of the certificates. For example, you can check whether a certificate is revoked.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeClientCertificateStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientCertificateStatusResponse
func (client *Client) DescribeClientCertificateStatusWithOptions(request *DescribeClientCertificateStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeClientCertificateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClientCertificateStatus"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClientCertificateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status information about client certificates and server certificates by using the unique identifiers of the certificates.
//
// Description:
//
// You can call the DescribeClientCertificateStatus operation to query the status information about multiple client certificates or server certificates at a time by using the unique identifiers of the certificates. For example, you can check whether a certificate is revoked.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeClientCertificateStatusRequest
//
// @return DescribeClientCertificateStatusResponse
func (client *Client) DescribeClientCertificateStatus(request *DescribeClientCertificateStatusRequest) (_result *DescribeClientCertificateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClientCertificateStatusResponse{}
	_body, _err := client.DescribeClientCertificateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status information about a private root certificate authority (CA) instance or a private intermediate CA instance that you purchase by using the Certificate Management Service console.
//
// Description:
//
// ## Usage notes
//
// You can call the GetCAInstanceStatus operation to query the status information of a private CA instance by using the ID of the instance. The instance is purchased by using the SSL Certificates Service console. The status information includes the status of the private CA instance, the number of certificates that can be issued by using the private CA instance, and the number of issued certificates.
//
// Before you call this operation, make sure that you have purchased a private CA by using the [SSL Certificates Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// @param request - GetCAInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCAInstanceStatusResponse
func (client *Client) GetCAInstanceStatusWithOptions(request *GetCAInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *GetCAInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCAInstanceStatus"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCAInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status information about a private root certificate authority (CA) instance or a private intermediate CA instance that you purchase by using the Certificate Management Service console.
//
// Description:
//
// ## Usage notes
//
// You can call the GetCAInstanceStatus operation to query the status information of a private CA instance by using the ID of the instance. The instance is purchased by using the SSL Certificates Service console. The status information includes the status of the private CA instance, the number of certificates that can be issued by using the private CA instance, and the number of issued certificates.
//
// Before you call this operation, make sure that you have purchased a private CA by using the [SSL Certificates Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](https://help.aliyun.com/document_detail/208553.html).
//
// @param request - GetCAInstanceStatusRequest
//
// @return GetCAInstanceStatusResponse
func (client *Client) GetCAInstanceStatus(request *GetCAInstanceStatusRequest) (_result *GetCAInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCAInstanceStatusResponse{}
	_body, _err := client.GetCAInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取证书列表
//
// @param request - ListCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCertResponse
func (client *Client) ListCertWithOptions(request *ListCertRequest, runtime *util.RuntimeOptions) (_result *ListCertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AfterDate)) {
		query["AfterDate"] = request.AfterDate
	}

	if !tea.BoolValue(util.IsUnset(request.BeforeDate)) {
		query["BeforeDate"] = request.BeforeDate
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceUuid)) {
		query["InstanceUuid"] = request.InstanceUuid
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCert"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取证书列表
//
// @param request - ListCertRequest
//
// @return ListCertResponse
func (client *Client) ListCert(request *ListCertRequest) (_result *ListCertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCertResponse{}
	_body, _err := client.ListCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about all client certificates and server certificates.
//
// Description:
//
// You can call the ListClientCertificate operation to perform a paged query of the details about all client certificates and server certificates that you create. The details include the unique identifier, serial number, user information, content, and status of each certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientCertificateResponse
func (client *Client) ListClientCertificateWithOptions(request *ListClientCertificateRequest, runtime *util.RuntimeOptions) (_result *ListClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClientCertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about all client certificates and server certificates.
//
// Description:
//
// You can call the ListClientCertificate operation to perform a paged query of the details about all client certificates and server certificates that you create. The details include the unique identifier, serial number, user information, content, and status of each certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListClientCertificateRequest
//
// @return ListClientCertificateResponse
func (client *Client) ListClientCertificate(request *ListClientCertificateRequest) (_result *ListClientCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClientCertificateResponse{}
	_body, _err := client.ListClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about all client certificates and server certificates that are revoked.
//
// Description:
//
// You can call the ListRevokeCertificate operation to perform a paged query of the details about all revoked client certificates and server certificates. The details include the unique identifier, serial number, and revocation date of each certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListRevokeCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRevokeCertificateResponse
func (client *Client) ListRevokeCertificateWithOptions(request *ListRevokeCertificateRequest, runtime *util.RuntimeOptions) (_result *ListRevokeCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRevokeCertificate"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRevokeCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about all client certificates and server certificates that are revoked.
//
// Description:
//
// You can call the ListRevokeCertificate operation to perform a paged query of the details about all revoked client certificates and server certificates. The details include the unique identifier, serial number, and revocation date of each certificate.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListRevokeCertificateRequest
//
// @return ListRevokeCertificateResponse
func (client *Client) ListRevokeCertificate(request *ListRevokeCertificateRequest) (_result *ListRevokeCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRevokeCertificateResponse{}
	_body, _err := client.ListRevokeCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the status of a root certificate authority (CA) certificate or an intermediate CA certificate from ISSUE to REVOKE.
//
// Description:
//
// After a CA certificate is created, the CA certificate is in the ISSUE state by default. You can call the UpdateCACertificateStatus operation to change the status of a CA certificate from ISSUE to REVOKE. If a CA certificate is in the ISSUE state, the CA certificate can be used to issue certificates. If a CA certificate is in the REVOKE state, the CA certificate cannot be used to issue certificates, and the certificates that are issued from the CA certificate become invalid.
//
// Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate] operation or an intermediate CA certificate by calling the [CreateSubCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateCACertificateStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCACertificateStatusResponse
func (client *Client) UpdateCACertificateStatusWithOptions(request *UpdateCACertificateStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateCACertificateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCACertificateStatus"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCACertificateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of a root certificate authority (CA) certificate or an intermediate CA certificate from ISSUE to REVOKE.
//
// Description:
//
// After a CA certificate is created, the CA certificate is in the ISSUE state by default. You can call the UpdateCACertificateStatus operation to change the status of a CA certificate from ISSUE to REVOKE. If a CA certificate is in the ISSUE state, the CA certificate can be used to issue certificates. If a CA certificate is in the REVOKE state, the CA certificate cannot be used to issue certificates, and the certificates that are issued from the CA certificate become invalid.
//
// Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate] operation or an intermediate CA certificate by calling the [CreateSubCACertificate] operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateCACertificateStatusRequest
//
// @return UpdateCACertificateStatusResponse
func (client *Client) UpdateCACertificateStatus(request *UpdateCACertificateStatusRequest) (_result *UpdateCACertificateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCACertificateStatusResponse{}
	_body, _err := client.UpdateCACertificateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传pca证书到SSL上传证书
//
// @param request - UploadPcaCertToCasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadPcaCertToCasResponse
func (client *Client) UploadPcaCertToCasWithOptions(request *UploadPcaCertToCasRequest, runtime *util.RuntimeOptions) (_result *UploadPcaCertToCasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadPcaCertToCas"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadPcaCertToCasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传pca证书到SSL上传证书
//
// @param request - UploadPcaCertToCasRequest
//
// @return UploadPcaCertToCasResponse
func (client *Client) UploadPcaCertToCas(request *UploadPcaCertToCasRequest) (_result *UploadPcaCertToCasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadPcaCertToCasResponse{}
	_body, _err := client.UploadPcaCertToCasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
