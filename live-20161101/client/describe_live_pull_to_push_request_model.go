// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePullToPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeLivePullToPushRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeLivePullToPushRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeLivePullToPushRequest
	GetRegionId() *string
	SetTaskId(v string) *DescribeLivePullToPushRequest
	GetTaskId() *string
}

type DescribeLivePullToPushRequest struct {
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
	// fd245384-4067-4f91-9d75-9666a6bc****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeLivePullToPushRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullToPushRequest) GoString() string {
	return s.String()
}

func (s *DescribeLivePullToPushRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLivePullToPushRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeLivePullToPushRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLivePullToPushRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeLivePullToPushRequest) SetOwnerId(v int64) *DescribeLivePullToPushRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLivePullToPushRequest) SetRegion(v string) *DescribeLivePullToPushRequest {
	s.Region = &v
	return s
}

func (s *DescribeLivePullToPushRequest) SetRegionId(v string) *DescribeLivePullToPushRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLivePullToPushRequest) SetTaskId(v string) *DescribeLivePullToPushRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeLivePullToPushRequest) Validate() error {
	return dara.Validate(s)
}
