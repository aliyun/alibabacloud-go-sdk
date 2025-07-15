// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoricalInstanceReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetHistoricalInstanceReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetHistoricalInstanceReportRequest
	GetInstanceId() *string
	SetMediaType(v string) *GetHistoricalInstanceReportRequest
	GetMediaType() *string
	SetStartTime(v int64) *GetHistoricalInstanceReportRequest
	GetStartTime() *int64
}

type GetHistoricalInstanceReportRequest struct {
	// example:
	//
	// 1532707199000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MediaType  *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetHistoricalInstanceReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalInstanceReportRequest) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetHistoricalInstanceReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHistoricalInstanceReportRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *GetHistoricalInstanceReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetHistoricalInstanceReportRequest) SetEndTime(v int64) *GetHistoricalInstanceReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetHistoricalInstanceReportRequest) SetInstanceId(v string) *GetHistoricalInstanceReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHistoricalInstanceReportRequest) SetMediaType(v string) *GetHistoricalInstanceReportRequest {
	s.MediaType = &v
	return s
}

func (s *GetHistoricalInstanceReportRequest) SetStartTime(v int64) *GetHistoricalInstanceReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetHistoricalInstanceReportRequest) Validate() error {
	return dara.Validate(s)
}
