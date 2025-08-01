// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateConfigRequest
	GetAcceptLanguage() *string
	SetAuthEnabled(v bool) *UpdateConfigRequest
	GetAuthEnabled() *bool
	SetAutopurgePurgeInterval(v string) *UpdateConfigRequest
	GetAutopurgePurgeInterval() *string
	SetAutopurgeSnapRetainCount(v string) *UpdateConfigRequest
	GetAutopurgeSnapRetainCount() *string
	SetClusterId(v string) *UpdateConfigRequest
	GetClusterId() *string
	SetConfigAuthEnabled(v bool) *UpdateConfigRequest
	GetConfigAuthEnabled() *bool
	SetConfigSecretEnabled(v bool) *UpdateConfigRequest
	GetConfigSecretEnabled() *bool
	SetConfigType(v string) *UpdateConfigRequest
	GetConfigType() *string
	SetConsoleUIEnabled(v bool) *UpdateConfigRequest
	GetConsoleUIEnabled() *bool
	SetEnable4lw(v bool) *UpdateConfigRequest
	GetEnable4lw() *bool
	SetEurekaSupported(v bool) *UpdateConfigRequest
	GetEurekaSupported() *bool
	SetExtendedTypesEnable(v string) *UpdateConfigRequest
	GetExtendedTypesEnable() *string
	SetInitLimit(v string) *UpdateConfigRequest
	GetInitLimit() *string
	SetInstanceId(v string) *UpdateConfigRequest
	GetInstanceId() *string
	SetJuteMaxbuffer(v string) *UpdateConfigRequest
	GetJuteMaxbuffer() *string
	SetMCPEnabled(v bool) *UpdateConfigRequest
	GetMCPEnabled() *bool
	SetMaxClientCnxns(v string) *UpdateConfigRequest
	GetMaxClientCnxns() *string
	SetMaxSessionTimeout(v string) *UpdateConfigRequest
	GetMaxSessionTimeout() *string
	SetMinSessionTimeout(v string) *UpdateConfigRequest
	GetMinSessionTimeout() *string
	SetNamingAuthEnabled(v bool) *UpdateConfigRequest
	GetNamingAuthEnabled() *bool
	SetOpenSuperAcl(v string) *UpdateConfigRequest
	GetOpenSuperAcl() *string
	SetPassWord(v string) *UpdateConfigRequest
	GetPassWord() *string
	SetRequestPars(v string) *UpdateConfigRequest
	GetRequestPars() *string
	SetSnapshotCount(v string) *UpdateConfigRequest
	GetSnapshotCount() *string
	SetSyncLimit(v string) *UpdateConfigRequest
	GetSyncLimit() *string
	SetTLSEnabled(v bool) *UpdateConfigRequest
	GetTLSEnabled() *bool
	SetTickTime(v string) *UpdateConfigRequest
	GetTickTime() *string
	SetUserName(v string) *UpdateConfigRequest
	GetUserName() *string
}

type UpdateConfigRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AuthEnabled    *bool   `json:"AuthEnabled,omitempty" xml:"AuthEnabled,omitempty"`
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
	// The ID of the cluster.
	//
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable Resource Access Management (RAM) authentication for a configuration center. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: enabled.
	//
	// 	- `false`: disabled.
	//
	// > Before you configure this parameter, you must call the QueryConfig operation to obtain the ConfigAuthSupported parameter value to check whether the instance supports the RAM authentication feature.
	//
	// example:
	//
	// true
	ConfigAuthEnabled *bool `json:"ConfigAuthEnabled,omitempty" xml:"ConfigAuthEnabled,omitempty"`
	// Specifies whether to enable configuration encryption for a configuration center. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: enabled.
	//
	// 	- `false`: disabled.
	//
	// > Before you configure this parameter, you must call the QueryConfig operation to obtain the ConfigSecretSupported parameter value to check whether the instance supports configuration encryption.
	//
	// example:
	//
	// true
	ConfigSecretEnabled *bool `json:"ConfigSecretEnabled,omitempty" xml:"ConfigSecretEnabled,omitempty"`
	// The format of the configuration. Supported formats include TEXT, JSON, XML, and HTML.
	//
	// example:
	//
	// TEXT
	ConfigType       *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	ConsoleUIEnabled *bool   `json:"ConsoleUIEnabled,omitempty" xml:"ConsoleUIEnabled,omitempty"`
	Enable4lw        *bool   `json:"Enable4lw,omitempty" xml:"Enable4lw,omitempty"`
	// example:
	//
	// true
	EurekaSupported *bool `json:"EurekaSupported,omitempty" xml:"EurekaSupported,omitempty"`
	// Specifies whether to enable the time to live (TTL) configuration. This parameter is valid for ZooKeeper instances.
	//
	// example:
	//
	// true
	ExtendedTypesEnable *string `json:"ExtendedTypesEnable,omitempty" xml:"ExtendedTypesEnable,omitempty"`
	// The maximum connection duration of the instance. This parameter is valid for ZooKeeper instances. Unit: seconds.
	//
	// example:
	//
	// 100
	InitLimit *string `json:"InitLimit,omitempty" xml:"InitLimit,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse_prepaid_public_cn-st2212****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum amount of data on each node. This parameter is valid for ZooKeeper instances. The default maximum data amount on each node is 1 megabyte. Unit: bytes.
	//
	// example:
	//
	// 1048575
	JuteMaxbuffer *string `json:"JuteMaxbuffer,omitempty" xml:"JuteMaxbuffer,omitempty"`
	// Specifies whether to enable Mesh Configuration Protocol (MCP). This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: enabled.
	//
	// 	- `false`: disabled.
	//
	// > Before you configure this parameter, you must call the QueryConfig operation to obtain the MCPSupported parameter value to check whether the instance supports MCP.
	//
	// example:
	//
	// true
	MCPEnabled *bool `json:"MCPEnabled,omitempty" xml:"MCPEnabled,omitempty"`
	// The number of connections between a client and a server. This parameter is valid for ZooKeeper instances.\\
	//
	// If this parameter is set to 0, no limits are imposed on the number of connections.
	//
	// example:
	//
	// 0
	MaxClientCnxns *string `json:"MaxClientCnxns,omitempty" xml:"MaxClientCnxns,omitempty"`
	// The maximum timeout period. This parameter is valid for ZooKeeper instances. Unit: seconds.
	//
	// example:
	//
	// 1000
	MaxSessionTimeout *string `json:"MaxSessionTimeout,omitempty" xml:"MaxSessionTimeout,omitempty"`
	// The minimum timeout period. This parameter is valid for ZooKeeper instances. Unit: seconds.
	//
	// example:
	//
	// 10
	MinSessionTimeout *string `json:"MinSessionTimeout,omitempty" xml:"MinSessionTimeout,omitempty"`
	// Specifies whether to enable RAM authentication for a registry. This parameter is valid for Nacos instances. Valid values:
	//
	// 	- `true`: enabled.
	//
	// 	- `false`: disabled.
	//
	// > Before you configure this parameter, you must call the QueryConfig operation to obtain the NamingAuthSupporte parameter value to check whether the instance supports the RAM authentication feature.
	//
	// example:
	//
	// true
	NamingAuthEnabled *bool `json:"NamingAuthEnabled,omitempty" xml:"NamingAuthEnabled,omitempty"`
	// Specifies whether to enable super permissions. This parameter is valid for ZooKeeper instances. Valid values:
	//
	// 	- `true`: enabled
	//
	// 	- `false`: disabled
	//
	// example:
	//
	// true
	OpenSuperAcl *string `json:"OpenSuperAcl,omitempty" xml:"OpenSuperAcl,omitempty"`
	// The password that corresponds to the username.
	//
	// > You must specify this parameter if OpenSuperAcl is set to true.
	//
	// example:
	//
	// password
	PassWord *string `json:"PassWord,omitempty" xml:"PassWord,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
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
	// example:
	//
	// true
	TLSEnabled *bool `json:"TLSEnabled,omitempty" xml:"TLSEnabled,omitempty"`
	// The time unit. This parameter is valid for ZooKeeper instances. Default value: 2000. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	TickTime *string `json:"TickTime,omitempty" xml:"TickTime,omitempty"`
	// The name of the user.
	//
	// > You must specify this parameter if OpenSuperAcl is set to true.
	//
	// example:
	//
	// name
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateConfigRequest) GetAuthEnabled() *bool {
	return s.AuthEnabled
}

func (s *UpdateConfigRequest) GetAutopurgePurgeInterval() *string {
	return s.AutopurgePurgeInterval
}

func (s *UpdateConfigRequest) GetAutopurgeSnapRetainCount() *string {
	return s.AutopurgeSnapRetainCount
}

func (s *UpdateConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateConfigRequest) GetConfigAuthEnabled() *bool {
	return s.ConfigAuthEnabled
}

func (s *UpdateConfigRequest) GetConfigSecretEnabled() *bool {
	return s.ConfigSecretEnabled
}

func (s *UpdateConfigRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *UpdateConfigRequest) GetConsoleUIEnabled() *bool {
	return s.ConsoleUIEnabled
}

func (s *UpdateConfigRequest) GetEnable4lw() *bool {
	return s.Enable4lw
}

func (s *UpdateConfigRequest) GetEurekaSupported() *bool {
	return s.EurekaSupported
}

func (s *UpdateConfigRequest) GetExtendedTypesEnable() *string {
	return s.ExtendedTypesEnable
}

func (s *UpdateConfigRequest) GetInitLimit() *string {
	return s.InitLimit
}

func (s *UpdateConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateConfigRequest) GetJuteMaxbuffer() *string {
	return s.JuteMaxbuffer
}

func (s *UpdateConfigRequest) GetMCPEnabled() *bool {
	return s.MCPEnabled
}

func (s *UpdateConfigRequest) GetMaxClientCnxns() *string {
	return s.MaxClientCnxns
}

func (s *UpdateConfigRequest) GetMaxSessionTimeout() *string {
	return s.MaxSessionTimeout
}

