// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructRefreshJobModel interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *OpenStructRefreshJobModel
	GetEndTime() *string
	SetName(v string) *OpenStructRefreshJobModel
	GetName() *string
	SetProcessid(v string) *OpenStructRefreshJobModel
	GetProcessid() *string
	SetRefreshInterval(v string) *OpenStructRefreshJobModel
	GetRefreshInterval() *string
	SetRefreshModel(v string) *OpenStructRefreshJobModel
	GetRefreshModel() *string
	SetResourceGroup(v string) *OpenStructRefreshJobModel
	GetResourceGroup() *string
	SetScheduledStartTime(v string) *OpenStructRefreshJobModel
	GetScheduledStartTime() *string
	SetSchemaName(v string) *OpenStructRefreshJobModel
	GetSchemaName() *string
	SetStartTime(v string) *OpenStructRefreshJobModel
	GetStartTime() *string
	SetStatus(v string) *OpenStructRefreshJobModel
	GetStatus() *string
}

type OpenStructRefreshJobModel struct {
	// The end time of the refresh task.
	//
	// example:
	//
	// 2025-01-15 10:12:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The materialized view name.
	//
	// example:
	//
	// demo_mv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2025103010xxxxxxx
	Processid *string `json:"Processid,omitempty" xml:"Processid,omitempty"`
	// The refresh interval.
	//
	// example:
	//
	// 30s
	RefreshInterval *string `json:"RefreshInterval,omitempty" xml:"RefreshInterval,omitempty"`
	// The refresh model.
	//
	// 	- FAST
	//
	// 	- COMPLETE
	//
	// example:
	//
	// FAST
	RefreshModel *string `json:"RefreshModel,omitempty" xml:"RefreshModel,omitempty"`
	// The resource group used for the refresh.
	//
	// example:
	//
	// user_default
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The scheduled start time.
	//
	// example:
	//
	// 2025-01-15 10:11:00
	ScheduledStartTime *string `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	// The database name.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The actual start time.
	//
	// example:
	//
	// 2025-01-15 10:11:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the queried refresh task.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s OpenStructRefreshJobModel) String() string {
	return dara.Prettify(s)
}

func (s OpenStructRefreshJobModel) GoString() string {
	return s.String()
}

func (s *OpenStructRefreshJobModel) GetEndTime() *string {
	return s.EndTime
}

func (s *OpenStructRefreshJobModel) GetName() *string {
	return s.Name
}

func (s *OpenStructRefreshJobModel) GetProcessid() *string {
	return s.Processid
}

func (s *OpenStructRefreshJobModel) GetRefreshInterval() *string {
	return s.RefreshInterval
}

func (s *OpenStructRefreshJobModel) GetRefreshModel() *string {
	return s.RefreshModel
}

func (s *OpenStructRefreshJobModel) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *OpenStructRefreshJobModel) GetScheduledStartTime() *string {
	return s.ScheduledStartTime
}

func (s *OpenStructRefreshJobModel) GetSchemaName() *string {
	return s.SchemaName
}

func (s *OpenStructRefreshJobModel) GetStartTime() *string {
	return s.StartTime
}

func (s *OpenStructRefreshJobModel) GetStatus() *string {
	return s.Status
}

func (s *OpenStructRefreshJobModel) SetEndTime(v string) *OpenStructRefreshJobModel {
	s.EndTime = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetName(v string) *OpenStructRefreshJobModel {
	s.Name = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetProcessid(v string) *OpenStructRefreshJobModel {
	s.Processid = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetRefreshInterval(v string) *OpenStructRefreshJobModel {
	s.RefreshInterval = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetRefreshModel(v string) *OpenStructRefreshJobModel {
	s.RefreshModel = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetResourceGroup(v string) *OpenStructRefreshJobModel {
	s.ResourceGroup = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetScheduledStartTime(v string) *OpenStructRefreshJobModel {
	s.ScheduledStartTime = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetSchemaName(v string) *OpenStructRefreshJobModel {
	s.SchemaName = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetStartTime(v string) *OpenStructRefreshJobModel {
	s.StartTime = &v
	return s
}

func (s *OpenStructRefreshJobModel) SetStatus(v string) *OpenStructRefreshJobModel {
	s.Status = &v
	return s
}

func (s *OpenStructRefreshJobModel) Validate() error {
	return dara.Validate(s)
}
