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
	// The Cloud Assistant Agent upgrade configuration.
	AgentUpgradeConfig *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig `json:"AgentUpgradeConfig,omitempty" xml:"AgentUpgradeConfig,omitempty" type:"Struct"`
	// The OSS delivery configuration.
	OssDeliveryConfig *ModifyCloudAssistantSettingsRequestOssDeliveryConfig `json:"OssDeliveryConfig,omitempty" xml:"OssDeliveryConfig,omitempty" type:"Struct"`
	OwnerAccount      *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The Cloud Assistant resource usage configuration. This parameter takes effect only when the Cloud Assistant Agent version meets the following minimum requirements:
	//
	// - Windows: 2.1.4.1065
	//
	// - Linux: 2.2.4.1065
	ResourceUsageConfig *ModifyCloudAssistantSettingsRequestResourceUsageConfig `json:"ResourceUsageConfig,omitempty" xml:"ResourceUsageConfig,omitempty" type:"Struct"`
	// The Cloud Assistant session feature configuration.
	SessionManagerConfig *ModifyCloudAssistantSettingsRequestSessionManagerConfig `json:"SessionManagerConfig,omitempty" xml:"SessionManagerConfig,omitempty" type:"Struct"`
	// The service configuration type. Valid values:
	//
	// - SessionManagerDelivery: session operation log delivery.
	//
	// - InvocationDelivery: task execution log delivery.
	//
	// - AgentUpgradeConfig: Cloud Assistant Agent upgrade configuration.
	//
	// - SessionManagerConfig: Cloud Assistant SessionManager configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// SessionManagerDelivery
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
	// The Simple Log Service (SLS) delivery configuration.
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
	// The list of time windows during which upgrades are allowed. The time can be specified down to the minute. The default time zone is UTC.
	//
	// The interval between time windows cannot be less than 1 hour.
	//
	// Format: Start time (HH:mm)-End time (HH:mm).
	//
	// Example: [
	//
	// "02:00-03:00",
	//
	// "05:00-06:00"
	//
	// ]
	//
	// This indicates that upgrades are allowed daily from 02:00 to 03:00 and from 05:00 to 06:00 in the UTC time zone.
	AllowedUpgradeWindow []*string `json:"AllowedUpgradeWindow,omitempty" xml:"AllowedUpgradeWindow,omitempty" type:"Repeated"`
	// Specifies whether the Cloud Assistant Agent checks for updates and performs an upgrade immediately upon startup. Default value: true.
	//
	// This parameter takes effect only when the Cloud Assistant Agent version meets the following minimum requirements:
	//
	// - Windows: 2.1.4.1065
	//
	// - Linux: 2.2.4.1065
	//
	// example:
	//
	// true
	BootstrapUpgrade *bool `json:"BootstrapUpgrade,omitempty" xml:"BootstrapUpgrade,omitempty"`
	// Specifies whether to prevent the Cloud Assistant Agent from checking for and performing updates. Default value: false.
	//
	// This parameter takes effect only when the Cloud Assistant Agent version meets the following minimum requirements:
	//
	// - Windows: 2.1.4.1065
	//
	// - Linux: 2.2.4.1065
	//
	// example:
	//
	// false
	DisableUpgrade *bool `json:"DisableUpgrade,omitempty" xml:"DisableUpgrade,omitempty"`
	// Specifies whether to enable the custom Agent upgrade configuration. If this parameter is set to false, the system attempts to upgrade the Agent every 30 minutes by default.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time zone for the allowed upgrade time windows. Default value: UTC.
	//
	// The time zone can be specified in the following formats:
	//
	// - Full time zone name, such as Asia/Shanghai or America/Los_Angeles.
	//
	// - GMT offset from Greenwich Mean Time, such as GMT+8:00 or GMT-7:00. Leading zeros are not supported for the hour value.
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
	// Specifies whether to enable delivery to OSS. Default value: false.
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
	// The ID of the customer master key (CMK) when the encryption method is set to KMS.
	//
	// example:
	//
	// a807****7a70e
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The OSS encryption method. Valid values:
	//
	// - Inherit: inherits the encryption method of the bucket.
	//
	// - OssManaged: OSS-managed encryption.
	//
	// - KMS: Key Management Service (KMS) encryption.
	//
	// example:
	//
	// Inherit
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The directory prefix of the OSS bucket. The following limits apply:
	//
	// - The prefix cannot exceed 254 characters in length.
	//
	// - The prefix cannot start with a forward slash (/) or a backslash (\\).
	//
	// > Note: Set this parameter to an empty string ("") if no directory prefix is required. If a prefix was previously configured and is no longer needed, set this parameter to an empty string ("") to clear it.
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
	// The maximum CPU usage allowed for the Cloud Assistant Agent main process.
	//
	// - Unit: percentage.
	//
	// - Valid values: 10 to 95.
	//
	// - Default value: 20.
	//
	// example:
	//
	// 20
	CpuLimit *int32 `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// Specifies whether to retain the script file in the Cloud Assistant directory after command execution is complete.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	KeepScriptFile *bool `json:"KeepScriptFile,omitempty" xml:"KeepScriptFile,omitempty"`
	// The maximum number of Cloud Assistant log files to retain.
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
	// The maximum size of a single Cloud Assistant log file. You must specify the unit (B|KB|MB).
	//
	// - Default value: 100MB.
	//
	// - Minimum value: 10MB.
	//
	// - Maximum value: 1024MB.
	//
	// example:
	//
	// 10MB
	LogSizeLimit *string `json:"LogSizeLimit,omitempty" xml:"LogSizeLimit,omitempty"`
	// The maximum memory usage allowed for the Cloud Assistant Agent main process. You must specify the unit (B|KB|MB).
	//
	// - Default value: 50MB.
	//
	// - Minimum value: 35MB.
	//
	// - Maximum value: 1024MB.
	//
	// example:
	//
	// 50MB
	MemoryLimit *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// The maximum number of consecutive times that CPU or memory resources usage can exceed the limit before the Cloud Assistant Agent automatically stops running.
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
	// Specifies whether to enable the Cloud Assistant session feature. Valid values:
	//
	// 	- true: Enabled.
	//
	// 	- false: Disabled.
	//
	// Note:
	//
	// 	- Enabling or disabling the session feature takes effect across all regions.
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
	// Specifies whether to enable delivery to SLS.
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
