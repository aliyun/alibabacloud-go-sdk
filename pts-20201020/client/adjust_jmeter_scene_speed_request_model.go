// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdjustJMeterSceneSpeedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *AdjustJMeterSceneSpeedRequest
	GetReportId() *string
	SetSpeed(v int32) *AdjustJMeterSceneSpeedRequest
	GetSpeed() *int32
}

type AdjustJMeterSceneSpeedRequest struct {
	// The ID of the report.
	//
	// This parameter is required.
	//
	// example:
	//
	// DYYPZIH
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The load to which you want to adjust.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
}

func (s AdjustJMeterSceneSpeedRequest) String() string {
	return dara.Prettify(s)
}

func (s AdjustJMeterSceneSpeedRequest) GoString() string {
	return s.String()
}

func (s *AdjustJMeterSceneSpeedRequest) GetReportId() *string {
	return s.ReportId
}

func (s *AdjustJMeterSceneSpeedRequest) GetSpeed() *int32 {
	return s.Speed
}

func (s *AdjustJMeterSceneSpeedRequest) SetReportId(v string) *AdjustJMeterSceneSpeedRequest {
	s.ReportId = &v
	return s
}

func (s *AdjustJMeterSceneSpeedRequest) SetSpeed(v int32) *AdjustJMeterSceneSpeedRequest {
	s.Speed = &v
	return s
}

func (s *AdjustJMeterSceneSpeedRequest) Validate() error {
	return dara.Validate(s)
}
