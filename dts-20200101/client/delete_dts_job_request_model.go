// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDtsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *DeleteDtsJobRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *DeleteDtsJobRequest
	GetDtsJobId() *string
	SetJobType(v string) *DeleteDtsJobRequest
	GetJobType() *string
	SetRegionId(v string) *DeleteDtsJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteDtsJobRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *DeleteDtsJobRequest
	GetSynchronizationDirection() *string
	SetZeroEtlJob(v bool) *DeleteDtsJobRequest
	GetZeroEtlJob() *bool
}

type DeleteDtsJobRequest struct {
	// The data migration, data synchronization, or subscribe instance ID.
	//
	// example:
	//
	// dtsl3m1213ye7l****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking task.
	//
	// example:
	//
	// l3m1213ye7l****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The task type of the DTS instance. Valid values:
	//
	// - **MIGRATION**: data migration.
	//
	// - **SYNC**: data synchronization.
	//
	// - **SUBSCRIBE**: change tracking.
	//
	// example:
	//
	// MIGRATION
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The region ID of the data migration or data synchronization instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The field reserved for special business scenarios. You do not need to specify this parameter.
	//
	// example:
	//
	// rg-aek26lwshij****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward synchronization.
	//
	// - **Reverse**: reverse synchronization.
	//
	// > - Default value: **Forward**.
	//
	// - You can set this parameter to **Reverse*	- to release the reverse synchronization link only when the synchronization topology of the data synchronization instance is two-way.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The field reserved for special business scenarios. You do not need to specify this parameter.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s DeleteDtsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDtsJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDtsJobRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DeleteDtsJobRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DeleteDtsJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *DeleteDtsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDtsJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteDtsJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DeleteDtsJobRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *DeleteDtsJobRequest) SetDtsInstanceId(v string) *DeleteDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DeleteDtsJobRequest) SetDtsJobId(v string) *DeleteDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *DeleteDtsJobRequest) SetJobType(v string) *DeleteDtsJobRequest {
	s.JobType = &v
	return s
}

func (s *DeleteDtsJobRequest) SetRegionId(v string) *DeleteDtsJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDtsJobRequest) SetResourceGroupId(v string) *DeleteDtsJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteDtsJobRequest) SetSynchronizationDirection(v string) *DeleteDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DeleteDtsJobRequest) SetZeroEtlJob(v bool) *DeleteDtsJobRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *DeleteDtsJobRequest) Validate() error {
	return dara.Validate(s)
}
