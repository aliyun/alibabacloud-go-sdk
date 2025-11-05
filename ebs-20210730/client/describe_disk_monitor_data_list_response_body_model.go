// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskMonitorDataListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorData(v []*DescribeDiskMonitorDataListResponseBodyMonitorData) *DescribeDiskMonitorDataListResponseBody
	GetMonitorData() []*DescribeDiskMonitorDataListResponseBodyMonitorData
	SetNextToken(v string) *DescribeDiskMonitorDataListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDiskMonitorDataListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDiskMonitorDataListResponseBody
	GetTotalCount() *int64
}

type DescribeDiskMonitorDataListResponseBody struct {
	// The near real-time monitoring data of the disks.
	MonitorData []*DescribeDiskMonitorDataListResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// e71d8a535bd9c****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskMonitorDataListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataListResponseBody) GetMonitorData() []*DescribeDiskMonitorDataListResponseBodyMonitorData {
	return s.MonitorData
}

func (s *DescribeDiskMonitorDataListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiskMonitorDataListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskMonitorDataListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDiskMonitorDataListResponseBody) SetMonitorData(v []*DescribeDiskMonitorDataListResponseBodyMonitorData) *DescribeDiskMonitorDataListResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBody) SetNextToken(v string) *DescribeDiskMonitorDataListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBody) SetRequestId(v string) *DescribeDiskMonitorDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBody) SetTotalCount(v int64) *DescribeDiskMonitorDataListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBody) Validate() error {
	if s.MonitorData != nil {
		for _, item := range s.MonitorData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiskMonitorDataListResponseBodyMonitorData struct {
	// The number of burst I/O operations.
	//
	// example:
	//
	// 2000
	BurstIOCount *int64 `json:"BurstIOCount,omitempty" xml:"BurstIOCount,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The beginning of the time range during which the performance of the disk bursts. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-06-01T08:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeDiskMonitorDataListResponseBodyMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataListResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataListResponseBodyMonitorData) GetBurstIOCount() *int64 {
	return s.BurstIOCount
}

func (s *DescribeDiskMonitorDataListResponseBodyMonitorData) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskMonitorDataListResponseBodyMonitorData) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeDiskMonitorDataListResponseBodyMonitorData) SetBurstIOCount(v int64) *DescribeDiskMonitorDataListResponseBodyMonitorData {
	s.BurstIOCount = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBodyMonitorData) SetDiskId(v string) *DescribeDiskMonitorDataListResponseBodyMonitorData {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBodyMonitorData) SetTimestamp(v string) *DescribeDiskMonitorDataListResponseBodyMonitorData {
	s.Timestamp = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponseBodyMonitorData) Validate() error {
	return dara.Validate(s)
}
