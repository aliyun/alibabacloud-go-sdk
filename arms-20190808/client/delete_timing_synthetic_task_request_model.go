// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTimingSyntheticTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteTimingSyntheticTaskRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteTimingSyntheticTaskRequest
	GetResourceGroupId() *string
	SetTaskId(v string) *DeleteTimingSyntheticTaskRequest
	GetTaskId() *string
}

type DeleteTimingSyntheticTaskRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2eq4peca****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the synthetic monitoring task.
	//
	// example:
	//
	// 5308a2691f59422c8c3b7aeccec9cd3b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteTimingSyntheticTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTimingSyntheticTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTimingSyntheticTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTimingSyntheticTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteTimingSyntheticTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteTimingSyntheticTaskRequest) SetRegionId(v string) *DeleteTimingSyntheticTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTimingSyntheticTaskRequest) SetResourceGroupId(v string) *DeleteTimingSyntheticTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteTimingSyntheticTaskRequest) SetTaskId(v string) *DeleteTimingSyntheticTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteTimingSyntheticTaskRequest) Validate() error {
	return dara.Validate(s)
}
