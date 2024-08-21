// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelCertificateForPackageRequestRequest struct {
	// The order ID.
	//
	// >  You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123451222
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
	// The request ID.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
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
	// The order ID.
	//
	// >  You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123451222
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
	// The request ID.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
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
	// >  This parameter is available only when you apply for OV certificates. For more information, see [Manage company profiles](https://help.aliyun.com/document_detail/198289.html). If you want to apply for a DV certificate, you do not need to add a company profile.
	//
	// If you specify a company name, the information about the company that is configured in the **Information Management*	- module is used. If you do not specify this parameter, the information about the most recent company that is added to the **Information Management*	- module is used.
	//
	// example:
	//
	// A company
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// The content of the certificate signing request (CSR) file that is manually generated by using OpenSSL or Keytool for the domain name. The key algorithm in the CSR file must be Rivest-Shamir-Adleman (RSA) or elliptic-curve cryptography (ECC), and the key length of the RSA algorithm must be greater than or equal to 2,048 characters. For more information about how to create a CSR file, see [Create a CSR file](https://help.aliyun.com/document_detail/313297.html). If you do not specify this parameter, Certificate Management Service automatically creates a CSR file.
	//
	// A CSR file contains the information about your server and company. When you apply for a certificate, you must submit the CSR file to the CA. The CA signs the CSR file by using the private key of the root certificate and generates a public key file to issue your certificate.
	//
	// >  The \\*\\*CN\\*\\	- field in CSR file specifies the domain name that is bound to the certificate. You must include the field in the parameter value.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The domain name that you want to bind to the certificate. The domain name must meet the following requirements:
	//
	// 	- The domain name must be a single domain name or a wildcard domain name. Example: `*.aliyundoc.com`.
	//
	// 	- You can specify multiple domain names. Separate multiple domain names with commas (,). You can specify a maximum of five domain names.
	//
	// 	- If you specify multiple domain names, the domain names must be only single domain names or only wildcard domain names. You cannot specify both single domain names and wildcard domain names.
	//
	// >  If you want to bind multiple domain names to the certificate, you must specify this parameter. You must specify at least one of the Domain parameter and the \\*\\*Csr\\*\\	- parameter. If you specify both the Domain parameter and the \\*\\*Csr\\*\\	- parameter, the value of the \\*\\*CN\\*\\	- field in the \\*\\*Csr\\*\\	- parameter is used as the domain name that is bound to the certificate.
	//
	// example:
	//
	// aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The email address of the applicant. After the CA receives your certificate application, the CA sends a verification email to the email address that you specify. You must log on to the mailbox, open the mail, and complete the verification of the domain name ownership based on the steps that are described in the email.
	//
	// If you do not specify this parameter, the information about the most recent contact that is added to the **Information Management*	- module is used. For more information about how to add a contact to the **Information Management*	- module, see [Manage contacts](https://help.aliyun.com/document_detail/198262.html).
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The contact phone number of the applicant. CA staff can call the phone number to confirm the information in your certificate application.
	//
	// If you do not specify this parameter, the information about the most recent contact that is added to the **Information Management*	- module is used. For more information about how to add a contact to the **Information Management*	- module, see [Manage contacts](https://help.aliyun.com/document_detail/198262.html).
	//
	// example:
	//
	// 1390000****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The specifications of the certificate that you want to apply for. Valid values:
	//
	// 	- **digicert-free-1-free*	- (default): DigiCert single-domain domain validated (DV) certificate, which is free and valid for 3 months. This value is available only on the China site (aliyun.com).
	//
	// 	- **symantec-free-1-free**: DigiCert single-domain DV certificate, which is free and valid for 1 year. This value is available only on the China site (aliyun.com).
	//
	// 	- **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	//
	// 	- **symantec-ov-1-personal**: DigiCert single-domain organization validated (OV) certificate.
	//
	// 	- **symantec-ov-w-personal**: DigiCert wildcard OV certificate.
	//
	// 	- **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	//
	// 	- **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	//
	// 	- **geotrust-ov-1-personal**: GeoTrust single-domain OV certificate.
	//
	// 	- **geotrust-ov-w-personal**: GeoTrust wildcard OV certificate.
	//
	// 	- **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	//
	// 	- **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	//
	// 	- **globalsign-ov-1-personal**: GlobalSign single-domain OV certificate.
	//
	// 	- **globalsign-ov-w-advanced**: GlobalSign wildcard OV certificate.
	//
	// 	- **cfca-ov-1-personal**: China Financial Certification Authority (CFCA) single-domain OV certificate, available only on the China site (aliyun.com).
	//
	// 	- **cfca-ev-w-advanced**: CFCA wildcard OV certificate, available only on the China site (aliyun.com).
	//
	// example:
	//
	// symantec-free-1-free
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the applicant.
	//
	// If you do not specify this parameter, the information about the most recent contact that is added to the **Information Management*	- module is used. For more information about how to add a contact to the **Information Management*	- module, see [Manage contacts](https://help.aliyun.com/document_detail/198262.html).
	//
	// example:
	//
	// Tom
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The verification method of the domain name ownership. Valid values:
	//
	// 	- **DNS**: DNS verification. If you use this method, you must add a TXT record to the DNS records of the domain name in the management platform of the domain name. You must have operation permissions on domain name resolution to verify the ownership of the domain name.
	//
	// 	- **FILE**: file verification. If you use this method, you must create a specified file on the DNS server. You must have administrative rights on the DNS server to verify the ownership of the domain name.
	//
	// For more information about the verification methods, see [Verify the ownership of a domain name](https://help.aliyun.com/document_detail/48016.html).
	//
	// example:
	//
	// DNS
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
	// >  You can use the ID to query the status of the certificate application order. For more information, see [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html).
	//
	// example:
	//
	// 2021010
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5890029B-938A-589E-98B9-3DEC7BA7C400
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
	// >  The domain name must match the certificate specifications that you specify for the **ProductCode*	- parameter. If you apply for a single-domain certificate, you must specify a single domain name for this parameter. If you apply for a wildcard certificate, you must specify a wildcard domain name such as `*.aliyundoc.com` for this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The contact email address of the applicant.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The phone number of the applicant.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The specifications of the certificate. Valid values:
	//
	// 	- **digicert-free-1-free*	- (default): DigiCert single-domain DV certificate, which is free and valid for 3 months.
	//
	// 	- **symantec-free-1-free**: DigiCert single-domain DV certificate, which is free and valid for 1 year. This value is available only on the China site (aliyun.com).
	//
	// 	- **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	//
	// 	- **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	//
	// 	- **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	//
	// 	- **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	//
	// 	- **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	//
	// example:
	//
	// symantec-free-1-free
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the applicant.
	//
	// This parameter is required.
	//
	// example:
	//
	// Tom
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The method to verify the ownership of a domain name. Valid values:
	//
	// 	- **DNS**: DNS verification. If you use this method, you must add a TXT record to the DNS records of the domain name in the management platform of the domain name. You must have operation permissions on domain name resolution to verify the ownership of the domain name.
	//
	// 	- **FILE**: file verification. If you use this method, you must create a specified file on the DNS server. You must have administrative rights on the DNS server to verify the ownership of the domain name.
	//
	// For more information about the verification methods, see [Verify the ownership of a domain name](https://help.aliyun.com/document_detail/48016.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// DNS
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
	// >  You can use the ID to query the status of the certificate application. For more information, see [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html).
	//
	// example:
	//
	// 98987582437920968
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
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
	// The content of the CSR file.\\
	//
	// The key algorithm in the CSR file must be Rivest-Shamir-Adleman (RSA) or elliptic-curve cryptography (ECC), and the key length of the RSA algorithm must be greater than or equal to 2,048 characters. For more information about how to create a CSR file, see [How do I create a CSR file?](https://help.aliyun.com/document_detail/42218.html)\\
	//
	// A CSR file contains the information about your server and company. When you apply for a certificate, you must submit the CSR file to the CA. The CA signs the CSR file by using the private key of the root certificate and generates a public key file to issue your certificate.
	//
	// >  The **CN*	- field in the CSR file specifies the domain name that is bound to the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The contact email address of the applicant.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The phone number of the applicant.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The specifications of the certificate that you want to apply for. Valid values:
	//
	// 	- **digicert-free-1-free*	- (default): DigiCert single-domain DV certificate in a three-month free trial, available only on the China site (aliyun.com).
	//
	// 	- **symantec-free-1-free**: DigiCert single-domain DV certificate in a one-year free trial, available only on the China site (aliyun.com).
	//
	// 	- **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	//
	// 	- **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	//
	// 	- **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	//
	// 	- **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	//
	// 	- **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	//
	// example:
	//
	// symantec-free-1-free
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the applicant.
	//
	// This parameter is required.
	//
	// example:
	//
	// Tom
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The method to verify the ownership of a domain name. Valid values:
	//
	// 	- **DNS**: DNS verification. If you use this method, you must add a TXT record to the DNS records of the domain name in the management platform of the domain name. You must have operation permissions on domain name resolution to verify the ownership of the domain name.
	//
	// 	- **FILE**: file verification. If you use this method, you must create a specified file on the DNS server. You must have administrative rights on the DNS server to verify the ownership of the domain name.
	//
	// For more information about the verification methods, see [Verify the ownership of a domain name](https://help.aliyun.com/document_detail/48016.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// DNS
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
	// >  You can use the ID to query the status of the certificate application. For more information, see [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html).
	//
	// example:
	//
	// 98987582437920968
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
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
	// The content of the CSR.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The unique identifier of the CSR. You can use this value to obtain the content of the CSR. For more information about the operation that you can call to obtain the content of a CSR, see [GetCsrDetail](https://help.aliyun.com/document_detail/2709720.html).
	//
	// example:
	//
	// 3365
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
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
	// The certificate IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12342649,12304338,12067351,9957285
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// The contact IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	// The type of the task.
	//
	// Valid values:
	//
	// 	- cloud
	//
	// 	- user
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// jobName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6591316,6585549,6190248,5811894,5775176,5772504
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The time when the task was scheduled.
	//
	// example:
	//
	// 1706613560008
	ScheduleTime *int64 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
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
	// The ID of the deployment task.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
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
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// example:
	//
	// aliyun
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// example:
	//
	// 365
	Days *int64 `json:"Days,omitempty" xml:"Days,omitempty"`
	// example:
	//
	// 1
	Immediately *int64 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// example:
	//
	// 12
	Months       *int64  `json:"Months,omitempty" xml:"Months,omitempty"`
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 273ae6bb538d538c70c01f81jh2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// example:
	//
	// 2
	SanType *int64 `json:"SanType,omitempty" xml:"SanType,omitempty"`
	// example:
	//
	// example.com
	SanValue *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1
	Years *int64 `json:"Years,omitempty" xml:"Years,omitempty"`
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
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// example:
	//
	// 190ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	ParentX509Certificate *string `json:"ParentX509Certificate,omitempty" xml:"ParentX509Certificate,omitempty"`
	// example:
	//
	// 8C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	RootX509Certificate *string `json:"RootX509Certificate,omitempty" xml:"RootX509Certificate,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
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
	// The encryption algorithm. Valid values:
	//
	// 	- **RSAES_OAEP_SHA_1**
	//
	// 	- **RSAES_OAEP_SHA_256**
	//
	// 	- **SM2PKE**
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAESOAEPSHA_1
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The unique identifier of the certificate. You can call the [ListCert](https://help.aliyun.com/document_detail/455806.html) operation to query the identifier.
	//
	// 	- If the certificate is an SSL certificate, the value of this parameter must be in the {Certificate ID}-cn-hangzhou format.
	//
	// 	- If the certificate is a private certificate, the value of this parameter must be the value of the Identifier field for the private certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The data that you want to decrypt. The value is encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The value type of the Message parameter. Valid values:
	//
	// 	- RAW: The returned result is raw data encoded in UTF-8.
	//
	// 	- Base64: The returned result is Base64-encoded data. This is the default value.
	//
	// example:
	//
	// Base64
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
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
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The data after decryption.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// >  After you call the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html), [CreateCertificateRequest](https://help.aliyun.com/document_detail/455292.html), or [CreateCertificateWithCsrRequest](https://help.aliyun.com/document_detail/455801.html) operation to submit a certificate application, you can obtain the ID of the certificate application order from the **OrderId*	- response parameter. You can also call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the order ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123451222
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
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
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
	// The ID of the CSR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3013
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
	// The request ID.
	//
	// example:
	//
	// D3F1FA43-1C26-50A2-8F0F-7A03851DBB46
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
	// The ID of the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
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
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
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
	// This parameter is required.
	//
	// example:
	//
	// ccaf0c629c2be1e2ab
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
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
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
	//
	// >  You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7562353
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
	//
	// example:
	//
	// 3E50D480-9765-5CFD-9650-9BACCECA5135
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
	// The ID of the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the worker for the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13
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
	// The request ID.
	//
	// example:
	//
	// EA69E364-5CBB-50E8-BF09-E8CAA396A4F8
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
	// >  You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123451222
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
	// The content of the certificate in the PEM format. For more information about the PEM format and how to convert certificate formats, see [What formats are used for mainstream digital certificates?](https://help.aliyun.com/document_detail/42214.html)
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **certificate**. The value certificate indicates that the certificate is issued.
	//
	// example:
	//
	// ——BEGIN CERTIFICATE—— …… ——END CERTIFICATE——
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The content that you need to write to the newly created file when you use the file verification method.
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **FILE**. The value domain_verify indicates that the verification of the domain name ownership is not complete, and the value FILE indicates that the file verification method is used.
	//
	// example:
	//
	// http://example.com/.well-known/pki-validation/fileauth.txt
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The domain name to be verified when you use the file verification method. You must connect to the DNS server of the domain name and create a file on the server. The file is specified by the **Uri*	- parameter.
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **FILE**. The value domain_verify indicates that the verification of the domain name ownership is not complete, and the value FILE indicates that the file verification method is used.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The private key of the certificate in the PEM format. For more information about the PEM format and how to convert certificate formats, see [What formats are used for mainstream digital certificates?](https://help.aliyun.com/document_detail/42214.html)
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **certificate**. The value certificate indicates that the certificate is issued.
	//
	// example:
	//
	// ——BEGIN RSA PRIVATE KEY—— …… ——END RSA PRIVATE KEY——
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The DNS record that you need to manage when you use the DNS verification method.
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **DNS**. The value domain_verify indicates that the verification of the domain name ownership is not complete, and the value DNS indicates that the DNS verification method is used.
	//
	// example:
	//
	// _dnsauth
	RecordDomain *string `json:"RecordDomain,omitempty" xml:"RecordDomain,omitempty"`
	// The type of the DNS record that you need to add when you use the DNS verification method. Valid values:
	//
	// 	- **TXT**
	//
	// 	- **CNAME**
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **DNS**. The value domain_verify indicates that the verification of the domain name ownership is not complete.
	//
	// example:
	//
	// TXT
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// You need to add a TXT record to the DNS records only when you use the DNS verification method.
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **DNS**. The value domain_verify indicates that the verification of the domain name ownership is not complete, and the value DNS indicates that the DNS verification method is used.
	//
	// example:
	//
	// 20200420000000223erigacv46uhaubchcm0o7spxi7i2isvjq59mlx9lucnkqcy
	RecordValue *string `json:"RecordValue,omitempty" xml:"RecordValue,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the certificate application order. Valid values:
	//
	// 	- **domain_verify**: **pending review**, which indicates that you have not completed the verification of the domain name ownership after you submit the certificate application.
	//
	//      >After you submit a certificate application, you must manually complete the verification of the domain name ownership. The CA reviews the certificate application only after the verification is complete. If you have not completed the verification of the domain name ownership, you can complete the verification based on the data returned by this operation.
	//
	// 	- **process**: **being reviewed**, which indicates that the certificate application is being reviewed by the CA.
	//
	// 	- **verify_fail**: **review failed**, which indicates that the certificate application failed to be reviewed.
	//
	//     >  If a certificate application fails to be reviewed, the information that you specified in the certificate application may be incorrect. We recommend that you call the [DeleteCertificateRequest](https://help.aliyun.com/document_detail/164109.html) operation to delete the certificate application order and resubmit a certificate application. After the order is deleted, the quota that is consumed for the order is released.
	//
	// 	- **certificate**: **issued**, which indicates that the certificate is issued.
	//
	// 	- **payed**: **pending application**, which indicates that you have not submitted a certificate application.
	//
	// 	- **unknow**: The status is **unknown**.
	//
	// example:
	//
	// domain_verify
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The file that you need to create on the DNS server when you use the file verification method. **The value of this parameter contains the file path and file name.**
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify*	- and the value of the **ValidateType*	- parameter is **FILE**. The value domain_verify indicates that the verification of the domain name ownership is not complete, and the value FILE indicates that the file verification method is used.
	//
	// example:
	//
	// /.well-known/pki-validation/fileauth.txt
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The verification method of the domain name ownership. Valid values:
	//
	// 	- **DNS**: DNS verification. If you use this method, you must add a TXT record to the DNS records of the domain name in the management platform of the domain name.
	//
	// 	- **FILE**: file verification. If you use this method, you must create a specified file on the DNS server.
	//
	// >  This parameter is returned only when the value of the **Type*	- parameter is **domain_verify**. The value domain_verify indicates that the verification of the domain name ownership is not complete.
	//
	// example:
	//
	// FILE
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
	// The AccessKey secret used to access cloud resources.
	//
	// >  You can call the [ListCloudAccess](https://help.aliyun.com/document_detail/2712219.html) operation to obtain the ID.
	//
	// example:
	//
	// AKID9*******XX
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
	// The response parameters.
	Data []*DescribeCloudResourceStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 09470F19-CEE8-5C63-BF2C-02B5E3F07A17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The cloud service provider.
	//
	// example:
	//
	// aliyun
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service.
	//
	// example:
	//
	// OSS
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The total number of cloud resources on which certificates are deployed.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The ID of the deployment job. The **ID*	- of the job is returned after you call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) operation. You can also call the [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
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
	// The information about the contact.
	CasContacts []*DescribeDeploymentJobResponseBodyCasContacts `json:"CasContacts,omitempty" xml:"CasContacts,omitempty" type:"Repeated"`
	// The domain names bound to the certificate of the deployment task.
	//
	// example:
	//
	// example.aliyundoc.com,demo.aliyundoc.com
	CertDomain *string `json:"CertDomain,omitempty" xml:"CertDomain,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **upload**: uploaded certificate
	//
	// 	- **buy**: purchased certificate
	//
	// 	- **free**: free certificate available only on the China site (aliyun.com)
	//
	// example:
	//
	// buy
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The configurations of the deployment task.
	//
	// example:
	//
	// {\\"shareCertIds\\":[],\\"certIds\\":[12342649,12304338,12067351,9957285]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Indicates whether the deployment job was deleted. Valid values:
	//
	// 	- **0**: not deleted
	//
	// 	- **1**: deleted
	//
	// example:
	//
	// 1
	Del *int32 `json:"Del,omitempty" xml:"Del,omitempty"`
	// The end time of the deployment job. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679541809000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the deployment job was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679541809000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the deployment job was last modified. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679541809000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the deployment job.
	//
	// example:
	//
	// 8888
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the deployment task.
	//
	// example:
	//
	// 14dcc8afc7578e1f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the deployment job. Valid values:
	//
	// 	- **cloud**: multi-cloud deployment job.
	//
	// 	- **trustee**: hosted deployment job available only on the China site (aliyun.com).
	//
	// 	- **user**: cloud service deployment job. The cloud server is not included.
	//
	// example:
	//
	// user
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the deployment task.
	//
	// example:
	//
	// auto-test-AXX
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cloud services included in the deployment task.
	//
	// example:
	//
	// CDN
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the deployment job includes the rollback worker. For example, if a cloud service in a deployment job has been rolled back, **1*	- is returned. Valid values:
	//
	// 	- **0**: The rollback worker is not included.
	//
	// 	- **1**: The rollback worker is included.
	//
	// example:
	//
	// 1
	Rollback *int32 `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
	// The time when the deployment job was scheduled. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1678083209335
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The start time of the deployment job. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679541809000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the deployment job. Valid values:
	//
	// 	- **pending**
	//
	// 	- **editing**
	//
	// 	- **scheduling**
	//
	// 	- **processing**
	//
	// 	- **error**
	//
	// 	- **success**
	//
	// example:
	//
	// editing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Alibaba Cloud account in which the deployment job is created.
	//
	// example:
	//
	// 166688437XXXX785
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The email address of the contact.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the contact.
	//
	// example:
	//
	// 3304
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The phone number of the contact.
	//
	// example:
	//
	// 139****0000
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The name of the contact.
	//
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The ID of the deployment task. You can call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) operation to obtain the ID of a deployment task from the **JobId*	- parameter. You can also call the [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation to obtain the ID of a deployment task.
	//
	// example:
	//
	// 8888
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
	// The total number of purchased resources.
	//
	// example:
	//
	// 500
	BuyCount *int32 `json:"BuyCount,omitempty" xml:"BuyCount,omitempty"`
	// The number of certificates involved in the deployment task.
	//
	// example:
	//
	// 20
	CertCount *int32 `json:"CertCount,omitempty" xml:"CertCount,omitempty"`
	// The number of resources consumed by worker tasks.
	//
	// example:
	//
	// 20
	CostCount *int32 `json:"CostCount,omitempty" xml:"CostCount,omitempty"`
	// The number of failed worker tasks, excluding rollback tasks.
	//
	// example:
	//
	// 20
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The total number of worker tasks that match the certificate.
	//
	// example:
	//
	// 20
	MatchWorkerCount *int32 `json:"MatchWorkerCount,omitempty" xml:"MatchWorkerCount,omitempty"`
	// The number of cloud resources to which certificates are deployed in the deployment task.
	ProductWorkerCount []*DescribeDeploymentJobStatusResponseBodyProductWorkerCount `json:"ProductWorkerCount,omitempty" xml:"ProductWorkerCount,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of resources.
	//
	// example:
	//
	// 4127
	ResourceCount *int32 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// The number of worker tasks that are rolled backed.
	//
	// example:
	//
	// 20
	RollbackCount *int32 `json:"RollbackCount,omitempty" xml:"RollbackCount,omitempty"`
	// The number of worker tasks that failed to be rolled back.
	//
	// example:
	//
	// 20
	RollbackFailedCount *int32 `json:"RollbackFailedCount,omitempty" xml:"RollbackFailedCount,omitempty"`
	// The number of worker tasks that are successfully rolled back.
	//
	// example:
	//
	// 20
	RollbackSuccessCount *int32 `json:"RollbackSuccessCount,omitempty" xml:"RollbackSuccessCount,omitempty"`
	// The number of successful worker tasks, excluding rollback tasks.
	//
	// example:
	//
	// 20
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The total number of used resources.
	//
	// example:
	//
	// 300
	UseCount *int32 `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
	// The total number of resources to which certificate are deployed in the current cloud service. The value indicates the total number of worker tasks.
	//
	// example:
	//
	// 20
	WorkerCount *int32 `json:"WorkerCount,omitempty" xml:"WorkerCount,omitempty"`
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
	// The total number of resources of a cloud service in the deployment task.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the cloud service. Valid values:
	//
	// 	- **SLB**: Classic Load Balancer (CLB). This value is supported only at the China site (aliyun.com).
	//
	// 	- **LIVE**: ApsaraVideo Live. This value is supported only at the China site (aliyun.com).
	//
	// 	- **webHosting**: Cloud Web Hosting. This value is supported only at the China site (aliyun.com).
	//
	// 	- **VOD**: ApsaraVideo VOD. This value is supported only at the China site (aliyun.com).
	//
	// 	- **CR**: Container Registry. This value is supported only at the China site (aliyun.com).
	//
	// 	- **DCDN**: Dynamic Content Delivery Network (DCDN).
	//
	// 	- **DDOS**: Anti-DDoS.
	//
	// 	- **CDN**: Alibaba Cloud CDN (CDN).
	//
	// 	- **ALB**: Application Load Balancer (ALB).
	//
	// 	- **APIGateway**: API Gateway.
	//
	// 	- **FC**: Function Compute.
	//
	// 	- **GA**: Global Accelerator (GA).
	//
	// 	- **MSE**: Microservices Engine (MSE).
	//
	// 	- **NLB**: Network Load Balancer (NLB).
	//
	// 	- **OSS**: Object Storage Service (OSS).
	//
	// 	- **SAE**: Serverless App Engine (SAE).
	//
	// 	- **TencentCDN**: Tencent Cloud Content Delivery Network (CDN).
	//
	// 	- **WAF**: Web Application Firewall (WAF).
	//
	// example:
	//
	// NLB
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
	// 	- **digicert-free-1-free*	- (default): DigiCert single-domain domain validated (DV) certificate in a three-month free trial, available only on the China site (aliyun.com).
	//
	// 	- **symantec-free-1-free**: DigiCert single-domain DV certificate in a one-year free trial, available only on the China site (aliyun.com).
	//
	// 	- **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	//
	// 	- **symantec-ov-1-personal**: DigiCert single-domain organization validated (OV) certificate.
	//
	// 	- **symantec-ov-w-personal**: DigiCert wildcard OV certificate.
	//
	// 	- **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	//
	// 	- **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	//
	// 	- **geotrust-ov-1-personal**: GeoTrust single-domain OV certificate.
	//
	// 	- **geotrust-ov-w-personal**: GeoTrust wildcard OV certificate.
	//
	// 	- **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	//
	// 	- **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	//
	// 	- **globalsign-ov-1-personal**: GlobalSign single-domain OV certificate.
	//
	// 	- **globalsign-ov-w-advanced**: GlobalSign wildcard OV certificate.
	//
	// 	- **cfca-ov-1-personal**: China Financial Certification Authority (CFCA) single-domain OV certificate, available only on the China site (aliyun.com).
	//
	// 	- **cfca-ev-w-advanced**: CFCA wildcard OV certificate, available only on the China site (aliyun.com).
	//
	// example:
	//
	// symantec-free-1-free
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
	//
	// example:
	//
	// 1
	IssuedCount *int64 `json:"IssuedCount,omitempty" xml:"IssuedCount,omitempty"`
	// The specifications of the certificate resource plan. Valid values:
	//
	// 	- **digicert-free-1-free**: DigiCert single-domain DV certificate in a three-month free trial, available only on the China site (aliyun.com).
	//
	// 	- **symantec-free-1-free**: DigiCert single-domain DV certificate in a one-year free trial, available only on the China site (aliyun.com).
	//
	// 	- **symantec-dv-1-starter**: DigiCert wildcard DV certificate.
	//
	// 	- **symantec-ov-1-personal**: DigiCert single-domain OV certificate.
	//
	// 	- **symantec-ov-w-personal**: DigiCert wildcard OV certificate.
	//
	// 	- **geotrust-dv-1-starter**: GeoTrust single-domain DV certificate.
	//
	// 	- **geotrust-dv-w-starter**: GeoTrust wildcard DV certificate.
	//
	// 	- **geotrust-ov-1-personal**: GeoTrust single-domain OV certificate.
	//
	// 	- **geotrust-ov-w-personal**: GeoTrust wildcard OV certificate.
	//
	// 	- **globalsign-dv-1-personal**: GlobalSign single-domain DV certificate.
	//
	// 	- **globalsign-dv-w-advanced**: GlobalSign wildcard DV certificate.
	//
	// 	- **globalsign-ov-1-personal**: GlobalSign single-domain OV certificate.
	//
	// 	- **globalsign-ov-w-advanced**: GlobalSign wildcard OV certificate.
	//
	// 	- **cfca-ov-1-personal**: CFCA single-domain OV certificate, available only on the China site (aliyun.com).
	//
	// 	- **cfca-ev-w-advanced**: CFCA wildcard OV certificate, available only on the China site (aliyun.com).
	//
	// example:
	//
	// symantec-free-1-free
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 10CFA380-1C58-45C7-8075-06215F3DB681
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of purchased certificate resource plans of the specified specifications.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of certificate applications that you submitted for certificates of the specified specifications.
	//
	// > : A successful call of the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/204087.html), [CreateCertificateRequest](https://help.aliyun.com/document_detail/164105.html), or [CreateCertificateWithCsrRequest](https://help.aliyun.com/document_detail/178732.html) operation is counted as one a certificate application, regardless of whether the certificate is issued.
	//
	// example:
	//
	// 2
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
	// The encryption algorithm. Valid values:
	//
	// 	- **RSAES_OAEP_SHA_1**
	//
	// 	- **RSAES_OAEP_SHA_256**
	//
	// 	- **SM2PKE**
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAESOAEPSHA_1
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The unique identifier of the certificate. You can call the [ListCert](https://help.aliyun.com/document_detail/455806.html) operation to obtain the identifier.
	//
	// 	- If the certificate is an SSL certificate, the value of this parameter must be in the {Certificate ID}-cn-hangzhou format.
	//
	// 	- If the certificate is a private certificate, the value of this parameter must be the value of the Identifier field for the private certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The value type of the Message parameter. Valid values:
	//
	// 	- RAW: The value of the Plaintext parameter is directly encrypted. This is the default value.
	//
	// 	- Base64: The value of the Plaintext parameter is Base64-encoded data. The data is decoded and then encrypted.
	//
	// example:
	//
	// RAW
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The data that you want to encrypt. The value of this parameter can be raw data or Base64-encoded data. For more information, see the description of the MessageType parameter. For example, if the hexadecimal data that you want to encrypt is `[0x31, 0x32, 0x33, 0x34]`, the Base64-encoded data is MTIzNA==. The size of data that can be encrypted varies based on the encryption algorithm that you use. The following list describes the relationship between the encryption algorithms and data sizes:
	//
	// 	- **RSAES_OAEP_SHA_1**: 214 bytes
	//
	// 	- **RSAES_OAEP_SHA_256**: 190 bytes
	//
	// 	- **SM2PKE**: 6,047 bytes
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234***
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
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
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The data after encryption. The value is encoded in Base64.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total quota for certificate repositories, including the free quota and purchased quota.
	//
	// example:
	//
	// 5000
	TotalQuota *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 1000
	UseCount *int64 `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
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
	// The ID of the CSR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3924
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
	// The content of the CSR.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----   ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The private key. Specify a Base64-encoded string.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----…… -----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 08F45EA0-66A7-4504-9B31-3589F5CE308D
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

func (s *GetCsrDetailResponseBody) SetPrivateKey(v string) *GetCsrDetailResponseBody {
	s.PrivateKey = &v
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
	// Specifies whether to filter return results. Valid values: true and false. Default value: false. **true*	- specifies that the Cert, Key, EncryptCert, EncryptPrivateKey, SignCert, and SignPrivateKey parameters are not returned. **false*	- specifies that the parameters are returned.
	//
	// example:
	//
	// false
	CertFilter *bool `json:"CertFilter,omitempty" xml:"CertFilter,omitempty"`
	// The ID of the certificate.
	//
	// >  You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6055048
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
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// Indicates whether the certificate was purchased from Alibaba Cloud. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	BuyInAliyun *bool `json:"BuyInAliyun,omitempty" xml:"BuyInAliyun,omitempty"`
	// The content of the certificate if the certificate does not use an SM algorithm. If certFilter is set to false, this parameter is returned. Otherwise, this parameter is not returned.
	//
	// example:
	//
	// ---BEGIN CERTIFICATE----- MIIF...... -----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The city of the company or organization to which the certificate purchaser belongs.
	//
	// example:
	//
	// hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The primary domain name that is bound to the certificate.
	//
	// example:
	//
	// *.com
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// The country or region of the company or organization to which the certificate purchaser belongs.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the encryption certificate if the certificate uses an SM algorithm and is encoded in the PEM format. If certFilter is set to false, this parameter is returned. Otherwise, this parameter is not returned.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICDzCCA***
	//
	// -----END CERTIFICATE-----
	EncryptCert *string `json:"EncryptCert,omitempty" xml:"EncryptCert,omitempty"`
	// The private key of the encryption certificate if the certificate uses an SM algorithm and is encoded in the PEM format. If certFilter is set to false, this parameter is returned. Otherwise, this parameter is not returned.
	//
	// example:
	//
	// -----BEGIN EC PRIVATE KEY-----
	//
	// MHcCAQEEI****
	//
	// -----END EC PRIVATE KEY-----
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	// The expiration date of the certificate.
	//
	// example:
	//
	// 2023-10-25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Indicates whether the certificate has expired. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The fingerprint of the certificate.
	//
	// example:
	//
	// 1D7801BBE772D5DE55CBF1F88AEB41A42402DA07
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 121345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cas-upload-50yf1q
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// Digicert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The private key of the certificate if the certificate does not use an SM algorithm. If certFilter is set to false, this parameter is returned. Otherwise, this parameter is not returned.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- MII.... -----END RSA PRIVATE KEY-----
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// cert_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 123456
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The name of the company or organization to which the certificate purchaser belongs.
	//
	// example:
	//
	// Alibaba
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The province of the company or organization to which the certificate purchaser belongs.
	//
	// example:
	//
	// zhejiang
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// All domain names that are bound to the certificate.
	//
	// example:
	//
	// *.com
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The certificate serial No.
	//
	// example:
	//
	// 06ea4879591ddf84e6c8b6ba43607ccf
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	// The certificate sha2 value.
	//
	// example:
	//
	// 840707695D5EE41323102DDC2CB4924AA561012FBDC4E1A6324147119ED3C339
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The content of the signing certificate if the certificate uses an SM algorithm and is encoded in the PEM format. If certFilter is set to false, this parameter is returned. Otherwise, this parameter is not returned.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICDzCCAbagAw****
	//
	// -----END CERTIFICATE-----
	SignCert *string `json:"SignCert,omitempty" xml:"SignCert,omitempty"`
	// The private key of the signing certificate if the certificate uses an SM algorithm and is encoded in the PEM format. If certFilter is set to false, this parameter is returned. Otherwise, this parameter is not returned.
	//
	// example:
	//
	// -----BEGIN EC PRIVATE KEY-----
	//
	// MHcCAQEEILR****
	//
	// -----END EC PRIVATE KEY-----
	SignPrivateKey *string `json:"SignPrivateKey,omitempty" xml:"SignPrivateKey,omitempty"`
	// The issuance date of the certificate.
	//
	// example:
	//
	// 2018-07-13
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

func (s *GetUserCertificateDetailResponseBody) SetInstanceId(v string) *GetUserCertificateDetailResponseBody {
	s.InstanceId = &v
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

func (s *GetUserCertificateDetailResponseBody) SetSerialNo(v string) *GetUserCertificateDetailResponseBody {
	s.SerialNo = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetSha2(v string) *GetUserCertificateDetailResponseBody {
	s.Sha2 = &v
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
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword for the query. You can enter a name, domain name, or Subject Alternative Name (SAN) extension. Fuzzy match is supported.
	//
	// example:
	//
	// test_name
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The number of entries to return on each page. Default value: 50.
	//
	// example:
	//
	// 50
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
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
	// The ID of the certificate repository. You can call the ListCertWarehouse API operation to query the IDs of certificate repositories.
	//
	// example:
	//
	// 12
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
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID of the certificate application repository.
	//
	// example:
	//
	// 14dcc8afc7578e1f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the certificate application repository. Fuzzy match is supported.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Default value: 50.
	//
	// example:
	//
	// 50
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The type of the certificate application repository. Valid values:
	//
	// 	- **ssl**: certificate application repository of SSL certificates
	//
	// 	- **uploadPCA**: certificate application repository of uploaded private certificates
	//
	// 	- **free**: certificate application repository of free certificates, available only on the China site (aliyun.com)
	//
	// 	- **aliyunPCA**: certificate application repository of private certificates purchased from Alibaba Cloud Private Certificate Authority (PCA), available only on the China site (aliyun.com)
	//
	// 	- **disable**: disabled certificate application repository
	//
	// example:
	//
	// aliyunPCA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// The certificate application repositories.
	CertWarehouseList []*ListCertWarehouseResponseBodyCertWarehouseList `json:"CertWarehouseList,omitempty" xml:"CertWarehouseList,omitempty" type:"Repeated"`
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
	// The timestamp when the certificate application repository expires. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID of the certificate application repository.
	//
	// example:
	//
	// 14dcc8afc7578e1f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the certificate application repository has expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsExpired *bool `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// The name of the certificate application repository.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The instance ID of the private CA.
	//
	// example:
	//
	// 14dcc8afc7578e1f
	PcaInstanceId *string `json:"PcaInstanceId,omitempty" xml:"PcaInstanceId,omitempty"`
	// The queries per second (QPS).
	//
	// example:
	//
	// 10
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The type of the certificate application repository. Valid values:
	//
	// 	- **ssl**: certificate application repository of SSL certificates
	//
	// 	- **uploadPCA**: certificate application repository of uploaded private certificates
	//
	// 	- **free**: certificate application repository of free certificates, available only on the China site (aliyun.com)
	//
	// 	- **aliyunPCA**: certificate application repository of private certificates purchased from Alibaba Cloud Private Certificate Authority (PCA), available only on the China site (aliyun.com)
	//
	// 	- **disable**: disabled certificate application repository
	//
	// example:
	//
	// aliyunPCA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the certificate application repository.
	//
	// example:
	//
	// 1
	WhId *int64 `json:"WhId,omitempty" xml:"WhId,omitempty"`
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
	// The cloud service provider.
	//
	// example:
	//
	// Tencent
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The AccessKey secret used to access the cloud service.
	//
	// example:
	//
	// 276
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The number of certificates per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
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
	// The list of the AccessKey pairs.
	CloudAccessList []*ListCloudAccessResponseBodyCloudAccessList `json:"CloudAccessList,omitempty" xml:"CloudAccessList,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D3F1FA43-1C26-50A2-8F0F-7A03851DBB46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of revoked certificates per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The AccessKey ID used to access the cloud service.
	//
	// example:
	//
	// 888
	AccessId *int64 `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The cloud service provider.
	//
	// example:
	//
	// Tencent
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The AccessKey secret used to access the cloud service.
	//
	// example:
	//
	// LTAI4G5KAZCJQqdwPBAXXXX
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The status of the service.
	//
	// example:
	//
	// normal
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
	CertIds []*int64 `json:"CertIds,omitempty" xml:"CertIds,omitempty" type:"Repeated"`
	// The cloud service provider.
	//
	// Valid values:
	//
	// 	- Tencent: Tencent Cloud
	//
	// 	- aliyun: Alibaba Cloud
	//
	// example:
	//
	// Tencent
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service. Only Content Delivery Network (CDN) is supported for Tencent Cloud.
	//
	// example:
	//
	// SLB
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The domain name bound to the cloud resource.
	//
	// example:
	//
	// cert-instanceId
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The AccessKey ID used to access cloud resources.
	//
	// example:
	//
	// 21
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The number of revoked certificates per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListCloudResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesRequest) SetCertIds(v []*int64) *ListCloudResourcesRequest {
	s.CertIds = v
	return s
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

type ListCloudResourcesShrinkRequest struct {
	CertIdsShrink *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// The cloud service provider.
	//
	// Valid values:
	//
	// 	- Tencent: Tencent Cloud
	//
	// 	- aliyun: Alibaba Cloud
	//
	// example:
	//
	// Tencent
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service. Only Content Delivery Network (CDN) is supported for Tencent Cloud.
	//
	// example:
	//
	// SLB
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The domain name bound to the cloud resource.
	//
	// example:
	//
	// cert-instanceId
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The AccessKey ID used to access cloud resources.
	//
	// example:
	//
	// 21
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The number of revoked certificates per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListCloudResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesShrinkRequest) SetCertIdsShrink(v string) *ListCloudResourcesShrinkRequest {
	s.CertIdsShrink = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetCloudName(v string) *ListCloudResourcesShrinkRequest {
	s.CloudName = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetCloudProduct(v string) *ListCloudResourcesShrinkRequest {
	s.CloudProduct = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetCurrentPage(v int32) *ListCloudResourcesShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetKeyword(v string) *ListCloudResourcesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetSecretId(v string) *ListCloudResourcesShrinkRequest {
	s.SecretId = &v
	return s
}

func (s *ListCloudResourcesShrinkRequest) SetShowSize(v int32) *ListCloudResourcesShrinkRequest {
	s.ShowSize = &v
	return s
}

type ListCloudResourcesResponseBody struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The response parameters.
	Data []*ListCloudResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of certificate authority (CA) certificates per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 440
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The end date of the certificate bound to the cloud resource. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1737795520000
	CertEndTime *string `json:"CertEndTime,omitempty" xml:"CertEndTime,omitempty"`
	// The ID of the certificate bound to the cloud resource.
	//
	// example:
	//
	// 12433121
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate bound to the cloud resource.
	//
	// example:
	//
	// shop.amsaudio.cn
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The start date of the certificate bound to the cloud resource. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1706259520000
	CertStartTime *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The AccessKey ID used to access cloud resources.
	//
	// >  This parameter is required only when you deploy certificates to services of multiple clouds.
	//
	// example:
	//
	// 1234
	CloudAccessId *string `json:"CloudAccessId,omitempty" xml:"CloudAccessId,omitempty"`
	// The cloud service provider of the cloud resource. Valid values:
	//
	// 	- **aliyun**: Alibaba Cloud
	//
	// 	- **Tencent**: Tencent Cloud
	//
	// example:
	//
	// aliyun
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service.
	//
	// example:
	//
	// ALB
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The region ID of the cloud service provider to which the cloud resource belongs.
	//
	// example:
	//
	// cn-hangzhou
	CloudRegion *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	// Indicates whether the cloud resource is the default resource. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 0
	DefaultResource *int32 `json:"DefaultResource,omitempty" xml:"DefaultResource,omitempty"`
	// The domain name bound to the cloud resource.
	//
	// example:
	//
	// www.tkgeo.ru
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether HTTPS is enabled for the cloud resource. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	EnableHttps *int32 `json:"EnableHttps,omitempty" xml:"EnableHttps,omitempty"`
	// The time when the cloud resource was created. The time is a timestamp in seconds.
	//
	// example:
	//
	// 1673423339000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the cloud resource was last modified. The time is a timestamp in seconds.
	//
	// example:
	//
	// 1696911946000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the cloud resource.
	//
	// example:
	//
	// 2356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// nlb-rv05agjc97ovm14il5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listener ID of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// lsn-jiugof6t23et66lsnc@443
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listening port of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 8047
	ListenerPort *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The region ID of the cloud resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the cloud resource.
	//
	// example:
	//
	// BUY
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether an Alibaba Cloud SSL certificate is used. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// >  This parameter is required only when you deploy certificates to services of multiple clouds.
	//
	// example:
	//
	// 1
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1666884372152785
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword used in the query. For example, you can specify a keyword in names, email addresses, and mobile phone numbers.
	//
	// example:
	//
	// 186
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of contacts per page.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
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
	// The contacts.
	ContactList []*ListContactResponseBodyContactList `json:"ContactList,omitempty" xml:"ContactList,omitempty" type:"Repeated"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword used in the fuzzy search.
	//
	// example:
	//
	// 186
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 31C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of certificates per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The ID of the contact.
	//
	// example:
	//
	// 519580
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the email address passed the verification.
	//
	// example:
	//
	// 1
	EmailStatus *int32 `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 139****8888
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Indicates whether the phone number was verified.
	//
	// example:
	//
	// 1
	MobileStatus *int32 `json:"MobileStatus,omitempty" xml:"MobileStatus,omitempty"`
	// The name of the contact.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The webhook URL of the chatbot.
	//
	// example:
	//
	// [\\"https://open.feishu.cn/open-apis/bot/v2/hook/XXX\\",\\"https://oapi.dingtalk.com/robot/send?access_token=XXX\\",\\"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=XXX\\"]
	Webhooks *string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
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
	// The algorithm. Valid values: RSA, ECC, and SM2.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword.
	//
	// example:
	//
	// test_name
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The number of entries per page. Default value: 50.
	//
	// example:
	//
	// 10
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the deployment task.
	//
	// Valid values:
	//
	// 	- cloud: multi-cloud deployment task.
	//
	// 	- user: cloud service deployment task. This type of task does not support Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// user
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The number of certificates per page. Default value: **50**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The status of the deployment task.
	//
	// Valid values:
	//
	// 	- success
	//
	// 	- pending
	//
	// 	- scheduling
	//
	// 	- processing
	//
	// 	- error
	//
	// 	- editing
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data returned for the request.
	Data []*ListDeploymentJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of deployment tasks per page. Default value: **50**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of deployment tasks returned.
	//
	// example:
	//
	// 7
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The domain names bound to the certificate of the deployment task.
	//
	// example:
	//
	// aliyundoc1.com,aliyundoc2.com,aliyundoc3.com
	CertDomain *string `json:"CertDomain,omitempty" xml:"CertDomain,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **upload**: uploaded certificate
	//
	// 	- **buy**: purchased certificate
	//
	// 	- **free**: free certificate, available only on the China site (aliyun.com)
	//
	// example:
	//
	// upload
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// Indicates whether the deployment task is deleted. Valid values:
	//
	// 	- **0**: not deleted
	//
	// 	- **1**: deleted
	//
	// example:
	//
	// 1
	Del *int32 `json:"Del,omitempty" xml:"Del,omitempty"`
	// The end time of the deployment task.
	//
	// example:
	//
	// 1606482979000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the deployment task was created.
	//
	// example:
	//
	// 1624343180000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the deployment task was last modified.
	//
	// example:
	//
	// 1606482979000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the deployment task. You can use the ID to query the details and status of the deployment task.
	//
	// example:
	//
	// 19975
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the deployment task.
	//
	// example:
	//
	// cas-job-user-0gvntn
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the deployment task.
	//
	// 	- **cloud**: multi-cloud deployment task.
	//
	// 	- **user**: cloud service deployment task. This type of task does not support ECS instances.
	//
	// example:
	//
	// user
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the deployment task.
	//
	// example:
	//
	// job-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cloud service included in the resources of the deployment task.
	//
	// example:
	//
	// NLB
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Indicates whether the rollback worker is included. For example, if a cloud service involved in a deployment task has been rolled back, **1*	- is returned. Valid values:
	//
	// 	- **0**: The rollback worker is not included.
	//
	// 	- **1**: The rollback worker is included.
	//
	// example:
	//
	// 1
	Rollback *int32 `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
	// The time when the deployment task was scheduled.
	//
	// example:
	//
	// 1606482979000
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The start time of the deployment task.
	//
	// example:
	//
	// 1606482979000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the deployment task. Valid values:
	//
	// 	- **pending**
	//
	// 	- **editing**
	//
	// 	- **scheduling**
	//
	// 	- **processing**
	//
	// 	- **error**
	//
	// 	- **success**
	//
	// example:
	//
	// scheduling
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 1666884372152785
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The ID of the deployment task. You can call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) operation to obtain the ID of a deployment task from the **JobId*	- parameter. You can also call the [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation to obtain the ID of a deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
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
	// The response parameters.
	Data []*ListDeploymentJobCertResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The algorithm of the certificate public key.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 11174100
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The instance ID of the certificate.
	//
	// example:
	//
	// cas-ivauto-2crxzi
	CertInstanceId *string `json:"CertInstanceId,omitempty" xml:"CertInstanceId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// edkog.shop
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The type of the certificate order. Valid values:
	//
	// 	- **upload**: uploaded certificate.
	//
	// 	- **buy**: purchased certificate.
	//
	// 	- **free**: free certificate. This value is available only on the China site (aliyun.com).
	//
	// example:
	//
	// buy
	CertOrderType *string `json:"CertOrderType,omitempty" xml:"CertOrderType,omitempty"`
	// The type of the certificate.
	//
	// example:
	//
	// DV
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// vaultwebhook.vault-webhook.svc
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// Indicates whether the certificate is hosted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsTrustee *bool `json:"IsTrustee,omitempty" xml:"IsTrustee,omitempty"`
	// The month in which the certificate is applied for.
	//
	// example:
	//
	// 12
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// The end time of the validity period of the certificate. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1683625266108
	NotAfterTime *int64 `json:"NotAfterTime,omitempty" xml:"NotAfterTime,omitempty"`
	// The start time of the validity period of the certificate. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1683625266108
	NotBeforeTime *int64 `json:"NotBeforeTime,omitempty" xml:"NotBeforeTime,omitempty"`
	// The ID of the certificate order.
	//
	// >  If CertId is returned, this parameter is not returned.
	//
	// example:
	//
	// 6127067
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The subject alternative name (SAN) extensions of the certificate.
	Sans []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
	// The status code of the certificate. Valid values:
	//
	// 	- **payed**: paid and pending application
	//
	// 	- **checking**: being validated
	//
	// 	- **checkedFail**: validation failed
	//
	// 	- **revoked**: revoked
	//
	// 	- **revokeChecking**: revocation request being validated
	//
	// 	- **issued**: issued (excluding hosted certificates that are issued, certificates that are about to expire, expired certificates, and uploaded certificates)
	//
	// 	- **trustee**: hosted and issued
	//
	// 	- **upload**: uploaded (excluding certificates that are about to expire and expired certificates)
	//
	// 	- **willExpired**: about to expire (including certificates issued by using the Certificate Management Service console and uploaded certificates)
	//
	// 	- **expired**: expired (including certificates issued by using the Certificate Management Service console and uploaded certificates)
	//
	// 	- **validity**: valid (including certificates that are not expired or revoked)
	//
	// 	- **refund**: refunded
	//
	// 	- **closed**: closed
	//
	// example:
	//
	// issued
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
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
	// The ID of the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
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
	// The response parameters.
	Data []*ListDeploymentJobResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The end date of the certificate bound to the cloud resource. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1681956830000
	CertEndTime *string `json:"CertEndTime,omitempty" xml:"CertEndTime,omitempty"`
	// The ID of the certificate bound to the cloud resource.
	//
	// example:
	//
	// 11599949
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate bound to the cloud resource.
	//
	// example:
	//
	// sc-SSL
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The start date of the certificate bound to the cloud resource. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1681956830000
	CertStartTime *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The AccessKey ID used to access cloud resources.
	//
	// >  This parameter is required only when you deploy certificates to services of multiple clouds.
	//
	// example:
	//
	// 1234
	CloudAccessId *string `json:"CloudAccessId,omitempty" xml:"CloudAccessId,omitempty"`
	// The cloud service provider of the cloud resource. Valid values:
	//
	// 	- **aliyun**: Alibaba Cloud
	//
	// 	- **Tencent**: Tencent Cloud
	//
	// example:
	//
	// aliyun
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service. Valid values:
	//
	// 	- **CDN**: Alibaba Cloud CDN (CDN). This value is supported only at the China site (aliyun.com).
	//
	// 	- **SLB**: Classic Load Balancer (CLB). This value is supported only at the China site (aliyun.com).
	//
	// 	- **DCDN**: Dynamic Content Delivery Network (DCDN). This value is supported only at the China site (aliyun.com).
	//
	// 	- **DDOS**: Anti-DDoS. This value is supported only at the China site (aliyun.com).
	//
	// 	- **LIVE**: ApsaraVideo Live. This value is supported only at the China site (aliyun.com).
	//
	// 	- **webHosting**: Cloud Web Hosting. This value is supported only at the China site (aliyun.com).
	//
	// 	- **VOD**: ApsaraVideo VOD. This value is supported only at the China site (aliyun.com).
	//
	// 	- **CR**: Container Registry. This value is supported only at the China site (aliyun.com).
	//
	// 	- **ALB**: Application Load Balancer (ALB).
	//
	// 	- **APIGateway**: API Gateway.
	//
	// 	- **FC**: Function Compute.
	//
	// 	- **GA**: Global Accelerator (GA).
	//
	// 	- **MSE**: Microservices Engine (MSE).
	//
	// 	- **NLB**: Network Load Balancer (NLB).
	//
	// 	- **OSS**: Object Storage Service (OSS).
	//
	// 	- **SAE**: Serverless App Engine (SAE).
	//
	// 	- **TencentCDN**: Tencent Cloud Content Delivery Network (CDN).
	//
	// 	- **WAF**: Web Application Firewall (WAF).
	//
	// example:
	//
	// SLB
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The region ID of the cloud service provider to which the cloud resource belongs.
	//
	// example:
	//
	// cn-hangzhou
	CloudRegion *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	// Indicates whether the cloud resource is the default resource. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 0
	DefaultResource *int32 `json:"DefaultResource,omitempty" xml:"DefaultResource,omitempty"`
	// The domain name bound to the cloud resource.
	//
	// example:
	//
	// aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether HTTPS is enabled for the cloud resource. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	EnableHttps *int32 `json:"EnableHttps,omitempty" xml:"EnableHttps,omitempty"`
	// The time when the cloud resource was created. The time is a timestamp in seconds.
	//
	// example:
	//
	// 1673423339000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the cloud resource was last modified. The time is in the timestamp format.
	//
	// example:
	//
	// 1681956830000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the cloud resource.
	//
	// example:
	//
	// 20979
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// cas-cn-m7r1qocw91at
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listener ID of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// lsn-vwdff0q20poq5xazb9@443
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listening port of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 8047
	ListenerPort *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The region ID of the cloud resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The other metadata related to the cloud resource.
	//
	// example:
	//
	// {\\"camera_model\\":\\"GIFSHOW [1267087617][OnePlus
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the cloud resource.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether an Alibaba Cloud SSL certificate is used. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// >  This parameter is required only when you deploy certificates to services of multiple clouds.
	//
	// example:
	//
	// 1
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1666884372152785
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The domain name that is bound or the ID of the resource. Fuzzy match is supported.
	//
	// example:
	//
	// cert-instanceId
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The type of the order. Default value: **CPACK**. Valid values:
	//
	// 	- **CPACK**: virtual resource order. If you set OrderType to CPACK, only the information about orders that are generated to consume the certificate quota is returned.
	//
	// 	- **BUY**: purchase order. If you set OrderType to BUY, only the information about purchase orders is returned. In most cases, this type of order can be ignored.
	//
	// 	- **UPLOAD**: uploaded certificate. If you set OrderType to UPLOAD, only uploaded certificates are returned.
	//
	// 	- **CERT**: certificate. If you set OrderType to CERT, both issued certificates and uploaded certificates are returned.
	//
	// example:
	//
	// CPACK
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group. You can call the [ListResources](https://help.aliyun.com/document_detail/2716559.html) operation to obtain the ID.
	//
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of entries to return on each page. Default value: 50.
	//
	// example:
	//
	// 10
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The certificate status of the order. Valid values:
	//
	// 	- **PAYED**: pending application. You can set Status to PAYED only if you set OrderType to CPACK or BUY.
	//
	// 	- **CHECKING**: validating. You can set Status to CHECKING only if you set OrderType to CPACK or BUY.
	//
	// 	- **CHECKED_FAIL**: validation failed. You can set Status to CHECKED_FAIL only if you set OrderType to CPACK or BUY.
	//
	// 	- **ISSUED**: issued.
	//
	// 	- **WILLEXPIRED**: about to expire.
	//
	// 	- **EXPIRED**: expired.
	//
	// 	- **NOTACTIVATED**: not activated. You can set Status to NOTACTIVATED only if you set OrderType to CPACK or BUY.
	//
	// 	- **REVOKED**: revoked. You can set Status to REVOKED only if you set OrderType to CPACK or BUY.
	//
	// If you set OrderType to CERT or UPLOAD and Status is left empty, valid certificates are returned by default, including issued certificates and certificates that are about to expire. If you set OrderType to CPACK or BUY and Status is left empty, all orders are returned by default.
	//
	// example:
	//
	// ISSUED
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
	// The certificates and orders.
	CertificateOrderList []*ListUserCertificateOrderResponseBodyCertificateOrderList `json:"CertificateOrderList,omitempty" xml:"CertificateOrderList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
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
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the Alibaba Cloud order. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 234567
	AliyunOrderId *int64 `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	// The time at which the order was placed. Unit: milliseconds. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 1634283958000
	BuyDate *int64 `json:"BuyDate,omitempty" xml:"BuyDate,omitempty"`
	// The time at which the certificate expires. Unit: milliseconds. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 1665819958000
	CertEndTime *int64 `json:"CertEndTime,omitempty" xml:"CertEndTime,omitempty"`
	// The time at which the certificate starts to take effect. Unit: milliseconds. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 1665819958000
	CertStartTime *int64 `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The type of the certificate. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// 	- **DV**: domain validated (DV) certificate
	//
	// 	- **EV**: extended validation (EV) certificate
	//
	// 	- **OV**: organization validated (OV) certificate **FREE**: free certificate, available only on the China site (aliyun.com)
	//
	// example:
	//
	// FREE
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The ID of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// 896521
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The city in which the organization is located. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// Hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The parent domain name of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// aliyun.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The domain name. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The total number of domain names that can be bound to the certificate. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 1
	DomainCount *int64 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The type of the domain name. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// 	- **ONE**: single domain name
	//
	// 	- **MULTIPLE**: multiple domain names
	//
	// 	- **WILDCARD**: single wildcard domain name
	//
	// 	- **M_WILDCARD**: multiple wildcard domain names
	//
	// 	- **MIX**: hybrid domain name
	//
	// example:
	//
	// ONE
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The time at which the certificate expires. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// 2022-11-17
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Indicates whether the certificate expires. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// true
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The fingerprint of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// CC6B3696E7C7CA715BD26E28E45FF3E3DF435C03
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// cas-instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issuer of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// MyIssuer
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The name of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// cert-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order ID. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 2345687
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The name of the organization that is associated with the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// Alibaba Cloud
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The ID of the third-party certificate authority (CA) order. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// ca-123456
	PartnerOrderId *string `json:"PartnerOrderId,omitempty" xml:"PartnerOrderId,omitempty"`
	// The specification ID of the order. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// bykj123456
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The specification name of the order. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// CFCA
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The province or autonomous region in which the organization is located. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// Zhejiang
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The ID of the resource group. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The brand of the certificate. Valid values: WoSign, CFCA, DigiCert, and vTrus. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// CFCA
	RootBrand *string `json:"RootBrand,omitempty" xml:"RootBrand,omitempty"`
	// All domain names that are bound to the certificate. Multiple domain names are separated by commas (,). This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// aliyun.com
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// 040a6e493cffdda6d744acf99b6576cf
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	// The SHA-2 value of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// 56B4DED2243A81DD909D7C39824FFE4DDBD87F91BFA46CD333FF212FE0E7CB11
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The type of the order. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// 	- **cpack**: virtual resource order
	//
	// 	- **buy**: purchase order
	//
	// example:
	//
	// buy
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time at which the certificate starts to take effect. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// 2021-11-16
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The certificate status of the order. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// 	- **PAYED**: pending application
	//
	// 	- **CHECKING**: reviewing
	//
	// 	- **CHECKED_FAIL**: review failed
	//
	// 	- **ISSUED**: issued
	//
	// 	- **WILLEXPIRED**: about to expire
	//
	// 	- **EXPIRED**: expired
	//
	// 	- **NOTACTIVATED**: not activated
	//
	// 	- **REVOKED**: revoked
	//
	// example:
	//
	// PAYED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The hosting status of the certificate. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// 	- **unTrustee**: not hosted
	//
	// 	- **trustee**: hosted
	//
	// example:
	//
	// unTrustee
	TrusteeStatus *string `json:"TrusteeStatus,omitempty" xml:"TrusteeStatus,omitempty"`
	// Indicates whether the certificate is an uploaded certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// false
	Upload *bool `json:"Upload,omitempty" xml:"Upload,omitempty"`
	// The number of wildcard domain names that can be bound to the certificate. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 0
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
	// The cloud service in the deployment task.
	//
	// example:
	//
	// NLB
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the deployment task. You can call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) operation to obtain the ID of a deployment task from the **ID*	- parameter. You can also call the [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation to obtain the ID of a deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The number of entries per page. Default value: 50.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The status of the worker task.
	//
	// Valid values:
	//
	// 	- rollback
	//
	// 	- rollback_error
	//
	// 	- success
	//
	// 	- rollback_success
	//
	// 	- pending
	//
	// 	- scheduling
	//
	// 	- processing
	//
	// 	- error
	//
	// 	- editing
	//
	// example:
	//
	// editing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The response parameters.
	Data []*ListWorkerResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3E50D480-9765-5CFD-9650-9BACCECA5135
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries per page. Default value: **50**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The domain name bound to the certificate in the worker task.
	//
	// example:
	//
	// www.example.com
	CertDomain *string `json:"CertDomain,omitempty" xml:"CertDomain,omitempty"`
	// The ID of the certificate in the worker task.
	//
	// example:
	//
	// 12073663
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The instance ID of the certificate in the worker task.
	//
	// example:
	//
	// lsn-jzk2h0xz5dmynuphb8@1883
	CertInstanceId *string `json:"CertInstanceId,omitempty" xml:"CertInstanceId,omitempty"`
	// The name of the certificate in the worker task.
	//
	// example:
	//
	// testCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The cloud service provider to which the cloud resource in the worker task belongs.
	//
	// >  This parameter is not returned if you deploy certificates to Alibaba Cloud services.
	//
	// example:
	//
	// aliyun
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service to which the cloud resource in the worker task belongs. Valid values:
	//
	// 	- **CDN**: Alibaba Cloud CDN (CDN). This value is supported only at the China site (aliyun.com).
	//
	// 	- **SLB**: Classic Load Balancer (CLB). This value is supported only at the China site (aliyun.com).
	//
	// 	- **DCDN**: Dynamic Content Delivery Network (DCDN). This value is supported only at the China site (aliyun.com).
	//
	// 	- **DDOS**: Anti-DDoS. This value is supported only at the China site (aliyun.com).
	//
	// 	- **LIVE**: ApsaraVideo Live. This value is supported only at the China site (aliyun.com).
	//
	// 	- **webHosting**: Cloud Web Hosting. This value is supported only at the China site (aliyun.com).
	//
	// 	- **VOD**: ApsaraVideo VOD. This value is supported only at the China site (aliyun.com).
	//
	// 	- **CR**: Container Registry. This value is supported only at the China site (aliyun.com).
	//
	// 	- **ALB**: Application Load Balancer (ALB).
	//
	// 	- **APIGateway**: API Gateway.
	//
	// 	- **FC**: Function Compute.
	//
	// 	- **GA**: Global Accelerator (GA).
	//
	// 	- **MSE**: Microservices Engine (MSE).
	//
	// 	- **NLB**: Network Load Balancer (NLB).
	//
	// 	- **OSS**: Object Storage Service (OSS).
	//
	// 	- **SAE**: Serverless App Engine (SAE).
	//
	// 	- **TencentCDN**: Tencent Cloud Content Delivery Network (CDN).
	//
	// 	- **WAF**: Web Application Firewall (WAF).
	//
	// example:
	//
	// SLB
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The original region ID of the cloud resource in the worker task. The value is the region ID defined by the cloud service provider. This parameter is required only when you deploy certificates to services of multiple clouds.
	//
	// example:
	//
	// cn-hangzhou
	CloudRegion *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	// Indicates whether the cloud resource in the worker task is the default resource. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 0
	DefaultResource *bool `json:"DefaultResource,omitempty" xml:"DefaultResource,omitempty"`
	// The time when the worker task was created. The time is a timestamp in seconds.
	//
	// example:
	//
	// 1680228896000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the worker task was last modified. The time is a timestamp in seconds.
	//
	// example:
	//
	// 1681956830000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the worker task.
	//
	// example:
	//
	// 22487
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the cloud resource in the worker task.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// cas-cn-0pp118nuu40b
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the deployment task to which the worker task belongs.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The listener ID of the cloud resource in the worker task.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// lsn-lhl8y7s1e1ngc3m3zz@81
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the namespace in SAE. This parameter is returned only if you deploy certificates to SAE.
	//
	// example:
	//
	// 32fed52a-4bf7-44f6-955f-e82ada68ef18
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The order ID of the worker task, which is the same as the order ID of the certificate.
	//
	// >  If the CertId parameter is returned, this parameter is not returned.
	//
	// example:
	//
	// 9349278
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The listening port of the cloud resource in the worker task.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 4431
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the cloud resource in the worker task.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the certificate that was originally bound to the cloud resource in the worker task.
	//
	// example:
	//
	// 123
	ResourceCertId *int64 `json:"ResourceCertId,omitempty" xml:"ResourceCertId,omitempty"`
	// The domain name that was originally bound to the cloud resource in the worker task.
	//
	// example:
	//
	// www.example.com
	ResourceDomain *string `json:"ResourceDomain,omitempty" xml:"ResourceDomain,omitempty"`
	// The ID of the cloud resource in the worker task.
	//
	// example:
	//
	// 1286999
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status of the worker task. Valid values:
	//
	// 	- **editing**
	//
	// 	- **pending**
	//
	// 	- **scheduling**
	//
	// 	- **processing**
	//
	// 	- **error**
	//
	// 	- **success**
	//
	// 	- **rollback**
	//
	// 	- **rollback_success**
	//
	// 	- **rollback_error**
	//
	// example:
	//
	// editing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Alibaba Cloud account to which the worker task belongs.
	//
	// example:
	//
	// 1666884372152785
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

type MoveResourceGroupRequest struct {
	// The region of the organization to which the owner of the certificate belongs. Valid values: ap-southeast-1 and cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmykgxu5d46ey
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cas-cn-4591d3xa****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.\\
	//
	// Default value: **instance**
	//
	// Valid values:
	//
	// 	- instance: certificate order
	//
	// 	- Certificate: certificate
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BCD2F6C-7A9D-47C1-8588-2CC6A4E0BE5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type RenewCertificateOrderForPackageRequestRequest struct {
	// The content of the certificate signing request (CSR) file that is manually generated for the domain name by using OpenSSL or Keytool. The key algorithm in the CSR file must be Rivest-Shamir-Adleman (RSA) or elliptic-curve cryptography (ECC), and the key length of the RSA algorithm must be greater than or equal to 2,048 characters. For more information about how to create a CSR file, see [How do I create a CSR file?](https://help.aliyun.com/document_detail/42218.html)
	//
	// If you do not specify this parameter, Certificate Management Service automatically generates a CSR file for the domain name in the certificate application order that you want to renew.
	//
	// A CSR file contains the information about your server and company. When you apply for a certificate, you must submit the CSR file to the CA. The CA signs the CSR file by using the private key of the root certificate and generates a public key file to issue your certificate.
	//
	// >  The **CN*	- field in the CSR file specifies the domain name that is bound to the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- MIIC1TCCAb0CAQAwgY8xCzAJBgNVBAYTAkNOMRIwEAYDVQQIDAlHdWFuZ3pob3Ux ETAPBgNVBAcMCFNoZW56aGVuMQ8wDQYDVQQKDAZDaGFjdW8xEDAOBgNVBAsMB0lU IERlcHQxFzAVBgNVBAMMDnd3dy5jaGFjdW8ubmV0MR0wGwYJKoZIhvcNAQkBFg44 MjkyNjY5QHFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALo7 atRvQf9tKo1NJ/MQqzHvIjHNhU+0MMerDq+tRlJ+a7Ro1r6IWNF5MB0Z*****	- -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The ID of the certificate application order that you want to renew.
	//
	// >  After you call the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html), [CreateCertificateRequest](https://help.aliyun.com/document_detail/455292.html), or [CreateCertificateWithCsrRequest](https://help.aliyun.com/document_detail/455801.html) operation to submit a certificate application, you can obtain the ID of the certificate application order from the **OrderId*	- response parameter. You can also call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the order ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123451222
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
	// >  You can use the ID to query the status of the certificate application order. For more information, see [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html).
	//
	// example:
	//
	// 323451222
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
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
	// The unique identifier of the client certificate or server certificate that you want to revoke.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
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
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
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
	// The unique identifier of the certificate. You can call the [ListCert](https://help.aliyun.com/document_detail/455806.html) operation to obtain the identifier.
	//
	// 	- If the certificate is an SSL certificate, the value of this parameter must be in the {Certificate ID}-cn-hangzhou format.
	//
	// 	- If the certificate is a private certificate, the value of this parameter must be the value of the Identifier field for the private certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccaf0c629c2be1e2abb63bb76b
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The data to sign. The value must be encoded in Base64.\\
	//
	// For example, if the hexadecimal data that you want to sign is [0x31, 0x32, 0x33, 0x34], set the parameter to the Base64-encoded value MTIzNA==. If you set MessageType to RAW, the size of the data must be less than 4 KB. If the size of the data is greater than 4 KB, you can set MessageType to DIGEST and set Message to the digest of the data. The digest is a hash value. You can compute the digest of the data on an on-premises machine. The certificate application repository uses the digest that you compute in your own certificate application system. The message digest algorithm that you use must match the specified signature algorithm. The following items describe the details:
	//
	// 	- If the signature algorithm is SHA256withRSA, SHA256withRSA/PSS, or SHA256withECDSA, the message digest algorithm must be SHA-256.
	//
	// 	- If the signature algorithm is SM3withSM2, the message digest algorithm must be SM3.
	//
	// This parameter is required.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The value type of the Message parameter. Valid values:
	//
	// 	- RAW: the raw data. This is the default value.
	//
	// 	- DIGEST: the message digest (hash value) of the raw data.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAW
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The signature algorithm. Valid values:
	//
	// 	- SHA256withRSA
	//
	// 	- SHA256withRSA/PSS
	//
	// 	- SHA256withECDSA
	//
	// 	- SM3withSM2
	//
	// This parameter is required.
	//
	// example:
	//
	// SHA256withRSA
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
	// The ID of the request.
	//
	// example:
	//
	// 1ed33293-2e48-6b14-861e-538e28e408eb
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature.
	//
	// example:
	//
	// eyaC0w3ROK5b3QcHmUtAhMY/sQjKu2t3uBfnf6J/gn7JfZtyxwcCUjzXbw5jmqJQRbj1te670Bshg9kUdanKhtHFhJjU5jX+ZMMBr6pH0gqQDJxR0K0yHXRc0Q5OQoUZ6BfpbI4Wt4jJvJSdCstz1vSg12CfEHS8Kd5qfhItK7Y=
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
	// The ID of the CSR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5209
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	// The private key content of the certificate in the PEM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- MII.... -----END RSA PRIVATE KEY-----
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
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
	// The ID of the certificate. Separate multiple certificate IDs with commas (,). You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the ID of a certificate from the **CertificateId*	- parameter.
	//
	// example:
	//
	// 9957285,12067351,12304338,12342649
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// The ID of the contact. Separate multiple contact IDs with commas (,). You can call the [ListContact](https://help.aliyun.com/document_detail/2712221.html) operation to obtain the ID of a contact from the **ContactId*	- parameter.
	//
	// example:
	//
	// 9957285,12067351,12304338,12342649
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	// The ID of the deployment task. You can call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) operation to obtain the ID of a deployment task from the JobId parameter. You can also call the [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation to obtain the ID of a deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the deployment task.
	//
	// example:
	//
	// cert-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the cloud resource. Separate multiple resource IDs with commas (,). You can call the [ListCloudResources](https://help.aliyun.com/document_detail/2712230.html) operation to obtain the ID of a cloud resource from the **Id*	- parameter.
	//
	// example:
	//
	// 9957285,12067351,12304338,12342649
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The time when the task starts. The value is a UNIX timestamp. If you do not specify this parameter, the system immediately starts the task after it is created.
	//
	// example:
	//
	// 1606482979000
	ScheduleTime *int64 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
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
	// The ID of the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The desired status.
	//
	// Valid values:
	//
	// 	- pending
	//
	// 	- scheduling
	//
	// 	- editing
	//
	// This parameter is required.
	//
	// example:
	//
	// editing
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
	// The response parameters.
	//
	// example:
	//
	// []
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA69E364-5CBB-50E8-BF09-E8CAA396A4F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the deployment task. You can call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) operation to obtain the ID of a deployment task from the **JobId*	- parameter. You can also call the [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation to obtain the ID of a deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The desired status.
	//
	// Valid values:
	//
	// 	- rollback
	//
	// This parameter is required.
	//
	// example:
	//
	// rollback
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the worker task. You can call the [ListWorkerResource](https://help.aliyun.com/document_detail/2712224.html) operation to obtain the ID of a worker task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	WorkerId *int64 `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
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
	// The response parameters.
	//
	// example:
	//
	// []
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The content of the CSR.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The private key content of the certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- MII.... -----END RSA PRIVATE KEY-----
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the CSR.
	//
	// example:
	//
	// name
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
	// The ID of the CSR.
	//
	// example:
	//
	// 2271
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
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
	//
	//     <RequestId>15C66C7B-671A-4297-9187-2C4477247A74</RequestId>
	//
	// </UploadPCACertResponse>
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIEJDCCAwygAwIBAgIQITRHItTLTQizTyd3K7AMRTANBgkqhkiG9w0BAQsFADBe ***************	- 5/akmr2GK/Y= -----END CERTIFICATE----- -----BEGIN CERTIFICATE----- MIIDuzCCAqOgAwIBAgIQSEIWDPfWTDKZcWNyL2O+fjANBgkqhkiG9w0BAQsFADBf ***************	- URUHyMW5Qd5Q9g6Y4sDOIm6It9TF7EjpwMs42R30agcRYzuUsN72ZFBYFJwnBX8= -----END CERTIFICATE----- -----BEGIN CERTIFICATE----- MIIDizCCAnOgAwIBAgIRAMfjPkDKfELTo07l3A3cUSYwDQYJKoZIhvcNAQELBQAw ********	- CjWTnYPhCcO2uIcnqMt7zCVs5LXBK/XSwlAXKMvKT0uuzw9VxeMfEabflKu0By8= -----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// UploadPCACert
	//
	// example:
	//
	// cert_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Uploads a private certificate to a certificate repository.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- MIIEowIBAAKCAQEA5SIfpNCBoiDrZhX1H39CHwQMVD0kBNeBTWfP9xkeesvfzbOz ******	- POVNFfDf9h7pJtQ5fRZNTYTDs/d+cH62Z3+nS74mNnEfff0nkvne -----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The ID of the repository.
	//
	// >  You can call the [ListCertWarehouse](https://help.aliyun.com/document_detail/455805.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
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
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 1ed65580-7e33-6a50-8630-dd13fdc009ee
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIF...... -----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The content of the encryption certificate in PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICDzCCA***
	//
	// -----END CERTIFICATE-----
	EncryptCert *string `json:"EncryptCert,omitempty" xml:"EncryptCert,omitempty"`
	// The private key of the encryption certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN EC PRIVATE KEY-----
	//
	// MHcCAQEEI****
	//
	// -----END EC PRIVATE KEY-----
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	// The private key of the certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICDzCCAbagAw****
	//
	// -----END CERTIFICATE-----
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the certificate. The name can contain up to 128 characters in length. The name can contain all types of characters, such as letters, digits, and underscores (_).
	//
	// >  The name must be unique within an Alibaba Cloud account.
	//
	// example:
	//
	// cert-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// the resource group id.
	//
	// example:
	//
	// rg-ae****vty
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The content of the signing certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICDzCCAbagAw****
	//
	// -----END CERTIFICATE-----
	SignCert *string `json:"SignCert,omitempty" xml:"SignCert,omitempty"`
	// The private key of the signing certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN EC PRIVATE KEY-----
	//
	// MHcCAQEEILR****
	//
	// -----END EC PRIVATE KEY-----
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
	//
	// example:
	//
	// 12345
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BDB81BA2-E1F5-4D08-A2DD-4BE2BF44C90E
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
	// The unique identifier of the certificate. You can call the [ListCert](https://help.aliyun.com/document_detail/455806.html) operation to obtain the unique identifier of a certificate.
	//
	// 	- If the certificate is an SSL certificate, the value of this parameter must be in the {Certificate ID}-cn-hangzhou format.
	//
	// 	- If the certificate is a private certificate, the value of this parameter must be the value of the Identifier field for the private certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5870821-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The data for which you want to verify the signature. The value must be encoded in Base64.\\
	//
	// For example, if the hexadecimal data that you want to verify is [0x31, 0x32, 0x33, 0x34], set the parameter to the Base64-encoded value MTIzNA==. If you set MessageType to RAW, the size of the data must be less than 4 KB. If the size of the data is greater than 4 KB, you can set MessageType to DIGEST and set Message to the digest of the data. The digest is also called hash value. You can compute the digest of the data on an on-premises machine. The certificate repository uses your certificate application system to compute the message digest. The message digest algorithm that is used must meet the requirements of the specified signature algorithm. The following signature algorithms require different message digest algorithms:
	//
	// 	- If the signature algorithm is SHA256withRSA, SHA256withRSA/PSS, or SHA256withECDSA, the message digest algorithm must be SHA-256.
	//
	// 	- If the signature algorithm is SM3withSM2, the message digest algorithm must be SM3.
	//
	// This parameter is required.
	//
	// example:
	//
	// MTIzNA==
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The value type of the Message parameter. Valid values:
	//
	// 	- **RAW**: the raw data. This is the default value.
	//
	// 	- **DIGEST**: the message digest of the raw data.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAW
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The signature value. The value must be encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// eyaC0w3ROK5b3QcHmUtAhMY/sQjKu2t3uBfnf6J/gn7JfZtyxwcCUjzXbw5jmqJQRbj1te670Bshg9kUdanKhtHFhJjU5jX+ZMMBr6pH0gqQDJxR0K0yHXRc0Q5OQoUZ6BfpbI4Wt4jJvJSdCstz1vSg12CfEHS8Kd5qfhItK7Y=
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
	// The signature algorithm. Valid values:
	//
	// 	- **SHA256withRSA**
	//
	// 	- **SHA256withRSA/PSS**
	//
	// 	- **SHA256withECDSA**
	//
	// 	- **SM3withSM2**
	//
	// This parameter is required.
	//
	// example:
	//
	// SHA256withRSA
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
	// The ID of the request.
	//
	// example:
	//
	// 1ed33293-2e48-6b14-861e-538e28e408eb
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the signature is valid. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SignatureValid *bool `json:"SignatureValid,omitempty" xml:"SignatureValid,omitempty"`
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

// Summary:
//
// Revokes an issued certificate and cancels the application order of the certificate.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CancelCertificateForPackageRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCertificateForPackageRequestResponse
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

// Summary:
//
// Revokes an issued certificate and cancels the application order of the certificate.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CancelCertificateForPackageRequestRequest
//
// @return CancelCertificateForPackageRequestResponse
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

// Summary:
//
// Cancels a certificate application order that is in the pending validation or being reviewed state.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CancelOrderRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelOrderRequestResponse
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

// Summary:
//
// Cancels a certificate application order that is in the pending validation or being reviewed state.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CancelOrderRequestRequest
//
// @return CancelOrderRequestResponse
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

// Summary:
//
// Submits a certificate application.
//
// Description:
//
//   Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455800.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// 	- After you call this operation to submit a certificate application and the certificate is issued, the certificate quota provided by the resource plan that you purchased is consumed. When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
// 	- After you call this operation to submit a certificate application, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
//
// @param request - CreateCertificateForPackageRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateForPackageRequestResponse
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

// Summary:
//
// Submits a certificate application.
//
// Description:
//
//   Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455800.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// 	- After you call this operation to submit a certificate application and the certificate is issued, the certificate quota provided by the resource plan that you purchased is consumed. When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
// 	- After you call this operation to submit a certificate application, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
//
// @param request - CreateCertificateForPackageRequestRequest
//
// @return CreateCertificateForPackageRequestResponse
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

// Summary:
//
// Purchases, applies for, and issues a domain validated (DV) certificate by using extended certificate services.
//
// Description:
//
//   You can call this operation to apply for only DV certificates. If you want to apply for an organization validated (OV) or extended validation (EV) certificate, we recommend that you call the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation. This operation allows you to apply for certificates of all specifications and specify the method to generate a certificate signing request (CSR) file.
//
// 	- Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// 	- When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate.
//
// 	- After you call this operation to submit a certificate application, Certificate Management Service automatically creates a CSR file for your application and consumes the certificate quota in the certificate resource plans of the specified specifications that you purchased. After you call this operation, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required to complete domain name verification, and manually complete the verification. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on your DNS server. Then, the certificate authority (CA) will review your certificate application.
//
// @param request - CreateCertificateRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateRequestResponse
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

// Summary:
//
// Purchases, applies for, and issues a domain validated (DV) certificate by using extended certificate services.
//
// Description:
//
//   You can call this operation to apply for only DV certificates. If you want to apply for an organization validated (OV) or extended validation (EV) certificate, we recommend that you call the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation. This operation allows you to apply for certificates of all specifications and specify the method to generate a certificate signing request (CSR) file.
//
// 	- Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// 	- When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate.
//
// 	- After you call this operation to submit a certificate application, Certificate Management Service automatically creates a CSR file for your application and consumes the certificate quota in the certificate resource plans of the specified specifications that you purchased. After you call this operation, you also need to call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required to complete domain name verification, and manually complete the verification. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on your DNS server. Then, the certificate authority (CA) will review your certificate application.
//
// @param request - CreateCertificateRequestRequest
//
// @return CreateCertificateRequestResponse
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

// Summary:
//
// Purchases, applies for, and issues a domain validated (DV) certificate by using a custom certificate signing request (CSR) file. You can use extended certificate services to purchase and apply for a DV certificate with a few clicks.
//
// Description:
//
//   You can use this operation to apply for only a domain validated (DV) certificate. You cannot use this operation to apply for an organization validated (OV) certificate. We recommend that you use the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation to apply for a certificate. You can use the CreateCertificateForPackageRequest operation to apply for certificates of all types and specify the CSR generation method.
//
// 	- Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// 	- When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
// 	- After you call this operation to submit a certificate application, the certificate quota of the required specifications that you purchased is consumed. After you call this operation, you must call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
//
// @param request - CreateCertificateWithCsrRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateWithCsrRequestResponse
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

// Summary:
//
// Purchases, applies for, and issues a domain validated (DV) certificate by using a custom certificate signing request (CSR) file. You can use extended certificate services to purchase and apply for a DV certificate with a few clicks.
//
// Description:
//
//   You can use this operation to apply for only a domain validated (DV) certificate. You cannot use this operation to apply for an organization validated (OV) certificate. We recommend that you use the [CreateCertificateForPackageRequest](https://help.aliyun.com/document_detail/455296.html) operation to apply for a certificate. You can use the CreateCertificateForPackageRequest operation to apply for certificates of all types and specify the CSR generation method.
//
// 	- Before you call this operation, make sure that you have purchased a certificate resource plan of the required specifications. For more information about how to purchase a certificate resource plan, see [Purchase a certificate resource plan](https://help.aliyun.com/document_detail/28542.html). You can call the [DescribePackageState](https://help.aliyun.com/document_detail/455803.html) operation to query the usage of a certificate resource plan of specified specifications, including the total number of certificate resource plans that you purchase, the number of certificate applications that you submit, and the number of certificates that are issued.
//
// 	- When you call this operation, you can use the **ProductCode*	- parameter to specify the specifications of the certificate that you want to apply for.
//
// 	- After you call this operation to submit a certificate application, the certificate quota of the required specifications that you purchased is consumed. After you call this operation, you must call the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to obtain the information that is required for domain name ownership verification and manually complete the verification. Then, your certificate application is reviewed by the certificate authority (CA). If you use the Domain Name System (DNS) verification method, you must complete the verification on your DNS service provider system. If you use the file verification method, you must complete the verification on the DNS server.
//
// @param request - CreateCertificateWithCsrRequestRequest
//
// @return CreateCertificateWithCsrRequestResponse
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

// Summary:
//
// Creates a certificate signing request (CSR). A CSR file contains the information about an SSL certificate that you want to apply for. The information includes the domain names that you want to bind to the certificate and the name and the geographical location of the certificate holder. When you submit a certificate application to a certificate authority (CA), you must provide a CSR. After the CA approves your certificate application, the CA uses the private key of the root CA to sign your CSR and generates a public key file. The public key file is the SSL certificate that the CA issues to you. The private key of the SSL certificate is generated when you create the CSR.
//
// @param request - CreateCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCsrResponse
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

// Summary:
//
// Creates a certificate signing request (CSR). A CSR file contains the information about an SSL certificate that you want to apply for. The information includes the domain names that you want to bind to the certificate and the name and the geographical location of the certificate holder. When you submit a certificate application to a certificate authority (CA), you must provide a CSR. After the CA approves your certificate application, the CA uses the private key of the root CA to sign your CSR and generates a public key file. The public key file is the SSL certificate that the CA issues to you. The private key of the SSL certificate is generated when you create the CSR.
//
// @param request - CreateCsrRequest
//
// @return CreateCsrResponse
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

// Summary:
//
// Creates a deployment task.
//
// @param request - CreateDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentJobResponse
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

// Summary:
//
// Creates a deployment task.
//
// @param request - CreateDeploymentJobRequest
//
// @return CreateDeploymentJobResponse
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

// @param request - CreateWHClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWHClientCertificateResponse
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

// @param request - CreateWHClientCertificateRequest
//
// @return CreateWHClientCertificateResponse
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

// Summary:
//
// Decrypts a certificate in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DecryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecryptResponse
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

// Summary:
//
// Decrypts a certificate in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DecryptRequest
//
// @return DecryptResponse
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

// Summary:
//
// Deletes an order in which the application for a domain validated (DV) certificate failed.
//
// Description:
//
// You can call this operation to delete a certificate application order only in the following scenarios:
//
// 	- The status of the order is **review failed**. You have called the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to query the status of the certificate application order and the value of the **Type*	- parameter is **verify_fail**.
//
// 	- The status of the order is **pending application**. You have called the [CancelOrderRequest](https://help.aliyun.com/document_detail/455299.html) operation to cancel a certificate application order whose status is pending review or being reviewed. The status of the certificate application order that is canceled in this case changes to **pending application**.
//
// @param request - DeleteCertificateRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCertificateRequestResponse
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

// Summary:
//
// Deletes an order in which the application for a domain validated (DV) certificate failed.
//
// Description:
//
// You can call this operation to delete a certificate application order only in the following scenarios:
//
// 	- The status of the order is **review failed**. You have called the [DescribeCertificateState](https://help.aliyun.com/document_detail/455800.html) operation to query the status of the certificate application order and the value of the **Type*	- parameter is **verify_fail**.
//
// 	- The status of the order is **pending application**. You have called the [CancelOrderRequest](https://help.aliyun.com/document_detail/455299.html) operation to cancel a certificate application order whose status is pending review or being reviewed. The status of the certificate application order that is canceled in this case changes to **pending application**.
//
// @param request - DeleteCertificateRequestRequest
//
// @return DeleteCertificateRequestResponse
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

// Summary:
//
// Deletes a certificate signing request (CSR) file.
//
// @param request - DeleteCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCsrResponse
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

// Summary:
//
// Deletes a certificate signing request (CSR) file.
//
// @param request - DeleteCsrRequest
//
// @return DeleteCsrResponse
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

// Summary:
//
// Deletes a deployment task.
//
// @param request - DeleteDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentJobResponse
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

// Summary:
//
// Deletes a deployment task.
//
// @param request - DeleteDeploymentJobRequest
//
// @return DeleteDeploymentJobResponse
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

// @param request - DeletePCACertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePCACertResponse
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

// @param request - DeletePCACertRequest
//
// @return DeletePCACertResponse
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

// Summary:
//
// Deletes an expired or uploaded certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteUserCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserCertificateResponse
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

// Summary:
//
// Deletes an expired or uploaded certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteUserCertificateRequest
//
// @return DeleteUserCertificateResponse
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

// Summary:
//
// Deletes the worker of a deployment task.
//
// @param request - DeleteWorkerResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkerResourceResponse
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

// Summary:
//
// Deletes the worker of a deployment task.
//
// @param request - DeleteWorkerResourceRequest
//
// @return DeleteWorkerResourceResponse
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

// Summary:
//
// Queries the status of a specified certificate application order.
//
// Description:
//
// If you do not complete the verification of the domain name ownership after you submit a certificate application, you can call this operation to obtain the information that is required to complete the verification. You can complete the verification of the domain name ownership based on the data returned. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on the DNS server.
//
// The certificate authority (CA) reviews your certificate application only after you complete the verification of the domain name ownership. After the CA approves your certificate application, the CA issues the certificate. If a certificate is issued, you can call this operation to obtain the CA certificate and private key of the certificate.
//
// @param request - DescribeCertificateStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertificateStateResponse
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

// Summary:
//
// Queries the status of a specified certificate application order.
//
// Description:
//
// If you do not complete the verification of the domain name ownership after you submit a certificate application, you can call this operation to obtain the information that is required to complete the verification. You can complete the verification of the domain name ownership based on the data returned. If you use the DNS verification method, you must complete the verification on the management platform of the domain name. If you use the file verification method, you must complete the verification on the DNS server.
//
// The certificate authority (CA) reviews your certificate application only after you complete the verification of the domain name ownership. After the CA approves your certificate application, the CA issues the certificate. If a certificate is issued, you can call this operation to obtain the CA certificate and private key of the certificate.
//
// @param request - DescribeCertificateStateRequest
//
// @return DescribeCertificateStateResponse
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

// Summary:
//
// Queries the number of third-party cloud resources on which you deployed certificates by using a multi-cloud deployment task.
//
// @param request - DescribeCloudResourceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudResourceStatusResponse
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

// Summary:
//
// Queries the number of third-party cloud resources on which you deployed certificates by using a multi-cloud deployment task.
//
// @param request - DescribeCloudResourceStatusRequest
//
// @return DescribeCloudResourceStatusResponse
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

// Summary:
//
// Queries the details of a deployment task. You can call the CreateDeploymentJob operation to create a deployment task and obtain the ID of the task.
//
// @param request - DescribeDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeploymentJobResponse
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

// Summary:
//
// Queries the details of a deployment task. You can call the CreateDeploymentJob operation to create a deployment task and obtain the ID of the task.
//
// @param request - DescribeDeploymentJobRequest
//
// @return DescribeDeploymentJobResponse
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

// Summary:
//
// Queries the number of worker tasks in a deployment task.
//
// @param request - DescribeDeploymentJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeploymentJobStatusResponse
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

// Summary:
//
// Queries the number of worker tasks in a deployment task.
//
// @param request - DescribeDeploymentJobStatusRequest
//
// @return DescribeDeploymentJobStatusResponse
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

// Summary:
//
// Queries the quota for domain validated (DV) certificates that you purchase and the quota usage.
//
// @param request - DescribePackageStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackageStateResponse
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

// Summary:
//
// Queries the quota for domain validated (DV) certificates that you purchase and the quota usage.
//
// @param request - DescribePackageStateRequest
//
// @return DescribePackageStateResponse
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

// Summary:
//
// Encrypts a certificate in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - EncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncryptResponse
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

// Summary:
//
// Encrypts a certificate in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - EncryptRequest
//
// @return EncryptResponse
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

// Summary:
//
// Queries the quota for certificate repositories.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetCertWarehouseQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCertWarehouseQuotaResponse
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

// Summary:
//
// Queries the quota for certificate repositories.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @return GetCertWarehouseQuotaResponse
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

// Summary:
//
// Obtains the content of a certificate signing request (CSR) file.
//
// @param request - GetCsrDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCsrDetailResponse
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

// Summary:
//
// Obtains the content of a certificate signing request (CSR) file.
//
// @param request - GetCsrDetailRequest
//
// @return GetCsrDetailResponse
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

// Summary:
//
// Queries the details of a certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetUserCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserCertificateDetailResponse
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

// Summary:
//
// Queries the details of a certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetUserCertificateDetailRequest
//
// @return GetUserCertificateDetailResponse
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

// Summary:
//
// Queries the certificates in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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

// Summary:
//
// Queries the certificates in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries certificate repositories.
//
// Description:
//
// You can call the ListCertWarehouse operation to query certificate repositories.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListCertWarehouseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCertWarehouseResponse
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

// Summary:
//
// Queries certificate repositories.
//
// Description:
//
// You can call the ListCertWarehouse operation to query certificate repositories.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListCertWarehouseRequest
//
// @return ListCertWarehouseResponse
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

// Summary:
//
// Queries a list of AccessKey pairs for multi-cloud deployment.
//
// @param request - ListCloudAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAccessResponse
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

// Summary:
//
// Queries a list of AccessKey pairs for multi-cloud deployment.
//
// @param request - ListCloudAccessRequest
//
// @return ListCloudAccessResponse
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

// Summary:
//
// Queries the certificate resources of a cloud service provider and cloud services.
//
// @param tmpReq - ListCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudResourcesResponse
func (client *Client) ListCloudResourcesWithOptions(tmpReq *ListCloudResourcesRequest, runtime *util.RuntimeOptions) (_result *ListCloudResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCloudResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CertIds)) {
		request.CertIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CertIds, tea.String("CertIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIdsShrink)) {
		query["CertIds"] = request.CertIdsShrink
	}

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

// Summary:
//
// Queries the certificate resources of a cloud service provider and cloud services.
//
// @param request - ListCloudResourcesRequest
//
// @return ListCloudResourcesResponse
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

// Summary:
//
// Queries a list of contacts.
//
// @param request - ListContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactResponse
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

// Summary:
//
// Queries a list of contacts.
//
// @param request - ListContactRequest
//
// @return ListContactResponse
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

// Summary:
//
// Queries certificate signing requests (CSRs).
//
// @param request - ListCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCsrResponse
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

// Summary:
//
// Queries certificate signing requests (CSRs).
//
// @param request - ListCsrRequest
//
// @return ListCsrResponse
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

// Summary:
//
// Queries a list of deployment tasks that are created.
//
// @param request - ListDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentJobResponse
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

// Summary:
//
// Queries a list of deployment tasks that are created.
//
// @param request - ListDeploymentJobRequest
//
// @return ListDeploymentJobResponse
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

// Summary:
//
// Queries the basic information about a deployment task. After you create a deployment task, you can call this operation to obtain the basic information about the deployment task, including the instance ID, type, and name of the certificate.
//
// @param request - ListDeploymentJobCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentJobCertResponse
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

// Summary:
//
// Queries the basic information about a deployment task. After you create a deployment task, you can call this operation to obtain the basic information about the deployment task, including the instance ID, type, and name of the certificate.
//
// @param request - ListDeploymentJobCertRequest
//
// @return ListDeploymentJobCertResponse
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

// Summary:
//
// Queries the cloud resources of cloud services in a deployment task.
//
// @param request - ListDeploymentJobResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentJobResourceResponse
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

// Summary:
//
// Queries the cloud resources of cloud services in a deployment task.
//
// @param request - ListDeploymentJobResourceRequest
//
// @return ListDeploymentJobResourceResponse
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

// Summary:
//
// Queries the certificates or certificate orders of users.
//
// Description:
//
// You can call the ListUserCertificateOrder operation to query the certificates or certificate orders of users. If you set OrderType to CERT or UPLOAD, certificates are returned. If you set OrderType to CPACK or BUY, certificate orders are returned.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListUserCertificateOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserCertificateOrderResponse
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

// Summary:
//
// Queries the certificates or certificate orders of users.
//
// Description:
//
// You can call the ListUserCertificateOrder operation to query the certificates or certificate orders of users. If you set OrderType to CERT or UPLOAD, certificates are returned. If you set OrderType to CPACK or BUY, certificate orders are returned.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListUserCertificateOrderRequest
//
// @return ListUserCertificateOrderResponse
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

// Summary:
//
// Queries the details about the worker tasks of a deployment task. Alibaba Cloud allows you to deploy multiple certificates at a time. Therefore, a deployment task may include multiple worker tasks in multiple cloud services. A worker task refers to a task that deploys a certificate to a cloud resource in a cloud service.
//
// @param request - ListWorkerResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkerResourceResponse
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

// Summary:
//
// Queries the details about the worker tasks of a deployment task. Alibaba Cloud allows you to deploy multiple certificates at a time. Therefore, a deployment task may include multiple worker tasks in multiple cloud services. A worker task refers to a task that deploys a certificate to a cloud resource in a cloud service.
//
// @param request - ListWorkerResourceRequest
//
// @return ListWorkerResourceResponse
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

// Summary:
//
// Changes the resource group to which a certificate or certificate order belongs.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2020-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a certificate or certificate order belongs.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a renewal application for an issued certificate.
//
// Description:
//
// You can call the RenewCertificateOrderForPackageRequest operation to submit a renewal application for a certificate only when the order of the certificate is in the expiring state. After the renewal is complete, a new certificate order whose status is pending application is generated. You must submit a certificate application for the new certificate order and install the new certificate after the new certificate is issued.
//
// >  You can call the [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html) operation to query the status of a certificate application order. If the value of the **Type*	- response parameter is **certificate**, the certificate is issued.
//
// @param request - RenewCertificateOrderForPackageRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewCertificateOrderForPackageRequestResponse
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

// Summary:
//
// Submits a renewal application for an issued certificate.
//
// Description:
//
// You can call the RenewCertificateOrderForPackageRequest operation to submit a renewal application for a certificate only when the order of the certificate is in the expiring state. After the renewal is complete, a new certificate order whose status is pending application is generated. You must submit a certificate application for the new certificate order and install the new certificate after the new certificate is issued.
//
// >  You can call the [DescribeCertificateState](https://help.aliyun.com/document_detail/164111.html) operation to query the status of a certificate application order. If the value of the **Type*	- response parameter is **certificate**, the certificate is issued.
//
// @param request - RenewCertificateOrderForPackageRequestRequest
//
// @return RenewCertificateOrderForPackageRequestResponse
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

// Summary:
//
// Revokes a client certificate or a server certificate in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - RevokeWHClientCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeWHClientCertificateResponse
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

// Summary:
//
// Revokes a client certificate or a server certificate in a certificate repository.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - RevokeWHClientCertificateRequest
//
// @return RevokeWHClientCertificateResponse
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

// Summary:
//
// Signs a private certificate in a certificate application repository.
//
// Description:
//
// You can call the Sign operation to sign a private certificate in a certificate application repository.
//
// ### Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SignResponse
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

// Summary:
//
// Signs a private certificate in a certificate application repository.
//
// Description:
//
// You can call the Sign operation to sign a private certificate in a certificate application repository.
//
// ### Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SignRequest
//
// @return SignResponse
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

// Summary:
//
// Updates the private key of a certificate signing request (CSR).
//
// @param request - UpdateCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCsrResponse
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

// Summary:
//
// Updates the private key of a certificate signing request (CSR).
//
// @param request - UpdateCsrRequest
//
// @return UpdateCsrResponse
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

// Summary:
//
// Updates a deployment task.
//
// @param request - UpdateDeploymentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentJobResponse
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

// Summary:
//
// Updates a deployment task.
//
// @param request - UpdateDeploymentJobRequest
//
// @return UpdateDeploymentJobResponse
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

// Summary:
//
// Updates the status of a deployment task.
//
// @param request - UpdateDeploymentJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentJobStatusResponse
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

// Summary:
//
// Updates the status of a deployment task.
//
// @param request - UpdateDeploymentJobStatusRequest
//
// @return UpdateDeploymentJobStatusResponse
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

// Summary:
//
// Rolls back or executes a worker task in a deployment task.
//
// @param request - UpdateWorkerResourceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkerResourceStatusResponse
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

// Summary:
//
// Rolls back or executes a worker task in a deployment task.
//
// @param request - UpdateWorkerResourceStatusRequest
//
// @return UpdateWorkerResourceStatusResponse
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

// Summary:
//
// Uploads a certificate signing request (CSR) file
//
// @param request - UploadCsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadCsrResponse
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

// Summary:
//
// Uploads a certificate signing request (CSR) file
//
// @param request - UploadCsrRequest
//
// @return UploadCsrResponse
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

// Summary:
//
// The private key of the certificate.
//
// Description:
//
// You can call this operation to upload a private certificate to a certificate repository.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UploadPCACertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadPCACertResponse
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

// Summary:
//
// The private key of the certificate.
//
// Description:
//
// You can call this operation to upload a private certificate to a certificate repository.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UploadPCACertRequest
//
// @return UploadPCACertResponse
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

// Summary:
//
// Uploads a certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UploadUserCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadUserCertificateResponse
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

// Summary:
//
// Uploads a certificate.
//
// Description:
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UploadUserCertificateRequest
//
// @return UploadUserCertificateResponse
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

// Summary:
//
// Verifies the signature of a private certificate in a certificate application repository.
//
// Description:
//
// You can call the Verify operation to verify the signature of a private certificate in a certificate application repository.
//
// ### Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyResponse
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

// Summary:
//
// Verifies the signature of a private certificate in a certificate application repository.
//
// Description:
//
// You can call the Verify operation to verify the signature of a private certificate in a certificate application repository.
//
// ### Limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifyRequest
//
// @return VerifyResponse
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
