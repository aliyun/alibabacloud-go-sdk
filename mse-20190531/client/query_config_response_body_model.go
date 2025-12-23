// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryConfigResponseBody
	GetCode() *int32
	SetData(v *QueryConfigResponseBodyData) *QueryConfigResponseBody
	GetData() *QueryConfigResponseBodyData
	SetHttpStatusCode(v int32) *QueryConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryConfigResponseBody
	GetSuccess() *bool
}

type QueryConfigResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8BD1E58D-0755-42AC-A599-E6B55112EC53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryConfigResponseBody) GetData() *QueryConfigResponseBodyData {
	return s.Data
}

func (s *QueryConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryConfigResponseBody) SetCode(v int32) *QueryConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryConfigResponseBody) SetData(v *QueryConfigResponseBodyData) *QueryConfigResponseBody {
	s.Data = v
	return s
}

func (s *QueryConfigResponseBody) SetHttpStatusCode(v int32) *QueryConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryConfigResponseBody) SetMessage(v string) *QueryConfigResponseBody {
	s.Message = &v
	return s
}

func (s *QueryConfigResponseBody) SetRequestId(v string) *QueryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConfigResponseBody) SetSuccess(v bool) *QueryConfigResponseBody {
	s.Success = &v
	return s
}