func (s *UpdateConfigRequest) GetMinSessionTimeout() *string {
	return s.MinSessionTimeout
}

func (s *UpdateConfigRequest) GetNamingAuthEnabled() *bool {
	return s.NamingAuthEnabled
}

func (s *UpdateConfigRequest) GetOpenSuperAcl() *string {
	return s.OpenSuperAcl
}

func (s *UpdateConfigRequest) GetPassWord() *string {
	return s.PassWord
}

func (s *UpdateConfigRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *UpdateConfigRequest) GetSnapshotCount() *string {
	return s.SnapshotCount
}

func (s *UpdateConfigRequest) GetSyncLimit() *string {
	return s.SyncLimit
}

func (s *UpdateConfigRequest) GetTLSEnabled() *bool {
	return s.TLSEnabled
}

func (s *UpdateConfigRequest) GetTickTime() *string {
	return s.TickTime
}

func (s *UpdateConfigRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateConfigRequest) SetAcceptLanguage(v string) *UpdateConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateConfigRequest) SetAuthEnabled(v bool) *UpdateConfigRequest {
	s.AuthEnabled = &v
	return s
}

func (s *UpdateConfigRequest) SetAutopurgePurgeInterval(v string) *UpdateConfigRequest {
	s.AutopurgePurgeInterval = &v
	return s
}

func (s *UpdateConfigRequest) SetAutopurgeSnapRetainCount(v string) *UpdateConfigRequest {
	s.AutopurgeSnapRetainCount = &v
	return s
}

func (s *UpdateConfigRequest) SetClusterId(v string) *UpdateConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateConfigRequest) SetConfigAuthEnabled(v bool) *UpdateConfigRequest {
	s.ConfigAuthEnabled = &v
	return s
}

func (s *UpdateConfigRequest) SetConfigSecretEnabled(v bool) *UpdateConfigRequest {
	s.ConfigSecretEnabled = &v
	return s
}

func (s *UpdateConfigRequest) SetConfigType(v string) *UpdateConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *UpdateConfigRequest) SetConsoleUIEnabled(v bool) *UpdateConfigRequest {
	s.ConsoleUIEnabled = &v
	return s
}

func (s *UpdateConfigRequest) SetEnable4lw(v bool) *UpdateConfigRequest {
	s.Enable4lw = &v
	return s
}

func (s *UpdateConfigRequest) SetEurekaSupported(v bool) *UpdateConfigRequest {
	s.EurekaSupported = &v
	return s
}

func (s *UpdateConfigRequest) SetExtendedTypesEnable(v string) *UpdateConfigRequest {
	s.ExtendedTypesEnable = &v
	return s
}

func (s *UpdateConfigRequest) SetInitLimit(v string) *UpdateConfigRequest {
	s.InitLimit = &v
	return s
}

func (s *UpdateConfigRequest) SetInstanceId(v string) *UpdateConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConfigRequest) SetJuteMaxbuffer(v string) *UpdateConfigRequest {
	s.JuteMaxbuffer = &v
	return s
}

func (s *UpdateConfigRequest) SetMCPEnabled(v bool) *UpdateConfigRequest {
	s.MCPEnabled = &v
	return s
}

func (s *UpdateConfigRequest) SetMaxClientCnxns(v string) *UpdateConfigRequest {
	s.MaxClientCnxns = &v
	return s
}

func (s *UpdateConfigRequest) SetMaxSessionTimeout(v string) *UpdateConfigRequest {
	s.MaxSessionTimeout = &v
	return s
}

func (s *UpdateConfigRequest) SetMinSessionTimeout(v string) *UpdateConfigRequest {
	s.MinSessionTimeout = &v
	return s
}

func (s *UpdateConfigRequest) SetNamingAuthEnabled(v bool) *UpdateConfigRequest {
	s.NamingAuthEnabled = &v
	return s
}

func (s *UpdateConfigRequest) SetOpenSuperAcl(v string) *UpdateConfigRequest {
	s.OpenSuperAcl = &v
	return s
}

func (s *UpdateConfigRequest) SetPassWord(v string) *UpdateConfigRequest {
	s.PassWord = &v
	return s
}

func (s *UpdateConfigRequest) SetRequestPars(v string) *UpdateConfigRequest {
	s.RequestPars = &v
	return s
}

func (s *UpdateConfigRequest) SetSnapshotCount(v string) *UpdateConfigRequest {
	s.SnapshotCount = &v
	return s
}

func (s *UpdateConfigRequest) SetSyncLimit(v string) *UpdateConfigRequest {
	s.SyncLimit = &v
	return s
}

func (s *UpdateConfigRequest) SetTLSEnabled(v bool) *UpdateConfigRequest {
	s.TLSEnabled = &v
	return s
}

func (s *UpdateConfigRequest) SetTickTime(v string) *UpdateConfigRequest {
	s.TickTime = &v
	return s
}

func (s *UpdateConfigRequest) SetUserName(v string) *UpdateConfigRequest {
	s.UserName = &v
	return s
}

func (s *UpdateConfigRequest) Validate() error {
	return dara.Validate(s)
}
