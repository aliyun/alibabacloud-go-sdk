// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCluster interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *Cluster
	GetClusterId() *string
	SetClusterName(v string) *Cluster
	GetClusterName() *string
	SetClusterState(v string) *Cluster
	GetClusterState() *string
	SetClusterType(v string) *Cluster
	GetClusterType() *string
	SetCreateTime(v int64) *Cluster
	GetCreateTime() *int64
	SetDeletionProtection(v bool) *Cluster
	GetDeletionProtection() *bool
	SetDeployMode(v string) *Cluster
	GetDeployMode() *string
	SetDescription(v string) *Cluster
	GetDescription() *string
	SetEmrDefaultRole(v string) *Cluster
	GetEmrDefaultRole() *string
	SetEndTime(v int64) *Cluster
	GetEndTime() *int64
	SetExpireTime(v int64) *Cluster
	GetExpireTime() *int64
	SetNodeAttributes(v *NodeAttributes) *Cluster
	GetNodeAttributes() *NodeAttributes
	SetPaymentType(v string) *Cluster
	GetPaymentType() *string
	SetReadyTime(v int64) *Cluster
	GetReadyTime() *int64
	SetRegionId(v string) *Cluster
	GetRegionId() *string
	SetReleaseVersion(v string) *Cluster
	GetReleaseVersion() *string
	SetResourceGroupId(v string) *Cluster
	GetResourceGroupId() *string
	SetSecurityMode(v string) *Cluster
	GetSecurityMode() *string
	SetStateChangeReason(v *ClusterStateChangeReason) *Cluster
	GetStateChangeReason() *ClusterStateChangeReason
	SetSubscriptionConfig(v *SubscriptionConfig) *Cluster
	GetSubscriptionConfig() *SubscriptionConfig
	SetTags(v []*Tag) *Cluster
	GetTags() []*Tag
}

