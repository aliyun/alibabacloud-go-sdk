// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaCertificates(v []*UpdateListenerAttributeRequestCaCertificates) *UpdateListenerAttributeRequest
	GetCaCertificates() []*UpdateListenerAttributeRequestCaCertificates
	SetCaEnabled(v bool) *UpdateListenerAttributeRequest
	GetCaEnabled() *bool
	SetCertificates(v []*UpdateListenerAttributeRequestCertificates) *UpdateListenerAttributeRequest
	GetCertificates() []*UpdateListenerAttributeRequestCertificates
	SetClientToken(v string) *UpdateListenerAttributeRequest
	GetClientToken() *string
	SetDefaultActions(v []*UpdateListenerAttributeRequestDefaultActions) *UpdateListenerAttributeRequest
	GetDefaultActions() []*UpdateListenerAttributeRequestDefaultActions
	SetDryRun(v bool) *UpdateListenerAttributeRequest
	GetDryRun() *bool
	SetGzipEnabled(v bool) *UpdateListenerAttributeRequest
	GetGzipEnabled() *bool
	SetHttp2Enabled(v bool) *UpdateListenerAttributeRequest
	GetHttp2Enabled() *bool
	SetIdleTimeout(v int32) *UpdateListenerAttributeRequest
	GetIdleTimeout() *int32
	SetListenerDescription(v string) *UpdateListenerAttributeRequest
	GetListenerDescription() *string
	SetListenerId(v string) *UpdateListenerAttributeRequest
	GetListenerId() *string
	SetQuicConfig(v *UpdateListenerAttributeRequestQuicConfig) *UpdateListenerAttributeRequest
	GetQuicConfig() *UpdateListenerAttributeRequestQuicConfig
	SetRequestTimeout(v int32) *UpdateListenerAttributeRequest
	GetRequestTimeout() *int32
	SetSecurityPolicyId(v string) *UpdateListenerAttributeRequest
	GetSecurityPolicyId() *string
	SetXForwardedForConfig(v *UpdateListenerAttributeRequestXForwardedForConfig) *UpdateListenerAttributeRequest
	GetXForwardedForConfig() *UpdateListenerAttributeRequestXForwardedForConfig
}

type UpdateListenerAttributeRequest struct {
	// The CA certificate. You can specify only one CA certificate.
	CaCertificates []*UpdateListenerAttributeRequestCaCertificates `json:"CaCertificates,omitempty" xml:"CaCertificates,omitempty" type:"Repeated"`
	// Specifies whether to enable mutual authentication. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CaEnabled *bool `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// The certificates. You can add at most 20 certificates.
	//
	// >  Only server certificates are supported.
	Certificates []*UpdateListenerAttributeRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The actions of the default forwarding rule.
	DefaultActions []*UpdateListenerAttributeRequestDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable GZIP compression for specific types of files. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// Specifies whether to enable HTTP/2. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > This parameter is available only when you create an HTTPS listener.
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The timeout period for idle connections. Unit: seconds. Valid values: **1 to 60**
	//
	// If no requests are received within the specified timeout period, ALB closes the current connection. When another request is received, ALB establishes a new connection.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The name of the listener.
	//
	// The name must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// HTTP_80
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The ID of the Application Load Balancer (ALB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The configuration information when the listener is associated with a QUIC listener.
	QuicConfig *UpdateListenerAttributeRequestQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// The timeout period of a request. Unit: seconds. Valid values: **1 to 180**.
	//
	// If no response is received from the backend server within the specified timeout period, ALB returns an `HTTP 504` error code to the client.
	//
	// example:
	//
	// 3
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The security policy ID. System security policies and custom security policies are supported.
	//
	// > This parameter is available only when you create an HTTPS listener.
	//
	// example:
	//
	// tls_cipher_policy_1_0
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The configurations of the X-Forwarded-For header.
	XForwardedForConfig *UpdateListenerAttributeRequestXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s UpdateListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequest) GetCaCertificates() []*UpdateListenerAttributeRequestCaCertificates {
	return s.CaCertificates
}

func (s *UpdateListenerAttributeRequest) GetCaEnabled() *bool {
	return s.CaEnabled
}

func (s *UpdateListenerAttributeRequest) GetCertificates() []*UpdateListenerAttributeRequestCertificates {
	return s.Certificates
}

func (s *UpdateListenerAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateListenerAttributeRequest) GetDefaultActions() []*UpdateListenerAttributeRequestDefaultActions {
	return s.DefaultActions
}

func (s *UpdateListenerAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateListenerAttributeRequest) GetGzipEnabled() *bool {
	return s.GzipEnabled
}

func (s *UpdateListenerAttributeRequest) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *UpdateListenerAttributeRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *UpdateListenerAttributeRequest) GetListenerDescription() *string {
	return s.ListenerDescription
}

