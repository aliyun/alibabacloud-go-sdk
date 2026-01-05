// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaCertificates(v []*CreateListenerRequestCaCertificates) *CreateListenerRequest
	GetCaCertificates() []*CreateListenerRequestCaCertificates
	SetCaEnabled(v bool) *CreateListenerRequest
	GetCaEnabled() *bool
	SetCertificates(v []*CreateListenerRequestCertificates) *CreateListenerRequest
	GetCertificates() []*CreateListenerRequestCertificates
	SetClientToken(v string) *CreateListenerRequest
	GetClientToken() *string
	SetDefaultActions(v []*CreateListenerRequestDefaultActions) *CreateListenerRequest
	GetDefaultActions() []*CreateListenerRequestDefaultActions
	SetDryRun(v bool) *CreateListenerRequest
	GetDryRun() *bool
	SetGzipEnabled(v bool) *CreateListenerRequest
	GetGzipEnabled() *bool
	SetHttp2Enabled(v bool) *CreateListenerRequest
	GetHttp2Enabled() *bool
	SetIdleTimeout(v int32) *CreateListenerRequest
	GetIdleTimeout() *int32
	SetListenerDescription(v string) *CreateListenerRequest
	GetListenerDescription() *string
	SetListenerPort(v int32) *CreateListenerRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *CreateListenerRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *CreateListenerRequest
	GetLoadBalancerId() *string
	SetQuicConfig(v *CreateListenerRequestQuicConfig) *CreateListenerRequest
	GetQuicConfig() *CreateListenerRequestQuicConfig
	SetRequestTimeout(v int32) *CreateListenerRequest
	GetRequestTimeout() *int32
	SetSecurityPolicyId(v string) *CreateListenerRequest
	GetSecurityPolicyId() *string
	SetTag(v []*CreateListenerRequestTag) *CreateListenerRequest
	GetTag() []*CreateListenerRequestTag
	SetXForwardedForConfig(v *CreateListenerRequestXForwardedForConfig) *CreateListenerRequest
	GetXForwardedForConfig() *CreateListenerRequestXForwardedForConfig
}

