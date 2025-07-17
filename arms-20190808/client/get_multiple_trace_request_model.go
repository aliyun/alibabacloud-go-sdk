// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipleTraceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetMultipleTraceRequest
	GetEndTime() *int64
	SetPageNumber(v int64) *GetMultipleTraceRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetMultipleTraceRequest
	GetPageSize() *int64
	SetRegionId(v string) *GetMultipleTraceRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetMultipleTraceRequest
	GetStartTime() *int64
	SetTraceIDs(v []*string) *GetMultipleTraceRequest
	GetTraceIDs() []*string
}

type GetMultipleTraceRequest struct {
	// The time when the trace ends. The value is a timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1663999380000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. Default value: `1`.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page, the maximum value is 1000.
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
	// The start time of the trace. The value is a timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1657692507000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trace IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac1400a115951745017447033d****
	TraceIDs []*string `json:"TraceIDs,omitempty" xml:"TraceIDs,omitempty" type:"Repeated"`
}

func (s GetMultipleTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceRequest) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetMultipleTraceRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetMultipleTraceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetMultipleTraceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMultipleTraceRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetMultipleTraceRequest) GetTraceIDs() []*string {
	return s.TraceIDs
}

func (s *GetMultipleTraceRequest) SetEndTime(v int64) *GetMultipleTraceRequest {
	s.EndTime = &v
	return s
}

func (s *GetMultipleTraceRequest) SetPageNumber(v int64) *GetMultipleTraceRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMultipleTraceRequest) SetPageSize(v int64) *GetMultipleTraceRequest {
	s.PageSize = &v
	return s
}

func (s *GetMultipleTraceRequest) SetRegionId(v string) *GetMultipleTraceRequest {
	s.RegionId = &v
	return s
}

func (s *GetMultipleTraceRequest) SetStartTime(v int64) *GetMultipleTraceRequest {
	s.StartTime = &v
	return s
}

func (s *GetMultipleTraceRequest) SetTraceIDs(v []*string) *GetMultipleTraceRequest {
	s.TraceIDs = v
	return s
}

func (s *GetMultipleTraceRequest) Validate() error {
	return dara.Validate(s)
}
