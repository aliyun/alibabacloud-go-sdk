// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClusterSummary interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ClusterSummary
	GetClusterId() *string
	SetClusterName(v string) *ClusterSummary
	GetClusterName() *string
	SetClusterState(v string) *ClusterSummary
	GetClusterState() *string
	SetClusterType(v string) *ClusterSummary
	GetClusterType() *string
	SetCreateTime(v int64) *ClusterSummary
	GetCreateTime() *int64
	SetDeletionProtection(v bool) *ClusterSummary
	GetDeletionProtection() *bool
	SetDescription(v string) *ClusterSummary
	GetDescription() *string
	SetEmrDefaultRole(v string) *ClusterSummary
	GetEmrDefaultRole() *string
	SetEndTime(v int64) *ClusterSummary
	GetEndTime() *int64
	SetExpireTime(v int64) *ClusterSummary
	GetExpireTime() *int64
	SetPaymentType(v string) *ClusterSummary
	GetPaymentType() *string
	SetReadyTime(v int64) *ClusterSummary
	GetReadyTime() *int64
	SetReleaseVersion(v string) *ClusterSummary
	GetReleaseVersion() *string
	SetResourceGroupId(v string) *ClusterSummary
	GetResourceGroupId() *string
	SetStateChangeReason(v *ClusterStateChangeReason) *ClusterSummary
	GetStateChangeReason() *ClusterStateChangeReason
	SetTags(v []*Tag) *ClusterSummary
	GetTags() []*Tag
}

type ClusterSummary struct {
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
	// The state of the cluster. Valid values:
	//
	// - `STARTING`: The cluster is starting.
	//
	// - `START_FAILED`: The cluster fails to be started.
	//
	// - `BOOTSTRAPPING`: The cluster is being initialized.
	//
	// - `RUNNING`: The cluster is running.
	//
	// - `TERMINATING`: The cluster is being terminated.
	//
	// - `TERMINATED`: The cluster is terminated.
	//
	// - `TERMINATED_WITH_ERRORS`: The cluster is terminated with errors.
	//
	// - `TERMINATE_FAILED`: The cluster fails to be terminated.
	//
	// example:
	//
	// RUNNING
	ClusterState *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	// The cluster type. Valid values:
	//
	// - `DATALAKE`: data lake.
	//
	// - `OLAP`: data analytics.
	//
	// - `DATAFLOW`: real-time dataflow.
	//
	// - `DATASERVING`: data serving.
	//
	// example:
	//
	// DATALAKE
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 1592837465788
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The release protection feature.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// EMR cluster
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The EMR service role.
	//
	// example:
	//
	// AliyunEMRDefaultRole
	EmrDefaultRole *string `json:"EmrDefaultRole,omitempty" xml:"EmrDefaultRole,omitempty"`
	// The time when the cluster was deleted.
	//
	// example:
	//
	// 1592837485788
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1592837475788
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The billing method. Valid values:
	//
	// - `PayAsYouGo`: pay-as-you-go.
	//
	// - `Subscription`: subscription.
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The time when the cluster is available.
	//
	// example:
	//
	// 1592837465788
	ReadyTime *int64 `json:"ReadyTime,omitempty" xml:"ReadyTime,omitempty"`
	// The E-MapReduce (EMR) release version.
	//
	// example:
	//
	// EMR-5.8.0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The reason for the failure.
	StateChangeReason *ClusterStateChangeReason `json:"StateChangeReason,omitempty" xml:"StateChangeReason,omitempty"`
	// The list of tags.
	Tags []*Tag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ClusterSummary) String() string {
	return dara.Prettify(s)
}

func (s ClusterSummary) GoString() string {
	return s.String()
}

func (s *ClusterSummary) GetClusterId() *string {
	return s.ClusterId
}

func (s *ClusterSummary) GetClusterName() *string {
	return s.ClusterName
}

func (s *ClusterSummary) GetClusterState() *string {
	return s.ClusterState
}

func (s *ClusterSummary) GetClusterType() *string {
	return s.ClusterType
}

func (s *ClusterSummary) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ClusterSummary) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ClusterSummary) GetDescription() *string {
	return s.Description
}

func (s *ClusterSummary) GetEmrDefaultRole() *string {
	return s.EmrDefaultRole
}

func (s *ClusterSummary) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClusterSummary) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ClusterSummary) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ClusterSummary) GetReadyTime() *int64 {
	return s.ReadyTime
}

func (s *ClusterSummary) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ClusterSummary) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ClusterSummary) GetStateChangeReason() *ClusterStateChangeReason {
	return s.StateChangeReason
}

func (s *ClusterSummary) GetTags() []*Tag {
	return s.Tags
}

func (s *ClusterSummary) SetClusterId(v string) *ClusterSummary {
	s.ClusterId = &v
	return s
}

func (s *ClusterSummary) SetClusterName(v string) *ClusterSummary {
	s.ClusterName = &v
	return s
}

func (s *ClusterSummary) SetClusterState(v string) *ClusterSummary {
	s.ClusterState = &v
	return s
}

func (s *ClusterSummary) SetClusterType(v string) *ClusterSummary {
	s.ClusterType = &v
	return s
}

func (s *ClusterSummary) SetCreateTime(v int64) *ClusterSummary {
	s.CreateTime = &v
	return s
}

func (s *ClusterSummary) SetDeletionProtection(v bool) *ClusterSummary {
	s.DeletionProtection = &v
	return s
}

func (s *ClusterSummary) SetDescription(v string) *ClusterSummary {
	s.Description = &v
	return s
}

func (s *ClusterSummary) SetEmrDefaultRole(v string) *ClusterSummary {
	s.EmrDefaultRole = &v
	return s
}

func (s *ClusterSummary) SetEndTime(v int64) *ClusterSummary {
	s.EndTime = &v
	return s
}

func (s *ClusterSummary) SetExpireTime(v int64) *ClusterSummary {
	s.ExpireTime = &v
	return s
}

func (s *ClusterSummary) SetPaymentType(v string) *ClusterSummary {
	s.PaymentType = &v
	return s
}

func (s *ClusterSummary) SetReadyTime(v int64) *ClusterSummary {
	s.ReadyTime = &v
	return s
}

func (s *ClusterSummary) SetReleaseVersion(v string) *ClusterSummary {
	s.ReleaseVersion = &v
	return s
}

func (s *ClusterSummary) SetResourceGroupId(v string) *ClusterSummary {
	s.ResourceGroupId = &v
	return s
}

func (s *ClusterSummary) SetStateChangeReason(v *ClusterStateChangeReason) *ClusterSummary {
	s.StateChangeReason = v
	return s
}

func (s *ClusterSummary) SetTags(v []*Tag) *ClusterSummary {
	s.Tags = v
	return s
}

func (s *ClusterSummary) Validate() error {
	if s.StateChangeReason != nil {
		if err := s.StateChangeReason.Validate(); err != nil {
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
