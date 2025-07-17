// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetTraceRequest
	GetEndTime() *int64
	SetPageNumber(v int64) *GetTraceRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetTraceRequest
	GetPageSize() *int64
	SetRegionId(v string) *GetTraceRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetTraceRequest
	GetStartTime() *int64
	SetTraceID(v string) *GetTraceRequest
	GetTraceID() *string
}

type GetTraceRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// > If the ID of the trace is 30 characters in length, this parameter is optional. Otherwise, this parameter is required.
	//
	// example:
	//
	// 1623827603000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on each page. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// > If the ID of the trace is 30 characters in length, this parameter is optional. Otherwise, this parameter is required.
	//
	// example:
	//
	// 1623827602000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trace ID. You can log on to the ARMS console and obtain the trace ID on the **Trace Query*	- page or **Interface Snapshot*	- tab.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac14001a15954493811405707d****
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTraceRequest) GoString() string {
	return s.String()
}

func (s *GetTraceRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetTraceRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetTraceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetTraceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTraceRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetTraceRequest) GetTraceID() *string {
	return s.TraceID
}

func (s *GetTraceRequest) SetEndTime(v int64) *GetTraceRequest {
	s.EndTime = &v
	return s
}

func (s *GetTraceRequest) SetPageNumber(v int64) *GetTraceRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTraceRequest) SetPageSize(v int64) *GetTraceRequest {
	s.PageSize = &v
	return s
}

func (s *GetTraceRequest) SetRegionId(v string) *GetTraceRequest {
	s.RegionId = &v
	return s
}

func (s *GetTraceRequest) SetStartTime(v int64) *GetTraceRequest {
	s.StartTime = &v
	return s
}

func (s *GetTraceRequest) SetTraceID(v string) *GetTraceRequest {
	s.TraceID = &v
	return s
}

func (s *GetTraceRequest) Validate() error {
	return dara.Validate(s)
}
