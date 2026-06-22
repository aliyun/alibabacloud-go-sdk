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
	// The application configurations. The number of array elements N can range from 1 to 1000.
	ApplicationConfigsShrink *string `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty"`
	// The list of applications. The number of array elements N can range from 1 to 100.
	//
	// This parameter is required.
	ApplicationsShrink *string `json:"Applications,omitempty" xml:"Applications,omitempty"`
	// The array of bootstrap scripts. The number of array elements N can range from 1 to 10.
	BootstrapScriptsShrink *string `json:"BootstrapScripts,omitempty" xml:"BootstrapScripts,omitempty"`
	// A client token to ensure the idempotence of the request. Multiple calls with the same client token return the same result and create only one cluster.
	//
	// example:
	//
	// A7D960FA-6DBA-5E07-8746-A63E3E4D****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster name. The name must be 1 to 128 characters in length. It must start with a letter or a Chinese character. It cannot start with http\\:// or https\\://. It can contain Chinese characters, letters, digits, colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// emrtest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster type. Valid values:
	//
	// - DATALAKE: new data lake.
	//
	// - OLAP: data analytics.
	//
	// - DATAFLOW: real-time data stream.
	//
	// - DATASERVING: DataService Studio.
	//
	// - CUSTOM: custom cluster.
	//
	// - HADOOP: legacy data lake. This value is not recommended. Use the new data lake cluster type instead.
	//
	// If you create an EMR cluster for the first time after 17:00 (UTC+8) on December 19, 2022, you cannot select HADOOP, DATA_SCIENCE, PRESTO, or ZOOKEEPER as the cluster type.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATALAKE
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Specifies whether to enable deletion protection for the cluster. Valid values:
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
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The deployment mode of applications in the cluster. Valid values:
	//
	// - NORMAL (default): non-high availability deployment. The cluster has one master node.
	//
	// - HA: high availability (HA) deployment. This deployment mode requires at least three master nodes.
	//
	// example:
	//
	// HA
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// Emr cluster for ETL
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The node attributes. These are the basic attributes of all ECS nodes in the cluster.
	NodeAttributesShrink *string `json:"NodeAttributes,omitempty" xml:"NodeAttributes,omitempty"`
	// The array of node group configurations. The number of array elements N can range from 1 to 100.
	//
	// This parameter is required.
	NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
	// The billing method. Valid values:
	//
	// - PayAsYouGo: pay-as-you-go.
	//
	// - Subscription: subscription.
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
	// The EMR release version. You can find the EMR release version on the EMR cluster purchase page.
	//
	// This parameter is required.
	//
	// example:
	//
	// EMR-5.16.0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Kerberos security mode of the cluster. Valid values:
	//
	// - NORMAL (default): normal mode. Kerberos is disabled.
	//
	// - KERBEROS: Kerberos mode. Kerberos is enabled.
	//
	// example:
	//
	// NORMAL
	SecurityMode *string `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
	// The subscription configurations. This parameter is required if you set PaymentType to Subscription.
	SubscriptionConfigShrink *string `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	// The tags. The number of array elements N can range from 0 to 20.
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
