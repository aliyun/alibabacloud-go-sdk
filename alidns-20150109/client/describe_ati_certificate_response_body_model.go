// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeAtiCertificateResponseBodyAccessDeniedDetail) *DescribeAtiCertificateResponseBody
	GetAccessDeniedDetail() *DescribeAtiCertificateResponseBodyAccessDeniedDetail
	SetAgentHost(v string) *DescribeAtiCertificateResponseBody
	GetAgentHost() *string
	SetAgentId(v string) *DescribeAtiCertificateResponseBody
	GetAgentId() *string
	SetAlgorithm(v string) *DescribeAtiCertificateResponseBody
	GetAlgorithm() *string
	SetCertPem(v string) *DescribeAtiCertificateResponseBody
	GetCertPem() *string
	SetCertType(v string) *DescribeAtiCertificateResponseBody
	GetCertType() *string
	SetCreateTimestamp(v int64) *DescribeAtiCertificateResponseBody
	GetCreateTimestamp() *int64
	SetIssuer(v string) *DescribeAtiCertificateResponseBody
	GetIssuer() *string
	SetNotAfter(v string) *DescribeAtiCertificateResponseBody
	GetNotAfter() *string
	SetNotBefore(v string) *DescribeAtiCertificateResponseBody
	GetNotBefore() *string
	SetRequestId(v string) *DescribeAtiCertificateResponseBody
	GetRequestId() *string
	SetSan(v string) *DescribeAtiCertificateResponseBody
	GetSan() *string
	SetSerialNumber(v string) *DescribeAtiCertificateResponseBody
	GetSerialNumber() *string
	SetSource(v string) *DescribeAtiCertificateResponseBody
	GetSource() *string
	SetStatus(v string) *DescribeAtiCertificateResponseBody
	GetStatus() *string
	SetSubject(v string) *DescribeAtiCertificateResponseBody
	GetSubject() *string
	SetTlsaFingerprint(v string) *DescribeAtiCertificateResponseBody
	GetTlsaFingerprint() *string
	SetUpdateTimestamp(v int64) *DescribeAtiCertificateResponseBody
	GetUpdateTimestamp() *int64
}

type DescribeAtiCertificateResponseBody struct {
	// The access denied details. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *DescribeAtiCertificateResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The agent host address.
	//
	// example:
	//
	// www.example.com
	AgentHost *string `json:"AgentHost,omitempty" xml:"AgentHost,omitempty"`
	// The agent ID. After CNNIC real-name authentication, CNNIC assigns a unified agent ID. The agent ID serves as the unique identifier that binds the agent to the real-name authenticated registrant.
	//
	// example:
	//
	// csp01860716@5e0964fd-951c-4e45-b518-d09d4d2db8ca
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The encryption algorithm of the certificate.
	//
	// example:
	//
	// RSA-2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The certificate file in PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIDrzCCApegAwIBAgIQCDvgVpBCRrGhdWrJWZHHSjANBgkqhkiG9w0BAQUFAD...
	//
	// （中间是一长串经过 Base64 编码的数据）
	//
	// ...
	//
	// -----END CERTIFICATE-----
	CertPem *string `json:"CertPem,omitempty" xml:"CertPem,omitempty"`
	// The certificate type. Valid values:
	//
	// - Server: server certificate.
	//
	// - Identity: identity certificate.
	//
	// example:
	//
	// Server
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The creation time of the health check template (timestamp).
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The issuer information of the certificate, identified in Distinguished Names (DN) format.
	//
	// example:
	//
	// DigiCert Inc
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The end time of the certificate validity period.
	//
	// example:
	//
	// 2027-05-09 02:19:49
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The start time of the certificate validity period.
	//
	// example:
	//
	// 2026-01-26 02:16:38
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Subject Alternative Name (SAN), which specifies additional identities for the certificate subject.
	//
	// example:
	//
	// "dNSName: example.com, dNSName: www.example.com",
	San *string `json:"San,omitempty" xml:"San,omitempty"`
	// The serial number that indicates the priority of the returned address. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The certificate source.
	//
	// example:
	//
	// BYOC
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The validity status of the certificate. Valid values:
	//
	// - Valid
	//
	// - Invalid
	//
	// example:
	//
	// Valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The certificate subject (owner), identified in DN format.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:1825725063814405:eventstreaming/dsadsad123213-trigger1
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The DNS TLSA record value that stores the certificate public key fingerprint.
	//
	// example:
	//
	// 3 1 1 2ea103e8c5ba3466ff7f94cc28336b40ce7432e55a2fc37e86b65e55737c45662
	TlsaFingerprint *string `json:"TlsaFingerprint,omitempty" xml:"TlsaFingerprint,omitempty"`
	// The update time (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeAtiCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAtiCertificateResponseBody) GetAccessDeniedDetail() *DescribeAtiCertificateResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeAtiCertificateResponseBody) GetAgentHost() *string {
	return s.AgentHost
}

