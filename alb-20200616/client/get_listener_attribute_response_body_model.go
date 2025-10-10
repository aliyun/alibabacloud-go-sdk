// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclConfig(v *GetListenerAttributeResponseBodyAclConfig) *GetListenerAttributeResponseBody
	GetAclConfig() *GetListenerAttributeResponseBodyAclConfig
	SetCaCertificates(v []*GetListenerAttributeResponseBodyCaCertificates) *GetListenerAttributeResponseBody
	GetCaCertificates() []*GetListenerAttributeResponseBodyCaCertificates
	SetCaEnabled(v bool) *GetListenerAttributeResponseBody
	GetCaEnabled() *bool
	SetCertificates(v []*GetListenerAttributeResponseBodyCertificates) *GetListenerAttributeResponseBody
	GetCertificates() []*GetListenerAttributeResponseBodyCertificates
	SetDefaultActions(v []*GetListenerAttributeResponseBodyDefaultActions) *GetListenerAttributeResponseBody
	GetDefaultActions() []*GetListenerAttributeResponseBodyDefaultActions
	SetGzipEnabled(v bool) *GetListenerAttributeResponseBody
	GetGzipEnabled() *bool
	SetHttp2Enabled(v bool) *GetListenerAttributeResponseBody
	GetHttp2Enabled() *bool
	SetIdleTimeout(v int32) *GetListenerAttributeResponseBody
	GetIdleTimeout() *int32
	SetListenerDescription(v string) *GetListenerAttributeResponseBody
	GetListenerDescription() *string
	SetListenerId(v string) *GetListenerAttributeResponseBody
	GetListenerId() *string
	SetListenerPort(v int32) *GetListenerAttributeResponseBody
	GetListenerPort() *int32
	SetListenerProtocol(v string) *GetListenerAttributeResponseBody
	GetListenerProtocol() *string
	SetListenerStatus(v string) *GetListenerAttributeResponseBody
	GetListenerStatus() *string
	SetLoadBalancerId(v string) *GetListenerAttributeResponseBody
	GetLoadBalancerId() *string
	SetLogConfig(v *GetListenerAttributeResponseBodyLogConfig) *GetListenerAttributeResponseBody
	GetLogConfig() *GetListenerAttributeResponseBodyLogConfig
	SetQuicConfig(v *GetListenerAttributeResponseBodyQuicConfig) *GetListenerAttributeResponseBody
	GetQuicConfig() *GetListenerAttributeResponseBodyQuicConfig
	SetRequestId(v string) *GetListenerAttributeResponseBody
	GetRequestId() *string
	SetRequestTimeout(v int32) *GetListenerAttributeResponseBody
	GetRequestTimeout() *int32
	SetSecurityPolicyId(v string) *GetListenerAttributeResponseBody
	GetSecurityPolicyId() *string
	SetTags(v []*GetListenerAttributeResponseBodyTags) *GetListenerAttributeResponseBody
	GetTags() []*GetListenerAttributeResponseBodyTags
	SetXForwardedForConfig(v *GetListenerAttributeResponseBodyXForwardedForConfig) *GetListenerAttributeResponseBody
	GetXForwardedForConfig() *GetListenerAttributeResponseBodyXForwardedForConfig
}

