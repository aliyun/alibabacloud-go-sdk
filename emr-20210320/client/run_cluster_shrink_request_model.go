// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigsShrink(v string) *RunClusterShrinkRequest
	GetApplicationConfigsShrink() *string
	SetApplicationsShrink(v string) *RunClusterShrinkRequest
	GetApplicationsShrink() *string
	SetBootstrapScriptsShrink(v string) *RunClusterShrinkRequest
	GetBootstrapScriptsShrink() *string
	SetClientToken(v string) *RunClusterShrinkRequest
	GetClientToken() *string
	SetClusterName(v string) *RunClusterShrinkRequest
	GetClusterName() *string
	SetClusterType(v string) *RunClusterShrinkRequest
	GetClusterType() *string
	SetDeletionProtection(v bool) *RunClusterShrinkRequest
	GetDeletionProtection() *bool
	SetDeployMode(v string) *RunClusterShrinkRequest
	GetDeployMode() *string
	SetDescription(v string) *RunClusterShrinkRequest
	GetDescription() *string
	SetNodeAttributesShrink(v string) *RunClusterShrinkRequest
	GetNodeAttributesShrink() *string
	SetNodeGroupsShrink(v string) *RunClusterShrinkRequest
	GetNodeGroupsShrink() *string
	SetPaymentType(v string) *RunClusterShrinkRequest
	GetPaymentType() *string
	SetPromotionsShrink(v string) *RunClusterShrinkRequest
	GetPromotionsShrink() *string
	SetRegionId(v string) *RunClusterShrinkRequest
	GetRegionId() *string
	SetReleaseVersion(v string) *RunClusterShrinkRequest
	GetReleaseVersion() *string
	SetResourceGroupId(v string) *RunClusterShrinkRequest
	GetResourceGroupId() *string
	SetSecurityMode(v string) *RunClusterShrinkRequest
	GetSecurityMode() *string
	SetSubscriptionConfigShrink(v string) *RunClusterShrinkRequest
	GetSubscriptionConfigShrink() *string
	SetTagsShrink(v string) *RunClusterShrinkRequest
	GetTagsShrink() *string
}

type RunClusterShrinkRequest struct {
	// The application configurations. Number of elements in the array: 1 to 1000.
	ApplicationConfigsShrink *string `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty"`
	// The list of services. Number of elements in the array: 1 to 100.
	//
	// This parameter is required.
	ApplicationsShrink *string `json:"Applications,omitempty" xml:"Applications,omitempty"`
	// The array of bootstrap scripts. Number of elements in the array: 1 to 10.
	BootstrapScriptsShrink *string `json:"BootstrapScripts,omitempty" xml:"BootstrapScripts,omitempty"`
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
	NodeAttributesShrink *string `json:"NodeAttributes,omitempty" xml:"NodeAttributes,omitempty"`
	// The array of configurations of the node groups. Number of elements in the array: 1 to 100.
	//
	// This parameter is required.
	NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
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
	PaymentType      *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	PromotionsShrink *string `json:"Promotions,omitempty" xml:"Promotions,omitempty"`
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
	SubscriptionConfigShrink *string `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	// The tag. Number of elements in the array: 0 to 20.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s RunClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunClusterShrinkRequest) GetApplicationConfigsShrink() *string {
	return s.ApplicationConfigsShrink
}

func (s *RunClusterShrinkRequest) GetApplicationsShrink() *string {
	return s.ApplicationsShrink
}

func (s *RunClusterShrinkRequest) GetBootstrapScriptsShrink() *string {
	return s.BootstrapScriptsShrink
}

func (s *RunClusterShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunClusterShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *RunClusterShrinkRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *RunClusterShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *RunClusterShrinkRequest) GetDeployMode() *string {
	return s.DeployMode
}

func (s *RunClusterShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *RunClusterShrinkRequest) GetNodeAttributesShrink() *string {
	return s.NodeAttributesShrink
}

func (s *RunClusterShrinkRequest) GetNodeGroupsShrink() *string {
	return s.NodeGroupsShrink
}

func (s *RunClusterShrinkRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *RunClusterShrinkRequest) GetPromotionsShrink() *string {
	return s.PromotionsShrink
}

func (s *RunClusterShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunClusterShrinkRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *RunClusterShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunClusterShrinkRequest) GetSecurityMode() *string {
	return s.SecurityMode
}

func (s *RunClusterShrinkRequest) GetSubscriptionConfigShrink() *string {
	return s.SubscriptionConfigShrink
}

func (s *RunClusterShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *RunClusterShrinkRequest) SetApplicationConfigsShrink(v string) *RunClusterShrinkRequest {
	s.ApplicationConfigsShrink = &v
	return s
}

func (s *RunClusterShrinkRequest) SetApplicationsShrink(v string) *RunClusterShrinkRequest {
	s.ApplicationsShrink = &v
	return s
}

func (s *RunClusterShrinkRequest) SetBootstrapScriptsShrink(v string) *RunClusterShrinkRequest {
	s.BootstrapScriptsShrink = &v
	return s
}

func (s *RunClusterShrinkRequest) SetClientToken(v string) *RunClusterShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RunClusterShrinkRequest) SetClusterName(v string) *RunClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *RunClusterShrinkRequest) SetClusterType(v string) *RunClusterShrinkRequest {
	s.ClusterType = &v
	return s
}

func (s *RunClusterShrinkRequest) SetDeletionProtection(v bool) *RunClusterShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *RunClusterShrinkRequest) SetDeployMode(v string) *RunClusterShrinkRequest {
	s.DeployMode = &v
	return s
}

func (s *RunClusterShrinkRequest) SetDescription(v string) *RunClusterShrinkRequest {
	s.Description = &v
	return s
}

func (s *RunClusterShrinkRequest) SetNodeAttributesShrink(v string) *RunClusterShrinkRequest {
	s.NodeAttributesShrink = &v
	return s
}

func (s *RunClusterShrinkRequest) SetNodeGroupsShrink(v string) *RunClusterShrinkRequest {
	s.NodeGroupsShrink = &v
	return s
}

func (s *RunClusterShrinkRequest) SetPaymentType(v string) *RunClusterShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *RunClusterShrinkRequest) SetPromotionsShrink(v string) *RunClusterShrinkRequest {
	s.PromotionsShrink = &v
	return s
}

func (s *RunClusterShrinkRequest) SetRegionId(v string) *RunClusterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunClusterShrinkRequest) SetReleaseVersion(v string) *RunClusterShrinkRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *RunClusterShrinkRequest) SetResourceGroupId(v string) *RunClusterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunClusterShrinkRequest) SetSecurityMode(v string) *RunClusterShrinkRequest {
	s.SecurityMode = &v
	return s
}

func (s *RunClusterShrinkRequest) SetSubscriptionConfigShrink(v string) *RunClusterShrinkRequest {
	s.SubscriptionConfigShrink = &v
	return s
}

func (s *RunClusterShrinkRequest) SetTagsShrink(v string) *RunClusterShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *RunClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
