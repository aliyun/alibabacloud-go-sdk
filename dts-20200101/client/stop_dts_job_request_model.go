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
	// The ID of the data migration, data synchronization, or change tracking instance.
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
	// The region ID of the Data Transmission Service (DTS) instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aekznwnajjh4d3a
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// >
	//
	// 	- Default value: **Forward**.
	//
	// 	- You can set this parameter to **Reverse*	- to stop the reverse synchronization task only when the topology is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// Whether it is a seamless integration (Zero-ETL) task, the value can be:
	//
	// - **false**: No. - **true**: Yes.
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
