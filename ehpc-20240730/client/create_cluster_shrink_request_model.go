// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalPackagesShrink(v string) *CreateClusterShrinkRequest
	GetAdditionalPackagesShrink() *string
	SetAddonsShrink(v string) *CreateClusterShrinkRequest
	GetAddonsShrink() *string
	SetClientVersion(v string) *CreateClusterShrinkRequest
	GetClientVersion() *string
	SetClusterCategory(v string) *CreateClusterShrinkRequest
	GetClusterCategory() *string
	SetClusterCredentialsShrink(v string) *CreateClusterShrinkRequest
	GetClusterCredentialsShrink() *string
	SetClusterCustomConfigurationShrink(v string) *CreateClusterShrinkRequest
	GetClusterCustomConfigurationShrink() *string
	SetClusterDescription(v string) *CreateClusterShrinkRequest
	GetClusterDescription() *string
	SetClusterMode(v string) *CreateClusterShrinkRequest
	GetClusterMode() *string
	SetClusterName(v string) *CreateClusterShrinkRequest
	GetClusterName() *string
	SetClusterVSwitchId(v string) *CreateClusterShrinkRequest
	GetClusterVSwitchId() *string
	SetClusterVpcId(v string) *CreateClusterShrinkRequest
	GetClusterVpcId() *string
	SetDeletionProtection(v bool) *CreateClusterShrinkRequest
	GetDeletionProtection() *bool
	SetGrowInterval(v int32) *CreateClusterShrinkRequest
	GetGrowInterval() *int32
	SetIdleInterval(v int32) *CreateClusterShrinkRequest
	GetIdleInterval() *int32
	SetIsEnterpriseSecurityGroup(v bool) *CreateClusterShrinkRequest
	GetIsEnterpriseSecurityGroup() *bool
	SetManagerShrink(v string) *CreateClusterShrinkRequest
	GetManagerShrink() *string
	SetMaxCoreCount(v int32) *CreateClusterShrinkRequest
	GetMaxCoreCount() *int32
	SetMaxCount(v int32) *CreateClusterShrinkRequest
	GetMaxCount() *int32
	SetQueuesShrink(v string) *CreateClusterShrinkRequest
	GetQueuesShrink() *string
	SetResourceGroupId(v string) *CreateClusterShrinkRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateClusterShrinkRequest
	GetSecurityGroupId() *string
	SetSharedStoragesShrink(v string) *CreateClusterShrinkRequest
	GetSharedStoragesShrink() *string
	SetTagsShrink(v string) *CreateClusterShrinkRequest
	GetTagsShrink() *string
}

type CreateClusterShrinkRequest struct {
	// A list of software to install in the cluster. You can specify up to 10 packages.
	AdditionalPackagesShrink *string `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty"`
	// The configuration of the custom service component for the cluster. Only one component is supported.
	AddonsShrink *string `json:"Addons,omitempty" xml:"Addons,omitempty"`
	// The version of the E-HPC client. By default, the latest version is used.
	//
	// example:
	//
	// 2.1.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// - Standard
	//
	// - Serverless
	//
	// example:
	//
	// Standard
	ClusterCategory *string `json:"ClusterCategory,omitempty" xml:"ClusterCategory,omitempty"`
	// The security credentials for the cluster.
	ClusterCredentialsShrink *string `json:"ClusterCredentials,omitempty" xml:"ClusterCredentials,omitempty"`
	// The post-processing script for the cluster.
	ClusterCustomConfigurationShrink *string `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty"`
	// The description of the cluster. The description must be 2 to 128 characters long and can contain letters, Chinese characters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// slurm22.05.8-cluster-20240718
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The cluster\\"s deployment type. Valid values:
	//
	// - Integrated: An integrated cluster.
	//
	// - Hybrid: A hybrid cloud cluster.
	//
	// - Custom: A custom cluster.
	//
	// example:
	//
	// Integrated
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The name of the cluster. The name must be 2 to 128 characters long and can contain letters, Chinese characters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// slurm22.05.8-cluster-20240718
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the VSwitch for the cluster. The VSwitch must be in the VPC specified by `ClusterVpcId`.
	//
	// Call the [DescribeVpcs](https://help.aliyun.com/document_detail/448581.html) operation to find available VPCs and VSwitches.
	//
	// example:
	//
	// vsw-f8za5p0mwzgdu3wgx****
	ClusterVSwitchId *string `json:"ClusterVSwitchId,omitempty" xml:"ClusterVSwitchId,omitempty"`
	// The ID of the VPC for the cluster.
	//
	// example:
	//
	// vpc-m5efjevmclc0xdmys****
	ClusterVpcId *string `json:"ClusterVpcId,omitempty" xml:"ClusterVpcId,omitempty"`
	// Specifies whether to enable deletion protection for the cluster. This feature prevents the cluster from being deleted via the console or the [DeleteCluster](https://help.aliyun.com/document_detail/424406.html) operation.
	//
	// - true: Enables deletion protection.
	//
	// - false: Disables deletion protection.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DeletionProtection *bool  `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	GrowInterval       *int32 `json:"GrowInterval,omitempty" xml:"GrowInterval,omitempty"`
	IdleInterval       *int32 `json:"IdleInterval,omitempty" xml:"IdleInterval,omitempty"`
	// Specifies whether to use an enterprise security group. Valid values:
	//
	// - true: The system automatically creates and uses an enterprise security group.
	//
	// - false: The system automatically creates and uses a security group.
	//
	// For more information about how to select a security group type, see [Security groups and enterprise security groups](https://help.aliyun.com/document_detail/605897.html).
	//
	// example:
	//
	// false
	IsEnterpriseSecurityGroup *bool `json:"IsEnterpriseSecurityGroup,omitempty" xml:"IsEnterpriseSecurityGroup,omitempty"`
	// Configuration for the cluster manager node.
	ManagerShrink *string `json:"Manager,omitempty" xml:"Manager,omitempty"`
	// The maximum number of CPU cores that the cluster can manage across all compute nodes. Valid values: 0 to 100,000.
	//
	// example:
	//
	// 10000
	MaxCoreCount *int32 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The maximum number of compute nodes that the cluster can manage. Valid values: 0 to 5,000.
	//
	// example:
	//
	// 500
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// Configuration for the cluster queues. You can specify up to 8 queues.
	QueuesShrink *string `json:"Queues,omitempty" xml:"Queues,omitempty"`
	// The ID of the resource group.
	//
	// Call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to find resource group IDs.
	//
	// example:
	//
	// rg-acfmxazb4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group for the cluster.
	//
	// Call the [DescribeSecurityGroups](https://help.aliyun.com/document_detail/25556.html) operation to find available security groups in the current region.
	//
	// example:
	//
	// sg-bp13n61xsydodfyg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Configuration for the cluster\\"s shared storage.
	SharedStoragesShrink *string `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty"`
	// The list of tags to add to the cluster. You can add up to 20 tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) GetAdditionalPackagesShrink() *string {
	return s.AdditionalPackagesShrink
}

