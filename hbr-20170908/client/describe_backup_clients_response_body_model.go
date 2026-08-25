// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClients(v []*DescribeBackupClientsResponseBodyClients) *DescribeBackupClientsResponseBody
	GetClients() []*DescribeBackupClientsResponseBodyClients
	SetCode(v string) *DescribeBackupClientsResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeBackupClientsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeBackupClientsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupClientsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackupClientsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupClientsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeBackupClientsResponseBody
	GetTotalCount() *int64
}

type DescribeBackupClientsResponseBody struct {
	// The list of backup clients.
	//
	// example:
	//
	// {\\"Client\\": []}
	Clients []*DescribeBackupClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// The response code. 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. The value "successful" is returned for successful requests. An error message is returned for failed requests.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Minimum value: 1. Maximum value: 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of backup clients that meet the specified conditions.
	//
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponseBody) GetClients() []*DescribeBackupClientsResponseBodyClients {
	return s.Clients
}

func (s *DescribeBackupClientsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeBackupClientsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackupClientsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupClientsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupClientsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupClientsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeBackupClientsResponseBody) SetClients(v []*DescribeBackupClientsResponseBodyClients) *DescribeBackupClientsResponseBody {
	s.Clients = v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetCode(v string) *DescribeBackupClientsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetMessage(v string) *DescribeBackupClientsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetPageNumber(v int32) *DescribeBackupClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetPageSize(v int32) *DescribeBackupClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetRequestId(v string) *DescribeBackupClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetSuccess(v bool) *DescribeBackupClientsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetTotalCount(v int64) *DescribeBackupClientsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) Validate() error {
	if s.Clients != nil {
		for _, item := range s.Clients {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupClientsResponseBodyClients struct {
	// Indicates whether the client is a hardware monitoring appliance client.
	//
	// - true: The client is a hardware monitoring appliance client.
	//
	// - false: The client is not a hardware monitoring appliance client.
	//
	// example:
	//
	// false
	Appliance *bool `json:"Appliance,omitempty" xml:"Appliance,omitempty"`
	// This parameter is valid only when **ClientType*	- is set to **ECS_CLIENT**. The system architecture of the backup client. Valid values:
	//
	// 	- **amd64**
	//
	// 	- **386**
	//
	// example:
	//
	// amd64
	ArchType *string `json:"ArchType,omitempty" xml:"ArchType,omitempty"`
	// The protection status of the backup client. Valid values:
	//
	// 	- **UNPROTECTED**: The server is not protected.
	//
	// 	- **PROTECTED**: The server is protected.
	//
	// example:
	//
	// PROTECTED
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The ID of the backup client.
	//
	// example:
	//
	// c-*********************
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The type of the backup client. The value **ECS_CLIENT*	- indicates an ECS File Backup client.
	//
	// example:
	//
	// ECS_CLIENT
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The version number of the backup client.
	//
	// example:
	//
	// 2.4.5
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The time when the backup client was created. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The hostname of the backup client.
	//
	// example:
	//
	// hostname
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The instance ID.
	//
	// - If the client type is ECS File Backup client, this parameter indicates the ECS instance ID.
	//
	// - If the client type is local file backup client, this parameter indicates the hardware fingerprint generated based on system information.
	//
	// example:
	//
	// i-*********************
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is valid only when **ClientType*	- is set to **ECS_CLIENT**. The name of the ECS instance.
	//
	// example:
	//
	// instancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The time of the last heartbeat of the backup client. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	LastHeartBeatTime *int64 `json:"LastHeartBeatTime,omitempty" xml:"LastHeartBeatTime,omitempty"`
	// The latest version number of the backup client.
	//
	// example:
	//
	// 2.4.5
	MaxClientVersion *string `json:"MaxClientVersion,omitempty" xml:"MaxClientVersion,omitempty"`
	// This parameter is valid only when **ClientType*	- is set to **ECS_CLIENT**. The operating system type of the client. Valid values:
	//
	// 	- **windows**
	//
	// 	- **linux**
	//
	// example:
	//
	// linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// This parameter is valid only when **ClientType*	- is set to **ECS_CLIENT**. The internal IP address of the ECS instance.
	//
	// example:
	//
	// 192.168.1.1
	PrivateIpV4 *string `json:"PrivateIpV4,omitempty" xml:"PrivateIpV4,omitempty"`
	// The configuration information of the backup client.
	Settings *DescribeBackupClientsResponseBodyClientsSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Struct"`
	// The status of the backup client. Valid values:
	//
	// 	- **REGISTERED**: The client is registered.
	//
	// 	- **ACTIVATED**: The client is activated.
	//
	// 	- **DEACTIVATED**: The client activation has expired.
	//
	// 	- **INSTALLING**: The client is being installed.
	//
	// 	- **INSTALL_FAILED**: The client installation failed.
	//
	// 	- **NOT_INSTALLED**: The client is not installed.
	//
	// 	- **UPGRADING**: The client is being upgraded.
	//
	// 	- **UPGRADE_FAILED**: The client upgrade failed.
	//
	// 	- **UNINSTALLING**: The client is being uninstalled.
	//
	// 	- **UNINSTALL_FAILED**: The client uninstallation failed.
	//
	// 	- **STOPPED**: The client service is stopped.
	//
	// 	- **UNKNOWN**: The client is disconnected.
	//
	// example:
	//
	// ACTIVATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	Tags []*DescribeBackupClientsResponseBodyClientsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the backup client was last updated. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// This parameter is valid only when **ClientType*	- is set to **ECS_CLIENT**. The zone ID.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeBackupClientsResponseBodyClients) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponseBodyClients) GetAppliance() *bool {
	return s.Appliance
}

func (s *DescribeBackupClientsResponseBodyClients) GetArchType() *string {
	return s.ArchType
}

func (s *DescribeBackupClientsResponseBodyClients) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupClientsResponseBodyClients) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeBackupClientsResponseBodyClients) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeBackupClientsResponseBodyClients) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeBackupClientsResponseBodyClients) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeBackupClientsResponseBodyClients) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeBackupClientsResponseBodyClients) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackupClientsResponseBodyClients) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeBackupClientsResponseBodyClients) GetLastHeartBeatTime() *int64 {
	return s.LastHeartBeatTime
}

