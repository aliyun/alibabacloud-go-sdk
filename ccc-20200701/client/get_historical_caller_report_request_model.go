// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoricalCallerReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallingNumber(v string) *GetHistoricalCallerReportRequest
	GetCallingNumber() *string
	SetInstanceId(v string) *GetHistoricalCallerReportRequest
	GetInstanceId() *string
	SetStartTime(v int64) *GetHistoricalCallerReportRequest
	GetStartTime() *int64
	SetStopTime(v int64) *GetHistoricalCallerReportRequest
	GetStopTime() *int64
}

type GetHistoricalCallerReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1900000****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1646841600000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1646928000000
	StopTime *int64 `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
}

func (s GetHistoricalCallerReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalCallerReportRequest) GoString() string {
	return s.String()
}

func (s *GetHistoricalCallerReportRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *GetHistoricalCallerReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHistoricalCallerReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetHistoricalCallerReportRequest) GetStopTime() *int64 {
	return s.StopTime
}

func (s *GetHistoricalCallerReportRequest) SetCallingNumber(v string) *GetHistoricalCallerReportRequest {
	s.CallingNumber = &v
	return s
}

func (s *GetHistoricalCallerReportRequest) SetInstanceId(v string) *GetHistoricalCallerReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHistoricalCallerReportRequest) SetStartTime(v int64) *GetHistoricalCallerReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetHistoricalCallerReportRequest) SetStopTime(v int64) *GetHistoricalCallerReportRequest {
	s.StopTime = &v
	return s
}

func (s *GetHistoricalCallerReportRequest) Validate() error {
	return dara.Validate(s)
}
