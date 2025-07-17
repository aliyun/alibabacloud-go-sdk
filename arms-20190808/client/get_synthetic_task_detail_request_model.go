// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetSyntheticTaskDetailRequest
	GetRegionId() *string
	SetTaskId(v string) *GetSyntheticTaskDetailRequest
	GetTaskId() *string
}

type GetSyntheticTaskDetailRequest struct {
	// The region ID. Default value: cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the synthetic monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetSyntheticTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSyntheticTaskDetailRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSyntheticTaskDetailRequest) SetRegionId(v string) *GetSyntheticTaskDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetSyntheticTaskDetailRequest) SetTaskId(v string) *GetSyntheticTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *GetSyntheticTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
