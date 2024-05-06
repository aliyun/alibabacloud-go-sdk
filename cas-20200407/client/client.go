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

type CancelCertificateForPackageRequestRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CancelCertificateForPackageRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateForPackageRequestRequest) GoString() string {
	return s.String()
}

func (s *CancelCertificateForPackageRequestRequest) SetOrderId(v int64) *CancelCertificateForPackageRequestRequest {
	s.OrderId = &v
	return s
}

type CancelCertificateForPackageRequestResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelCertificateForPackageRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateForPackageRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCertificateForPackageRequestResponseBody) SetRequestId(v string) *CancelCertificateForPackageRequestResponseBody {
	s.RequestId = &v
	return s
}

type CancelCertificateForPackageRequestResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCertificateForPackageRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCertificateForPackageRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateForPackageRequestResponse) GoString() string {
	return s.String()
}

func (s *CancelCertificateForPackageRequestResponse) SetHeaders(v map[string]*string) *CancelCertificateForPackageRequestResponse {
	s.Headers = v
	return s
}

func (s *CancelCertificateForPackageRequestResponse) SetStatusCode(v int32) *CancelCertificateForPackageRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCertificateForPackageRequestResponse) SetBody(v *CancelCertificateForPackageRequestResponseBody) *CancelCertificateForPackageRequestResponse {
	s.Body = v
	return s
}

type CancelOrderRequestRequest struct {
	// The ID of the certificate application order that you want to cancel.
	//
	// >  After you call the [CreateCertificateForPackageRequest](~~CreateCertificateForPackageRequest~~), [CreateCertificateRequest](~~CreateCertificateRequest~~), or [CreateCertificateWithCsrRequest](~~CreateCertificateWithCsrRequest~~) operation to submit a certificate application, you can obtain the ID of the certificate application order from the **OrderId** response parameter.
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CancelOrderRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRequestRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRequestRequest) SetOrderId(v int64) *CancelOrderRequestRequest {
	s.OrderId = &v
	return s
}

type CancelOrderRequestResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOrderRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderRequestResponseBody) SetRequestId(v string) *CancelOrderRequestResponseBody {
	s.RequestId = &v
	return s
}

type CancelOrderRequestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOrderRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOrderRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRequestResponse) GoString() string {
	return s.String()
}

func (s *CancelOrderRequestResponse) SetHeaders(v map[string]*string) *CancelOrderRequestResponse {
	s.Headers = v
	return s
}

func (s *CancelOrderRequestResponse) SetStatusCode(v int32) *CancelOrderRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOrderRequestResponse) SetBody(v *CancelOrderRequestResponseBody) *CancelOrderRequestResponse {
	s.Body = v
	return s
}

type CreateCertificateForPackageRequestRequest struct {
	// The company name of the certificate application.
	//
	// > This parameter is available only when you apply for OV certificates. If you want to apply for an OV certificate, you must add a company profile to the **Information Management** module of the [Certificate Management Service console](https://yundun.console.aliyun.com/?p=cas#/). For more information, see [Manage company profiles](~~198289~~). If you want to apply for a DV certificate, you do not need to add a company profile.
	//
	// If you specify a company name, the information about the company that is configured in the **Information Management** module is used. If you do not specify this parameter, the information about the most recent company that is added to the **Information Management** module is used.
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// The content of the certificate signing request (CSR) file that is manually generated for the domain name by using OpenSSL or Keytool. The key algorithm in the CSR file must be Rivest-Shamir-Adleman (RSA) or elliptic-curve cryptography (ECC), and the key length of the RSA algorithm must be greater than or equal to 2,048 characters. For more information about how to create a CSR file, see [Create a CSR file](~~313297~~). If you do not specify this parameter, Certificate Management Service automatically creates a CSR file.
	//
	// A CSR file contains the information about your server and company. When you apply for a certificate, you must submit the CSR file to the CA. The CA signs the CSR file by using the private key of the root certificate and generates a public key file to issue your certificate.
	//
	// >
	//
	// The **CN** field in the CSR file specifies the domain name that you want to bind to the certificate. You must include the field in the parameter value.
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The domain name that you want to bind to the certificate. The domain name must meet the following requirements:
	//
	// *   The domain name must be a single domain name or a wildcard domain name. Example: `*.aliyundoc.com`.
	// *   You can specify multiple domain names. Separate multiple domain names with commas (,). You can specify a maximum of five domain names.
	// *   If you specify multiple domain names, the domain names must be only single domain names or only wildcard domain names. You cannot specify both single domain names and wildcard domain names.
	//
	// >
	//
	// If you want to bind multiple domain names to the certificate, you must specify this parameter. You must specify at least one of the Domain parameter and the **Csr** parameter. If you specify both the Domain parameter and the **Csr** parameter, the value of the **CN** field in the **Csr** parameter is used as the domain name that can be bound to the certificate.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The email address of the applicant. After the CA receives your certificate application, the CA sends a verification email to the email address that you specify. You must log on to the mailbox, open the mail, and complete the verification of the domain name ownership based on the steps that are described in the email.
	//
	// If you do not specify this parameter, the information about the most recent contact that is added to the **Information Management** module is used. For more information about how to add a contact to the **Information Management** module, see [Manage contacts](~~198262~~).
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The phone number of the applicant. CA staff can call the phone number to confirm the information in your certificate application.
	//
	// If you do not specify this parameter, the information about the most recent contact that is added to the **Information Management** module is used. For more information about how to add a contact to the **Information Management** module, see [Manage contacts](~~198262~~).
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The specifications of the certificate. Valid values:
	//
	// *   **digicert-free-1-free**: DigiCert single-domain domain validated (DV) certificate in 3 months free trial. This is the default value.
	// *   **symantec-free-1-free**: DigiCert single-domain domain validated (DV) certificate in 1 year free trial.
	// *   **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	// *   **symantec-ov-1-personal**: DigiCert single-domain organization validated (OV) certificate.
	// *   **symantec-ov-w-personal**: DigiCert wildcard OV certificate.
	// *   **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	// *   **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	// *   **geotrust-ov-1-personal**: GeoTrust single-domain OV certificate.
	// *   **geotrust-ov-w-personal**: GeoTrust wildcard OV certificate.
	// *   **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	// *   **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	// *   **globalsign-ov-1-personal**: GlobalSign single-domain OV certificate.
	// *   **globalsign-ov-w-advanced**: GlobalSign wildcard OV certificate.
	// *   **cfca-ov-1-personal**: China Financial Certification Authority (CFCA) single-domain OV certificate.
	// *   **cfca-ev-w-advanced**: CFCA wildcard OV certificate.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the applicant.
	//
	// If you do not specify this parameter, the information about the most recent contact that is added to the **Information Management** module is used. For more information about how to add a contact to the **Information Management** module, see [Manage contacts](~~198262~~).
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The verification method of the domain name ownership. Valid values:
	//
	// *   **DNS**: DNS verification. If you use this method, you must add a TXT record to the DNS records of the domain name in the management platform of the domain name. You must have operation permissions on domain name resolution to verify the ownership of the domain name.
	// *   **FILE**: file verification. If you use this method, you must create a specified file on the DNS server. You must have administrative rights on the DNS server to verify the ownership of the domain name.
	//
	// For more information about the verification methods, see [Verify the ownership of a domain name](~~48016~~).
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
}

func (s CreateCertificateForPackageRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateForPackageRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateForPackageRequestRequest) SetCompanyName(v string) *CreateCertificateForPackageRequestRequest {
	s.CompanyName = &v
	return s
}

func (s *CreateCertificateForPackageRequestRequest) SetCsr(v string) *CreateCertificateForPackageRequestRequest {
	s.Csr = &v
	return s
}

func (s *CreateCertificateForPackageRequestRequest) SetDomain(v string) *CreateCertificateForPackageRequestRequest {
	s.Domain = &v
	return s
}

func (s *CreateCertificateForPackageRequestRequest) SetEmail(v string) *CreateCertificateForPackageRequestRequest {
	s.Email = &v
	return s
}

func (s *CreateCertificateForPackageRequestRequest) SetPhone(v string) *CreateCertificateForPackageRequestRequest {
	s.Phone = &v
	return s
}

func (s *CreateCertificateForPackageRequestRequest) SetProductCode(v string) *CreateCertificateForPackageRequestRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateCertificateForPackageRequestRequest) SetUsername(v string) *CreateCertificateForPackageRequestRequest {
	s.Username = &v
	return s
}

func (s *CreateCertificateForPackageRequestRequest) SetValidateType(v string) *CreateCertificateForPackageRequestRequest {
	s.ValidateType = &v
	return s
}

type CreateCertificateForPackageRequestResponseBody struct {
	// The ID of the certificate application order.
	//
	// > You can use the ID to query the status of the certificate application order. For more information, see [DescribeCertificateState](~~455800~~).
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertificateForPackageRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateForPackageRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateForPackageRequestResponseBody) SetOrderId(v int64) *CreateCertificateForPackageRequestResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCertificateForPackageRequestResponseBody) SetRequestId(v string) *CreateCertificateForPackageRequestResponseBody {
	s.RequestId = &v
	return s
}

type CreateCertificateForPackageRequestResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertificateForPackageRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCertificateForPackageRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateForPackageRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateForPackageRequestResponse) SetHeaders(v map[string]*string) *CreateCertificateForPackageRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateForPackageRequestResponse) SetStatusCode(v int32) *CreateCertificateForPackageRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertificateForPackageRequestResponse) SetBody(v *CreateCertificateForPackageRequestResponseBody) *CreateCertificateForPackageRequestResponse {
	s.Body = v
	return s
}

type CreateCertificateRequestRequest struct {
	// The domain name that you want to bind to the certificate. You can specify only one domain name.
	//
	// > The domain name must match the certificate specifications that you specify for the **ProductCode** parameter. If you apply for a single-domain certificate, you must specify a single domain name for this parameter. If you apply for a wildcard certificate, you must specify a wildcard domain name such as `*.aliyundoc.com` for this parameter.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The email address of the applicant.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The phone number of the applicant.
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The specifications of the certificate. Valid values:
	//
	// *   **digicert-free-1-free**: DigiCert single-domain DV certificate in 3 months free trial. This is the default value.
	// *   **symantec-free-1-free**: DigiCert single-domain DV certificate in 1 year free trial.
	// *   **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	// *   **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	// *   **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	// *   **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	// *   **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the applicant.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The verification method of the domain name ownership. Valid values:
	//
	// *   **DNS**: DNS verification. If you use this method, you must add a TXT record to the DNS records of the domain name in the management platform of the domain name. You must have operation permissions on domain name resolution to verify the ownership of the domain name.
	// *   **FILE**: file verification. If you use this method, you must create a specified file on the DNS server. You must have administrative rights on the DNS server to verify the ownership of the domain name.
	//
	// For more information about the verification methods, see [Verify the ownership of a domain name](~~48016~~).
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
}

func (s CreateCertificateRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequestRequest) SetDomain(v string) *CreateCertificateRequestRequest {
	s.Domain = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetEmail(v string) *CreateCertificateRequestRequest {
	s.Email = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetPhone(v string) *CreateCertificateRequestRequest {
	s.Phone = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetProductCode(v string) *CreateCertificateRequestRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetUsername(v string) *CreateCertificateRequestRequest {
	s.Username = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetValidateType(v string) *CreateCertificateRequestRequest {
	s.ValidateType = &v
	return s
}

type CreateCertificateRequestResponseBody struct {
	// The ID of the certificate application order.
	//
	// > You can use the ID to query the status of the certificate application. For more information, see [DescribeCertificateState](~~455800~~).
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertificateRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequestResponseBody) SetOrderId(v int64) *CreateCertificateRequestResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCertificateRequestResponseBody) SetRequestId(v string) *CreateCertificateRequestResponseBody {
	s.RequestId = &v
	return s
}

type CreateCertificateRequestResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertificateRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCertificateRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequestResponse) SetHeaders(v map[string]*string) *CreateCertificateRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateRequestResponse) SetStatusCode(v int32) *CreateCertificateRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertificateRequestResponse) SetBody(v *CreateCertificateRequestResponseBody) *CreateCertificateRequestResponse {
	s.Body = v
	return s
}

type CreateCertificateWithCsrRequestRequest struct {
	// The content of the existing CSR file.\
	// The key algorithm in the CSR file must be Rivest-Shamir-Adleman (RSA) or elliptic-curve cryptography (ECC), and the key length of the RSA algorithm must be greater than or equal to 2,048 characters. For more information about how to create a CSR file, see [How do I create a CSR file?](~~42218~~) You can also create a CSR in the [Certificate Management Service console](https://yundunnext.console.aliyun.com/?\&p=cas). For more information, see [Create a CSR](~~313297~~).\
	// A CSR file contains the information about your server and company. When you apply for a certificate, you must submit the CSR file to the CA. The CA signs the CSR file by using the private key of the root certificate and generates a public key file to issue your certificate.
	//
	// >  The **CN** field in the CSR file specifies the domain name that is bound to the certificate.
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The contact email address of the applicant.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The phone number of the applicant.
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The specifications of the certificate. Valid values:
	//
	// *   **digicert-free-1-free**: DigiCert single-domain DV certificate in 3 months free trial. This is the default value.
	// *   **symantec-free-1-free**: DigiCert single-domain DV certificate in 1 year free trial.
	// *   **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	// *   **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	// *   **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	// *   **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	// *   **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the applicant.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The method to verify the ownership of a domain name. Valid values:
	//
	// *   **DNS**: DNS verification. If you use this method, you must add a TXT record to the DNS records of the domain name in the management platform of the domain name. You must have operation permissions on domain name resolution to verify the ownership of the domain name.
	// *   **FILE**: file verification. If you use this method, you must create a specified file on the DNS server. You must have administrative rights on the DNS server to verify the ownership of the domain name.
	//
	// For more information about the verification methods, see [Verify the ownership of a domain name](~~48016~~).
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
}

func (s CreateCertificateWithCsrRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateWithCsrRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithCsrRequestRequest) SetCsr(v string) *CreateCertificateWithCsrRequestRequest {
	s.Csr = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetEmail(v string) *CreateCertificateWithCsrRequestRequest {
	s.Email = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetPhone(v string) *CreateCertificateWithCsrRequestRequest {
	s.Phone = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetProductCode(v string) *CreateCertificateWithCsrRequestRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetUsername(v string) *CreateCertificateWithCsrRequestRequest {
	s.Username = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetValidateType(v string) *CreateCertificateWithCsrRequestRequest {
	s.ValidateType = &v
	return s
}

type CreateCertificateWithCsrRequestResponseBody struct {
	// The ID of the certificate application order.
	//
	// >  You can use the ID to query the status of the certificate application. For more information, see [DescribeCertificateState](~~164111~~).
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertificateWithCsrRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateWithCsrRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithCsrRequestResponseBody) SetOrderId(v int64) *CreateCertificateWithCsrRequestResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCertificateWithCsrRequestResponseBody) SetRequestId(v string) *CreateCertificateWithCsrRequestResponseBody {
	s.RequestId = &v
	return s
}

type CreateCertificateWithCsrRequestResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertificateWithCsrRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCertificateWithCsrRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateWithCsrRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithCsrRequestResponse) SetHeaders(v map[string]*string) *CreateCertificateWithCsrRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateWithCsrRequestResponse) SetStatusCode(v int32) *CreateCertificateWithCsrRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertificateWithCsrRequestResponse) SetBody(v *CreateCertificateWithCsrRequestResponseBody) *CreateCertificateWithCsrRequestResponse {
	s.Body = v
	return s
}

type CreateCsrRequest struct {
	Algorithm   *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CommonName  *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CorpName    *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	KeySize     *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Locality    *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Sans        *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
}

func (s CreateCsrRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCsrRequest) GoString() string {
	return s.String()
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

type CreateCsrResponseBody struct {
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// CSR ID。
	CsrId     *int64  `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCsrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCsrResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCsrResponseBody) SetCsr(v string) *CreateCsrResponseBody {
	s.Csr = &v
	return s
}

func (s *CreateCsrResponseBody) SetCsrId(v int64) *CreateCsrResponseBody {
	s.CsrId = &v
	return s
}

func (s *CreateCsrResponseBody) SetRequestId(v string) *CreateCsrResponseBody {
	s.RequestId = &v
	return s
}

type CreateCsrResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCsrResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCsrResponse) GoString() string {
	return s.String()
}

func (s *CreateCsrResponse) SetHeaders(v map[string]*string) *CreateCsrResponse {
	s.Headers = v
	return s
}

func (s *CreateCsrResponse) SetStatusCode(v int32) *CreateCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCsrResponse) SetBody(v *CreateCsrResponseBody) *CreateCsrResponse {
	s.Body = v
	return s
}

type CreateDeploymentJobRequest struct {
	CertIds      *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	ContactIds   *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	JobType      *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceIds  *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ScheduleTime *int64  `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s CreateDeploymentJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentJobRequest) SetCertIds(v string) *CreateDeploymentJobRequest {
	s.CertIds = &v
	return s
}

func (s *CreateDeploymentJobRequest) SetContactIds(v string) *CreateDeploymentJobRequest {
	s.ContactIds = &v
	return s
}

func (s *CreateDeploymentJobRequest) SetJobType(v string) *CreateDeploymentJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateDeploymentJobRequest) SetName(v string) *CreateDeploymentJobRequest {
	s.Name = &v
	return s
}

func (s *CreateDeploymentJobRequest) SetResourceIds(v string) *CreateDeploymentJobRequest {
	s.ResourceIds = &v
	return s
}

func (s *CreateDeploymentJobRequest) SetScheduleTime(v int64) *CreateDeploymentJobRequest {
	s.ScheduleTime = &v
	return s
}

type CreateDeploymentJobResponseBody struct {
	JobId     *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeploymentJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentJobResponseBody) SetJobId(v int64) *CreateDeploymentJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateDeploymentJobResponseBody) SetRequestId(v string) *CreateDeploymentJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateDeploymentJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentJobResponse) SetHeaders(v map[string]*string) *CreateDeploymentJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentJobResponse) SetStatusCode(v int32) *CreateDeploymentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentJobResponse) SetBody(v *CreateDeploymentJobResponseBody) *CreateDeploymentJobResponse {
	s.Body = v
	return s
}

