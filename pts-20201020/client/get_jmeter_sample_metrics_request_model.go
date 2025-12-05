// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterSampleMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *GetJMeterSampleMetricsRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *GetJMeterSampleMetricsRequest
	GetEndTime() *int64
	SetReportId(v string) *GetJMeterSampleMetricsRequest
	GetReportId() *string
	SetSamplerId(v int32) *GetJMeterSampleMetricsRequest
	GetSamplerId() *int32
}

type GetJMeterSampleMetricsRequest struct {
	// The beginning of the time range to query.
	//
	// example:
	//
	// 1637157070000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 1637157073000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The report ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7R4RE352
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The sampler ID. This parameter value starts from 0. If this parameter value is -1, the data of the whole scenario is returned.
	//
	// example:
	//
	// 0
	SamplerId *int32 `json:"SamplerId,omitempty" xml:"SamplerId,omitempty"`
}

func (s GetJMeterSampleMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterSampleMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetJMeterSampleMetricsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetJMeterSampleMetricsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetJMeterSampleMetricsRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetJMeterSampleMetricsRequest) GetSamplerId() *int32 {
	return s.SamplerId
}

func (s *GetJMeterSampleMetricsRequest) SetBeginTime(v int64) *GetJMeterSampleMetricsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetJMeterSampleMetricsRequest) SetEndTime(v int64) *GetJMeterSampleMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJMeterSampleMetricsRequest) SetReportId(v string) *GetJMeterSampleMetricsRequest {
	s.ReportId = &v
	return s
}

func (s *GetJMeterSampleMetricsRequest) SetSamplerId(v int32) *GetJMeterSampleMetricsRequest {
	s.SamplerId = &v
	return s
}

func (s *GetJMeterSampleMetricsRequest) Validate() error {
	return dara.Validate(s)
}
