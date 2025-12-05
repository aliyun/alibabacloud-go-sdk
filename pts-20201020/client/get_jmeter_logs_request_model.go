// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIndex(v int32) *GetJMeterLogsRequest
	GetAgentIndex() *int32
	SetBeginTime(v int64) *GetJMeterLogsRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *GetJMeterLogsRequest
	GetEndTime() *int64
	SetKeyword(v string) *GetJMeterLogsRequest
	GetKeyword() *string
	SetLevel(v string) *GetJMeterLogsRequest
	GetLevel() *string
	SetPageNumber(v int32) *GetJMeterLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetJMeterLogsRequest
	GetPageSize() *int32
	SetReportId(v string) *GetJMeterLogsRequest
	GetReportId() *string
	SetThread(v string) *GetJMeterLogsRequest
	GetThread() *string
}

type GetJMeterLogsRequest struct {
	// Specifies that the operational logs of the stress tester identified as number 0 are returned if the agent index is invalid.
	//
	// example:
	//
	// 0
	AgentIndex *int32 `json:"AgentIndex,omitempty" xml:"AgentIndex,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	//
	// example:
	//
	// 1637115306000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query. Unit: seconds.
	//
	// example:
	//
	// 1637115309000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The log level. Valid values:
	//
	// 	- ERROR
	//
	// 	- WARN
	//
	// 	- INFO (default)
	//
	// 	- DEBUG
	//
	// 	- TRACE
	//
	// example:
	//
	// INFO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The report ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// KS2YE3J2
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The thread name.
	//
	// example:
	//
	// main
	Thread *string `json:"Thread,omitempty" xml:"Thread,omitempty"`
}

func (s GetJMeterLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterLogsRequest) GoString() string {
	return s.String()
}

func (s *GetJMeterLogsRequest) GetAgentIndex() *int32 {
	return s.AgentIndex
}

func (s *GetJMeterLogsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetJMeterLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetJMeterLogsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetJMeterLogsRequest) GetLevel() *string {
	return s.Level
}

func (s *GetJMeterLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetJMeterLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetJMeterLogsRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetJMeterLogsRequest) GetThread() *string {
	return s.Thread
}

func (s *GetJMeterLogsRequest) SetAgentIndex(v int32) *GetJMeterLogsRequest {
	s.AgentIndex = &v
	return s
}

func (s *GetJMeterLogsRequest) SetBeginTime(v int64) *GetJMeterLogsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetJMeterLogsRequest) SetEndTime(v int64) *GetJMeterLogsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJMeterLogsRequest) SetKeyword(v string) *GetJMeterLogsRequest {
	s.Keyword = &v
	return s
}

func (s *GetJMeterLogsRequest) SetLevel(v string) *GetJMeterLogsRequest {
	s.Level = &v
	return s
}

func (s *GetJMeterLogsRequest) SetPageNumber(v int32) *GetJMeterLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJMeterLogsRequest) SetPageSize(v int32) *GetJMeterLogsRequest {
	s.PageSize = &v
	return s
}

func (s *GetJMeterLogsRequest) SetReportId(v string) *GetJMeterLogsRequest {
	s.ReportId = &v
	return s
}

func (s *GetJMeterLogsRequest) SetThread(v string) *GetJMeterLogsRequest {
	s.Thread = &v
	return s
}

func (s *GetJMeterLogsRequest) Validate() error {
	return dara.Validate(s)
}
