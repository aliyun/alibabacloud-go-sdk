// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartLivePullToPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *RestartLivePullToPushRequest
	GetOwnerId() *int64
	SetRegion(v string) *RestartLivePullToPushRequest
	GetRegion() *string
	SetRegionId(v string) *RestartLivePullToPushRequest
	GetRegionId() *string
	SetTaskId(v string) *RestartLivePullToPushRequest
	GetTaskId() *string
}

type RestartLivePullToPushRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the live center. Valid values:
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
	// preregion
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6f869419-0692-4fd5-8cf0-66cab659****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartLivePullToPushRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartLivePullToPushRequest) GoString() string {
	return s.String()
}

func (s *RestartLivePullToPushRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartLivePullToPushRequest) GetRegion() *string {
	return s.Region
}

func (s *RestartLivePullToPushRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartLivePullToPushRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RestartLivePullToPushRequest) SetOwnerId(v int64) *RestartLivePullToPushRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartLivePullToPushRequest) SetRegion(v string) *RestartLivePullToPushRequest {
	s.Region = &v
	return s
}

func (s *RestartLivePullToPushRequest) SetRegionId(v string) *RestartLivePullToPushRequest {
	s.RegionId = &v
	return s
}

func (s *RestartLivePullToPushRequest) SetTaskId(v string) *RestartLivePullToPushRequest {
	s.TaskId = &v
	return s
}

func (s *RestartLivePullToPushRequest) Validate() error {
	return dara.Validate(s)
}
