// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessSoarStrategyTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyTaskId(v int64) *ProcessSoarStrategyTaskRequest
	GetStrategyTaskId() *int64
	SetTaskAction(v string) *ProcessSoarStrategyTaskRequest
	GetTaskAction() *string
}

type ProcessSoarStrategyTaskRequest struct {
	// ID of the strategy task.
	//
	// > You can obtain this parameter by calling the [DescribeSoarStrategyTasks](~~DescribeSoarStrategyTasks~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	StrategyTaskId *int64 `json:"StrategyTaskId,omitempty" xml:"StrategyTaskId,omitempty"`
	// Task action status. Values:
	//
	// - SCHEDULE: Schedule
	//
	// - PAUSE: Pause
	//
	// This parameter is required.
	//
	// example:
	//
	// SCHEDULE
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s ProcessSoarStrategyTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ProcessSoarStrategyTaskRequest) GoString() string {
	return s.String()
}

func (s *ProcessSoarStrategyTaskRequest) GetStrategyTaskId() *int64 {
	return s.StrategyTaskId
}

func (s *ProcessSoarStrategyTaskRequest) GetTaskAction() *string {
	return s.TaskAction
}

func (s *ProcessSoarStrategyTaskRequest) SetStrategyTaskId(v int64) *ProcessSoarStrategyTaskRequest {
	s.StrategyTaskId = &v
	return s
}

func (s *ProcessSoarStrategyTaskRequest) SetTaskAction(v string) *ProcessSoarStrategyTaskRequest {
	s.TaskAction = &v
	return s
}

func (s *ProcessSoarStrategyTaskRequest) Validate() error {
	return dara.Validate(s)
}