type CreateWHClientCertificateRequest struct {
	AfterTime        *int64  `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeTime       *int64  `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	Country          *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Csr              *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	Days             *int64  `json:"Days,omitempty" xml:"Days,omitempty"`
	Immediately      *int64  `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Months           *int64  `json:"Months,omitempty" xml:"Months,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	SanType          *int64  `json:"SanType,omitempty" xml:"SanType,omitempty"`
	SanValue         *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	Years            *int64  `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateWHClientCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWHClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateWHClientCertificateRequest) SetAfterTime(v int64) *CreateWHClientCertificateRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetAlgorithm(v string) *CreateWHClientCertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetBeforeTime(v int64) *CreateWHClientCertificateRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetCommonName(v string) *CreateWHClientCertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetCountry(v string) *CreateWHClientCertificateRequest {
	s.Country = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetCsr(v string) *CreateWHClientCertificateRequest {
	s.Csr = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetDays(v int64) *CreateWHClientCertificateRequest {
	s.Days = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetImmediately(v int64) *CreateWHClientCertificateRequest {
	s.Immediately = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetLocality(v string) *CreateWHClientCertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetMonths(v int64) *CreateWHClientCertificateRequest {
	s.Months = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetOrganization(v string) *CreateWHClientCertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetOrganizationUnit(v string) *CreateWHClientCertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetParentIdentifier(v string) *CreateWHClientCertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetSanType(v int64) *CreateWHClientCertificateRequest {
	s.SanType = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetSanValue(v string) *CreateWHClientCertificateRequest {
	s.SanValue = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetState(v string) *CreateWHClientCertificateRequest {
	s.State = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetYears(v int64) *CreateWHClientCertificateRequest {
	s.Years = &v
	return s
}

type CreateWHClientCertificateResponseBody struct {
	CertificateChain      *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	Identifier            *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	ParentX509Certificate *string `json:"ParentX509Certificate,omitempty" xml:"ParentX509Certificate,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootX509Certificate   *string `json:"RootX509Certificate,omitempty" xml:"RootX509Certificate,omitempty"`
	X509Certificate       *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s CreateWHClientCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWHClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWHClientCertificateResponseBody) SetCertificateChain(v string) *CreateWHClientCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateWHClientCertificateResponseBody) SetIdentifier(v string) *CreateWHClientCertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateWHClientCertificateResponseBody) SetParentX509Certificate(v string) *CreateWHClientCertificateResponseBody {
	s.ParentX509Certificate = &v
	return s
}

func (s *CreateWHClientCertificateResponseBody) SetRequestId(v string) *CreateWHClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWHClientCertificateResponseBody) SetRootX509Certificate(v string) *CreateWHClientCertificateResponseBody {
	s.RootX509Certificate = &v
	return s
}

func (s *CreateWHClientCertificateResponseBody) SetX509Certificate(v string) *CreateWHClientCertificateResponseBody {
	s.X509Certificate = &v
	return s
}

type CreateWHClientCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWHClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWHClientCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWHClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateWHClientCertificateResponse) SetHeaders(v map[string]*string) *CreateWHClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateWHClientCertificateResponse) SetStatusCode(v int32) *CreateWHClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWHClientCertificateResponse) SetBody(v *CreateWHClientCertificateResponseBody) *CreateWHClientCertificateResponse {
	s.Body = v
	return s
}

type DecryptRequest struct {
	Algorithm      *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	MessageType    *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s DecryptRequest) String() string {
	return tea.Prettify(s)
}

func (s DecryptRequest) GoString() string {
	return s.String()
}

func (s *DecryptRequest) SetAlgorithm(v string) *DecryptRequest {
	s.Algorithm = &v
	return s
}

func (s *DecryptRequest) SetCertIdentifier(v string) *DecryptRequest {
	s.CertIdentifier = &v
	return s
}

func (s *DecryptRequest) SetCiphertextBlob(v string) *DecryptRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *DecryptRequest) SetMessageType(v string) *DecryptRequest {
	s.MessageType = &v
	return s
}

type DecryptResponseBody struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	Plaintext      *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DecryptResponseBody) GoString() string {
	return s.String()
}

func (s *DecryptResponseBody) SetCertIdentifier(v string) *DecryptResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *DecryptResponseBody) SetPlaintext(v string) *DecryptResponseBody {
	s.Plaintext = &v
	return s
}

func (s *DecryptResponseBody) SetRequestId(v string) *DecryptResponseBody {
	s.RequestId = &v
	return s
}

type DecryptResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DecryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DecryptResponse) String() string {
	return tea.Prettify(s)
}

func (s DecryptResponse) GoString() string {
	return s.String()
}

func (s *DecryptResponse) SetHeaders(v map[string]*string) *DecryptResponse {
	s.Headers = v
	return s
}

func (s *DecryptResponse) SetStatusCode(v int32) *DecryptResponse {
	s.StatusCode = &v
	return s
}

func (s *DecryptResponse) SetBody(v *DecryptResponseBody) *DecryptResponse {
	s.Body = v
	return s
}

type DeleteCertificateRequestRequest struct {
	// The ID of the certificate application order that you want to delete.
	//
	// >  After you call the [CreateCertificateForPackageRequest](~~455296~~), [CreateCertificateRequest](~~455292~~), or [CreateCertificateWithCsrRequest](~~455801~~) operation to submit a certificate application, you can obtain the ID of the certificate application order from the **OrderId** response parameter.
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DeleteCertificateRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertificateRequestRequest) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequestRequest) SetOrderId(v int64) *DeleteCertificateRequestRequest {
	s.OrderId = &v
	return s
}

type DeleteCertificateRequestResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCertificateRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertificateRequestResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequestResponseBody) SetRequestId(v string) *DeleteCertificateRequestResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCertificateRequestResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCertificateRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCertificateRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertificateRequestResponse) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequestResponse) SetHeaders(v map[string]*string) *DeleteCertificateRequestResponse {
	s.Headers = v
	return s
}

func (s *DeleteCertificateRequestResponse) SetStatusCode(v int32) *DeleteCertificateRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCertificateRequestResponse) SetBody(v *DeleteCertificateRequestResponseBody) *DeleteCertificateRequestResponse {
	s.Body = v
	return s
}

type DeleteCsrRequest struct {
	// CSR ID。
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
}

func (s DeleteCsrRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCsrRequest) GoString() string {
	return s.String()
}

func (s *DeleteCsrRequest) SetCsrId(v int64) *DeleteCsrRequest {
	s.CsrId = &v
	return s
}

type DeleteCsrResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCsrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCsrResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCsrResponseBody) SetRequestId(v string) *DeleteCsrResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCsrResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCsrResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCsrResponse) GoString() string {
	return s.String()
}

func (s *DeleteCsrResponse) SetHeaders(v map[string]*string) *DeleteCsrResponse {
	s.Headers = v
	return s
}

func (s *DeleteCsrResponse) SetStatusCode(v int32) *DeleteCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCsrResponse) SetBody(v *DeleteCsrResponseBody) *DeleteCsrResponse {
	s.Body = v
	return s
}

type DeleteDeploymentJobRequest struct {
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteDeploymentJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentJobRequest) SetJobId(v int64) *DeleteDeploymentJobRequest {
	s.JobId = &v
	return s
}

type DeleteDeploymentJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeploymentJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentJobResponseBody) SetRequestId(v string) *DeleteDeploymentJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDeploymentJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeploymentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeploymentJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeploymentJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentJobResponse) SetHeaders(v map[string]*string) *DeleteDeploymentJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeploymentJobResponse) SetStatusCode(v int32) *DeleteDeploymentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeploymentJobResponse) SetBody(v *DeleteDeploymentJobResponseBody) *DeleteDeploymentJobResponse {
	s.Body = v
	return s
}

type DeletePCACertRequest struct {
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DeletePCACertRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePCACertRequest) GoString() string {
	return s.String()
}

func (s *DeletePCACertRequest) SetIdentifier(v string) *DeletePCACertRequest {
	s.Identifier = &v
	return s
}

type DeletePCACertResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePCACertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePCACertResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePCACertResponseBody) SetRequestId(v string) *DeletePCACertResponseBody {
	s.RequestId = &v
	return s
}

type DeletePCACertResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePCACertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePCACertResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePCACertResponse) GoString() string {
	return s.String()
}

func (s *DeletePCACertResponse) SetHeaders(v map[string]*string) *DeletePCACertResponse {
	s.Headers = v
	return s
}

func (s *DeletePCACertResponse) SetStatusCode(v int32) *DeletePCACertResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePCACertResponse) SetBody(v *DeletePCACertResponseBody) *DeletePCACertResponse {
	s.Body = v
	return s
}

type DeleteUserCertificateRequest struct {
	// The ID of the certificate.
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
}

func (s DeleteUserCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserCertificateRequest) SetCertId(v int64) *DeleteUserCertificateRequest {
	s.CertId = &v
	return s
}

type DeleteUserCertificateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserCertificateResponseBody) SetRequestId(v string) *DeleteUserCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserCertificateResponse) SetHeaders(v map[string]*string) *DeleteUserCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserCertificateResponse) SetStatusCode(v int32) *DeleteUserCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserCertificateResponse) SetBody(v *DeleteUserCertificateResponseBody) *DeleteUserCertificateResponse {
	s.Body = v
	return s
}

type DeleteWorkerResourceRequest struct {
	JobId    *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	WorkerId *int64 `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s DeleteWorkerResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkerResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkerResourceRequest) SetJobId(v int64) *DeleteWorkerResourceRequest {
	s.JobId = &v
	return s
}

func (s *DeleteWorkerResourceRequest) SetWorkerId(v int64) *DeleteWorkerResourceRequest {
	s.WorkerId = &v
	return s
}

type DeleteWorkerResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkerResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkerResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkerResourceResponseBody) SetRequestId(v string) *DeleteWorkerResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWorkerResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkerResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkerResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkerResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkerResourceResponse) SetHeaders(v map[string]*string) *DeleteWorkerResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkerResourceResponse) SetStatusCode(v int32) *DeleteWorkerResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkerResourceResponse) SetBody(v *DeleteWorkerResourceResponseBody) *DeleteWorkerResourceResponse {
	s.Body = v
	return s
}

type DescribeCertificateStateRequest struct {
	// The ID of the certificate application order that you want to query.
	//
	// > After you call the [CreateCertificateForPackageRequest](~~455296~~), [CreateCertificateRequest](~~455292~~), or [CreateCertificateWithCsrRequest](~~455801~~) operation to submit a certificate application, you can obtain the ID of the certificate application order from the **OrderId** response parameter.
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DescribeCertificateStateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateStateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificateStateRequest) SetOrderId(v int64) *DescribeCertificateStateRequest {
	s.OrderId = &v
	return s
}

type DescribeCertificateStateResponseBody struct {
	// The content of the certificate in the PEM format. For more information about the PEM format and how to convert certificate formats, see [What formats are used for mainstream digital certificates?](~~42214~~)
	//
	// > This parameter is returned only when the value of the **Type** parameter is **certificate**. The value certificate indicates that the certificate is issued.
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The content that you need to write to the newly created file when you use the file verification method.
	//
	// > This parameter is returned only when the value of the **Type** parameter is **domain\_verify** and the value of the **ValidateType** parameter is **FILE**. The value domain\_verify indicates that the verification of the domain name ownership is not complete, and the value FILE indicates that the file verification method is used.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The domain name to be verified when you use the file verification method. You must connect to the DNS server of the domain name and create a file on the server. The file is specified by the **Uri** parameter.
	//
	// > This parameter is returned only when the value of the **Type** parameter is **domain\_verify** and the value of the **ValidateType** parameter is **FILE**. The value domain\_verify indicates that the verification of the domain name ownership is not complete, and the value FILE indicates that the file verification method is used.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The private key of the certificate in the PEM format. For more information about the PEM format and how to convert certificate formats, see [What formats are used for mainstream digital certificates?](~~42214~~)
	//
	// > This parameter is returned only when the value of the **Type** parameter is **certificate**. The value certificate indicates that the certificate is issued.
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The DNS record that you need to manage when you use the DNS verification method.
	//
	// > This parameter is returned only when the value of the **Type** parameter is **domain\_verify** and the value of the **ValidateType** parameter is **DNS**. The value domain\_verify indicates that the verification of the domain name ownership is not complete, and the value DNS indicates that the DNS verification method is used.
	RecordDomain *string `json:"RecordDomain,omitempty" xml:"RecordDomain,omitempty"`
	// The type of the DNS record that you need to add when you use the DNS verification method. Valid values:
	//
	// *   **TXT**
	// *   **CNAME**
	//
	// > This parameter is returned only when the value of the **Type** parameter is **domain\_verify** and the value of the **ValidateType** parameter is **DNS**. The value domain\_verify indicates that the verification of the domain name ownership is not complete.
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// You need to add a TXT record to the DNS records only when you use the DNS verification method.
	//
	// > This parameter is returned only when the value of the **Type** parameter is **domain\_verify** and the value of the **ValidateType** parameter is **DNS**. The value domain\_verify indicates that the verification of the domain name ownership is not complete, and the value DNS indicates that the DNS verification method is used.
	RecordValue *string `json:"RecordValue,omitempty" xml:"RecordValue,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the certificate application order. Valid values:
	//
	// *   **domain_verify**: **pending review**, which indicates that you have not completed the verification of the domain name ownership after you submit the certificate application.
	//
	//     > After you submit a certificate application, you must manually complete the verification of the domain name ownership. The CA reviews the certificate application only after the verification is complete. If you have not completed the verification of the domain name ownership, you can complete the verification based on the data returned by this operation.
	//
	// *   **process**: **being reviewed**, which indicates that the certificate application is being reviewed by the CA.
	//
	// *   **verify_fail**: **review failed**, which indicates that the certificate application failed to be reviewed.
	//
	//     > If a certificate application fails to be reviewed, the information that you specified in the certificate application may be incorrect. We recommend that you call the [DeleteCertificateRequest](~~455294~~) operation to delete the certificate application order and resubmit a certificate application. After the order is deleted, the quota that is consumed for the order is released.
	//
	// *   **certificate**: **issued**, which indicates that the certificate is issued.
	// *   **payed**: **pending application**, which indicates that you have not submitted a certificate application.
	// *   **unknow**: The status is **unknown**.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The file that you need to create on the DNS server when you use the file verification method. **The value of this parameter contains the file path and file name.**
	//
	// > This parameter is returned only when the value of the **Type** parameter is **domain\_verify** and the value of the **ValidateType** parameter is **FILE**. The value domain\_verify indicates that the verification of the domain name ownership is not complete, and the value FILE indicates that the file verification method is used.
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The verification method of the domain name ownership. Valid values:
	//
	// *   **DNS**: DNS verification. If you use this method, you must add a TXT record to the DNS records of the domain name in the management platform of the domain name.
	// *   **FILE**: file verification. If you use this method, you must create a specified file on the DNS server.
	//
	// > This parameter is returned only when the value of the **Type** parameter is **domain\_verify**. The value domain\_verify indicates that the verification of the domain name ownership is not complete.
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
}

func (s DescribeCertificateStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateStateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificateStateResponseBody) SetCertificate(v string) *DescribeCertificateStateResponseBody {
	s.Certificate = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetContent(v string) *DescribeCertificateStateResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetDomain(v string) *DescribeCertificateStateResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetPrivateKey(v string) *DescribeCertificateStateResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRecordDomain(v string) *DescribeCertificateStateResponseBody {
	s.RecordDomain = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRecordType(v string) *DescribeCertificateStateResponseBody {
	s.RecordType = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRecordValue(v string) *DescribeCertificateStateResponseBody {
	s.RecordValue = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRequestId(v string) *DescribeCertificateStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetType(v string) *DescribeCertificateStateResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetUri(v string) *DescribeCertificateStateResponseBody {
	s.Uri = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetValidateType(v string) *DescribeCertificateStateResponseBody {
	s.ValidateType = &v
	return s
}

type DescribeCertificateStateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertificateStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertificateStateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateStateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificateStateResponse) SetHeaders(v map[string]*string) *DescribeCertificateStateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificateStateResponse) SetStatusCode(v int32) *DescribeCertificateStateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertificateStateResponse) SetBody(v *DescribeCertificateStateResponseBody) *DescribeCertificateStateResponse {
	s.Body = v
	return s
}

type DescribeCloudResourceStatusRequest struct {
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s DescribeCloudResourceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudResourceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceStatusRequest) SetSecretId(v string) *DescribeCloudResourceStatusRequest {
	s.SecretId = &v
	return s
}

type DescribeCloudResourceStatusResponseBody struct {
	Data      []*DescribeCloudResourceStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudResourceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudResourceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceStatusResponseBody) SetData(v []*DescribeCloudResourceStatusResponseBodyData) *DescribeCloudResourceStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudResourceStatusResponseBody) SetRequestId(v string) *DescribeCloudResourceStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCloudResourceStatusResponseBodyData struct {
	CloudName    *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	TotalCount   *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudResourceStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudResourceStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceStatusResponseBodyData) SetCloudName(v string) *DescribeCloudResourceStatusResponseBodyData {
	s.CloudName = &v
	return s
}

func (s *DescribeCloudResourceStatusResponseBodyData) SetCloudProduct(v string) *DescribeCloudResourceStatusResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *DescribeCloudResourceStatusResponseBodyData) SetTotalCount(v int64) *DescribeCloudResourceStatusResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeCloudResourceStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudResourceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudResourceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudResourceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceStatusResponse) SetHeaders(v map[string]*string) *DescribeCloudResourceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudResourceStatusResponse) SetStatusCode(v int32) *DescribeCloudResourceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudResourceStatusResponse) SetBody(v *DescribeCloudResourceStatusResponseBody) *DescribeCloudResourceStatusResponse {
	s.Body = v
	return s
}

type DescribeDeploymentJobRequest struct {
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeDeploymentJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobRequest) SetJobId(v int64) *DescribeDeploymentJobRequest {
	s.JobId = &v
	return s
}

