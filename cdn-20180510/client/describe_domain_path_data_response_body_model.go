// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainPathDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainPathDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainPathDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainPathDataResponseBody
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeDomainPathDataResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDomainPathDataResponseBody
	GetPageSize() *int32
	SetPathDataPerInterval(v *DescribeDomainPathDataResponseBodyPathDataPerInterval) *DescribeDomainPathDataResponseBody
	GetPathDataPerInterval() *DescribeDomainPathDataResponseBodyPathDataPerInterval
	SetRequestId(v string) *DescribeDomainPathDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainPathDataResponseBody
	GetStartTime() *string
	SetTotalCount(v int32) *DescribeDomainPathDataResponseBody
	GetTotalCount() *int32
}

type DescribeDomainPathDataResponseBody struct {
	// The time interval. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2017-09-30T17:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number of the returned page. Pages start from page **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A list of bandwidth values collected at each time interval.
	PathDataPerInterval *DescribeDomainPathDataResponseBodyPathDataPerInterval `json:"PathDataPerInterval,omitempty" xml:"PathDataPerInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 809A6F10-8238-4A49-BE00-4B2D6305686E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2017-09-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainPathDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainPathDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainPathDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainPathDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainPathDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainPathDataResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDomainPathDataResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDomainPathDataResponseBody) GetPathDataPerInterval() *DescribeDomainPathDataResponseBodyPathDataPerInterval {
	return s.PathDataPerInterval
}

func (s *DescribeDomainPathDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainPathDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainPathDataResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDomainPathDataResponseBody) SetDataInterval(v string) *DescribeDomainPathDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetDomainName(v string) *DescribeDomainPathDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetEndTime(v string) *DescribeDomainPathDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetPageNumber(v int32) *DescribeDomainPathDataResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetPageSize(v int32) *DescribeDomainPathDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetPathDataPerInterval(v *DescribeDomainPathDataResponseBodyPathDataPerInterval) *DescribeDomainPathDataResponseBody {
	s.PathDataPerInterval = v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetRequestId(v string) *DescribeDomainPathDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetStartTime(v string) *DescribeDomainPathDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetTotalCount(v int32) *DescribeDomainPathDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainPathDataResponseBodyPathDataPerInterval struct {
	UsageData []*DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainPathDataResponseBodyPathDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainPathDataResponseBodyPathDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerInterval) GetUsageData() []*DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData {
	return s.UsageData
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerInterval) SetUsageData(v []*DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) *DescribeDomainPathDataResponseBodyPathDataPerInterval {
	s.UsageData = v
	return s
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData struct {
	// The number of visits to the URL.
	//
	// example:
	//
	// 10
	Acc *int32 `json:"Acc,omitempty" xml:"Acc,omitempty"`
	// The path.
	//
	// example:
	//
	// /path/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The point in time.
	//
	// example:
	//
	// 2017-09-30T16:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The amount of network traffic. Unit: bytes.
	//
	// example:
	//
	// 346
	Traffic *int32 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) GetAcc() *int32 {
	return s.Acc
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) GetPath() *string {
	return s.Path
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) GetTime() *string {
	return s.Time
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) GetTraffic() *int32 {
	return s.Traffic
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) SetAcc(v int32) *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData {
	s.Acc = &v
	return s
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) SetPath(v string) *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData {
	s.Path = &v
	return s
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) SetTime(v string) *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData {
	s.Time = &v
	return s
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) SetTraffic(v int32) *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData {
	s.Traffic = &v
	return s
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) Validate() error {
	return dara.Validate(s)
}
