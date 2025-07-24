// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigs(v []*ApplicationConfig) *RunClusterRequest
	GetApplicationConfigs() []*ApplicationConfig
	SetApplications(v []*Application) *RunClusterRequest
	GetApplications() []*Application
	SetBootstrapScripts(v []*Script) *RunClusterRequest
	GetBootstrapScripts() []*Script
	SetClientToken(v string) *RunClusterRequest
	GetClientToken() *string
	SetClusterName(v string) *RunClusterRequest
	GetClusterName() *string
	SetClusterType(v string) *RunClusterRequest
	GetClusterType() *string
	SetDeletionProtection(v bool) *RunClusterRequest
	GetDeletionProtection() *bool
	SetDeployMode(v string) *RunClusterRequest
	GetDeployMode() *string
	SetDescription(v string) *RunClusterRequest
	GetDescription() *string
	SetNodeAttributes(v *NodeAttributes) *RunClusterRequest
	GetNodeAttributes() *NodeAttributes
	SetNodeGroups(v []*NodeGroupConfig) *RunClusterRequest
	GetNodeGroups() []*NodeGroupConfig
	SetPaymentType(v string) *RunClusterRequest
	GetPaymentType() *string
	SetPromotions(v []*Promotion) *RunClusterRequest
	GetPromotions() []*Promotion
	SetRegionId(v string) *RunClusterRequest
	GetRegionId() *string
	SetReleaseVersion(v string) *RunClusterRequest
	GetReleaseVersion() *string
	SetResourceGroupId(v string) *RunClusterRequest
	GetResourceGroupId() *string
	SetSecurityMode(v string) *RunClusterRequest
	GetSecurityMode() *string
	SetSubscriptionConfig(v *SubscriptionConfig) *RunClusterRequest
	GetSubscriptionConfig() *SubscriptionConfig
	SetTags(v []*Tag) *RunClusterRequest
	GetTags() []*Tag
}

type RunClusterRequest struct {
	// The application configurations. Number of elements in the array: 1 to 1000.
	ApplicationConfigs []*ApplicationConfig `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
	// The list of services. Number of elements in the array: 1 to 100.
	//
	// This parameter is required.
	Applications []*Application `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The array of bootstrap scripts. Number of elements in the array: 1 to 10.
	BootstrapScripts []*Script `json:"BootstrapScripts,omitempty" xml:"BootstrapScripts,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. The same ClientToken value for multiple calls to the RunCluster operation results in identical responses. Only one cluster can be created by using the same ClientToken value.
	//
	// example:
	//
	// A7D960FA-6DBA-5E07-8746-A63E3E4D****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster name. The name must be 1 to 128 characters in length. The name must start with a letter but cannot start with http:// or https://. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// emrtest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- DATALAKE
	//
	// 	- OLAP
	//
	// 	- DATAFLOW
	//
	// 	- DATASERVING
	//
	// 	- CUSTOM
	//
	// 	- HADOOP: We recommend that you set this parameter to DATALAKE rather than HADOOP.
	//
	// If the first time you create an EMR cluster is after 17:00 (UTC+8) on December 19, 2022, you cannot create a Hadoop, Data Science, Presto, or ZooKeeper cluster.
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
	// 	- HA: high availability mode. A cluster that contains at least three master nodes is created.
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
	NodeAttributes *NodeAttributes `json:"NodeAttributes,omitempty" xml:"NodeAttributes,omitempty"`
	// The array of configurations of the node groups. Number of elements in the array: 1 to 100.
	//
	// This parameter is required.
	NodeGroups []*NodeGroupConfig `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// The billing method of the cluster. Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Subscription
	//
	// Default value: PayAsYouGo.
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string      `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	Promotions  []*Promotion `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The EMR version. You can query available EMR versions in the EMR console.
	//
	// This parameter is required.
	//
	// example:
	//
	// EMR-5.16.0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security mode of the cluster. Valid values:
	//
	// 	- NORMAL: regular mode. Kerberos authentication is disabled. This is the default value.
	//
	// 	- KERBEROS: Kerberos mode. Kerberos authentication is enabled.
	//
	// example:
	//
	// NORMAL
	SecurityMode *string `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
	// The subscription configurations. This parameter is required when the PaymentType parameter is set to Subscription.
	SubscriptionConfig *SubscriptionConfig `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	// The tag. Number of elements in the array: 0 to 20.
	Tags []*Tag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RunClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s RunClusterRequest) GoString() string {
	return s.String()
}

