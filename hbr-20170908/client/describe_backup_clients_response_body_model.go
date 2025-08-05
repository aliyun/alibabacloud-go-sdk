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
	// The queried backup clients.
	//
	// example:
	//
	// {\\"Client\\": []}
	Clients []*DescribeBackupClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned HBR clients that meet the specified conditions.
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
	return dara.Validate(s)
}

type DescribeBackupClientsResponseBodyClients struct {
	// Indicates whether the client is installed on an all-in-one PC that integrates hardware and monitoring program. Valid values:
	//
	// 	- true: The client is installed on an all-in-one PC that integrates hardware and monitoring program.
	//
	// 	- false: The client is not installed on an all-in-one PC that integrates hardware and monitoring program.
	//
	// example:
	//
	// false
	Appliance *bool `json:"Appliance,omitempty" xml:"Appliance,omitempty"`
	// This parameter is valid only if the **ClientType*	- parameter is set to **ECS_CLIENT**. This parameter indicates the system architecture where the backup client resides. Valid values:
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
	// 	- **UNPROTECTED**: The backup client is not protected.
	//
	// 	- **PROTECTED**: The backup client is protected.
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
	// The type of the backup client. Valid value: **ECS_CLIENT**, which indicates a client for ECS file backup.
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
	// The time when the backup client was created. The value is a UNIX timestamp. Unit: seconds.
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
	// 	- If the client is used to back up ECS files, this parameter indicates the ID of an ECS instance.
	//
	// 	- If the client is used to back up on-premises files, this parameter indicates the hardware fingerprint that is generated based on the system information.
	//
	// example:
	//
	// i-*********************
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is valid only if the **ClientType*	- parameter is set to **ECS_CLIENT**. This parameter indicates the name of the ECS instance.
	//
	// example:
	//
	// instancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The last heartbeat time of the backup client. The value is a UNIX timestamp. Unit: seconds.
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
	// This parameter is valid only if the **ClientType*	- parameter is set to **ECS_CLIENT**. This parameter indicates the operating system type of the backup client. Valid values:
	//
	// 	- **windows**
	//
	// 	- **linux**
	//
	// example:
	//
	// linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// This parameter is valid only if the **ClientType*	- parameter is set to **ECS_CLIENT**. This parameter indicates the internal IP address of the ECS instance.
	//
	// example:
	//
	// 192.168.1.1
	PrivateIpV4 *string `json:"PrivateIpV4,omitempty" xml:"PrivateIpV4,omitempty"`
	// The configuration information of the backup client.
	Settings *DescribeBackupClientsResponseBodyClientsSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Struct"`
	// The status of the backup client. Valid values:
	//
	// 	- **REGISTERED**: The backup client is registered.
	//
	// 	- **ACTIVATED**: The backup client is activated.
	//
	// 	- **DEACTIVATED**: The backup client fails to be activated.
	//
	// 	- **INSTALLING**: The backup client is being installed.
	//
	// 	- **INSTALL_FAILED**: The backup client fails to be installed.
	//
	// 	- **NOT_INSTALLED**: The backup client is not installed.
	//
	// 	- **UPGRADING**: The backup client is being upgraded.
	//
	// 	- **UPGRADE_FAILED**: The backup client fails to be upgraded.
	//
	// 	- **UNINSTALLING**: The backup client is being uninstalled.
	//
	// 	- **UNINSTALL_FAILED**: The backup client fails to be uninstalled.
	//
	// 	- **STOPPED**: The backup client is out of service.
	//
	// 	- **UNKNOWN**: The backup client is disconnected.
	//
	// example:
	//
	// ACTIVATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	Tags []*DescribeBackupClientsResponseBodyClientsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the backup client was updated. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// This parameter is valid only if the **ClientType*	- parameter is set to **ECS_CLIENT**. This parameter indicates the zone of the backup client.
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
	return dara.Validate(s)
}

type DescribeBackupClientsResponseBodyClientsSettings struct {
	// Indicates whether alerts are generated for partially completed jobs. This parameter is valid only for on-premises file backup and ECS file backup.
	//
	// example:
	//
	// false
	AlertOnPartialComplete *bool `json:"AlertOnPartialComplete,omitempty" xml:"AlertOnPartialComplete,omitempty"`
	// The type of the endpoint on the data plane. Valid values:
	//
	// 	- **PUBLIC**: Internet
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **CLASSIC**: classic network
	//
	// example:
	//
	// VPC
	DataNetworkType *string `json:"DataNetworkType,omitempty" xml:"DataNetworkType,omitempty"`
	// The proxy configuration on the data plane. Valid values:
	//
	// 	- **DISABLE**: The proxy is not used.
	//
	// 	- **USE_CONTROL_PROXY*	- (default): The configuration is the same as that on the control plane.
	//
	// 	- **CUSTOM**: The configuration is customized (HTTP).
	//
	// example:
	//
	// USE_CONTROL_PROXY
	DataProxySetting *string `json:"DataProxySetting,omitempty" xml:"DataProxySetting,omitempty"`
	// The number of CPU cores used by a single backup job. The value 0 indicates that the number is unlimited.
	//
	// example:
	//
	// 1
	MaxCpuCore *string `json:"MaxCpuCore,omitempty" xml:"MaxCpuCore,omitempty"`
	// The maximum memory that can be used by the client. Unit: bytes. Only V2.13.0 and later are supported.
	//
	// example:
	//
	// 0
	MaxMemory *int64 `json:"MaxMemory,omitempty" xml:"MaxMemory,omitempty"`
	// The number of concurrent backup jobs. The value 0 indicates that the number is unlimited.
	//
	// example:
	//
	// 1
	MaxWorker *string `json:"MaxWorker,omitempty" xml:"MaxWorker,omitempty"`
	// The custom host IP address of the proxy server on the data plane.
	//
	// example:
	//
	// 192.168.11.100
	ProxyHost *string `json:"ProxyHost,omitempty" xml:"ProxyHost,omitempty"`
	// The custom password of the proxy server on the data plane.
	//
	// example:
	//
	// ******
	ProxyPassword *string `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	// The custom host port of the proxy server on the data plane.
	//
	// example:
	//
	// 3128
	ProxyPort *int32 `json:"ProxyPort,omitempty" xml:"ProxyPort,omitempty"`
	// The custom username of the proxy server on the data plane.
	//
	// example:
	//
	// user
	ProxyUser *string `json:"ProxyUser,omitempty" xml:"ProxyUser,omitempty"`
	// Indicates whether data on the data plane is transmitted over HTTPS. Valid values:
	//
	// 	- true: Data is transmitted over HTTPS.
	//
	// 	- false: Data is transmitted over HTTP.
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
	// The tag key of the backup vault. Valid values of N: 1 to 20
	//
	// 	- The tag key cannot start with `aliyun` or `acs:`.
	//
	// 	- The tag key cannot contain `http://` or `https://`.
	//
	// 	- The tag key cannot be an empty string.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the backup vault. Valid values of N: 1 to 20
	//
	// 	- The tag value cannot start with `aliyun` or `acs:`.
	//
	// 	- The tag value cannot contain `http://` or `https://`.
	//
	// 	- The tag value cannot be an empty string.
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
