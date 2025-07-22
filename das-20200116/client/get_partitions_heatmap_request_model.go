// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPartitionsHeatmapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *GetPartitionsHeatmapRequest
	GetConsoleContext() *string
	SetInstanceId(v string) *GetPartitionsHeatmapRequest
	GetInstanceId() *string
	SetTimeRange(v string) *GetPartitionsHeatmapRequest
	GetTimeRange() *string
	SetType(v string) *GetPartitionsHeatmapRequest
	GetType() *string
}

type GetPartitionsHeatmapRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxc-hzrciqy62c****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time range to be queried. Valid values:
	//
	// 	- **LAST_ONE_HOURS**: the last hour.
	//
	// 	- **LAST_SIX_HOURS**: the last six hours.
	//
	// 	- **LAST_ONE_DAYS**: the last day.
	//
	// 	- **LAST_THREE_DAYS**: the last three days.
	//
	// 	- **LAST_SEVEN_DAYS**: the last seven days.
	//
	// example:
	//
	// LAST_SIX_HOURS
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
	// The type of the data to be queried. Valid values:
	//
	// 	- **READ_ROWS**: the read rows.
	//
	// 	- **WRITTEN_ROWS**: the written rows.
	//
	// 	- **READ_WRITTEN_ROWS**: the read and written rows.
	//
	// 	- **UPDATE_ROWS**: the updated rows.
	//
	// 	- **INSERTED_ROWS**: the inserted rows.
	//
	// 	- **DELETED_ROWS**: the deleted rows.
	//
	// 	- **READ_ROWS_WITH_DN**: the read rows returned from a data node.
	//
	// 	- **WRITTEN_ROWS_WITH_DN**: the written rows returned from a data node.
	//
	// 	- **READ_WRITTEN_ROWS_WITH_DN**: the read and written rows returned from a data node.
	//
	// example:
	//
	// WRITTEN_ROWS_WITH_DN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPartitionsHeatmapRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPartitionsHeatmapRequest) GoString() string {
	return s.String()
}

func (s *GetPartitionsHeatmapRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *GetPartitionsHeatmapRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPartitionsHeatmapRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *GetPartitionsHeatmapRequest) GetType() *string {
	return s.Type
}

func (s *GetPartitionsHeatmapRequest) SetConsoleContext(v string) *GetPartitionsHeatmapRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetPartitionsHeatmapRequest) SetInstanceId(v string) *GetPartitionsHeatmapRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPartitionsHeatmapRequest) SetTimeRange(v string) *GetPartitionsHeatmapRequest {
	s.TimeRange = &v
	return s
}

func (s *GetPartitionsHeatmapRequest) SetType(v string) *GetPartitionsHeatmapRequest {
	s.Type = &v
	return s
}

func (s *GetPartitionsHeatmapRequest) Validate() error {
	return dara.Validate(s)
}
