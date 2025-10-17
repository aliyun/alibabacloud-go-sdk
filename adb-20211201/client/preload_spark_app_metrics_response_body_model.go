// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadSparkAppMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PreloadSparkAppMetricsResponseBodyData) *PreloadSparkAppMetricsResponseBody
	GetData() *PreloadSparkAppMetricsResponseBodyData
	SetRequestId(v string) *PreloadSparkAppMetricsResponseBody
	GetRequestId() *string
}

type PreloadSparkAppMetricsResponseBody struct {
	// The returned data.
	Data *PreloadSparkAppMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 84489769-3065-5A28-A4CB-977CD426F1C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreloadSparkAppMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreloadSparkAppMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadSparkAppMetricsResponseBody) GetData() *PreloadSparkAppMetricsResponseBodyData {
	return s.Data
}

func (s *PreloadSparkAppMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreloadSparkAppMetricsResponseBody) SetData(v *PreloadSparkAppMetricsResponseBodyData) *PreloadSparkAppMetricsResponseBody {
	s.Data = v
	return s
}

func (s *PreloadSparkAppMetricsResponseBody) SetRequestId(v string) *PreloadSparkAppMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreloadSparkAppMetricsResponseBodyData struct {
	// The ID of the Spark application.
	//
	// example:
	//
	// s202212181815shaccb8be0000253
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The retry ID of the Spark application.
	//
	// example:
	//
	// s202301061000hz57d797b0000201-0001
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// The event log path.
	//
	// example:
	//
	// oss://path/to/eventLog
	EventLogPath *string `json:"EventLogPath,omitempty" xml:"EventLogPath,omitempty"`
	// Indicates whether parsing is complete. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// True
	Finished *bool `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// The metrics.
	ScanMetrics *PreloadSparkAppMetricsResponseBodyDataScanMetrics `json:"ScanMetrics,omitempty" xml:"ScanMetrics,omitempty" type:"Struct"`
}

func (s PreloadSparkAppMetricsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PreloadSparkAppMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PreloadSparkAppMetricsResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *PreloadSparkAppMetricsResponseBodyData) GetAttemptId() *string {
	return s.AttemptId
}

func (s *PreloadSparkAppMetricsResponseBodyData) GetEventLogPath() *string {
	return s.EventLogPath
}

func (s *PreloadSparkAppMetricsResponseBodyData) GetFinished() *bool {
	return s.Finished
}

func (s *PreloadSparkAppMetricsResponseBodyData) GetScanMetrics() *PreloadSparkAppMetricsResponseBodyDataScanMetrics {
	return s.ScanMetrics
}

func (s *PreloadSparkAppMetricsResponseBodyData) SetAppId(v string) *PreloadSparkAppMetricsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyData) SetAttemptId(v string) *PreloadSparkAppMetricsResponseBodyData {
	s.AttemptId = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyData) SetEventLogPath(v string) *PreloadSparkAppMetricsResponseBodyData {
	s.EventLogPath = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyData) SetFinished(v bool) *PreloadSparkAppMetricsResponseBodyData {
	s.Finished = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyData) SetScanMetrics(v *PreloadSparkAppMetricsResponseBodyDataScanMetrics) *PreloadSparkAppMetricsResponseBodyData {
	s.ScanMetrics = v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyData) Validate() error {
	if s.ScanMetrics != nil {
		if err := s.ScanMetrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreloadSparkAppMetricsResponseBodyDataScanMetrics struct {
	// The number of rows scanned.
	//
	// example:
	//
	// 1000
	OutputRowsCount *int64 `json:"OutputRowsCount,omitempty" xml:"OutputRowsCount,omitempty"`
	// The size of the scanned data. Unit: bytes.
	//
	// example:
	//
	// 10000
	TotalReadFileSizeInByte *int64 `json:"TotalReadFileSizeInByte,omitempty" xml:"TotalReadFileSizeInByte,omitempty"`
}

func (s PreloadSparkAppMetricsResponseBodyDataScanMetrics) String() string {
	return dara.Prettify(s)
}

func (s PreloadSparkAppMetricsResponseBodyDataScanMetrics) GoString() string {
	return s.String()
}

func (s *PreloadSparkAppMetricsResponseBodyDataScanMetrics) GetOutputRowsCount() *int64 {
	return s.OutputRowsCount
}

func (s *PreloadSparkAppMetricsResponseBodyDataScanMetrics) GetTotalReadFileSizeInByte() *int64 {
	return s.TotalReadFileSizeInByte
}

func (s *PreloadSparkAppMetricsResponseBodyDataScanMetrics) SetOutputRowsCount(v int64) *PreloadSparkAppMetricsResponseBodyDataScanMetrics {
	s.OutputRowsCount = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyDataScanMetrics) SetTotalReadFileSizeInByte(v int64) *PreloadSparkAppMetricsResponseBodyDataScanMetrics {
	s.TotalReadFileSizeInByte = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyDataScanMetrics) Validate() error {
	return dara.Validate(s)
}
