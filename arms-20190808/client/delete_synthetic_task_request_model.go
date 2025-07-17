// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSyntheticTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteSyntheticTaskRequest
	GetRegionId() *string
	SetTaskIds(v []*string) *DeleteSyntheticTaskRequest
	GetTaskIds() []*string
}

type DeleteSyntheticTaskRequest struct {
	// The region ID. Default value: cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task IDs.
	//
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s DeleteSyntheticTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSyntheticTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteSyntheticTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSyntheticTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *DeleteSyntheticTaskRequest) SetRegionId(v string) *DeleteSyntheticTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSyntheticTaskRequest) SetTaskIds(v []*string) *DeleteSyntheticTaskRequest {
	s.TaskIds = v
	return s
}

func (s *DeleteSyntheticTaskRequest) Validate() error {
	return dara.Validate(s)
}
