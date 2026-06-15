// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudAssistantSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentUpgradeConfig(v *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) *ModifyCloudAssistantSettingsRequest
	GetAgentUpgradeConfig() *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig
	SetOssDeliveryConfig(v *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) *ModifyCloudAssistantSettingsRequest
	GetOssDeliveryConfig() *ModifyCloudAssistantSettingsRequestOssDeliveryConfig
	SetOwnerAccount(v string) *ModifyCloudAssistantSettingsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCloudAssistantSettingsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCloudAssistantSettingsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCloudAssistantSettingsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCloudAssistantSettingsRequest
	GetResourceOwnerId() *int64
	SetResourceUsageConfig(v *ModifyCloudAssistantSettingsRequestResourceUsageConfig) *ModifyCloudAssistantSettingsRequest
	GetResourceUsageConfig() *ModifyCloudAssistantSettingsRequestResourceUsageConfig
	SetSessionManagerConfig(v *ModifyCloudAssistantSettingsRequestSessionManagerConfig) *ModifyCloudAssistantSettingsRequest
	GetSessionManagerConfig() *ModifyCloudAssistantSettingsRequestSessionManagerConfig
	SetSettingType(v string) *ModifyCloudAssistantSettingsRequest
	GetSettingType() *string
	SetSlsDeliveryConfig(v *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) *ModifyCloudAssistantSettingsRequest
	GetSlsDeliveryConfig() *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig
}

type ModifyCloudAssistantSettingsRequest struct {
	// The configurations of upgrading the Cloud Assistant agent.
	AgentUpgradeConfig *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig `json:"AgentUpgradeConfig,omitempty" xml:"AgentUpgradeConfig,omitempty" type:"Struct"`
	// The configurations of delivering records to OSS.
	OssDeliveryConfig *ModifyCloudAssistantSettingsRequestOssDeliveryConfig `json:"OssDeliveryConfig,omitempty" xml:"OssDeliveryConfig,omitempty" type:"Struct"`
	OwnerAccount      *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The configurations of resource usage for Cloud Assistant. This setting takes effect only when the version of the Cloud Assistant agent is not earlier than the following versions:
	//
	// - Windows: 2.1.4.1065
	//
	// - Linux: 2.2.4.1065
	ResourceUsageConfig *ModifyCloudAssistantSettingsRequestResourceUsageConfig `json:"ResourceUsageConfig,omitempty" xml:"ResourceUsageConfig,omitempty" type:"Struct"`
	// The configurations of the Session Manager feature.
	SessionManagerConfig *ModifyCloudAssistantSettingsRequestSessionManagerConfig `json:"SessionManagerConfig,omitempty" xml:"SessionManagerConfig,omitempty" type:"Struct"`
	// The type of the service configurations. Valid values:
	//
	// - `SessionManagerDelivery`: the configurations of delivering session records.
	//
	// - `InvocationDelivery`: the configurations of delivering command execution records.
	//
	// - `AgentUpgradeConfig`: the configurations of upgrading the Cloud Assistant agent.
	//
	// - `SessionManagerConfig`: the configurations of Cloud Assistant Session Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// SessionManagerDelivery
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
	// The configurations of delivering records to SLS.
	SlsDeliveryConfig *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig `json:"SlsDeliveryConfig,omitempty" xml:"SlsDeliveryConfig,omitempty" type:"Struct"`
}

func (s ModifyCloudAssistantSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequest) GetAgentUpgradeConfig() *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig {
	return s.AgentUpgradeConfig
}

func (s *ModifyCloudAssistantSettingsRequest) GetOssDeliveryConfig() *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	return s.OssDeliveryConfig
}

func (s *ModifyCloudAssistantSettingsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCloudAssistantSettingsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCloudAssistantSettingsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudAssistantSettingsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCloudAssistantSettingsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCloudAssistantSettingsRequest) GetResourceUsageConfig() *ModifyCloudAssistantSettingsRequestResourceUsageConfig {
	return s.ResourceUsageConfig
}

