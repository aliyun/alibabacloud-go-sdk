// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDRPlanConfigurationDetail interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigs(v []*ApplicationConfig) *DRPlanConfigurationDetail
	GetApplicationConfigs() []*ApplicationConfig
	SetApplications(v []*Application) *DRPlanConfigurationDetail
	GetApplications() []*Application
	SetBootstrapScripts(v []*Script) *DRPlanConfigurationDetail
	GetBootstrapScripts() []*Script
	SetClusterName(v string) *DRPlanConfigurationDetail
	GetClusterName() *string
	SetClusterType(v string) *DRPlanConfigurationDetail
	GetClusterType() *string
	SetDeletionProtection(v bool) *DRPlanConfigurationDetail
	GetDeletionProtection() *bool
	SetDeployMode(v string) *DRPlanConfigurationDetail
	GetDeployMode() *string
	SetDescription(v string) *DRPlanConfigurationDetail
	GetDescription() *string
	SetLogCollectStrategy(v string) *DRPlanConfigurationDetail
	GetLogCollectStrategy() *string
	SetManagedScalingPolicy(v *DRPlanConfigurationDetailManagedScalingPolicy) *DRPlanConfigurationDetail
	GetManagedScalingPolicy() *DRPlanConfigurationDetailManagedScalingPolicy
	SetMetaStoreType(v string) *DRPlanConfigurationDetail
	GetMetaStoreType() *string
	SetNodeAttributes(v *NodeAttributes) *DRPlanConfigurationDetail
	GetNodeAttributes() *NodeAttributes
	SetNodeGroups(v []*NodeGroupConfig) *DRPlanConfigurationDetail
	GetNodeGroups() []*NodeGroupConfig
	SetPaymentType(v string) *DRPlanConfigurationDetail
	GetPaymentType() *string
	SetRegionId(v string) *DRPlanConfigurationDetail
	GetRegionId() *string
	SetReleaseVersion(v string) *DRPlanConfigurationDetail
	GetReleaseVersion() *string
	SetResourceGroupId(v string) *DRPlanConfigurationDetail
	GetResourceGroupId() *string
	SetScalingPolicies(v []*ScalingPolicy) *DRPlanConfigurationDetail
	GetScalingPolicies() []*ScalingPolicy
	SetSecurityMode(v string) *DRPlanConfigurationDetail
	GetSecurityMode() *string
	SetSubscriptionConfig(v *SubscriptionConfig) *DRPlanConfigurationDetail
	GetSubscriptionConfig() *SubscriptionConfig
	SetTags(v []*Tag) *DRPlanConfigurationDetail
	GetTags() []*Tag
}

