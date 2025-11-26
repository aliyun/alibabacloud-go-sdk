// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDRPlanConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigs(v []*ApplicationConfig) *DRPlanConfiguration
	GetApplicationConfigs() []*ApplicationConfig
	SetApplications(v []*Application) *DRPlanConfiguration
	GetApplications() []*Application
	SetBootstrapScripts(v []*Script) *DRPlanConfiguration
	GetBootstrapScripts() []*Script
	SetClusterName(v string) *DRPlanConfiguration
	GetClusterName() *string
	SetClusterType(v string) *DRPlanConfiguration
	GetClusterType() *string
	SetDeletionProtection(v bool) *DRPlanConfiguration
	GetDeletionProtection() *bool
	SetDeployMode(v string) *DRPlanConfiguration
	GetDeployMode() *string
	SetDescription(v string) *DRPlanConfiguration
	GetDescription() *string
	SetLogCollectStrategy(v string) *DRPlanConfiguration
	GetLogCollectStrategy() *string
	SetManagedScalingPolicy(v *DRPlanConfigurationManagedScalingPolicy) *DRPlanConfiguration
	GetManagedScalingPolicy() *DRPlanConfigurationManagedScalingPolicy
	SetNodeAttributes(v *NodeAttributes) *DRPlanConfiguration
	GetNodeAttributes() *NodeAttributes
	SetNodeGroups(v []*NodeGroupConfig) *DRPlanConfiguration
	GetNodeGroups() []*NodeGroupConfig
	SetPaymentType(v string) *DRPlanConfiguration
	GetPaymentType() *string
	SetRegionId(v string) *DRPlanConfiguration
	GetRegionId() *string
	SetReleaseVersion(v string) *DRPlanConfiguration
	GetReleaseVersion() *string
	SetResourceGroupId(v string) *DRPlanConfiguration
	GetResourceGroupId() *string
	SetSecurityMode(v string) *DRPlanConfiguration
	GetSecurityMode() *string
	SetSubscriptionConfig(v *SubscriptionConfig) *DRPlanConfiguration
	GetSubscriptionConfig() *SubscriptionConfig
	SetTags(v []*DRPlanConfigurationTags) *DRPlanConfiguration
	GetTags() []*DRPlanConfigurationTags
}

type DRPlanConfiguration struct {
	ApplicationConfigs []*ApplicationConfig `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
	// This parameter is required.
	Applications     []*Application `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	BootstrapScripts []*Script      `json:"BootstrapScripts,omitempty" xml:"BootstrapScripts,omitempty" type:"Repeated"`
	// This parameter is required.
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	ClusterType          *string                                  `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	DeletionProtection   *bool                                    `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	DeployMode           *string                                  `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	Description          *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	LogCollectStrategy   *string                                  `json:"LogCollectStrategy,omitempty" xml:"LogCollectStrategy,omitempty"`
	ManagedScalingPolicy *DRPlanConfigurationManagedScalingPolicy `json:"ManagedScalingPolicy,omitempty" xml:"ManagedScalingPolicy,omitempty" type:"Struct"`
	// This parameter is required.
	NodeAttributes *NodeAttributes `json:"NodeAttributes,omitempty" xml:"NodeAttributes,omitempty"`
	// This parameter is required.
	NodeGroups []*NodeGroupConfig `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// This parameter is required.
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ReleaseVersion  *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	SecurityMode       *string                    `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
	SubscriptionConfig *SubscriptionConfig        `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	Tags               []*DRPlanConfigurationTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DRPlanConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DRPlanConfiguration) GoString() string {
	return s.String()
}

func (s *DRPlanConfiguration) GetApplicationConfigs() []*ApplicationConfig {
	return s.ApplicationConfigs
}

func (s *DRPlanConfiguration) GetApplications() []*Application {
	return s.Applications
}

func (s *DRPlanConfiguration) GetBootstrapScripts() []*Script {
	return s.BootstrapScripts
}

func (s *DRPlanConfiguration) GetClusterName() *string {
	return s.ClusterName
}

func (s *DRPlanConfiguration) GetClusterType() *string {
	return s.ClusterType
}

func (s *DRPlanConfiguration) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DRPlanConfiguration) GetDeployMode() *string {
	return s.DeployMode
}

func (s *DRPlanConfiguration) GetDescription() *string {
	return s.Description
}

func (s *DRPlanConfiguration) GetLogCollectStrategy() *string {
	return s.LogCollectStrategy
}

func (s *DRPlanConfiguration) GetManagedScalingPolicy() *DRPlanConfigurationManagedScalingPolicy {
	return s.ManagedScalingPolicy
}

func (s *DRPlanConfiguration) GetNodeAttributes() *NodeAttributes {
	return s.NodeAttributes
}

func (s *DRPlanConfiguration) GetNodeGroups() []*NodeGroupConfig {
	return s.NodeGroups
}

func (s *DRPlanConfiguration) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DRPlanConfiguration) GetRegionId() *string {
	return s.RegionId
}

func (s *DRPlanConfiguration) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *DRPlanConfiguration) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DRPlanConfiguration) GetSecurityMode() *string {
	return s.SecurityMode
}

func (s *DRPlanConfiguration) GetSubscriptionConfig() *SubscriptionConfig {
	return s.SubscriptionConfig
}

func (s *DRPlanConfiguration) GetTags() []*DRPlanConfigurationTags {
	return s.Tags
}

