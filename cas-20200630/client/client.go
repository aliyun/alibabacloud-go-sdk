// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateCertificateWithExtensionRequest struct {
	AfterTime                *int64  `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	AlgorithmKeySize         *int32  `json:"AlgorithmKeySize,omitempty" xml:"AlgorithmKeySize,omitempty"`
	AliasName                *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	AppendCrl                *bool   `json:"AppendCrl,omitempty" xml:"AppendCrl,omitempty"`
	BasicConstraintsCritical *bool   `json:"BasicConstraintsCritical,omitempty" xml:"BasicConstraintsCritical,omitempty"`
	BeforeTime               *int64  `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	CertType                 *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CommonName               *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CountryCode              *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	CsrPemString             *string `json:"CsrPemString,omitempty" xml:"CsrPemString,omitempty"`
	Locality                 *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Organization             *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit         *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	ParentIdentifier         *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	Sans                     *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	State                    *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateCertificateWithExtensionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateWithExtensionRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithExtensionRequest) SetAfterTime(v int64) *CreateCertificateWithExtensionRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetAlgorithmKeySize(v int32) *CreateCertificateWithExtensionRequest {
	s.AlgorithmKeySize = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetAliasName(v string) *CreateCertificateWithExtensionRequest {
	s.AliasName = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetAppendCrl(v bool) *CreateCertificateWithExtensionRequest {
	s.AppendCrl = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetBasicConstraintsCritical(v bool) *CreateCertificateWithExtensionRequest {
	s.BasicConstraintsCritical = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetBeforeTime(v int64) *CreateCertificateWithExtensionRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetCertType(v string) *CreateCertificateWithExtensionRequest {
	s.CertType = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetCommonName(v string) *CreateCertificateWithExtensionRequest {
	s.CommonName = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetCountryCode(v string) *CreateCertificateWithExtensionRequest {
	s.CountryCode = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetCsrPemString(v string) *CreateCertificateWithExtensionRequest {
	s.CsrPemString = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetLocality(v string) *CreateCertificateWithExtensionRequest {
	s.Locality = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetOrganization(v string) *CreateCertificateWithExtensionRequest {
	s.Organization = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetOrganizationUnit(v string) *CreateCertificateWithExtensionRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetParentIdentifier(v string) *CreateCertificateWithExtensionRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetSans(v string) *CreateCertificateWithExtensionRequest {
	s.Sans = &v
	return s
}

func (s *CreateCertificateWithExtensionRequest) SetState(v string) *CreateCertificateWithExtensionRequest {
	s.State = &v
	return s
}

type CreateCertificateWithExtensionResponseBody struct {
	Certificate      *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	Identifier       *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertificateWithExtensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateWithExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithExtensionResponseBody) SetCertificate(v string) *CreateCertificateWithExtensionResponseBody {
	s.Certificate = &v
	return s
}

func (s *CreateCertificateWithExtensionResponseBody) SetCertificateChain(v string) *CreateCertificateWithExtensionResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateCertificateWithExtensionResponseBody) SetIdentifier(v string) *CreateCertificateWithExtensionResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateCertificateWithExtensionResponseBody) SetRequestId(v string) *CreateCertificateWithExtensionResponseBody {
	s.RequestId = &v
	return s
}

type CreateCertificateWithExtensionResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCertificateWithExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCertificateWithExtensionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateWithExtensionResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithExtensionResponse) SetHeaders(v map[string]*string) *CreateCertificateWithExtensionResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateWithExtensionResponse) SetBody(v *CreateCertificateWithExtensionResponseBody) *CreateCertificateWithExtensionResponse {
	s.Body = v
	return s
}

type CreateClientCertificateRequest struct {
	AfterTime        *int64  `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeTime       *int64  `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	Days             *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	SanType          *int32  `json:"SanType,omitempty" xml:"SanType,omitempty"`
	SanValue         *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
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

func (s *CreateClientCertificateRequest) SetDays(v int32) *CreateClientCertificateRequest {
	s.Days = &v
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

type CreateClientCertificateResponseBody struct {
	CertificateChain      *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	Identifier            *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	ParentX509Certificate *string `json:"ParentX509Certificate,omitempty" xml:"ParentX509Certificate,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootX509Certificate   *string `json:"RootX509Certificate,omitempty" xml:"RootX509Certificate,omitempty"`
	X509Certificate       *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
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

func (s *CreateClientCertificateResponseBody) SetParentX509Certificate(v string) *CreateClientCertificateResponseBody {
	s.ParentX509Certificate = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetRequestId(v string) *CreateClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetRootX509Certificate(v string) *CreateClientCertificateResponseBody {
	s.RootX509Certificate = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetX509Certificate(v string) *CreateClientCertificateResponseBody {
	s.X509Certificate = &v
	return s
}

type CreateClientCertificateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateClientCertificateResponse) SetBody(v *CreateClientCertificateResponseBody) *CreateClientCertificateResponse {
	s.Body = v
	return s
}

type CreateClientCertificateWithCsrRequest struct {
	AfterTime        *int64  `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	BeforeTime       *int64  `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	Csr              *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	Days             *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	SanType          *int32  `json:"SanType,omitempty" xml:"SanType,omitempty"`
	SanValue         *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
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

func (s *CreateClientCertificateWithCsrRequest) SetBeforeTime(v int64) *CreateClientCertificateWithCsrRequest {
	s.BeforeTime = &v
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

type CreateClientCertificateWithCsrResponseBody struct {
	CertificateChain      *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	Identifier            *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	ParentX509Certificate *string `json:"ParentX509Certificate,omitempty" xml:"ParentX509Certificate,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootX509Certificate   *string `json:"RootX509Certificate,omitempty" xml:"RootX509Certificate,omitempty"`
	X509Certificate       *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
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

func (s *CreateClientCertificateWithCsrResponseBody) SetParentX509Certificate(v string) *CreateClientCertificateWithCsrResponseBody {
	s.ParentX509Certificate = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetRequestId(v string) *CreateClientCertificateWithCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetRootX509Certificate(v string) *CreateClientCertificateWithCsrResponseBody {
	s.RootX509Certificate = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetX509Certificate(v string) *CreateClientCertificateWithCsrResponseBody {
	s.X509Certificate = &v
	return s
}

type CreateClientCertificateWithCsrResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClientCertificateWithCsrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateClientCertificateWithCsrResponse) SetBody(v *CreateClientCertificateWithCsrResponseBody) *CreateClientCertificateWithCsrResponse {
	s.Body = v
	return s
}

type CreateRevokeClientCertificateRequest struct {
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRevokeClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateRevokeClientCertificateResponse) SetBody(v *CreateRevokeClientCertificateResponseBody) *CreateRevokeClientCertificateResponse {
	s.Body = v
	return s
}

type CreateRootCACertificateRequest struct {
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CountryCode      *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	Years            *int32  `json:"Years,omitempty" xml:"Years,omitempty"`
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
	Certificate      *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	Identifier       *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRootCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateRootCACertificateResponse) SetBody(v *CreateRootCACertificateResponseBody) *CreateRootCACertificateResponse {
	s.Body = v
	return s
}

type CreateServerCertificateRequest struct {
	AfterTime        *int64  `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeTime       *int64  `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	Days             *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	Domain           *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
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

func (s *CreateServerCertificateRequest) SetDays(v int32) *CreateServerCertificateRequest {
	s.Days = &v
	return s
}

func (s *CreateServerCertificateRequest) SetDomain(v string) *CreateServerCertificateRequest {
	s.Domain = &v
	return s
}

func (s *CreateServerCertificateRequest) SetParentIdentifier(v string) *CreateServerCertificateRequest {
	s.ParentIdentifier = &v
	return s
}

type CreateServerCertificateResponseBody struct {
	CertificateChain      *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	Identifier            *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	ParentX509Certificate *string `json:"ParentX509Certificate,omitempty" xml:"ParentX509Certificate,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootX509Certificate   *string `json:"RootX509Certificate,omitempty" xml:"RootX509Certificate,omitempty"`
	X509Certificate       *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
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

func (s *CreateServerCertificateResponseBody) SetParentX509Certificate(v string) *CreateServerCertificateResponseBody {
	s.ParentX509Certificate = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetRequestId(v string) *CreateServerCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetRootX509Certificate(v string) *CreateServerCertificateResponseBody {
	s.RootX509Certificate = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetX509Certificate(v string) *CreateServerCertificateResponseBody {
	s.X509Certificate = &v
	return s
}

type CreateServerCertificateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServerCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateServerCertificateResponse) SetBody(v *CreateServerCertificateResponseBody) *CreateServerCertificateResponse {
	s.Body = v
	return s
}

type CreateServerCertificateWithCsrRequest struct {
	AfterTime        *int64  `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	BeforeTime       *int64  `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	Csr              *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	Days             *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	Domain           *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
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

func (s *CreateServerCertificateWithCsrRequest) SetBeforeTime(v int64) *CreateServerCertificateWithCsrRequest {
	s.BeforeTime = &v
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

func (s *CreateServerCertificateWithCsrRequest) SetParentIdentifier(v string) *CreateServerCertificateWithCsrRequest {
	s.ParentIdentifier = &v
	return s
}

type CreateServerCertificateWithCsrResponseBody struct {
	CertificateChain      *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	Identifier            *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	ParentX509Certificate *string `json:"ParentX509Certificate,omitempty" xml:"ParentX509Certificate,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootX509Certificate   *string `json:"RootX509Certificate,omitempty" xml:"RootX509Certificate,omitempty"`
	X509Certificate       *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
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

func (s *CreateServerCertificateWithCsrResponseBody) SetParentX509Certificate(v string) *CreateServerCertificateWithCsrResponseBody {
	s.ParentX509Certificate = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetRequestId(v string) *CreateServerCertificateWithCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetRootX509Certificate(v string) *CreateServerCertificateWithCsrResponseBody {
	s.RootX509Certificate = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetX509Certificate(v string) *CreateServerCertificateWithCsrResponseBody {
	s.X509Certificate = &v
	return s
}

type CreateServerCertificateWithCsrResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServerCertificateWithCsrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateServerCertificateWithCsrResponse) SetBody(v *CreateServerCertificateWithCsrResponseBody) *CreateServerCertificateWithCsrResponse {
	s.Body = v
	return s
}

type CreateSubCACertificateRequest struct {
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CountryCode      *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	Years            *int32  `json:"Years,omitempty" xml:"Years,omitempty"`
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

func (s *CreateSubCACertificateRequest) SetState(v string) *CreateSubCACertificateRequest {
	s.State = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetYears(v int32) *CreateSubCACertificateRequest {
	s.Years = &v
	return s
}

type CreateSubCACertificateResponseBody struct {
	Certificate      *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	Identifier       *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSubCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateSubCACertificateResponse) SetBody(v *CreateSubCACertificateResponseBody) *CreateSubCACertificateResponse {
	s.Body = v
	return s
}

type DeleteClientCertificateRequest struct {
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteClientCertificateResponse) SetBody(v *DeleteClientCertificateResponseBody) *DeleteClientCertificateResponse {
	s.Body = v
	return s
}

type DescribeCACertificateRequest struct {
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
	Certificate *DescribeCACertificateResponseBodyCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty" type:"Struct"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type DescribeCACertificateResponseBodyCertificate struct {
	AfterDate        *int64  `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeDate       *int64  `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CertificateType  *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CountryCode      *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Identifier       *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	KeySize          *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	Sans             *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Sha2             *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	SignAlgorithm    *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubjectDN        *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	X509Certificate  *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
	Years            *int32  `json:"Years,omitempty" xml:"Years,omitempty"`
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

func (s *DescribeCACertificateResponseBodyCertificate) SetYears(v int32) *DescribeCACertificateResponseBodyCertificate {
	s.Years = &v
	return s
}

type DescribeCACertificateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCACertificateResponse) SetBody(v *DescribeCACertificateResponseBody) *DescribeCACertificateResponse {
	s.Body = v
	return s
}

type DescribeCACertificateCountResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCACertificateCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCACertificateCountResponse) SetBody(v *DescribeCACertificateCountResponseBody) *DescribeCACertificateCountResponse {
	s.Body = v
	return s
}

type DescribeCACertificateListRequest struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ShowSize    *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
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
	CertificateList []*DescribeCACertificateListResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	CurrentPage     *int32                                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount       *int32                                                  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	RequestId       *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize        *int32                                                  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount      *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AfterDate        *int64  `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeDate       *int64  `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CertificateType  *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CountryCode      *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Identifier       *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	KeySize          *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	Sans             *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Sha2             *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	SignAlgorithm    *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubjectDN        *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	X509Certificate  *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
	Years            *int32  `json:"Years,omitempty" xml:"Years,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCACertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCACertificateListResponse) SetBody(v *DescribeCACertificateListResponseBody) *DescribeCACertificateListResponse {
	s.Body = v
	return s
}

type DescribeCertificatePrivateKeyRequest struct {
	EncryptedCode *string `json:"EncryptedCode,omitempty" xml:"EncryptedCode,omitempty"`
	Identifier    *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
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
	EncryptedData *string `json:"EncryptedData,omitempty" xml:"EncryptedData,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCertificatePrivateKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCertificatePrivateKeyResponse) SetBody(v *DescribeCertificatePrivateKeyResponseBody) *DescribeCertificatePrivateKeyResponse {
	s.Body = v
	return s
}

type DescribeClientCertificateRequest struct {
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
	Certificate *DescribeClientCertificateResponseBodyCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty" type:"Struct"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AfterDate        *int64  `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeDate       *int64  `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CertificateType  *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CountryCode      *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Days             *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	Identifier       *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	KeySize          *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	Sans             *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Sha2             *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	SignAlgorithm    *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubjectDN        *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	X509Certificate  *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeClientCertificateResponse) SetBody(v *DescribeClientCertificateResponseBody) *DescribeClientCertificateResponse {
	s.Body = v
	return s
}

type DescribeClientCertificateForSerialNumberRequest struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DescribeClientCertificateForSerialNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateForSerialNumberRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateForSerialNumberRequest) SetSerialNumber(v string) *DescribeClientCertificateForSerialNumberRequest {
	s.SerialNumber = &v
	return s
}

type DescribeClientCertificateForSerialNumberResponseBody struct {
	CertificateList []*DescribeClientCertificateForSerialNumberResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	RequestId       *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientCertificateForSerialNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateForSerialNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateForSerialNumberResponseBody) SetCertificateList(v []*DescribeClientCertificateForSerialNumberResponseBodyCertificateList) *DescribeClientCertificateForSerialNumberResponseBody {
	s.CertificateList = v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBody) SetRequestId(v string) *DescribeClientCertificateForSerialNumberResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClientCertificateForSerialNumberResponseBodyCertificateList struct {
	AfterDate        *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeDate       *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CountryCode      *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Identifier       *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	KeySize          *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	Sans             *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Sha2             *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	SignAlgorithm    *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubjectDN        *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	X509Certificate  *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
	Years            *int32  `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s DescribeClientCertificateForSerialNumberResponseBodyCertificateList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetAfterDate(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.AfterDate = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetAlgorithm(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetBeforeDate(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.BeforeDate = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetCommonName(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetCountryCode(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.CountryCode = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetIdentifier(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Identifier = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetKeySize(v int32) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetLocality(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Locality = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetMd5(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Md5 = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetOrganization(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Organization = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetOrganizationUnit(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.OrganizationUnit = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetSans(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Sans = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetSerialNumber(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetSha2(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Sha2 = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetSignAlgorithm(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.SignAlgorithm = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetState(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.State = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetStatus(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Status = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetSubjectDN(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.SubjectDN = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetX509Certificate(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.X509Certificate = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetYears(v int32) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Years = &v
	return s
}

type DescribeClientCertificateForSerialNumberResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClientCertificateForSerialNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClientCertificateForSerialNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateForSerialNumberResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateForSerialNumberResponse) SetHeaders(v map[string]*string) *DescribeClientCertificateForSerialNumberResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponse) SetBody(v *DescribeClientCertificateForSerialNumberResponseBody) *DescribeClientCertificateForSerialNumberResponse {
	s.Body = v
	return s
}

type DescribeClientCertificateStatusRequest struct {
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
	CertificateStatus []*DescribeClientCertificateStatusResponseBodyCertificateStatus `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty" type:"Repeated"`
	RequestId         *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	RevokeTime   *int64  `json:"RevokeTime,omitempty" xml:"RevokeTime,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClientCertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeClientCertificateStatusResponse) SetBody(v *DescribeClientCertificateStatusResponseBody) *DescribeClientCertificateStatusResponse {
	s.Body = v
	return s
}

type DescribeClientCertificateStatusForSerialNumberRequest struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DescribeClientCertificateStatusForSerialNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateStatusForSerialNumberRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusForSerialNumberRequest) SetSerialNumber(v string) *DescribeClientCertificateStatusForSerialNumberRequest {
	s.SerialNumber = &v
	return s
}

type DescribeClientCertificateStatusForSerialNumberResponseBody struct {
	CertificateStatus []*DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty" type:"Repeated"`
	RequestId         *string                                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientCertificateStatusForSerialNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateStatusForSerialNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBody) SetCertificateStatus(v []*DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) *DescribeClientCertificateStatusForSerialNumberResponseBody {
	s.CertificateStatus = v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBody) SetRequestId(v string) *DescribeClientCertificateStatusForSerialNumberResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus struct {
	RevokeTime   *int64  `json:"RevokeTime,omitempty" xml:"RevokeTime,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) SetRevokeTime(v int64) *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus {
	s.RevokeTime = &v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) SetSerialNumber(v string) *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) SetStatus(v string) *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus {
	s.Status = &v
	return s
}

type DescribeClientCertificateStatusForSerialNumberResponse struct {
	Headers map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClientCertificateStatusForSerialNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClientCertificateStatusForSerialNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientCertificateStatusForSerialNumberResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusForSerialNumberResponse) SetHeaders(v map[string]*string) *DescribeClientCertificateStatusForSerialNumberResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponse) SetBody(v *DescribeClientCertificateStatusForSerialNumberResponseBody) *DescribeClientCertificateStatusForSerialNumberResponse {
	s.Body = v
	return s
}

type GetCAInstanceStatusRequest struct {
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
	InstanceStatusList []*GetCAInstanceStatusResponseBodyInstanceStatusList `json:"InstanceStatusList,omitempty" xml:"InstanceStatusList,omitempty" type:"Repeated"`
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AfterTime       *int64  `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	BeforeTime      *int64  `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	CertIssuedCount *int32  `json:"CertIssuedCount,omitempty" xml:"CertIssuedCount,omitempty"`
	CertTotalCount  *int32  `json:"CertTotalCount,omitempty" xml:"CertTotalCount,omitempty"`
	Identifier      *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UseExpireTime   *int64  `json:"UseExpireTime,omitempty" xml:"UseExpireTime,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCAInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetCAInstanceStatusResponse) SetBody(v *GetCAInstanceStatusResponseBody) *GetCAInstanceStatusResponse {
	s.Body = v
	return s
}

type ListCACertificateLogRequest struct {
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s ListCACertificateLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCACertificateLogRequest) GoString() string {
	return s.String()
}

func (s *ListCACertificateLogRequest) SetIdentifier(v string) *ListCACertificateLogRequest {
	s.Identifier = &v
	return s
}

type ListCACertificateLogResponseBody struct {
	LogList   []*ListCACertificateLogResponseBodyLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCACertificateLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCACertificateLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListCACertificateLogResponseBody) SetLogList(v []*ListCACertificateLogResponseBodyLogList) *ListCACertificateLogResponseBody {
	s.LogList = v
	return s
}

func (s *ListCACertificateLogResponseBody) SetRequestId(v string) *ListCACertificateLogResponseBody {
	s.RequestId = &v
	return s
}

type ListCACertificateLogResponseBodyLogList struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	OpType     *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
}

func (s ListCACertificateLogResponseBodyLogList) String() string {
	return tea.Prettify(s)
}

func (s ListCACertificateLogResponseBodyLogList) GoString() string {
	return s.String()
}

func (s *ListCACertificateLogResponseBodyLogList) SetContent(v string) *ListCACertificateLogResponseBodyLogList {
	s.Content = &v
	return s
}

func (s *ListCACertificateLogResponseBodyLogList) SetCreateTime(v int64) *ListCACertificateLogResponseBodyLogList {
	s.CreateTime = &v
	return s
}

func (s *ListCACertificateLogResponseBodyLogList) SetIdentifier(v string) *ListCACertificateLogResponseBodyLogList {
	s.Identifier = &v
	return s
}

func (s *ListCACertificateLogResponseBodyLogList) SetOpType(v string) *ListCACertificateLogResponseBodyLogList {
	s.OpType = &v
	return s
}

type ListCACertificateLogResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCACertificateLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCACertificateLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCACertificateLogResponse) GoString() string {
	return s.String()
}

func (s *ListCACertificateLogResponse) SetHeaders(v map[string]*string) *ListCACertificateLogResponse {
	s.Headers = v
	return s
}

func (s *ListCACertificateLogResponse) SetBody(v *ListCACertificateLogResponseBody) *ListCACertificateLogResponse {
	s.Body = v
	return s
}

type ListClientCertificateRequest struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ShowSize    *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
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
	CertificateList []*ListClientCertificateResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	CurrentPage     *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount       *int32                                              `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize        *int32                                              `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount      *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListClientCertificateResponseBody) SetTotalCount(v int32) *ListClientCertificateResponseBody {
	s.TotalCount = &v
	return s
}

type ListClientCertificateResponseBodyCertificateList struct {
	AfterDate        *int64  `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeDate       *int64  `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CertificateType  *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CountryCode      *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Days             *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	Identifier       *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	KeySize          *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	Sans             *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Sha2             *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	SignAlgorithm    *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubjectDN        *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	X509Certificate  *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListClientCertificateResponse) SetBody(v *ListClientCertificateResponseBody) *ListClientCertificateResponse {
	s.Body = v
	return s
}

type ListRevokeCertificateRequest struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ShowSize    *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
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
	CertificateList []*ListRevokeCertificateResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	CurrentPage     *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount       *int32                                              `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize        *int32                                              `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount      *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListRevokeCertificateResponseBody) SetTotalCount(v int32) *ListRevokeCertificateResponseBody {
	s.TotalCount = &v
	return s
}

type ListRevokeCertificateResponseBodyCertificateList struct {
	AfterDate        *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	Algorithm        *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeDate       *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CountryCode      *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Identifier       *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	KeySize          *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Locality         *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	RevokeDate       *string `json:"RevokeDate,omitempty" xml:"RevokeDate,omitempty"`
	Sans             *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Sha2             *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	SignAlgorithm    *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	SubjectDN        *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
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

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSubjectDN(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.SubjectDN = &v
	return s
}

type ListRevokeCertificateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRevokeCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRevokeCertificateResponse) SetBody(v *ListRevokeCertificateResponseBody) *ListRevokeCertificateResponse {
	s.Body = v
	return s
}

type UpdateCACertificateStatusRequest struct {
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCACertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) CreateCertificateWithExtensionWithOptions(request *CreateCertificateWithExtensionRequest, runtime *util.RuntimeOptions) (_result *CreateCertificateWithExtensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AfterTime"] = request.AfterTime
	query["AlgorithmKeySize"] = request.AlgorithmKeySize
	query["AliasName"] = request.AliasName
	query["AppendCrl"] = request.AppendCrl
	query["BasicConstraintsCritical"] = request.BasicConstraintsCritical
	query["BeforeTime"] = request.BeforeTime
	query["CertType"] = request.CertType
	query["CommonName"] = request.CommonName
	query["CountryCode"] = request.CountryCode
	query["CsrPemString"] = request.CsrPemString
	query["Locality"] = request.Locality
	query["Organization"] = request.Organization
	query["OrganizationUnit"] = request.OrganizationUnit
	query["ParentIdentifier"] = request.ParentIdentifier
	query["Sans"] = request.Sans
	query["State"] = request.State
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCertificateWithExtension"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCertificateWithExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCertificateWithExtension(request *CreateCertificateWithExtensionRequest) (_result *CreateCertificateWithExtensionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertificateWithExtensionResponse{}
	_body, _err := client.CreateCertificateWithExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClientCertificateWithOptions(request *CreateClientCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AfterTime"] = request.AfterTime
	query["Algorithm"] = request.Algorithm
	query["BeforeTime"] = request.BeforeTime
	query["CommonName"] = request.CommonName
	query["Days"] = request.Days
	query["ParentIdentifier"] = request.ParentIdentifier
	query["SanType"] = request.SanType
	query["SanValue"] = request.SanValue
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

func (client *Client) CreateClientCertificateWithCsrWithOptions(request *CreateClientCertificateWithCsrRequest, runtime *util.RuntimeOptions) (_result *CreateClientCertificateWithCsrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AfterTime"] = request.AfterTime
	query["BeforeTime"] = request.BeforeTime
	query["Csr"] = request.Csr
	query["Days"] = request.Days
	query["ParentIdentifier"] = request.ParentIdentifier
	query["SanType"] = request.SanType
	query["SanValue"] = request.SanValue
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

func (client *Client) CreateRevokeClientCertificateWithOptions(request *CreateRevokeClientCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateRevokeClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Identifier"] = request.Identifier
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

func (client *Client) CreateRootCACertificateWithOptions(request *CreateRootCACertificateRequest, runtime *util.RuntimeOptions) (_result *CreateRootCACertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Algorithm"] = request.Algorithm
	query["CommonName"] = request.CommonName
	query["CountryCode"] = request.CountryCode
	query["Locality"] = request.Locality
	query["Organization"] = request.Organization
	query["OrganizationUnit"] = request.OrganizationUnit
	query["State"] = request.State
	query["Years"] = request.Years
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

func (client *Client) CreateServerCertificateWithOptions(request *CreateServerCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateServerCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AfterTime"] = request.AfterTime
	query["Algorithm"] = request.Algorithm
	query["BeforeTime"] = request.BeforeTime
	query["CommonName"] = request.CommonName
	query["Days"] = request.Days
	query["Domain"] = request.Domain
	query["ParentIdentifier"] = request.ParentIdentifier
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

func (client *Client) CreateServerCertificateWithCsrWithOptions(request *CreateServerCertificateWithCsrRequest, runtime *util.RuntimeOptions) (_result *CreateServerCertificateWithCsrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AfterTime"] = request.AfterTime
	query["BeforeTime"] = request.BeforeTime
	query["Csr"] = request.Csr
	query["Days"] = request.Days
	query["Domain"] = request.Domain
	query["ParentIdentifier"] = request.ParentIdentifier
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

func (client *Client) CreateSubCACertificateWithOptions(request *CreateSubCACertificateRequest, runtime *util.RuntimeOptions) (_result *CreateSubCACertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Algorithm"] = request.Algorithm
	query["CommonName"] = request.CommonName
	query["CountryCode"] = request.CountryCode
	query["Locality"] = request.Locality
	query["Organization"] = request.Organization
	query["OrganizationUnit"] = request.OrganizationUnit
	query["ParentIdentifier"] = request.ParentIdentifier
	query["State"] = request.State
	query["Years"] = request.Years
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

func (client *Client) DeleteClientCertificateWithOptions(request *DeleteClientCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Identifier"] = request.Identifier
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

func (client *Client) DescribeCACertificateWithOptions(request *DescribeCACertificateRequest, runtime *util.RuntimeOptions) (_result *DescribeCACertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Identifier"] = request.Identifier
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

func (client *Client) DescribeCACertificateListWithOptions(request *DescribeCACertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeCACertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CurrentPage"] = request.CurrentPage
	query["ShowSize"] = request.ShowSize
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

func (client *Client) DescribeCertificatePrivateKeyWithOptions(request *DescribeCertificatePrivateKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeCertificatePrivateKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EncryptedCode"] = request.EncryptedCode
	query["Identifier"] = request.Identifier
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

func (client *Client) DescribeClientCertificateWithOptions(request *DescribeClientCertificateRequest, runtime *util.RuntimeOptions) (_result *DescribeClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Identifier"] = request.Identifier
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

func (client *Client) DescribeClientCertificateForSerialNumberWithOptions(request *DescribeClientCertificateForSerialNumberRequest, runtime *util.RuntimeOptions) (_result *DescribeClientCertificateForSerialNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["SerialNumber"] = request.SerialNumber
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClientCertificateForSerialNumber"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClientCertificateForSerialNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClientCertificateForSerialNumber(request *DescribeClientCertificateForSerialNumberRequest) (_result *DescribeClientCertificateForSerialNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClientCertificateForSerialNumberResponse{}
	_body, _err := client.DescribeClientCertificateForSerialNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClientCertificateStatusWithOptions(request *DescribeClientCertificateStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeClientCertificateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Identifier"] = request.Identifier
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

func (client *Client) DescribeClientCertificateStatusForSerialNumberWithOptions(request *DescribeClientCertificateStatusForSerialNumberRequest, runtime *util.RuntimeOptions) (_result *DescribeClientCertificateStatusForSerialNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["SerialNumber"] = request.SerialNumber
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClientCertificateStatusForSerialNumber"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClientCertificateStatusForSerialNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClientCertificateStatusForSerialNumber(request *DescribeClientCertificateStatusForSerialNumberRequest) (_result *DescribeClientCertificateStatusForSerialNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClientCertificateStatusForSerialNumberResponse{}
	_body, _err := client.DescribeClientCertificateStatusForSerialNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCAInstanceStatusWithOptions(request *GetCAInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *GetCAInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["InstanceId"] = request.InstanceId
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

func (client *Client) ListCACertificateLogWithOptions(request *ListCACertificateLogRequest, runtime *util.RuntimeOptions) (_result *ListCACertificateLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Identifier"] = request.Identifier
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCACertificateLog"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCACertificateLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCACertificateLog(request *ListCACertificateLogRequest) (_result *ListCACertificateLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCACertificateLogResponse{}
	_body, _err := client.ListCACertificateLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClientCertificateWithOptions(request *ListClientCertificateRequest, runtime *util.RuntimeOptions) (_result *ListClientCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CurrentPage"] = request.CurrentPage
	query["ShowSize"] = request.ShowSize
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

func (client *Client) ListRevokeCertificateWithOptions(request *ListRevokeCertificateRequest, runtime *util.RuntimeOptions) (_result *ListRevokeCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CurrentPage"] = request.CurrentPage
	query["ShowSize"] = request.ShowSize
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

func (client *Client) UpdateCACertificateStatusWithOptions(request *UpdateCACertificateStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateCACertificateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Identifier"] = request.Identifier
	query["Status"] = request.Status
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
