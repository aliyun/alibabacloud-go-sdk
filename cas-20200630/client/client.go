// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
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
	// >  The **BeforeTime** and **AfterTime** parameters must be both empty or both specified.
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the client certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// *   **RSA\_1024**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_2048**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_4096**: The signature algorithm is Sha256WithRSA.
	// *   **ECC\_256**: The signature algorithm is Sha256WithECDSA.
	// *   **ECC\_384**: The signature algorithm is Sha256WithECDSA.
	// *   **ECC\_512**: The signature algorithm is Sha256WithECDSA.
	// *   **SM2\_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the client certificate must be the same with the encryption algorithm of the intermediate CA certificate. The key length can be different. For example, if the key algorithm of the intermediate CA certificate is RSA\_2048, the key algorithm of the client certificate must be RSA\_1024, RSA\_2048, or RSA\_4096.
	//
	// >  You can call the [DescribeCACertificate](~~328096~~) operation to query the key algorithm of an intermediate CA certificate.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the client certificate. This value is a UNIX timestamp. The default value is the time when you call this operation. Unit: seconds.
	//
	// >  The **BeforeTime** and **AfterTime** parameters must be both empty or both specified.
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The name of the client certificate user. In most cases, the user of a client certificate is an individual, a company, an organization, or an application. We recommend that you enter the common name of a user. Examples: Bob, Alibaba, Alibaba Cloud password platform, and Tmall Genie.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country in which the organization is located. Default value: CN.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the CSR file. You can generate a CSR file by using the OpenSSL tool or Keytool. For more information, see [How do I create a CSR file?](~~42218~~) You can also create a CSR file in the Certificate Management Service console. For more information, see [Create a CSR](~~313297~~).
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The validity period of the client certificate. Unit: days. You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime** parameters. The **BeforeTime** and **AfterTime** parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// *   If you specify the **Days** parameter, you can specify both the **BeforeTime** and **AfterTime** parameters or leave them both empty.********
	// *   If you do not specify the **Days** parameter, you must specify both the **BeforeTime** and **AfterTime** parameters.
	//
	// >
	//
	// *   If you specify the **Days**, **BeforeTime**, and **AfterTime** parameters together, the validity period of the client certificate is determined by the value of the **Days** parameter.
	//
	// *   The validity period of the client certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](~~328096~~) operation to query the validity period of an intermediate CA certificate.
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// Specifies whether to return the certificate. Valid values:
	//
	// *   **0**: does not return the certificate. This is the default value.
	// *   **1**: returns the certificate.
	// *   **2**: returns the certificate and the certificate chain of the certificate.
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The name of the city in which the organization is located. The value can contain letters. The default value is the name of the city in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the client certificate. Unit: months.
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate CA certificate from which the client certificate is issued.
	//
	// >  You can call the [DescribeCACertificateList](~~328095~~) operation to query the unique identifier of an intermediate CA certificate.
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The type of the Subject Alternative Name (SAN) extension that is supported by the client certificate. Valid values:
	//
	// *   **1**: an email address
	// *   **6**: a Uniform Resource Identifier (URI)
	SanType *int32 `json:"SanType,omitempty" xml:"SanType,omitempty"`
	// The content of the extension. You can specify multiple SAN extensions. If you want to specify multiple SAN extensions, separate them with commas (,).
	SanValue *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the client certificate. Unit: years.
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

func (s *CreateClientCertificateRequest) SetCsr(v string) *CreateClientCertificateRequest {
	s.Csr = &v
	return s
}