func (s *DescribeAtiCertificateResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribeAtiCertificateResponseBody) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DescribeAtiCertificateResponseBody) GetCertPem() *string {
	return s.CertPem
}

func (s *DescribeAtiCertificateResponseBody) GetCertType() *string {
	return s.CertType
}

func (s *DescribeAtiCertificateResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeAtiCertificateResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeAtiCertificateResponseBody) GetNotAfter() *string {
	return s.NotAfter
}

func (s *DescribeAtiCertificateResponseBody) GetNotBefore() *string {
	return s.NotBefore
}

func (s *DescribeAtiCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAtiCertificateResponseBody) GetSan() *string {
	return s.San
}

func (s *DescribeAtiCertificateResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeAtiCertificateResponseBody) GetSource() *string {
	return s.Source
}

func (s *DescribeAtiCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAtiCertificateResponseBody) GetSubject() *string {
	return s.Subject
}

func (s *DescribeAtiCertificateResponseBody) GetTlsaFingerprint() *string {
	return s.TlsaFingerprint
}

func (s *DescribeAtiCertificateResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeAtiCertificateResponseBody) SetAccessDeniedDetail(v *DescribeAtiCertificateResponseBodyAccessDeniedDetail) *DescribeAtiCertificateResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetAgentHost(v string) *DescribeAtiCertificateResponseBody {
	s.AgentHost = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetAgentId(v string) *DescribeAtiCertificateResponseBody {
	s.AgentId = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetAlgorithm(v string) *DescribeAtiCertificateResponseBody {
	s.Algorithm = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetCertPem(v string) *DescribeAtiCertificateResponseBody {
	s.CertPem = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetCertType(v string) *DescribeAtiCertificateResponseBody {
	s.CertType = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetCreateTimestamp(v int64) *DescribeAtiCertificateResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetIssuer(v string) *DescribeAtiCertificateResponseBody {
	s.Issuer = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetNotAfter(v string) *DescribeAtiCertificateResponseBody {
	s.NotAfter = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetNotBefore(v string) *DescribeAtiCertificateResponseBody {
	s.NotBefore = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetRequestId(v string) *DescribeAtiCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetSan(v string) *DescribeAtiCertificateResponseBody {
	s.San = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetSerialNumber(v string) *DescribeAtiCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetSource(v string) *DescribeAtiCertificateResponseBody {
	s.Source = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetStatus(v string) *DescribeAtiCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetSubject(v string) *DescribeAtiCertificateResponseBody {
	s.Subject = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetTlsaFingerprint(v string) *DescribeAtiCertificateResponseBody {
	s.TlsaFingerprint = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) SetUpdateTimestamp(v int64) *DescribeAtiCertificateResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeAtiCertificateResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAtiCertificateResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// CreateUser
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
	//
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authorization principal.
	//
	// example:
	//
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encrypted complete diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaNIARXXXXUQwNjE0LUQzN0XXXXVEQy1BQzExLTMzXXXXNTkxRjk1Ng==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason for the authentication failure. Valid values:
	//
	// - ExplicitDeny: Explicit deny.
	//
	// - ImplicitDeny: Implicit deny.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeAtiCertificateResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiCertificateResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeAtiCertificateResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeAtiCertificateResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeAtiCertificateResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeAtiCertificateResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeAtiCertificateResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeAtiCertificateResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeAtiCertificateResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeAtiCertificateResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