type DescribeDeploymentJobResponseBody struct {
	CasContacts  []*DescribeDeploymentJobResponseBodyCasContacts `json:"CasContacts,omitempty" xml:"CasContacts,omitempty" type:"Repeated"`
	CertDomain   *string                                         `json:"CertDomain,omitempty" xml:"CertDomain,omitempty"`
	CertType     *string                                         `json:"CertType,omitempty" xml:"CertType,omitempty"`
	Config       *string                                         `json:"Config,omitempty" xml:"Config,omitempty"`
	Del          *int32                                          `json:"Del,omitempty" xml:"Del,omitempty"`
	EndTime      *string                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GmtCreate    *string                                         `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string                                         `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id           *int64                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobType      *string                                         `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Name         *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductName  *string                                         `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rollback     *int32                                          `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
	ScheduleTime *string                                         `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	StartTime    *string                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId       *int64                                          `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeDeploymentJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeploymentJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobResponseBody) SetCasContacts(v []*DescribeDeploymentJobResponseBodyCasContacts) *DescribeDeploymentJobResponseBody {
	s.CasContacts = v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetCertDomain(v string) *DescribeDeploymentJobResponseBody {
	s.CertDomain = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetCertType(v string) *DescribeDeploymentJobResponseBody {
	s.CertType = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetConfig(v string) *DescribeDeploymentJobResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetDel(v int32) *DescribeDeploymentJobResponseBody {
	s.Del = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetEndTime(v string) *DescribeDeploymentJobResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetGmtCreate(v string) *DescribeDeploymentJobResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetGmtModified(v string) *DescribeDeploymentJobResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetId(v int64) *DescribeDeploymentJobResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetInstanceId(v string) *DescribeDeploymentJobResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetJobType(v string) *DescribeDeploymentJobResponseBody {
	s.JobType = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetName(v string) *DescribeDeploymentJobResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetProductName(v string) *DescribeDeploymentJobResponseBody {
	s.ProductName = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetRequestId(v string) *DescribeDeploymentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetRollback(v int32) *DescribeDeploymentJobResponseBody {
	s.Rollback = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetScheduleTime(v string) *DescribeDeploymentJobResponseBody {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetStartTime(v string) *DescribeDeploymentJobResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetStatus(v string) *DescribeDeploymentJobResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetUserId(v int64) *DescribeDeploymentJobResponseBody {
	s.UserId = &v
	return s
}

type DescribeDeploymentJobResponseBodyCasContacts struct {
	Email  *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDeploymentJobResponseBodyCasContacts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeploymentJobResponseBodyCasContacts) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) SetEmail(v string) *DescribeDeploymentJobResponseBodyCasContacts {
	s.Email = &v
	return s
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) SetId(v string) *DescribeDeploymentJobResponseBodyCasContacts {
	s.Id = &v
	return s
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) SetMobile(v string) *DescribeDeploymentJobResponseBodyCasContacts {
	s.Mobile = &v
	return s
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) SetName(v string) *DescribeDeploymentJobResponseBodyCasContacts {
	s.Name = &v
	return s
}

type DescribeDeploymentJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeploymentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeploymentJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeploymentJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobResponse) SetHeaders(v map[string]*string) *DescribeDeploymentJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeploymentJobResponse) SetStatusCode(v int32) *DescribeDeploymentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeploymentJobResponse) SetBody(v *DescribeDeploymentJobResponseBody) *DescribeDeploymentJobResponse {
	s.Body = v
	return s
}

type DescribeDeploymentJobStatusRequest struct {
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeDeploymentJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeploymentJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobStatusRequest) SetJobId(v int64) *DescribeDeploymentJobStatusRequest {
	s.JobId = &v
	return s
}

type DescribeDeploymentJobStatusResponseBody struct {
	BuyCount             *int32                                                       `json:"BuyCount,omitempty" xml:"BuyCount,omitempty"`
	CertCount            *int32                                                       `json:"CertCount,omitempty" xml:"CertCount,omitempty"`
	CostCount            *int32                                                       `json:"CostCount,omitempty" xml:"CostCount,omitempty"`
	FailedCount          *int32                                                       `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	MatchWorkerCount     *int32                                                       `json:"MatchWorkerCount,omitempty" xml:"MatchWorkerCount,omitempty"`
	ProductWorkerCount   []*DescribeDeploymentJobStatusResponseBodyProductWorkerCount `json:"ProductWorkerCount,omitempty" xml:"ProductWorkerCount,omitempty" type:"Repeated"`
	RequestId            *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceCount        *int32                                                       `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	RollbackCount        *int32                                                       `json:"RollbackCount,omitempty" xml:"RollbackCount,omitempty"`
	RollbackFailedCount  *int32                                                       `json:"RollbackFailedCount,omitempty" xml:"RollbackFailedCount,omitempty"`
	RollbackSuccessCount *int32                                                       `json:"RollbackSuccessCount,omitempty" xml:"RollbackSuccessCount,omitempty"`
	SuccessCount         *int32                                                       `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	UseCount             *int32                                                       `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
	WorkerCount          *int32                                                       `json:"WorkerCount,omitempty" xml:"WorkerCount,omitempty"`
}

func (s DescribeDeploymentJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeploymentJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobStatusResponseBody) SetBuyCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.BuyCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetCertCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.CertCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetCostCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.CostCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetFailedCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.FailedCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetMatchWorkerCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.MatchWorkerCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetProductWorkerCount(v []*DescribeDeploymentJobStatusResponseBodyProductWorkerCount) *DescribeDeploymentJobStatusResponseBody {
	s.ProductWorkerCount = v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetRequestId(v string) *DescribeDeploymentJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetResourceCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.ResourceCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetRollbackCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.RollbackCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetRollbackFailedCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.RollbackFailedCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetRollbackSuccessCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.RollbackSuccessCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetSuccessCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetUseCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.UseCount = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBody) SetWorkerCount(v int32) *DescribeDeploymentJobStatusResponseBody {
	s.WorkerCount = &v
	return s
}

type DescribeDeploymentJobStatusResponseBodyProductWorkerCount struct {
	Count       *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s DescribeDeploymentJobStatusResponseBodyProductWorkerCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeploymentJobStatusResponseBodyProductWorkerCount) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobStatusResponseBodyProductWorkerCount) SetCount(v int32) *DescribeDeploymentJobStatusResponseBodyProductWorkerCount {
	s.Count = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponseBodyProductWorkerCount) SetProductName(v string) *DescribeDeploymentJobStatusResponseBodyProductWorkerCount {
	s.ProductName = &v
	return s
}

type DescribeDeploymentJobStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeploymentJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeploymentJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeploymentJobStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobStatusResponse) SetHeaders(v map[string]*string) *DescribeDeploymentJobStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeploymentJobStatusResponse) SetStatusCode(v int32) *DescribeDeploymentJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeploymentJobStatusResponse) SetBody(v *DescribeDeploymentJobStatusResponseBody) *DescribeDeploymentJobStatusResponse {
	s.Body = v
	return s
}

type DescribePackageStateRequest struct {
	// The specifications of the certificate resource plan. Valid values:
	//
	// *   **digicert-free-1-free**: DigiCert single-domain DV certificate in 3 months free trial. This is the default value.
	// *   **symantec-free-1-free**: DigiCert single-domain DV certificate in 1 year free trial.
	// *   **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	// *   **symantec-ov-1-personal**: DigiCert single-domain organization validated (OV) certificate.
	// *   **symantec-ov-w-personal**: DigiCert wildcard OV certificate.
	// *   **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	// *   **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	// *   **geotrust-ov-1-personal**: GeoTrust single-domain OV certificate.
	// *   **geotrust-ov-w-personal**: GeoTrust wildcard OV certificate.
	// *   **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	// *   **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	// *   **globalsign-ov-1-personal**: GlobalSign single-domain OV certificate.
	// *   **globalsign-ov-w-advanced**: GlobalSign wildcard OV certificate.
	// *   **cfca-ov-1-personal**: China Financial Certification Authority (CFCA) single-domain OV certificate.
	// *   **cfca-ev-w-advanced**: CFCA wildcard OV certificate.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribePackageStateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageStateRequest) GoString() string {
	return s.String()
}

func (s *DescribePackageStateRequest) SetProductCode(v string) *DescribePackageStateRequest {
	s.ProductCode = &v
	return s
}

type DescribePackageStateResponseBody struct {
	// The number of issued certificates of the specified specifications.
	IssuedCount *int64 `json:"IssuedCount,omitempty" xml:"IssuedCount,omitempty"`
	// The specifications of the certificate. Valid values:
	//
	// *   **symantec-free-1-free**: DigiCert single-domain DV certificate in 3 months free trial.
	// *   **symantec-free-1-free**: DigiCert single-domain DV certificate in 1 year free trial.
	// *   **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	// *   **symantec-ov-1-personal**: DigiCert single-domain OV certificate.
	// *   **symantec-ov-w-personal**: DigiCert wildcard OV certificate.
	// *   **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	// *   **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	// *   **geotrust-ov-1-personal**: GeoTrust single-domain OV certificate.
	// *   **geotrust-ov-w-personal**: GeoTrust wildcard OV certificate.
	// *   **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	// *   **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	// *   **globalsign-ov-1-personal**: GlobalSign single-domain OV certificate.
	// *   **globalsign-ov-w-advanced**: GlobalSign wildcard OV certificate.
	// *   **cfca-ov-1-personal**: CFCA single-domain OV certificate.
	// *   **cfca-ev-w-advanced**: CFCA wildcard OV certificate.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of purchased certificate resource plans of the specified specifications.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of certificate applications that you submitted for certificates of the specified specifications.
	//
	// > A successful call of the [CreateCertificateForPackageRequest](~~455296~~), [CreateCertificateRequest](~~455292~~), or [CreateCertificateWithCsrRequest](~~455801~~) operation is counted as one a certificate application, regardless of whether the certificate is issued.
	UsedCount *int64 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribePackageStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageStateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackageStateResponseBody) SetIssuedCount(v int64) *DescribePackageStateResponseBody {
	s.IssuedCount = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetProductCode(v string) *DescribePackageStateResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetRequestId(v string) *DescribePackageStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetTotalCount(v int64) *DescribePackageStateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetUsedCount(v int64) *DescribePackageStateResponseBody {
	s.UsedCount = &v
	return s
}

type DescribePackageStateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackageStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePackageStateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageStateResponse) GoString() string {
	return s.String()
}

func (s *DescribePackageStateResponse) SetHeaders(v map[string]*string) *DescribePackageStateResponse {
	s.Headers = v
	return s
}

func (s *DescribePackageStateResponse) SetStatusCode(v int32) *DescribePackageStateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackageStateResponse) SetBody(v *DescribePackageStateResponseBody) *DescribePackageStateResponse {
	s.Body = v
	return s
}

type EncryptRequest struct {
	Algorithm      *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	MessageType    *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Plaintext      *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s EncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s EncryptRequest) GoString() string {
	return s.String()
}

func (s *EncryptRequest) SetAlgorithm(v string) *EncryptRequest {
	s.Algorithm = &v
	return s
}

func (s *EncryptRequest) SetCertIdentifier(v string) *EncryptRequest {
	s.CertIdentifier = &v
	return s
}

func (s *EncryptRequest) SetMessageType(v string) *EncryptRequest {
	s.MessageType = &v
	return s
}

func (s *EncryptRequest) SetPlaintext(v string) *EncryptRequest {
	s.Plaintext = &v
	return s
}

type EncryptResponseBody struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EncryptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EncryptResponseBody) GoString() string {
	return s.String()
}

func (s *EncryptResponseBody) SetCertIdentifier(v string) *EncryptResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *EncryptResponseBody) SetCiphertextBlob(v string) *EncryptResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *EncryptResponseBody) SetRequestId(v string) *EncryptResponseBody {
	s.RequestId = &v
	return s
}

type EncryptResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s EncryptResponse) GoString() string {
	return s.String()
}

func (s *EncryptResponse) SetHeaders(v map[string]*string) *EncryptResponse {
	s.Headers = v
	return s
}

func (s *EncryptResponse) SetStatusCode(v int32) *EncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *EncryptResponse) SetBody(v *EncryptResponseBody) *EncryptResponse {
	s.Body = v
	return s
}

type GetCertWarehouseQuotaResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalQuota *int64  `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	UseCount   *int64  `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
}

func (s GetCertWarehouseQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCertWarehouseQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertWarehouseQuotaResponseBody) SetRequestId(v string) *GetCertWarehouseQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertWarehouseQuotaResponseBody) SetTotalQuota(v int64) *GetCertWarehouseQuotaResponseBody {
	s.TotalQuota = &v
	return s
}

func (s *GetCertWarehouseQuotaResponseBody) SetUseCount(v int64) *GetCertWarehouseQuotaResponseBody {
	s.UseCount = &v
	return s
}

type GetCertWarehouseQuotaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCertWarehouseQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCertWarehouseQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCertWarehouseQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetCertWarehouseQuotaResponse) SetHeaders(v map[string]*string) *GetCertWarehouseQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetCertWarehouseQuotaResponse) SetStatusCode(v int32) *GetCertWarehouseQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCertWarehouseQuotaResponse) SetBody(v *GetCertWarehouseQuotaResponseBody) *GetCertWarehouseQuotaResponse {
	s.Body = v
	return s
}

type GetCsrDetailRequest struct {
	// CSR ID。
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
}

func (s GetCsrDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCsrDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCsrDetailRequest) SetCsrId(v int64) *GetCsrDetailRequest {
	s.CsrId = &v
	return s
}

type GetCsrDetailResponseBody struct {
	Csr       *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCsrDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCsrDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCsrDetailResponseBody) SetCsr(v string) *GetCsrDetailResponseBody {
	s.Csr = &v
	return s
}

func (s *GetCsrDetailResponseBody) SetRequestId(v string) *GetCsrDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetCsrDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCsrDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCsrDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCsrDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCsrDetailResponse) SetHeaders(v map[string]*string) *GetCsrDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCsrDetailResponse) SetStatusCode(v int32) *GetCsrDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCsrDetailResponse) SetBody(v *GetCsrDetailResponseBody) *GetCsrDetailResponse {
	s.Body = v
	return s
}

type GetUserCertificateDetailRequest struct {
	// If true, the Cert, Key, EncryptCert, EncryptPrivateKey, SignCert, SignPrivateKey will return null, default is false.
	CertFilter *bool `json:"CertFilter,omitempty" xml:"CertFilter,omitempty"`
	// The ID of the certificate.
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
}

func (s GetUserCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetUserCertificateDetailRequest) SetCertFilter(v bool) *GetUserCertificateDetailRequest {
	s.CertFilter = &v
	return s
}

func (s *GetUserCertificateDetailRequest) SetCertId(v int64) *GetUserCertificateDetailRequest {
	s.CertId = &v
	return s
}

type GetUserCertificateDetailResponseBody struct {
	// The algorithm.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// Indicates whether the certificate was purchased from Alibaba Cloud. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	BuyInAliyun *bool `json:"BuyInAliyun,omitempty" xml:"BuyInAliyun,omitempty"`
	// The content of the certificate.
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The city of the company or organization to which the certificate purchaser belongs.
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The parent domain name that is bound to the certificate.
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// The country or region of the company or organization to which the certificate purchaser belongs.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the encryption certificate in PEM format.
	EncryptCert *string `json:"EncryptCert,omitempty" xml:"EncryptCert,omitempty"`
	// The private key of the encryption certificate in the PEM format.
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	// The expiration date of the certificate.
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Indicates whether the certificate has expired. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The fingerprint of the certificate.
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The ID of the certificate.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The private key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the certificate.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the certificate application order.
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The name of the company or organization to which the certificate purchaser belongs.
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The province of the company or organization to which the certificate purchaser belongs.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the certificate belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// All domain names that are bound to the certificate.
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The content of the signing certificate in the PEM format.
	SignCert *string `json:"SignCert,omitempty" xml:"SignCert,omitempty"`
	// The private key of the signing certificate in the PEM format.
	SignPrivateKey *string `json:"SignPrivateKey,omitempty" xml:"SignPrivateKey,omitempty"`
	// The issuance date of the certificate.
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetUserCertificateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserCertificateDetailResponseBody) SetAlgorithm(v string) *GetUserCertificateDetailResponseBody {
	s.Algorithm = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetBuyInAliyun(v bool) *GetUserCertificateDetailResponseBody {
	s.BuyInAliyun = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCert(v string) *GetUserCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCity(v string) *GetUserCertificateDetailResponseBody {
	s.City = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCommon(v string) *GetUserCertificateDetailResponseBody {
	s.Common = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCountry(v string) *GetUserCertificateDetailResponseBody {
	s.Country = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetEncryptCert(v string) *GetUserCertificateDetailResponseBody {
	s.EncryptCert = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetEncryptPrivateKey(v string) *GetUserCertificateDetailResponseBody {
	s.EncryptPrivateKey = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetEndDate(v string) *GetUserCertificateDetailResponseBody {
	s.EndDate = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetExpired(v bool) *GetUserCertificateDetailResponseBody {
	s.Expired = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetFingerprint(v string) *GetUserCertificateDetailResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetId(v int64) *GetUserCertificateDetailResponseBody {
	s.Id = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetIssuer(v string) *GetUserCertificateDetailResponseBody {
	s.Issuer = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetKey(v string) *GetUserCertificateDetailResponseBody {
	s.Key = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetName(v string) *GetUserCertificateDetailResponseBody {
	s.Name = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetOrderId(v int64) *GetUserCertificateDetailResponseBody {
	s.OrderId = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetOrgName(v string) *GetUserCertificateDetailResponseBody {
	s.OrgName = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetProvince(v string) *GetUserCertificateDetailResponseBody {
	s.Province = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetRequestId(v string) *GetUserCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetResourceGroupId(v string) *GetUserCertificateDetailResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetSans(v string) *GetUserCertificateDetailResponseBody {
	s.Sans = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetSignCert(v string) *GetUserCertificateDetailResponseBody {
	s.SignCert = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetSignPrivateKey(v string) *GetUserCertificateDetailResponseBody {
	s.SignPrivateKey = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetStartDate(v string) *GetUserCertificateDetailResponseBody {
	s.StartDate = &v
	return s
}

type GetUserCertificateDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserCertificateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *GetUserCertificateDetailResponse) SetHeaders(v map[string]*string) *GetUserCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *GetUserCertificateDetailResponse) SetStatusCode(v int32) *GetUserCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserCertificateDetailResponse) SetBody(v *GetUserCertificateDetailResponseBody) *GetUserCertificateDetailResponse {
	s.Body = v
	return s
}

type ListCertRequest struct {
	// The type of the certificate.
	//
	// *   **CA**: the CA certificate.
	// *   **CERT**: a issued certificate.
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The number of the page to return. Default value: 1.
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword that is used for queries. The value can be a name, domain name, or subject alternative name (SAN) attribute. Fuzzy match is supported.
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The number of entries to return on each page. Default value: 50.
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The source of the certificate. Valid values:
	//
	// *   **upload**: uploaded certificate
	// *   **aliyun**: Alibaba Cloud certificate
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The status of the certificate. Valid values:
	//
	// *   **ISSUE**: issued
	// *   **REVOKE**: revoked
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the certificate repository. You can call the ListCertWarehouse API operation to query the IDs of certificate repositories.
	WarehouseId *int64 `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
}

