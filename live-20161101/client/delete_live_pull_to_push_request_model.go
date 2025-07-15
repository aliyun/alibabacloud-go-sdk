// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePullToPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteLivePullToPushRequest
	GetOwnerId() *int64
	SetRegion(v string) *DeleteLivePullToPushRequest
	GetRegion() *string
	SetRegionId(v string) *DeleteLivePullToPushRequest
	GetRegionId() *string
	SetTaskId(v string) *DeleteLivePullToPushRequest
	GetTaskId() *string
}

type DeleteLivePullToPushRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the task is started. Valid values:
	//
	// 	- ap-southeast-1: Singapore
	//
	// 	- ap-southeast-5: Indonesia (Jakarta)
	//
	// 	- cn-beijing: China (Beijing)
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3efb43c5-18ff-49eb-92a6-005f6521****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteLivePullToPushRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePullToPushRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivePullToPushRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLivePullToPushRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteLivePullToPushRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLivePullToPushRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteLivePullToPushRequest) SetOwnerId(v int64) *DeleteLivePullToPushRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLivePullToPushRequest) SetRegion(v string) *DeleteLivePullToPushRequest {
	s.Region = &v
	return s
}

func (s *DeleteLivePullToPushRequest) SetRegionId(v string) *DeleteLivePullToPushRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLivePullToPushRequest) SetTaskId(v string) *DeleteLivePullToPushRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteLivePullToPushRequest) Validate() error {
	return dara.Validate(s)
}
