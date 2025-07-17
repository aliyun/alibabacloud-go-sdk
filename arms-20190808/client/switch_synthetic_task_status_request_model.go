// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSyntheticTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSwitchStatus(v int64) *SwitchSyntheticTaskStatusRequest
	GetSwitchStatus() *int64
	SetTaskIds(v []*int64) *SwitchSyntheticTaskStatusRequest
	GetTaskIds() []*int64
}

type SwitchSyntheticTaskStatusRequest struct {
	// Specifies whether to start or stop the task. Valid values:
	//
	// 	- **0**: stops the task
	//
	// 	- **1**: starts the task
	//
	// example:
	//
	// 0
	SwitchStatus *int64 `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
	// The task IDs. You can specify up to 30 task IDs at a time.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s SwitchSyntheticTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchSyntheticTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *SwitchSyntheticTaskStatusRequest) GetSwitchStatus() *int64 {
	return s.SwitchStatus
}

func (s *SwitchSyntheticTaskStatusRequest) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *SwitchSyntheticTaskStatusRequest) SetSwitchStatus(v int64) *SwitchSyntheticTaskStatusRequest {
	s.SwitchStatus = &v
	return s
}

func (s *SwitchSyntheticTaskStatusRequest) SetTaskIds(v []*int64) *SwitchSyntheticTaskStatusRequest {
	s.TaskIds = v
	return s
}

func (s *SwitchSyntheticTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
