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
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetGatewayConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F6092602-C357-4750-89D9-E572FBEA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// {"requestMetadata":[],"responseHeaders":[],"requestHeaders":["x-test"],"responseMetadata":[]}
	AccessLogHeader *string `json:"AccessLogHeader,omitempty" xml:"AccessLogHeader,omitempty"`
	// example:
	//
	// 32768
	DownstreamConnectionBufferLimits *int32 `json:"DownstreamConnectionBufferLimits,omitempty" xml:"DownstreamConnectionBufferLimits,omitempty"`
	// example:
	//
	// 100
	DownstreamHttp2MaxConcurrentStream *int32 `json:"DownstreamHttp2MaxConcurrentStream,omitempty" xml:"DownstreamHttp2MaxConcurrentStream,omitempty"`
	// example:
	//
	// 180
	DownstreamIdleTime *int32 `json:"DownstreamIdleTime,omitempty" xml:"DownstreamIdleTime,omitempty"`
	// example:
	//
	// false
	EnableCustomAuthConfigPush *string `json:"EnableCustomAuthConfigPush,omitempty" xml:"EnableCustomAuthConfigPush,omitempty"`
	// example:
	//
	// true
	EnableGenerateRequestId *bool `json:"EnableGenerateRequestId,omitempty" xml:"EnableGenerateRequestId,omitempty"`
	// example:
	//
	// false
	EnableGzip                   *bool `json:"EnableGzip,omitempty" xml:"EnableGzip,omitempty"`
	EnableGzipHardwareAccelerate *bool `json:"EnableGzipHardwareAccelerate,omitempty" xml:"EnableGzipHardwareAccelerate,omitempty"`
	// example:
	//
	// true
	EnableHardwareAccelerate *bool `json:"EnableHardwareAccelerate,omitempty" xml:"EnableHardwareAccelerate,omitempty"`
	// example:
	//
	// false
	EnableHttp2 *bool `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
	// example:
	//
	// false
	EnableHttp3                   *bool                                                          `json:"EnableHttp3,omitempty" xml:"EnableHttp3,omitempty"`
	EnableK8sSourceWorkloadFilter *GetGatewayConfigResponseBodyDataEnableK8sSourceWorkloadFilter `json:"EnableK8sSourceWorkloadFilter,omitempty" xml:"EnableK8sSourceWorkloadFilter,omitempty" type:"Struct"`
	// example:
	//
	// true
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// example:
	//
	// true
	EnableSlashMerge *bool `json:"EnableSlashMerge,omitempty" xml:"EnableSlashMerge,omitempty"`
	// example:
	//
	// true
	EnableWaf             *bool                                                  `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	EnableXffTrustedCidrs *GetGatewayConfigResponseBodyDataEnableXffTrustedCidrs `json:"EnableXffTrustedCidrs,omitempty" xml:"EnableXffTrustedCidrs,omitempty" type:"Struct"`
	// example:
	//
	// gw-61f465fa2dd044069e2208c4912*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 1048576
	InitialConnectionWindowSize *int32 `json:"InitialConnectionWindowSize,omitempty" xml:"InitialConnectionWindowSize,omitempty"`
	// example:
	//
	// 65535
	InitialStreamWindowSize *int32 `json:"InitialStreamWindowSize,omitempty" xml:"InitialStreamWindowSize,omitempty"`
	// example:
	//
	// 0
	KeepaliveHeaderTimeout *int32 `json:"KeepaliveHeaderTimeout,omitempty" xml:"KeepaliveHeaderTimeout,omitempty"`
	// example:
	//
	// ""
	LogFilterConfig *string `json:"LogFilterConfig,omitempty" xml:"LogFilterConfig,omitempty"`
	// example:
	//
	// "EnableProxyProtocol,EnableCustomAuthConfigPush"
	NoSupportedConfigList *string `json:"NoSupportedConfigList,omitempty" xml:"NoSupportedConfigList,omitempty"`
	// example:
	//
	// KEEP_UNCHANGED
	PathWithEscapedSlashes *string `json:"PathWithEscapedSlashes,omitempty" xml:"PathWithEscapedSlashes,omitempty"`
	// example:
	//
	// true
	PreserveHeaderFormat *bool                                             `json:"PreserveHeaderFormat,omitempty" xml:"PreserveHeaderFormat,omitempty"`
	SlsConfigDetails     *GetGatewayConfigResponseBodyDataSlsConfigDetails `json:"SlsConfigDetails,omitempty" xml:"SlsConfigDetails,omitempty" type:"Struct"`
	// example:
	//
	// true
	SupportWaf *bool `json:"SupportWaf,omitempty" xml:"SupportWaf,omitempty"`
	// example:
	//
	// 30
	UpstreamIdleTimeout *int32 `json:"UpstreamIdleTimeout,omitempty" xml:"UpstreamIdleTimeout,omitempty"`
	// example:
	//
	// 30
	WebsocketTermGracePeriod *int32 `json:"WebsocketTermGracePeriod,omitempty" xml:"WebsocketTermGracePeriod,omitempty"`
	// example:
	//
	// 0
	XffTrustedNum *int32                                         `json:"XffTrustedNum,omitempty" xml:"XffTrustedNum,omitempty"`
	XtraceDetails *GetGatewayConfigResponseBodyDataXtraceDetails `json:"XtraceDetails,omitempty" xml:"XtraceDetails,omitempty" type:"Struct"`
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

func (s *GetGatewayConfigResponseBodyData) GetLogFilterConfig() *string {
	return s.LogFilterConfig
}

func (s *GetGatewayConfigResponseBodyData) GetNoSupportedConfigList() *string {
	return s.NoSupportedConfigList
}

func (s *GetGatewayConfigResponseBodyData) GetPathWithEscapedSlashes() *string {
	return s.PathWithEscapedSlashes
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

func (s *GetGatewayConfigResponseBodyData) SetLogFilterConfig(v string) *GetGatewayConfigResponseBodyData {
	s.LogFilterConfig = &v
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
	Enable     *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	FilterOpt  *string `json:"FilterOpt,omitempty" xml:"FilterOpt,omitempty"`
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
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
	Enable        *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
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
	// example:
	//
	// 500
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-61f465fa2dd044069e2208c4912*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// ""
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// ""
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 0
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	LogOn *bool `json:"LogOn,omitempty" xml:"LogOn,omitempty"`
	// example:
	//
	// mse_gw_access_log
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// false
	NginxCompatible *bool `json:"NginxCompatible,omitempty" xml:"NginxCompatible,omitempty"`
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
	// example:
	//
	// 500
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-61f465fa2dd044069e2208c4912*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 2024-08-12T09:39:07.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-08-12T09:39:07.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 255
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 100
	Sample *int32 `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// example:
	//
	// 333
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 8888
	ServicePort *string `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// example:
	//
	// true
	TraceOn *bool `json:"TraceOn,omitempty" xml:"TraceOn,omitempty"`
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
