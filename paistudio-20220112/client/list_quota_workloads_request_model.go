// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaWorkloadsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeforeWorkloadId(v string) *ListQuotaWorkloadsRequest
	GetBeforeWorkloadId() *string
	SetGmtDequeuedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest
	GetGmtDequeuedTimeRange() *TimeRangeFilter
	SetGmtEnqueuedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest
	GetGmtEnqueuedTimeRange() *TimeRangeFilter
	SetGmtPositionModifiedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest
	GetGmtPositionModifiedTimeRange() *TimeRangeFilter
	SetInstanceId(v string) *ListQuotaWorkloadsRequest
	GetInstanceId() *string
	SetNodeName(v string) *ListQuotaWorkloadsRequest
	GetNodeName() *string
	SetOrder(v string) *ListQuotaWorkloadsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListQuotaWorkloadsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListQuotaWorkloadsRequest
	GetPageSize() *int32
	SetPosition(v int32) *ListQuotaWorkloadsRequest
	GetPosition() *int32
	SetPriority(v int32) *ListQuotaWorkloadsRequest
	GetPriority() *int32
	SetShowOwn(v bool) *ListQuotaWorkloadsRequest
	GetShowOwn() *bool
	SetSortBy(v string) *ListQuotaWorkloadsRequest
	GetSortBy() *string
	SetStatus(v string) *ListQuotaWorkloadsRequest
	GetStatus() *string
	SetSubQuotaIds(v string) *ListQuotaWorkloadsRequest
	GetSubQuotaIds() *string
	SetUseOversoldResource(v bool) *ListQuotaWorkloadsRequest
	GetUseOversoldResource() *bool
	SetUserIds(v string) *ListQuotaWorkloadsRequest
	GetUserIds() *string
	SetWithHistoricalData(v bool) *ListQuotaWorkloadsRequest
	GetWithHistoricalData() *bool
	SetWorkloadCreatedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest
	GetWorkloadCreatedTimeRange() *TimeRangeFilter
	SetWorkloadIds(v string) *ListQuotaWorkloadsRequest
	GetWorkloadIds() *string
	SetWorkloadNames(v string) *ListQuotaWorkloadsRequest
	GetWorkloadNames() *string
	SetWorkloadStatuses(v string) *ListQuotaWorkloadsRequest
	GetWorkloadStatuses() *string
	SetWorkloadType(v string) *ListQuotaWorkloadsRequest
	GetWorkloadType() *string
	SetWorkspaceIds(v string) *ListQuotaWorkloadsRequest
	GetWorkspaceIds() *string
}

type ListQuotaWorkloadsRequest struct {
	// Queries the queuing information for instances before a specified workload ID in the queue.
	//
	// example:
	//
	// dsw654433**
	BeforeWorkloadId *string `json:"BeforeWorkloadId,omitempty" xml:"BeforeWorkloadId,omitempty"`
	// The range search for workload dequeue time.
	GmtDequeuedTimeRange *TimeRangeFilter `json:"GmtDequeuedTimeRange,omitempty" xml:"GmtDequeuedTimeRange,omitempty"`
	// The range search for workload enqueue time.
	GmtEnqueuedTimeRange *TimeRangeFilter `json:"GmtEnqueuedTimeRange,omitempty" xml:"GmtEnqueuedTimeRange,omitempty"`
	// The range search for workload queue position change time.
	GmtPositionModifiedTimeRange *TimeRangeFilter `json:"GmtPositionModifiedTimeRange,omitempty" xml:"GmtPositionModifiedTimeRange,omitempty"`
	InstanceId                   *string          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Searches only for nodes on the specified node.
	//
	// example:
	//
	// lrn482781276**
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The sort order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number of the query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Specifies whether to display only the workloads associated with the current resource quota.
	//
	// example:
	//
	// true
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// The field by which to sort the returned results.
	//
	// example:
	//
	// GmtCreatedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The status in the queue.
	//
	// example:
	//
	// Enqueued
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of sub-resource quota IDs.
	//
	// example:
	//
	// quota123446**,quota644322**
	SubQuotaIds         *string `json:"SubQuotaIds,omitempty" xml:"SubQuotaIds,omitempty"`
	UseOversoldResource *bool   `json:"UseOversoldResource,omitempty" xml:"UseOversoldResource,omitempty"`
	// The list of user IDs.
	//
	// example:
	//
	// 290438938**,238290930**
	UserIds            *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	WithHistoricalData *bool   `json:"WithHistoricalData,omitempty" xml:"WithHistoricalData,omitempty"`
	// The range search for workload creation time.
	WorkloadCreatedTimeRange *TimeRangeFilter `json:"WorkloadCreatedTimeRange,omitempty" xml:"WorkloadCreatedTimeRange,omitempty"`
	// Queries the queuing information for specified workload IDs.
	//
	// example:
	//
	// dlc123445**
	WorkloadIds   *string `json:"WorkloadIds,omitempty" xml:"WorkloadIds,omitempty"`
	WorkloadNames *string `json:"WorkloadNames,omitempty" xml:"WorkloadNames,omitempty"`
	// Filters by workload status.
	//
	// example:
	//
	// Pending
	WorkloadStatuses *string `json:"WorkloadStatuses,omitempty" xml:"WorkloadStatuses,omitempty"`
	// Queries the queuing information for instances of a specified sub-product.
	//
	// example:
	//
	// dlc
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
	// Queries the queuing information for workloads in a specified workspace.
	//
	// example:
	//
	// 1866**
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s ListQuotaWorkloadsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaWorkloadsRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaWorkloadsRequest) GetBeforeWorkloadId() *string {
	return s.BeforeWorkloadId
}

