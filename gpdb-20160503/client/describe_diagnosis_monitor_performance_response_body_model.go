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
	// Details of query execution.
	Performances []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The threshold for the number of queries.
	//
	// example:
	//
	// 10000
	PerformancesThreshold *int32 `json:"PerformancesThreshold,omitempty" xml:"PerformancesThreshold,omitempty"`
	// Indicates whether the queries are truncated when the number of queries exceeds the threshold. Valid values:
	//
	// 	- **true**: The queries are truncated.
	//
	// 	- **false**: The queries are not truncated.
	//
	// example:
	//
	// false
	PerformancesTruncated *bool `json:"PerformancesTruncated,omitempty" xml:"PerformancesTruncated,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
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
	// The execution duration of the query. Unit: milliseconds.
	//
	// example:
	//
	// 1
	Cost *int32 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adbtest
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The ID of the query. It is a unique identifier of the query.
	//
	// example:
	//
	// 2022042612465401000000012903151998970
	QueryID *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
	// The start time of the query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1651877940000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution state of the query. Valid values:
	//
	// 	- **running**: The query is being executed.
	//
	// 	- **finished**: The query is complete.
	//
	// example:
	//
	// finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// adbpguser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetCost() *int32 {
	return s.Cost
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetQueryID() *string {
	return s.QueryID
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GetUser() *string {
	return s.User
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetCost(v int32) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Cost = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetDatabase(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetQueryID(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.QueryID = &v
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

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetUser(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.User = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) Validate() error {
	return dara.Validate(s)
}