func (s *CreateClusterShrinkRequest) GetAddonsShrink() *string {
	return s.AddonsShrink
}

func (s *CreateClusterShrinkRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *CreateClusterShrinkRequest) GetClusterCategory() *string {
	return s.ClusterCategory
}

func (s *CreateClusterShrinkRequest) GetClusterCredentialsShrink() *string {
	return s.ClusterCredentialsShrink
}

func (s *CreateClusterShrinkRequest) GetClusterCustomConfigurationShrink() *string {
	return s.ClusterCustomConfigurationShrink
}

func (s *CreateClusterShrinkRequest) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *CreateClusterShrinkRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *CreateClusterShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterShrinkRequest) GetClusterVSwitchId() *string {
	return s.ClusterVSwitchId
}

func (s *CreateClusterShrinkRequest) GetClusterVpcId() *string {
	return s.ClusterVpcId
}

func (s *CreateClusterShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateClusterShrinkRequest) GetGrowInterval() *int32 {
	return s.GrowInterval
}

func (s *CreateClusterShrinkRequest) GetIdleInterval() *int32 {
	return s.IdleInterval
}

func (s *CreateClusterShrinkRequest) GetIsEnterpriseSecurityGroup() *bool {
	return s.IsEnterpriseSecurityGroup
}

func (s *CreateClusterShrinkRequest) GetManagerShrink() *string {
	return s.ManagerShrink
}

func (s *CreateClusterShrinkRequest) GetMaxCoreCount() *int32 {
	return s.MaxCoreCount
}

func (s *CreateClusterShrinkRequest) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *CreateClusterShrinkRequest) GetQueuesShrink() *string {
	return s.QueuesShrink
}

func (s *CreateClusterShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClusterShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateClusterShrinkRequest) GetSharedStoragesShrink() *string {
	return s.SharedStoragesShrink
}

func (s *CreateClusterShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateClusterShrinkRequest) SetAdditionalPackagesShrink(v string) *CreateClusterShrinkRequest {
	s.AdditionalPackagesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetAddonsShrink(v string) *CreateClusterShrinkRequest {
	s.AddonsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClientVersion(v string) *CreateClusterShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterCategory(v string) *CreateClusterShrinkRequest {
	s.ClusterCategory = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterCredentialsShrink(v string) *CreateClusterShrinkRequest {
	s.ClusterCredentialsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterCustomConfigurationShrink(v string) *CreateClusterShrinkRequest {
	s.ClusterCustomConfigurationShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterDescription(v string) *CreateClusterShrinkRequest {
	s.ClusterDescription = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterMode(v string) *CreateClusterShrinkRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterName(v string) *CreateClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterVSwitchId(v string) *CreateClusterShrinkRequest {
	s.ClusterVSwitchId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterVpcId(v string) *CreateClusterShrinkRequest {
	s.ClusterVpcId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetDeletionProtection(v bool) *CreateClusterShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetGrowInterval(v int32) *CreateClusterShrinkRequest {
	s.GrowInterval = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetIdleInterval(v int32) *CreateClusterShrinkRequest {
	s.IdleInterval = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetIsEnterpriseSecurityGroup(v bool) *CreateClusterShrinkRequest {
	s.IsEnterpriseSecurityGroup = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetManagerShrink(v string) *CreateClusterShrinkRequest {
	s.ManagerShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetMaxCoreCount(v int32) *CreateClusterShrinkRequest {
	s.MaxCoreCount = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetMaxCount(v int32) *CreateClusterShrinkRequest {
	s.MaxCount = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetQueuesShrink(v string) *CreateClusterShrinkRequest {
	s.QueuesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetResourceGroupId(v string) *CreateClusterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetSecurityGroupId(v string) *CreateClusterShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetSharedStoragesShrink(v string) *CreateClusterShrinkRequest {
	s.SharedStoragesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetTagsShrink(v string) *CreateClusterShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
