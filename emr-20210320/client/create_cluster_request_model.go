// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigs(v []*ApplicationConfig) *CreateClusterRequest
	GetApplicationConfigs() []*ApplicationConfig
	SetApplications(v []*Application) *CreateClusterRequest
	GetApplications() []*Application
	SetBootstrapScripts(v []*Script) *CreateClusterRequest
	GetBootstrapScripts() []*Script
	SetClientToken(v string) *CreateClusterRequest
	GetClientToken() *string
	SetClusterName(v string) *CreateClusterRequest
	GetClusterName() *string
	SetClusterType(v string) *CreateClusterRequest
	GetClusterType() *string
	SetDeletionProtection(v bool) *CreateClusterRequest
	GetDeletionProtection() *bool
	SetDeployMode(v string) *CreateClusterRequest
	GetDeployMode() *string
	SetDescription(v string) *CreateClusterRequest
	GetDescription() *string
	SetNodeAttributes(v *NodeAttributes) *CreateClusterRequest
	GetNodeAttributes() *NodeAttributes
	SetNodeGroups(v []*NodeGroupConfig) *CreateClusterRequest
	GetNodeGroups() []*NodeGroupConfig
	SetPaymentType(v string) *CreateClusterRequest
	GetPaymentType() *string
	SetRegionId(v string) *CreateClusterRequest
	GetRegionId() *string
	SetReleaseVersion(v string) *CreateClusterRequest
	GetReleaseVersion() *string
	SetResourceGroupId(v string) *CreateClusterRequest
	GetResourceGroupId() *string
	SetSecurityMode(v string) *CreateClusterRequest
	GetSecurityMode() *string
	SetSubscriptionConfig(v *SubscriptionConfig) *CreateClusterRequest
	GetSubscriptionConfig() *SubscriptionConfig
	SetTags(v []*Tag) *CreateClusterRequest
	GetTags() []*Tag
}

type CreateClusterRequest struct {
	// The service configurations. Number of elements in the array: 1 to 1,000.
	ApplicationConfigs []*ApplicationConfig `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
	// The services. Number of elements in the array: 1 to 100.
	//
	// This parameter is required.
	Applications []*Application `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The array of bootstrap scripts. Number of elements in the array: 1 to 10.
	BootstrapScripts []*Script `json:"BootstrapScripts,omitempty" xml:"BootstrapScripts,omitempty" type:"Repeated"`
	// The idempotent client token. If you call the same ClientToken multiple times, the returned results are the same. Only one cluster can be created with the same ClientToken.
	//
	// example:
	//
	// A7D960FA-6DBA-5E07-8746-A63E3E4D****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the cluster. The name must be 1 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// emrtest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- DATALAKE: data lake
	//
	// 	- OLAP: online analytical processing (OLAP)
	//
	// 	- DATAFLOW: Dataflow
	//
	// 	- DATASERVING: DataServing
	//
	// 	- CUSTOM: a custom hybrid cluster.
	//
	// 	- HADOOP: the old data lake. We recommend that you use the new data lake.
	//
	// If you create an EMR cluster for the first time after 17:00 (UTC +8) on December 19, 2022, you cannot select the HADOOP, DATA_SCIENCE, PRESTO, or ZOOKEEPER cluster type.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATALAKE
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Specifies whether to enable release protection for the cluster. Valid values:
	//
	// 	- true: enables release protection for the cluster.
	//
	// 	- false: disables release protection for the cluster.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The deployment mode of master nodes in the cluster. Valid values:
	//
	// 	- NORMAL: regular mode. This is the default value. A cluster that contains only one master node is created.
	//
	// 	- HA: high availability (HA) mode. A cluster that contains three master nodes is created.
	//
	// example:
	//
	// HA
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// Emr cluster for ETL
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The node attributes. The basic attributes of all ECS nodes in the cluster.
	//
	// This parameter is required.
	NodeAttributes *NodeAttributes `json:"NodeAttributes,omitempty" xml:"NodeAttributes,omitempty"`
	// The array of configurations of the node groups. Number of elements in the array: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	NodeGroups []*NodeGroupConfig `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// The billing cycle of the instance. Valid values:
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// 	- Subscription: subscription
	//
	// Default value: PayAsYouGo.
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The EMR version. You can query available E-MapReduce (EMR) versions in the EMR console.
	//
	// This parameter is required.
	//
	// example:
	//
	// EMR-5.8.0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// The ID of the resource group to which to assign the ENI.
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security mode of the cluster. Valid values:
	//
	// 	- NORMAL: disables Kerberos authentication for the cluster. This is the default value.
	//
	// 	- KERBEROS: enables Kerberos authentication for the cluster.
	//
	// example:
	//
	// NORMAL
	SecurityMode *string `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
	// The subscription configurations. This parameter takes effect only if you set the PaymentType parameter to Subscription.
	SubscriptionConfig *SubscriptionConfig `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	// The tag. Number of elements in the array: 0 to 20.
	//
	// example:
	//
	// A7D960FA-6DBA-5E07-8746-A63E3E4D****
	Tags []*Tag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetApplicationConfigs() []*ApplicationConfig {
	return s.ApplicationConfigs
}

func (s *CreateClusterRequest) GetApplications() []*Application {
	return s.Applications
}

func (s *CreateClusterRequest) GetBootstrapScripts() []*Script {
	return s.BootstrapScripts
}

func (s *CreateClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreateClusterRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateClusterRequest) GetDeployMode() *string {
	return s.DeployMode
}

func (s *CreateClusterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateClusterRequest) GetNodeAttributes() *NodeAttributes {
	return s.NodeAttributes
}

func (s *CreateClusterRequest) GetNodeGroups() []*NodeGroupConfig {
	return s.NodeGroups
}

func (s *CreateClusterRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateClusterRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *CreateClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClusterRequest) GetSecurityMode() *string {
	return s.SecurityMode
}

func (s *CreateClusterRequest) GetSubscriptionConfig() *SubscriptionConfig {
	return s.SubscriptionConfig
}

func (s *CreateClusterRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *CreateClusterRequest) SetApplicationConfigs(v []*ApplicationConfig) *CreateClusterRequest {
	s.ApplicationConfigs = v
	return s
}

func (s *CreateClusterRequest) SetApplications(v []*Application) *CreateClusterRequest {
	s.Applications = v
	return s
}

func (s *CreateClusterRequest) SetBootstrapScripts(v []*Script) *CreateClusterRequest {
	s.BootstrapScripts = v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetDeletionProtection(v bool) *CreateClusterRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateClusterRequest) SetDeployMode(v string) *CreateClusterRequest {
	s.DeployMode = &v
	return s
}

func (s *CreateClusterRequest) SetDescription(v string) *CreateClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterRequest) SetNodeAttributes(v *NodeAttributes) *CreateClusterRequest {
	s.NodeAttributes = v
	return s
}

func (s *CreateClusterRequest) SetNodeGroups(v []*NodeGroupConfig) *CreateClusterRequest {
	s.NodeGroups = v
	return s
}

func (s *CreateClusterRequest) SetPaymentType(v string) *CreateClusterRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetReleaseVersion(v string) *CreateClusterRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityMode(v string) *CreateClusterRequest {
	s.SecurityMode = &v
	return s
}

func (s *CreateClusterRequest) SetSubscriptionConfig(v *SubscriptionConfig) *CreateClusterRequest {
	s.SubscriptionConfig = v
	return s
}

func (s *CreateClusterRequest) SetTags(v []*Tag) *CreateClusterRequest {
	s.Tags = v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	return dara.Validate(s)
}
