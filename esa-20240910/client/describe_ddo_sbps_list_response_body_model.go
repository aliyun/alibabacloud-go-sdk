// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSBpsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v int32) *DescribeDDoSBpsListResponseBody
	GetDataInterval() *int32
	SetDataModule(v []*DescribeDDoSBpsListResponseBodyDataModule) *DescribeDDoSBpsListResponseBody
	GetDataModule() []*DescribeDDoSBpsListResponseBodyDataModule
	SetEndTime(v string) *DescribeDDoSBpsListResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDDoSBpsListResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDDoSBpsListResponseBody
	GetStartTime() *string
}

type DescribeDDoSBpsListResponseBody struct {
	// The interval between each piece of data, in seconds.
	//
	// Generated based on the interval between StartTime and EndTime: less than 1 hour, 60s; 1 hour or more but less than 1 day, 300s; 1 day or more but less than a week, 1800s; 1 week or more, 3600s.
	//
	// example:
	//
	// 300
	DataInterval *int32 `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// A list of network bandwidth data for each time interval.
	DataModule []*DescribeDDoSBpsListResponseBodyDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
	// The end time for fetching data. In ISO8601 format, using UTC+0, formatted as: yyyy-MM-ddTHH:mm:ssZ.
	//
	// The end time must be later than the start time, and the span between start and end times should not exceed 31 days.
	//
	// example:
	//
	// 2023-05-18T06:19:42Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 156A6B-677B1A-4297B7-9187B7-2B44792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time for fetching data. In ISO8601 format, using UTC, formatted as: YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2023-05-14T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSBpsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSBpsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSBpsListResponseBody) GetDataInterval() *int32 {
	return s.DataInterval
}

func (s *DescribeDDoSBpsListResponseBody) GetDataModule() []*DescribeDDoSBpsListResponseBodyDataModule {
	return s.DataModule
}

func (s *DescribeDDoSBpsListResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSBpsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSBpsListResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSBpsListResponseBody) SetDataInterval(v int32) *DescribeDDoSBpsListResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBody) SetDataModule(v []*DescribeDDoSBpsListResponseBodyDataModule) *DescribeDDoSBpsListResponseBody {
	s.DataModule = v
	return s
}

func (s *DescribeDDoSBpsListResponseBody) SetEndTime(v string) *DescribeDDoSBpsListResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBody) SetRequestId(v string) *DescribeDDoSBpsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBody) SetStartTime(v string) *DescribeDDoSBpsListResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBody) Validate() error {
	if s.DataModule != nil {
		for _, item := range s.DataModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDDoSBpsListResponseBodyDataModule struct {
	// Attack bandwidth, in bps.
	//
	// example:
	//
	// 9000000000
	AttackBps *int64 `json:"AttackBps,omitempty" xml:"AttackBps,omitempty"`
	// Attack PPS.
	//
	// example:
	//
	// 9000000
	AttackPps *int64 `json:"AttackPps,omitempty" xml:"AttackPps,omitempty"`
	// Normal business bandwidth, in bps.
	//
	// example:
	//
	// 1000000000
	NormalBps *int64 `json:"NormalBps,omitempty" xml:"NormalBps,omitempty"`
	// Normal business PPS.
	//
	// example:
	//
	// 1000000
	NormalPps *int64 `json:"NormalPps,omitempty" xml:"NormalPps,omitempty"`
	// The timestamp of this data, in ISO8601 format, using UTC+0, formatted as: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-05-14T17:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// Total bandwidth, in bps.
	//
	// example:
	//
	// 10000000000
	TotalBps *int64 `json:"TotalBps,omitempty" xml:"TotalBps,omitempty"`
	// Total PPS.
	//
	// example:
	//
	// 100000000
	TotalPps *int64 `json:"TotalPps,omitempty" xml:"TotalPps,omitempty"`
}

func (s DescribeDDoSBpsListResponseBodyDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSBpsListResponseBodyDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetAttackBps() *int64 {
	return s.AttackBps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetAttackPps() *int64 {
	return s.AttackPps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetNormalBps() *int64 {
	return s.NormalBps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetNormalPps() *int64 {
	return s.NormalPps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetTotalBps() *int64 {
	return s.TotalBps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) GetTotalPps() *int64 {
	return s.TotalPps
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetAttackBps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.AttackBps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetAttackPps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.AttackPps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetNormalBps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.NormalBps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetNormalPps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.NormalPps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetTimeStamp(v string) *DescribeDDoSBpsListResponseBodyDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetTotalBps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.TotalBps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) SetTotalPps(v int64) *DescribeDDoSBpsListResponseBodyDataModule {
	s.TotalPps = &v
	return s
}

func (s *DescribeDDoSBpsListResponseBodyDataModule) Validate() error {
	return dara.Validate(s)
}