func (s *DescribeBackupClientsResponseBodyClients) GetMaxClientVersion() *string {
	return s.MaxClientVersion
}

func (s *DescribeBackupClientsResponseBodyClients) GetOsType() *string {
	return s.OsType
}

func (s *DescribeBackupClientsResponseBodyClients) GetPrivateIpV4() *string {
	return s.PrivateIpV4
}

func (s *DescribeBackupClientsResponseBodyClients) GetSettings() *DescribeBackupClientsResponseBodyClientsSettings {
	return s.Settings
}

func (s *DescribeBackupClientsResponseBodyClients) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupClientsResponseBodyClients) GetTags() []*DescribeBackupClientsResponseBodyClientsTags {
	return s.Tags
}

func (s *DescribeBackupClientsResponseBodyClients) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeBackupClientsResponseBodyClients) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeBackupClientsResponseBodyClients) SetAppliance(v bool) *DescribeBackupClientsResponseBodyClients {
	s.Appliance = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetArchType(v string) *DescribeBackupClientsResponseBodyClients {
	s.ArchType = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetBackupStatus(v string) *DescribeBackupClientsResponseBodyClients {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetClientId(v string) *DescribeBackupClientsResponseBodyClients {
	s.ClientId = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetClientType(v string) *DescribeBackupClientsResponseBodyClients {
	s.ClientType = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetClientVersion(v string) *DescribeBackupClientsResponseBodyClients {
	s.ClientVersion = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetCreatedTime(v int64) *DescribeBackupClientsResponseBodyClients {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetHostname(v string) *DescribeBackupClientsResponseBodyClients {
	s.Hostname = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetInstanceId(v string) *DescribeBackupClientsResponseBodyClients {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetInstanceName(v string) *DescribeBackupClientsResponseBodyClients {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetLastHeartBeatTime(v int64) *DescribeBackupClientsResponseBodyClients {
	s.LastHeartBeatTime = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetMaxClientVersion(v string) *DescribeBackupClientsResponseBodyClients {
	s.MaxClientVersion = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetOsType(v string) *DescribeBackupClientsResponseBodyClients {
	s.OsType = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetPrivateIpV4(v string) *DescribeBackupClientsResponseBodyClients {
	s.PrivateIpV4 = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetSettings(v *DescribeBackupClientsResponseBodyClientsSettings) *DescribeBackupClientsResponseBodyClients {
	s.Settings = v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetStatus(v string) *DescribeBackupClientsResponseBodyClients {
	s.Status = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetTags(v []*DescribeBackupClientsResponseBodyClientsTags) *DescribeBackupClientsResponseBodyClients {
	s.Tags = v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetUpdatedTime(v int64) *DescribeBackupClientsResponseBodyClients {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetZoneId(v string) *DescribeBackupClientsResponseBodyClients {
	s.ZoneId = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) Validate() error {
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupClientsResponseBodyClientsSettings struct {
	// Specifies whether to send alerts for partially completed jobs. This setting applies only to File Backup and ECS File Backup Essential Edition.
	//
	// example:
	//
	// false
	AlertOnPartialComplete *bool `json:"AlertOnPartialComplete,omitempty" xml:"AlertOnPartialComplete,omitempty"`
	// The type of the data plane endpoint. Valid values:
	//
	// - **PUBLIC**: public network.
	//
	// - **VPC**: VPC network.
	//
	// - **CLASSIC**: classic network.
	//
	// example:
	//
	// VPC
	DataNetworkType *string `json:"DataNetworkType,omitempty" xml:"DataNetworkType,omitempty"`
	// The data plane proxy setting. Valid values:
	//
	// - **DISABLE**: No proxy is used.
	//
	// - **USE_CONTROL_PROXY*	- (default): The same configuration as the control plane is used.
	//
	// - **CUSTOM**: A custom configuration is used (HTTP protocol).
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
	MaxCpuCore *string `json:"MaxCpuCore,omitempty" xml:"MaxCpuCore,omitempty"`
	// The maximum memory that the client can use. Unit: bytes. Only versions 2.13.0 and later are supported.
	//
	// example:
	//
	// 0
	MaxMemory *int64 `json:"MaxMemory,omitempty" xml:"MaxMemory,omitempty"`
	// The number of concurrent workers for a single backup job. A value of 0 indicates no limit.
	//
	// example:
	//
	// 1
	MaxWorker *string `json:"MaxWorker,omitempty" xml:"MaxWorker,omitempty"`
	// The IP address of the custom data plane proxy server.
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
	// The port of the custom data plane proxy server.
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
	// Indicates whether HTTPS is used to transmit data plane data.
	//
	// - true: HTTPS is used for transmission.
	//
	// - false: HTTP is used for transmission.
	//
	// example:
	//
	// false
	UseHttps *string `json:"UseHttps,omitempty" xml:"UseHttps,omitempty"`
}

func (s DescribeBackupClientsResponseBodyClientsSettings) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsResponseBodyClientsSettings) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetAlertOnPartialComplete() *bool {
	return s.AlertOnPartialComplete
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetDataNetworkType() *string {
	return s.DataNetworkType
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetDataProxySetting() *string {
	return s.DataProxySetting
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetMaxCpuCore() *string {
	return s.MaxCpuCore
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetMaxMemory() *int64 {
	return s.MaxMemory
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetMaxWorker() *string {
	return s.MaxWorker
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetProxyHost() *string {
	return s.ProxyHost
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetProxyPassword() *string {
	return s.ProxyPassword
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetProxyPort() *int32 {
	return s.ProxyPort
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetProxyUser() *string {
	return s.ProxyUser
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) GetUseHttps() *string {
	return s.UseHttps
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetAlertOnPartialComplete(v bool) *DescribeBackupClientsResponseBodyClientsSettings {
	s.AlertOnPartialComplete = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetDataNetworkType(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.DataNetworkType = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetDataProxySetting(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.DataProxySetting = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetMaxCpuCore(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.MaxCpuCore = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetMaxMemory(v int64) *DescribeBackupClientsResponseBodyClientsSettings {
	s.MaxMemory = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetMaxWorker(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.MaxWorker = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetProxyHost(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.ProxyHost = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetProxyPassword(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.ProxyPassword = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetProxyPort(v int32) *DescribeBackupClientsResponseBodyClientsSettings {
	s.ProxyPort = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetProxyUser(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.ProxyUser = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetUseHttps(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.UseHttps = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupClientsResponseBodyClientsTags struct {
	// The tag key of the backup vault. Valid values of N: 1 to 20.
	//
	// - The tag key cannot start with `aliyun` or `acs:`.
	//
	// - The tag key cannot contain `http://` or `https://`.
	//
	// - The tag key cannot be an empty string.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the backup vault. Valid values of N: 1 to 20.
	//
	// - The tag value cannot start with `aliyun` or `acs:`.
	//
	// - The tag value cannot contain `http://` or `https://`.
	//
	// - The tag value cannot be an empty string.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBackupClientsResponseBodyClientsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsResponseBodyClientsTags) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponseBodyClientsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeBackupClientsResponseBodyClientsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeBackupClientsResponseBodyClientsTags) SetKey(v string) *DescribeBackupClientsResponseBodyClientsTags {
	s.Key = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsTags) SetValue(v string) *DescribeBackupClientsResponseBodyClientsTags {
	s.Value = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsTags) Validate() error {
	return dara.Validate(s)
}
