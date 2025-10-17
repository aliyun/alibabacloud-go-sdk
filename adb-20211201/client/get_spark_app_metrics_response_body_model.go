// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSparkAppMetricsResponseBodyData) *GetSparkAppMetricsResponseBody
	GetData() *GetSparkAppMetricsResponseBodyData
	SetRequestId(v string) *GetSparkAppMetricsResponseBody
	GetRequestId() *string
}

type GetSparkAppMetricsResponseBody struct {
	// The returned data.
	Data *GetSparkAppMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppMetricsResponseBody) GetData() *GetSparkAppMetricsResponseBodyData {
	return s.Data
}

func (s *GetSparkAppMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkAppMetricsResponseBody) SetData(v *GetSparkAppMetricsResponseBodyData) *GetSparkAppMetricsResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppMetricsResponseBody) SetRequestId(v string) *GetSparkAppMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkAppMetricsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSparkAppMetricsResponseBodyData struct {
	// The ID of the Spark application.
	//
	// example:
	//
	// s202302051515shfa865f80003691
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The attempt ID of the Spark application.
	//
	// example:
	//
	// s202301061000hz57d797b0000201-0001
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// The path of the event log.
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
	ScanMetrics *GetSparkAppMetricsResponseBodyDataScanMetrics `json:"ScanMetrics,omitempty" xml:"ScanMetrics,omitempty" type:"Struct"`
}

func (s GetSparkAppMetricsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkAppMetricsResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkAppMetricsResponseBodyData) GetAttemptId() *string {
	return s.AttemptId
}

func (s *GetSparkAppMetricsResponseBodyData) GetEventLogPath() *string {
	return s.EventLogPath
}

func (s *GetSparkAppMetricsResponseBodyData) GetFinished() *bool {
	return s.Finished
}

func (s *GetSparkAppMetricsResponseBodyData) GetScanMetrics() *GetSparkAppMetricsResponseBodyDataScanMetrics {
	return s.ScanMetrics
}

func (s *GetSparkAppMetricsResponseBodyData) SetAppId(v string) *GetSparkAppMetricsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyData) SetAttemptId(v string) *GetSparkAppMetricsResponseBodyData {
	s.AttemptId = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyData) SetEventLogPath(v string) *GetSparkAppMetricsResponseBodyData {
	s.EventLogPath = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyData) SetFinished(v bool) *GetSparkAppMetricsResponseBodyData {
	s.Finished = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyData) SetScanMetrics(v *GetSparkAppMetricsResponseBodyDataScanMetrics) *GetSparkAppMetricsResponseBodyData {
	s.ScanMetrics = v
	return s
}

func (s *GetSparkAppMetricsResponseBodyData) Validate() error {
	if s.ScanMetrics != nil {
		if err := s.ScanMetrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSparkAppMetricsResponseBodyDataScanMetrics struct {
	// The number of scanned rows.
	//
	// example:
	//
	// 1000
	OutputRowsCount *int64 `json:"OutputRowsCount,omitempty" xml:"OutputRowsCount,omitempty"`
	// The number of scanned bytes.
	//
	// example:
	//
	// 10000
	TotalReadFileSizeInByte *int64 `json:"TotalReadFileSizeInByte,omitempty" xml:"TotalReadFileSizeInByte,omitempty"`
}

func (s GetSparkAppMetricsResponseBodyDataScanMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppMetricsResponseBodyDataScanMetrics) GoString() string {
	return s.String()
}

func (s *GetSparkAppMetricsResponseBodyDataScanMetrics) GetOutputRowsCount() *int64 {
	return s.OutputRowsCount
}

func (s *GetSparkAppMetricsResponseBodyDataScanMetrics) GetTotalReadFileSizeInByte() *int64 {
	return s.TotalReadFileSizeInByte
}

func (s *GetSparkAppMetricsResponseBodyDataScanMetrics) SetOutputRowsCount(v int64) *GetSparkAppMetricsResponseBodyDataScanMetrics {
	s.OutputRowsCount = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyDataScanMetrics) SetTotalReadFileSizeInByte(v int64) *GetSparkAppMetricsResponseBodyDataScanMetrics {
	s.TotalReadFileSizeInByte = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyDataScanMetrics) Validate() error {
	return dara.Validate(s)
}