type DRPlanConfigurationDetail struct {
	ApplicationConfigs []*ApplicationConfig `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
	Applications       []*Application       `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	BootstrapScripts   []*Script            `json:"BootstrapScripts,omitempty" xml:"BootstrapScripts,omitempty" type:"Repeated"`
	ClusterName        *string              `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	ClusterType        *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	DeletionProtection *bool   `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// This parameter is required.
	DeployMode  *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	LogCollectStrategy   *string                                        `json:"LogCollectStrategy,omitempty" xml:"LogCollectStrategy,omitempty"`
	ManagedScalingPolicy *DRPlanConfigurationDetailManagedScalingPolicy `json:"ManagedScalingPolicy,omitempty" xml:"ManagedScalingPolicy,omitempty" type:"Struct"`
	// This parameter is required.
	MetaStoreType *string `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	// This parameter is required.
	NodeAttributes *NodeAttributes    `json:"NodeAttributes,omitempty" xml:"NodeAttributes,omitempty"`
	NodeGroups     []*NodeGroupConfig `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// This parameter is required.
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ReleaseVersion  *string          `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	ResourceGroupId *string          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScalingPolicies []*ScalingPolicy `json:"ScalingPolicies,omitempty" xml:"ScalingPolicies,omitempty" type:"Repeated"`
	// This parameter is required.
	SecurityMode       *string             `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
	SubscriptionConfig *SubscriptionConfig `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	Tags               []*Tag              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DRPlanConfigurationDetail) String() string {
	return dara.Prettify(s)
}

func (s DRPlanConfigurationDetail) GoString() string {
	return s.String()
}

func (s *DRPlanConfigurationDetail) GetApplicationConfigs() []*ApplicationConfig {
	return s.ApplicationConfigs
}

func (s *DRPlanConfigurationDetail) GetApplications() []*Application {
	return s.Applications
}

func (s *DRPlanConfigurationDetail) GetBootstrapScripts() []*Script {
	return s.BootstrapScripts
}

func (s *DRPlanConfigurationDetail) GetClusterName() *string {
	return s.ClusterName
}

func (s *DRPlanConfigurationDetail) GetClusterType() *string {
	return s.ClusterType
}

func (s *DRPlanConfigurationDetail) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DRPlanConfigurationDetail) GetDeployMode() *string {
	return s.DeployMode
}

func (s *DRPlanConfigurationDetail) GetDescription() *string {
	return s.Description
}

func (s *DRPlanConfigurationDetail) GetLogCollectStrategy() *string {
	return s.LogCollectStrategy
}

func (s *DRPlanConfigurationDetail) GetManagedScalingPolicy() *DRPlanConfigurationDetailManagedScalingPolicy {
	return s.ManagedScalingPolicy
}

func (s *DRPlanConfigurationDetail) GetMetaStoreType() *string {
	return s.MetaStoreType
}

func (s *DRPlanConfigurationDetail) GetNodeAttributes() *NodeAttributes {
	return s.NodeAttributes
}

func (s *DRPlanConfigurationDetail) GetNodeGroups() []*NodeGroupConfig {
	return s.NodeGroups
}

func (s *DRPlanConfigurationDetail) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DRPlanConfigurationDetail) GetRegionId() *string {
	return s.RegionId
}

func (s *DRPlanConfigurationDetail) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *DRPlanConfigurationDetail) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DRPlanConfigurationDetail) GetScalingPolicies() []*ScalingPolicy {
	return s.ScalingPolicies
}

func (s *DRPlanConfigurationDetail) GetSecurityMode() *string {
	return s.SecurityMode
}

func (s *DRPlanConfigurationDetail) GetSubscriptionConfig() *SubscriptionConfig {
	return s.SubscriptionConfig
}

func (s *DRPlanConfigurationDetail) GetTags() []*Tag {
	return s.Tags
}

func (s *DRPlanConfigurationDetail) SetApplicationConfigs(v []*ApplicationConfig) *DRPlanConfigurationDetail {
	s.ApplicationConfigs = v
	return s
}

func (s *DRPlanConfigurationDetail) SetApplications(v []*Application) *DRPlanConfigurationDetail {
	s.Applications = v
	return s
}

func (s *DRPlanConfigurationDetail) SetBootstrapScripts(v []*Script) *DRPlanConfigurationDetail {
	s.BootstrapScripts = v
	return s
}

func (s *DRPlanConfigurationDetail) SetClusterName(v string) *DRPlanConfigurationDetail {
	s.ClusterName = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetClusterType(v string) *DRPlanConfigurationDetail {
	s.ClusterType = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetDeletionProtection(v bool) *DRPlanConfigurationDetail {
	s.DeletionProtection = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetDeployMode(v string) *DRPlanConfigurationDetail {
	s.DeployMode = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetDescription(v string) *DRPlanConfigurationDetail {
	s.Description = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetLogCollectStrategy(v string) *DRPlanConfigurationDetail {
	s.LogCollectStrategy = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetManagedScalingPolicy(v *DRPlanConfigurationDetailManagedScalingPolicy) *DRPlanConfigurationDetail {
	s.ManagedScalingPolicy = v
	return s
}

func (s *DRPlanConfigurationDetail) SetMetaStoreType(v string) *DRPlanConfigurationDetail {
	s.MetaStoreType = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetNodeAttributes(v *NodeAttributes) *DRPlanConfigurationDetail {
	s.NodeAttributes = v
	return s
}

func (s *DRPlanConfigurationDetail) SetNodeGroups(v []*NodeGroupConfig) *DRPlanConfigurationDetail {
	s.NodeGroups = v
	return s
}

func (s *DRPlanConfigurationDetail) SetPaymentType(v string) *DRPlanConfigurationDetail {
	s.PaymentType = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetRegionId(v string) *DRPlanConfigurationDetail {
	s.RegionId = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetReleaseVersion(v string) *DRPlanConfigurationDetail {
	s.ReleaseVersion = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetResourceGroupId(v string) *DRPlanConfigurationDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetScalingPolicies(v []*ScalingPolicy) *DRPlanConfigurationDetail {
	s.ScalingPolicies = v
	return s
}

func (s *DRPlanConfigurationDetail) SetSecurityMode(v string) *DRPlanConfigurationDetail {
	s.SecurityMode = &v
	return s
}

func (s *DRPlanConfigurationDetail) SetSubscriptionConfig(v *SubscriptionConfig) *DRPlanConfigurationDetail {
	s.SubscriptionConfig = v
	return s
}

func (s *DRPlanConfigurationDetail) SetTags(v []*Tag) *DRPlanConfigurationDetail {
	s.Tags = v
	return s
}

func (s *DRPlanConfigurationDetail) Validate() error {
	return dara.Validate(s)
}

type DRPlanConfigurationDetailManagedScalingPolicy struct {
	Constraints *ManagedScalingConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
}

func (s DRPlanConfigurationDetailManagedScalingPolicy) String() string {
	return dara.Prettify(s)
}

func (s DRPlanConfigurationDetailManagedScalingPolicy) GoString() string {
	return s.String()
}

func (s *DRPlanConfigurationDetailManagedScalingPolicy) GetConstraints() *ManagedScalingConstraints {
	return s.Constraints
}

func (s *DRPlanConfigurationDetailManagedScalingPolicy) SetConstraints(v *ManagedScalingConstraints) *DRPlanConfigurationDetailManagedScalingPolicy {
	s.Constraints = v
	return s
}

func (s *DRPlanConfigurationDetailManagedScalingPolicy) Validate() error {
	return dara.Validate(s)
}
