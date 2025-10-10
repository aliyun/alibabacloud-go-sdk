// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody
	GetListeners() []*ListListenersResponseBodyListeners
	SetMaxResults(v int32) *ListListenersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListListenersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListListenersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListListenersResponseBody
	GetTotalCount() *int32
}

type ListListenersResponseBody struct {
	// The listeners.
	Listeners []*ListListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position where the query stopped. If this parameter is not returned, all data is queried.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBody) GetListeners() []*ListListenersResponseBodyListeners {
	return s.Listeners
}

func (s *ListListenersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListListenersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListListenersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListListenersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListListenersResponseBody) SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersResponseBody) SetMaxResults(v int32) *ListListenersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenersResponseBody) SetNextToken(v string) *ListListenersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenersResponseBody) SetRequestId(v string) *ListListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersResponseBody) SetTotalCount(v int32) *ListListenersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListListenersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListeners struct {
	// The default actions in the forwarding rules.
	DefaultActions []*ListListenersResponseBodyListenersDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
	// Indicates whether GZIP compression is enabled to compress specific types of files. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// Indicates whether HTTP/2 is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// false
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The timeout period of an idle connection. Unit: seconds. Valid values: **1 to 60**.
	//
	// If no request is received within the specified timeout period, ALB closes the connection. ALB establishes the connection again when a new connection request is received.
	//
	// example:
	//
	// 3
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The name of the listener.
	//
	// example:
	//
	// test
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listener ID.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The frontend port that is used by the ALB instance. Valid values: **1 to 65535**.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol of the instance. Valid values:
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// 	- **QUIC**
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The status of the listener. Valid values:
	//
	// 	- **Provisioning**: The listener is being created.
	//
	// 	- **Running**: The listener is running.
	//
	// 	- **Configuring**: The listener is being configured.
	//
	// 	- **Stopped**: The listener is disabled.
	//
	// example:
	//
	// Running
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// The ALB instance ID.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The logging configurations.
	LogConfig *ListListenersResponseBodyListenersLogConfig `json:"LogConfig,omitempty" xml:"LogConfig,omitempty" type:"Struct"`
	// The configurations of the QUIC listener associated with the ALB instance.
	QuicConfig *ListListenersResponseBodyListenersQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// The timeout period of a request. Unit: seconds. Valid values: **1 to 180**.
	//
	// If no responses are received from the backend server within the specified timeout period, ALB returns an `HTTP 504` error code to the client.
	//
	// example:
	//
	// 34
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The security policy.
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// tls_cipher_policy_1_1
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The tags.
	Tags []*ListListenersResponseBodyListenersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The configuration of the `XForward` header.
	XForwardedForConfig *ListListenersResponseBodyListenersXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s ListListenersResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListeners) GetDefaultActions() []*ListListenersResponseBodyListenersDefaultActions {
	return s.DefaultActions
}

func (s *ListListenersResponseBodyListeners) GetGzipEnabled() *bool {
	return s.GzipEnabled
}

func (s *ListListenersResponseBodyListeners) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *ListListenersResponseBodyListeners) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *ListListenersResponseBodyListeners) GetListenerDescription() *string {
	return s.ListenerDescription
}

func (s *ListListenersResponseBodyListeners) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListListenersResponseBodyListeners) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *ListListenersResponseBodyListeners) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *ListListenersResponseBodyListeners) GetListenerStatus() *string {
	return s.ListenerStatus
}

func (s *ListListenersResponseBodyListeners) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListListenersResponseBodyListeners) GetLogConfig() *ListListenersResponseBodyListenersLogConfig {
	return s.LogConfig
}

func (s *ListListenersResponseBodyListeners) GetQuicConfig() *ListListenersResponseBodyListenersQuicConfig {
	return s.QuicConfig
}

func (s *ListListenersResponseBodyListeners) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *ListListenersResponseBodyListeners) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *ListListenersResponseBodyListeners) GetTags() []*ListListenersResponseBodyListenersTags {
	return s.Tags
}

func (s *ListListenersResponseBodyListeners) GetXForwardedForConfig() *ListListenersResponseBodyListenersXForwardedForConfig {
	return s.XForwardedForConfig
}

