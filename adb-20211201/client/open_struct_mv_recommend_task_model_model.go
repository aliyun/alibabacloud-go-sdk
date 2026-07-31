// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructMvRecommendTaskModel interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *OpenStructMvRecommendTaskModel
	GetCreatedTime() *string
	SetDescription(v string) *OpenStructMvRecommendTaskModel
	GetDescription() *string
	SetLastRunAt(v string) *OpenStructMvRecommendTaskModel
	GetLastRunAt() *string
	SetMinRewriteQueryCount(v int32) *OpenStructMvRecommendTaskModel
	GetMinRewriteQueryCount() *int32
	SetMinRewriteQueryPattern(v int32) *OpenStructMvRecommendTaskModel
	GetMinRewriteQueryPattern() *int32
	SetScanQueriesRange(v int32) *OpenStructMvRecommendTaskModel
	GetScanQueriesRange() *int32
	SetSchedulingSettings(v string) *OpenStructMvRecommendTaskModel
	GetSchedulingSettings() *string
	SetSlowQueryThreshold(v int32) *OpenStructMvRecommendTaskModel
	GetSlowQueryThreshold() *int32
	SetTaskName(v string) *OpenStructMvRecommendTaskModel
	GetTaskName() *string
}

type OpenStructMvRecommendTaskModel struct {
	// The creation time.
	//
	// example:
	//
	// 2024-12-12 23:59
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Description.
	//
	// example:
	//
	// task desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Last execution time.
	//
	// example:
	//
	// 2024-12-13 00:10
	LastRunAt *string `json:"LastRunAt,omitempty" xml:"LastRunAt,omitempty"`
	// The minimum number of slow queries that match the pattern.
	//
	// example:
	//
	// 3
	MinRewriteQueryCount *int32 `json:"MinRewriteQueryCount,omitempty" xml:"MinRewriteQueryCount,omitempty"`
	// The number of minimum acceleration patterns.
	//
	// example:
	//
	// 5
	MinRewriteQueryPattern *int32 `json:"MinRewriteQueryPattern,omitempty" xml:"MinRewriteQueryPattern,omitempty"`
	// The time range for scanning data. Unit: days. Default value: 3.
	//
	// example:
	//
	// 3
	ScanQueriesRange *int32 `json:"ScanQueriesRange,omitempty" xml:"ScanQueriesRange,omitempty"`
	// The execution schedule of the task.
	//
	// example:
	//
	// The 12:30 every day
	SchedulingSettings *string `json:"SchedulingSettings,omitempty" xml:"SchedulingSettings,omitempty"`
	// The wait threshold for slow queries.
	//
	// example:
	//
	// 2
	SlowQueryThreshold *int32 `json:"SlowQueryThreshold,omitempty" xml:"SlowQueryThreshold,omitempty"`
	// The name of the recommendation task.
	//
	// example:
	//
	// mv_task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s OpenStructMvRecommendTaskModel) String() string {
	return dara.Prettify(s)
}

func (s OpenStructMvRecommendTaskModel) GoString() string {
	return s.String()
}

func (s *OpenStructMvRecommendTaskModel) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *OpenStructMvRecommendTaskModel) GetDescription() *string {
	return s.Description
}

func (s *OpenStructMvRecommendTaskModel) GetLastRunAt() *string {
	return s.LastRunAt
}

func (s *OpenStructMvRecommendTaskModel) GetMinRewriteQueryCount() *int32 {
	return s.MinRewriteQueryCount
}

func (s *OpenStructMvRecommendTaskModel) GetMinRewriteQueryPattern() *int32 {
	return s.MinRewriteQueryPattern
}

func (s *OpenStructMvRecommendTaskModel) GetScanQueriesRange() *int32 {
	return s.ScanQueriesRange
}

func (s *OpenStructMvRecommendTaskModel) GetSchedulingSettings() *string {
	return s.SchedulingSettings
}

func (s *OpenStructMvRecommendTaskModel) GetSlowQueryThreshold() *int32 {
	return s.SlowQueryThreshold
}

func (s *OpenStructMvRecommendTaskModel) GetTaskName() *string {
	return s.TaskName
}

func (s *OpenStructMvRecommendTaskModel) SetCreatedTime(v string) *OpenStructMvRecommendTaskModel {
	s.CreatedTime = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) SetDescription(v string) *OpenStructMvRecommendTaskModel {
	s.Description = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) SetLastRunAt(v string) *OpenStructMvRecommendTaskModel {
	s.LastRunAt = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) SetMinRewriteQueryCount(v int32) *OpenStructMvRecommendTaskModel {
	s.MinRewriteQueryCount = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) SetMinRewriteQueryPattern(v int32) *OpenStructMvRecommendTaskModel {
	s.MinRewriteQueryPattern = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) SetScanQueriesRange(v int32) *OpenStructMvRecommendTaskModel {
	s.ScanQueriesRange = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) SetSchedulingSettings(v string) *OpenStructMvRecommendTaskModel {
	s.SchedulingSettings = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) SetSlowQueryThreshold(v int32) *OpenStructMvRecommendTaskModel {
	s.SlowQueryThreshold = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) SetTaskName(v string) *OpenStructMvRecommendTaskModel {
	s.TaskName = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) Validate() error {
	return dara.Validate(s)
}
