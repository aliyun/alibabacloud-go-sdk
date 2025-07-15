// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLivePullToPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *StopLivePullToPushRequest
	GetOwnerId() *int64
	SetRegion(v string) *StopLivePullToPushRequest
	GetRegion() *string
	SetRegionId(v string) *StopLivePullToPushRequest
	GetRegionId() *string
	SetTaskId(v string) *StopLivePullToPushRequest
	GetTaskId() *string
}

type StopLivePullToPushRequest struct {
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
	// 3bb44350-0c34-49c7-8c5e-cba5e6c0****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopLivePullToPushRequest) String() string {
	return dara.Prettify(s)
}

func (s StopLivePullToPushRequest) GoString() string {
	return s.String()
}

func (s *StopLivePullToPushRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopLivePullToPushRequest) GetRegion() *string {
	return s.Region
}

func (s *StopLivePullToPushRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopLivePullToPushRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopLivePullToPushRequest) SetOwnerId(v int64) *StopLivePullToPushRequest {
	s.OwnerId = &v
	return s
}

func (s *StopLivePullToPushRequest) SetRegion(v string) *StopLivePullToPushRequest {
	s.Region = &v
	return s
}

func (s *StopLivePullToPushRequest) SetRegionId(v string) *StopLivePullToPushRequest {
	s.RegionId = &v
	return s
}

func (s *StopLivePullToPushRequest) SetTaskId(v string) *StopLivePullToPushRequest {
	s.TaskId = &v
	return s
}

func (s *StopLivePullToPushRequest) Validate() error {
	return dara.Validate(s)
}