type CreateListenerRequest struct {
	// The certificate authority (CA) certificates. You can specify only one CA certificate.
	CaCertificates []*CreateListenerRequestCaCertificates `json:"CaCertificates,omitempty" xml:"CaCertificates,omitempty" type:"Repeated"`
	// Specifies whether to enable mutual authentication. Valid values:
	//
	// 	- **true**: enables mutual authentication.
	//
	// 	- **false*	- (default): disables mutual authentication.
	//
	// example:
	//
	// false
	CaEnabled *bool `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// The details about each certificate.
	Certificates []*CreateListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses the value of **RequestId*	- as the value of **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The actions of the forwarding rule.
	//
	// This parameter is required.
	DefaultActions []*CreateListenerRequestDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
	// Specifies whether to perform only a precheck. Valid values:
	//
	// 	- **true**: prechecks the request without creating a listener. The system checks the required parameters, request syntax, and limits. If the request fails the precheck, an error code is returned based on the cause of the failure. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the API request. If the request passes the precheck, a 2xx HTTP status code is returned and the system proceeds to create a listener.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable `Gzip` compression to compress specific types of files. Valid values:
	//
	// 	- **true*	- (default): enables Gzip compression.
	//
	// 	- **false**: disables Gzip compression.
	//
	// example:
	//
	// true
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// Specifies whether to enable `HTTP/2`. Valid values:
	//
	// 	- **true*	- (default): enables HTTP/2.
	//
	// 	- **false**: disables HTTP/2.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The timeout period of an idle connection. Unit: seconds.
	//
	// Valid values: **1 to 60**.
	//
	// Default value: **15**.
	//
	// If no requests are received within the specified timeout period, ALB closes the current connection. When a new request is received, ALB establishes a new connection.
	//
	// example:
	//
	// 3
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The name of the listener.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_). Regular expressions are supported.
	//
	// example:
	//
	// HTTP_80
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The frontend port that is used by the ALB instance.
	//
	// Valid values: **1 to 65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol.
	//
	// Valid values: **HTTP**, **HTTPS**, and **QUIC**.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the ALB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-n5qw04uq8vavfe****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// Select a QUIC listener and associate it with the ALB instance.
	QuicConfig *CreateListenerRequestQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// The timeout period of a request. Unit: seconds.
	//
	// Valid values: **1 to 180**.
	//
	// Default value: **60**.
	//
	// If no response is received from the backend server during the request timeout period, ALB sends an `HTTP 504` error code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The ID of the security policy. System security policies and custom security policies are supported.
	//
	// Default value: **tls_cipher_policy_1_0*	- (system security policy).
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// tls_cipher_policy_1_0
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The tags.
	Tag []*CreateListenerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The configuration of the XForward header.
	XForwardedForConfig *CreateListenerRequestXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s CreateListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerRequest) GetCaCertificates() []*CreateListenerRequestCaCertificates {
	return s.CaCertificates
}

func (s *CreateListenerRequest) GetCaEnabled() *bool {
	return s.CaEnabled
}

func (s *CreateListenerRequest) GetCertificates() []*CreateListenerRequestCertificates {
	return s.Certificates
}

func (s *CreateListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateListenerRequest) GetDefaultActions() []*CreateListenerRequestDefaultActions {
	return s.DefaultActions
}

func (s *CreateListenerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateListenerRequest) GetGzipEnabled() *bool {
	return s.GzipEnabled
}

func (s *CreateListenerRequest) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *CreateListenerRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *CreateListenerRequest) GetListenerDescription() *string {
	return s.ListenerDescription
}

func (s *CreateListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateListenerRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *CreateListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateListenerRequest) GetQuicConfig() *CreateListenerRequestQuicConfig {
	return s.QuicConfig
}

func (s *CreateListenerRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *CreateListenerRequest) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *CreateListenerRequest) GetTag() []*CreateListenerRequestTag {
	return s.Tag
}

func (s *CreateListenerRequest) GetXForwardedForConfig() *CreateListenerRequestXForwardedForConfig {
	return s.XForwardedForConfig
}

func (s *CreateListenerRequest) SetCaCertificates(v []*CreateListenerRequestCaCertificates) *CreateListenerRequest {
	s.CaCertificates = v
	return s
}

func (s *CreateListenerRequest) SetCaEnabled(v bool) *CreateListenerRequest {
	s.CaEnabled = &v
	return s
}

func (s *CreateListenerRequest) SetCertificates(v []*CreateListenerRequestCertificates) *CreateListenerRequest {
	s.Certificates = v
	return s
}

func (s *CreateListenerRequest) SetClientToken(v string) *CreateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateListenerRequest) SetDefaultActions(v []*CreateListenerRequestDefaultActions) *CreateListenerRequest {
	s.DefaultActions = v
	return s
}

func (s *CreateListenerRequest) SetDryRun(v bool) *CreateListenerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateListenerRequest) SetGzipEnabled(v bool) *CreateListenerRequest {
	s.GzipEnabled = &v
	return s
}

func (s *CreateListenerRequest) SetHttp2Enabled(v bool) *CreateListenerRequest {
	s.Http2Enabled = &v
	return s
}

func (s *CreateListenerRequest) SetIdleTimeout(v int32) *CreateListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateListenerRequest) SetListenerDescription(v string) *CreateListenerRequest {
	s.ListenerDescription = &v
	return s
}

func (s *CreateListenerRequest) SetListenerPort(v int32) *CreateListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateListenerRequest) SetListenerProtocol(v string) *CreateListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *CreateListenerRequest) SetLoadBalancerId(v string) *CreateListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateListenerRequest) SetQuicConfig(v *CreateListenerRequestQuicConfig) *CreateListenerRequest {
	s.QuicConfig = v
	return s
}

func (s *CreateListenerRequest) SetRequestTimeout(v int32) *CreateListenerRequest {
	s.RequestTimeout = &v
	return s
}

func (s *CreateListenerRequest) SetSecurityPolicyId(v string) *CreateListenerRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *CreateListenerRequest) SetTag(v []*CreateListenerRequestTag) *CreateListenerRequest {
	s.Tag = v
	return s
}

func (s *CreateListenerRequest) SetXForwardedForConfig(v *CreateListenerRequestXForwardedForConfig) *CreateListenerRequest {
	s.XForwardedForConfig = v
	return s
}

func (s *CreateListenerRequest) Validate() error {
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.XForwardedForConfig != nil {
		if err := s.XForwardedForConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateListenerRequestCaCertificates struct {
	// The ID of the CA certificate.
	//
	// >  This parameter is required if **CaEnabled*	- is set to **true**.
	//
	// example:
	//
	// 123157*******
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s CreateListenerRequestCaCertificates) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestCaCertificates) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCaCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CreateListenerRequestCaCertificates) SetCertificateId(v string) *CreateListenerRequestCaCertificates {
	s.CertificateId = &v
	return s
}

func (s *CreateListenerRequestCaCertificates) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestCertificates struct {
	// The ID of the certificate. Only server certificates are supported. You can specify up to 20 certificate IDs.
	//
	// example:
	//
	// 103705*******
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s CreateListenerRequestCertificates) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CreateListenerRequestCertificates) SetCertificateId(v string) *CreateListenerRequestCertificates {
	s.CertificateId = &v
	return s
}

func (s *CreateListenerRequestCertificates) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestDefaultActions struct {
	// The configuration of the forwarding action. You can specify at most 20 actions.
	//
	// This parameter is required.
	ForwardGroupConfig *CreateListenerRequestDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The action type. You can specify only one action type. Valid value:
	//
	// **ForwardGroup**: forwards requests to multiple Server groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// ForwardGroup
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateListenerRequestDefaultActions) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestDefaultActions) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestDefaultActions) GetForwardGroupConfig() *CreateListenerRequestDefaultActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *CreateListenerRequestDefaultActions) GetType() *string {
	return s.Type
}

func (s *CreateListenerRequestDefaultActions) SetForwardGroupConfig(v *CreateListenerRequestDefaultActionsForwardGroupConfig) *CreateListenerRequestDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateListenerRequestDefaultActions) SetType(v string) *CreateListenerRequestDefaultActions {
	s.Type = &v
	return s
}

func (s *CreateListenerRequestDefaultActions) Validate() error {
	if s.ForwardGroupConfig != nil {
		if err := s.ForwardGroupConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateListenerRequestDefaultActionsForwardGroupConfig struct {
	// The destination server group to which requests are forwarded.
	//
	// This parameter is required.
	ServerGroupTuples []*CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestDefaultActionsForwardGroupConfig) GetServerGroupTuples() []*CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *CreateListenerRequestDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) *CreateListenerRequestDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *CreateListenerRequestDefaultActionsForwardGroupConfig) Validate() error {
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

type CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// The ID of the server group to which requests are forwarded.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-8ilqs4axp6******
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestQuicConfig struct {
	// The ID of the QUIC listener that you want to associate with the HTTPS listener. Only HTTPS listeners support this parameter. This parameter is required when **QuicUpgradeEnabled*	- is set to **true**.
	//
	// >  The HTTPS listener and the QUIC listener must be added to the same ALB instance. Make sure that the QUIC listener is not associated with any other listeners.
	//
	// example:
	//
	// lsn-o4u54y73wq7b******
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// Specifies whether to enable QUIC upgrade. Valid values:
	//
	// 	- **true**: enables QUIC upgrade.
	//
	// 	- **false*	- (default): disables QUIC upgrade.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// false
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s CreateListenerRequestQuicConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestQuicConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestQuicConfig) GetQuicListenerId() *string {
	return s.QuicListenerId
}

func (s *CreateListenerRequestQuicConfig) GetQuicUpgradeEnabled() *bool {
	return s.QuicUpgradeEnabled
}

func (s *CreateListenerRequestQuicConfig) SetQuicListenerId(v string) *CreateListenerRequestQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *CreateListenerRequestQuicConfig) SetQuicUpgradeEnabled(v bool) *CreateListenerRequestQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

func (s *CreateListenerRequestQuicConfig) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateListenerRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateListenerRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateListenerRequestTag) SetKey(v string) *CreateListenerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateListenerRequestTag) SetValue(v string) *CreateListenerRequestTag {
	s.Value = &v
	return s
}

func (s *CreateListenerRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestXForwardedForConfig struct {
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertClientVerifyEnabled*	- is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (_), and digits.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// test_client-verify-alias_123456
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-clientverify` header to retrieve the verification result of the client certificate. Valid values:
	//
	// 	- **true**: uses the X-Forwarded-Clientcert-clientverify header.
	//
	// 	- **false*	- (default): does not use the X-Forwarded-Clientcert-clientverify header.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertFingerprintEnabled*	- is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (_), and digits.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// test_finger-print-alias_123456
	XForwardedForClientCertFingerprintAlias *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-fingerprint` header to retrieve the fingerprint of the client certificate. Valid values:
	//
	// 	- **true**: uses the X-Forwarded-Clientcert-fingerprint header.
	//
	// 	- **false*	- (default): does not use the X-Forwarded-Clientcert-fingerprint header.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertIssuerDNEnabled*	- is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (_), and digits.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// test_issue-dn-alias_123456
	XForwardedForClientCertIssuerDNAlias *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-issuerdn` header to retrieve information about the authority that issues the client certificate. Valid values:
	//
	// 	- **true**: uses the X-Forwarded-Clientcert-issuerdn header.
	//
	// 	- **false*	- (default): does not use the X-Forwarded-Clientcert-issuerdn header.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertSubjectDNEnabled*	- is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (_), and digits.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// test_subject-dn-alias_123456
	XForwardedForClientCertSubjectDNAlias *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	// Specifies whether to use the `X-Forwarded-Clientcert-subjectdn` header to retrieve information about the owner of the client certificate. Valid values:
	//
	// 	- **true**: uses the X-Forwarded-Clientcert-subjectdn header.
	//
	// 	- **false*	- (default): does not use the X-Forwarded-Clientcert-subjectdn header.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// Specifies whether to allow the ALB instance to retrieve client IP addresses from the `X-Forwarded-For` header. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  HTTP and HTTPS listeners support this parameter.
	//
	// example:
	//
	// false
	XForwardedForClientSourceIpsEnabled *bool `json:"XForwardedForClientSourceIpsEnabled,omitempty" xml:"XForwardedForClientSourceIpsEnabled,omitempty"`
	// The trusted proxy IP address.
	//
	// ALB traverses `X-Forwarded-For` backwards and selects the first IP address that is not in the trusted IP list as the originating IP address of the client, which will be throttled if source IP address throttling is enabled.
	//
	// example:
	//
	// 10.1.1.0/24
	XForwardedForClientSourceIpsTrusted *string `json:"XForwardedForClientSourceIpsTrusted,omitempty" xml:"XForwardedForClientSourceIpsTrusted,omitempty"`
	// Specifies whether to use the `X-Forwarded-Client-srcport` header to retrieve the client port. Valid values:
	//
	// 	- **true**: uses the X-Forwarded-Client-srcport header.
	//
	// 	- **false*	- (default): does not use the X-Forwarded-Client-srcport header.
	//
	// >  HTTP and HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-For` header to retrieve client IP addresses. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// >  HTTP and HTTPS listeners support this parameter.
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
	// Specifies whether to use the `X-Forwarded-Proto` header to retrieve the listener protocol of the ALB instance. Valid values:
	//
	// 	- **true**: uses the X-Forwarded-Proto header.
	//
	// 	- **false*	- (default): does not use the X-Forwarded-Proto header.
	//
	// >  HTTP, HTTPS, and QUIC listeners support this parameter.
	//
	// example:
	//
	// false
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// Specifies whether to use the `SLB-ID` header to retrieve the ID of the ALB instance. Valid values:
	//
	// 	- **true**: uses the SLB-ID header.
	//
	// 	- **false*	- (default): does not use the SLB-ID header.
	//
	// >  HTTP, HTTPS, and QUIC listeners support this parameter.
	//
	// example:
	//
	// false
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-Port` header to retrieve the listener port of the ALB instance. Valid values:
	//
	// 	- **true**: uses the X-Forwarded-Port header.
	//
	// 	- **false*	- (default): does not use the X-Forwarded-Port header.
	//
	// >  HTTP, HTTPS, and QUIC listeners support this parameter.
	//
	// example:
	//
	// false
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s CreateListenerRequestXForwardedForConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientCertClientVerifyAlias() *string {
	return s.XForwardedForClientCertClientVerifyAlias
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientCertClientVerifyEnabled() *bool {
	return s.XForwardedForClientCertClientVerifyEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientCertFingerprintAlias() *string {
	return s.XForwardedForClientCertFingerprintAlias
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientCertFingerprintEnabled() *bool {
	return s.XForwardedForClientCertFingerprintEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientCertIssuerDNAlias() *string {
	return s.XForwardedForClientCertIssuerDNAlias
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientCertIssuerDNEnabled() *bool {
	return s.XForwardedForClientCertIssuerDNEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientCertSubjectDNAlias() *string {
	return s.XForwardedForClientCertSubjectDNAlias
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientCertSubjectDNEnabled() *bool {
	return s.XForwardedForClientCertSubjectDNEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientSourceIpsEnabled() *bool {
	return s.XForwardedForClientSourceIpsEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientSourceIpsTrusted() *string {
	return s.XForwardedForClientSourceIpsTrusted
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForClientSrcPortEnabled() *bool {
	return s.XForwardedForClientSrcPortEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForEnabled() *bool {
	return s.XForwardedForEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForHostEnabled() *bool {
	return s.XForwardedForHostEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForProcessingMode() *string {
	return s.XForwardedForProcessingMode
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForProtoEnabled() *bool {
	return s.XForwardedForProtoEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForSLBIdEnabled() *bool {
	return s.XForwardedForSLBIdEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForSLBPortEnabled() *bool {
	return s.XForwardedForSLBPortEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientSourceIpsEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientSourceIpsEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientSourceIpsTrusted(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientSourceIpsTrusted = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForHostEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForHostEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForProcessingMode(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForProcessingMode = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) Validate() error {
	return dara.Validate(s)
}
