// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertMarginThreshold(v int32) *CreateBaselineRequest
	GetAlertMarginThreshold() *int32
	SetBaselineName(v string) *CreateBaselineRequest
	GetBaselineName() *string
	SetBaselineType(v string) *CreateBaselineRequest
	GetBaselineType() *string
	SetNodeIds(v string) *CreateBaselineRequest
	GetNodeIds() *string
	SetOvertimeSettings(v []*CreateBaselineRequestOvertimeSettings) *CreateBaselineRequest
	GetOvertimeSettings() []*CreateBaselineRequestOvertimeSettings
	SetOwner(v string) *CreateBaselineRequest
	GetOwner() *string
	SetPriority(v int32) *CreateBaselineRequest
	GetPriority() *int32
	SetProjectId(v int64) *CreateBaselineRequest
	GetProjectId() *int64
}

type CreateBaselineRequest struct {
	// The alert margin threshold of the baseline. Unit: minutes.
	//
	// example:
	//
	// 30
	AlertMarginThreshold *int32 `json:"AlertMarginThreshold,omitempty" xml:"AlertMarginThreshold,omitempty"`
	// The name of the baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// BaselineName
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The type of the baseline. Valid values: DAILY and HOURLY.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAILY
	BaselineType *string `json:"BaselineType,omitempty" xml:"BaselineType,omitempty"`
	// The ancestor nodes of nodes in the baseline.
	//
	// example:
	//
	// 210001233239,210001236482
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The settings of the committed completion time of the baseline.
	//
	// This parameter is required.
	OvertimeSettings []*CreateBaselineRequestOvertimeSettings `json:"OvertimeSettings,omitempty" xml:"OvertimeSettings,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account used by the baseline owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: {1,3,5,7,8}.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,3,5,7,8
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBaselineRequest) GoString() string {
	return s.String()
}

func (s *CreateBaselineRequest) GetAlertMarginThreshold() *int32 {
	return s.AlertMarginThreshold
}

func (s *CreateBaselineRequest) GetBaselineName() *string {
	return s.BaselineName
}

func (s *CreateBaselineRequest) GetBaselineType() *string {
	return s.BaselineType
}

func (s *CreateBaselineRequest) GetNodeIds() *string {
	return s.NodeIds
}

func (s *CreateBaselineRequest) GetOvertimeSettings() []*CreateBaselineRequestOvertimeSettings {
	return s.OvertimeSettings
}

func (s *CreateBaselineRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateBaselineRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateBaselineRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateBaselineRequest) SetAlertMarginThreshold(v int32) *CreateBaselineRequest {
	s.AlertMarginThreshold = &v
	return s
}

func (s *CreateBaselineRequest) SetBaselineName(v string) *CreateBaselineRequest {
	s.BaselineName = &v
	return s
}

func (s *CreateBaselineRequest) SetBaselineType(v string) *CreateBaselineRequest {
	s.BaselineType = &v
	return s
}

func (s *CreateBaselineRequest) SetNodeIds(v string) *CreateBaselineRequest {
	s.NodeIds = &v
	return s
}

func (s *CreateBaselineRequest) SetOvertimeSettings(v []*CreateBaselineRequestOvertimeSettings) *CreateBaselineRequest {
	s.OvertimeSettings = v
	return s
}

func (s *CreateBaselineRequest) SetOwner(v string) *CreateBaselineRequest {
	s.Owner = &v
	return s
}

func (s *CreateBaselineRequest) SetPriority(v int32) *CreateBaselineRequest {
	s.Priority = &v
	return s
}

func (s *CreateBaselineRequest) SetProjectId(v int64) *CreateBaselineRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateBaselineRequest) Validate() error {
	return dara.Validate(s)
}

type CreateBaselineRequestOvertimeSettings struct {
	// The cycle that corresponds to the committed completion time. For a day-level baseline, set this parameter to 1. For an hour-level baseline, set this parameter to a value that is no more than 24.
	//
	// example:
	//
	// 1
	Cycle *int32 `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The committed completion time in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 00:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s CreateBaselineRequestOvertimeSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateBaselineRequestOvertimeSettings) GoString() string {
	return s.String()
}

func (s *CreateBaselineRequestOvertimeSettings) GetCycle() *int32 {
	return s.Cycle
}

func (s *CreateBaselineRequestOvertimeSettings) GetTime() *string {
	return s.Time
}

func (s *CreateBaselineRequestOvertimeSettings) SetCycle(v int32) *CreateBaselineRequestOvertimeSettings {
	s.Cycle = &v
	return s
}

func (s *CreateBaselineRequestOvertimeSettings) SetTime(v string) *CreateBaselineRequestOvertimeSettings {
	s.Time = &v
	return s
}

func (s *CreateBaselineRequestOvertimeSettings) Validate() error {
	return dara.Validate(s)
}