type GetListenerAttributeResponseBody struct {
	// The configurations of the access control lists (ACLs).
	AclConfig *GetListenerAttributeResponseBodyAclConfig `json:"AclConfig,omitempty" xml:"AclConfig,omitempty" type:"Struct"`
	// A list of default CA certificates.
	CaCertificates []*GetListenerAttributeResponseBodyCaCertificates `json:"CaCertificates,omitempty" xml:"CaCertificates,omitempty" type:"Repeated"`
	// Indicates whether mutual authentication is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CaEnabled *bool `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// A list of certificates.
	Certificates []*GetListenerAttributeResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The actions of the default forwarding rule.
	DefaultActions []*GetListenerAttributeResponseBodyDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
	// Indicates whether GZIP compression is enabled to compress specific types of files. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// Indicates whether HTTP/2 is enabled. Valid values:
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
	// The timeout period of an idle connection. Unit: seconds.
	//
	// If no requests are received within the specified timeout period, Application Load Balancer (ALB) closes the current connection. When a request is received, ALB establishes a new connection.
	//
	// example:
	//
	// 2
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The name of the listener.
	//
	// example:
	//
	// test
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The frontend port that is used by the ALB instance.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol. Valid values: **HTTP**, **HTTPS**, and **QUIC**.
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The status of the listener. Valid values:
	//
	// 	- **Provisioning**
	//
	// 	- **Running**
	//
	// 	- **Configuring**
	//
	// 	- **Stopped**
	//
	// example:
	//
	// Running
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// The ALB instance ID.
	//
	// example:
	//
	// lb-bp1o94dp5i6ea****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The logging configuration.
	LogConfig *GetListenerAttributeResponseBodyLogConfig `json:"LogConfig,omitempty" xml:"LogConfig,omitempty" type:"Struct"`
	// The configuration information when the listener is associated with a QUIC listener.
	QuicConfig *GetListenerAttributeResponseBodyQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timeout period of a request. Unit: seconds.
	//
	// If no responses are received from the backend server within the specified timeout period, ALB returns an `HTTP 504` error code to the client.
	//
	// example:
	//
	// 34
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The security policy.
	//
	// > This parameter is available only when you create an HTTPS listener.
	//
	// example:
	//
	// tls_cipher_policy_1_1
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The tags.
	Tags []*GetListenerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The configuration of the XForward headers.
	XForwardedForConfig *GetListenerAttributeResponseBodyXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s GetListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBody) GetAclConfig() *GetListenerAttributeResponseBodyAclConfig {
	return s.AclConfig
}

func (s *GetListenerAttributeResponseBody) GetCaCertificates() []*GetListenerAttributeResponseBodyCaCertificates {
	return s.CaCertificates
}

func (s *GetListenerAttributeResponseBody) GetCaEnabled() *bool {
	return s.CaEnabled
}

func (s *GetListenerAttributeResponseBody) GetCertificates() []*GetListenerAttributeResponseBodyCertificates {
	return s.Certificates
}

func (s *GetListenerAttributeResponseBody) GetDefaultActions() []*GetListenerAttributeResponseBodyDefaultActions {
	return s.DefaultActions
}

func (s *GetListenerAttributeResponseBody) GetGzipEnabled() *bool {
	return s.GzipEnabled
}

func (s *GetListenerAttributeResponseBody) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *GetListenerAttributeResponseBody) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *GetListenerAttributeResponseBody) GetListenerDescription() *string {
	return s.ListenerDescription
}

func (s *GetListenerAttributeResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *GetListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *GetListenerAttributeResponseBody) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *GetListenerAttributeResponseBody) GetListenerStatus() *string {
	return s.ListenerStatus
}

func (s *GetListenerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *GetListenerAttributeResponseBody) GetLogConfig() *GetListenerAttributeResponseBodyLogConfig {
	return s.LogConfig
}

func (s *GetListenerAttributeResponseBody) GetQuicConfig() *GetListenerAttributeResponseBodyQuicConfig {
	return s.QuicConfig
}

func (s *GetListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetListenerAttributeResponseBody) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *GetListenerAttributeResponseBody) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *GetListenerAttributeResponseBody) GetTags() []*GetListenerAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetListenerAttributeResponseBody) GetXForwardedForConfig() *GetListenerAttributeResponseBodyXForwardedForConfig {
	return s.XForwardedForConfig
}

