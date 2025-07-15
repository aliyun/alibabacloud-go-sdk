// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalInstanceReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListIntervalInstanceReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListIntervalInstanceReportRequest
	GetInstanceId() *string
	SetInterval(v string) *ListIntervalInstanceReportRequest
	GetInterval() *string
	SetStartTime(v int64) *ListIntervalInstanceReportRequest
	GetStartTime() *int64
}

type ListIntervalInstanceReportRequest struct {
	// example:
	//
	// 1620316799000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Hourly
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 1620230400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListIntervalInstanceReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalInstanceReportRequest) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListIntervalInstanceReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIntervalInstanceReportRequest) GetInterval() *string {
	return s.Interval
}

func (s *ListIntervalInstanceReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListIntervalInstanceReportRequest) SetEndTime(v int64) *ListIntervalInstanceReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) SetInstanceId(v string) *ListIntervalInstanceReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) SetInterval(v string) *ListIntervalInstanceReportRequest {
	s.Interval = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) SetStartTime(v int64) *ListIntervalInstanceReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) Validate() error {
	return dara.Validate(s)
}