func (s *QueryConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConfigResponseBodyData struct {
	// Indicates whether Simple Authentication and Security Layer (SASL) forced identity authentication is enabled for the ZooKeeper instance.
	AuthEnabled *bool `json:"AuthEnabled,omitempty" xml:"AuthEnabled,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// none
	AutopurgePurgeInterval *string `json:"AutopurgePurgeInterval,omitempty" xml:"AutopurgePurgeInterval,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// none
	AutopurgeSnapRetainCount *string `json:"AutopurgeSnapRetainCount,omitempty" xml:"AutopurgeSnapRetainCount,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// name
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Indicates whether RAM authentication of a configuration center is enabled. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: enabled.
	//
	// 	- `false`: disabled.
	//
	// example:
	//
	// true
	ConfigAuthEnabled *bool `json:"ConfigAuthEnabled,omitempty" xml:"ConfigAuthEnabled,omitempty"`
	// Indicates whether RAM authentication is supported by a configuration center of the instance. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: supported.
	//
	// 	- `false`: not supported.
	//
	// example:
	//
	// true
	ConfigAuthSupported *bool `json:"ConfigAuthSupported,omitempty" xml:"ConfigAuthSupported,omitempty"`
	// The maximum size of contents in a configuration. Unit: KB.
	//
	// example:
	//
	// 100
	ConfigContentLimit *int64 `json:"ConfigContentLimit,omitempty" xml:"ConfigContentLimit,omitempty"`
	// Indicates whether configuration encryption of a configuration center is enabled by the instance. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: enabled.
	//
	// 	- `false`: disabled.
	//
	// example:
	//
	// true
	ConfigSecretEnabled *bool `json:"ConfigSecretEnabled,omitempty" xml:"ConfigSecretEnabled,omitempty"`
	// Indicates whether configuration encryption of a configuration center is supported by the instance. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: supported.
	//
	// 	- `false`: not supported.
	//
	// example:
	//
	// true
	ConfigSecretSupported *bool `json:"ConfigSecretSupported,omitempty" xml:"ConfigSecretSupported,omitempty"`
	// Indicates whether the Nacos open source console is enabled.
	//
	// example:
	//
	// true
	ConsoleUIEnabled *bool `json:"ConsoleUIEnabled,omitempty" xml:"ConsoleUIEnabled,omitempty"`
	Enable4lw        *bool `json:"Enable4lw,omitempty" xml:"Enable4lw,omitempty"`
	// Indicates whether access port 8761 was enabled for Eureka. If this port is disabled, applications cannot use the Eureka protocol for service registration and discovery.
	//
	// example:
	//
	// true
	EurekaSupported *bool `json:"EurekaSupported,omitempty" xml:"EurekaSupported,omitempty"`
	// Indicates whether the time to live (TTL) configuration is enabled. This parameter is valid for ZooKeeper instances.
	//
	// example:
	//
	// true
	ExtendedTypesEnable *bool `json:"ExtendedTypesEnable,omitempty" xml:"ExtendedTypesEnable,omitempty"`
	// The maximum connection duration of the instance. Unit: seconds. This parameter is valid for ZooKeeper instances.
	//
	// example:
	//
	// 100
	InitLimit *string `json:"InitLimit,omitempty" xml:"InitLimit,omitempty"`
	// The maximum amount of data on each node. This parameter is valid for ZooKeeper instances. Unit: bytes.
	//
	// example:
	//
	// 1
	JuteMaxbuffer *string `json:"JuteMaxbuffer,omitempty" xml:"JuteMaxbuffer,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// none
	JvmFlagsCustom *string `json:"JvmFlagsCustom,omitempty" xml:"JvmFlagsCustom,omitempty"`
	// Indicates whether Mesh Configuration Protocol (MCP) is enabled. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: enabled.
	//
	// 	- `false`: disabled.
	//
	// example:
	//
	// true
	MCPEnabled *bool `json:"MCPEnabled,omitempty" xml:"MCPEnabled,omitempty"`
	// Indicates whether MCP is supported. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: supported.
	//
	// 	- `false`: not supported.
	//
	// example:
	//
	// true
	MCPSupported *bool `json:"MCPSupported,omitempty" xml:"MCPSupported,omitempty"`
	// The number of connections between a client and a server. This parameter is valid for ZooKeeper instances.\\
	//
	// If this parameter is set to 0, no limits are imposed on the number of connections.
	//
	// example:
	//
	// 0
	MaxClientCnxns *string `json:"MaxClientCnxns,omitempty" xml:"MaxClientCnxns,omitempty"`
	// The maximum timeout period. This parameter is valid for ZooKeeper instances.
	//
	// example:
	//
	// -1
	MaxSessionTimeout *string `json:"MaxSessionTimeout,omitempty" xml:"MaxSessionTimeout,omitempty"`
	// The minimum timeout period. This parameter is valid for ZooKeeper instances.
	//
	// example:
	//
	// -1
	MinSessionTimeout *string `json:"MinSessionTimeout,omitempty" xml:"MinSessionTimeout,omitempty"`
	// The runtime configuration of the Nacos instance.
	NacosRunningEnv *QueryConfigResponseBodyDataNacosRunningEnv `json:"NacosRunningEnv,omitempty" xml:"NacosRunningEnv,omitempty" type:"Struct"`
	// Indicates whether RAM authentication of a registry is enabled. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: enabled.
	//
	// 	- `false`: disabled.
	//
	// example:
	//
	// false
	NamingAuthEnabled *bool `json:"NamingAuthEnabled,omitempty" xml:"NamingAuthEnabled,omitempty"`
	// Indicates whether RAM authentication of services is supported by the instance. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: supported.
	//
	// 	- `false`: not supported.
	//
	// example:
	//
	// true
	NamingAuthSupported *bool `json:"NamingAuthSupported,omitempty" xml:"NamingAuthSupported,omitempty"`
	// Indicates whether service creation is supported for the instance. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: supported.
	//
	// 	- `false`: not supported.
	//
	// example:
	//
	// true
	NamingCreateServiceSupported *bool `json:"NamingCreateServiceSupported,omitempty" xml:"NamingCreateServiceSupported,omitempty"`
	// Indicates whether super permissions are enabled. This parameter is valid for ZooKeeper instances. Valid values:
	//
	// 	- `true`: enabled.
	//
	// 	- `false`: disabled.
	//
	// example:
	//
	// true
	OpenSuperAcl *bool `json:"OpenSuperAcl,omitempty" xml:"OpenSuperAcl,omitempty"`
	// The password that corresponds to the username. This parameter is valid only if OpenSuperAcl is set to true.
	//
	// example:
	//
	// password
	PassWord                    *string `json:"PassWord,omitempty" xml:"PassWord,omitempty"`
	PrometheusSdProtocolEnabled *string `json:"PrometheusSdProtocolEnabled,omitempty" xml:"PrometheusSdProtocolEnabled,omitempty"`
	// Indicates whether the instance was restarted and new configurations have taken effect. Valid values:
	//
	// 	- `true`: The restart was successful.
	//
	// 	- `false`: The restart failed.
	//
	// example:
	//
	// true
	RestartFlag *bool `json:"RestartFlag,omitempty" xml:"RestartFlag,omitempty"`
	// The frequency for generating snapshots. This parameter is valid for ZooKeeper instances.
	//
	// example:
	//
	// 100000
	SnapshotCount *string `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	// The connection timeout period of the instance. This parameter is valid for ZooKeeper instances. Unit: seconds.
	//
	// example:
	//
	// 10
	SyncLimit *string `json:"SyncLimit,omitempty" xml:"SyncLimit,omitempty"`
	// MSE Nacos supports TLS transmission link encryption since version 2.1.2.1. Nacos clients must be upgraded to version 2.2.1 or later. After TLS is enabled, the system performance will decrease by about 10%. You must evaluate the system capacity. For more information about the relevant operations, see Nacos TLS transmission encryption.
	//
	// example:
	//
	// true
	TLSEnabled *bool `json:"TLSEnabled,omitempty" xml:"TLSEnabled,omitempty"`
	// The time unit of the engine. This parameter is valid for ZooKeeper instances. Default value: 2000. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	TickTime *string `json:"TickTime,omitempty" xml:"TickTime,omitempty"`
	// The username of the user. This parameter is valid only if OpenSuperAcl is set to true.
	//
	// example:
	//
	// name
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s QueryConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryConfigResponseBodyData) GetAuthEnabled() *bool {
	return s.AuthEnabled
}

func (s *QueryConfigResponseBodyData) GetAutopurgePurgeInterval() *string {
	return s.AutopurgePurgeInterval
}

func (s *QueryConfigResponseBodyData) GetAutopurgeSnapRetainCount() *string {
	return s.AutopurgeSnapRetainCount
}

func (s *QueryConfigResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *QueryConfigResponseBodyData) GetConfigAuthEnabled() *bool {
	return s.ConfigAuthEnabled
}

func (s *QueryConfigResponseBodyData) GetConfigAuthSupported() *bool {
	return s.ConfigAuthSupported
}

func (s *QueryConfigResponseBodyData) GetConfigContentLimit() *int64 {
	return s.ConfigContentLimit
}

func (s *QueryConfigResponseBodyData) GetConfigSecretEnabled() *bool {
	return s.ConfigSecretEnabled
}

func (s *QueryConfigResponseBodyData) GetConfigSecretSupported() *bool {
	return s.ConfigSecretSupported
}

func (s *QueryConfigResponseBodyData) GetConsoleUIEnabled() *bool {
	return s.ConsoleUIEnabled
}

func (s *QueryConfigResponseBodyData) GetEnable4lw() *bool {
	return s.Enable4lw
}

func (s *QueryConfigResponseBodyData) GetEurekaSupported() *bool {
	return s.EurekaSupported
}

func (s *QueryConfigResponseBodyData) GetExtendedTypesEnable() *bool {
	return s.ExtendedTypesEnable
}

func (s *QueryConfigResponseBodyData) GetInitLimit() *string {
	return s.InitLimit
}

func (s *QueryConfigResponseBodyData) GetJuteMaxbuffer() *string {
	return s.JuteMaxbuffer
}

func (s *QueryConfigResponseBodyData) GetJvmFlagsCustom() *string {
	return s.JvmFlagsCustom
}

func (s *QueryConfigResponseBodyData) GetMCPEnabled() *bool {
	return s.MCPEnabled
}

func (s *QueryConfigResponseBodyData) GetMCPSupported() *bool {
	return s.MCPSupported
}

func (s *QueryConfigResponseBodyData) GetMaxClientCnxns() *string {
	return s.MaxClientCnxns
}

func (s *QueryConfigResponseBodyData) GetMaxSessionTimeout() *string {
	return s.MaxSessionTimeout
}

func (s *QueryConfigResponseBodyData) GetMinSessionTimeout() *string {
	return s.MinSessionTimeout
}

func (s *QueryConfigResponseBodyData) GetNacosRunningEnv() *QueryConfigResponseBodyDataNacosRunningEnv {
	return s.NacosRunningEnv
}

func (s *QueryConfigResponseBodyData) GetNamingAuthEnabled() *bool {
	return s.NamingAuthEnabled
}

func (s *QueryConfigResponseBodyData) GetNamingAuthSupported() *bool {
	return s.NamingAuthSupported
}

func (s *QueryConfigResponseBodyData) GetNamingCreateServiceSupported() *bool {
	return s.NamingCreateServiceSupported
}

func (s *QueryConfigResponseBodyData) GetOpenSuperAcl() *bool {
	return s.OpenSuperAcl
}

func (s *QueryConfigResponseBodyData) GetPassWord() *string {
	return s.PassWord
}

func (s *QueryConfigResponseBodyData) GetPrometheusSdProtocolEnabled() *string {
	return s.PrometheusSdProtocolEnabled
}

func (s *QueryConfigResponseBodyData) GetRestartFlag() *bool {
	return s.RestartFlag
}

func (s *QueryConfigResponseBodyData) GetSnapshotCount() *string {
	return s.SnapshotCount
}

func (s *QueryConfigResponseBodyData) GetSyncLimit() *string {
	return s.SyncLimit
}

func (s *QueryConfigResponseBodyData) GetTLSEnabled() *bool {
	return s.TLSEnabled
}

func (s *QueryConfigResponseBodyData) GetTickTime() *string {
	return s.TickTime
}

func (s *QueryConfigResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *QueryConfigResponseBodyData) SetAuthEnabled(v bool) *QueryConfigResponseBodyData {
	s.AuthEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetAutopurgePurgeInterval(v string) *QueryConfigResponseBodyData {
	s.AutopurgePurgeInterval = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetAutopurgeSnapRetainCount(v string) *QueryConfigResponseBodyData {
	s.AutopurgeSnapRetainCount = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetClusterName(v string) *QueryConfigResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetConfigAuthEnabled(v bool) *QueryConfigResponseBodyData {
	s.ConfigAuthEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetConfigAuthSupported(v bool) *QueryConfigResponseBodyData {
	s.ConfigAuthSupported = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetConfigContentLimit(v int64) *QueryConfigResponseBodyData {
	s.ConfigContentLimit = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetConfigSecretEnabled(v bool) *QueryConfigResponseBodyData {
	s.ConfigSecretEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetConfigSecretSupported(v bool) *QueryConfigResponseBodyData {
	s.ConfigSecretSupported = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetConsoleUIEnabled(v bool) *QueryConfigResponseBodyData {
	s.ConsoleUIEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetEnable4lw(v bool) *QueryConfigResponseBodyData {
	s.Enable4lw = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetEurekaSupported(v bool) *QueryConfigResponseBodyData {
	s.EurekaSupported = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetExtendedTypesEnable(v bool) *QueryConfigResponseBodyData {
	s.ExtendedTypesEnable = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetInitLimit(v string) *QueryConfigResponseBodyData {
	s.InitLimit = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetJuteMaxbuffer(v string) *QueryConfigResponseBodyData {
	s.JuteMaxbuffer = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetJvmFlagsCustom(v string) *QueryConfigResponseBodyData {
	s.JvmFlagsCustom = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetMCPEnabled(v bool) *QueryConfigResponseBodyData {
	s.MCPEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetMCPSupported(v bool) *QueryConfigResponseBodyData {
	s.MCPSupported = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetMaxClientCnxns(v string) *QueryConfigResponseBodyData {
	s.MaxClientCnxns = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetMaxSessionTimeout(v string) *QueryConfigResponseBodyData {
	s.MaxSessionTimeout = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetMinSessionTimeout(v string) *QueryConfigResponseBodyData {
	s.MinSessionTimeout = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetNacosRunningEnv(v *QueryConfigResponseBodyDataNacosRunningEnv) *QueryConfigResponseBodyData {
	s.NacosRunningEnv = v
	return s
}

func (s *QueryConfigResponseBodyData) SetNamingAuthEnabled(v bool) *QueryConfigResponseBodyData {
	s.NamingAuthEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetNamingAuthSupported(v bool) *QueryConfigResponseBodyData {
	s.NamingAuthSupported = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetNamingCreateServiceSupported(v bool) *QueryConfigResponseBodyData {
	s.NamingCreateServiceSupported = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetOpenSuperAcl(v bool) *QueryConfigResponseBodyData {
	s.OpenSuperAcl = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetPassWord(v string) *QueryConfigResponseBodyData {
	s.PassWord = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetPrometheusSdProtocolEnabled(v string) *QueryConfigResponseBodyData {
	s.PrometheusSdProtocolEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetRestartFlag(v bool) *QueryConfigResponseBodyData {
	s.RestartFlag = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetSnapshotCount(v string) *QueryConfigResponseBodyData {
	s.SnapshotCount = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetSyncLimit(v string) *QueryConfigResponseBodyData {
	s.SyncLimit = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetTLSEnabled(v bool) *QueryConfigResponseBodyData {
	s.TLSEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetTickTime(v string) *QueryConfigResponseBodyData {
	s.TickTime = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetUserName(v string) *QueryConfigResponseBodyData {
	s.UserName = &v
	return s
}

func (s *QueryConfigResponseBodyData) Validate() error {
	if s.NacosRunningEnv != nil {
		if err := s.NacosRunningEnv.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConfigResponseBodyDataNacosRunningEnv struct {
	// Indicates whether empty list protection is enabled.
	//
	// example:
	//
	// true
	EmptyProtect *bool                                                  `json:"emptyProtect,omitempty" xml:"emptyProtect,omitempty"`
	FenceEnabled *bool                                                  `json:"fenceEnabled,omitempty" xml:"fenceEnabled,omitempty"`
	FencePolicy  *QueryConfigResponseBodyDataNacosRunningEnvFencePolicy `json:"fencePolicy,omitempty" xml:"fencePolicy,omitempty" type:"Struct"`
	GrayAuth     *string                                                `json:"grayAuth,omitempty" xml:"grayAuth,omitempty"`
}

func (s QueryConfigResponseBodyDataNacosRunningEnv) String() string {
	return dara.Prettify(s)
}

func (s QueryConfigResponseBodyDataNacosRunningEnv) GoString() string {
	return s.String()
}

func (s *QueryConfigResponseBodyDataNacosRunningEnv) GetEmptyProtect() *bool {
	return s.EmptyProtect
}

func (s *QueryConfigResponseBodyDataNacosRunningEnv) GetFenceEnabled() *bool {
	return s.FenceEnabled
}

func (s *QueryConfigResponseBodyDataNacosRunningEnv) GetFencePolicy() *QueryConfigResponseBodyDataNacosRunningEnvFencePolicy {
	return s.FencePolicy
}

func (s *QueryConfigResponseBodyDataNacosRunningEnv) GetGrayAuth() *string {
	return s.GrayAuth
}

func (s *QueryConfigResponseBodyDataNacosRunningEnv) SetEmptyProtect(v bool) *QueryConfigResponseBodyDataNacosRunningEnv {
	s.EmptyProtect = &v
	return s
}

func (s *QueryConfigResponseBodyDataNacosRunningEnv) SetFenceEnabled(v bool) *QueryConfigResponseBodyDataNacosRunningEnv {
	s.FenceEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyDataNacosRunningEnv) SetFencePolicy(v *QueryConfigResponseBodyDataNacosRunningEnvFencePolicy) *QueryConfigResponseBodyDataNacosRunningEnv {
	s.FencePolicy = v
	return s
}

func (s *QueryConfigResponseBodyDataNacosRunningEnv) SetGrayAuth(v string) *QueryConfigResponseBodyDataNacosRunningEnv {
	s.GrayAuth = &v
	return s
}

func (s *QueryConfigResponseBodyDataNacosRunningEnv) Validate() error {
	if s.FencePolicy != nil {
		if err := s.FencePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConfigResponseBodyDataNacosRunningEnvFencePolicy struct {
	EnabledModules  []*string          `json:"enabledModules,omitempty" xml:"enabledModules,omitempty" type:"Repeated"`
	InterceptPolicy map[string]*string `json:"interceptPolicy,omitempty" xml:"interceptPolicy,omitempty"`
}

func (s QueryConfigResponseBodyDataNacosRunningEnvFencePolicy) String() string {
	return dara.Prettify(s)
}

func (s QueryConfigResponseBodyDataNacosRunningEnvFencePolicy) GoString() string {
	return s.String()
}

func (s *QueryConfigResponseBodyDataNacosRunningEnvFencePolicy) GetEnabledModules() []*string {
	return s.EnabledModules
}

func (s *QueryConfigResponseBodyDataNacosRunningEnvFencePolicy) GetInterceptPolicy() map[string]*string {
	return s.InterceptPolicy
}

func (s *QueryConfigResponseBodyDataNacosRunningEnvFencePolicy) SetEnabledModules(v []*string) *QueryConfigResponseBodyDataNacosRunningEnvFencePolicy {
	s.EnabledModules = v
	return s
}

func (s *QueryConfigResponseBodyDataNacosRunningEnvFencePolicy) SetInterceptPolicy(v map[string]*string) *QueryConfigResponseBodyDataNacosRunningEnvFencePolicy {
	s.InterceptPolicy = v
	return s
}

func (s *QueryConfigResponseBodyDataNacosRunningEnvFencePolicy) Validate() error {
	return dara.Validate(s)
}
