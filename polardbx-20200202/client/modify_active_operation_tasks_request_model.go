// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ModifyActiveOperationTasksRequest
	GetIds() *string
	SetImmediateStart(v int64) *ModifyActiveOperationTasksRequest
	GetImmediateStart() *int64
	SetRegionId(v string) *ModifyActiveOperationTasksRequest
	GetRegionId() *string
	SetSwitchTime(v string) *ModifyActiveOperationTasksRequest
	GetSwitchTime() *string
}

type ModifyActiveOperationTasksRequest struct {
	// The O&M event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// Specifies whether to immediately execute the event. Valid values:
	//
	// - 1: immediately execute
	//
	// - 0: execute at the specified time.
	//
	// example:
	//
	// 1
	ImmediateStart *int64 `json:"ImmediateStart,omitempty" xml:"ImmediateStart,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The switchover start time in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2021-08-15T12:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyActiveOperationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksRequest) GetIds() *string {
	return s.Ids
}

func (s *ModifyActiveOperationTasksRequest) GetImmediateStart() *int64 {
	return s.ImmediateStart
}

func (s *ModifyActiveOperationTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyActiveOperationTasksRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyActiveOperationTasksRequest) SetIds(v string) *ModifyActiveOperationTasksRequest {
	s.Ids = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetImmediateStart(v int64) *ModifyActiveOperationTasksRequest {
	s.ImmediateStart = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetRegionId(v string) *ModifyActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetSwitchTime(v string) *ModifyActiveOperationTasksRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) Validate() error {
	return dara.Validate(s)
}
