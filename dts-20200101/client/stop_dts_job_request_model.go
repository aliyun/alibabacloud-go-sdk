// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDtsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *StopDtsJobRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *StopDtsJobRequest
	GetDtsJobId() *string
	SetRegionId(v string) *StopDtsJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StopDtsJobRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *StopDtsJobRequest
	GetSynchronizationDirection() *string
	SetZeroEtlJob(v bool) *StopDtsJobRequest
	GetZeroEtlJob() *bool
}

type StopDtsJobRequest struct {
	// The ID of the data migration, data synchronization, or subscribe instance.
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
	// The region where the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm2jr35xnjh7i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: Forward synchronization.
	//
	// - **Reverse**: Reverse synchronization.
	//
	// > - Default value: **Forward**.
	//
	// - You can specify **Reverse*	- to stop the reverse synchronization link only when the topology of the data synchronization instance is two-way.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// Specifies whether the task is a seamless integration (Zero-ETL) task. Valid values:
	//
	// - **false**: No.
	//
	// - **true**: Yes.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s StopDtsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDtsJobRequest) GoString() string {
	return s.String()
}

func (s *StopDtsJobRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *StopDtsJobRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *StopDtsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDtsJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StopDtsJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *StopDtsJobRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *StopDtsJobRequest) SetDtsInstanceId(v string) *StopDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *StopDtsJobRequest) SetDtsJobId(v string) *StopDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *StopDtsJobRequest) SetRegionId(v string) *StopDtsJobRequest {
	s.RegionId = &v
	return s
}

func (s *StopDtsJobRequest) SetResourceGroupId(v string) *StopDtsJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StopDtsJobRequest) SetSynchronizationDirection(v string) *StopDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *StopDtsJobRequest) SetZeroEtlJob(v bool) *StopDtsJobRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *StopDtsJobRequest) Validate() error {
	return dara.Validate(s)
}