func (s ListCertRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCertRequest) GoString() string {
	return s.String()
}

func (s *ListCertRequest) SetCertType(v string) *ListCertRequest {
	s.CertType = &v
	return s
}

func (s *ListCertRequest) SetCurrentPage(v int64) *ListCertRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCertRequest) SetKeyWord(v string) *ListCertRequest {
	s.KeyWord = &v
	return s
}

func (s *ListCertRequest) SetShowSize(v int64) *ListCertRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCertRequest) SetSourceType(v string) *ListCertRequest {
	s.SourceType = &v
	return s
}

func (s *ListCertRequest) SetStatus(v string) *ListCertRequest {
	s.Status = &v
	return s
}

func (s *ListCertRequest) SetWarehouseId(v int64) *ListCertRequest {
	s.WarehouseId = &v
	return s
}

type ListCertResponseBody struct {
	// The certificates.
	CertList []*ListCertResponseBodyCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	// The page number of the returned page. Default value: 1.
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned per page. Default value: 50.
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCertResponseBody) GoString() string {
	return s.String()
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

type ListCertResponseBodyCertList struct {
	// The expiration time of the certificate. The value is a UNIX timestamp. Unit: milliseconds.
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The issuance time of the certificate. The value is a UNIX timestamp. Unit: milliseconds.
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the certificate.
	//
	// *   **CA**: the CA certificate.
	// *   **CERT**: a issued certificate.
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The domain name.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// Indicates whether the certificate contains a private key. Valid values:
	//
	// *   **true**
	// *   **false**
	ExistPrivateKey *bool `json:"ExistPrivateKey,omitempty" xml:"ExistPrivateKey,omitempty"`
	// The unique identifier of the certificate.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The issuer of the certificate.
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// All domain names that are bound to the certificate. Multiple domain names are separated by commas (,).
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The source of the certificate. Valid values:
	//
	// *   **upload**: uploaded certificate
	// *   **aliyun**: Alibaba Cloud certificate
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The status of the certificate. Valid values:
	//
	// *   **ISSUE**: issued
	// *   **REVOKE**: revoked
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the certificate application repository.
	WhId *int64 `json:"WhId,omitempty" xml:"WhId,omitempty"`
	// The instance ID of the certificate application repository.
	WhInstanceId *string `json:"WhInstanceId,omitempty" xml:"WhInstanceId,omitempty"`
}

func (s ListCertResponseBodyCertList) String() string {
	return tea.Prettify(s)
}

func (s ListCertResponseBodyCertList) GoString() string {
	return s.String()
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

type ListCertWarehouseRequest struct {
	CurrentPage *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShowSize    *int64  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCertWarehouseRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCertWarehouseRequest) GoString() string {
	return s.String()
}

func (s *ListCertWarehouseRequest) SetCurrentPage(v int64) *ListCertWarehouseRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCertWarehouseRequest) SetInstanceId(v string) *ListCertWarehouseRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCertWarehouseRequest) SetName(v string) *ListCertWarehouseRequest {
	s.Name = &v
	return s
}

func (s *ListCertWarehouseRequest) SetShowSize(v int64) *ListCertWarehouseRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCertWarehouseRequest) SetType(v string) *ListCertWarehouseRequest {
	s.Type = &v
	return s
}

type ListCertWarehouseResponseBody struct {
	CertWarehouseList []*ListCertWarehouseResponseBodyCertWarehouseList `json:"CertWarehouseList,omitempty" xml:"CertWarehouseList,omitempty" type:"Repeated"`
	CurrentPage       *int64                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId         *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize          *int64                                            `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount        *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertWarehouseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCertWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertWarehouseResponseBody) SetCertWarehouseList(v []*ListCertWarehouseResponseBodyCertWarehouseList) *ListCertWarehouseResponseBody {
	s.CertWarehouseList = v
	return s
}

func (s *ListCertWarehouseResponseBody) SetCurrentPage(v int64) *ListCertWarehouseResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCertWarehouseResponseBody) SetRequestId(v string) *ListCertWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertWarehouseResponseBody) SetShowSize(v int64) *ListCertWarehouseResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCertWarehouseResponseBody) SetTotalCount(v int64) *ListCertWarehouseResponseBody {
	s.TotalCount = &v
	return s
}

type ListCertWarehouseResponseBodyCertWarehouseList struct {
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsExpired     *bool   `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PcaInstanceId *string `json:"PcaInstanceId,omitempty" xml:"PcaInstanceId,omitempty"`
	Qps           *int64  `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WhId          *int64  `json:"WhId,omitempty" xml:"WhId,omitempty"`
}

func (s ListCertWarehouseResponseBodyCertWarehouseList) String() string {
	return tea.Prettify(s)
}

func (s ListCertWarehouseResponseBodyCertWarehouseList) GoString() string {
	return s.String()
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetEndTime(v int64) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.EndTime = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetInstanceId(v string) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.InstanceId = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetIsExpired(v bool) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.IsExpired = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetName(v string) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.Name = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetPcaInstanceId(v string) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.PcaInstanceId = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetQps(v int64) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.Qps = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetType(v string) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.Type = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetWhId(v int64) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.WhId = &v
	return s
}

type ListCertWarehouseResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCertWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCertWarehouseResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCertWarehouseResponse) GoString() string {
	return s.String()
}

func (s *ListCertWarehouseResponse) SetHeaders(v map[string]*string) *ListCertWarehouseResponse {
	s.Headers = v
	return s
}

func (s *ListCertWarehouseResponse) SetStatusCode(v int32) *ListCertWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCertWarehouseResponse) SetBody(v *ListCertWarehouseResponseBody) *ListCertWarehouseResponse {
	s.Body = v
	return s
}

type ListCloudAccessRequest struct {
	CloudName   *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	SecretId    *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	ShowSize    *int32  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListCloudAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudAccessRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAccessRequest) SetCloudName(v string) *ListCloudAccessRequest {
	s.CloudName = &v
	return s
}

func (s *ListCloudAccessRequest) SetCurrentPage(v int32) *ListCloudAccessRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudAccessRequest) SetSecretId(v string) *ListCloudAccessRequest {
	s.SecretId = &v
	return s
}

func (s *ListCloudAccessRequest) SetShowSize(v int32) *ListCloudAccessRequest {
	s.ShowSize = &v
	return s
}

type ListCloudAccessResponseBody struct {
	CloudAccessList []*ListCloudAccessResponseBodyCloudAccessList `json:"CloudAccessList,omitempty" xml:"CloudAccessList,omitempty" type:"Repeated"`
	CurrentPage     *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize        *int32                                        `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount      *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCloudAccessResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAccessResponseBody) SetCloudAccessList(v []*ListCloudAccessResponseBodyCloudAccessList) *ListCloudAccessResponseBody {
	s.CloudAccessList = v
	return s
}

func (s *ListCloudAccessResponseBody) SetCurrentPage(v int32) *ListCloudAccessResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudAccessResponseBody) SetRequestId(v string) *ListCloudAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAccessResponseBody) SetShowSize(v int32) *ListCloudAccessResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCloudAccessResponseBody) SetTotalCount(v int32) *ListCloudAccessResponseBody {
	s.TotalCount = &v
	return s
}

type ListCloudAccessResponseBodyCloudAccessList struct {
	AccessId      *int64  `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	CloudName     *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	SecretId      *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s ListCloudAccessResponseBodyCloudAccessList) String() string {
	return tea.Prettify(s)
}

func (s ListCloudAccessResponseBodyCloudAccessList) GoString() string {
	return s.String()
}

func (s *ListCloudAccessResponseBodyCloudAccessList) SetAccessId(v int64) *ListCloudAccessResponseBodyCloudAccessList {
	s.AccessId = &v
	return s
}

func (s *ListCloudAccessResponseBodyCloudAccessList) SetCloudName(v string) *ListCloudAccessResponseBodyCloudAccessList {
	s.CloudName = &v
	return s
}

func (s *ListCloudAccessResponseBodyCloudAccessList) SetSecretId(v string) *ListCloudAccessResponseBodyCloudAccessList {
	s.SecretId = &v
	return s
}

func (s *ListCloudAccessResponseBodyCloudAccessList) SetServiceStatus(v string) *ListCloudAccessResponseBodyCloudAccessList {
	s.ServiceStatus = &v
	return s
}

type ListCloudAccessResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCloudAccessResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAccessResponse) SetHeaders(v map[string]*string) *ListCloudAccessResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAccessResponse) SetStatusCode(v int32) *ListCloudAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAccessResponse) SetBody(v *ListCloudAccessResponseBody) *ListCloudAccessResponse {
	s.Body = v
	return s
}

type ListCloudResourcesRequest struct {
	CloudName    *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	SecretId     *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	ShowSize     *int32  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListCloudResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesRequest) SetCloudName(v string) *ListCloudResourcesRequest {
	s.CloudName = &v
	return s
}

func (s *ListCloudResourcesRequest) SetCloudProduct(v string) *ListCloudResourcesRequest {
	s.CloudProduct = &v
	return s
}

func (s *ListCloudResourcesRequest) SetCurrentPage(v int32) *ListCloudResourcesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudResourcesRequest) SetKeyword(v string) *ListCloudResourcesRequest {
	s.Keyword = &v
	return s
}

func (s *ListCloudResourcesRequest) SetSecretId(v string) *ListCloudResourcesRequest {
	s.SecretId = &v
	return s
}

func (s *ListCloudResourcesRequest) SetShowSize(v int32) *ListCloudResourcesRequest {
	s.ShowSize = &v
	return s
}

type ListCloudResourcesResponseBody struct {
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*ListCloudResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize    *int32                                `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	Total       *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCloudResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesResponseBody) SetCurrentPage(v int32) *ListCloudResourcesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudResourcesResponseBody) SetData(v []*ListCloudResourcesResponseBodyData) *ListCloudResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListCloudResourcesResponseBody) SetRequestId(v string) *ListCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudResourcesResponseBody) SetShowSize(v int32) *ListCloudResourcesResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCloudResourcesResponseBody) SetTotal(v int64) *ListCloudResourcesResponseBody {
	s.Total = &v
	return s
}