func (s *ModifyCloudAssistantSettingsRequest) GetSessionManagerConfig() *ModifyCloudAssistantSettingsRequestSessionManagerConfig {
	return s.SessionManagerConfig
}

func (s *ModifyCloudAssistantSettingsRequest) GetSettingType() *string {
	return s.SettingType
}

func (s *ModifyCloudAssistantSettingsRequest) GetSlsDeliveryConfig() *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig {
	return s.SlsDeliveryConfig
}

func (s *ModifyCloudAssistantSettingsRequest) SetAgentUpgradeConfig(v *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) *ModifyCloudAssistantSettingsRequest {
	s.AgentUpgradeConfig = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetOssDeliveryConfig(v *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) *ModifyCloudAssistantSettingsRequest {
	s.OssDeliveryConfig = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetOwnerAccount(v string) *ModifyCloudAssistantSettingsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetOwnerId(v int64) *ModifyCloudAssistantSettingsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetRegionId(v string) *ModifyCloudAssistantSettingsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetResourceOwnerAccount(v string) *ModifyCloudAssistantSettingsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetResourceOwnerId(v int64) *ModifyCloudAssistantSettingsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetResourceUsageConfig(v *ModifyCloudAssistantSettingsRequestResourceUsageConfig) *ModifyCloudAssistantSettingsRequest {
	s.ResourceUsageConfig = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetSessionManagerConfig(v *ModifyCloudAssistantSettingsRequestSessionManagerConfig) *ModifyCloudAssistantSettingsRequest {
	s.SessionManagerConfig = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetSettingType(v string) *ModifyCloudAssistantSettingsRequest {
	s.SettingType = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetSlsDeliveryConfig(v *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) *ModifyCloudAssistantSettingsRequest {
	s.SlsDeliveryConfig = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) Validate() error {
	if s.AgentUpgradeConfig != nil {
		if err := s.AgentUpgradeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OssDeliveryConfig != nil {
		if err := s.OssDeliveryConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceUsageConfig != nil {
		if err := s.ResourceUsageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SessionManagerConfig != nil {
		if err := s.SessionManagerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SlsDeliveryConfig != nil {
		if err := s.SlsDeliveryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCloudAssistantSettingsRequestAgentUpgradeConfig struct {
	// A list of time windows during which the agent is allowed to be upgraded. The time windows are accurate to minutes and are in UTC by default.
	//
	// The interval between two consecutive time windows must be at least 1 hour.
	//
	// Format: StartTime(HH:mm)-EndTime(HH:mm).
	//
	// For example, [
	//
	// "02:00-03:00",
	//
	// "05:00-06:00"
	//
	// ]
	//
	// indicates that the agent can be upgraded from 2:00 to 3:00 and from 5:00 to 6:00 every day in UTC.
	AllowedUpgradeWindow []*string `json:"AllowedUpgradeWindow,omitempty" xml:"AllowedUpgradeWindow,omitempty" type:"Repeated"`
	// Specifies whether to immediately check the version and perform an update when the Cloud Assistant agent is started. Default value: true.
	//
	// This setting takes effect only when the version of the Cloud Assistant agent is not earlier than the following versions:
	//
	// - Windows: 2.1.4.1065
	//
	// - Linux: 2.2.4.1065
	//
	// example:
	//
	// true
	BootstrapUpgrade *bool `json:"BootstrapUpgrade,omitempty" xml:"BootstrapUpgrade,omitempty"`
	// Specifies whether to disallow the Cloud Assistant agent to check for or perform updates. Default value: false.
	//
	// This setting takes effect only when the version of the Cloud Assistant agent is not earlier than the following versions:
	//
	// - Windows: 2.1.4.1065
	//
	// - Linux: 2.2.4.1065
	//
	// example:
	//
	// false
	DisableUpgrade *bool `json:"DisableUpgrade,omitempty" xml:"DisableUpgrade,omitempty"`
	// Specifies whether to enable custom upgrade configurations for the agent. If you set this parameter to false, the agent attempts to upgrade every 30 minutes by default.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time zone of the time windows for agent upgrade. Default value: UTC.
	//
	// The following formats are supported for the time zone:
	//
	// - Time zone name: for example, Asia/Shanghai (China/Shanghai time) and America/Los_Angeles (US/Los Angeles time).
	//
	// - Offset from Greenwich Mean Time (GMT): for example, GMT+8:00 (UTC+8) and GMT-7:00 (UTC-7). The hour part cannot have a leading zero.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) GetAllowedUpgradeWindow() []*string {
	return s.AllowedUpgradeWindow
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) GetBootstrapUpgrade() *bool {
	return s.BootstrapUpgrade
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) GetDisableUpgrade() *bool {
	return s.DisableUpgrade
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) SetAllowedUpgradeWindow(v []*string) *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig {
	s.AllowedUpgradeWindow = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) SetBootstrapUpgrade(v bool) *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig {
	s.BootstrapUpgrade = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) SetDisableUpgrade(v bool) *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig {
	s.DisableUpgrade = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) SetEnabled(v bool) *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig {
	s.Enabled = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) SetTimeZone(v string) *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig {
	s.TimeZone = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCloudAssistantSettingsRequestOssDeliveryConfig struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// example-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// Specifies whether to enable the feature of delivering records to OSS. Default value: false.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The OSS encryption algorithm. Valid values:
	//
	// - AES256
	//
	// - SM4
	//
	// example:
	//
	// AES256
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" xml:"EncryptionAlgorithm,omitempty"`
	// The ID of the customer master key (CMK) when KMS encryption is used.
	//
	// example:
	//
	// a807****7a70e
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The OSS encryption mode. Valid values:
	//
	// - Inherit: inherits the bucket encryption.
	//
	// - OssManaged: uses OSS-managed server-side encryption.
	//
	// - KMS: uses KMS encryption.
	//
	// example:
	//
	// Inherit
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The prefix of the directory in the OSS bucket. The following limits apply:
	//
	// - The prefix can be up to 254 characters in length.
	//
	// - The prefix cannot start with a forward slash (/) or a backslash ().
	//
	// Note: If you want to deliver records to the root directory of the bucket, enter "". To clear the prefix that is previously set, enter "".
	//
	// example:
	//
	// sessionmanager/audit
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ModifyCloudAssistantSettingsRequestOssDeliveryConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetEncryptionAlgorithm() *string {
	return s.EncryptionAlgorithm
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetEncryptionKeyId() *string {
	return s.EncryptionKeyId
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetEncryptionType() *string {
	return s.EncryptionType
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetBucketName(v string) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.BucketName = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetEnabled(v bool) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.Enabled = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetEncryptionAlgorithm(v string) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.EncryptionAlgorithm = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetEncryptionKeyId(v string) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.EncryptionKeyId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetEncryptionType(v string) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.EncryptionType = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetPrefix(v string) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.Prefix = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCloudAssistantSettingsRequestResourceUsageConfig struct {
	// The maximum CPU usage that is allowed for the main process of the Cloud Assistant agent.
	//
	// - Unit: %.
	//
	// - Valid values: 10 to 95.
	//
	// - Default value: 20.
	//
	// example:
	//
	// 20
	CpuLimit *int32 `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// Specifies whether to retain the script file of a command in the Cloud Assistant directory after the command execution is complete.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	KeepScriptFile *bool `json:"KeepScriptFile,omitempty" xml:"KeepScriptFile,omitempty"`
	// The maximum number of Cloud Assistant log files that can be retained.
	//
	// - Default value: 30.
	//
	// - Minimum value: 7.
	//
	// - Maximum value: 365.
	//
	// example:
	//
	// 30
	LogFileCountLimit *int32 `json:"LogFileCountLimit,omitempty" xml:"LogFileCountLimit,omitempty"`
	// The maximum size of a single Cloud Assistant log file. You must specify a unit (B, KB, or MB).
	//
	// - Default value: 100 MB.
	//
	// - Minimum value: 10 MB.
	//
	// - Maximum value: 1024 MB.
	//
	// example:
	//
	// 10MB
	LogSizeLimit *string `json:"LogSizeLimit,omitempty" xml:"LogSizeLimit,omitempty"`
	// The maximum memory usage that is allowed for the main process of the Cloud Assistant agent. You must specify a unit (B, KB, or MB).
	//
	// - Default value: 50 MB.
	//
	// - Minimum value: 35 MB.
	//
	// - Maximum value: 1024 MB.
	//
	// example:
	//
	// 50MB
	MemoryLimit *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// The maximum number of consecutive times that CPU or memory usage can exceed the specified limits. If the limits are consecutively exceeded for the specified number of times, the Cloud Assistant agent is automatically stopped.
	//
	// - Default value: 3.
	//
	// - Minimum value: 3.
	//
	// example:
	//
	// 3
	OverloadLimit *int32 `json:"OverloadLimit,omitempty" xml:"OverloadLimit,omitempty"`
}

func (s ModifyCloudAssistantSettingsRequestResourceUsageConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequestResourceUsageConfig) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) GetCpuLimit() *int32 {
	return s.CpuLimit
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) GetKeepScriptFile() *bool {
	return s.KeepScriptFile
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) GetLogFileCountLimit() *int32 {
	return s.LogFileCountLimit
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) GetLogSizeLimit() *string {
	return s.LogSizeLimit
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) GetMemoryLimit() *string {
	return s.MemoryLimit
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) GetOverloadLimit() *int32 {
	return s.OverloadLimit
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) SetCpuLimit(v int32) *ModifyCloudAssistantSettingsRequestResourceUsageConfig {
	s.CpuLimit = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) SetKeepScriptFile(v bool) *ModifyCloudAssistantSettingsRequestResourceUsageConfig {
	s.KeepScriptFile = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) SetLogFileCountLimit(v int32) *ModifyCloudAssistantSettingsRequestResourceUsageConfig {
	s.LogFileCountLimit = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) SetLogSizeLimit(v string) *ModifyCloudAssistantSettingsRequestResourceUsageConfig {
	s.LogSizeLimit = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) SetMemoryLimit(v string) *ModifyCloudAssistantSettingsRequestResourceUsageConfig {
	s.MemoryLimit = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) SetOverloadLimit(v int32) *ModifyCloudAssistantSettingsRequestResourceUsageConfig {
	s.OverloadLimit = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestResourceUsageConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCloudAssistantSettingsRequestSessionManagerConfig struct {
	// The switch for the Session Manager feature. Valid values:
	//
	// - true: enables the feature.
	//
	// - false: disables the feature.
	//
	// Note:
	//
	// - After you enable or disable the Session Manager feature, the setting takes effect for all regions.
	//
	// example:
	//
	// true
	SessionManagerEnabled *bool `json:"SessionManagerEnabled,omitempty" xml:"SessionManagerEnabled,omitempty"`
}

func (s ModifyCloudAssistantSettingsRequestSessionManagerConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequestSessionManagerConfig) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequestSessionManagerConfig) GetSessionManagerEnabled() *bool {
	return s.SessionManagerEnabled
}

func (s *ModifyCloudAssistantSettingsRequestSessionManagerConfig) SetSessionManagerEnabled(v bool) *ModifyCloudAssistantSettingsRequestSessionManagerConfig {
	s.SessionManagerEnabled = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestSessionManagerConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCloudAssistantSettingsRequestSlsDeliveryConfig struct {
	// Specifies whether to enable the feature of delivering records to SLS.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the SLS Logstore.
	//
	// example:
	//
	// example-logstore
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	// The name of the SLS project.
	//
	// example:
	//
	// example-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) GetProjectName() *string {
	return s.ProjectName
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) SetEnabled(v bool) *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig {
	s.Enabled = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) SetLogstoreName(v string) *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig {
	s.LogstoreName = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) SetProjectName(v string) *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig {
	s.ProjectName = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) Validate() error {
	return dara.Validate(s)
}