type Cluster struct {
	// The cluster ID.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// emrtest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster state. Valid values:
	//
	// - `STARTING`: The cluster is starting.
	//
	// - `START_FAILED`: The cluster failed to start.
	//
	// - `BOOTSTRAPPING`: The cluster is running bootstrap actions.
	//
	// - `RUNNING`: The cluster is running.
	//
	// - `TERMINATING`: The cluster is terminating.
	//
	// - `TERMINATED`: The cluster is terminated.
	//
	// - `TERMINATED_WITH_ERRORS`: The cluster terminated due to errors.
	//
	// - `TERMINATE_FAILED`: The cluster failed to terminate.
	//
	// example:
	//
	// RUNNING
	ClusterState *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	// The cluster type. Valid values:
	//
	// - `DATALAKE`: New data lake.
	//
	// - `OLAP`: Data analysis.
	//
	// - `DATAFLOW`: Real-time data flow.
	//
	// - `DATASERVING`: Data serving.
	//
	// - `CUSTOM`: Custom cluster.
	//
	// - `HADOOP`: Legacy data lake.
	//
	// example:
	//
	// DATALAKE
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The time when the cluster was created. The time is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1592837465784
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether deletion protection is enabled for the cluster.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The application deployment mode. Valid values:
	//
	// - `NORMAL`: Standard deployment.
	//
	// - `HA`: High-availability deployment.
	//
	// example:
	//
	// HA
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// EMR cluster
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The default role for E-MapReduce.
	//
	// example:
	//
	// AliyunEMRDefaultRole
	EmrDefaultRole *string `json:"EmrDefaultRole,omitempty" xml:"EmrDefaultRole,omitempty"`
	// The time when the cluster was deleted. The time is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1592837465784
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the cluster is scheduled to expire. The time is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1592837465784
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The node attributes.
	NodeAttributes *NodeAttributes `json:"NodeAttributes,omitempty" xml:"NodeAttributes,omitempty"`
	// The billing method. Valid values:
	//
	// - `PayAsYouGo`: Pay-as-you-go.
	//
	// - `Subscription`: Subscription.
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The time when the cluster became ready. The time is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1592837465784
	ReadyTime *int64 `json:"ReadyTime,omitempty" xml:"ReadyTime,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The release version of E-MapReduce.
	//
	// example:
	//
	// EMR-5.8.0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security mode of the cluster. Valid values:
	//
	// - `NORMAL`: Kerberos is disabled.
	//
	// - `KERBEROS`: Kerberos is enabled.
	//
	// example:
	//
	// NORMAL
	SecurityMode *string `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
	// The reason for the cluster state change.
	StateChangeReason *ClusterStateChangeReason `json:"StateChangeReason,omitempty" xml:"StateChangeReason,omitempty"`
	// The subscription configuration.
	SubscriptionConfig *SubscriptionConfig `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	// The cluster tags.
	Tags []*Tag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s Cluster) String() string {
	return dara.Prettify(s)
}

func (s Cluster) GoString() string {
	return s.String()
}

func (s *Cluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *Cluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *Cluster) GetClusterState() *string {
	return s.ClusterState
}

func (s *Cluster) GetClusterType() *string {
	return s.ClusterType
}

func (s *Cluster) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Cluster) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *Cluster) GetDeployMode() *string {
	return s.DeployMode
}

func (s *Cluster) GetDescription() *string {
	return s.Description
}

func (s *Cluster) GetEmrDefaultRole() *string {
	return s.EmrDefaultRole
}

func (s *Cluster) GetEndTime() *int64 {
	return s.EndTime
}

func (s *Cluster) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *Cluster) GetNodeAttributes() *NodeAttributes {
	return s.NodeAttributes
}

func (s *Cluster) GetPaymentType() *string {
	return s.PaymentType
}

func (s *Cluster) GetReadyTime() *int64 {
	return s.ReadyTime
}

func (s *Cluster) GetRegionId() *string {
	return s.RegionId
}

func (s *Cluster) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *Cluster) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *Cluster) GetSecurityMode() *string {
	return s.SecurityMode
}

func (s *Cluster) GetStateChangeReason() *ClusterStateChangeReason {
	return s.StateChangeReason
}

func (s *Cluster) GetSubscriptionConfig() *SubscriptionConfig {
	return s.SubscriptionConfig
}

func (s *Cluster) GetTags() []*Tag {
	return s.Tags
}

func (s *Cluster) SetClusterId(v string) *Cluster {
	s.ClusterId = &v
	return s
}

func (s *Cluster) SetClusterName(v string) *Cluster {
	s.ClusterName = &v
	return s
}

func (s *Cluster) SetClusterState(v string) *Cluster {
	s.ClusterState = &v
	return s
}

func (s *Cluster) SetClusterType(v string) *Cluster {
	s.ClusterType = &v
	return s
}

func (s *Cluster) SetCreateTime(v int64) *Cluster {
	s.CreateTime = &v
	return s
}

func (s *Cluster) SetDeletionProtection(v bool) *Cluster {
	s.DeletionProtection = &v
	return s
}

func (s *Cluster) SetDeployMode(v string) *Cluster {
	s.DeployMode = &v
	return s
}

func (s *Cluster) SetDescription(v string) *Cluster {
	s.Description = &v
	return s
}

func (s *Cluster) SetEmrDefaultRole(v string) *Cluster {
	s.EmrDefaultRole = &v
	return s
}

func (s *Cluster) SetEndTime(v int64) *Cluster {
	s.EndTime = &v
	return s
}

func (s *Cluster) SetExpireTime(v int64) *Cluster {
	s.ExpireTime = &v
	return s
}

func (s *Cluster) SetNodeAttributes(v *NodeAttributes) *Cluster {
	s.NodeAttributes = v
	return s
}

func (s *Cluster) SetPaymentType(v string) *Cluster {
	s.PaymentType = &v
	return s
}

func (s *Cluster) SetReadyTime(v int64) *Cluster {
	s.ReadyTime = &v
	return s
}

func (s *Cluster) SetRegionId(v string) *Cluster {
	s.RegionId = &v
	return s
}

func (s *Cluster) SetReleaseVersion(v string) *Cluster {
	s.ReleaseVersion = &v
	return s
}

func (s *Cluster) SetResourceGroupId(v string) *Cluster {
	s.ResourceGroupId = &v
	return s
}

func (s *Cluster) SetSecurityMode(v string) *Cluster {
	s.SecurityMode = &v
	return s
}

func (s *Cluster) SetStateChangeReason(v *ClusterStateChangeReason) *Cluster {
	s.StateChangeReason = v
	return s
}

func (s *Cluster) SetSubscriptionConfig(v *SubscriptionConfig) *Cluster {
	s.SubscriptionConfig = v
	return s
}

func (s *Cluster) SetTags(v []*Tag) *Cluster {
	s.Tags = v
	return s
}

func (s *Cluster) Validate() error {
	if s.NodeAttributes != nil {
		if err := s.NodeAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.StateChangeReason != nil {
		if err := s.StateChangeReason.Validate(); err != nil {
			return err
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