type ListCloudResourcesResponseBodyData struct {
	CertEndTime     *string `json:"CertEndTime,omitempty" xml:"CertEndTime,omitempty"`
	CertId          *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName        *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertStartTime   *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	CloudAccessId   *string `json:"CloudAccessId,omitempty" xml:"CloudAccessId,omitempty"`
	CloudName       *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	CloudProduct    *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	CloudRegion     *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	DefaultResource *int32  `json:"DefaultResource,omitempty" xml:"DefaultResource,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EnableHttps     *int32  `json:"EnableHttps,omitempty" xml:"EnableHttps,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ListenerId      *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	ListenerPort    *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UseSsl          *int32  `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	UserId          *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListCloudResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCloudResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesResponseBodyData) SetCertEndTime(v string) *ListCloudResourcesResponseBodyData {
	s.CertEndTime = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCertId(v int64) *ListCloudResourcesResponseBodyData {
	s.CertId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCertName(v string) *ListCloudResourcesResponseBodyData {
	s.CertName = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCertStartTime(v string) *ListCloudResourcesResponseBodyData {
	s.CertStartTime = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCloudAccessId(v string) *ListCloudResourcesResponseBodyData {
	s.CloudAccessId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCloudName(v string) *ListCloudResourcesResponseBodyData {
	s.CloudName = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCloudProduct(v string) *ListCloudResourcesResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCloudRegion(v string) *ListCloudResourcesResponseBodyData {
	s.CloudRegion = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetDefaultResource(v int32) *ListCloudResourcesResponseBodyData {
	s.DefaultResource = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetDomain(v string) *ListCloudResourcesResponseBodyData {
	s.Domain = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetEnableHttps(v int32) *ListCloudResourcesResponseBodyData {
	s.EnableHttps = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetGmtCreate(v string) *ListCloudResourcesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetGmtModified(v string) *ListCloudResourcesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetId(v int64) *ListCloudResourcesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetInstanceId(v string) *ListCloudResourcesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetListenerId(v string) *ListCloudResourcesResponseBodyData {
	s.ListenerId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetListenerPort(v string) *ListCloudResourcesResponseBodyData {
	s.ListenerPort = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetRegionId(v string) *ListCloudResourcesResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetStatus(v string) *ListCloudResourcesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetUseSsl(v int32) *ListCloudResourcesResponseBodyData {
	s.UseSsl = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetUserId(v int64) *ListCloudResourcesResponseBodyData {
	s.UserId = &v
	return s
}

type ListCloudResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesResponse) SetHeaders(v map[string]*string) *ListCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudResourcesResponse) SetStatusCode(v int32) *ListCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudResourcesResponse) SetBody(v *ListCloudResourcesResponseBody) *ListCloudResourcesResponse {
	s.Body = v
	return s
}

type ListContactRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	ShowSize    *int32  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListContactRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContactRequest) GoString() string {
	return s.String()
}

func (s *ListContactRequest) SetCurrentPage(v int32) *ListContactRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListContactRequest) SetKeyword(v string) *ListContactRequest {
	s.Keyword = &v
	return s
}

func (s *ListContactRequest) SetShowSize(v int32) *ListContactRequest {
	s.ShowSize = &v
	return s
}

type ListContactResponseBody struct {
	ContactList []*ListContactResponseBodyContactList `json:"ContactList,omitempty" xml:"ContactList,omitempty" type:"Repeated"`
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Keyword     *string                               `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize    *int32                                `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount  *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContactResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactResponseBody) SetContactList(v []*ListContactResponseBodyContactList) *ListContactResponseBody {
	s.ContactList = v
	return s
}

func (s *ListContactResponseBody) SetCurrentPage(v int32) *ListContactResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListContactResponseBody) SetKeyword(v string) *ListContactResponseBody {
	s.Keyword = &v
	return s
}

func (s *ListContactResponseBody) SetRequestId(v string) *ListContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContactResponseBody) SetShowSize(v int32) *ListContactResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListContactResponseBody) SetTotalCount(v int64) *ListContactResponseBody {
	s.TotalCount = &v
	return s
}

type ListContactResponseBodyContactList struct {
	ContactId    *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EmailStatus  *int32  `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	Mobile       *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	MobileStatus *int32  `json:"MobileStatus,omitempty" xml:"MobileStatus,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Webhooks     *string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
}

func (s ListContactResponseBodyContactList) String() string {
	return tea.Prettify(s)
}

func (s ListContactResponseBodyContactList) GoString() string {
	return s.String()
}

func (s *ListContactResponseBodyContactList) SetContactId(v int64) *ListContactResponseBodyContactList {
	s.ContactId = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetEmail(v string) *ListContactResponseBodyContactList {
	s.Email = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetEmailStatus(v int32) *ListContactResponseBodyContactList {
	s.EmailStatus = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetMobile(v string) *ListContactResponseBodyContactList {
	s.Mobile = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetMobileStatus(v int32) *ListContactResponseBodyContactList {
	s.MobileStatus = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetName(v string) *ListContactResponseBodyContactList {
	s.Name = &v
	return s
}

func (s *ListContactResponseBodyContactList) SetWebhooks(v string) *ListContactResponseBodyContactList {
	s.Webhooks = &v
	return s
}

type ListContactResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContactResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContactResponse) GoString() string {
	return s.String()
}

func (s *ListContactResponse) SetHeaders(v map[string]*string) *ListContactResponse {
	s.Headers = v
	return s
}

func (s *ListContactResponse) SetStatusCode(v int32) *ListContactResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContactResponse) SetBody(v *ListContactResponseBody) *ListContactResponse {
	s.Body = v
	return s
}

type ListCsrRequest struct {
	Algorithm   *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CurrentPage *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	KeyWord     *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	ShowSize    *int64  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListCsrRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCsrRequest) GoString() string {
	return s.String()
}

func (s *ListCsrRequest) SetAlgorithm(v string) *ListCsrRequest {
	s.Algorithm = &v
	return s
}

func (s *ListCsrRequest) SetCurrentPage(v int64) *ListCsrRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCsrRequest) SetKeyWord(v string) *ListCsrRequest {
	s.KeyWord = &v
	return s
}

func (s *ListCsrRequest) SetShowSize(v int64) *ListCsrRequest {
	s.ShowSize = &v
	return s
}

type ListCsrResponseBody struct {
	CsrList     []*ListCsrResponseBodyCsrList `json:"CsrList,omitempty" xml:"CsrList,omitempty" type:"Repeated"`
	CurrentPage *int64                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize    *int64                        `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount  *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCsrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCsrResponseBody) GoString() string {
	return s.String()
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

type ListCsrResponseBodyCsrList struct {
	Algorithm   *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CommonName  *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CorpName    *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// CSR ID。
	CsrId         *int64  `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	Department    *string `json:"Department,omitempty" xml:"Department,omitempty"`
	HasPrivateKey *bool   `json:"HasPrivateKey,omitempty" xml:"HasPrivateKey,omitempty"`
	KeySha2       *string `json:"KeySha2,omitempty" xml:"KeySha2,omitempty"`
	KeySize       *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Locality      *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Province      *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Sans          *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
}

func (s ListCsrResponseBodyCsrList) String() string {
	return tea.Prettify(s)
}

func (s ListCsrResponseBodyCsrList) GoString() string {
	return s.String()
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

type ListCsrResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCsrResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCsrResponse) GoString() string {
	return s.String()
}

func (s *ListCsrResponse) SetHeaders(v map[string]*string) *ListCsrResponse {
	s.Headers = v
	return s
}

func (s *ListCsrResponse) SetStatusCode(v int32) *ListCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCsrResponse) SetBody(v *ListCsrResponseBody) *ListCsrResponse {
	s.Body = v
	return s
}

type ListDeploymentJobRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	JobType     *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	ShowSize    *int32  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobRequest) SetCurrentPage(v int32) *ListDeploymentJobRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDeploymentJobRequest) SetJobType(v string) *ListDeploymentJobRequest {
	s.JobType = &v
	return s
}

func (s *ListDeploymentJobRequest) SetShowSize(v int32) *ListDeploymentJobRequest {
	s.ShowSize = &v
	return s
}

func (s *ListDeploymentJobRequest) SetStatus(v string) *ListDeploymentJobRequest {
	s.Status = &v
	return s
}

type ListDeploymentJobResponseBody struct {
	CurrentPage *int32                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*ListDeploymentJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize    *int32                               `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	Total       *int64                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeploymentJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResponseBody) SetCurrentPage(v int32) *ListDeploymentJobResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListDeploymentJobResponseBody) SetData(v []*ListDeploymentJobResponseBodyData) *ListDeploymentJobResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentJobResponseBody) SetRequestId(v string) *ListDeploymentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentJobResponseBody) SetShowSize(v int32) *ListDeploymentJobResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListDeploymentJobResponseBody) SetTotal(v int64) *ListDeploymentJobResponseBody {
	s.Total = &v
	return s
}

type ListDeploymentJobResponseBodyData struct {
	CertDomain   *string `json:"CertDomain,omitempty" xml:"CertDomain,omitempty"`
	CertType     *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	Del          *int32  `json:"Del,omitempty" xml:"Del,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobType      *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductName  *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Rollback     *int32  `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDeploymentJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResponseBodyData) SetCertDomain(v string) *ListDeploymentJobResponseBodyData {
	s.CertDomain = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetCertType(v string) *ListDeploymentJobResponseBodyData {
	s.CertType = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetDel(v int32) *ListDeploymentJobResponseBodyData {
	s.Del = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetEndTime(v string) *ListDeploymentJobResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetGmtCreate(v string) *ListDeploymentJobResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetGmtModified(v string) *ListDeploymentJobResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetId(v int64) *ListDeploymentJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetInstanceId(v string) *ListDeploymentJobResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetJobType(v string) *ListDeploymentJobResponseBodyData {
	s.JobType = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetName(v string) *ListDeploymentJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetProductName(v string) *ListDeploymentJobResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetRollback(v int32) *ListDeploymentJobResponseBodyData {
	s.Rollback = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetScheduleTime(v string) *ListDeploymentJobResponseBodyData {
	s.ScheduleTime = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetStartTime(v string) *ListDeploymentJobResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetStatus(v string) *ListDeploymentJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetUserId(v int64) *ListDeploymentJobResponseBodyData {
	s.UserId = &v
	return s
}

type ListDeploymentJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResponse) SetHeaders(v map[string]*string) *ListDeploymentJobResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentJobResponse) SetStatusCode(v int32) *ListDeploymentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentJobResponse) SetBody(v *ListDeploymentJobResponseBody) *ListDeploymentJobResponse {
	s.Body = v
	return s
}

type ListDeploymentJobCertRequest struct {
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s ListDeploymentJobCertRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobCertRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobCertRequest) SetJobId(v int64) *ListDeploymentJobCertRequest {
	s.JobId = &v
	return s
}

type ListDeploymentJobCertResponseBody struct {
	Data      []*ListDeploymentJobCertResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeploymentJobCertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobCertResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobCertResponseBody) SetData(v []*ListDeploymentJobCertResponseBodyData) *ListDeploymentJobCertResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentJobCertResponseBody) SetRequestId(v string) *ListDeploymentJobCertResponseBody {
	s.RequestId = &v
	return s
}

type ListDeploymentJobCertResponseBodyData struct {
	Algorithm      *string   `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CertId         *int64    `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertInstanceId *string   `json:"CertInstanceId,omitempty" xml:"CertInstanceId,omitempty"`
	CertName       *string   `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertOrderType  *string   `json:"CertOrderType,omitempty" xml:"CertOrderType,omitempty"`
	CertType       *string   `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CommonName     *string   `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	IsTrustee      *bool     `json:"IsTrustee,omitempty" xml:"IsTrustee,omitempty"`
	Month          *int32    `json:"Month,omitempty" xml:"Month,omitempty"`
	NotAfterTime   *int64    `json:"NotAfterTime,omitempty" xml:"NotAfterTime,omitempty"`
	NotBeforeTime  *int64    `json:"NotBeforeTime,omitempty" xml:"NotBeforeTime,omitempty"`
	OrderId        *int64    `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Sans           []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
	StatusCode     *string   `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListDeploymentJobCertResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobCertResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobCertResponseBodyData) SetAlgorithm(v string) *ListDeploymentJobCertResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCertId(v int64) *ListDeploymentJobCertResponseBodyData {
	s.CertId = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCertInstanceId(v string) *ListDeploymentJobCertResponseBodyData {
	s.CertInstanceId = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCertName(v string) *ListDeploymentJobCertResponseBodyData {
	s.CertName = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCertOrderType(v string) *ListDeploymentJobCertResponseBodyData {
	s.CertOrderType = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCertType(v string) *ListDeploymentJobCertResponseBodyData {
	s.CertType = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCommonName(v string) *ListDeploymentJobCertResponseBodyData {
	s.CommonName = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetIsTrustee(v bool) *ListDeploymentJobCertResponseBodyData {
	s.IsTrustee = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetMonth(v int32) *ListDeploymentJobCertResponseBodyData {
	s.Month = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetNotAfterTime(v int64) *ListDeploymentJobCertResponseBodyData {
	s.NotAfterTime = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetNotBeforeTime(v int64) *ListDeploymentJobCertResponseBodyData {
	s.NotBeforeTime = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetOrderId(v int64) *ListDeploymentJobCertResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetSans(v []*string) *ListDeploymentJobCertResponseBodyData {
	s.Sans = v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetStatusCode(v string) *ListDeploymentJobCertResponseBodyData {
	s.StatusCode = &v
	return s
}

type ListDeploymentJobCertResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentJobCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentJobCertResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobCertResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobCertResponse) SetHeaders(v map[string]*string) *ListDeploymentJobCertResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentJobCertResponse) SetStatusCode(v int32) *ListDeploymentJobCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentJobCertResponse) SetBody(v *ListDeploymentJobCertResponseBody) *ListDeploymentJobCertResponse {
	s.Body = v
	return s
}

type ListDeploymentJobResourceRequest struct {
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s ListDeploymentJobResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobResourceRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResourceRequest) SetJobId(v int64) *ListDeploymentJobResourceRequest {
	s.JobId = &v
	return s
}

type ListDeploymentJobResourceResponseBody struct {
	Data      []*ListDeploymentJobResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeploymentJobResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResourceResponseBody) SetData(v []*ListDeploymentJobResourceResponseBodyData) *ListDeploymentJobResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentJobResourceResponseBody) SetRequestId(v string) *ListDeploymentJobResourceResponseBody {
	s.RequestId = &v
	return s
}

type ListDeploymentJobResourceResponseBodyData struct {
	CertEndTime     *string `json:"CertEndTime,omitempty" xml:"CertEndTime,omitempty"`
	CertId          *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName        *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertStartTime   *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	CloudAccessId   *string `json:"CloudAccessId,omitempty" xml:"CloudAccessId,omitempty"`
	CloudName       *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	CloudProduct    *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	CloudRegion     *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	DefaultResource *int32  `json:"DefaultResource,omitempty" xml:"DefaultResource,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EnableHttps     *int32  `json:"EnableHttps,omitempty" xml:"EnableHttps,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ListenerId      *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	ListenerPort    *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UseSsl          *int32  `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	UserId          *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDeploymentJobResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCertEndTime(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CertEndTime = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCertId(v int64) *ListDeploymentJobResourceResponseBodyData {
	s.CertId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCertName(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CertName = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCertStartTime(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CertStartTime = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCloudAccessId(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CloudAccessId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCloudName(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CloudName = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCloudProduct(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetCloudRegion(v string) *ListDeploymentJobResourceResponseBodyData {
	s.CloudRegion = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetDefaultResource(v int32) *ListDeploymentJobResourceResponseBodyData {
	s.DefaultResource = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetDomain(v string) *ListDeploymentJobResourceResponseBodyData {
	s.Domain = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetEnableHttps(v int32) *ListDeploymentJobResourceResponseBodyData {
	s.EnableHttps = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetGmtCreate(v string) *ListDeploymentJobResourceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetGmtModified(v string) *ListDeploymentJobResourceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetId(v int64) *ListDeploymentJobResourceResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetInstanceId(v string) *ListDeploymentJobResourceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetListenerId(v string) *ListDeploymentJobResourceResponseBodyData {
	s.ListenerId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetListenerPort(v string) *ListDeploymentJobResourceResponseBodyData {
	s.ListenerPort = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetRegionId(v string) *ListDeploymentJobResourceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetRemark(v string) *ListDeploymentJobResourceResponseBodyData {
	s.Remark = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetStatus(v string) *ListDeploymentJobResourceResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetUseSsl(v int32) *ListDeploymentJobResourceResponseBodyData {
	s.UseSsl = &v
	return s
}

func (s *ListDeploymentJobResourceResponseBodyData) SetUserId(v int64) *ListDeploymentJobResourceResponseBodyData {
	s.UserId = &v
	return s
}

type ListDeploymentJobResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentJobResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentJobResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentJobResourceResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResourceResponse) SetHeaders(v map[string]*string) *ListDeploymentJobResourceResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentJobResourceResponse) SetStatusCode(v int32) *ListDeploymentJobResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentJobResourceResponse) SetBody(v *ListDeploymentJobResourceResponseBody) *ListDeploymentJobResourceResponse {
	s.Body = v
	return s
}

type ListUserCertificateOrderRequest struct {
	// The number of the page to return.
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The domain names that are bound or the ID of the order. Fuzzy match is supported.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The type of the order. Valid values:
	//
	// *   **CPACK**: virtual resource order. If you set OrderType to CPACK, only the information about orders that are generated to consume the certificate quota is returned.
	// *   **BUY**: purchase order. If you set OrderType to BUY, only the information about purchase orders is returned. In most cases, this type of order can be ignored.
	// *   **UPLOAD**: uploaded certificate. If you set OrderType to UPLOAD, only uploaded certificates are returned.
	// *   **CERT**: certificate. If you set OrderType to CERT, both issued certificates and uploaded certificates are returned.
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of entries to return on each page. Default value: 50.
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The certificate status of the order. Valid values:
	//
	// *   **PAYED**: pending application. You can set Status to PAYED only if you set OrderType to CPACK or BUY.
	// *   **CHECKING**: reviewing. You can set Status to CHECKING only if you set OrderType to CPACK or BUY.
	// *   **CHECKED_FAIL**: review failed. You can set Status to CHECKED_FAIL only if you set OrderType to CPACK or BUY.
	// *   **ISSUED**: issued.
	// *   **WILLEXPIRED**: about to expire.
	// *   **EXPIRED**: expired.
	// *   **NOTACTIVATED**: not activated. You can set Status to NOTACTIVATED only if you set OrderType to CPACK or BUY.
	// *   **REVOKED**: revoked. You can set Status to REVOKED only if you set OrderType to CPACK or BUY.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListUserCertificateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserCertificateOrderRequest) GoString() string {
	return s.String()
}

func (s *ListUserCertificateOrderRequest) SetCurrentPage(v int64) *ListUserCertificateOrderRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserCertificateOrderRequest) SetKeyword(v string) *ListUserCertificateOrderRequest {
	s.Keyword = &v
	return s
}

func (s *ListUserCertificateOrderRequest) SetOrderType(v string) *ListUserCertificateOrderRequest {
	s.OrderType = &v
	return s
}

func (s *ListUserCertificateOrderRequest) SetResourceGroupId(v string) *ListUserCertificateOrderRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListUserCertificateOrderRequest) SetShowSize(v int64) *ListUserCertificateOrderRequest {
	s.ShowSize = &v
	return s
}

func (s *ListUserCertificateOrderRequest) SetStatus(v string) *ListUserCertificateOrderRequest {
	s.Status = &v
	return s
}

type ListUserCertificateOrderResponseBody struct {
	// An array that consists of the information about the certificates and orders.
	CertificateOrderList []*ListUserCertificateOrderResponseBodyCertificateOrderList `json:"CertificateOrderList,omitempty" xml:"CertificateOrderList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned per page.
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserCertificateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserCertificateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserCertificateOrderResponseBody) SetCertificateOrderList(v []*ListUserCertificateOrderResponseBodyCertificateOrderList) *ListUserCertificateOrderResponseBody {
	s.CertificateOrderList = v
	return s
}

func (s *ListUserCertificateOrderResponseBody) SetCurrentPage(v int64) *ListUserCertificateOrderResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserCertificateOrderResponseBody) SetRequestId(v string) *ListUserCertificateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBody) SetShowSize(v int64) *ListUserCertificateOrderResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListUserCertificateOrderResponseBody) SetTotalCount(v int64) *ListUserCertificateOrderResponseBody {
	s.TotalCount = &v
	return s
}

type ListUserCertificateOrderResponseBodyCertificateOrderList struct {
	// The algorithm. This parameter is returned only if OrderType is set to CPACK or BUY.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the Alibaba Cloud order. This parameter is returned only if OrderType is set to CPACK or BUY.
	AliyunOrderId *int64 `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	// The time at which the order was placed. Unit: milliseconds. This parameter is returned only if OrderType is set to CPACK or BUY.
	BuyDate *int64 `json:"BuyDate,omitempty" xml:"BuyDate,omitempty"`
	// The time at which the certificate expires. Unit: milliseconds. This parameter is returned only if OrderType is set to CPACK or BUY.
	CertEndTime *int64 `json:"CertEndTime,omitempty" xml:"CertEndTime,omitempty"`
	// The time at which the certificate starts to take effect. Unit: milliseconds. This parameter is returned only if OrderType is set to CPACK or BUY.
	CertStartTime *int64 `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The type of the certificate. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// *   **DV**: domain validated (DV) certificate
	// *   **EV**: extended validation (EV) certificate
	// *   **OV**: organization validated (OV) certificate
	// *   **FREE**: free certificate
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The ID of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The city in which the organization is located. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The parent domain name of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The domain name. This parameter is returned only if OrderType is set to CPACK or BUY.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The total number of domain names that can be bound to the certificate. This parameter is returned only if OrderType is set to CPACK or BUY.
	DomainCount *int64 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The type of the domain name. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// *   **ONE**: single domain name
	// *   **MULTIPLE**: multiple domain names
	// *   **WILDCARD**: single wildcard domain name
	// *   **M_WILDCARD**: multiple wildcard domain names
	// *   **MIX**: hybrid domain name
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The time at which the certificate expires. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Indicates whether the certificate expires. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The fingerprint of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The ID of the resource.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issuer of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The name of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order ID. This parameter is returned only if OrderType is set to CPACK or BUY.
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The name of the organization that is associated with the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The ID of the certificate authority (CA) order. This parameter is returned only if OrderType is set to CPACK or BUY.
	PartnerOrderId *string `json:"PartnerOrderId,omitempty" xml:"PartnerOrderId,omitempty"`
	// The specification ID of the order. This parameter is returned only if OrderType is set to CPACK or BUY.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The specification name of the order. This parameter is returned only if OrderType is set to CPACK or BUY.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The name of the province or autonomous region in which the organization is located. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The ID of the resource group. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The brand of the certificate. Valid values: WoSign, CFCA, DigiCert, and vTrus. This parameter is returned only if OrderType is set to CPACK or BUY.
	RootBrand *string `json:"RootBrand,omitempty" xml:"RootBrand,omitempty"`
	// All domain names that are bound to the certificate. Multiple domain names are separated by commas (,). This parameter is returned only if OrderType is set to CERT or UPLOAD.
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	// The SHA-2 value of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The type of the order. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// *   **cpack**: virtual resource order
	// *   **buy**: purchase order
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time at which the certificate starts to take effect. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The certificate status of the order. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// *   **PAYED**: pending application
	// *   **CHECKING**: reviewing
	// *   **CHECKED_FAIL**: review failed
	// *   **ISSUED**: issued
	// *   **WILLEXPIRED**: about to expire
	// *   **EXPIRED**: expired
	// *   **NOTACTIVATED**: not activated
	// *   **REVOKED**: revoked
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The hosting status of the certificate. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// *   **unTrustee**: not hosted
	// *   **trustee**: hosted
	TrusteeStatus *string `json:"TrusteeStatus,omitempty" xml:"TrusteeStatus,omitempty"`
	// Indicates whether the certificate is an uploaded certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	Upload *bool `json:"Upload,omitempty" xml:"Upload,omitempty"`
	// The number of wildcard domain names that can be bound to the certificate. This parameter is returned only if OrderType is set to CPACK or BUY.
	WildDomainCount *int64 `json:"WildDomainCount,omitempty" xml:"WildDomainCount,omitempty"`
}

func (s ListUserCertificateOrderResponseBodyCertificateOrderList) String() string {
	return tea.Prettify(s)
}

func (s ListUserCertificateOrderResponseBodyCertificateOrderList) GoString() string {
	return s.String()
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetAlgorithm(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Algorithm = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetAliyunOrderId(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.AliyunOrderId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetBuyDate(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.BuyDate = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCertEndTime(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.CertEndTime = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCertStartTime(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.CertStartTime = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCertType(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.CertType = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCertificateId(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.CertificateId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCity(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.City = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCommonName(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.CommonName = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCountry(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Country = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetDomain(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Domain = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetDomainCount(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.DomainCount = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetDomainType(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.DomainType = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetEndDate(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.EndDate = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetExpired(v bool) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Expired = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetFingerprint(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Fingerprint = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetInstanceId(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.InstanceId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetIssuer(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Issuer = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetName(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Name = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetOrderId(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.OrderId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetOrgName(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.OrgName = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetPartnerOrderId(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.PartnerOrderId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetProductCode(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.ProductCode = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetProductName(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.ProductName = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetProvince(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Province = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetResourceGroupId(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetRootBrand(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.RootBrand = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetSans(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Sans = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetSerialNo(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.SerialNo = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetSha2(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Sha2 = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetSourceType(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.SourceType = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetStartDate(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.StartDate = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetStatus(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Status = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetTrusteeStatus(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.TrusteeStatus = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetUpload(v bool) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Upload = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetWildDomainCount(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.WildDomainCount = &v
	return s
}

type ListUserCertificateOrderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserCertificateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserCertificateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserCertificateOrderResponse) GoString() string {
	return s.String()
}

func (s *ListUserCertificateOrderResponse) SetHeaders(v map[string]*string) *ListUserCertificateOrderResponse {
	s.Headers = v
	return s
}

func (s *ListUserCertificateOrderResponse) SetStatusCode(v int32) *ListUserCertificateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserCertificateOrderResponse) SetBody(v *ListUserCertificateOrderResponseBody) *ListUserCertificateOrderResponse {
	s.Body = v
	return s
}

type ListWorkerResourceRequest struct {
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	JobId        *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ShowSize     *int32  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWorkerResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkerResourceRequest) GoString() string {
	return s.String()
}

func (s *ListWorkerResourceRequest) SetCloudProduct(v string) *ListWorkerResourceRequest {
	s.CloudProduct = &v
	return s
}

func (s *ListWorkerResourceRequest) SetCurrentPage(v int32) *ListWorkerResourceRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListWorkerResourceRequest) SetJobId(v int64) *ListWorkerResourceRequest {
	s.JobId = &v
	return s
}

func (s *ListWorkerResourceRequest) SetShowSize(v int32) *ListWorkerResourceRequest {
	s.ShowSize = &v
	return s
}

func (s *ListWorkerResourceRequest) SetStatus(v string) *ListWorkerResourceRequest {
	s.Status = &v
	return s
}

type ListWorkerResourceResponseBody struct {
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*ListWorkerResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize    *int32                                `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	Total       *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListWorkerResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkerResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkerResourceResponseBody) SetCurrentPage(v int32) *ListWorkerResourceResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListWorkerResourceResponseBody) SetData(v []*ListWorkerResourceResponseBodyData) *ListWorkerResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkerResourceResponseBody) SetRequestId(v string) *ListWorkerResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkerResourceResponseBody) SetShowSize(v int32) *ListWorkerResourceResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListWorkerResourceResponseBody) SetTotal(v int64) *ListWorkerResourceResponseBody {
	s.Total = &v
	return s
}

type ListWorkerResourceResponseBodyData struct {
	CertDomain      *string `json:"CertDomain,omitempty" xml:"CertDomain,omitempty"`
	CertId          *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertInstanceId  *string `json:"CertInstanceId,omitempty" xml:"CertInstanceId,omitempty"`
	CertName        *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CloudName       *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	CloudProduct    *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	CloudRegion     *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	DefaultResource *bool   `json:"DefaultResource,omitempty" xml:"DefaultResource,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ListenerId      *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	NamespaceId     *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	OrderId         *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceCertId  *int64  `json:"ResourceCertId,omitempty" xml:"ResourceCertId,omitempty"`
	ResourceDomain  *string `json:"ResourceDomain,omitempty" xml:"ResourceDomain,omitempty"`
	ResourceId      *int64  `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId          *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListWorkerResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWorkerResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkerResourceResponseBodyData) SetCertDomain(v string) *ListWorkerResourceResponseBodyData {
	s.CertDomain = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCertId(v int64) *ListWorkerResourceResponseBodyData {
	s.CertId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCertInstanceId(v string) *ListWorkerResourceResponseBodyData {
	s.CertInstanceId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCertName(v string) *ListWorkerResourceResponseBodyData {
	s.CertName = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCloudName(v string) *ListWorkerResourceResponseBodyData {
	s.CloudName = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCloudProduct(v string) *ListWorkerResourceResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetCloudRegion(v string) *ListWorkerResourceResponseBodyData {
	s.CloudRegion = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetDefaultResource(v bool) *ListWorkerResourceResponseBodyData {
	s.DefaultResource = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetGmtCreate(v string) *ListWorkerResourceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetGmtModified(v string) *ListWorkerResourceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetId(v int64) *ListWorkerResourceResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetInstanceId(v string) *ListWorkerResourceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetJobId(v int64) *ListWorkerResourceResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetListenerId(v string) *ListWorkerResourceResponseBodyData {
	s.ListenerId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetNamespaceId(v string) *ListWorkerResourceResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetOrderId(v int64) *ListWorkerResourceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetPort(v int32) *ListWorkerResourceResponseBodyData {
	s.Port = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetRegionId(v string) *ListWorkerResourceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetResourceCertId(v int64) *ListWorkerResourceResponseBodyData {
	s.ResourceCertId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetResourceDomain(v string) *ListWorkerResourceResponseBodyData {
	s.ResourceDomain = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetResourceId(v int64) *ListWorkerResourceResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetStatus(v string) *ListWorkerResourceResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListWorkerResourceResponseBodyData) SetUserId(v int64) *ListWorkerResourceResponseBodyData {
	s.UserId = &v
	return s
}

type ListWorkerResourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkerResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkerResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkerResourceResponse) GoString() string {
	return s.String()
}

func (s *ListWorkerResourceResponse) SetHeaders(v map[string]*string) *ListWorkerResourceResponse {
	s.Headers = v
	return s
}

func (s *ListWorkerResourceResponse) SetStatusCode(v int32) *ListWorkerResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkerResourceResponse) SetBody(v *ListWorkerResourceResponseBody) *ListWorkerResourceResponse {
	s.Body = v
	return s
}

type RenewCertificateOrderForPackageRequestRequest struct {
	// The content of the certificate signing request (CSR) file that is manually generated for the domain name by using OpenSSL or Keytool. The key algorithm in the CSR file must be Rivest-Shamir-Adleman (RSA) or elliptic-curve cryptography (ECC), and the key length of the RSA algorithm must be greater than or equal to 2,048 characters. For more information about how to create a CSR file, see [How do I create a CSR file?](~~42218~~)
	//
	// If you do not specify this parameter, Certificate Management Service automatically generates a CSR file for the domain name in the certificate application order that you want to renew.
	//
	// A CSR file contains the information about your server and company. When you apply for a certificate, you must submit the CSR file to the CA. The CA signs the CSR file by using the private key of the root certificate and generates a public key file to issue your certificate.
	//
	// > The **CN** field in the CSR file specifies the domain name that is bound to the certificate.
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The ID of the certificate application order that you want to renew.
	//
	// > After you call the [CreateCertificateForPackageRequest](~~455296~~), [CreateCertificateRequest](~~455292~~), or [CreateCertificateWithCsrRequest](~~455801~~) operation to submit a certificate application, you can obtain the ID of the certificate application order from the **OrderId** response parameter.
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewCertificateOrderForPackageRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewCertificateOrderForPackageRequestRequest) GoString() string {
	return s.String()
}

func (s *RenewCertificateOrderForPackageRequestRequest) SetCsr(v string) *RenewCertificateOrderForPackageRequestRequest {
	s.Csr = &v
	return s
}

func (s *RenewCertificateOrderForPackageRequestRequest) SetOrderId(v int64) *RenewCertificateOrderForPackageRequestRequest {
	s.OrderId = &v
	return s
}

type RenewCertificateOrderForPackageRequestResponseBody struct {
	// The ID of the certificate application order that is renewed.
	//
	// > You can use the ID to query the status of the certificate application. For more information, see [DescribeCertificateState](~~455800~~).
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewCertificateOrderForPackageRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewCertificateOrderForPackageRequestResponseBody) GoString() string {
	return s.String()
}

func (s *RenewCertificateOrderForPackageRequestResponseBody) SetOrderId(v int64) *RenewCertificateOrderForPackageRequestResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewCertificateOrderForPackageRequestResponseBody) SetRequestId(v string) *RenewCertificateOrderForPackageRequestResponseBody {
	s.RequestId = &v
	return s
}

type RenewCertificateOrderForPackageRequestResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewCertificateOrderForPackageRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewCertificateOrderForPackageRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewCertificateOrderForPackageRequestResponse) GoString() string {
	return s.String()
}

func (s *RenewCertificateOrderForPackageRequestResponse) SetHeaders(v map[string]*string) *RenewCertificateOrderForPackageRequestResponse {
	s.Headers = v
	return s
}

func (s *RenewCertificateOrderForPackageRequestResponse) SetStatusCode(v int32) *RenewCertificateOrderForPackageRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewCertificateOrderForPackageRequestResponse) SetBody(v *RenewCertificateOrderForPackageRequestResponseBody) *RenewCertificateOrderForPackageRequestResponse {
	s.Body = v
	return s
}

type RevokeWHClientCertificateRequest struct {
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s RevokeWHClientCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeWHClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *RevokeWHClientCertificateRequest) SetIdentifier(v string) *RevokeWHClientCertificateRequest {
	s.Identifier = &v
	return s
}

type RevokeWHClientCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeWHClientCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeWHClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeWHClientCertificateResponseBody) SetRequestId(v string) *RevokeWHClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

type RevokeWHClientCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeWHClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeWHClientCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeWHClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *RevokeWHClientCertificateResponse) SetHeaders(v map[string]*string) *RevokeWHClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *RevokeWHClientCertificateResponse) SetStatusCode(v int32) *RevokeWHClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeWHClientCertificateResponse) SetBody(v *RevokeWHClientCertificateResponseBody) *RevokeWHClientCertificateResponse {
	s.Body = v
	return s
}

type SignRequest struct {
	CertIdentifier   *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageType      *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	SigningAlgorithm *string `json:"SigningAlgorithm,omitempty" xml:"SigningAlgorithm,omitempty"`
}

func (s SignRequest) String() string {
	return tea.Prettify(s)
}

func (s SignRequest) GoString() string {
	return s.String()
}

func (s *SignRequest) SetCertIdentifier(v string) *SignRequest {
	s.CertIdentifier = &v
	return s
}

func (s *SignRequest) SetMessage(v string) *SignRequest {
	s.Message = &v
	return s
}

func (s *SignRequest) SetMessageType(v string) *SignRequest {
	s.MessageType = &v
	return s
}

func (s *SignRequest) SetSigningAlgorithm(v string) *SignRequest {
	s.SigningAlgorithm = &v
	return s
}

type SignResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s SignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SignResponseBody) GoString() string {
	return s.String()
}

func (s *SignResponseBody) SetRequestId(v string) *SignResponseBody {
	s.RequestId = &v
	return s
}

func (s *SignResponseBody) SetSignature(v string) *SignResponseBody {
	s.Signature = &v
	return s
}

type SignResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SignResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SignResponse) String() string {
	return tea.Prettify(s)
}

func (s SignResponse) GoString() string {
	return s.String()
}

func (s *SignResponse) SetHeaders(v map[string]*string) *SignResponse {
	s.Headers = v
	return s
}

func (s *SignResponse) SetStatusCode(v int32) *SignResponse {
	s.StatusCode = &v
	return s
}

func (s *SignResponse) SetBody(v *SignResponseBody) *SignResponse {
	s.Body = v
	return s
}

type UpdateCsrRequest struct {
	// CSR ID。
	CsrId *int64  `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s UpdateCsrRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCsrRequest) GoString() string {
	return s.String()
}

func (s *UpdateCsrRequest) SetCsrId(v int64) *UpdateCsrRequest {
	s.CsrId = &v
	return s
}

func (s *UpdateCsrRequest) SetKey(v string) *UpdateCsrRequest {
	s.Key = &v
	return s
}

type UpdateCsrResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCsrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCsrResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCsrResponseBody) SetRequestId(v string) *UpdateCsrResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCsrResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCsrResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCsrResponse) GoString() string {
	return s.String()
}

func (s *UpdateCsrResponse) SetHeaders(v map[string]*string) *UpdateCsrResponse {
	s.Headers = v
	return s
}

func (s *UpdateCsrResponse) SetStatusCode(v int32) *UpdateCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCsrResponse) SetBody(v *UpdateCsrResponseBody) *UpdateCsrResponse {
	s.Body = v
	return s
}

type UpdateDeploymentJobRequest struct {
	CertIds      *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	ContactIds   *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	JobId        *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceIds  *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ScheduleTime *int64  `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s UpdateDeploymentJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobRequest) SetCertIds(v string) *UpdateDeploymentJobRequest {
	s.CertIds = &v
	return s
}

func (s *UpdateDeploymentJobRequest) SetContactIds(v string) *UpdateDeploymentJobRequest {
	s.ContactIds = &v
	return s
}

func (s *UpdateDeploymentJobRequest) SetJobId(v int64) *UpdateDeploymentJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateDeploymentJobRequest) SetName(v string) *UpdateDeploymentJobRequest {
	s.Name = &v
	return s
}

func (s *UpdateDeploymentJobRequest) SetResourceIds(v string) *UpdateDeploymentJobRequest {
	s.ResourceIds = &v
	return s
}

func (s *UpdateDeploymentJobRequest) SetScheduleTime(v int64) *UpdateDeploymentJobRequest {
	s.ScheduleTime = &v
	return s
}

type UpdateDeploymentJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeploymentJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobResponseBody) SetRequestId(v string) *UpdateDeploymentJobResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDeploymentJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobResponse) SetHeaders(v map[string]*string) *UpdateDeploymentJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentJobResponse) SetStatusCode(v int32) *UpdateDeploymentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentJobResponse) SetBody(v *UpdateDeploymentJobResponseBody) *UpdateDeploymentJobResponse {
	s.Body = v
	return s
}

type UpdateDeploymentJobStatusRequest struct {
	JobId  *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateDeploymentJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentJobStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobStatusRequest) SetJobId(v int64) *UpdateDeploymentJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *UpdateDeploymentJobStatusRequest) SetStatus(v string) *UpdateDeploymentJobStatusRequest {
	s.Status = &v
	return s
}

type UpdateDeploymentJobStatusResponseBody struct {
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeploymentJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobStatusResponseBody) SetData(v interface{}) *UpdateDeploymentJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDeploymentJobStatusResponseBody) SetRequestId(v string) *UpdateDeploymentJobStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDeploymentJobStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeploymentJobStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobStatusResponse) SetHeaders(v map[string]*string) *UpdateDeploymentJobStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentJobStatusResponse) SetStatusCode(v int32) *UpdateDeploymentJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentJobStatusResponse) SetBody(v *UpdateDeploymentJobStatusResponseBody) *UpdateDeploymentJobStatusResponse {
	s.Body = v
	return s
}

type UpdateWorkerResourceStatusRequest struct {
	JobId    *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkerId *int64  `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s UpdateWorkerResourceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkerResourceStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResourceStatusRequest) SetJobId(v int64) *UpdateWorkerResourceStatusRequest {
	s.JobId = &v
	return s
}

func (s *UpdateWorkerResourceStatusRequest) SetStatus(v string) *UpdateWorkerResourceStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateWorkerResourceStatusRequest) SetWorkerId(v int64) *UpdateWorkerResourceStatusRequest {
	s.WorkerId = &v
	return s
}

type UpdateWorkerResourceStatusResponseBody struct {
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWorkerResourceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkerResourceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResourceStatusResponseBody) SetData(v interface{}) *UpdateWorkerResourceStatusResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWorkerResourceStatusResponseBody) SetRequestId(v string) *UpdateWorkerResourceStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWorkerResourceStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkerResourceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkerResourceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkerResourceStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResourceStatusResponse) SetHeaders(v map[string]*string) *UpdateWorkerResourceStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkerResourceStatusResponse) SetStatusCode(v int32) *UpdateWorkerResourceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkerResourceStatusResponse) SetBody(v *UpdateWorkerResourceStatusResponseBody) *UpdateWorkerResourceStatusResponse {
	s.Body = v
	return s
}

type UploadCsrRequest struct {
	Csr  *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	Key  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UploadCsrRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCsrRequest) GoString() string {
	return s.String()
}

func (s *UploadCsrRequest) SetCsr(v string) *UploadCsrRequest {
	s.Csr = &v
	return s
}

func (s *UploadCsrRequest) SetKey(v string) *UploadCsrRequest {
	s.Key = &v
	return s
}

func (s *UploadCsrRequest) SetName(v string) *UploadCsrRequest {
	s.Name = &v
	return s
}

type UploadCsrResponseBody struct {
	// CSR ID。
	CsrId     *int64  `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadCsrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadCsrResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCsrResponseBody) SetCsrId(v int64) *UploadCsrResponseBody {
	s.CsrId = &v
	return s
}

func (s *UploadCsrResponseBody) SetRequestId(v string) *UploadCsrResponseBody {
	s.RequestId = &v
	return s
}

type UploadCsrResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadCsrResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCsrResponse) GoString() string {
	return s.String()
}

func (s *UploadCsrResponse) SetHeaders(v map[string]*string) *UploadCsrResponse {
	s.Headers = v
	return s
}

func (s *UploadCsrResponse) SetStatusCode(v int32) *UploadCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadCsrResponse) SetBody(v *UploadCsrResponseBody) *UploadCsrResponse {
	s.Body = v
	return s
}

type UploadPCACertRequest struct {
	// <UploadPCACertResponse>
	//     <RequestId>15C66C7B-671A-4297-9187-2C4477247A74</RequestId>
	// </UploadPCACertResponse>
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// UploadPCACert
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Uploads a private certificate to a certificate repository.
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// {
	//     "RequestId": "15C66C7B-671A-4297-9187-2C4477247A74"
	// }
	WarehouseId *int64 `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
}

func (s UploadPCACertRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadPCACertRequest) GoString() string {
	return s.String()
}

func (s *UploadPCACertRequest) SetCert(v string) *UploadPCACertRequest {
	s.Cert = &v
	return s
}

func (s *UploadPCACertRequest) SetName(v string) *UploadPCACertRequest {
	s.Name = &v
	return s
}

func (s *UploadPCACertRequest) SetPrivateKey(v string) *UploadPCACertRequest {
	s.PrivateKey = &v
	return s
}

func (s *UploadPCACertRequest) SetWarehouseId(v int64) *UploadPCACertRequest {
	s.WarehouseId = &v
	return s
}

type UploadPCACertResponseBody struct {
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadPCACertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadPCACertResponseBody) GoString() string {
	return s.String()
}

func (s *UploadPCACertResponseBody) SetIdentifier(v string) *UploadPCACertResponseBody {
	s.Identifier = &v
	return s
}

func (s *UploadPCACertResponseBody) SetRequestId(v string) *UploadPCACertResponseBody {
	s.RequestId = &v
	return s
}

type UploadPCACertResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadPCACertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadPCACertResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadPCACertResponse) GoString() string {
	return s.String()
}

func (s *UploadPCACertResponse) SetHeaders(v map[string]*string) *UploadPCACertResponse {
	s.Headers = v
	return s
}

func (s *UploadPCACertResponse) SetStatusCode(v int32) *UploadPCACertResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadPCACertResponse) SetBody(v *UploadPCACertResponseBody) *UploadPCACertResponse {
	s.Body = v
	return s
}

type UploadUserCertificateRequest struct {
	// The content of the certificate in the PEM format.
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The content of the encryption certificate in PEM format.
	EncryptCert *string `json:"EncryptCert,omitempty" xml:"EncryptCert,omitempty"`
	// The private key of the encryption certificate in the PEM format.
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	// The private key of the certificate in the PEM format.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the certificate. The name can contain up to 128 characters in length. The name can contain all types of characters, such as letters, digits, and underscores (\_).
	//
	// >  The name must be unique within an Alibaba Cloud account.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// the resource group id.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The content of the signing certificate in the PEM format.
	SignCert *string `json:"SignCert,omitempty" xml:"SignCert,omitempty"`
	// The private key of the signing certificate in the PEM format.
	SignPrivateKey *string `json:"SignPrivateKey,omitempty" xml:"SignPrivateKey,omitempty"`
}

func (s UploadUserCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadUserCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadUserCertificateRequest) SetCert(v string) *UploadUserCertificateRequest {
	s.Cert = &v
	return s
}

func (s *UploadUserCertificateRequest) SetEncryptCert(v string) *UploadUserCertificateRequest {
	s.EncryptCert = &v
	return s
}

func (s *UploadUserCertificateRequest) SetEncryptPrivateKey(v string) *UploadUserCertificateRequest {
	s.EncryptPrivateKey = &v
	return s
}

func (s *UploadUserCertificateRequest) SetKey(v string) *UploadUserCertificateRequest {
	s.Key = &v
	return s
}

func (s *UploadUserCertificateRequest) SetName(v string) *UploadUserCertificateRequest {
	s.Name = &v
	return s
}

func (s *UploadUserCertificateRequest) SetResourceGroupId(v string) *UploadUserCertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UploadUserCertificateRequest) SetSignCert(v string) *UploadUserCertificateRequest {
	s.SignCert = &v
	return s
}

func (s *UploadUserCertificateRequest) SetSignPrivateKey(v string) *UploadUserCertificateRequest {
	s.SignPrivateKey = &v
	return s
}

type UploadUserCertificateResponseBody struct {
	// The ID of the certificate.
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadUserCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadUserCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadUserCertificateResponseBody) SetCertId(v int64) *UploadUserCertificateResponseBody {
	s.CertId = &v
	return s
}

func (s *UploadUserCertificateResponseBody) SetRequestId(v string) *UploadUserCertificateResponseBody {
	s.RequestId = &v
	return s
}

type UploadUserCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadUserCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadUserCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadUserCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadUserCertificateResponse) SetHeaders(v map[string]*string) *UploadUserCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadUserCertificateResponse) SetStatusCode(v int32) *UploadUserCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadUserCertificateResponse) SetBody(v *UploadUserCertificateResponseBody) *UploadUserCertificateResponse {
	s.Body = v
	return s
}

type VerifyRequest struct {
	CertIdentifier   *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageType      *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	SignatureValue   *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
	SigningAlgorithm *string `json:"SigningAlgorithm,omitempty" xml:"SigningAlgorithm,omitempty"`
}

func (s VerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyRequest) GoString() string {
	return s.String()
}

func (s *VerifyRequest) SetCertIdentifier(v string) *VerifyRequest {
	s.CertIdentifier = &v
	return s
}

func (s *VerifyRequest) SetMessage(v string) *VerifyRequest {
	s.Message = &v
	return s
}

func (s *VerifyRequest) SetMessageType(v string) *VerifyRequest {
	s.MessageType = &v
	return s
}

func (s *VerifyRequest) SetSignatureValue(v string) *VerifyRequest {
	s.SignatureValue = &v
	return s
}

func (s *VerifyRequest) SetSigningAlgorithm(v string) *VerifyRequest {
	s.SigningAlgorithm = &v
	return s
}

type VerifyResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignatureValid *bool   `json:"SignatureValid,omitempty" xml:"SignatureValid,omitempty"`
}

func (s VerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyResponseBody) SetRequestId(v string) *VerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyResponseBody) SetSignatureValid(v bool) *VerifyResponseBody {
	s.SignatureValid = &v
	return s
}

type VerifyResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyResponse) GoString() string {
	return s.String()
}

func (s *VerifyResponse) SetHeaders(v map[string]*string) *VerifyResponse {
	s.Headers = v
	return s
}

func (s *VerifyResponse) SetStatusCode(v int32) *VerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyResponse) SetBody(v *VerifyResponseBody) *VerifyResponse {
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

/**
 * Revokes an issued certificate and cancels the application order of the certificate.
 *
 * @param request CancelCertificateForPackageRequestRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelCertificateForPackageRequestResponse
 */
func (client *Client) CancelCertificateForPackageRequestWithOptions(request *CancelCertificateForPackageRequestRequest, runtime *util.RuntimeOptions) (_result *CancelCertificateForPackageRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelCertificateForPackageRequest"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelCertificateForPackageRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Revokes an issued certificate and cancels the application order of the certificate.
 *
 * @param request CancelCertificateForPackageRequestRequest
 * @return CancelCertificateForPackageRequestResponse
 */
func (client *Client) CancelCertificateForPackageRequest(request *CancelCertificateForPackageRequestRequest) (_result *CancelCertificateForPackageRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelCertificateForPackageRequestResponse{}
	_body, _err := client.CancelCertificateForPackageRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the CancelOrderRequest operation to cancel a certificate application order only in the following scenarios:
 * *   The order is in the **pending validation** state. You have submitted a certificate application but the verification of the domain name ownership is not complete.
 * *   The order is in the **being reviewed** state. You have submitted a certificate application and the verification of the domain name ownership is complete, but the certificate authority (CA) does not complete the review of the certificate application.
 * After a certificate application order is canceled, the status of the order changes to the **pending application** state. In this case, you can call the [DeleteCertificateRequest](~~164109~~) operation to delete the certificate application order. Then, the consumed certificate quota is returned to you.
 *
 * @param request CancelOrderRequestRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelOrderRequestResponse
 */
func (client *Client) CancelOrderRequestWithOptions(request *CancelOrderRequestRequest, runtime *util.RuntimeOptions) (_result *CancelOrderRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelOrderRequest"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelOrderRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the CancelOrderRequest operation to cancel a certificate application order only in the following scenarios:
 * *   The order is in the **pending validation** state. You have submitted a certificate application but the verification of the domain name ownership is not complete.
 * *   The order is in the **being reviewed** state. You have submitted a certificate application and the verification of the domain name ownership is complete, but the certificate authority (CA) does not complete the review of the certificate application.
 * After a certificate application order is canceled, the status of the order changes to the **pending application** state. In this case, you can call the [DeleteCertificateRequest](~~164109~~) operation to delete the certificate application order. Then, the consumed certificate quota is returned to you.
 *
 * @param request CancelOrderRequestRequest
 * @return CancelOrderRequestResponse
 */
func (client *Client) CancelOrderRequest(request *CancelOrderRequestRequest) (_result *CancelOrderRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelOrderRequestResponse{}
	_body, _err := client.CancelOrderRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](~~28542~~). You can call the [DescribePackageState](~~455800~~) operation to query the usage of certificate resource plans of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that are submitted, and the number of certificates that are issued.
 * *   After you call this operation to submit a certificate application and the certificate is issued, the certificate quota provided by the resource plan that you purchased is consumed. When you call this operation, you can use the **ProductCode** parameter to specify the specifications of the certificate that you want to apply for.
 * *   After you call this operation to submit a certificate application, you also need to call the [DescribeCertificateState](~~455800~~) operation to obtain the information that is required to complete the verification of the domain name ownership, and complete the verification. If you use the DNS verification method, you must complete the verification in the management platform of the domain name. If you use the file verification method, you must complete the verification in the DNS server. Then, the certificate application order will be reviewed by the certificate authority (CA).
 *
 * @param request CreateCertificateForPackageRequestRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateCertificateForPackageRequestResponse
 */
func (client *Client) CreateCertificateForPackageRequestWithOptions(request *CreateCertificateForPackageRequestRequest, runtime *util.RuntimeOptions) (_result *CreateCertificateForPackageRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompanyName)) {
		query["CompanyName"] = request.CompanyName
	}

	if !tea.BoolValue(util.IsUnset(request.Csr)) {
		query["Csr"] = request.Csr
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.ValidateType)) {
		query["ValidateType"] = request.ValidateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCertificateForPackageRequest"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCertificateForPackageRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](~~28542~~). You can call the [DescribePackageState](~~455800~~) operation to query the usage of certificate resource plans of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that are submitted, and the number of certificates that are issued.
 * *   After you call this operation to submit a certificate application and the certificate is issued, the certificate quota provided by the resource plan that you purchased is consumed. When you call this operation, you can use the **ProductCode** parameter to specify the specifications of the certificate that you want to apply for.
 * *   After you call this operation to submit a certificate application, you also need to call the [DescribeCertificateState](~~455800~~) operation to obtain the information that is required to complete the verification of the domain name ownership, and complete the verification. If you use the DNS verification method, you must complete the verification in the management platform of the domain name. If you use the file verification method, you must complete the verification in the DNS server. Then, the certificate application order will be reviewed by the certificate authority (CA).
 *
 * @param request CreateCertificateForPackageRequestRequest
 * @return CreateCertificateForPackageRequestResponse
 */
func (client *Client) CreateCertificateForPackageRequest(request *CreateCertificateForPackageRequestRequest) (_result *CreateCertificateForPackageRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertificateForPackageRequestResponse{}
	_body, _err := client.CreateCertificateForPackageRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   You can call this operation to apply for only DV certificates. If you want to apply for an organization validated (OV) or extended validation (EV) certificate, we recommend that you call the [CreateCertificateForPackageRequest](~~455296~~) operation. This operation allows you to apply for certificates of all specifications and specify the method to generate a certificate signing request (CSR) file.
 * *   Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](~~28542~~). You can call the [DescribePackageState](~~455803~~) operation to query the usage of certificate resource plans of specified specifications, including the total number of purchased certificate resource plans of the specified specifications, the number of times that certificate applications have been submitted, and the number of times that certificates have been issued.
 * *   When you call this operation, you can use the **ProductCode** parameter to specify the specifications of the certificate.
 * *   After you call this operation to submit a certificate application, Certificate Management Service automatically creates a CSR file for your application and consumes the certificate quota in the certificate resource plans of the specified specifications that you purchased. After you call this operation, you also need to call the [DescribeCertificateState](~~455800~~) operation to obtain the information that is required to complete domain name verification, and manually complete the verification. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on your DNS server. Then, the certificate authority (CA) will review your certificate application.
 *
 * @param request CreateCertificateRequestRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateCertificateRequestResponse
 */
func (client *Client) CreateCertificateRequestWithOptions(request *CreateCertificateRequestRequest, runtime *util.RuntimeOptions) (_result *CreateCertificateRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.ValidateType)) {
		query["ValidateType"] = request.ValidateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCertificateRequest"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCertificateRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   You can call this operation to apply for only DV certificates. If you want to apply for an organization validated (OV) or extended validation (EV) certificate, we recommend that you call the [CreateCertificateForPackageRequest](~~455296~~) operation. This operation allows you to apply for certificates of all specifications and specify the method to generate a certificate signing request (CSR) file.
 * *   Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](~~28542~~). You can call the [DescribePackageState](~~455803~~) operation to query the usage of certificate resource plans of specified specifications, including the total number of purchased certificate resource plans of the specified specifications, the number of times that certificate applications have been submitted, and the number of times that certificates have been issued.
 * *   When you call this operation, you can use the **ProductCode** parameter to specify the specifications of the certificate.
 * *   After you call this operation to submit a certificate application, Certificate Management Service automatically creates a CSR file for your application and consumes the certificate quota in the certificate resource plans of the specified specifications that you purchased. After you call this operation, you also need to call the [DescribeCertificateState](~~455800~~) operation to obtain the information that is required to complete domain name verification, and manually complete the verification. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on your DNS server. Then, the certificate authority (CA) will review your certificate application.
 *
 * @param request CreateCertificateRequestRequest
 * @return CreateCertificateRequestResponse
 */
func (client *Client) CreateCertificateRequest(request *CreateCertificateRequestRequest) (_result *CreateCertificateRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertificateRequestResponse{}
	_body, _err := client.CreateCertificateRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   You can call the CreateCertificateWithCsrRequest operation to apply only for DV certificates. We recommend that you call the [CreateCertificateForPackageRequest](~~455296~~) operation to submit a certificate application. This operation allows you to apply for certificates of all specifications and specify the method to generate a CSR file.
 * *   Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](~~28542~~). You can call the [DescribePackageState](~~164110~~) operation to query the usage of certificate resource plans of specified specifications. The usage information includes the total number of purchased certificate resource plans of the specified specifications, the number of times that certificate applications are submitted, and the number of times that certificates are issued.
 * *   When you call this operation, you can use the **ProductCode** parameter to specify the specifications of the certificate.
 * *   After you call this operation to submit a certificate application, the certificate quota of the required specifications that you purchased is consumed. After you call this operation, you also need to call the [DescribeCertificateState](~~164111~~) operation to obtain the information that is required to complete domain name verification, and manually complete the verification. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on your DNS server. The certificate authority (CA) starts to review your certificate application only after the domain name verification is complete.
 *
 * @param request CreateCertificateWithCsrRequestRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateCertificateWithCsrRequestResponse
 */
func (client *Client) CreateCertificateWithCsrRequestWithOptions(request *CreateCertificateWithCsrRequestRequest, runtime *util.RuntimeOptions) (_result *CreateCertificateWithCsrRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Csr)) {
		query["Csr"] = request.Csr
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.ValidateType)) {
		query["ValidateType"] = request.ValidateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCertificateWithCsrRequest"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCertificateWithCsrRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   You can call the CreateCertificateWithCsrRequest operation to apply only for DV certificates. We recommend that you call the [CreateCertificateForPackageRequest](~~455296~~) operation to submit a certificate application. This operation allows you to apply for certificates of all specifications and specify the method to generate a CSR file.
 * *   Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](~~28542~~). You can call the [DescribePackageState](~~164110~~) operation to query the usage of certificate resource plans of specified specifications. The usage information includes the total number of purchased certificate resource plans of the specified specifications, the number of times that certificate applications are submitted, and the number of times that certificates are issued.
 * *   When you call this operation, you can use the **ProductCode** parameter to specify the specifications of the certificate.
 * *   After you call this operation to submit a certificate application, the certificate quota of the required specifications that you purchased is consumed. After you call this operation, you also need to call the [DescribeCertificateState](~~164111~~) operation to obtain the information that is required to complete domain name verification, and manually complete the verification. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on your DNS server. The certificate authority (CA) starts to review your certificate application only after the domain name verification is complete.
 *
 * @param request CreateCertificateWithCsrRequestRequest
 * @return CreateCertificateWithCsrRequestResponse
 */
func (client *Client) CreateCertificateWithCsrRequest(request *CreateCertificateWithCsrRequestRequest) (_result *CreateCertificateWithCsrRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertificateWithCsrRequestResponse{}
	_body, _err := client.CreateCertificateWithCsrRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCsrWithOptions(request *CreateCsrRequest, runtime *util.RuntimeOptions) (_result *CreateCsrResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.CorpName)) {
		query["CorpName"] = request.CorpName
	}

	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		query["CountryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.Department)) {
		query["Department"] = request.Department
	}

	if !tea.BoolValue(util.IsUnset(request.KeySize)) {
		query["KeySize"] = request.KeySize
	}

	if !tea.BoolValue(util.IsUnset(request.Locality)) {
		query["Locality"] = request.Locality
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.Sans)) {
		query["Sans"] = request.Sans
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCsr"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCsr(request *CreateCsrRequest) (_result *CreateCsrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCsrResponse{}
	_body, _err := client.CreateCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeploymentJobWithOptions(request *CreateDeploymentJobRequest, runtime *util.RuntimeOptions) (_result *CreateDeploymentJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIds)) {
		query["CertIds"] = request.CertIds
	}

	if !tea.BoolValue(util.IsUnset(request.ContactIds)) {
		query["ContactIds"] = request.ContactIds
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleTime)) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeploymentJob"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeploymentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeploymentJob(request *CreateDeploymentJobRequest) (_result *CreateDeploymentJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeploymentJobResponse{}
	_body, _err := client.CreateDeploymentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWHClientCertificateWithOptions(request *CreateWHClientCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateWHClientCertificateResponse, _err error) {
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
		Action:      tea.String("CreateWHClientCertificate"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWHClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWHClientCertificate(request *CreateWHClientCertificateRequest) (_result *CreateWHClientCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWHClientCertificateResponse{}
	_body, _err := client.CreateWHClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DecryptWithOptions(request *DecryptRequest, runtime *util.RuntimeOptions) (_result *DecryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.CiphertextBlob)) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Decrypt"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DecryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Decrypt(request *DecryptRequest) (_result *DecryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DecryptResponse{}
	_body, _err := client.DecryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to delete a certificate application order only in the following scenarios:
 * *   The status of the order is review failed. You have called the [DescribeCertificateState](~~455800~~)  operation to query the status of the certificate application order and the value of the **Type** parameter is **verify_fail**.
 * *   The status of the order is **pending application**. You have called the [CancelOrderRequest](~~455299~~) operation to cancel a certificate application order whose status is pending review or being reviewed. The status of the certificate application order that is canceled in this case changes to **pending application**.
 *
 * @param request DeleteCertificateRequestRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteCertificateRequestResponse
 */
func (client *Client) DeleteCertificateRequestWithOptions(request *DeleteCertificateRequestRequest, runtime *util.RuntimeOptions) (_result *DeleteCertificateRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCertificateRequest"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCertificateRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to delete a certificate application order only in the following scenarios:
 * *   The status of the order is review failed. You have called the [DescribeCertificateState](~~455800~~)  operation to query the status of the certificate application order and the value of the **Type** parameter is **verify_fail**.
 * *   The status of the order is **pending application**. You have called the [CancelOrderRequest](~~455299~~) operation to cancel a certificate application order whose status is pending review or being reviewed. The status of the certificate application order that is canceled in this case changes to **pending application**.
 *
 * @param request DeleteCertificateRequestRequest
 * @return DeleteCertificateRequestResponse
 */
func (client *Client) DeleteCertificateRequest(request *DeleteCertificateRequestRequest) (_result *DeleteCertificateRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCertificateRequestResponse{}
	_body, _err := client.DeleteCertificateRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCsrWithOptions(request *DeleteCsrRequest, runtime *util.RuntimeOptions) (_result *DeleteCsrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CsrId)) {
		query["CsrId"] = request.CsrId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCsr"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCsr(request *DeleteCsrRequest) (_result *DeleteCsrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCsrResponse{}
	_body, _err := client.DeleteCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeploymentJobWithOptions(request *DeleteDeploymentJobRequest, runtime *util.RuntimeOptions) (_result *DeleteDeploymentJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDeploymentJob"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeploymentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeploymentJob(request *DeleteDeploymentJobRequest) (_result *DeleteDeploymentJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeploymentJobResponse{}
	_body, _err := client.DeleteDeploymentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePCACertWithOptions(request *DeletePCACertRequest, runtime *util.RuntimeOptions) (_result *DeletePCACertResponse, _err error) {
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
		Action:      tea.String("DeletePCACert"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePCACertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePCACert(request *DeletePCACertRequest) (_result *DeletePCACertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePCACertResponse{}
	_body, _err := client.DeletePCACertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteUserCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteUserCertificateResponse
 */
func (client *Client) DeleteUserCertificateWithOptions(request *DeleteUserCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteUserCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertId)) {
		query["CertId"] = request.CertId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserCertificate"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteUserCertificateRequest
 * @return DeleteUserCertificateResponse
 */
func (client *Client) DeleteUserCertificate(request *DeleteUserCertificateRequest) (_result *DeleteUserCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserCertificateResponse{}
	_body, _err := client.DeleteUserCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkerResourceWithOptions(request *DeleteWorkerResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteWorkerResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerId)) {
		query["WorkerId"] = request.WorkerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkerResource"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkerResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkerResource(request *DeleteWorkerResourceRequest) (_result *DeleteWorkerResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWorkerResourceResponse{}
	_body, _err := client.DeleteWorkerResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you do not complete the verification of the domain name ownership after you submit a certificate application, you can call this operation to obtain the information that is required to complete the verification. You can complete the verification of the domain name ownership based on the data returned. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on the DNS server.
 * The certificate authority (CA) reviews your certificate application only after you complete the verification of the domain name ownership. After the CA approves your certificate application, the CA issues the certificate. If a certificate is issued, you can call this operation to obtain the CA certificate and private key of the certificate.
 *
 * @param request DescribeCertificateStateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeCertificateStateResponse
 */
func (client *Client) DescribeCertificateStateWithOptions(request *DescribeCertificateStateRequest, runtime *util.RuntimeOptions) (_result *DescribeCertificateStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCertificateState"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCertificateStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you do not complete the verification of the domain name ownership after you submit a certificate application, you can call this operation to obtain the information that is required to complete the verification. You can complete the verification of the domain name ownership based on the data returned. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on the DNS server.
 * The certificate authority (CA) reviews your certificate application only after you complete the verification of the domain name ownership. After the CA approves your certificate application, the CA issues the certificate. If a certificate is issued, you can call this operation to obtain the CA certificate and private key of the certificate.
 *
 * @param request DescribeCertificateStateRequest
 * @return DescribeCertificateStateResponse
 */
func (client *Client) DescribeCertificateState(request *DescribeCertificateStateRequest) (_result *DescribeCertificateStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCertificateStateResponse{}
	_body, _err := client.DescribeCertificateStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudResourceStatusWithOptions(request *DescribeCloudResourceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudResourceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecretId)) {
		query["SecretId"] = request.SecretId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudResourceStatus"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudResourceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudResourceStatus(request *DescribeCloudResourceStatusRequest) (_result *DescribeCloudResourceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudResourceStatusResponse{}
	_body, _err := client.DescribeCloudResourceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeploymentJobWithOptions(request *DescribeDeploymentJobRequest, runtime *util.RuntimeOptions) (_result *DescribeDeploymentJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeploymentJob"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeploymentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeploymentJob(request *DescribeDeploymentJobRequest) (_result *DescribeDeploymentJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeploymentJobResponse{}
	_body, _err := client.DescribeDeploymentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeploymentJobStatusWithOptions(request *DescribeDeploymentJobStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDeploymentJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeploymentJobStatus"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeploymentJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeploymentJobStatus(request *DescribeDeploymentJobStatusRequest) (_result *DescribeDeploymentJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeploymentJobStatusResponse{}
	_body, _err := client.DescribeDeploymentJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackageStateWithOptions(request *DescribePackageStateRequest, runtime *util.RuntimeOptions) (_result *DescribePackageStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePackageState"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePackageStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackageState(request *DescribePackageStateRequest) (_result *DescribePackageStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackageStateResponse{}
	_body, _err := client.DescribePackageStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EncryptWithOptions(request *EncryptRequest, runtime *util.RuntimeOptions) (_result *EncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.Plaintext)) {
		query["Plaintext"] = request.Plaintext
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Encrypt"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EncryptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Encrypt(request *EncryptRequest) (_result *EncryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EncryptResponse{}
	_body, _err := client.EncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCertWarehouseQuotaWithOptions(runtime *util.RuntimeOptions) (_result *GetCertWarehouseQuotaResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetCertWarehouseQuota"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCertWarehouseQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCertWarehouseQuota() (_result *GetCertWarehouseQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCertWarehouseQuotaResponse{}
	_body, _err := client.GetCertWarehouseQuotaWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCsrDetailWithOptions(request *GetCsrDetailRequest, runtime *util.RuntimeOptions) (_result *GetCsrDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CsrId)) {
		query["CsrId"] = request.CsrId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCsrDetail"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCsrDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCsrDetail(request *GetCsrDetailRequest) (_result *GetCsrDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCsrDetailResponse{}
	_body, _err := client.GetCsrDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request GetUserCertificateDetailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetUserCertificateDetailResponse
 */
func (client *Client) GetUserCertificateDetailWithOptions(request *GetUserCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *GetUserCertificateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertFilter)) {
		query["CertFilter"] = request.CertFilter
	}

	if !tea.BoolValue(util.IsUnset(request.CertId)) {
		query["CertId"] = request.CertId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserCertificateDetail"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request GetUserCertificateDetailRequest
 * @return GetUserCertificateDetailResponse
 */
func (client *Client) GetUserCertificateDetail(request *GetUserCertificateDetailRequest) (_result *GetUserCertificateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserCertificateDetailResponse{}
	_body, _err := client.GetUserCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ListCertRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListCertResponse
 */
func (client *Client) ListCertWithOptions(request *ListCertRequest, runtime *util.RuntimeOptions) (_result *ListCertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		query["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WarehouseId)) {
		query["WarehouseId"] = request.WarehouseId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCert"),
		Version:     tea.String("2020-04-07"),
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

/**
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ListCertRequest
 * @return ListCertResponse
 */
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

func (client *Client) ListCertWarehouseWithOptions(request *ListCertWarehouseRequest, runtime *util.RuntimeOptions) (_result *ListCertWarehouseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCertWarehouse"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCertWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCertWarehouse(request *ListCertWarehouseRequest) (_result *ListCertWarehouseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCertWarehouseResponse{}
	_body, _err := client.ListCertWarehouseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCloudAccessWithOptions(request *ListCloudAccessRequest, runtime *util.RuntimeOptions) (_result *ListCloudAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloudName)) {
		query["CloudName"] = request.CloudName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.SecretId)) {
		query["SecretId"] = request.SecretId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCloudAccess"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCloudAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCloudAccess(request *ListCloudAccessRequest) (_result *ListCloudAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCloudAccessResponse{}
	_body, _err := client.ListCloudAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCloudResourcesWithOptions(request *ListCloudResourcesRequest, runtime *util.RuntimeOptions) (_result *ListCloudResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloudName)) {
		query["CloudName"] = request.CloudName
	}

	if !tea.BoolValue(util.IsUnset(request.CloudProduct)) {
		query["CloudProduct"] = request.CloudProduct
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.SecretId)) {
		query["SecretId"] = request.SecretId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCloudResources"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCloudResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCloudResources(request *ListCloudResourcesRequest) (_result *ListCloudResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCloudResourcesResponse{}
	_body, _err := client.ListCloudResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContactWithOptions(request *ListContactRequest, runtime *util.RuntimeOptions) (_result *ListContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListContact"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContact(request *ListContactRequest) (_result *ListContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContactResponse{}
	_body, _err := client.ListContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCsrWithOptions(request *ListCsrRequest, runtime *util.RuntimeOptions) (_result *ListCsrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCsr"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCsr(request *ListCsrRequest) (_result *ListCsrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCsrResponse{}
	_body, _err := client.ListCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeploymentJobWithOptions(request *ListDeploymentJobRequest, runtime *util.RuntimeOptions) (_result *ListDeploymentJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeploymentJob"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeploymentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeploymentJob(request *ListDeploymentJobRequest) (_result *ListDeploymentJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeploymentJobResponse{}
	_body, _err := client.ListDeploymentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeploymentJobCertWithOptions(request *ListDeploymentJobCertRequest, runtime *util.RuntimeOptions) (_result *ListDeploymentJobCertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeploymentJobCert"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeploymentJobCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeploymentJobCert(request *ListDeploymentJobCertRequest) (_result *ListDeploymentJobCertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeploymentJobCertResponse{}
	_body, _err := client.ListDeploymentJobCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeploymentJobResourceWithOptions(request *ListDeploymentJobResourceRequest, runtime *util.RuntimeOptions) (_result *ListDeploymentJobResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeploymentJobResource"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeploymentJobResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeploymentJobResource(request *ListDeploymentJobResourceRequest) (_result *ListDeploymentJobResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeploymentJobResourceResponse{}
	_body, _err := client.ListDeploymentJobResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ListUserCertificateOrder operation to query the certificates or certificate orders of users. If you set OrderType to CERT or UPLOAD, certificates are returned. If you set OrderType to CPACK or BUY, certificate orders are returned.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ListUserCertificateOrderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListUserCertificateOrderResponse
 */
func (client *Client) ListUserCertificateOrderWithOptions(request *ListUserCertificateOrderRequest, runtime *util.RuntimeOptions) (_result *ListUserCertificateOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserCertificateOrder"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserCertificateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ListUserCertificateOrder operation to query the certificates or certificate orders of users. If you set OrderType to CERT or UPLOAD, certificates are returned. If you set OrderType to CPACK or BUY, certificate orders are returned.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ListUserCertificateOrderRequest
 * @return ListUserCertificateOrderResponse
 */
func (client *Client) ListUserCertificateOrder(request *ListUserCertificateOrderRequest) (_result *ListUserCertificateOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserCertificateOrderResponse{}
	_body, _err := client.ListUserCertificateOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkerResourceWithOptions(request *ListWorkerResourceRequest, runtime *util.RuntimeOptions) (_result *ListWorkerResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloudProduct)) {
		query["CloudProduct"] = request.CloudProduct
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkerResource"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkerResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkerResource(request *ListWorkerResourceRequest) (_result *ListWorkerResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkerResourceResponse{}
	_body, _err := client.ListWorkerResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to submit a renewal application for a certificate only when the order of the certificate is in the expiring state. After the renewal is complete, a new certificate order whose status is pending application is generated. You must submit a certificate application for the new certificate order and install the new certificate after the new certificate is issued.
 * > You can call the [DescribeCertificateState](~~455800~~) operation to query the status of a certificate application order. If the value of the **Type** response parameter is **certificate**, the certificate is issued.
 *
 * @param request RenewCertificateOrderForPackageRequestRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RenewCertificateOrderForPackageRequestResponse
 */
func (client *Client) RenewCertificateOrderForPackageRequestWithOptions(request *RenewCertificateOrderForPackageRequestRequest, runtime *util.RuntimeOptions) (_result *RenewCertificateOrderForPackageRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Csr)) {
		query["Csr"] = request.Csr
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewCertificateOrderForPackageRequest"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewCertificateOrderForPackageRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to submit a renewal application for a certificate only when the order of the certificate is in the expiring state. After the renewal is complete, a new certificate order whose status is pending application is generated. You must submit a certificate application for the new certificate order and install the new certificate after the new certificate is issued.
 * > You can call the [DescribeCertificateState](~~455800~~) operation to query the status of a certificate application order. If the value of the **Type** response parameter is **certificate**, the certificate is issued.
 *
 * @param request RenewCertificateOrderForPackageRequestRequest
 * @return RenewCertificateOrderForPackageRequestResponse
 */
func (client *Client) RenewCertificateOrderForPackageRequest(request *RenewCertificateOrderForPackageRequestRequest) (_result *RenewCertificateOrderForPackageRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewCertificateOrderForPackageRequestResponse{}
	_body, _err := client.RenewCertificateOrderForPackageRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeWHClientCertificateWithOptions(request *RevokeWHClientCertificateRequest, runtime *util.RuntimeOptions) (_result *RevokeWHClientCertificateResponse, _err error) {
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
		Action:      tea.String("RevokeWHClientCertificate"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeWHClientCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeWHClientCertificate(request *RevokeWHClientCertificateRequest) (_result *RevokeWHClientCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeWHClientCertificateResponse{}
	_body, _err := client.RevokeWHClientCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SignWithOptions(request *SignRequest, runtime *util.RuntimeOptions) (_result *SignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.SigningAlgorithm)) {
		query["SigningAlgorithm"] = request.SigningAlgorithm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Sign"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Sign(request *SignRequest) (_result *SignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SignResponse{}
	_body, _err := client.SignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCsrWithOptions(request *UpdateCsrRequest, runtime *util.RuntimeOptions) (_result *UpdateCsrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CsrId)) {
		query["CsrId"] = request.CsrId
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCsr"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCsr(request *UpdateCsrRequest) (_result *UpdateCsrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCsrResponse{}
	_body, _err := client.UpdateCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeploymentJobWithOptions(request *UpdateDeploymentJobRequest, runtime *util.RuntimeOptions) (_result *UpdateDeploymentJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIds)) {
		query["CertIds"] = request.CertIds
	}

	if !tea.BoolValue(util.IsUnset(request.ContactIds)) {
		query["ContactIds"] = request.ContactIds
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleTime)) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDeploymentJob"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeploymentJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeploymentJob(request *UpdateDeploymentJobRequest) (_result *UpdateDeploymentJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeploymentJobResponse{}
	_body, _err := client.UpdateDeploymentJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeploymentJobStatusWithOptions(request *UpdateDeploymentJobStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateDeploymentJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDeploymentJobStatus"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeploymentJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeploymentJobStatus(request *UpdateDeploymentJobStatusRequest) (_result *UpdateDeploymentJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeploymentJobStatusResponse{}
	_body, _err := client.UpdateDeploymentJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkerResourceStatusWithOptions(request *UpdateWorkerResourceStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkerResourceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerId)) {
		query["WorkerId"] = request.WorkerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkerResourceStatus"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkerResourceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkerResourceStatus(request *UpdateWorkerResourceStatusRequest) (_result *UpdateWorkerResourceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkerResourceStatusResponse{}
	_body, _err := client.UpdateWorkerResourceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadCsrWithOptions(request *UploadCsrRequest, runtime *util.RuntimeOptions) (_result *UploadCsrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Csr)) {
		query["Csr"] = request.Csr
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadCsr"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadCsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadCsr(request *UploadCsrRequest) (_result *UploadCsrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadCsrResponse{}
	_body, _err := client.UploadCsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The unique identifier of the certificate.
 *
 * @param request UploadPCACertRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UploadPCACertResponse
 */
func (client *Client) UploadPCACertWithOptions(request *UploadPCACertRequest, runtime *util.RuntimeOptions) (_result *UploadPCACertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cert)) {
		query["Cert"] = request.Cert
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.WarehouseId)) {
		query["WarehouseId"] = request.WarehouseId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadPCACert"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadPCACertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The unique identifier of the certificate.
 *
 * @param request UploadPCACertRequest
 * @return UploadPCACertResponse
 */
func (client *Client) UploadPCACert(request *UploadPCACertRequest) (_result *UploadPCACertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadPCACertResponse{}
	_body, _err := client.UploadPCACertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UploadUserCertificateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UploadUserCertificateResponse
 */
func (client *Client) UploadUserCertificateWithOptions(request *UploadUserCertificateRequest, runtime *util.RuntimeOptions) (_result *UploadUserCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cert)) {
		query["Cert"] = request.Cert
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptCert)) {
		query["EncryptCert"] = request.EncryptCert
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptPrivateKey)) {
		query["EncryptPrivateKey"] = request.EncryptPrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SignCert)) {
		query["SignCert"] = request.SignCert
	}

	if !tea.BoolValue(util.IsUnset(request.SignPrivateKey)) {
		query["SignPrivateKey"] = request.SignPrivateKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadUserCertificate"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadUserCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UploadUserCertificateRequest
 * @return UploadUserCertificateResponse
 */
func (client *Client) UploadUserCertificate(request *UploadUserCertificateRequest) (_result *UploadUserCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadUserCertificateResponse{}
	_body, _err := client.UploadUserCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyWithOptions(request *VerifyRequest, runtime *util.RuntimeOptions) (_result *VerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureValue)) {
		query["SignatureValue"] = request.SignatureValue
	}

	if !tea.BoolValue(util.IsUnset(request.SigningAlgorithm)) {
		query["SigningAlgorithm"] = request.SigningAlgorithm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Verify"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Verify(request *VerifyRequest) (_result *VerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyResponse{}
	_body, _err := client.VerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
