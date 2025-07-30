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
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
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
	// The type of the Data Transmission Service (DTS) task. Valid values:
	//
	// 	- **MIGRATION**: data migration task
	//
	// 	- **SYNC**: data synchronization task
	//
	// 	- **SUBSCRIBE**: change tracking task
	//
	// example:
	//
	// MIGRATION
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aek26lwshijfk3q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The dynamic error code. This parameter will be removed in the future.
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