func (s *RunClusterRequest) GetApplicationConfigs() []*ApplicationConfig {
	return s.ApplicationConfigs
}

func (s *RunClusterRequest) GetApplications() []*Application {
	return s.Applications
}

func (s *RunClusterRequest) GetBootstrapScripts() []*Script {
	return s.BootstrapScripts
}

func (s *RunClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *RunClusterRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *RunClusterRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *RunClusterRequest) GetDeployMode() *string {
	return s.DeployMode
}

func (s *RunClusterRequest) GetDescription() *string {
	return s.Description
}

func (s *RunClusterRequest) GetNodeAttributes() *NodeAttributes {
	return s.NodeAttributes
}

func (s *RunClusterRequest) GetNodeGroups() []*NodeGroupConfig {
	return s.NodeGroups
}

func (s *RunClusterRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *RunClusterRequest) GetPromotions() []*Promotion {
	return s.Promotions
}

func (s *RunClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunClusterRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *RunClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunClusterRequest) GetSecurityMode() *string {
	return s.SecurityMode
}

func (s *RunClusterRequest) GetSubscriptionConfig() *SubscriptionConfig {
	return s.SubscriptionConfig
}

func (s *RunClusterRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *RunClusterRequest) SetApplicationConfigs(v []*ApplicationConfig) *RunClusterRequest {
	s.ApplicationConfigs = v
	return s
}

func (s *RunClusterRequest) SetApplications(v []*Application) *RunClusterRequest {
	s.Applications = v
	return s
}

func (s *RunClusterRequest) SetBootstrapScripts(v []*Script) *RunClusterRequest {
	s.BootstrapScripts = v
	return s
}

func (s *RunClusterRequest) SetClientToken(v string) *RunClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *RunClusterRequest) SetClusterName(v string) *RunClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *RunClusterRequest) SetClusterType(v string) *RunClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *RunClusterRequest) SetDeletionProtection(v bool) *RunClusterRequest {
	s.DeletionProtection = &v
	return s
}

func (s *RunClusterRequest) SetDeployMode(v string) *RunClusterRequest {
	s.DeployMode = &v
	return s
}

func (s *RunClusterRequest) SetDescription(v string) *RunClusterRequest {
	s.Description = &v
	return s
}

func (s *RunClusterRequest) SetNodeAttributes(v *NodeAttributes) *RunClusterRequest {
	s.NodeAttributes = v
	return s
}

func (s *RunClusterRequest) SetNodeGroups(v []*NodeGroupConfig) *RunClusterRequest {
	s.NodeGroups = v
	return s
}

func (s *RunClusterRequest) SetPaymentType(v string) *RunClusterRequest {
	s.PaymentType = &v
	return s
}

func (s *RunClusterRequest) SetPromotions(v []*Promotion) *RunClusterRequest {
	s.Promotions = v
	return s
}

func (s *RunClusterRequest) SetRegionId(v string) *RunClusterRequest {
	s.RegionId = &v
	return s
}

func (s *RunClusterRequest) SetReleaseVersion(v string) *RunClusterRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *RunClusterRequest) SetResourceGroupId(v string) *RunClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunClusterRequest) SetSecurityMode(v string) *RunClusterRequest {
	s.SecurityMode = &v
	return s
}

func (s *RunClusterRequest) SetSubscriptionConfig(v *SubscriptionConfig) *RunClusterRequest {
	s.SubscriptionConfig = v
	return s
}

func (s *RunClusterRequest) SetTags(v []*Tag) *RunClusterRequest {
	s.Tags = v
	return s
}

func (s *RunClusterRequest) Validate() error {
	return dara.Validate(s)
}