func (s *ListQuotaWorkloadsRequest) GetGmtDequeuedTimeRange() *TimeRangeFilter {
	return s.GmtDequeuedTimeRange
}

func (s *ListQuotaWorkloadsRequest) GetGmtEnqueuedTimeRange() *TimeRangeFilter {
	return s.GmtEnqueuedTimeRange
}

func (s *ListQuotaWorkloadsRequest) GetGmtPositionModifiedTimeRange() *TimeRangeFilter {
	return s.GmtPositionModifiedTimeRange
}

func (s *ListQuotaWorkloadsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListQuotaWorkloadsRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *ListQuotaWorkloadsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListQuotaWorkloadsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQuotaWorkloadsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQuotaWorkloadsRequest) GetPosition() *int32 {
	return s.Position
}

func (s *ListQuotaWorkloadsRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ListQuotaWorkloadsRequest) GetShowOwn() *bool {
	return s.ShowOwn
}

func (s *ListQuotaWorkloadsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListQuotaWorkloadsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListQuotaWorkloadsRequest) GetSubQuotaIds() *string {
	return s.SubQuotaIds
}

func (s *ListQuotaWorkloadsRequest) GetUseOversoldResource() *bool {
	return s.UseOversoldResource
}

func (s *ListQuotaWorkloadsRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *ListQuotaWorkloadsRequest) GetWithHistoricalData() *bool {
	return s.WithHistoricalData
}

func (s *ListQuotaWorkloadsRequest) GetWorkloadCreatedTimeRange() *TimeRangeFilter {
	return s.WorkloadCreatedTimeRange
}

func (s *ListQuotaWorkloadsRequest) GetWorkloadIds() *string {
	return s.WorkloadIds
}

func (s *ListQuotaWorkloadsRequest) GetWorkloadNames() *string {
	return s.WorkloadNames
}

func (s *ListQuotaWorkloadsRequest) GetWorkloadStatuses() *string {
	return s.WorkloadStatuses
}

func (s *ListQuotaWorkloadsRequest) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *ListQuotaWorkloadsRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *ListQuotaWorkloadsRequest) SetBeforeWorkloadId(v string) *ListQuotaWorkloadsRequest {
	s.BeforeWorkloadId = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetGmtDequeuedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest {
	s.GmtDequeuedTimeRange = v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetGmtEnqueuedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest {
	s.GmtEnqueuedTimeRange = v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetGmtPositionModifiedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest {
	s.GmtPositionModifiedTimeRange = v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetInstanceId(v string) *ListQuotaWorkloadsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetNodeName(v string) *ListQuotaWorkloadsRequest {
	s.NodeName = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetOrder(v string) *ListQuotaWorkloadsRequest {
	s.Order = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetPageNumber(v int32) *ListQuotaWorkloadsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetPageSize(v int32) *ListQuotaWorkloadsRequest {
	s.PageSize = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetPosition(v int32) *ListQuotaWorkloadsRequest {
	s.Position = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetPriority(v int32) *ListQuotaWorkloadsRequest {
	s.Priority = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetShowOwn(v bool) *ListQuotaWorkloadsRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetSortBy(v string) *ListQuotaWorkloadsRequest {
	s.SortBy = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetStatus(v string) *ListQuotaWorkloadsRequest {
	s.Status = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetSubQuotaIds(v string) *ListQuotaWorkloadsRequest {
	s.SubQuotaIds = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetUseOversoldResource(v bool) *ListQuotaWorkloadsRequest {
	s.UseOversoldResource = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetUserIds(v string) *ListQuotaWorkloadsRequest {
	s.UserIds = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWithHistoricalData(v bool) *ListQuotaWorkloadsRequest {
	s.WithHistoricalData = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkloadCreatedTimeRange(v *TimeRangeFilter) *ListQuotaWorkloadsRequest {
	s.WorkloadCreatedTimeRange = v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkloadIds(v string) *ListQuotaWorkloadsRequest {
	s.WorkloadIds = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkloadNames(v string) *ListQuotaWorkloadsRequest {
	s.WorkloadNames = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkloadStatuses(v string) *ListQuotaWorkloadsRequest {
	s.WorkloadStatuses = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkloadType(v string) *ListQuotaWorkloadsRequest {
	s.WorkloadType = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) SetWorkspaceIds(v string) *ListQuotaWorkloadsRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListQuotaWorkloadsRequest) Validate() error {
	if s.GmtDequeuedTimeRange != nil {
		if err := s.GmtDequeuedTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.GmtEnqueuedTimeRange != nil {
		if err := s.GmtEnqueuedTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.GmtPositionModifiedTimeRange != nil {
		if err := s.GmtPositionModifiedTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.WorkloadCreatedTimeRange != nil {
		if err := s.WorkloadCreatedTimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}
