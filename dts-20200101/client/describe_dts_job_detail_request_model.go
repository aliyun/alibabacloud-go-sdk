// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsJobDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbObjectOutputType(v string) *DescribeDtsJobDetailRequest
	GetDbObjectOutputType() *string
	SetDtsInstanceID(v string) *DescribeDtsJobDetailRequest
	GetDtsInstanceID() *string
	SetDtsJobId(v string) *DescribeDtsJobDetailRequest
	GetDtsJobId() *string
	SetRegionId(v string) *DescribeDtsJobDetailRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDtsJobDetailRequest
	GetResourceGroupId() *string
	SetSyncSubJobHistory(v bool) *DescribeDtsJobDetailRequest
	GetSyncSubJobHistory() *bool
	SetSynchronizationDirection(v string) *DescribeDtsJobDetailRequest
	GetSynchronizationDirection() *string
	SetZeroEtlJob(v bool) *DescribeDtsJobDetailRequest
	GetZeroEtlJob() *bool
}

type DescribeDtsJobDetailRequest struct {
	DbObjectOutputType *string `json:"DbObjectOutputType,omitempty" xml:"DbObjectOutputType,omitempty"`
	// The instance ID of the data migration, data synchronization, or subscribe instance.
	//
	// example:
	//
	// dtsta7w132u12h****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking task.
	//
	// example:
	//
	// ta7w132u12h****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The ID of the region in which the task resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to return information about all synchronization subtasks. Default value: **false**, which returns only the synchronization subtask that is in progress or the most recently executed synchronization subtask.
	//
	// example:
	//
	// false
	SyncSubJobHistory *bool `json:"SyncSubJobHistory,omitempty" xml:"SyncSubJobHistory,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// > - Default value: **Forward**.
	//
	// - The value **Reverse*	- takes effect only when the synchronization topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// Specifies whether the task is a zero-ETL task. Valid values:
	//
	// - **true**: The task is a zero-ETL task.
	//
	// - **false**: The task is not a zero-ETL task.
	//
	// example:
	//
	// false
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s DescribeDtsJobDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailRequest) GetDbObjectOutputType() *string {
	return s.DbObjectOutputType
}

func (s *DescribeDtsJobDetailRequest) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDtsJobDetailRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDtsJobDetailRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDtsJobDetailRequest) GetSyncSubJobHistory() *bool {
	return s.SyncSubJobHistory
}

func (s *DescribeDtsJobDetailRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeDtsJobDetailRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *DescribeDtsJobDetailRequest) SetDbObjectOutputType(v string) *DescribeDtsJobDetailRequest {
	s.DbObjectOutputType = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) SetDtsInstanceID(v string) *DescribeDtsJobDetailRequest {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) SetDtsJobId(v string) *DescribeDtsJobDetailRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) SetRegionId(v string) *DescribeDtsJobDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) SetResourceGroupId(v string) *DescribeDtsJobDetailRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) SetSyncSubJobHistory(v bool) *DescribeDtsJobDetailRequest {
	s.SyncSubJobHistory = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) SetSynchronizationDirection(v string) *DescribeDtsJobDetailRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) SetZeroEtlJob(v bool) *DescribeDtsJobDetailRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) Validate() error {
	return dara.Validate(s)
}
