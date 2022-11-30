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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelCertificateForPackageRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelOrderRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CompanyName  *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	Csr          *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone        *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ProductCode  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Username     *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCertificateForPackageRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone        *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ProductCode  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Username     *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCertificateRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Csr          *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone        *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ProductCode  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Username     *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCertificateWithCsrRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWHClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DecryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCertificateRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePCACertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeCertificateStateRequest struct {
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
	Certificate  *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	PrivateKey   *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RecordDomain *string `json:"RecordDomain,omitempty" xml:"RecordDomain,omitempty"`
	RecordType   *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	RecordValue  *string `json:"RecordValue,omitempty" xml:"RecordValue,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri          *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCertificateStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribePackageStateRequest struct {
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
	IssuedCount *int64  `json:"IssuedCount,omitempty" xml:"IssuedCount,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UsedCount   *int64  `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePackageStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EncryptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCertWarehouseQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetUserCertificateDetailRequest struct {
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
}

func (s GetUserCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetUserCertificateDetailRequest) SetCertId(v int64) *GetUserCertificateDetailRequest {
	s.CertId = &v
	return s
}

type GetUserCertificateDetailResponseBody struct {
	BuyInAliyun       *bool   `json:"BuyInAliyun,omitempty" xml:"BuyInAliyun,omitempty"`
	Cert              *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	City              *string `json:"City,omitempty" xml:"City,omitempty"`
	Common            *string `json:"Common,omitempty" xml:"Common,omitempty"`
	Country           *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EncryptCert       *string `json:"EncryptCert,omitempty" xml:"EncryptCert,omitempty"`
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	EndDate           *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Expired           *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Fingerprint       *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Issuer            *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Key               *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderId           *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrgName           *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	Province          *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Sans              *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SignCert          *string `json:"SignCert,omitempty" xml:"SignCert,omitempty"`
	SignPrivateKey    *string `json:"SignPrivateKey,omitempty" xml:"SignPrivateKey,omitempty"`
	StartDate         *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetUserCertificateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserCertificateDetailResponseBody) GoString() string {
	return s.String()
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CurrentPage *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	KeyWord     *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	ShowSize    *int64  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	SourceType  *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WarehouseId *int64  `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
}

func (s ListCertRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCertRequest) GoString() string {
	return s.String()
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
	CertList    []*ListCertResponseBodyCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	CurrentPage *int64                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize    *int64                          `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount  *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AfterDate       *int64  `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	BeforeDate      *int64  `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CommonName      *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	ExistPrivateKey *bool   `json:"ExistPrivateKey,omitempty" xml:"ExistPrivateKey,omitempty"`
	Identifier      *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Issuer          *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Sans            *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WhId            *int64  `json:"WhId,omitempty" xml:"WhId,omitempty"`
	WhInstanceId    *string `json:"WhInstanceId,omitempty" xml:"WhInstanceId,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCertWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListUserCertificateOrderRequest struct {
	CurrentPage     *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Keyword         *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OrderType       *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ShowSize        *int64  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	CertificateOrderList []*ListUserCertificateOrderResponseBodyCertificateOrderList `json:"CertificateOrderList,omitempty" xml:"CertificateOrderList,omitempty" type:"Repeated"`
	CurrentPage          *int64                                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize             *int64                                                      `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount           *int64                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Algorithm       *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	AliyunOrderId   *int64  `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	BuyDate         *int64  `json:"BuyDate,omitempty" xml:"BuyDate,omitempty"`
	CertEndTime     *int64  `json:"CertEndTime,omitempty" xml:"CertEndTime,omitempty"`
	CertStartTime   *int64  `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	CertType        *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertificateId   *int64  `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	City            *string `json:"City,omitempty" xml:"City,omitempty"`
	CommonName      *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	Country         *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainCount     *int64  `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	DomainType      *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EndDate         *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Expired         *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Fingerprint     *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Issuer          *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderId         *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrgName         *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	PartnerOrderId  *string `json:"PartnerOrderId,omitempty" xml:"PartnerOrderId,omitempty"`
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName     *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Province        *string `json:"Province,omitempty" xml:"Province,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RootBrand       *string `json:"RootBrand,omitempty" xml:"RootBrand,omitempty"`
	Sans            *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SerialNo        *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Sha2            *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartDate       *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TrusteeStatus   *string `json:"TrusteeStatus,omitempty" xml:"TrusteeStatus,omitempty"`
	Upload          *bool   `json:"Upload,omitempty" xml:"Upload,omitempty"`
	WildDomainCount *int64  `json:"WildDomainCount,omitempty" xml:"WildDomainCount,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserCertificateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type RenewCertificateOrderForPackageRequestRequest struct {
	Csr     *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	OrderId *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenewCertificateOrderForPackageRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevokeWHClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SignResponseBody  `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UploadPCACertRequest struct {
	Cert        *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PrivateKey  *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	WarehouseId *int64  `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadPCACertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Cert              *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	EncryptCert       *string `json:"EncryptCert,omitempty" xml:"EncryptCert,omitempty"`
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	Key               *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SignCert          *string `json:"SignCert,omitempty" xml:"SignCert,omitempty"`
	SignPrivateKey    *string `json:"SignPrivateKey,omitempty" xml:"SignPrivateKey,omitempty"`
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
	CertId    *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadUserCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) GetUserCertificateDetailWithOptions(request *GetUserCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *GetUserCertificateDetailResponse, _err error) {
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

func (client *Client) ListCertWithOptions(request *ListCertRequest, runtime *util.RuntimeOptions) (_result *ListCertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
