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
	// 集群ID。
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 集群名称。
	//
	// example:
	//
	// emrtest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// 集群状态。取值范围：
	//
	// - STARTING：启动中。
	//
	// - START_FAILED：启动失败。
	//
	// - BOOTSTRAPPING：引导操作初始化。
	//
	// - RUNNING：运行中。
	//
	// - TERMINATING：终止中。
	//
	// - TERMINATED：已终止。
	//
	// - TERMINATED_WITH_ERRORS：发生异常导致终止。
	//
	// - TERMINATE_FAILED：终止失败。
	//
	// example:
	//
	// RUNNING
	ClusterState *string `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	// 集群类型。取值范围：
	//
	// - DATALAKE：新版数据湖。
	//
	// - OLAP：数据分析。
	//
	// - DATAFLOW：实时数据流。
	//
	// - DATASERVING：数据服务。
	//
	// example:
	//
	// DATALAKE
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// 创建时间。
	//
	// example:
	//
	// 1592837465784
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeletionProtection *bool   `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// EMR服务角色。
	EmrDefaultRole *string `json:"EmrDefaultRole,omitempty" xml:"EmrDefaultRole,omitempty"`
	// 删除时间。
	//
	// example:
	//
	// 1592837465784
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 过期时间。
	//
	// example:
	//
	// 1592837465784
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// 付费类型。取值范围：
	//
	// - PayAsYouGo：后付费。
	//
	// - Subscription：预付费。
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 可用时间。
	//
	// example:
	//
	// 1592837465784
	ReadyTime *int64 `json:"ReadyTime,omitempty" xml:"ReadyTime,omitempty"`
	// EMR发行版。
	//
	// example:
	//
	// EMR-5.8.0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// 资源组ID。
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 失败原因。
	StateChangeReason *ClusterStateChangeReason `json:"StateChangeReason,omitempty" xml:"StateChangeReason,omitempty"`
	// 标签列表。
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
	return dara.Validate(s)
}