func (s *UpdateListenerAttributeRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateListenerAttributeRequest) GetQuicConfig() *UpdateListenerAttributeRequestQuicConfig {
	return s.QuicConfig
}

func (s *UpdateListenerAttributeRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *UpdateListenerAttributeRequest) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *UpdateListenerAttributeRequest) GetXForwardedForConfig() *UpdateListenerAttributeRequestXForwardedForConfig {
	return s.XForwardedForConfig
}

func (s *UpdateListenerAttributeRequest) SetCaCertificates(v []*UpdateListenerAttributeRequestCaCertificates) *UpdateListenerAttributeRequest {
	s.CaCertificates = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCaEnabled(v bool) *UpdateListenerAttributeRequest {
	s.CaEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCertificates(v []*UpdateListenerAttributeRequestCertificates) *UpdateListenerAttributeRequest {
	s.Certificates = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetClientToken(v string) *UpdateListenerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetDefaultActions(v []*UpdateListenerAttributeRequestDefaultActions) *UpdateListenerAttributeRequest {
	s.DefaultActions = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetDryRun(v bool) *UpdateListenerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetGzipEnabled(v bool) *UpdateListenerAttributeRequest {
	s.GzipEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetHttp2Enabled(v bool) *UpdateListenerAttributeRequest {
	s.Http2Enabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetIdleTimeout(v int32) *UpdateListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerDescription(v string) *UpdateListenerAttributeRequest {
	s.ListenerDescription = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerId(v string) *UpdateListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetQuicConfig(v *UpdateListenerAttributeRequestQuicConfig) *UpdateListenerAttributeRequest {
	s.QuicConfig = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetRequestTimeout(v int32) *UpdateListenerAttributeRequest {
	s.RequestTimeout = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetSecurityPolicyId(v string) *UpdateListenerAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetXForwardedForConfig(v *UpdateListenerAttributeRequestXForwardedForConfig) *UpdateListenerAttributeRequest {
	s.XForwardedForConfig = v
	return s
}

func (s *UpdateListenerAttributeRequest) Validate() error {
	if s.CaCertificates != nil {
		for _, item := range s.CaCertificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DefaultActions != nil {
		for _, item := range s.DefaultActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QuicConfig != nil {
		if err := s.QuicConfig.Validate(); err != nil {
			return err
		}
	}
	if s.XForwardedForConfig != nil {
		if err := s.XForwardedForConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateListenerAttributeRequestCaCertificates struct {
	// The ID of the CA certificate.
	//
	// >  This parameter is required if **CaEnabled*	- is set to **true**.
	//
	// example:
	//
	// 123359******
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s UpdateListenerAttributeRequestCaCertificates) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequestCaCertificates) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestCaCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *UpdateListenerAttributeRequestCaCertificates) SetCertificateId(v string) *UpdateListenerAttributeRequestCaCertificates {
	s.CertificateId = &v
	return s
}

func (s *UpdateListenerAttributeRequestCaCertificates) Validate() error {
	return dara.Validate(s)
}

type UpdateListenerAttributeRequestCertificates struct {
	// The certificate ID.
	//
	// example:
	//
	// 12315790212_166f8204689_1714763408_70998****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s UpdateListenerAttributeRequestCertificates) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequestCertificates) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *UpdateListenerAttributeRequestCertificates) SetCertificateId(v string) *UpdateListenerAttributeRequestCertificates {
	s.CertificateId = &v
	return s
}

func (s *UpdateListenerAttributeRequestCertificates) Validate() error {
	return dara.Validate(s)
}

type UpdateListenerAttributeRequestDefaultActions struct {
	// The forwarding action. This parameter takes effect only when you set **Type*	- to **ForwardGroup**. You can specify at most 20 actions.
	ForwardGroupConfig *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The action type. You can specify only one type.
	//
	// Set the value to **ForwardGroup**, which specifies that requests are forwarded to multiple server groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// ForwardGroup
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateListenerAttributeRequestDefaultActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequestDefaultActions) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestDefaultActions) GetForwardGroupConfig() *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *UpdateListenerAttributeRequestDefaultActions) GetType() *string {
	return s.Type
}

func (s *UpdateListenerAttributeRequestDefaultActions) SetForwardGroupConfig(v *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) *UpdateListenerAttributeRequestDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *UpdateListenerAttributeRequestDefaultActions) SetType(v string) *UpdateListenerAttributeRequestDefaultActions {
	s.Type = &v
	return s
}

func (s *UpdateListenerAttributeRequestDefaultActions) Validate() error {
	if s.ForwardGroupConfig != nil {
		if err := s.ForwardGroupConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig struct {
	// The server groups to which requests are forwarded.
	//
	// This parameter is required.
	ServerGroupTuples []*UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) GetServerGroupTuples() []*UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) Validate() error {
	if s.ServerGroupTuples != nil {
		for _, item := range s.ServerGroupTuples {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// The ID of the server group to which requests are forwarded.
	//
	// This parameter is required.
	//
	// example:
	//
	// rsp-cige6j5e7p****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type UpdateListenerAttributeRequestQuicConfig struct {
	// The QUIC listener ID. This parameter is required if **QuicUpgradeEnabled*	- is set to **true**. Only HTTPS listeners support this parameter.
	//
	// > You must add the HTTPS listener and the QUIC listener to the same ALB instance. In addition, make sure that the QUIC listener has never been associated with another listener.
	//
	// example:
	//
	// lsn-333
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// Specifies whether to enable QUIC upgrade. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// false
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s UpdateListenerAttributeRequestQuicConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequestQuicConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestQuicConfig) GetQuicListenerId() *string {
	return s.QuicListenerId
}

func (s *UpdateListenerAttributeRequestQuicConfig) GetQuicUpgradeEnabled() *bool {
	return s.QuicUpgradeEnabled
}

func (s *UpdateListenerAttributeRequestQuicConfig) SetQuicListenerId(v string) *UpdateListenerAttributeRequestQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequestQuicConfig) SetQuicUpgradeEnabled(v bool) *UpdateListenerAttributeRequestQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestQuicConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateListenerAttributeRequestXForwardedForConfig struct {
	// The name of the custom header. The header takes effect only when you set **XForwardedForClientCertClientVerifyEnabled **to **true**.
	//
	// The name must be 1 to 40 characters in length. It can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// >  This parameter is only available for HTTPS listeners.
	//
	// example:
	//
	// test_client-verify-alias_123456
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-clientverify` header to retrieve the verification result of the client certificate. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is only available for HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// The name of the custom header. The header takes effect only when you set **XForwardedForClientCertFingerprintEnabled*	- to **true**.
	//
	// The name must be 1 to 40 characters in length. It can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// >  This parameter is only available for HTTPS listeners.
	//
	// example:
	//
	// test_finger-print-alias_123456
	XForwardedForClientCertFingerprintAlias *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-fingerprint` header to retrieve the fingerprint of the client certificate. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is only available for HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// The name of the custom header. The header takes effect only when you set **XForwardedForClientCertIssuerDNEnabled*	- to **true**.
	//
	// The name must be 1 to 40 characters in length. It can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// >  This parameter is only available for HTTPS listeners.
	//
	// example:
	//
	// test_issue-dn-alias_123456
	XForwardedForClientCertIssuerDNAlias *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-issuerdn` header to retrieve information about the authority that issues the client certificate. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is only available for HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// The name of the custom header. This parameter is valid only if the **XForwardedForClientCertSubjectDNEnabled*	- parameter is set to **true**.
	//
	// The name must be 1 to 40 characters in length. It can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// >  This parameter is only available for HTTPS listeners.
	//
	// example:
	//
	// test_subject-dn-alias_123456
	XForwardedForClientCertSubjectDNAlias *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-subjectdn` header to retrieve information about the owner of the client certificate. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is only available for HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// Specifies whether to use the X-Forwarded-For header to preserve client IP addresses. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is only available for HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForClientSourceIpsEnabled *bool `json:"XForwardedForClientSourceIpsEnabled,omitempty" xml:"XForwardedForClientSourceIpsEnabled,omitempty"`
	// The trusted proxy IP address.
	//
	// ALB instances traverse the IP addresses in the `X-Forwarded-For` header from the rightmost IP address to the leftmost IP address. The first IP address that is not on the trusted IP address list is considered the client IP address. Requests from the client IP address are throttled.
	//
	// example:
	//
	// 10.1.1.0/24
	XForwardedForClientSourceIpsTrusted *string `json:"XForwardedForClientSourceIpsTrusted,omitempty" xml:"XForwardedForClientSourceIpsTrusted,omitempty"`
	// Specifies whether to use the `X-Forwarded-Client-srcport` header to retrieve the client port. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is only available for HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-For` header to retrieve the client IP address. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// > 	- If this parameter is set to **true**, the default value of the **XForwardedForProcessingMode*	- parameter is **append**. You can change it to **remove**.
	//
	// > 	- If this parameter is set to **false**, the `X-Forwarded-For` header in the request is not modified in any way before the request is sent to backend servers.
	//
	// > 	- This parameter is only available for HTTP and HTTPS listeners.
	//
	// example:
	//
	// true
	XForwardedForEnabled *bool `json:"XForwardedForEnabled,omitempty" xml:"XForwardedForEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-Host` header to retrieve the client domain name. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  This parameter is available for HTTP, HTTPS, and QUIC listeners.
	//
	// example:
	//
	// false
	XForwardedForHostEnabled *bool `json:"XForwardedForHostEnabled,omitempty" xml:"XForwardedForHostEnabled,omitempty"`
	// Specifies how the `X-Forwarded-For` header is processed. This parameter takes effect only when **XForwardedForEnabled*	- is set to **true**. Valid values:
	//
	// 	- **append*	- (default)
	//
	// 	- **remove**
	//
	// > 	- If this parameter is set to **append**, ALB appends the IP address of the last hop to the existing `X-Forwarded-For` header in the request before the request is sent to backend servers.
	//
	// > 	- If this parameter is set to **remove**, ALB removes the `X-Forwarded-For` header in the request before the request is sent to backend servers, no matter whether the request carries the `X-Forwarded-For` header.
	//
	// > 	- This parameter is only available for HTTP and HTTPS listeners.
	//
	// example:
	//
	// append
	XForwardedForProcessingMode *string `json:"XForwardedForProcessingMode,omitempty" xml:"XForwardedForProcessingMode,omitempty"`
	// Specifies whether to use the `X-Forwarded-Proto` header to retrieve the listener protocol. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is available for HTTP, HTTPS, and QUIC listeners.
	//
	// example:
	//
	// false
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// Specifies whether to use the `SLB-ID` header to retrieve the ID of the ALB instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is available for HTTP, HTTPS, and QUIC listeners.
	//
	// example:
	//
	// false
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-Port` header to retrieve the listener port of the ALB instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is available for HTTP, HTTPS, and QUIC listeners.
	//
	// example:
	//
	// false
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s UpdateListenerAttributeRequestXForwardedForConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequestXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientCertClientVerifyAlias() *string {
	return s.XForwardedForClientCertClientVerifyAlias
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientCertClientVerifyEnabled() *bool {
	return s.XForwardedForClientCertClientVerifyEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientCertFingerprintAlias() *string {
	return s.XForwardedForClientCertFingerprintAlias
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientCertFingerprintEnabled() *bool {
	return s.XForwardedForClientCertFingerprintEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientCertIssuerDNAlias() *string {
	return s.XForwardedForClientCertIssuerDNAlias
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientCertIssuerDNEnabled() *bool {
	return s.XForwardedForClientCertIssuerDNEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientCertSubjectDNAlias() *string {
	return s.XForwardedForClientCertSubjectDNAlias
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientCertSubjectDNEnabled() *bool {
	return s.XForwardedForClientCertSubjectDNEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientSourceIpsEnabled() *bool {
	return s.XForwardedForClientSourceIpsEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientSourceIpsTrusted() *string {
	return s.XForwardedForClientSourceIpsTrusted
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForClientSrcPortEnabled() *bool {
	return s.XForwardedForClientSrcPortEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForEnabled() *bool {
	return s.XForwardedForEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForHostEnabled() *bool {
	return s.XForwardedForHostEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForProcessingMode() *string {
	return s.XForwardedForProcessingMode
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForProtoEnabled() *bool {
	return s.XForwardedForProtoEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForSLBIdEnabled() *bool {
	return s.XForwardedForSLBIdEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) GetXForwardedForSLBPortEnabled() *bool {
	return s.XForwardedForSLBPortEnabled
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientSourceIpsEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientSourceIpsEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientSourceIpsTrusted(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientSourceIpsTrusted = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForHostEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForHostEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForProcessingMode(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForProcessingMode = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) Validate() error {
	return dara.Validate(s)
}