func (s *DRPlanConfiguration) SetApplicationConfigs(v []*ApplicationConfig) *DRPlanConfiguration {
	s.ApplicationConfigs = v
	return s
}

func (s *DRPlanConfiguration) SetApplications(v []*Application) *DRPlanConfiguration {
	s.Applications = v
	return s
}

func (s *DRPlanConfiguration) SetBootstrapScripts(v []*Script) *DRPlanConfiguration {
	s.BootstrapScripts = v
	return s
}

func (s *DRPlanConfiguration) SetClusterName(v string) *DRPlanConfiguration {
	s.ClusterName = &v
	return s
}

func (s *DRPlanConfiguration) SetClusterType(v string) *DRPlanConfiguration {
	s.ClusterType = &v
	return s
}

func (s *DRPlanConfiguration) SetDeletionProtection(v bool) *DRPlanConfiguration {
	s.DeletionProtection = &v
	return s
}

func (s *DRPlanConfiguration) SetDeployMode(v string) *DRPlanConfiguration {
	s.DeployMode = &v
	return s
}

func (s *DRPlanConfiguration) SetDescription(v string) *DRPlanConfiguration {
	s.Description = &v
	return s
}

func (s *DRPlanConfiguration) SetLogCollectStrategy(v string) *DRPlanConfiguration {
	s.LogCollectStrategy = &v
	return s
}

func (s *DRPlanConfiguration) SetManagedScalingPolicy(v *DRPlanConfigurationManagedScalingPolicy) *DRPlanConfiguration {
	s.ManagedScalingPolicy = v
	return s
}

func (s *DRPlanConfiguration) SetNodeAttributes(v *NodeAttributes) *DRPlanConfiguration {
	s.NodeAttributes = v
	return s
}

func (s *DRPlanConfiguration) SetNodeGroups(v []*NodeGroupConfig) *DRPlanConfiguration {
	s.NodeGroups = v
	return s
}

func (s *DRPlanConfiguration) SetPaymentType(v string) *DRPlanConfiguration {
	s.PaymentType = &v
	return s
}

func (s *DRPlanConfiguration) SetRegionId(v string) *DRPlanConfiguration {
	s.RegionId = &v
	return s
}

func (s *DRPlanConfiguration) SetReleaseVersion(v string) *DRPlanConfiguration {
	s.ReleaseVersion = &v
	return s
}

func (s *DRPlanConfiguration) SetResourceGroupId(v string) *DRPlanConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *DRPlanConfiguration) SetSecurityMode(v string) *DRPlanConfiguration {
	s.SecurityMode = &v
	return s
}

func (s *DRPlanConfiguration) SetSubscriptionConfig(v *SubscriptionConfig) *DRPlanConfiguration {
	s.SubscriptionConfig = v
	return s
}

func (s *DRPlanConfiguration) SetTags(v []*DRPlanConfigurationTags) *DRPlanConfiguration {
	s.Tags = v
	return s
}

func (s *DRPlanConfiguration) Validate() error {
	if s.ApplicationConfigs != nil {
		for _, item := range s.ApplicationConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.BootstrapScripts != nil {
		for _, item := range s.BootstrapScripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ManagedScalingPolicy != nil {
		if err := s.ManagedScalingPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.NodeAttributes != nil {
		if err := s.NodeAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.NodeGroups != nil {
		for _, item := range s.NodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubscriptionConfig != nil {
		if err := s.SubscriptionConfig.Validate(); err != nil {
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

type DRPlanConfigurationManagedScalingPolicy struct {
	Constraints *ManagedScalingConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
}

func (s DRPlanConfigurationManagedScalingPolicy) String() string {
	return dara.Prettify(s)
}

func (s DRPlanConfigurationManagedScalingPolicy) GoString() string {
	return s.String()
}

func (s *DRPlanConfigurationManagedScalingPolicy) GetConstraints() *ManagedScalingConstraints {
	return s.Constraints
}

func (s *DRPlanConfigurationManagedScalingPolicy) SetConstraints(v *ManagedScalingConstraints) *DRPlanConfigurationManagedScalingPolicy {
	s.Constraints = v
	return s
}

func (s *DRPlanConfigurationManagedScalingPolicy) Validate() error {
	if s.Constraints != nil {
		if err := s.Constraints.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DRPlanConfigurationTags struct {
	// 标签键。必填参数，不允许为空字符串。最多支持128个字符，不能以aliyun和acs:开头，不能包含http://或https://。
	//
	// This parameter is required.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值。非必填，可以为空字符串。最多支持128个字符，不能以acs:开头，不能包含http://或者https://。
	//
	// example:
	//
	// IT
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DRPlanConfigurationTags) String() string {
	return dara.Prettify(s)
}

func (s DRPlanConfigurationTags) GoString() string {
	return s.String()
}

func (s *DRPlanConfigurationTags) GetKey() *string {
	return s.Key
}

func (s *DRPlanConfigurationTags) GetValue() *string {
	return s.Value
}

func (s *DRPlanConfigurationTags) SetKey(v string) *DRPlanConfigurationTags {
	s.Key = &v
	return s
}

func (s *DRPlanConfigurationTags) SetValue(v string) *DRPlanConfigurationTags {
	s.Value = &v
	return s
}

func (s *DRPlanConfigurationTags) Validate() error {
	return dara.Validate(s)
}
