// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisMonitorPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPerformances(v []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) *DescribeDiagnosisMonitorPerformanceResponseBody
	GetPerformances() []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances
	SetPerformancesThreshold(v int32) *DescribeDiagnosisMonitorPerformanceResponseBody
	GetPerformancesThreshold() *int32
	SetPerformancesTruncated(v bool) *DescribeDiagnosisMonitorPerformanceResponseBody
	GetPerformancesTruncated() *bool
	SetRequestId(v string) *DescribeDiagnosisMonitorPerformanceResponseBody
	GetRequestId() *string
}

type DescribeDiagnosisMonitorPerformanceResponseBody struct {
	// The monitoring information about queries displayed in Gantt charts.
	Performances []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The threshold for the number of queries displayed in a Gantt chart. Default value: 10000.
	//
	// >  Up to 10,000 queries can be displayed in a Gantt chart even if more queries exist.
	//
	// example:
	//
	// 10000
	PerformancesThreshold *int32 `json:"PerformancesThreshold,omitempty" xml:"PerformancesThreshold,omitempty"`
	// Indicates whether all queries are returned. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	PerformancesTruncated *bool `json:"PerformancesTruncated,omitempty" xml:"PerformancesTruncated,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0F1AC5FD-16E9-5399-B81F-5AC434B1D9F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) GetPerformances() []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	return s.Performances
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) GetPerformancesThreshold() *int32 {
	return s.PerformancesThreshold
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) GetPerformancesTruncated() *bool {
	return s.PerformancesTruncated
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformances(v []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformancesThreshold(v int32) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.PerformancesThreshold = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformancesTruncated(v bool) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.PerformancesTruncated = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetRequestId(v string) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) Validate() error {
	if s.Performances != nil {
		for _, item := range s.Performances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnosisMonitorPerformanceResponseBodyPerformances struct {
	// The total execution duration. Unit: milliseconds.
	//
	// >  This value is the cumulative value of the `QueuedTime`, `TotalPlanningTime`, and `ExecutionTime` parameters.
	//
	// example:
	//
	// 252
	Cost *int64 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The peak memory of the query. Unit: bytes.
	//
	// example:
	//
	// 123
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 202210311015270330101470300315153****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The IP address of the AnalyticDB for MySQL frontend node on which the SQL statement is executed.
	//
	// example:
	//
	// 192.168.XX.XX
	RcHost *string `json:"RcHost,omitempty" xml:"RcHost,omitempty"`
	// The number of rows scanned.
	//
	// example:
	//
	// 2345
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The amount of scanned data. Unit: bytes.
	//
	// example:
	//
	// 123
	ScanSize *int64 `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	// The execution start time of the SQL statement. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1669011260000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the SQL statement. Valid values:
	//
	// 	- **running**
	//
	// 	- **finished**
	//
	// 	- **failed**
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The database account that is used to submit the query.
	//
	// example:
	//
	// rpt
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetCost() *int64 {
	return s.Cost
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetPeakMemory() *int64 {
	return s.PeakMemory
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetRcHost() *string {
	return s.RcHost
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetScanRows() *int64 {
	return s.ScanRows
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetScanSize() *int64 {
	return s.ScanSize
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetCost(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Cost = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetPeakMemory(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetProcessId(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetRcHost(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.RcHost = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetScanRows(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.ScanRows = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetScanSize(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.ScanSize = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetStartTime(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetStatus(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetUserName(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.UserName = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) Validate() error {
	return dara.Validate(s)
}
