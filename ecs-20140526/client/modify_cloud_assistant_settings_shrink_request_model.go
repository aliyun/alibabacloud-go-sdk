// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudAssistantSettingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentUpgradeConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetAgentUpgradeConfigShrink() *string
	SetOssDeliveryConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetOssDeliveryConfigShrink() *string
	SetOwnerAccount(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCloudAssistantSettingsShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCloudAssistantSettingsShrinkRequest
	GetResourceOwnerId() *int64
	SetResourceUsageConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetResourceUsageConfigShrink() *string
	SetSessionManagerConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetSessionManagerConfigShrink() *string
	SetSettingType(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetSettingType() *string
	SetSlsDeliveryConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetSlsDeliveryConfigShrink() *string
}

type ModifyCloudAssistantSettingsShrinkRequest struct {
	// The configurations of upgrading the Cloud Assistant agent.
	AgentUpgradeConfigShrink *string `json:"AgentUpgradeConfig,omitempty" xml:"AgentUpgradeConfig,omitempty"`
	// The configurations of delivering records to OSS.
	OssDeliveryConfigShrink *string `json:"OssDeliveryConfig,omitempty" xml:"OssDeliveryConfig,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	ResourceUsageConfigShrink *string `json:"ResourceUsageConfig,omitempty" xml:"ResourceUsageConfig,omitempty"`
	// The configurations of the Session Manager feature.
	SessionManagerConfigShrink *string `json:"SessionManagerConfig,omitempty" xml:"SessionManagerConfig,omitempty"`
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
	SlsDeliveryConfigShrink *string `json:"SlsDeliveryConfig,omitempty" xml:"SlsDeliveryConfig,omitempty"`
}

func (s ModifyCloudAssistantSettingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetAgentUpgradeConfigShrink() *string {
	return s.AgentUpgradeConfigShrink
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetOssDeliveryConfigShrink() *string {
	return s.OssDeliveryConfigShrink
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetResourceUsageConfigShrink() *string {
	return s.ResourceUsageConfigShrink
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetSessionManagerConfigShrink() *string {
	return s.SessionManagerConfigShrink
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetSettingType() *string {
	return s.SettingType
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetSlsDeliveryConfigShrink() *string {
	return s.SlsDeliveryConfigShrink
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetAgentUpgradeConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.AgentUpgradeConfigShrink = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetOssDeliveryConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.OssDeliveryConfigShrink = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetOwnerAccount(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetOwnerId(v int64) *ModifyCloudAssistantSettingsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetRegionId(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetResourceOwnerAccount(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetResourceOwnerId(v int64) *ModifyCloudAssistantSettingsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetResourceUsageConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.ResourceUsageConfigShrink = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetSessionManagerConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.SessionManagerConfigShrink = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetSettingType(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.SettingType = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetSlsDeliveryConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.SlsDeliveryConfigShrink = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
