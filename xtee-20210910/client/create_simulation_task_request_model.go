// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimulationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateSimulationTaskRequest
	GetLang() *string
	SetDataSourceConfig(v string) *CreateSimulationTaskRequest
	GetDataSourceConfig() *string
	SetDataSourceType(v string) *CreateSimulationTaskRequest
	GetDataSourceType() *string
	SetEndTime(v int64) *CreateSimulationTaskRequest
	GetEndTime() *int64
	SetEventCode(v string) *CreateSimulationTaskRequest
	GetEventCode() *string
	SetFiltersStr(v string) *CreateSimulationTaskRequest
	GetFiltersStr() *string
	SetRegId(v string) *CreateSimulationTaskRequest
	GetRegId() *string
	SetRulesStr(v string) *CreateSimulationTaskRequest
	GetRulesStr() *string
	SetRunTask(v bool) *CreateSimulationTaskRequest
	GetRunTask() *bool
	SetStartTime(v int64) *CreateSimulationTaskRequest
	GetStartTime() *int64
	SetTaskName(v string) *CreateSimulationTaskRequest
	GetTaskName() *string
}

type CreateSimulationTaskRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Data source configuration
	//
	// example:
	//
	// {}
	DataSourceConfig *string `json:"dataSourceConfig,omitempty" xml:"dataSourceConfig,omitempty"`
	// Data source type
	//
	// example:
	//
	// SLS
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735541040000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Event code
	//
	// This parameter is required.
	//
	// example:
	//
	// de_anbwns2231
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Filters
	//
	// example:
	//
	// {"left":"score","operate":"bw","right":"222,333"}
	FiltersStr *string `json:"filtersStr,omitempty" xml:"filtersStr,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Rules list
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"100234\\"]
	RulesStr *string `json:"rulesStr,omitempty" xml:"rulesStr,omitempty"`
	// Whether to run the task directly
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	RunTask *bool `json:"runTask,omitempty" xml:"runTask,omitempty"`
	// Start time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1730448000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Task name
	//
	// This parameter is required.
	//
	// example:
	//
	// 仿真任务
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s CreateSimulationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSimulationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSimulationTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSimulationTaskRequest) GetDataSourceConfig() *string {
	return s.DataSourceConfig
}

func (s *CreateSimulationTaskRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CreateSimulationTaskRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateSimulationTaskRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *CreateSimulationTaskRequest) GetFiltersStr() *string {
	return s.FiltersStr
}

func (s *CreateSimulationTaskRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateSimulationTaskRequest) GetRulesStr() *string {
	return s.RulesStr
}

func (s *CreateSimulationTaskRequest) GetRunTask() *bool {
	return s.RunTask
}

func (s *CreateSimulationTaskRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateSimulationTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateSimulationTaskRequest) SetLang(v string) *CreateSimulationTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateSimulationTaskRequest) SetDataSourceConfig(v string) *CreateSimulationTaskRequest {
	s.DataSourceConfig = &v
	return s
}

func (s *CreateSimulationTaskRequest) SetDataSourceType(v string) *CreateSimulationTaskRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateSimulationTaskRequest) SetEndTime(v int64) *CreateSimulationTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateSimulationTaskRequest) SetEventCode(v string) *CreateSimulationTaskRequest {
	s.EventCode = &v
	return s
}

func (s *CreateSimulationTaskRequest) SetFiltersStr(v string) *CreateSimulationTaskRequest {
	s.FiltersStr = &v
	return s
}

func (s *CreateSimulationTaskRequest) SetRegId(v string) *CreateSimulationTaskRequest {
	s.RegId = &v
	return s
}

func (s *CreateSimulationTaskRequest) SetRulesStr(v string) *CreateSimulationTaskRequest {
	s.RulesStr = &v
	return s
}

func (s *CreateSimulationTaskRequest) SetRunTask(v bool) *CreateSimulationTaskRequest {
	s.RunTask = &v
	return s
}

func (s *CreateSimulationTaskRequest) SetStartTime(v int64) *CreateSimulationTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSimulationTaskRequest) SetTaskName(v string) *CreateSimulationTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateSimulationTaskRequest) Validate() error {
	return dara.Validate(s)
}