func (s *CreateClientCertificateRequest) SetDays(v int32) *CreateClientCertificateRequest {
	s.Days = &v
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
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the client certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the server certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the client certificate.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// >  The **BeforeTime** and **AfterTime** parameters must be both empty or both specified.
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the client certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// *   **RSA\_1024**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_2048**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_4096**: The signature algorithm is Sha256WithRSA.
	// *   **ECC\_256**: The signature algorithm is Sha256WithECDSA.
	// *   **ECC\_384**: The signature algorithm is Sha256WithECDSA.
	// *   **ECC\_512**: The signature algorithm is Sha256WithECDSA.
	// *   **SM2\_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the client certificate must be the same with the encryption algorithm of the intermediate CA certificate. The key length can be different. For example, if the key algorithm of the intermediate CA certificate is RSA\_2048, the key algorithm of the client certificate must be RSA\_1024, RSA\_2048, or RSA\_4096.
	//
	// >  You can call the [DescribeCACertificate](~~328096~~) operation to query the key algorithm of an intermediate CA certificate.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the client certificate. This value is a UNIX timestamp. The default value is the time when you call this operation. Unit: seconds.
	//
	// >  The **BeforeTime** and **AfterTime** parameters must be both empty or both specified.
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The common name of the certificate. The value can contain letters.
	//
	// >  If you specify the **CsrPemString** parameter, the value of the **CommonName** parameter is determined by the **CsrPemString** parameter.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located, such as **CN** and **US**.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the CSR file. You can generate a CSR file by using the OpenSSL tool or Keytool. For more information, see [How do I create a CSR file?](~~42218~~) You can also create a CSR file in the Certificate Management Service console. For more information, see [Create a CSR](~~313297~~).
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The content of the CSR file. You can generate a CSR file by using the OpenSSL tool or Keytool. For more information, see [How do I create a CSR file?](~~42218~~) You can also create a CSR file in the Certificate Management Service console. For more information, see [Create a CSR](~~313297~~).
	Csr1 *string `json:"Csr1,omitempty" xml:"Csr1,omitempty"`
	// The validity period of the client certificate. Unit: days. You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime** parameters. The **BeforeTime** and **AfterTime** parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// *   If you specify the **Days** parameter, you can specify both the **BeforeTime** and **AfterTime** parameters or leave them both empty.********
	// *   If you do not specify the **Days** parameter, you must specify both the **BeforeTime** and **AfterTime** parameters.
	//
	// >
	//
	// *   If you specify the **Days**, **BeforeTime**, and **AfterTime** parameters together, the validity period of the client certificate is determined by the value of the **Days** parameter.
	//
	// *   The validity period of the client certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](~~328096~~) operation to query the validity period of an intermediate CA certificate.
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// Specifies whether to return the certificate. Valid values:
	//
	// *   **0**: does not return the certificate. This is the default value.
	// *   **1**: returns the certificate.
	// *   **2**: returns the certificate and the certificate chain of the certificate.
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The name of the city in which the organization is located. The value can contain letters. The default value is the name of the city in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the client certificate. Unit: months.
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate CA certificate from which the client certificate is issued.
	//
	// >  You can call the [DescribeCACertificateList](~~328095~~) operation to query the unique identifier of an intermediate CA certificate.
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The type of the Subject Alternative Name (SAN) extension that is supported by the client certificate. Valid values:
	//
	// *   **1**: an email address
	// *   **6**: a Uniform Resource Identifier (URI)
	SanType *int32 `json:"SanType,omitempty" xml:"SanType,omitempty"`
	// The content of the extension. You can specify multiple SAN extensions. If you want to specify multiple SAN extensions, separate them with commas (,).
	SanValue *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the client certificate. Unit: years.
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

func (s *CreateClientCertificateWithCsrRequest) SetCsr1(v string) *CreateClientCertificateWithCsrRequest {
	s.Csr1 = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetDays(v int32) *CreateClientCertificateWithCsrRequest {
	s.Days = &v
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
	// The certificate chain of the client certificate.
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the client certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the server certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the client certificate.
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s CreateClientCertificateWithCsrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClientCertificateWithCsrResponseBody) GoString() string {
	return s.String()
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateClientCertificateWithCsrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApiPassthrough   *CreateCustomCertificateRequestApiPassthrough `json:"ApiPassthrough,omitempty" xml:"ApiPassthrough,omitempty" type:"Struct"`
	Csr              *string                                       `json:"Csr,omitempty" xml:"Csr,omitempty"`
	Immediately      *int32                                        `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	ParentIdentifier *string                                       `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	Validity         *string                                       `json:"Validity,omitempty" xml:"Validity,omitempty"`
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
	Extensions *CreateCustomCertificateRequestApiPassthroughExtensions `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Struct"`
	Subject    *CreateCustomCertificateRequestApiPassthroughSubject    `json:"Subject,omitempty" xml:"Subject,omitempty" type:"Struct"`
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

func (s *CreateCustomCertificateRequestApiPassthrough) SetSubject(v *CreateCustomCertificateRequestApiPassthroughSubject) *CreateCustomCertificateRequestApiPassthrough {
	s.Subject = v
	return s
}

type CreateCustomCertificateRequestApiPassthroughExtensions struct {
	ExtendedKeyUsages       []*string                                                                        `json:"ExtendedKeyUsages,omitempty" xml:"ExtendedKeyUsages,omitempty" type:"Repeated"`
	KeyUsage                *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage                  `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty" type:"Struct"`
	SubjectAlternativeNames []*CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
}

func (s CreateCustomCertificateRequestApiPassthroughExtensions) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughExtensions) GoString() string {
	return s.String()
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
	ContentCommitment *bool `json:"ContentCommitment,omitempty" xml:"ContentCommitment,omitempty"`
	DataEncipherment  *bool `json:"DataEncipherment,omitempty" xml:"DataEncipherment,omitempty"`
	DecipherOnly      *bool `json:"DecipherOnly,omitempty" xml:"DecipherOnly,omitempty"`
	DigitalSignature  *bool `json:"DigitalSignature,omitempty" xml:"DigitalSignature,omitempty"`
	EncipherOnly      *bool `json:"EncipherOnly,omitempty" xml:"EncipherOnly,omitempty"`
	KeyAgreement      *bool `json:"KeyAgreement,omitempty" xml:"KeyAgreement,omitempty"`
	KeyEncipherment   *bool `json:"KeyEncipherment,omitempty" xml:"KeyEncipherment,omitempty"`
	NonRepudiation    *bool `json:"NonRepudiation,omitempty" xml:"NonRepudiation,omitempty"`
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
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	Country          *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
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

type CreateCustomCertificateResponseBody struct {
	Certificate      *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	Identifier       *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// >  You can call the [ListClientCertificate](~~330884~~) operation to query the unique identifiers of all client certificates and server certificates.
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRevokeClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   **RSA\_1024**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_2048**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_4096**: The signature algorithm is Sha256WithRSA.
	// *   **ECC\_256**: The signature algorithm is Sha256WithECDSA.
	// *   **ECC\_384**: The signature algorithm is Sha256WithECDSA.
	// *   **ECC\_512**: The signature algorithm is Sha256WithECDSA.
	// *   **SM2\_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the root CA certificate must be consistent with the **encryption algorithm** of the private root CA instance that you purchase. For example, if the **encryption algorithm** of the private root CA instance that you purchase is **RSA**, the key algorithm of the root CA certificate must be **RSA\_1024**, **RSA\_2048**, or **RSA\_4096**.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The common name or abbreviation of the organization. The value can contain letters.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country or region in which the organization is located. You can enter an alpha-2 code. For example, you can use **CN** to indicate China and use **US** to indicate the United States.
	//
	// For more information about country codes, see the **"Country codes"** section of the [Manage company profiles](~~198289~~) topic.
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The name of the city in which the organization is located. The value can contain letters.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the organization that is associated with the root CA certificate. You can enter the name of your enterprise or company. The value can contain letters.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization. The value can contain letters.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The value can contain letters.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the root CA certificate. Unit: years.
	//
	// >  We recommend that you set this parameter to a value from 5 to 10.
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
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain of the root CA certificate.
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the root CA certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRootCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// >  The **BeforeTime** and **AfterTime** parameters must be both empty or both specified.
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the server certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// *   **RSA\_1024**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_2048**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_4096**: The signature algorithm is Sha256WithRSA.
	// *   **ECC\_256**: The signature algorithm is Sha256WithECDSA.
	// *   **ECC\_384**: The signature algorithm is Sha256WithECDSA.
	// *   **ECC\_512**: The signature algorithm is Sha256WithECDSA.
	// *   **SM2\_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the server certificate must be the same as the encryption algorithm of the intermediate CA certificate. The key length can be different. For example, if the key algorithm of the intermediate CA certificate is RSA\_2048, the key algorithm of the server certificate must be RSA\_1024, RSA\_2048, or RSA\_4096.
	//
	// >  You can call the [DescribeCACertificate](~~328096~~) operation to query the key algorithm of an intermediate CA certificate.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the server certificate. This value is a UNIX timestamp. The default value is the time when you call this operation. Unit: seconds.
	//
	// >  The **BeforeTime** and **AfterTime** parameters must be both empty or both specified.
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The name of the certificate user. The user of a server certificate is a server. We recommend that you enter the domain name or IP address of the server.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located, such as CN or US.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the CSR file. You can generate a CSR file by using the OpenSSL tool or Keytool. For more information, see [How do I create a CSR file?](~~42218~~) You can also create a CSR file in the Certificate Management Service console. For more information, see [Create a CSR](~~313297~~).
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The validity period of the server certificate. Unit: days. You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime** parameters. The **BeforeTime** and **AfterTime** parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// *   If you specify the **Days** parameter, you can specify both the **BeforeTime** and **AfterTime** parameters or leave them both empty.********
	// *   If you do not specify the **Days** parameter, you must specify both the **BeforeTime** and **AfterTime** parameters.
	//
	// >
	//
	// *   If you specify the **Days**, **BeforeTime**, and **AfterTime** parameters together, the validity period of the server certificate is determined by the value of the **Days** parameter.
	//
	// *   The validity period of the server certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](~~328096~~) operation to query the validity period of an intermediate CA certificate.
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The additional domain names and additional IP addresses of the server certificate. After you add additional domain names and additional IP addresses to a certificate, you can apply the certificate to the domain names and IP addresses.
	//
	// Separate multiple domain names and multiple IP addresses with commas (,).
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to return the certificate. Valid values:
	//
	// *   **0**: does not return the certificate. This is the default value.
	// *   **1**: returns the certificate.
	// *   **2**: returns the certificate and the certificate chain of the certificate.
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The name of the city in which the organization is located. The value can contain letters. The default value is the name of the city in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the server certificate. Unit: months.
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate CA certificate from which the server certificate is issued.
	//
	// >  You can call the [DescribeCACertificateList](~~328095~~) operation to query the unique identifier of an intermediate CA certificate.
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the server certificate. Unit: years.
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

func (s *CreateServerCertificateRequest) SetCsr(v string) *CreateServerCertificateRequest {
	s.Csr = &v
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
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the server certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the server certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the server certificate.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServerCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// >  The **BeforeTime** and **AfterTime** parameters must be both empty or both specified.
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the server certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// *   **RSA\_1024**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_2048**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_4096**: The signature algorithm is Sha256WithRSA.
	// *   **ECC\_256**: The signature algorithm is Sha256WithECDSA.
	// *   **ECC\_384**: The signature algorithm is Sha256WithECDSA.
	// *   **ECC\_512**: The signature algorithm is Sha256WithECDSA.
	// *   **SM2\_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the server certificate must be the same as the encryption algorithm of the intermediate CA certificate. The key length can be different. For example, if the key algorithm of the intermediate CA certificate is RSA\_2048, the key algorithm of the server certificate must be RSA\_1024, RSA\_2048, or RSA\_4096.
	//
	// >  You can call the [DescribeCACertificate](~~328096~~) operation to query the key algorithm of an intermediate CA certificate.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the server certificate. This value is a UNIX timestamp. The default value is the time when you call this operation. Unit: seconds.
	//
	// >  The **BeforeTime** and **AfterTime** parameters must be both empty or both specified.
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The common name of the certificate. The value can contain letters.
	//
	// >  If you specify the **CsrPemString** parameter, the value of the **CommonName** parameter is determined by the **CsrPemString** parameter.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located, such as **CN**.
	//
	// >  This parameter is available and required only when the **RegistrantProfileId** parameter is not specified. In this case, you must specify this parameter. If this parameter is not specified, the domain name fails to be registered.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the CSR file. You can generate a CSR file by using the OpenSSL tool or Keytool. For more information, see [How do I create a CSR file?](~~42218~~) You can also create a CSR file in the Certificate Management Service console. For more information, see [Create a CSR](~~313297~~).
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The content of the CSR file. You can generate a CSR file by using the OpenSSL tool or Keytool. For more information, see [How do I create a CSR file?](~~42218~~) You can also create a CSR file in the Certificate Management Service console. For more information, see [Create a CSR](~~313297~~).
	Csr1 *string `json:"Csr1,omitempty" xml:"Csr1,omitempty"`
	// The validity period of the server certificate. Unit: days. You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime** parameters. The **BeforeTime** and **AfterTime** parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// *   If you specify the **Days** parameter, you can specify both the **BeforeTime** and **AfterTime** parameters or leave them both empty.********
	// *   If you do not specify the **Days** parameter, you must specify both the **BeforeTime** and **AfterTime** parameters.
	//
	// >
	//
	// *   If you specify the **Days**, **BeforeTime**, and **AfterTime** parameters together, the validity period of the server certificate is determined by the value of the **Days** parameter.
	//
	// *   The validity period of the server certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](~~328096~~) operation to query the validity period of an intermediate CA certificate.
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The additional domain names or additional IP addresses of the server certificate. After you add additional domain names and additional IP addresses to a certificate, you can apply the certificate to the domain names and IP addresses.
	//
	// You can specify multiple domain names and IP addresses. If you specify multiple domain names and IP addresses, separate them with commas (,).
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to return the certificate. Valid values:
	//
	// *   **0**: does not return the certificate. This is the default value.
	// *   **1**: returns the certificate.
	// *   **2**: returns the certificate and the certificate chain of the certificate.
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The name of the city in which the organization is located. The value can contain letters. The default value is the name of the city in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the server certificate. Unit: months.
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate CA certificate from which the server certificate is issued.
	//
	// >  You can call the [DescribeCACertificateList](~~328095~~) operation to query the unique identifier of an intermediate CA certificate.
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the server certificate. Unit: years.
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

func (s *CreateServerCertificateWithCsrRequest) SetCsr1(v string) *CreateServerCertificateWithCsrRequest {
	s.Csr1 = &v
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
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the server certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the server certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the server certificate.
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServerCertificateWithCsrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   **RSA\_1024**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_2048**: The signature algorithm is Sha256WithRSA.
	// *   **RSA\_4096**: The signature algorithm is Sha256WithRSA.
	// *   **ECC\_256**: The signature algorithm is Sha256WithECDSA.
	// *   **SM2\_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of an intermediate CA certificate must be consistent with the encryption algorithm of a root CA certificate. The length of the keys can be different. For example, if the key algorithm of the root CA certificate is **RSA\_2048**, the key algorithm of the intermediate CA certificate must be **RSA\_1024**, **RSA\_2048**, or **RSA\_4096**.
	//
	// >  You can call the [DescribeCACertificate](~~328096~~) operation to query the key algorithm of a root CA certificate.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The common name or abbreviation of the organization. The value can contain letters.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country or region in which the organization is located. You can enter an alpha-2 or alpha-3 code. For example, you can use **CN** to indicate China and use **US** to indicate the United States.
	//
	// For more information about country codes, see the **"Country codes"** section of the [Manage company profiles](~~198289~~) topic.
	CountryCode       *string   `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	ExtendedKeyUsages []*string `json:"ExtendedKeyUsages,omitempty" xml:"ExtendedKeyUsages,omitempty" type:"Repeated"`
	// The name of the city in which the organization is located. The value can contain letters.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the organization that is associated with the intermediate CA certificate. You can enter the name of your enterprise or company. The value can contain letters.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization. The value can contain letters.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the root CA certificate.
	//
	// >  You can call the [DescribeCACertificateList](~~328095~~) operation to query the unique identifiers of all CA certificates.
	ParentIdentifier  *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	PathLenConstraint *int32  `json:"PathLenConstraint,omitempty" xml:"PathLenConstraint,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The value can contain letters.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the intermediate CA certificate. Unit: years.
	//
	// We recommend that you set this parameter to 5 to 10.
	//
	// >  The validity period of the intermediate CA certificate cannot exceed the validity period of the root CA certificate. You can call the [DescribeCACertificate](~~328095~~) operation to query the validity period of a root CA certificate.
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
	// The intermediate CA certificate in the PEM format.
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain of the intermediate CA certificate.
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the intermediate CA certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSubCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// >  You can call the [ListClientCertificate](~~330884~~) operation to query the unique identifiers and status of all client certificates and server certificates.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// >  You can call the [DescribeCACertificateList](~~328095~~) operation to query the unique identifiers of all CA certificates.
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The validity period of the CA certificate. Unit: years.
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
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The encryption algorithm of the CA certificate. Valid values:
	//
	// *   **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	// *   **ECC**: the elliptic curve cryptography (ECC) algorithm.
	// *   **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance date of the CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the CA certificate. Valid values:
	//
	// *   **ROOT**: root CA certificate
	// *   **SUB_ROOT**: intermediate CA certificate
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name or abbreviation of the organization that is associated with the CA certificate.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located.
	//
	// For more information about country codes, see the **"Country codes"** section of the [Manage company profiles](~~198289~~) topic.
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The status of the certificate revocation list (CRL) feature.
	CrlStatus *string `json:"CrlStatus,omitempty" xml:"CrlStatus,omitempty"`
	// The address of the CRL.
	CrlUrl *string `json:"CrlUrl,omitempty" xml:"CrlUrl,omitempty"`
	// The unique identifier of the CA certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the CA certificate.
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the CA certificate.
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization that is associated with the CA certificate.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization that is associated with the CA certificate.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the root CA certificate from which the CA certificate is issued.
	//
	// >  This parameter is returned only if the value of the **CertificateType** parameter is **SUB_ROOT**. The value SUB_ROOT indicates an intermediate CA certificate.
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// This parameter is deprecated.
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the CA certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the CA certificate.
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the CA certificate.
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the CA certificate. Valid values:
	//
	// *   **ISSUE**: The CA certificate is issued.
	// *   **REVOKE**: The CA certificate is revoked.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user attribute of the CA certificate, which contains the following information:
	//
	// *   **C**: the country code in which the organization is located
	// *   **O**: the name of the organization
	// *   **OU**: the name of the department or branch in the organization
	// *   **L**: the name of the city in which the organization is located
	// *   **ST**: the name of the province, municipality, or autonomous region in which the organization is located
	// *   **CN**: the common name or abbreviation of the organization
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// The content of the CA certificate.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of created CA certificates, which includes root CA certificates and intermediate CA certificates.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCACertificateCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The number of the page to return. Default value: **1**.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of CA certificates to return on each page. Default value: **20**.
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s DescribeCACertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateListRequest) SetCurrentPage(v int32) *DescribeCACertificateListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetShowSize(v int32) *DescribeCACertificateListRequest {
	s.ShowSize = &v
	return s
}

type DescribeCACertificateListResponseBody struct {
	// An array that consists of the details about the CA certificate.
	CertificateList []*DescribeCACertificateListResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of returned pages.
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of CA certificates returned per page.
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of root CA certificates and intermediate CA certificates that are returned.
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
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The encryption algorithm of the CA certificate. Valid values:
	//
	// *   **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	// *   **ECC**: the elliptic curve cryptography (ECC) algorithm.
	// *   **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance date of the CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the CA certificate. Valid values:
	//
	// *   **ROOT**: root CA certificate
	// *   **SUB_ROOT**: intermediate CA certificate
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name or abbreviation of the organization that is associated with the CA certificate.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located.
	//
	// For more information about country codes, see the **"Country codes"** section of the [Manage company profiles](~~198289~~) topic.
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The unique identifier of the CA certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the CA certificate.
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the CA certificate.
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization that is associated with the CA certificate.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization that is associated with the CA certificate.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the root CA certificate from which the CA certificate is issued.
	//
	// >  This parameter is returned only if the value of the **CertificateType** parameter is **SUB_ROOT**. The value SUB_ROOT indicates an intermediate CA certificate.
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// This parameter is deprecated.
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the CA certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the CA certificate.
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the CA certificate.
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the CA certificate. Valid values:
	//
	// *   **ISSUE**: The CA certificate is issued.
	// *   **REVOKE**: The CA certificate is revoked.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Distinguished Name (DN) attribute of the CA certificate, which indicates the user information of the certificate. The DN attribute contains the following information:
	//
	// *   **C**: the country code in which the organization is located
	// *   **O**: the name of the organization
	// *   **OU**: the name of the department or branch in the organization
	// *   **L**: the name of the city in which the organization is located
	//
	// <props="china">- **ST**: the name of the province, municipality, or autonomous region in which the organization is located</props> <props="intl">- **ST**: the name of the province or state in which the organization is located</props>
	//
	// *   **CN**: the common name or abbreviation of the organization
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// The content of the CA certificate.
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
	// The validity period of the CA certificate. Unit: years.
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

func (s *DescribeCACertificateListResponseBodyCertificateList) SetX509Certificate(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.X509Certificate = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetYears(v int32) *DescribeCACertificateListResponseBodyCertificateList {
	s.Years = &v
	return s
}

type DescribeCACertificateListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCACertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// **
	//
	// **Warning** You must remember the password that you specify. The password is required to decrypt the encrypted private key. If you forget the password, the encrypted private key that is returned cannot be decrypted. You must call this operation again.
	EncryptedCode *string `json:"EncryptedCode,omitempty" xml:"EncryptedCode,omitempty"`
	// The unique identifier of the client certificate or server certificate that you want to query.
	//
	// >  You can call the [ListClientCertificate](~~330884~~) operation to query the unique identifiers of all client certificates and server certificates.
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
	EncryptedData *string `json:"EncryptedData,omitempty" xml:"EncryptedData,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCertificatePrivateKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// >  You can call the [ListClientCertificate](~~330884~~) operation to query the unique identifiers of all client certificates and server certificates.
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
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The type of the encryption algorithm of the certificate. Valid values:
	//
	// *   **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	// *   **ECC**: the elliptic curve cryptography (ECC) algorithm.
	// *   **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance date of the certificate. This value is a UNIX timestamp. Unit: milliseconds.
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the certificate. Valid values:
	//
	// *   **CLIENT**: client certificate
	// *   **SERVER**: server certificate
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name of the certificate.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// For more information about country codes, see the **"Country codes"** section of the [Manage company profiles](~~198289~~) topic.
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The validity period of the certificate. Unit: days.
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The unique identifier of the certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the certificate.
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the certificate.
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization. The organization is associated with the intermediate certificate from which the certificate is issued.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department in the organization. The organization is associated with the intermediate certificate authority (CA) certificate from which the certificate is issued.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate certificate from which the client certificate is issued.
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The subject alternative name (SAN) extension of the certificate. The value indicates additional information, including the additional domain names or IP addresses that are associated with the certificate.
	//
	// The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that corresponds to a SAN extension. A SAN extension struct contains the following parameters:
	//
	// *   **Type**: the type of the extension. Data type: integer. Valid values:
	//
	//     *   **1**: an email address
	//     *   **2**: a domain name
	//     *   **6**: a Uniform Resource Identifier (URI)
	//     *   **7**: an IP address
	//
	// *   **Value**: the value of the extension. Data type: string.
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the certificate.
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the certificate. Valid values:
	//
	// *   **ISSUE**: issued
	// *   **REVOKE**: revoked
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The distinguished name (DN) extension of the certificate, which indicates the user of the certificate. The DN extension includes the following information:
	//
	// *   **C**: the country
	// *   **O**: the organization
	// *   **OU**: the department
	// *   **L**: the city
	// *   **ST**: the province, municipality, or autonomous region
	// *   **CN**: the common name
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// The content of the certificate.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// >  You can call the [ListClientCertificate](~~330884~~) operation to query the unique identifiers of all client certificates and server certificates.
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
	// >  This parameter is returned only when the value of the **Status** parameter is **revoked**. The value revoked indicates that the certificate is revoked.
	RevokeTime *int64 `json:"RevokeTime,omitempty" xml:"RevokeTime,omitempty"`
	// The serial number of the certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The status of the certificate. Valid values:
	//
	// *   **good**: The certificate is not revoked.
	// *   **revoked**: The certificate is revoked.
	// *   **unknown**: The server cannot determine the status of the certificate.
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeClientCertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The ID of the private CA instance.
	//
	// >  After you purchase a private CA instance by using the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist), you can click **Details** for the private CA instance on the **Private Certificates** page to obtain the ID of the private CA instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCAInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCAInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCAInstanceStatusRequest) SetInstanceId(v string) *GetCAInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

type GetCAInstanceStatusResponseBody struct {
	// An array that consists of the status information about the private CA instance.
	InstanceStatusList []*GetCAInstanceStatusResponseBodyInstanceStatusList `json:"InstanceStatusList,omitempty" xml:"InstanceStatusList,omitempty" type:"Repeated"`
	// The ID of the request.
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
	// >  This parameter is returned only when the value of the **Status** parameter is **USED** or **REVOKE**. The value USED indicates that the private CA instance is enabled, and the value REVOKE indicates that the instance is revoked.
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The issuance date of the private CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// >  This parameter is returned only when the value of the **Status** parameter is **USED** or **REVOKE**. The value USED indicates that the private CA instance is enabled, and the value REVOKE indicates that the instance is revoked.
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The number of certificates that are issued by using the private CA instance.
	CertIssuedCount *int32 `json:"CertIssuedCount,omitempty" xml:"CertIssuedCount,omitempty"`
	// The number of certificates that can be issued by using the private CA instance.
	//
	// For a private root CA instance whose **Type** is **ROOT**, this parameter indicates the number of intermediate CA certificates that can be issued. For a private intermediate CA instance whose **Type** is **SUB_ROOT**, this parameter indicates the total number of client certificates and server certificates that can be issued
	CertTotalCount *int32 `json:"CertTotalCount,omitempty" xml:"CertTotalCount,omitempty"`
	// The unique identifier of the private CA certificate.
	//
	// >  This parameter is returned only when the value of the **Status** parameter is **USED** or **REVOKE**. The value USED indicates that the private CA instance is enabled, and the value REVOKE indicates that the instance is revoked.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the private CA instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the private CA instance. Valid values:
	//
	// *   **BUY**: The private CA instance is purchased but is not enabled.
	// *   **USED**: The private CA instance is enabled.
	// *   **REFUND**: The payment of the private CA instance is refunded.
	// *   **REVOKE**: The private CA instance is revoked.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the private CA instance. Valid values:
	//
	// *   **ROOT**: root CA instance
	// *   **SUB_ROOT**: intermediate CA instance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The expiration date of the private CA instance. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// >  This parameter corresponds to the duration that you select when you purchase the private CA instance. The duration indicates the subscription period of the Private Certificate Authority (PCA) service.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCAInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListClientCertificateRequest struct {
	// The number of the page to return. Default value: **1**.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of certificates to return on each page. Default value: **20**.
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

func (s *ListClientCertificateRequest) SetShowSize(v int32) *ListClientCertificateRequest {
	s.ShowSize = &v
	return s
}

type ListClientCertificateResponseBody struct {
	// An array that consists of the details about all client certificates and server certificates.
	CertificateList []*ListClientCertificateResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// The page number of the current page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of pages returned.
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of certificates that are returned per page.
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The number of client certificates and server certificates that are returned.
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
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The type of the encryption algorithm of the certificate. Valid values:
	//
	// *   **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	// *   **ECC**: the elliptic curve cryptography (ECC) algorithm.
	// *   **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance date of the certificate. This value is a UNIX timestamp. Unit: milliseconds.
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the certificate. Valid values:
	//
	// *   **CLIENT**: client certificate
	// *   **SERVER**: server certificate
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name of the certificate.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// For more information about country codes, see the **"Country codes"** section of the [Manage company profiles](~~198289~~) topic.
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The validity period of the certificate. Unit: days.
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The unique identifier of the certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the certificate.
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the certificate.
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization. The organization is associated with the intermediate certificate from which the certificate is issued.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department in the organization. The organization is associated with the intermediate certificate authority (CA) certificate from which the certificate is issued.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate certificate from which the client certificate is issued.
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The subject alternative name (SAN) extension of the certificate. The value indicates additional information, including the additional domain names or IP addresses that are associated with the certificate.
	//
	// The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that corresponds to a SAN extension. A SAN extension struct contains the following parameters:
	//
	// *   **Type**: the type of the extension. Data type: integer. Valid values:
	//
	//     *   **1**: an email address
	//     *   **2**: a domain name
	//     *   **6**: a Uniform Resource Identifier (URI)
	//     *   **7**: an IP address
	//
	// *   **Value**: the value of the extension. Data type: string.
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the certificate.
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the certificate. Valid values:
	//
	// *   **ISSUE**: issued
	// *   **REVOKE**: revoked
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The distinguished name (DN) extension of the certificate, which indicates the user of the certificate. The DN extension includes the following information:
	//
	// *   **C**: the country
	// *   **O**: the organization
	// *   **OU**: the department
	// *   **L**: the city
	// *   **ST**: the province, municipality, or autonomous region
	// *   **CN**: the common name
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// The content of the certificate.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of revoked certificates to return on each page. Default value: **20**.
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
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of pages returned.
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of revoked certificates that are returned per page.
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of revoked client certificates and server certificates that are returned.
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
	AfterDate *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The type of the encryption algorithm of the certificate. Valid values:
	//
	// *   **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	// *   **ECC**: the elliptic curve cryptography (ECC) algorithm.
	// *   **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance date of the certificate. The date is in the `yyyy-MM-ddT00:00Z` format. For example, the value `2021-01-01T00:00Z` indicates January 1, 2021.
	BeforeDate *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the certificate.
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name of the certificate.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// For more information about country codes, see the **"Country codes"** section of the [Manage company profiles](~~198289~~) topic.
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The unique identifier of the certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the certificate.
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the certificate.
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization. The organization is associated with the intermediate certificate from which the certificate is issued.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department in the organization. The organization is associated with the intermediate certificate authority (CA) certificate from which the certificate is issued.
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The identifier of the root certificate.
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The date on which the certificate was revoked. The value is in the `yyyy-MM-ddT00:00Z` format. For example, the value `2021-09-01T00:00Z` indicates September 1, 2021.
	RevokeDate *string `json:"RevokeDate,omitempty" xml:"RevokeDate,omitempty"`
	// The subject alternative name (SAN) extension of the certificate.
	//
	// The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that corresponds to a SAN extension. A SAN extension struct contains the following parameters:
	//
	// *   **Type**: the type of the extension. Data type: integer. Valid values:
	//
	//     *   **1**: an email address
	//     *   **2**: a domain name
	//     *   **6**: a Uniform Resource Identifier (URI)
	//     *   **7**: an IP address
	//
	// *   **Value**: the value of the extension. Data type: string.
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate.
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the certificate.
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The distinguished name (DN) extension of the certificate, which indicates the user of the certificate. The DN extension includes the following information:
	//
	// *   **C**: the country
	// *   **O**: the organization
	// *   **OU**: the department
	// *   **L**: the city
	// *   **ST**: the province, municipality, or autonomous region
	// *   **CN**: the common name
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRevokeCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// >  You can call the [DescribeCACertificateList](~~328095~~) operation to query the unique identifiers of all CA certificates.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The state to which you want to change the CA certificate. Set to the value to **REVOKE**. After this operation is called, the status of the CA certificate is changed to **REVOKE**.
	//
	// >  You can call this operation only if the status of a CA certificate is **ISSUE**. You can call the [DescribeCACertificate](~~328096~~) operation to query the status of a CA certificate.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCACertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
		"ap-southeast-1":              tea.String("cas.aliyuncs.com"),
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
		"cn-yushanfang":               tea.String("cas.aliyuncs.com"),
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

/**
 * Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~328093~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation. Only intermediate CA certificates can be used to issue client certificates.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateClientCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateClientCertificateResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.Csr)) {
		query["Csr"] = request.Csr
	}

	if !tea.BoolValue(util.IsUnset(request.Days)) {
		query["Days"] = request.Days
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

/**
 * Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~328093~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation. Only intermediate CA certificates can be used to issue client certificates.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateClientCertificateRequest
 * @return CreateClientCertificateResponse
 */
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

/**
 * Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~328093~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation. Only intermediate CA certificates can be used to issue client certificates.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateClientCertificateWithCsrRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateClientCertificateWithCsrResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.Csr1)) {
		query["Csr1"] = request.Csr1
	}

	if !tea.BoolValue(util.IsUnset(request.Days)) {
		query["Days"] = request.Days
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

/**
 * Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~328093~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation. Only intermediate CA certificates can be used to issue client certificates.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateClientCertificateWithCsrRequest
 * @return CreateClientCertificateWithCsrResponse
 */
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

/**
 * After a client certificate or a server certificate is revoked, the client or the server on which the certificate is installed cannot establish HTTPS connections with other devices.
 * After a client certificate or a server certificate is revoked, you can call the [DeleteClientCertificate](~~330880~~) operation to permanently delete the certificate.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateRevokeClientCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateRevokeClientCertificateResponse
 */
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

/**
 * After a client certificate or a server certificate is revoked, the client or the server on which the certificate is installed cannot establish HTTPS connections with other devices.
 * After a client certificate or a server certificate is revoked, you can call the [DeleteClientCertificate](~~330880~~) operation to permanently delete the certificate.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateRevokeClientCertificateRequest
 * @return CreateRevokeClientCertificateResponse
 */
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

/**
 * You can call the CreateRootCACertificate operation to create a self-signed root CA certificate. A root CA certificate is the trust anchor in a chain of trust for private certificates that are used within an enterprise. You must create a root CA certificate before you can use the root CA certificate to issue intermediate CA certificates. Then, you can use the intermediate CA certificates to issue client certificates and server certificates.
 * Before you call this operation, make sure that you have purchased a private root CA instance by using the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](~~208553~~).
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateRootCACertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateRootCACertificateResponse
 */
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

/**
 * You can call the CreateRootCACertificate operation to create a self-signed root CA certificate. A root CA certificate is the trust anchor in a chain of trust for private certificates that are used within an enterprise. You must create a root CA certificate before you can use the root CA certificate to issue intermediate CA certificates. Then, you can use the intermediate CA certificates to issue client certificates and server certificates.
 * Before you call this operation, make sure that you have purchased a private root CA instance by using the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](~~208553~~).
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateRootCACertificateRequest
 * @return CreateRootCACertificateResponse
 */
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

/**
 * Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~328093~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation. Only intermediate CA certificates can be used to issue server certificates.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateServerCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateServerCertificateResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.Csr)) {
		query["Csr"] = request.Csr
	}

	if !tea.BoolValue(util.IsUnset(request.Days)) {
		query["Days"] = request.Days
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
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

/**
 * Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~328093~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation. Only intermediate CA certificates can be used to issue server certificates.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateServerCertificateRequest
 * @return CreateServerCertificateResponse
 */
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

/**
 * Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~328093~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation. Only intermediate CA certificates can be used to issue server certificates.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateServerCertificateWithCsrRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateServerCertificateWithCsrResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.Csr1)) {
		query["Csr1"] = request.Csr1
	}

	if !tea.BoolValue(util.IsUnset(request.Days)) {
		query["Days"] = request.Days
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
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

/**
 * Before you call this operation, make sure that you have created a root certificate authority (CA) certificate by calling the [CreateRootCACertificate](~~328093~~) operation and an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation. Only intermediate CA certificates can be used to issue server certificates.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateServerCertificateWithCsrRequest
 * @return CreateServerCertificateWithCsrResponse
 */
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

/**
 * You can call the CreateSubCACertificate operation to issue an intermediate CA certificate by using an existing root CA certificate. Intermediate CA certificates can be used to issue client certificates and server certificates.
 * Before you call this operation, make sure that you have created a root CA certificate by calling the [CreateRootCACertificate](~~328093~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateSubCACertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateSubCACertificateResponse
 */
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

/**
 * You can call the CreateSubCACertificate operation to issue an intermediate CA certificate by using an existing root CA certificate. Intermediate CA certificates can be used to issue client certificates and server certificates.
 * Before you call this operation, make sure that you have created a root CA certificate by calling the [CreateRootCACertificate](~~328093~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateSubCACertificateRequest
 * @return CreateSubCACertificateResponse
 */
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

/**
 * Before you call this operation, you must call the [CreateRevokeClientCertificate](~~330876~~) operation to revoke a client certificate or a server certificate.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteClientCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteClientCertificateResponse
 */
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

/**
 * Before you call this operation, you must call the [CreateRevokeClientCertificate](~~330876~~) operation to revoke a client certificate or a server certificate.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteClientCertificateRequest
 * @return DeleteClientCertificateResponse
 */
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

/**
 * You can call the DescribeCACertificate operation to query the details about a root CA certificate or an intermediate CA certificate by using the unique identifier of the root CA certificate or intermediate CA certificate. The details include the serial number, user information, and content of a CA certificate.
 * Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate](~~328093~~) operation or an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeCACertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeCACertificateResponse
 */
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

/**
 * You can call the DescribeCACertificate operation to query the details about a root CA certificate or an intermediate CA certificate by using the unique identifier of the root CA certificate or intermediate CA certificate. The details include the serial number, user information, and content of a CA certificate.
 * Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate](~~328093~~) operation or an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeCACertificateRequest
 * @return DescribeCACertificateResponse
 */
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

/**
 * You can call the DescribeCACertificateCount operation to query the number of created CA certificates, which includes root CA certificates and intermediate CA certificates.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeCACertificateCountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeCACertificateCountResponse
 */
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

/**
 * You can call the DescribeCACertificateCount operation to query the number of created CA certificates, which includes root CA certificates and intermediate CA certificates.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @return DescribeCACertificateCountResponse
 */
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

/**
 * You can call the DescribeCACertificateList operation to perform a paged query of the details about all CA certificates that you create. The details include the unique identifier, serial number, user information, and content of each root CA certificate or intermediate CA certificate.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeCACertificateListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeCACertificateListResponse
 */
func (client *Client) DescribeCACertificateListWithOptions(request *DescribeCACertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeCACertificateListResponse, _err error) {
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

/**
 * You can call the DescribeCACertificateList operation to perform a paged query of the details about all CA certificates that you create. The details include the unique identifier, serial number, user information, and content of each root CA certificate or intermediate CA certificate.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeCACertificateListRequest
 * @return DescribeCACertificateListResponse
 */
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

/**
 * ## Usage notes
 * You can call the DescribeCertificatePrivateKey operation to obtain the encrypted private key of a client certificate or a server certificate. The certificate is issued based on a system-generated certificate signing request (CSR). Before you call this operation, make sure that you have issued a client certificate or a server certificate by calling the following operation:
 * *   [CreateClientCertificate](~~330873~~)
 * *   [CreateServerCertificate](~~330877~~)
 * To ensure the security of private key transmission, the DescribeCertificatePrivateKey operation encrypts the private key by using the private key password that you specify and returns the encrypted private key. The private key password is an string that is used to encrypt the private key. After you obtain the encrypted private key of the certificate, you can use the following methods to decrypt the private key:
 * *   If the encryption algorithm of the certificate is RSA, you must run the `openssl rsa -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
 * *   If the encryption algorithm of the certificate is ECC, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
 * *   If the encryption algorithm of the certificate is SM2, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
 * >  You can call the [DescribeClientCertificate](~~329929~~) operation to query the encryption algorithm type of a client certificate or a server certificate.
 * ## Limits
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeCertificatePrivateKeyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeCertificatePrivateKeyResponse
 */
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

/**
 * ## Usage notes
 * You can call the DescribeCertificatePrivateKey operation to obtain the encrypted private key of a client certificate or a server certificate. The certificate is issued based on a system-generated certificate signing request (CSR). Before you call this operation, make sure that you have issued a client certificate or a server certificate by calling the following operation:
 * *   [CreateClientCertificate](~~330873~~)
 * *   [CreateServerCertificate](~~330877~~)
 * To ensure the security of private key transmission, the DescribeCertificatePrivateKey operation encrypts the private key by using the private key password that you specify and returns the encrypted private key. The private key password is an string that is used to encrypt the private key. After you obtain the encrypted private key of the certificate, you can use the following methods to decrypt the private key:
 * *   If the encryption algorithm of the certificate is RSA, you must run the `openssl rsa -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
 * *   If the encryption algorithm of the certificate is ECC, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [OpenSSL](https://www.openssl.org/source/) or [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
 * *   If the encryption algorithm of the certificate is SM2, you must run the `openssl ec -in <Encrypted private key file> -passin pass:<Private key password> -out <Decrypted private key file>` command in the computer on which [BabaSSL](https://github.com/BabaSSL/BabaSSL) is installed.
 * >  You can call the [DescribeClientCertificate](~~329929~~) operation to query the encryption algorithm type of a client certificate or a server certificate.
 * ## Limits
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeCertificatePrivateKeyRequest
 * @return DescribeCertificatePrivateKeyResponse
 */
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

/**
 * You can call the DescribeClientCertificate operation to query the details about a client certificate or a server certificate by using the unique identifier of the certificate. The details include the serial number, user information, content, and status of each certificate.
 * Before you call this operation, make sure that you have created a client certificate or a server certificate.
 * For more information about how to call an operation to create a client certificate, see the following topics:
 * *   [CreateClientCertificate](~~330873~~)
 * *   [CreateClientCertificateWithCsr](~~330875~~)
 * For more information about how to call an operation to create a server certificate, see the following topics:
 * *   [CreateServerCertificate](~~330877~~)
 * *   [CreateServerCertificateWithCsr](~~330878~~)
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeClientCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeClientCertificateResponse
 */
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

/**
 * You can call the DescribeClientCertificate operation to query the details about a client certificate or a server certificate by using the unique identifier of the certificate. The details include the serial number, user information, content, and status of each certificate.
 * Before you call this operation, make sure that you have created a client certificate or a server certificate.
 * For more information about how to call an operation to create a client certificate, see the following topics:
 * *   [CreateClientCertificate](~~330873~~)
 * *   [CreateClientCertificateWithCsr](~~330875~~)
 * For more information about how to call an operation to create a server certificate, see the following topics:
 * *   [CreateServerCertificate](~~330877~~)
 * *   [CreateServerCertificateWithCsr](~~330878~~)
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeClientCertificateRequest
 * @return DescribeClientCertificateResponse
 */
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

/**
 * You can call the DescribeClientCertificateStatus operation to query the status information about multiple client certificates or server certificates at a time by using the unique identifiers of the certificates. For example, you can check whether a certificate is revoked.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeClientCertificateStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeClientCertificateStatusResponse
 */
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

/**
 * You can call the DescribeClientCertificateStatus operation to query the status information about multiple client certificates or server certificates at a time by using the unique identifiers of the certificates. For example, you can check whether a certificate is revoked.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeClientCertificateStatusRequest
 * @return DescribeClientCertificateStatusResponse
 */
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

/**
 * You can call the GetCAInstanceStatus operation to query the status information about a private CA instance by using the ID of the instance. The instance is purchased by using the Certificate Management Service console. The status information includes the status of the private CA instance, the number of certificates that can be issued by using the private CA instance, and the number of issued certificates.
 * Before you call this operation, make sure that you have purchased a private CA by using the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](~~208553~~).
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request GetCAInstanceStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetCAInstanceStatusResponse
 */
func (client *Client) GetCAInstanceStatusWithOptions(request *GetCAInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *GetCAInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

/**
 * You can call the GetCAInstanceStatus operation to query the status information about a private CA instance by using the ID of the instance. The instance is purchased by using the Certificate Management Service console. The status information includes the status of the private CA instance, the number of certificates that can be issued by using the private CA instance, and the number of issued certificates.
 * Before you call this operation, make sure that you have purchased a private CA by using the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist). For more information, see [Create a private CA](~~208553~~).
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request GetCAInstanceStatusRequest
 * @return GetCAInstanceStatusResponse
 */
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

/**
 * You can call the ListClientCertificate operation to perform a paged query of the details about all client certificates and server certificates that you create. The details include the unique identifier, serial number, user information, content, and status of each certificate.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ListClientCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListClientCertificateResponse
 */
func (client *Client) ListClientCertificateWithOptions(request *ListClientCertificateRequest, runtime *util.RuntimeOptions) (_result *ListClientCertificateResponse, _err error) {
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

/**
 * You can call the ListClientCertificate operation to perform a paged query of the details about all client certificates and server certificates that you create. The details include the unique identifier, serial number, user information, content, and status of each certificate.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ListClientCertificateRequest
 * @return ListClientCertificateResponse
 */
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

/**
 * You can call the ListRevokeCertificate operation to perform a paged query of the details about all revoked client certificates and server certificates. The details include the unique identifier, serial number, and revocation date of each certificate.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ListRevokeCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListRevokeCertificateResponse
 */
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

/**
 * You can call the ListRevokeCertificate operation to perform a paged query of the details about all revoked client certificates and server certificates. The details include the unique identifier, serial number, and revocation date of each certificate.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ListRevokeCertificateRequest
 * @return ListRevokeCertificateResponse
 */
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

/**
 * After a CA certificate is created, the CA certificate is in the ISSUE state by default. You can call the UpdateCACertificateStatus operation to change the status of a CA certificate from ISSUE to REVOKE. If a CA certificate is in the ISSUE state, the CA certificate can be used to issue certificates. If a CA certificate is in the REVOKE state, the CA certificate cannot be used to issue certificates, and the certificates that are issued from the CA certificate become invalid.
 * Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate](~~328093~~) operation or an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UpdateCACertificateStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateCACertificateStatusResponse
 */
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

/**
 * After a CA certificate is created, the CA certificate is in the ISSUE state by default. You can call the UpdateCACertificateStatus operation to change the status of a CA certificate from ISSUE to REVOKE. If a CA certificate is in the ISSUE state, the CA certificate can be used to issue certificates. If a CA certificate is in the REVOKE state, the CA certificate cannot be used to issue certificates, and the certificates that are issued from the CA certificate become invalid.
 * Before you call this operation, make sure that you have created a root CA by calling the [CreateRootCACertificate](~~328093~~) operation or an intermediate CA certificate by calling the [CreateSubCACertificate](~~328094~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UpdateCACertificateStatusRequest
 * @return UpdateCACertificateStatusResponse
 */
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
