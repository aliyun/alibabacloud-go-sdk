// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCloneMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterCloneMeta(v *GetClusterCloneMetaResponseBodyClusterCloneMeta) *GetClusterCloneMetaResponseBody
	GetClusterCloneMeta() *GetClusterCloneMetaResponseBodyClusterCloneMeta
	SetRequestId(v string) *GetClusterCloneMetaResponseBody
	GetRequestId() *string
}

type GetClusterCloneMetaResponseBody struct {
	// Cluster clone metadata.
	ClusterCloneMeta *GetClusterCloneMetaResponseBodyClusterCloneMeta `json:"ClusterCloneMeta,omitempty" xml:"ClusterCloneMeta,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterCloneMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCloneMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterCloneMetaResponseBody) GetClusterCloneMeta() *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	return s.ClusterCloneMeta
}

func (s *GetClusterCloneMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterCloneMetaResponseBody) SetClusterCloneMeta(v *GetClusterCloneMetaResponseBodyClusterCloneMeta) *GetClusterCloneMetaResponseBody {
	s.ClusterCloneMeta = v
	return s
}

func (s *GetClusterCloneMetaResponseBody) SetRequestId(v string) *GetClusterCloneMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterCloneMetaResponseBody) Validate() error {
	if s.ClusterCloneMeta != nil {
		if err := s.ClusterCloneMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterCloneMetaResponseBodyClusterCloneMeta struct {
	// The modified configuration items.
	ApplicationConfigs []*ApplicationConfig `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
	// The cluster applications.
	Applications []*Application `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// An array of bootstrap scripts. The number of elements in the array: 1 to 10.
	BootstrapScripts []*Script `json:"BootstrapScripts,omitempty" xml:"BootstrapScripts,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// c-b933c5aac7f7****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// emrtest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- STARTING
	//
	// 	- START_FAILED
	//
	// 	- BOOTSTRAPPING
	//
	// 	- RUNNING
	//
	// 	- TERMINATING
	//
	// 	- TERMINATED
	//
	// 	- TERMINATED_WITH_ERRORS
	//
	// 	- TERMINATE_FAILED
	//
	// example:
	//
	// RUNNING
	ClusterState *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	// The cluster type. Valid values:
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
	// 	- HADOOP
	//
	// example:
	//
	// DATALAKE
	ClusterType       *string            `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CollationTimeZone *CollationTimeZone `json:"CollationTimeZone,omitempty" xml:"CollationTimeZone,omitempty"`
	// Indicates whether release protection is enabled for the cluster. Valid values:
	//
	// 	- true: Release protection is enabled for the cluster.
	//
	// 	- false: Release protection is disabled for the cluster.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The deployment mode of master nodes in the cluster. Valid values:
	//
	// 	- NORMAL: regular mode.
	//
	// 	- HA: high availability mode.
	//
	// example:
	//
	// HA
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The EMR service role.
	//
	// example:
	//
	// AliyunEMRDefaultRole
	EmrDefaultRole *string `json:"EmrDefaultRole,omitempty" xml:"EmrDefaultRole,omitempty"`
	// Indicates whether the service configurations of a Hadoop cluster that you made during cluster creation can be cloned. Valid values:
	//
	// 	- False
	//
	// 	- True
	//
	// example:
	//
	// True
	ExistCloneConfig *bool `json:"ExistCloneConfig,omitempty" xml:"ExistCloneConfig,omitempty"`
	// The node attributes.
	NodeAttributes *NodeAttributes `json:"NodeAttributes,omitempty" xml:"NodeAttributes,omitempty"`
	// The node groups. The number of elements in the array: 1 to 100.
	NodeGroups []*NodeGroup `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// The billing method of the cluster. Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Subscription
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The EMR version.
	//
	// example:
	//
	// EMR-5.6.0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The auto scaling policies of each node group in the cluster.
	ScalingPolicies []*GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies `json:"ScalingPolicies,omitempty" xml:"ScalingPolicies,omitempty" type:"Repeated"`
	// The security mode of the cluster. Valid values:
	//
	// 	- NORMAL: regular mode. Kerberos is not enabled.
	//
	// 	- KERBEROS: Kerberos mode. Kerberos is enabled.
	//
	// example:
	//
	// NORMAL
	SecurityMode *string `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
	// The subscription configurations.
	SubscriptionConfig *SubscriptionConfig `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	// The list of cluster tags.
	Tags []*Tag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetClusterCloneMetaResponseBodyClusterCloneMeta) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCloneMetaResponseBodyClusterCloneMeta) GoString() string {
	return s.String()
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetApplicationConfigs() []*ApplicationConfig {
	return s.ApplicationConfigs
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetApplications() []*Application {
	return s.Applications
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetBootstrapScripts() []*Script {
	return s.BootstrapScripts
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetClusterState() *string {
	return s.ClusterState
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetClusterType() *string {
	return s.ClusterType
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetCollationTimeZone() *CollationTimeZone {
	return s.CollationTimeZone
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetDeployMode() *string {
	return s.DeployMode
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetEmrDefaultRole() *string {
	return s.EmrDefaultRole
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetExistCloneConfig() *bool {
	return s.ExistCloneConfig
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetNodeAttributes() *NodeAttributes {
	return s.NodeAttributes
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetNodeGroups() []*NodeGroup {
	return s.NodeGroups
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetScalingPolicies() []*GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies {
	return s.ScalingPolicies
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetSecurityMode() *string {
	return s.SecurityMode
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetSubscriptionConfig() *SubscriptionConfig {
	return s.SubscriptionConfig
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) GetTags() []*Tag {
	return s.Tags
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetApplicationConfigs(v []*ApplicationConfig) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.ApplicationConfigs = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetApplications(v []*Application) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.Applications = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetBootstrapScripts(v []*Script) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.BootstrapScripts = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetClusterId(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.ClusterId = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetClusterName(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.ClusterName = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetClusterState(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.ClusterState = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetClusterType(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.ClusterType = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetCollationTimeZone(v *CollationTimeZone) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.CollationTimeZone = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetDeletionProtection(v bool) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.DeletionProtection = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetDeployMode(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.DeployMode = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetEmrDefaultRole(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.EmrDefaultRole = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetExistCloneConfig(v bool) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.ExistCloneConfig = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetNodeAttributes(v *NodeAttributes) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.NodeAttributes = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetNodeGroups(v []*NodeGroup) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.NodeGroups = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetPaymentType(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.PaymentType = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetRegionId(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.RegionId = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetReleaseVersion(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.ReleaseVersion = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetResourceGroupId(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.ResourceGroupId = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetScalingPolicies(v []*GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.ScalingPolicies = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetSecurityMode(v string) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.SecurityMode = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetSubscriptionConfig(v *SubscriptionConfig) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.SubscriptionConfig = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) SetTags(v []*Tag) *GetClusterCloneMetaResponseBodyClusterCloneMeta {
	s.Tags = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMeta) Validate() error {
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
	if s.CollationTimeZone != nil {
		if err := s.CollationTimeZone.Validate(); err != nil {
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
	if s.ScalingPolicies != nil {
		for _, item := range s.ScalingPolicies {
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

type GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies struct {
	// The cluster ID.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The maximum and minimum number of nodes in the node group.
	Constraints *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty" type:"Struct"`
	// The node group ID.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The name of the node group.
	//
	// example:
	//
	// emr-etltask
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The ID of the auto scaling policy.
	//
	// example:
	//
	// asp-asduwe23znl***
	ScalingPolicyId *string `json:"ScalingPolicyId,omitempty" xml:"ScalingPolicyId,omitempty"`
	// The type of the auto scaling policy.
	ScalingPolicyType *string `json:"ScalingPolicyType,omitempty" xml:"ScalingPolicyType,omitempty"`
	// The auto scaling rules.
	ScalingRules []*GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules `json:"ScalingRules,omitempty" xml:"ScalingRules,omitempty" type:"Repeated"`
}

func (s GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) GoString() string {
	return s.String()
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) GetConstraints() *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints {
	return s.Constraints
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) GetScalingPolicyId() *string {
	return s.ScalingPolicyId
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) GetScalingPolicyType() *string {
	return s.ScalingPolicyType
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) GetScalingRules() []*GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules {
	return s.ScalingRules
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) SetClusterId(v string) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies {
	s.ClusterId = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) SetConstraints(v *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies {
	s.Constraints = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) SetNodeGroupId(v string) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies {
	s.NodeGroupId = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) SetNodeGroupName(v string) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies {
	s.NodeGroupName = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) SetScalingPolicyId(v string) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies {
	s.ScalingPolicyId = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) SetScalingPolicyType(v string) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies {
	s.ScalingPolicyType = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) SetScalingRules(v []*GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies {
	s.ScalingRules = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPolicies) Validate() error {
	if s.Constraints != nil {
		if err := s.Constraints.Validate(); err != nil {
			return err
		}
	}
	if s.ScalingRules != nil {
		for _, item := range s.ScalingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints struct {
	// The maximum number of nodes in the node group. Default value: 2000.
	//
	// example:
	//
	// 200
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The maximum number of pay-as-you-go nodes in the node group.
	//
	// example:
	//
	// 200
	MaxOnDemandCapacity *int32 `json:"MaxOnDemandCapacity,omitempty" xml:"MaxOnDemandCapacity,omitempty"`
	// The minimum number of nodes in the node group. Default value: 0.
	//
	// example:
	//
	// 50
	MinCapacity *int32 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints) GoString() string {
	return s.String()
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints) GetMaxOnDemandCapacity() *int32 {
	return s.MaxOnDemandCapacity
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints) SetMaxCapacity(v int32) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints {
	s.MaxCapacity = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints) SetMaxOnDemandCapacity(v int32) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints {
	s.MaxOnDemandCapacity = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints) SetMinCapacity(v int32) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints {
	s.MinCapacity = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesConstraints) Validate() error {
	return dara.Validate(s)
}

type GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules struct {
	// The scaling type. This parameter is required. Valid values:
	//
	// 	- SCALE_OUT
	//
	// 	- SCALE_IN
	//
	// example:
	//
	// SCALE_OUT
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	// The adjustment value of the auto scaling rule. This parameter is required. The parameter value must be a positive integer, which indicates the number of instances to be added or removed.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// The description of load-based scaling.
	MetricsTrigger *MetricsTrigger `json:"MetricsTrigger,omitempty" xml:"MetricsTrigger,omitempty"`
	// The name of the auto scaling rule.
	//
	// example:
	//
	// scaling-out-memory
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The description of time-based scaling.
	TimeTrigger *TimeTrigger `json:"TimeTrigger,omitempty" xml:"TimeTrigger,omitempty"`
	// The trigger mode of the auto scaling rule. This parameter is required. Valid values:
	//
	// 	- TIME_TRIGGER: time-based scaling.
	//
	// 	- METRICS_TRIGGER: load-based scaling.
	//
	// example:
	//
	// TIME_TRIGGER
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) GoString() string {
	return s.String()
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) GetActivityType() *string {
	return s.ActivityType
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) GetMetricsTrigger() *MetricsTrigger {
	return s.MetricsTrigger
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) GetRuleName() *string {
	return s.RuleName
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) GetTimeTrigger() *TimeTrigger {
	return s.TimeTrigger
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) SetActivityType(v string) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules {
	s.ActivityType = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) SetAdjustmentValue(v int32) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules {
	s.AdjustmentValue = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) SetMetricsTrigger(v *MetricsTrigger) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules {
	s.MetricsTrigger = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) SetRuleName(v string) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules {
	s.RuleName = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) SetTimeTrigger(v *TimeTrigger) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules {
	s.TimeTrigger = v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) SetTriggerType(v string) *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules {
	s.TriggerType = &v
	return s
}

func (s *GetClusterCloneMetaResponseBodyClusterCloneMetaScalingPoliciesScalingRules) Validate() error {
	if s.MetricsTrigger != nil {
		if err := s.MetricsTrigger.Validate(); err != nil {
			return err
		}
	}
	if s.TimeTrigger != nil {
		if err := s.TimeTrigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}
