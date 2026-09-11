// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendDtsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *SuspendDtsJobRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *SuspendDtsJobRequest
	GetDtsJobId() *string
	SetRegionId(v string) *SuspendDtsJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SuspendDtsJobRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *SuspendDtsJobRequest
	GetSynchronizationDirection() *string
	SetZeroEtlJob(v bool) *SuspendDtsJobRequest
	GetZeroEtlJob() *bool
}

type SuspendDtsJobRequest struct {
	// The migration, synchronization, or subscribe instance ID.
	//
	// example:
	//
	// dtsl3m1213ye7l****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data migration, synchronization, or change tracking task.
	//
	// > You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// example:
	//
	// l3m1213ye7l****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// > - The default value is **Forward**.
	//
	// - You can set this parameter to **Reverse*	- to suspend the reverse synchronization link only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// Specifies whether the node is a seamless integration (Zero-ETL) node. Valid values:
	//
	// - **false**: no.
	//
	// - **true**: yes.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s SuspendDtsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendDtsJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendDtsJobRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *SuspendDtsJobRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *SuspendDtsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SuspendDtsJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SuspendDtsJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *SuspendDtsJobRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *SuspendDtsJobRequest) SetDtsInstanceId(v string) *SuspendDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *SuspendDtsJobRequest) SetDtsJobId(v string) *SuspendDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *SuspendDtsJobRequest) SetRegionId(v string) *SuspendDtsJobRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendDtsJobRequest) SetResourceGroupId(v string) *SuspendDtsJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SuspendDtsJobRequest) SetSynchronizationDirection(v string) *SuspendDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *SuspendDtsJobRequest) SetZeroEtlJob(v bool) *SuspendDtsJobRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *SuspendDtsJobRequest) Validate() error {
	return dara.Validate(s)
}
