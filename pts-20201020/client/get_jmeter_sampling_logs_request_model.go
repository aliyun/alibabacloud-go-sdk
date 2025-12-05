// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterSamplingLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *GetJMeterSamplingLogsRequest
	GetAgentId() *int64
	SetBeginTime(v int64) *GetJMeterSamplingLogsRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *GetJMeterSamplingLogsRequest
	GetEndTime() *int64
	SetKeyword(v string) *GetJMeterSamplingLogsRequest
	GetKeyword() *string
	SetMaxRT(v int32) *GetJMeterSamplingLogsRequest
	GetMaxRT() *int32
	SetMinRT(v int32) *GetJMeterSamplingLogsRequest
	GetMinRT() *int32
	SetPageNumber(v int32) *GetJMeterSamplingLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetJMeterSamplingLogsRequest
	GetPageSize() *int32
	SetReportId(v string) *GetJMeterSamplingLogsRequest
	GetReportId() *string
	SetResponseCode(v string) *GetJMeterSamplingLogsRequest
	GetResponseCode() *string
	SetSamplerId(v int32) *GetJMeterSamplingLogsRequest
	GetSamplerId() *int32
	SetSuccess(v bool) *GetJMeterSamplingLogsRequest
	GetSuccess() *bool
	SetThread(v string) *GetJMeterSamplingLogsRequest
	GetThread() *string
}

type GetJMeterSamplingLogsRequest struct {
	// The ID of the load generator. This parameter is disabled.
	//
	// example:
	//
	// 14238000
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1637157073000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1637157076000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword. You can use the keyword in the value of **SceneName*	- for fuzzy searching or use the value of **SceneID*	- for exact searching.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum response time. Unit: ms.
	//
	// example:
	//
	// 1000
	MaxRT *int32 `json:"MaxRT,omitempty" xml:"MaxRT,omitempty"`
	// The minimum response time. Unit: ms.
	//
	// example:
	//
	// 0
	MinRT *int32 `json:"MinRT,omitempty" xml:"MinRT,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the report.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7R4RE352
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The response code.
	//
	// example:
	//
	// 200
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// The ID of the sampler. The value starts from 0.
	//
	// example:
	//
	// 0
	SamplerId *int32 `json:"SamplerId,omitempty" xml:"SamplerId,omitempty"`
	// Specifies whether the sampling is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The name of the thread. Fuzzy searching is supported. If you specify multiple threads, separate the threads with spaces.
	//
	// example:
	//
	// main
	Thread *string `json:"Thread,omitempty" xml:"Thread,omitempty"`
}

func (s GetJMeterSamplingLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterSamplingLogsRequest) GoString() string {
	return s.String()
}

func (s *GetJMeterSamplingLogsRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *GetJMeterSamplingLogsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetJMeterSamplingLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetJMeterSamplingLogsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetJMeterSamplingLogsRequest) GetMaxRT() *int32 {
	return s.MaxRT
}

func (s *GetJMeterSamplingLogsRequest) GetMinRT() *int32 {
	return s.MinRT
}

func (s *GetJMeterSamplingLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetJMeterSamplingLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetJMeterSamplingLogsRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetJMeterSamplingLogsRequest) GetResponseCode() *string {
	return s.ResponseCode
}

func (s *GetJMeterSamplingLogsRequest) GetSamplerId() *int32 {
	return s.SamplerId
}

func (s *GetJMeterSamplingLogsRequest) GetSuccess() *bool {
	return s.Success
}

func (s *GetJMeterSamplingLogsRequest) GetThread() *string {
	return s.Thread
}

func (s *GetJMeterSamplingLogsRequest) SetAgentId(v int64) *GetJMeterSamplingLogsRequest {
	s.AgentId = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetBeginTime(v int64) *GetJMeterSamplingLogsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetEndTime(v int64) *GetJMeterSamplingLogsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetKeyword(v string) *GetJMeterSamplingLogsRequest {
	s.Keyword = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetMaxRT(v int32) *GetJMeterSamplingLogsRequest {
	s.MaxRT = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetMinRT(v int32) *GetJMeterSamplingLogsRequest {
	s.MinRT = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetPageNumber(v int32) *GetJMeterSamplingLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetPageSize(v int32) *GetJMeterSamplingLogsRequest {
	s.PageSize = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetReportId(v string) *GetJMeterSamplingLogsRequest {
	s.ReportId = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetResponseCode(v string) *GetJMeterSamplingLogsRequest {
	s.ResponseCode = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetSamplerId(v int32) *GetJMeterSamplingLogsRequest {
	s.SamplerId = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetSuccess(v bool) *GetJMeterSamplingLogsRequest {
	s.Success = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetThread(v string) *GetJMeterSamplingLogsRequest {
	s.Thread = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) Validate() error {
	return dara.Validate(s)
}
