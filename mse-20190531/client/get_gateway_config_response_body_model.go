// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetGatewayConfigResponseBody
	GetCode() *int32
	SetData(v *GetGatewayConfigResponseBodyData) *GetGatewayConfigResponseBody
	GetData() *GetGatewayConfigResponseBodyData
	SetDynamicCode(v string) *GetGatewayConfigResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetGatewayConfigResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetGatewayConfigResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GetGatewayConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGatewayConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGatewayConfigResponseBody
	GetSuccess() *bool
}

type GetGatewayConfigResponseBody struct {
	// The code used to synchronize with the success field.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data overview.
	Data *GetGatewayConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// This field is not used. Ignore this field.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// This field is not used. Ignore this field.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F6092602-C357-4750-89D9-E572FBEA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// - `true`: The request is successful.
	//
	// - `false`: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayConfigResponseBody) GetData() *GetGatewayConfigResponseBodyData {
	return s.Data
}

func (s *GetGatewayConfigResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetGatewayConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetGatewayConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetGatewayConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayConfigResponseBody) SetCode(v int32) *GetGatewayConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayConfigResponseBody) SetData(v *GetGatewayConfigResponseBodyData) *GetGatewayConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayConfigResponseBody) SetDynamicCode(v string) *GetGatewayConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetGatewayConfigResponseBody) SetDynamicMessage(v string) *GetGatewayConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetGatewayConfigResponseBody) SetErrorCode(v string) *GetGatewayConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGatewayConfigResponseBody) SetHttpStatusCode(v int32) *GetGatewayConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayConfigResponseBody) SetMessage(v string) *GetGatewayConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayConfigResponseBody) SetRequestId(v string) *GetGatewayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayConfigResponseBody) SetSuccess(v bool) *GetGatewayConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayConfigResponseBodyData struct {
	// The custom log format.
	//
	// example:
	//
	// {"requestMetadata":[],"responseHeaders":[],"requestHeaders":["x-test"],"responseMetadata":[]}
	AccessLogHeader *string `json:"AccessLogHeader,omitempty" xml:"AccessLogHeader,omitempty"`
	// The buffer size for a single gateway connection. Configuring this parameter affects throughput and gateway memory usage. Unit: bytes.
	//
	// example:
	//
	// 32768
	DownstreamConnectionBufferLimits *int32 `json:"DownstreamConnectionBufferLimits,omitempty" xml:"DownstreamConnectionBufferLimits,omitempty"`
	// The maximum number of concurrent streams on a single connection when the client uses HTTP/2. This parameter takes effect at the request level.
	//
	// example:
	//
	// 100
	DownstreamHttp2MaxConcurrentStream *int32 `json:"DownstreamHttp2MaxConcurrentStream,omitempty" xml:"DownstreamHttp2MaxConcurrentStream,omitempty"`
	// The idle timeout period for connections from the client to the gateway. Unit: seconds. If no request is received within the specified period, the gateway disconnects the connection. Valid values: 0 to 600.
	//
	// example:
	//
	// 180
	DownstreamIdleTime *int32 `json:"DownstreamIdleTime,omitempty" xml:"DownstreamIdleTime,omitempty"`
	// Specifies whether to enable custom authentication configuration push for self-managed authentication services. When enabled, changes to authentication rules do not cause connection interruptions. This feature is applicable to WebSocket and online business scenarios.
	//
	// example:
	//
	// false
	EnableCustomAuthConfigPush *string `json:"EnableCustomAuthConfigPush,omitempty" xml:"EnableCustomAuthConfigPush,omitempty"`
	// Specifies whether to generate a request ID in the request header for request tracking. This parameter takes effect at the request level.
	//
	// example:
	//
	// true
	EnableGenerateRequestId *bool `json:"EnableGenerateRequestId,omitempty" xml:"EnableGenerateRequestId,omitempty"`
	// Specifies whether to enable Gzip compression for requests and responses. Enabling this feature reduces gateway traffic but increases gateway CPU consumption.
	//
	// example:
	//
	// false
	EnableGzip *bool `json:"EnableGzip,omitempty" xml:"EnableGzip,omitempty"`
	// Specifies whether to enable Gzip hardware acceleration.
	EnableGzipHardwareAccelerate *bool `json:"EnableGzipHardwareAccelerate,omitempty" xml:"EnableGzipHardwareAccelerate,omitempty"`
	// Specifies whether to enable hardware acceleration. This configuration fails if the region does not support hardware acceleration at the underlying layer or if hardware acceleration was not specified during purchase.
	//
	// example:
	//
	// true
	EnableHardwareAccelerate *bool `json:"EnableHardwareAccelerate,omitempty" xml:"EnableHardwareAccelerate,omitempty"`
	// Specifies whether to use HTTP/2 during negotiation between the server and the client. This parameter takes effect at the request level. Changes take effect after 1 to 2 minutes.
	//
	// example:
	//
	// false
	EnableHttp2 *bool `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
	// Specifies whether the HTTP/3 protocol is supported.
	//
	// example:
	//
	// false
	EnableHttp3 *bool `json:"EnableHttp3,omitempty" xml:"EnableHttp3,omitempty"`
	// The tag filtering configuration for Kubernetes service source workloads.
	EnableK8sSourceWorkloadFilter *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter `json:"EnableK8sSourceWorkloadFilter,omitempty" xml:"EnableK8sSourceWorkloadFilter,omitempty" type:"Struct"`
	// Specifies whether to enable the PROXY protocol. If the gateway ingress traffic is routed through a Network Load Balancer (NLB) and this feature is disabled, the real IP address of the client cannot be obtained. Enabling this feature does not affect non-PROXY protocol requests.
	//
	// example:
	//
	// true
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// Specifies whether to merge consecutive slashes in a request URL. For example, if this feature is enabled, the extra slash in `www.a.com//b` is merged, making it equivalent to `www.a.com/b`. This parameter takes effect at the request level.
	//
	// example:
	//
	// true
	EnableSlashMerge *bool `json:"EnableSlashMerge,omitempty" xml:"EnableSlashMerge,omitempty"`
	// Specifies whether WAF is enabled.
	//
	// example:
	//
	// true
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// The configuration for trusted X-Forwarded-For source CIDR blocks.
	EnableXffTrustedCidrs *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs `json:"EnableXffTrustedCidrs,omitempty" xml:"EnableXffTrustedCidrs,omitempty" type:"Struct"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-61f465fa2dd044069e2208c4912*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The initial connection-level window size negotiated between the gateway and the client when HTTP/2 is used. Unit: bytes. This parameter takes effect at the request level.
	//
	// example:
	//
	// 1048576
	InitialConnectionWindowSize *int32 `json:"InitialConnectionWindowSize,omitempty" xml:"InitialConnectionWindowSize,omitempty"`
	// The initial stream window size negotiated between the gateway and the client when HTTP/2 is used. Unit: bytes. This parameter takes effect at the request level.
	//
	// example:
	//
	// 65535
	InitialStreamWindowSize *int32 `json:"InitialStreamWindowSize,omitempty" xml:"InitialStreamWindowSize,omitempty"`
	// The keep-alive time in the Keep-Alive response header. Unit: seconds. If this parameter is set to 0, the response header is not returned.
	//
	// example:
	//
	// 0
	KeepaliveHeaderTimeout *int32 `json:"KeepaliveHeaderTimeout,omitempty" xml:"KeepaliveHeaderTimeout,omitempty"`
	// Indicates whether lightweight metric collection is enabled.
	LiteMetrics *bool `json:"LiteMetrics,omitempty" xml:"LiteMetrics,omitempty"`
	// Not available for use.
	//
	// example:
	//
	// ""
	LogFilterConfig *string `json:"LogFilterConfig,omitempty" xml:"LogFilterConfig,omitempty"`
	// The log collection configuration for request and response headers and bodies. The content is a JSON string.
	//
	// example:
	//
	// {"request":{"headers":{"enabled":true}}}
	LogRequestResponseConfig *string `json:"LogRequestResponseConfig,omitempty" xml:"LogRequestResponseConfig,omitempty"`
	// The list of unsupported configurations.
	//
	// example:
	//
	// "EnableProxyProtocol,EnableCustomAuthConfigPush"
	NoSupportedConfigList *string `json:"NoSupportedConfigList,omitempty" xml:"NoSupportedConfigList,omitempty"`
	// The action to take when a URI path contains escaped slashes such as %2F, %2f, %5C, or %5c. Default value: KEEP_UNCHANGED. Valid values:
	//
	// - KEEP_UNCHANGED
	//
	// - REJECT_REQUEST
	//
	// - UNESCAPE_AND_REDIRECT
	//
	// - UNESCAPE_AND_FORWARD
	//
	// example:
	//
	// KEEP_UNCHANGED
	PathWithEscapedSlashes *string `json:"PathWithEscapedSlashes,omitempty" xml:"PathWithEscapedSlashes,omitempty"`
	// Specifies whether to preserve the request ID passed by the client.
	PreserveExternalRequestID *bool `json:"PreserveExternalRequestID,omitempty" xml:"PreserveExternalRequestID,omitempty"`
	// Specifies whether to preserve the original case of request and response headers. In HTTP/1.1, headers are case-insensitive. To ensure compatibility with HTTP/2, headers are converted to lowercase by default.
	//
	// example:
	//
	// true
	PreserveHeaderFormat *bool `json:"PreserveHeaderFormat,omitempty" xml:"PreserveHeaderFormat,omitempty"`
	// The log delivery information.
	SlsConfigDetails *GetGatewayConfigResponseBodyDataSlsConfigDetails `json:"SlsConfigDetails,omitempty" xml:"SlsConfigDetails,omitempty" type:"Struct"`
	// Indicates whether the gateway supports Web Application Firewall (WAF).
	//
	// example:
	//
	// true
	SupportWaf *bool `json:"SupportWaf,omitempty" xml:"SupportWaf,omitempty"`
	// The default idle timeout period for upstream connections.
	//
	// example:
	//
	// 30
	UpstreamIdleTimeout *int32 `json:"UpstreamIdleTimeout,omitempty" xml:"UpstreamIdleTimeout,omitempty"`
	// The keep-alive time maintained for WebSocket connections when connections must be closed during gateway instance upgrades or restarts. Unit: seconds.
	//
	// example:
	//
	// 30
	WebsocketTermGracePeriod *int32 `json:"WebsocketTermGracePeriod,omitempty" xml:"WebsocketTermGracePeriod,omitempty"`
	// The number of trusted proxies in front of the gateway. This parameter affects whether the gateway uses request headers such as X-Forwarded-For and X-Request-ID generated by the client. This parameter takes effect at the request level.
	//
	// example:
	//
	// 0
	XffTrustedNum *int32 `json:"XffTrustedNum,omitempty" xml:"XffTrustedNum,omitempty"`
	// The Tracing Analysis information.
	XtraceDetails *GetGatewayConfigResponseBodyDataXtraceDetails `json:"XtraceDetails,omitempty" xml:"XtraceDetails,omitempty" type:"Struct"`
	// The compression algorithm used when compression is enabled. Valid values: brotli and gzip. If this parameter is set and EnableGzip is disabled, EnableGzip is automatically enabled.
	//
	// example:
	//
	// ["gzip"]
	ZipAlgorithm *string `json:"ZipAlgorithm,omitempty" xml:"ZipAlgorithm,omitempty"`
}

func (s GetGatewayConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayConfigResponseBodyData) GetAccessLogHeader() *string {
	return s.AccessLogHeader
}

func (s *GetGatewayConfigResponseBodyData) GetDownstreamConnectionBufferLimits() *int32 {
	return s.DownstreamConnectionBufferLimits
}

func (s *GetGatewayConfigResponseBodyData) GetDownstreamHttp2MaxConcurrentStream() *int32 {
	return s.DownstreamHttp2MaxConcurrentStream
}

func (s *GetGatewayConfigResponseBodyData) GetDownstreamIdleTime() *int32 {
	return s.DownstreamIdleTime
}

func (s *GetGatewayConfigResponseBodyData) GetEnableCustomAuthConfigPush() *string {
	return s.EnableCustomAuthConfigPush
}

func (s *GetGatewayConfigResponseBodyData) GetEnableGenerateRequestId() *bool {
	return s.EnableGenerateRequestId
}

func (s *GetGatewayConfigResponseBodyData) GetEnableGzip() *bool {
	return s.EnableGzip
}

func (s *GetGatewayConfigResponseBodyData) GetEnableGzipHardwareAccelerate() *bool {
	return s.EnableGzipHardwareAccelerate
}

func (s *GetGatewayConfigResponseBodyData) GetEnableHardwareAccelerate() *bool {
	return s.EnableHardwareAccelerate
}

func (s *GetGatewayConfigResponseBodyData) GetEnableHttp2() *bool {
	return s.EnableHttp2
}

func (s *GetGatewayConfigResponseBodyData) GetEnableHttp3() *bool {
	return s.EnableHttp3
}

func (s *GetGatewayConfigResponseBodyData) GetEnableK8sSourceWorkloadFilter() *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter {
	return s.EnableK8sSourceWorkloadFilter
}

func (s *GetGatewayConfigResponseBodyData) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *GetGatewayConfigResponseBodyData) GetEnableSlashMerge() *bool {
	return s.EnableSlashMerge
}

func (s *GetGatewayConfigResponseBodyData) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *GetGatewayConfigResponseBodyData) GetEnableXffTrustedCidrs() *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs {
	return s.EnableXffTrustedCidrs
}

func (s *GetGatewayConfigResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayConfigResponseBodyData) GetInitialConnectionWindowSize() *int32 {
	return s.InitialConnectionWindowSize
}

func (s *GetGatewayConfigResponseBodyData) GetInitialStreamWindowSize() *int32 {
	return s.InitialStreamWindowSize
}

func (s *GetGatewayConfigResponseBodyData) GetKeepaliveHeaderTimeout() *int32 {
	return s.KeepaliveHeaderTimeout
}

func (s *GetGatewayConfigResponseBodyData) GetLiteMetrics() *bool {
	return s.LiteMetrics
}

func (s *GetGatewayConfigResponseBodyData) GetLogFilterConfig() *string {
	return s.LogFilterConfig
}

func (s *GetGatewayConfigResponseBodyData) GetLogRequestResponseConfig() *string {
	return s.LogRequestResponseConfig
}

func (s *GetGatewayConfigResponseBodyData) GetNoSupportedConfigList() *string {
	return s.NoSupportedConfigList
}

func (s *GetGatewayConfigResponseBodyData) GetPathWithEscapedSlashes() *string {
	return s.PathWithEscapedSlashes
}

func (s *GetGatewayConfigResponseBodyData) GetPreserveExternalRequestID() *bool {
	return s.PreserveExternalRequestID
}

func (s *GetGatewayConfigResponseBodyData) GetPreserveHeaderFormat() *bool {
	return s.PreserveHeaderFormat
}

func (s *GetGatewayConfigResponseBodyData) GetSlsConfigDetails() *GetGatewayConfigResponseBodyDataSlsConfigDetails {
	return s.SlsConfigDetails
}

func (s *GetGatewayConfigResponseBodyData) GetSupportWaf() *bool {
	return s.SupportWaf
}

func (s *GetGatewayConfigResponseBodyData) GetUpstreamIdleTimeout() *int32 {
	return s.UpstreamIdleTimeout
}

func (s *GetGatewayConfigResponseBodyData) GetWebsocketTermGracePeriod() *int32 {
	return s.WebsocketTermGracePeriod
}

func (s *GetGatewayConfigResponseBodyData) GetXffTrustedNum() *int32 {
	return s.XffTrustedNum
}

func (s *GetGatewayConfigResponseBodyData) GetXtraceDetails() *GetGatewayConfigResponseBodyDataXtraceDetails {
	return s.XtraceDetails
}

func (s *GetGatewayConfigResponseBodyData) GetZipAlgorithm() *string {
	return s.ZipAlgorithm
}

func (s *GetGatewayConfigResponseBodyData) SetAccessLogHeader(v string) *GetGatewayConfigResponseBodyData {
	s.AccessLogHeader = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetDownstreamConnectionBufferLimits(v int32) *GetGatewayConfigResponseBodyData {
	s.DownstreamConnectionBufferLimits = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetDownstreamHttp2MaxConcurrentStream(v int32) *GetGatewayConfigResponseBodyData {
	s.DownstreamHttp2MaxConcurrentStream = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetDownstreamIdleTime(v int32) *GetGatewayConfigResponseBodyData {
	s.DownstreamIdleTime = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableCustomAuthConfigPush(v string) *GetGatewayConfigResponseBodyData {
	s.EnableCustomAuthConfigPush = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableGenerateRequestId(v bool) *GetGatewayConfigResponseBodyData {
	s.EnableGenerateRequestId = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableGzip(v bool) *GetGatewayConfigResponseBodyData {
	s.EnableGzip = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableGzipHardwareAccelerate(v bool) *GetGatewayConfigResponseBodyData {
	s.EnableGzipHardwareAccelerate = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableHardwareAccelerate(v bool) *GetGatewayConfigResponseBodyData {
	s.EnableHardwareAccelerate = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableHttp2(v bool) *GetGatewayConfigResponseBodyData {
	s.EnableHttp2 = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableHttp3(v bool) *GetGatewayConfigResponseBodyData {
	s.EnableHttp3 = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableK8sSourceWorkloadFilter(v *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) *GetGatewayConfigResponseBodyData {
	s.EnableK8sSourceWorkloadFilter = v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableProxyProtocol(v bool) *GetGatewayConfigResponseBodyData {
	s.EnableProxyProtocol = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableSlashMerge(v bool) *GetGatewayConfigResponseBodyData {
	s.EnableSlashMerge = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableWaf(v bool) *GetGatewayConfigResponseBodyData {
	s.EnableWaf = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetEnableXffTrustedCidrs(v *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs) *GetGatewayConfigResponseBodyData {
	s.EnableXffTrustedCidrs = v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetGatewayUniqueId(v string) *GetGatewayConfigResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetInitialConnectionWindowSize(v int32) *GetGatewayConfigResponseBodyData {
	s.InitialConnectionWindowSize = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetInitialStreamWindowSize(v int32) *GetGatewayConfigResponseBodyData {
	s.InitialStreamWindowSize = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetKeepaliveHeaderTimeout(v int32) *GetGatewayConfigResponseBodyData {
	s.KeepaliveHeaderTimeout = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetLiteMetrics(v bool) *GetGatewayConfigResponseBodyData {
	s.LiteMetrics = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetLogFilterConfig(v string) *GetGatewayConfigResponseBodyData {
	s.LogFilterConfig = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetLogRequestResponseConfig(v string) *GetGatewayConfigResponseBodyData {
	s.LogRequestResponseConfig = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetNoSupportedConfigList(v string) *GetGatewayConfigResponseBodyData {
	s.NoSupportedConfigList = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetPathWithEscapedSlashes(v string) *GetGatewayConfigResponseBodyData {
	s.PathWithEscapedSlashes = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetPreserveExternalRequestID(v bool) *GetGatewayConfigResponseBodyData {
	s.PreserveExternalRequestID = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetPreserveHeaderFormat(v bool) *GetGatewayConfigResponseBodyData {
	s.PreserveHeaderFormat = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetSlsConfigDetails(v *GetGatewayConfigResponseBodyDataSlsConfigDetails) *GetGatewayConfigResponseBodyData {
	s.SlsConfigDetails = v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetSupportWaf(v bool) *GetGatewayConfigResponseBodyData {
	s.SupportWaf = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetUpstreamIdleTimeout(v int32) *GetGatewayConfigResponseBodyData {
	s.UpstreamIdleTimeout = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetWebsocketTermGracePeriod(v int32) *GetGatewayConfigResponseBodyData {
	s.WebsocketTermGracePeriod = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetXffTrustedNum(v int32) *GetGatewayConfigResponseBodyData {
	s.XffTrustedNum = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetXtraceDetails(v *GetGatewayConfigResponseBodyDataXtraceDetails) *GetGatewayConfigResponseBodyData {
	s.XtraceDetails = v
	return s
}

func (s *GetGatewayConfigResponseBodyData) SetZipAlgorithm(v string) *GetGatewayConfigResponseBodyData {
	s.ZipAlgorithm = &v
	return s
}

func (s *GetGatewayConfigResponseBodyData) Validate() error {
	if s.EnableK8sSourceWorkloadFilter != nil {
		if err := s.EnableK8sSourceWorkloadFilter.Validate(); err != nil {
			return err
		}
	}
	if s.EnableXffTrustedCidrs != nil {
		if err := s.EnableXffTrustedCidrs.Validate(); err != nil {
			return err
		}
	}
	if s.SlsConfigDetails != nil {
		if err := s.SlsConfigDetails.Validate(); err != nil {
			return err
		}
	}
	if s.XtraceDetails != nil {
		if err := s.XtraceDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter struct {
	// Specifies whether to enable tag filtering for Kubernetes service source workloads.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The tag filtering method. Valid values: EQUAL (matches the specified tag) and NOT_EQUAL (excludes the specified tag).
	//
	// example:
	//
	// EQUAL
	FilterOpt *string `json:"FilterOpt,omitempty" xml:"FilterOpt,omitempty"`
	// The tag key used to filter Kubernetes workloads.
	//
	// example:
	//
	// key
	LabelKey *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	// The tag value used to filter Kubernetes workloads.
	//
	// example:
	//
	// value
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) GoString() string {
	return s.String()
}

func (s *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) GetEnable() *bool {
	return s.Enable
}

func (s *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) GetFilterOpt() *string {
	return s.FilterOpt
}

func (s *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) GetLabelKey() *string {
	return s.LabelKey
}

func (s *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) GetLabelValue() *string {
	return s.LabelValue
}

func (s *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) SetEnable(v bool) *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter {
	s.Enable = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) SetFilterOpt(v string) *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter {
	s.FilterOpt = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) SetLabelKey(v string) *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter {
	s.LabelKey = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) SetLabelValue(v string) *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter {
	s.LabelValue = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter) Validate() error {
	return dara.Validate(s)
}

type GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs struct {
	// Specifies whether to enable the configuration for trusted X-Forwarded-For source CIDR blocks.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The list of trusted source IP addresses or CIDR blocks.
	//
	// example:
	//
	// 10.0.0.0/8
	IpListContent *string `json:"IpListContent,omitempty" xml:"IpListContent,omitempty"`
}

func (s GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs) GoString() string {
	return s.String()
}

func (s *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs) GetEnable() *bool {
	return s.Enable
}

func (s *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs) GetIpListContent() *string {
	return s.IpListContent
}

func (s *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs) SetEnable(v bool) *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs {
	s.Enable = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs) SetIpListContent(v string) *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs {
	s.IpListContent = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs) Validate() error {
	return dara.Validate(s)
}

type GetGatewayConfigResponseBodyDataSlsConfigDetails struct {
	// The gateway ID.
	//
	// example:
	//
	// 500
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-61f465fa2dd044069e2208c4912*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Not available for use.
	//
	// example:
	//
	// ""
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Not available for use.
	//
	// example:
	//
	// ""
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Not available for use.
	//
	// example:
	//
	// 0
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether log delivery is enabled.
	//
	// example:
	//
	// true
	LogOn *bool `json:"LogOn,omitempty" xml:"LogOn,omitempty"`
	// The destination Logstore for log delivery. Valid values:
	//
	// - mse_gw_access_log: Nginx Ingress compatibility is disabled.
	//
	// - nginx-ingress: Nginx Ingress compatibility is enabled.
	//
	// example:
	//
	// mse_gw_access_log
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// Specifies whether Nginx Ingress compatibility is enabled.
	//
	// example:
	//
	// false
	NginxCompatible *bool `json:"NginxCompatible,omitempty" xml:"NginxCompatible,omitempty"`
	// The destination Simple Log Service project for log delivery.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetGatewayConfigResponseBodyDataSlsConfigDetails) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayConfigResponseBodyDataSlsConfigDetails) GoString() string {
	return s.String()
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) GetLogOn() *bool {
	return s.LogOn
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) GetNginxCompatible() *bool {
	return s.NginxCompatible
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) SetGatewayId(v int64) *GetGatewayConfigResponseBodyDataSlsConfigDetails {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) SetGatewayUniqueId(v string) *GetGatewayConfigResponseBodyDataSlsConfigDetails {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) SetGmtCreate(v string) *GetGatewayConfigResponseBodyDataSlsConfigDetails {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) SetGmtModified(v string) *GetGatewayConfigResponseBodyDataSlsConfigDetails {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) SetId(v int64) *GetGatewayConfigResponseBodyDataSlsConfigDetails {
	s.Id = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) SetLogOn(v bool) *GetGatewayConfigResponseBodyDataSlsConfigDetails {
	s.LogOn = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) SetLogStoreName(v string) *GetGatewayConfigResponseBodyDataSlsConfigDetails {
	s.LogStoreName = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) SetNginxCompatible(v bool) *GetGatewayConfigResponseBodyDataSlsConfigDetails {
	s.NginxCompatible = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) SetProjectName(v string) *GetGatewayConfigResponseBodyDataSlsConfigDetails {
	s.ProjectName = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataSlsConfigDetails) Validate() error {
	return dara.Validate(s)
}

type GetGatewayConfigResponseBodyDataXtraceDetails struct {
	// The gateway ID.
	//
	// example:
	//
	// 500
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-61f465fa2dd044069e2208c4912*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The time when the Tracing Analysis information was created.
	//
	// example:
	//
	// 2024-08-12T09:39:07.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the Tracing Analysis information was modified.
	//
	// example:
	//
	// 2024-08-12T09:39:07.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the Tracing Analysis information.
	//
	// example:
	//
	// 255
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The sampling rate (%).
	//
	// example:
	//
	// 100
	Sample *int32 `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// The SkyWalking service ID for self-managed SkyWalking.
	//
	// example:
	//
	// 333
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The SkyWalking service port for self-managed SkyWalking.
	//
	// example:
	//
	// 8888
	ServicePort *string `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// Specifies whether Tracing Analysis is enabled.
	//
	// example:
	//
	// true
	TraceOn *bool `json:"TraceOn,omitempty" xml:"TraceOn,omitempty"`
	// The Tracing Analysis type. Valid values:
	//
	// - XTRACE: Managed Service for OpenTelemetry with the Zipkin Tracing Analysis protocol.
	//
	// - SKYWALKING: Self-managed SkyWalking.
	//
	// - OPENTELEMETRY: Managed Service for OpenTelemetry with the OpenTelemetry Tracing Analysis protocol.
	//
	// - OTSKYWALKING: Managed Service for OpenTelemetry with the SkyWalking Tracing Analysis protocol.
	//
	// example:
	//
	// OPENTELEMETRY
	TraceType *string `json:"TraceType,omitempty" xml:"TraceType,omitempty"`
}

func (s GetGatewayConfigResponseBodyDataXtraceDetails) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayConfigResponseBodyDataXtraceDetails) GoString() string {
	return s.String()
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) GetSample() *int32 {
	return s.Sample
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) GetServicePort() *string {
	return s.ServicePort
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) GetTraceOn() *bool {
	return s.TraceOn
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) GetTraceType() *string {
	return s.TraceType
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) SetGatewayId(v int64) *GetGatewayConfigResponseBodyDataXtraceDetails {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) SetGatewayUniqueId(v string) *GetGatewayConfigResponseBodyDataXtraceDetails {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) SetGmtCreate(v string) *GetGatewayConfigResponseBodyDataXtraceDetails {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) SetGmtModified(v string) *GetGatewayConfigResponseBodyDataXtraceDetails {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) SetId(v int64) *GetGatewayConfigResponseBodyDataXtraceDetails {
	s.Id = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) SetSample(v int32) *GetGatewayConfigResponseBodyDataXtraceDetails {
	s.Sample = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) SetServiceId(v int64) *GetGatewayConfigResponseBodyDataXtraceDetails {
	s.ServiceId = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) SetServicePort(v string) *GetGatewayConfigResponseBodyDataXtraceDetails {
	s.ServicePort = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) SetTraceOn(v bool) *GetGatewayConfigResponseBodyDataXtraceDetails {
	s.TraceOn = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) SetTraceType(v string) *GetGatewayConfigResponseBodyDataXtraceDetails {
	s.TraceType = &v
	return s
}

func (s *GetGatewayConfigResponseBodyDataXtraceDetails) Validate() error {
	return dara.Validate(s)
}
