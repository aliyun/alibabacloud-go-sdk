// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJobStepCheckpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *ModifyJobStepCheckpointRequest
	GetDtsJobId() *string
	SetJobStepId(v string) *ModifyJobStepCheckpointRequest
	GetJobStepId() *string
	SetNewCheckPoint(v int64) *ModifyJobStepCheckpointRequest
	GetNewCheckPoint() *int64
	SetRegionId(v string) *ModifyJobStepCheckpointRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyJobStepCheckpointRequest
	GetResourceGroupId() *string
}

type ModifyJobStepCheckpointRequest struct {
	// The ID of the migration, synchronization, or tracking task.
	//
	// example:
	//
	// l5512es7w15****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The task step ID, which corresponds to the **synchronization*	- step ID returned by the ListJobStep API.
	//
	// example:
	//
	// wn3z4ukia9wi9xu_0004_0000
	JobStepId *string `json:"JobStepId,omitempty" xml:"JobStepId,omitempty"`
	// The new offset as a UNIX timestamp in seconds. The specified value must fall within the offset range of the incremental data collection step.
	//
	// example:
	//
	// 1760406***
	NewCheckPoint *int64 `json:"NewCheckPoint,omitempty" xml:"NewCheckPoint,omitempty"`
	// The region where the DTS instance resides.
	//
	// > For regions that support dedicated clusters, see [DTS dedicated cluster](https://help.aliyun.com/document_detail/417481.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2ilvoxlrd***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyJobStepCheckpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobStepCheckpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyJobStepCheckpointRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyJobStepCheckpointRequest) GetJobStepId() *string {
	return s.JobStepId
}

func (s *ModifyJobStepCheckpointRequest) GetNewCheckPoint() *int64 {
	return s.NewCheckPoint
}

func (s *ModifyJobStepCheckpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyJobStepCheckpointRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyJobStepCheckpointRequest) SetDtsJobId(v string) *ModifyJobStepCheckpointRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyJobStepCheckpointRequest) SetJobStepId(v string) *ModifyJobStepCheckpointRequest {
	s.JobStepId = &v
	return s
}

func (s *ModifyJobStepCheckpointRequest) SetNewCheckPoint(v int64) *ModifyJobStepCheckpointRequest {
	s.NewCheckPoint = &v
	return s
}

func (s *ModifyJobStepCheckpointRequest) SetRegionId(v string) *ModifyJobStepCheckpointRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyJobStepCheckpointRequest) SetResourceGroupId(v string) *ModifyJobStepCheckpointRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyJobStepCheckpointRequest) Validate() error {
	return dara.Validate(s)
}
