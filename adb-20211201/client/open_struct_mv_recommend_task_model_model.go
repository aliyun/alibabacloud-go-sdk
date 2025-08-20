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
	SetScanQueriesRange(v int32) *OpenStructMvRecommendTaskModel
	GetScanQueriesRange() *int32
	SetSchedulingSettings(v string) *OpenStructMvRecommendTaskModel
	GetSchedulingSettings() *string
	SetTaskName(v string) *OpenStructMvRecommendTaskModel
	GetTaskName() *string
}

type OpenStructMvRecommendTaskModel struct {
	CreatedTime        *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LastRunAt          *string `json:"LastRunAt,omitempty" xml:"LastRunAt,omitempty"`
	ScanQueriesRange   *int32  `json:"ScanQueriesRange,omitempty" xml:"ScanQueriesRange,omitempty"`
	SchedulingSettings *string `json:"SchedulingSettings,omitempty" xml:"SchedulingSettings,omitempty"`
	TaskName           *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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

func (s *OpenStructMvRecommendTaskModel) GetScanQueriesRange() *int32 {
	return s.ScanQueriesRange
}

func (s *OpenStructMvRecommendTaskModel) GetSchedulingSettings() *string {
	return s.SchedulingSettings
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

func (s *OpenStructMvRecommendTaskModel) SetScanQueriesRange(v int32) *OpenStructMvRecommendTaskModel {
	s.ScanQueriesRange = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) SetSchedulingSettings(v string) *OpenStructMvRecommendTaskModel {
	s.SchedulingSettings = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) SetTaskName(v string) *OpenStructMvRecommendTaskModel {
	s.TaskName = &v
	return s
}

func (s *OpenStructMvRecommendTaskModel) Validate() error {
	return dara.Validate(s)
}