func (s *ListListenersResponseBodyListeners) SetDefaultActions(v []*ListListenersResponseBodyListenersDefaultActions) *ListListenersResponseBodyListeners {
	s.DefaultActions = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetGzipEnabled(v bool) *ListListenersResponseBodyListeners {
	s.GzipEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetHttp2Enabled(v bool) *ListListenersResponseBodyListeners {
	s.Http2Enabled = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetIdleTimeout(v int32) *ListListenersResponseBodyListeners {
	s.IdleTimeout = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerDescription(v string) *ListListenersResponseBodyListeners {
	s.ListenerDescription = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerId(v string) *ListListenersResponseBodyListeners {
	s.ListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerPort(v int32) *ListListenersResponseBodyListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerProtocol(v string) *ListListenersResponseBodyListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerStatus(v string) *ListListenersResponseBodyListeners {
	s.ListenerStatus = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetLoadBalancerId(v string) *ListListenersResponseBodyListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetLogConfig(v *ListListenersResponseBodyListenersLogConfig) *ListListenersResponseBodyListeners {
	s.LogConfig = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetQuicConfig(v *ListListenersResponseBodyListenersQuicConfig) *ListListenersResponseBodyListeners {
	s.QuicConfig = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetRequestTimeout(v int32) *ListListenersResponseBodyListeners {
	s.RequestTimeout = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetSecurityPolicyId(v string) *ListListenersResponseBodyListeners {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetTags(v []*ListListenersResponseBodyListenersTags) *ListListenersResponseBodyListeners {
	s.Tags = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetXForwardedForConfig(v *ListListenersResponseBodyListenersXForwardedForConfig) *ListListenersResponseBodyListeners {
	s.XForwardedForConfig = v
	return s
}

func (s *ListListenersResponseBodyListeners) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersDefaultActions struct {
	// The configuration of the forwarding rule action. This parameter takes effect only when the action is **ForwardGroup**.
	ForwardGroupConfig *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The action. **ForwardGroup**: forwards requests to multiple server groups.
	//
	// example:
	//
	// ForwardGroup
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListListenersResponseBodyListenersDefaultActions) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersDefaultActions) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersDefaultActions) GetForwardGroupConfig() *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *ListListenersResponseBodyListenersDefaultActions) GetType() *string {
	return s.Type
}

func (s *ListListenersResponseBodyListenersDefaultActions) SetForwardGroupConfig(v *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) *ListListenersResponseBodyListenersDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *ListListenersResponseBodyListenersDefaultActions) SetType(v string) *ListListenersResponseBodyListenersDefaultActions {
	s.Type = &v
	return s
}

func (s *ListListenersResponseBodyListenersDefaultActions) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig struct {
	// The server groups to which requests are forwarded.
	ServerGroupTuples []*ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) GetServerGroupTuples() []*ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// The ID of the server group to which requests are forwarded.
	//
	// example:
	//
	// rsp-cige6j****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersLogConfig struct {
	// Indicates whether custom headers are carried in the access log. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// The configurations of xtrace.
	AccessLogTracingConfig *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig `json:"AccessLogTracingConfig,omitempty" xml:"AccessLogTracingConfig,omitempty" type:"Struct"`
}

func (s ListListenersResponseBodyListenersLogConfig) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersLogConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersLogConfig) GetAccessLogRecordCustomizedHeadersEnabled() *bool {
	return s.AccessLogRecordCustomizedHeadersEnabled
}

func (s *ListListenersResponseBodyListenersLogConfig) GetAccessLogTracingConfig() *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig {
	return s.AccessLogTracingConfig
}

func (s *ListListenersResponseBodyListenersLogConfig) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *ListListenersResponseBodyListenersLogConfig {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfig) SetAccessLogTracingConfig(v *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) *ListListenersResponseBodyListenersLogConfig {
	s.AccessLogTracingConfig = v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfig) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig struct {
	// Indicates whether xtrace is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter can be set to **true*	- only when the access log feature of ALB is enabled by setting **AccessLogEnabled*	- to true.
	//
	// example:
	//
	// true
	TracingEnabled *bool `json:"TracingEnabled,omitempty" xml:"TracingEnabled,omitempty"`
	// The sampling rate of xtrace. Valid values: **1 to 10000**.
	//
	// >  This parameter takes effect when **TracingEnabled*	- is set to **true**.
	//
	// example:
	//
	// 100
	TracingSample *int32 `json:"TracingSample,omitempty" xml:"TracingSample,omitempty"`
	// The type of xtrace. The value is set to **Zipkin**.
	//
	// >  This parameter takes effect when **TracingEnabled*	- is set to **true**.
	//
	// example:
	//
	// Zipkin
	TracingType *string `json:"TracingType,omitempty" xml:"TracingType,omitempty"`
}

func (s ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) GetTracingEnabled() *bool {
	return s.TracingEnabled
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) GetTracingSample() *int32 {
	return s.TracingSample
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) GetTracingType() *string {
	return s.TracingType
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) SetTracingEnabled(v bool) *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig {
	s.TracingEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) SetTracingSample(v int32) *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig {
	s.TracingSample = &v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) SetTracingType(v string) *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig {
	s.TracingType = &v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersQuicConfig struct {
	// The ID of the QUIC listener associated with the ALB instance. This parameter is required if the **QuicUpgradeEnabled*	- parameter is set to **true**. Only HTTPS listeners support this parameter.
	//
	// >  The existing listener and QUIC listener must be to the same ALB instance, and the QUIC listener has not been associated with an ALB instance.
	//
	// example:
	//
	// lsr-bp1bpn908w4nbw****
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// Indicates whether QUIC upgrade is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s ListListenersResponseBodyListenersQuicConfig) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersQuicConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersQuicConfig) GetQuicListenerId() *string {
	return s.QuicListenerId
}

func (s *ListListenersResponseBodyListenersQuicConfig) GetQuicUpgradeEnabled() *bool {
	return s.QuicUpgradeEnabled
}

func (s *ListListenersResponseBodyListenersQuicConfig) SetQuicListenerId(v string) *ListListenersResponseBodyListenersQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListenersQuicConfig) SetQuicUpgradeEnabled(v bool) *ListListenersResponseBodyListenersQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersQuicConfig) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersTags struct {
	// The tag key. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListListenersResponseBodyListenersTags) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersTags) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersTags) GetKey() *string {
	return s.Key
}

func (s *ListListenersResponseBodyListenersTags) GetValue() *string {
	return s.Value
}

func (s *ListListenersResponseBodyListenersTags) SetKey(v string) *ListListenersResponseBodyListenersTags {
	s.Key = &v
	return s
}

func (s *ListListenersResponseBodyListenersTags) SetValue(v string) *ListListenersResponseBodyListenersTags {
	s.Value = &v
	return s
}

func (s *ListListenersResponseBodyListenersTags) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersXForwardedForConfig struct {
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertClientVerifyEnabled*	- is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// test_client-verify-alias_123456
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// Indicates whether the `X-Forwarded-Clientcert-clientverify` header is used to obtain the verification result of the client certificate. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertFingerprintEnabled*	- is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// >  Only HTTPS listeners support this parameter.
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
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertIssuerDNEnabled*	- is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// >  Only HTTPS listeners support this parameter.
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
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// The name of the custom header. This parameter takes effect only when **XForwardedForClientCertSubjectDNEnabled*	- is set to **true**.
	//
	// The name must be 1 to 40 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// >  Only HTTPS listeners support this parameter.
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
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// Indicates whether the X-Forwarded-For header is used to preserver client IP addresses for the ALB instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is returned only for HTTP and HTTPS listeners.
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
	// Indicates whether the `X-Forwarded-Client-Port` header is used to retrieve the client port. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is returned only for HTTP and HTTPS listeners.
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
	// > 	- If this parameter is set to **true**, the default value of the **XForwardedForProcessingMode*	- parameter is **append**. You can change it to **remove**.
	//
	// > 	- If this parameter is set to **false**, the `X-Forwarded-For` header in the request is not modified in any way before the request is sent to backend servers.
	//
	// > 	- Both HTTP and HTTPS listeners support this parameter.
	//
	// example:
	//
	// true
	XForwardedForEnabled *bool `json:"XForwardedForEnabled,omitempty" xml:"XForwardedForEnabled,omitempty"`
	// Specifies whether to use the `X-Forwarded-Host` header to retrieve client domain names. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  HTTP, HTTPS, and QUIC listeners all support this parameter.
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
	// > 	- Both HTTP and HTTPS listeners support this parameter.
	//
	// example:
	//
	// append
	XForwardedForProcessingMode *string `json:"XForwardedForProcessingMode,omitempty" xml:"XForwardedForProcessingMode,omitempty"`
	// Indicates whether the `X-Forwarded-Proto` header is used to retrieve the listener protocol. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is supported by HTTP, HTTPS, and QUIC listeners.
	//
	// example:
	//
	// true
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// Specifies whether to use the `SLB-ID` header to retrieve the ID of the ALB instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is supported by HTTP, HTTPS, and QUIC listeners.
	//
	// example:
	//
	// true
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// Indicates whether the `X-Forwarded-Port` header is used to retrieve the listener port of the ALB instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is supported by HTTP, HTTPS, and QUIC listeners.
	//
	// example:
	//
	// true
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s ListListenersResponseBodyListenersXForwardedForConfig) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientCertClientVerifyAlias() *string {
	return s.XForwardedForClientCertClientVerifyAlias
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientCertClientVerifyEnabled() *bool {
	return s.XForwardedForClientCertClientVerifyEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientCertFingerprintAlias() *string {
	return s.XForwardedForClientCertFingerprintAlias
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientCertFingerprintEnabled() *bool {
	return s.XForwardedForClientCertFingerprintEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientCertIssuerDNAlias() *string {
	return s.XForwardedForClientCertIssuerDNAlias
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientCertIssuerDNEnabled() *bool {
	return s.XForwardedForClientCertIssuerDNEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientCertSubjectDNAlias() *string {
	return s.XForwardedForClientCertSubjectDNAlias
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientCertSubjectDNEnabled() *bool {
	return s.XForwardedForClientCertSubjectDNEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientSourceIpsEnabled() *bool {
	return s.XForwardedForClientSourceIpsEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientSourceIpsTrusted() *string {
	return s.XForwardedForClientSourceIpsTrusted
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForClientSrcPortEnabled() *bool {
	return s.XForwardedForClientSrcPortEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForEnabled() *bool {
	return s.XForwardedForEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForHostEnabled() *bool {
	return s.XForwardedForHostEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForProcessingMode() *string {
	return s.XForwardedForProcessingMode
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForProtoEnabled() *bool {
	return s.XForwardedForProtoEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForSLBIdEnabled() *bool {
	return s.XForwardedForSLBIdEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForSLBPortEnabled() *bool {
	return s.XForwardedForSLBPortEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientSourceIpsEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientSourceIpsEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientSourceIpsTrusted(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientSourceIpsTrusted = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForHostEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForHostEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForProcessingMode(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForProcessingMode = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) Validate() error {
	return dara.Validate(s)
}
