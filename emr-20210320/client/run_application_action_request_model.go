// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunApplicationActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionName(v string) *RunApplicationActionRequest
	GetActionName() *string
	SetBatchSize(v int32) *RunApplicationActionRequest
	GetBatchSize() *int32
	SetClusterId(v string) *RunApplicationActionRequest
	GetClusterId() *string
	SetComponentInstanceSelector(v *ComponentInstanceSelector) *RunApplicationActionRequest
	GetComponentInstanceSelector() *ComponentInstanceSelector
	SetDescription(v string) *RunApplicationActionRequest
	GetDescription() *string
	SetExecuteStrategy(v string) *RunApplicationActionRequest
	GetExecuteStrategy() *string
	SetInterval(v int64) *RunApplicationActionRequest
	GetInterval() *int64
	SetRegionId(v string) *RunApplicationActionRequest
	GetRegionId() *string
	SetRollingExecute(v bool) *RunApplicationActionRequest
	GetRollingExecute() *bool
}

type RunApplicationActionRequest struct {
	// The name of the action. Valid values:
	//
	// 	- start
	//
	// 	- stop
	//
	// 	- config
	//
	// 	- restart
	//
	// 	- refresh_queues
	//
	// 	- refresh_labels
	//
	// This parameter is required.
	//
	// example:
	//
	// start
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The number of applications in each batch.
	//
	// example:
	//
	// 1
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-C95F0A39D8FF****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The operation object.
	//
	// This parameter is required.
	ComponentInstanceSelector *ComponentInstanceSelector `json:"ComponentInstanceSelector,omitempty" xml:"ComponentInstanceSelector,omitempty"`
	// The description of the execution.
	//
	// example:
	//
	// 运行描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution policy. Valid values:
	//
	// 	- FAILED_BLOCK: The system stops the execution if the execution fails.
	//
	// 	- FAILED_CONTINUE: The system continues the execution if the execution fails.
	//
	// example:
	//
	// FAILED_CONTINUE
	ExecuteStrategy *string `json:"ExecuteStrategy,omitempty" xml:"ExecuteStrategy,omitempty"`
	// The interval for rolling execution. Unit: seconds.
	//
	// example:
	//
	// 10
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable rolling execution.
	//
	// example:
	//
	// true
	RollingExecute *bool `json:"RollingExecute,omitempty" xml:"RollingExecute,omitempty"`
}

func (s RunApplicationActionRequest) String() string {
	return dara.Prettify(s)
}

func (s RunApplicationActionRequest) GoString() string {
	return s.String()
}

func (s *RunApplicationActionRequest) GetActionName() *string {
	return s.ActionName
}

func (s *RunApplicationActionRequest) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *RunApplicationActionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RunApplicationActionRequest) GetComponentInstanceSelector() *ComponentInstanceSelector {
	return s.ComponentInstanceSelector
}

func (s *RunApplicationActionRequest) GetDescription() *string {
	return s.Description
}

func (s *RunApplicationActionRequest) GetExecuteStrategy() *string {
	return s.ExecuteStrategy
}

func (s *RunApplicationActionRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *RunApplicationActionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunApplicationActionRequest) GetRollingExecute() *bool {
	return s.RollingExecute
}

func (s *RunApplicationActionRequest) SetActionName(v string) *RunApplicationActionRequest {
	s.ActionName = &v
	return s
}

func (s *RunApplicationActionRequest) SetBatchSize(v int32) *RunApplicationActionRequest {
	s.BatchSize = &v
	return s
}

func (s *RunApplicationActionRequest) SetClusterId(v string) *RunApplicationActionRequest {
	s.ClusterId = &v
	return s
}

func (s *RunApplicationActionRequest) SetComponentInstanceSelector(v *ComponentInstanceSelector) *RunApplicationActionRequest {
	s.ComponentInstanceSelector = v
	return s
}

func (s *RunApplicationActionRequest) SetDescription(v string) *RunApplicationActionRequest {
	s.Description = &v
	return s
}

func (s *RunApplicationActionRequest) SetExecuteStrategy(v string) *RunApplicationActionRequest {
	s.ExecuteStrategy = &v
	return s
}

func (s *RunApplicationActionRequest) SetInterval(v int64) *RunApplicationActionRequest {
	s.Interval = &v
	return s
}

func (s *RunApplicationActionRequest) SetRegionId(v string) *RunApplicationActionRequest {
	s.RegionId = &v
	return s
}

func (s *RunApplicationActionRequest) SetRollingExecute(v bool) *RunApplicationActionRequest {
	s.RollingExecute = &v
	return s
}

func (s *RunApplicationActionRequest) Validate() error {
	if s.ComponentInstanceSelector != nil {
		if err := s.ComponentInstanceSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}