func (s *GetListenerAttributeResponseBody) SetAclConfig(v *GetListenerAttributeResponseBodyAclConfig) *GetListenerAttributeResponseBody {
	s.AclConfig = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCaCertificates(v []*GetListenerAttributeResponseBodyCaCertificates) *GetListenerAttributeResponseBody {
	s.CaCertificates = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCaEnabled(v bool) *GetListenerAttributeResponseBody {
	s.CaEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCertificates(v []*GetListenerAttributeResponseBodyCertificates) *GetListenerAttributeResponseBody {
	s.Certificates = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetDefaultActions(v []*GetListenerAttributeResponseBodyDefaultActions) *GetListenerAttributeResponseBody {
	s.DefaultActions = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetGzipEnabled(v bool) *GetListenerAttributeResponseBody {
	s.GzipEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetHttp2Enabled(v bool) *GetListenerAttributeResponseBody {
	s.Http2Enabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetIdleTimeout(v int32) *GetListenerAttributeResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerDescription(v string) *GetListenerAttributeResponseBody {
	s.ListenerDescription = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerId(v string) *GetListenerAttributeResponseBody {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerPort(v int32) *GetListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerProtocol(v string) *GetListenerAttributeResponseBody {
	s.ListenerProtocol = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerStatus(v string) *GetListenerAttributeResponseBody {
	s.ListenerStatus = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetLoadBalancerId(v string) *GetListenerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetLogConfig(v *GetListenerAttributeResponseBodyLogConfig) *GetListenerAttributeResponseBody {
	s.LogConfig = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetQuicConfig(v *GetListenerAttributeResponseBodyQuicConfig) *GetListenerAttributeResponseBody {
	s.QuicConfig = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRequestId(v string) *GetListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRequestTimeout(v int32) *GetListenerAttributeResponseBody {
	s.RequestTimeout = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetSecurityPolicyId(v string) *GetListenerAttributeResponseBody {
	s.SecurityPolicyId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetTags(v []*GetListenerAttributeResponseBodyTags) *GetListenerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetXForwardedForConfig(v *GetListenerAttributeResponseBodyXForwardedForConfig) *GetListenerAttributeResponseBody {
	s.XForwardedForConfig = v
	return s
}

func (s *GetListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyAclConfig struct {
	// The IDs of the ACLs that are associated with the listener.
	AclRelations []*GetListenerAttributeResponseBodyAclConfigAclRelations `json:"AclRelations,omitempty" xml:"AclRelations,omitempty" type:"Repeated"`
	// The type of the ACL. Valid values:
	//
	// 	- **White**: a whitelist. Only requests from the IP addresses or CIDR blocks in the network ACL are forwarded. Whitelists are applicable to scenarios in which you want to allow only specific IP addresses to access an application. Your service may be adversely affected if the whitelist is not properly configured. If a whitelist is configured for a listener, only requests from IP addresses that are on the whitelist are forwarded by the listener.
	//
	//     If you enable a whitelist but do not add an IP address to the whitelist, the listener forwards all requests.
	//
	// 	- **Black**: a blacklist. Requests from the IP addresses or CIDR blocks in the network ACL are denied. Blacklists are suitable for scenarios in which you want to deny access from specific IP addresses or CIDR blocks to an application.
	//
	//     If a blacklist is configured for a listener but no IP addresses are added to the blacklist, the listener forwards all requests.
	//
	// example:
	//
	// White
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
}

func (s GetListenerAttributeResponseBodyAclConfig) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyAclConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyAclConfig) GetAclRelations() []*GetListenerAttributeResponseBodyAclConfigAclRelations {
	return s.AclRelations
}

func (s *GetListenerAttributeResponseBodyAclConfig) GetAclType() *string {
	return s.AclType
}

func (s *GetListenerAttributeResponseBodyAclConfig) SetAclRelations(v []*GetListenerAttributeResponseBodyAclConfigAclRelations) *GetListenerAttributeResponseBodyAclConfig {
	s.AclRelations = v
	return s
}

func (s *GetListenerAttributeResponseBodyAclConfig) SetAclType(v string) *GetListenerAttributeResponseBodyAclConfig {
	s.AclType = &v
	return s
}

func (s *GetListenerAttributeResponseBodyAclConfig) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyAclConfigAclRelations struct {
	// The ID of the ACL that is associated with the listener.
	//
	// example:
	//
	// acl-doc****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Indicates whether the ACL is associated with the listener. Valid values:
	//
	// 	- **Associating**
	//
	// 	- **Associated**
	//
	// 	- **Dissociating**
	//
	// example:
	//
	// Associating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerAttributeResponseBodyAclConfigAclRelations) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyAclConfigAclRelations) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyAclConfigAclRelations) GetAclId() *string {
	return s.AclId
}

func (s *GetListenerAttributeResponseBodyAclConfigAclRelations) GetStatus() *string {
	return s.Status
}

func (s *GetListenerAttributeResponseBodyAclConfigAclRelations) SetAclId(v string) *GetListenerAttributeResponseBodyAclConfigAclRelations {
	s.AclId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyAclConfigAclRelations) SetStatus(v string) *GetListenerAttributeResponseBodyAclConfigAclRelations {
	s.Status = &v
	return s
}

func (s *GetListenerAttributeResponseBodyAclConfigAclRelations) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyCaCertificates struct {
	// The ID of the default CA certificate.
	//
	// example:
	//
	// 139a00604bd-cn-east-hangzho****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// Indicates whether the certificate is a default certificate: Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The status of the certificate.
	//
	// example:
	//
	// Associated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerAttributeResponseBodyCaCertificates) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyCaCertificates) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyCaCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *GetListenerAttributeResponseBodyCaCertificates) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetListenerAttributeResponseBodyCaCertificates) GetStatus() *string {
	return s.Status
}

func (s *GetListenerAttributeResponseBodyCaCertificates) SetCertificateId(v string) *GetListenerAttributeResponseBodyCaCertificates {
	s.CertificateId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyCaCertificates) SetIsDefault(v bool) *GetListenerAttributeResponseBodyCaCertificates {
	s.IsDefault = &v
	return s
}

func (s *GetListenerAttributeResponseBodyCaCertificates) SetStatus(v string) *GetListenerAttributeResponseBodyCaCertificates {
	s.Status = &v
	return s
}

func (s *GetListenerAttributeResponseBodyCaCertificates) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyCertificates struct {
	// The ID of the certificate. Only server certificates are supported.
	//
	// example:
	//
	// 12315790212_166f8204689_1714763408_70998****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s GetListenerAttributeResponseBodyCertificates) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *GetListenerAttributeResponseBodyCertificates) SetCertificateId(v string) *GetListenerAttributeResponseBodyCertificates {
	s.CertificateId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyCertificates) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyDefaultActions struct {
	// The configuration of the ForwardGroup action. This parameter is returned and takes effect when Type is set to **ForwardGroup**.
	ForwardGroupConfig *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The type of the action.
	//
	// If **ForwardGroup*	- is returned, requests are forwarded to multiple vServer groups.
	//
	// example:
	//
	// ForwardGroup
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetListenerAttributeResponseBodyDefaultActions) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyDefaultActions) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyDefaultActions) GetForwardGroupConfig() *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *GetListenerAttributeResponseBodyDefaultActions) GetType() *string {
	return s.Type
}

func (s *GetListenerAttributeResponseBodyDefaultActions) SetForwardGroupConfig(v *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) *GetListenerAttributeResponseBodyDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *GetListenerAttributeResponseBodyDefaultActions) SetType(v string) *GetListenerAttributeResponseBodyDefaultActions {
	s.Type = &v
	return s
}

func (s *GetListenerAttributeResponseBodyDefaultActions) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig struct {
	// The server group to which requests are forwarded.
	ServerGroupTuples []*GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) GetServerGroupTuples() []*GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// The ID of the server group to which requests are forwarded.
	//
	// example:
	//
	// rsp-cige6j****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyLogConfig struct {
	// Indicates whether custom headers are recorded in the access log. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// The configuration of Xtrace. Xtrace is used to record requests sent to ALB.
	AccessLogTracingConfig *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig `json:"AccessLogTracingConfig,omitempty" xml:"AccessLogTracingConfig,omitempty" type:"Struct"`
}

func (s GetListenerAttributeResponseBodyLogConfig) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyLogConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyLogConfig) GetAccessLogRecordCustomizedHeadersEnabled() *bool {
	return s.AccessLogRecordCustomizedHeadersEnabled
}

func (s *GetListenerAttributeResponseBodyLogConfig) GetAccessLogTracingConfig() *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig {
	return s.AccessLogTracingConfig
}

func (s *GetListenerAttributeResponseBodyLogConfig) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *GetListenerAttributeResponseBodyLogConfig {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfig) SetAccessLogTracingConfig(v *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) *GetListenerAttributeResponseBodyLogConfig {
	s.AccessLogTracingConfig = v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfig) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig struct {
	// Indicates whether Xtrace is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > You can set this parameter to **true*	- only if the AccessLogEnabled parameter is set to true.
	//
	// example:
	//
	// true
	TracingEnabled *bool `json:"TracingEnabled,omitempty" xml:"TracingEnabled,omitempty"`
	// The sampling rate of Xtrace. Valid values: 1 to 10000.
	//
	// > If **TracingEnabled*	- is set to **true**, this parameter is valid.
	//
	// example:
	//
	// 100
	TracingSample *int32 `json:"TracingSample,omitempty" xml:"TracingSample,omitempty"`
	// The Xtrace type. Supported Xtrace type: **Zipkin**.
	//
	// > If **TracingEnabled*	- is set to **true**, this parameter is valid.
	//
	// example:
	//
	// Zipkin
	TracingType *string `json:"TracingType,omitempty" xml:"TracingType,omitempty"`
}

func (s GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) GetTracingEnabled() *bool {
	return s.TracingEnabled
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) GetTracingSample() *int32 {
	return s.TracingSample
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) GetTracingType() *string {
	return s.TracingType
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) SetTracingEnabled(v bool) *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig {
	s.TracingEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) SetTracingSample(v int32) *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig {
	s.TracingSample = &v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) SetTracingType(v string) *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig {
	s.TracingType = &v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyQuicConfig struct {
	// The ID of the QUIC listener. This parameter is returned when **QuicUpgradeEnabled*	- is set to **true**. Only HTTPS listeners support this parameter.
	//
	// > You must associate the HTTPS listener and the QUIC listener with the same ALB instance. In addition, make sure that the QUIC listener has never been associated with another listener.
	//
	// example:
	//
	// lsn-333
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// Indicates whether QUIC upgrade is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s GetListenerAttributeResponseBodyQuicConfig) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyQuicConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyQuicConfig) GetQuicListenerId() *string {
	return s.QuicListenerId
}

func (s *GetListenerAttributeResponseBodyQuicConfig) GetQuicUpgradeEnabled() *bool {
	return s.QuicUpgradeEnabled
}

func (s *GetListenerAttributeResponseBodyQuicConfig) SetQuicListenerId(v string) *GetListenerAttributeResponseBodyQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyQuicConfig) SetQuicUpgradeEnabled(v bool) *GetListenerAttributeResponseBodyQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyQuicConfig) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyTags struct {
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

func (s GetListenerAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetListenerAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetListenerAttributeResponseBodyTags) SetKey(v string) *GetListenerAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetListenerAttributeResponseBodyTags) SetValue(v string) *GetListenerAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetListenerAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type GetListenerAttributeResponseBodyXForwardedForConfig struct {
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertClientVerifyEnabled*	- is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	//
	// example:
	//
	// test_client-verify-alias_123456
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-clientverify` header is used to retrieve the verification result of the client certificate. Valid values:
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
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertFingerprintEnabled*	- is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	//
	// example:
	//
	// test_finger-print-alias_123456
	XForwardedForClientCertFingerprintAlias *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-fingerprint` header is used to retrieve the fingerprint of the client certificate. Valid values:
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
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertIssuerDNEnabled*	- is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	//
	// example:
	//
	// test_issue-dn-alias_123456
	XForwardedForClientCertIssuerDNAlias *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-issuerdn` header is used to retrieve information about the authority that issues the client certificate. Valid values:
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
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertSubjectDNEnabled*	- is set to **true**.
	//
	// The name is 1 to 40 characters in length, and can contain lowercase letters, hyphens (-), underscores (_), and digits.
	//
	// > This parameter is available only when you create an HTTPS listener.
	//
	// example:
	//
	// test_subject-dn-alias_123456
	XForwardedForClientCertSubjectDNAlias *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-subjectdn` header is used to retrieve information about the owner of the client certificate. Valid values:
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
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// Indicates whether the `X-Forwarded-Client-Ip` header is used to retrieve the source port of the ALB instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > This parameter is available only when you create an HTTP, HTTPS, or QUIC listener.
	//
	// example:
	//
	// false
	XForwardedForClientSourceIpsEnabled *bool `json:"XForwardedForClientSourceIpsEnabled,omitempty" xml:"XForwardedForClientSourceIpsEnabled,omitempty"`
	// The trusted proxy IP address.
	//
	// ALB traverses `X-Forwarded-For` backward and selects the first IP address that is not on the trusted IP address list as the real IP address of the client. The IP address is used in source IP address throttling.
	//
	// example:
	//
	// 10.1.1.0/24
	XForwardedForClientSourceIpsTrusted *string `json:"XForwardedForClientSourceIpsTrusted,omitempty" xml:"XForwardedForClientSourceIpsTrusted,omitempty"`
	// Indicates whether the `X-Forwarded-Client-Port` header is used to retrieve the client port. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > This parameter is available only when you create an HTTP or HTTPS listener.
	//
	// example:
	//
	// true
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	// Indicates whether the `X-Forwarded-For` header is used to retrieve the client IP address. Valid values:
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
	// Indicates whether the `X-Forwarded-Proto` header is used to retrieve the listening protocol. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > This parameter is available only when you create an HTTP, HTTPS, or QUIC listener.
	//
	// example:
	//
	// true
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// Indicates whether the `SLB-ID` header is used to retrieve the ID of the CLB instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > This parameter is available only when you create an HTTP, HTTPS, or QUIC listener.
	//
	// example:
	//
	// true
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// Indicates whether the `X-Forwarded-Port` header is used to retrieve the listening port of the ALB instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > This parameter is available only when you create an HTTP, HTTPS, or QUIC listener.
	//
	// example:
	//
	// true
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s GetListenerAttributeResponseBodyXForwardedForConfig) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeResponseBodyXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientCertClientVerifyAlias() *string {
	return s.XForwardedForClientCertClientVerifyAlias
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientCertClientVerifyEnabled() *bool {
	return s.XForwardedForClientCertClientVerifyEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientCertFingerprintAlias() *string {
	return s.XForwardedForClientCertFingerprintAlias
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientCertFingerprintEnabled() *bool {
	return s.XForwardedForClientCertFingerprintEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientCertIssuerDNAlias() *string {
	return s.XForwardedForClientCertIssuerDNAlias
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientCertIssuerDNEnabled() *bool {
	return s.XForwardedForClientCertIssuerDNEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientCertSubjectDNAlias() *string {
	return s.XForwardedForClientCertSubjectDNAlias
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientCertSubjectDNEnabled() *bool {
	return s.XForwardedForClientCertSubjectDNEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientSourceIpsEnabled() *bool {
	return s.XForwardedForClientSourceIpsEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientSourceIpsTrusted() *string {
	return s.XForwardedForClientSourceIpsTrusted
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForClientSrcPortEnabled() *bool {
	return s.XForwardedForClientSrcPortEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForEnabled() *bool {
	return s.XForwardedForEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForHostEnabled() *bool {
	return s.XForwardedForHostEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForProcessingMode() *string {
	return s.XForwardedForProcessingMode
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForProtoEnabled() *bool {
	return s.XForwardedForProtoEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForSLBIdEnabled() *bool {
	return s.XForwardedForSLBIdEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) GetXForwardedForSLBPortEnabled() *bool {
	return s.XForwardedForSLBPortEnabled
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientSourceIpsEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientSourceIpsEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientSourceIpsTrusted(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientSourceIpsTrusted = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForHostEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForHostEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForProcessingMode(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForProcessingMode = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) Validate() error {
	return dara.Validate(s)
}
