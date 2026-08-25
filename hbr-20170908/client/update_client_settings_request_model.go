// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertOnPartialComplete(v bool) *UpdateClientSettingsRequest
	GetAlertOnPartialComplete() *bool
	SetClientId(v string) *UpdateClientSettingsRequest
	GetClientId() *string
	SetDataNetworkType(v string) *UpdateClientSettingsRequest
	GetDataNetworkType() *string
	SetDataProxySetting(v string) *UpdateClientSettingsRequest
	GetDataProxySetting() *string
	SetMaxCpuCore(v int32) *UpdateClientSettingsRequest
	GetMaxCpuCore() *int32
	SetMaxMemory(v int64) *UpdateClientSettingsRequest
	GetMaxMemory() *int64
	SetMaxWorker(v int32) *UpdateClientSettingsRequest
	GetMaxWorker() *int32
	SetProxyHost(v string) *UpdateClientSettingsRequest
	GetProxyHost() *string
	SetProxyPassword(v string) *UpdateClientSettingsRequest
	GetProxyPassword() *string
	SetProxyPort(v int32) *UpdateClientSettingsRequest
	GetProxyPort() *int32
	SetProxyUser(v string) *UpdateClientSettingsRequest
	GetProxyUser() *string
	SetResourceGroupId(v string) *UpdateClientSettingsRequest
	GetResourceGroupId() *string
	SetUseHttps(v bool) *UpdateClientSettingsRequest
	GetUseHttps() *bool
	SetVaultId(v string) *UpdateClientSettingsRequest
	GetVaultId() *string
}

type UpdateClientSettingsRequest struct {
	// Specifies whether to trigger an alert for partially completed jobs. This parameter takes effect only for local File Backup and ECS File Backup Essential Edition.
	//
	// example:
	//
	// false
	AlertOnPartialComplete *bool `json:"AlertOnPartialComplete,omitempty" xml:"AlertOnPartialComplete,omitempty"`
	// The backup client ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-*********************
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The data plane access point type. Valid values:
	//
	// 	- **PUBLIC**: public network
	//
	// 	- **VPC**: VPC network
	//
	// 	- **CLASSIC**: classic network
	//
	// example:
	//
	// VPC
	DataNetworkType *string `json:"DataNetworkType,omitempty" xml:"DataNetworkType,omitempty"`
	// The data plane proxy setting. Valid values:
	//
	// 	- **DISABLE**: does not use a proxy.
	//
	// 	- **USE_CONTROL_PROXY*	- (default): uses the same configuration as the control plane.
	//
	// 	- **CUSTOM**: uses a custom configuration (HTTP protocol).
	//
	// example:
	//
	// USE_CONTROL_PROXY
	DataProxySetting *string `json:"DataProxySetting,omitempty" xml:"DataProxySetting,omitempty"`
	// The number of CPU cores used by a single backup job. A value of 0 indicates no limit.
	//
	// example:
	//
	// 1
	MaxCpuCore *int32 `json:"MaxCpuCore,omitempty" xml:"MaxCpuCore,omitempty"`
	// The maximum memory that the client can use, in bytes. Only version 2.13.0 and later are supported.
	//
	// example:
	//
	// 4096
	MaxMemory *int64 `json:"MaxMemory,omitempty" xml:"MaxMemory,omitempty"`
	// The number of concurrent workers for a single backup job. A value of 0 indicates no limit.
	//
	// example:
	//
	// 1
	MaxWorker *int32 `json:"MaxWorker,omitempty" xml:"MaxWorker,omitempty"`
	// The IP address of the custom data plane proxy server host.
	//
	// example:
	//
	// 192.168.11.100
	ProxyHost *string `json:"ProxyHost,omitempty" xml:"ProxyHost,omitempty"`
	// The password of the custom data plane proxy server.
	//
	// example:
	//
	// ******
	ProxyPassword *string `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	// The port of the custom data plane proxy server host.
	//
	// example:
	//
	// 3128
	ProxyPort *int32 `json:"ProxyPort,omitempty" xml:"ProxyPort,omitempty"`
	// The username of the custom data plane proxy server.
	//
	// example:
	//
	// user
	ProxyUser *string `json:"ProxyUser,omitempty" xml:"ProxyUser,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-*********************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to use HTTPS to transmit data plane data.
	//
	// - true: Uses HTTPS for transmission.
	//
	// - false: Uses HTTP for transmission.
	//
	// example:
	//
	// false
	UseHttps *bool `json:"UseHttps,omitempty" xml:"UseHttps,omitempty"`
	// The backup vault ID. This parameter is required for legacy clients.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateClientSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientSettingsRequest) GetAlertOnPartialComplete() *bool {
	return s.AlertOnPartialComplete
}

func (s *UpdateClientSettingsRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateClientSettingsRequest) GetDataNetworkType() *string {
	return s.DataNetworkType
}

func (s *UpdateClientSettingsRequest) GetDataProxySetting() *string {
	return s.DataProxySetting
}

func (s *UpdateClientSettingsRequest) GetMaxCpuCore() *int32 {
	return s.MaxCpuCore
}

func (s *UpdateClientSettingsRequest) GetMaxMemory() *int64 {
	return s.MaxMemory
}

func (s *UpdateClientSettingsRequest) GetMaxWorker() *int32 {
	return s.MaxWorker
}

func (s *UpdateClientSettingsRequest) GetProxyHost() *string {
	return s.ProxyHost
}

func (s *UpdateClientSettingsRequest) GetProxyPassword() *string {
	return s.ProxyPassword
}

func (s *UpdateClientSettingsRequest) GetProxyPort() *int32 {
	return s.ProxyPort
}

func (s *UpdateClientSettingsRequest) GetProxyUser() *string {
	return s.ProxyUser
}

func (s *UpdateClientSettingsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateClientSettingsRequest) GetUseHttps() *bool {
	return s.UseHttps
}

func (s *UpdateClientSettingsRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *UpdateClientSettingsRequest) SetAlertOnPartialComplete(v bool) *UpdateClientSettingsRequest {
	s.AlertOnPartialComplete = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetClientId(v string) *UpdateClientSettingsRequest {
	s.ClientId = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetDataNetworkType(v string) *UpdateClientSettingsRequest {
	s.DataNetworkType = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetDataProxySetting(v string) *UpdateClientSettingsRequest {
	s.DataProxySetting = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetMaxCpuCore(v int32) *UpdateClientSettingsRequest {
	s.MaxCpuCore = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetMaxMemory(v int64) *UpdateClientSettingsRequest {
	s.MaxMemory = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetMaxWorker(v int32) *UpdateClientSettingsRequest {
	s.MaxWorker = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetProxyHost(v string) *UpdateClientSettingsRequest {
	s.ProxyHost = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetProxyPassword(v string) *UpdateClientSettingsRequest {
	s.ProxyPassword = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetProxyPort(v int32) *UpdateClientSettingsRequest {
	s.ProxyPort = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetProxyUser(v string) *UpdateClientSettingsRequest {
	s.ProxyUser = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetResourceGroupId(v string) *UpdateClientSettingsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetUseHttps(v bool) *UpdateClientSettingsRequest {
	s.UseHttps = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetVaultId(v string) *UpdateClientSettingsRequest {
	s.VaultId = &v
	return s
}

func (s *UpdateClientSettingsRequest) Validate() error {
	return dara.Validate(s)
}
