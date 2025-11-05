// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorData(v []*DescribeDiskMonitorDataResponseBodyMonitorData) *DescribeDiskMonitorDataResponseBody
	GetMonitorData() []*DescribeDiskMonitorDataResponseBodyMonitorData
	SetRequestId(v string) *DescribeDiskMonitorDataResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDiskMonitorDataResponseBody
	GetTotalCount() *int64
}

type DescribeDiskMonitorDataResponseBody struct {
	// The near real-time monitoring data of the disk.
	MonitorData []*DescribeDiskMonitorDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Repeated"`
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
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataResponseBody) GetMonitorData() []*DescribeDiskMonitorDataResponseBodyMonitorData {
	return s.MonitorData
}

func (s *DescribeDiskMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskMonitorDataResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDiskMonitorDataResponseBody) SetMonitorData(v []*DescribeDiskMonitorDataResponseBodyMonitorData) *DescribeDiskMonitorDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeDiskMonitorDataResponseBody) SetRequestId(v string) *DescribeDiskMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBody) SetTotalCount(v int64) *DescribeDiskMonitorDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBody) Validate() error {
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

type DescribeDiskMonitorDataResponseBodyMonitorData struct {
	// The percentage of BPS.
	//
	// example:
	//
	// 80(%)
	BPSPercent *int64 `json:"BPSPercent,omitempty" xml:"BPSPercent,omitempty"`
	// The number of burst I/O operations.
	//
	// example:
	//
	// 0
	BurstIOCount *int64 `json:"BurstIOCount,omitempty" xml:"BurstIOCount,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-bp1bq5g3dxxo1x4o****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The percentage of IOPS.
	//
	// example:
	//
	// 80(%)
	IOPSPercent *int64 `json:"IOPSPercent,omitempty" xml:"IOPSPercent,omitempty"`
	// The read bandwidth of the disk. Unit: MByte/s.
	//
	// example:
	//
	// 10
	ReadBPS *int64 `json:"ReadBPS,omitempty" xml:"ReadBPS,omitempty"`
	// Read IO block size. Unit: Bytes
	//
	// example:
	//
	// 4096
	ReadBlockSize *int64 `json:"ReadBlockSize,omitempty" xml:"ReadBlockSize,omitempty"`
	// The maximum number of read IOPS.
	//
	// example:
	//
	// 2000
	ReadIOPS *int64 `json:"ReadIOPS,omitempty" xml:"ReadIOPS,omitempty"`
	// Read IO latency. Unit:  microsecond
	//
	// example:
	//
	// 100
	ReadLatency *int64 `json:"ReadLatency,omitempty" xml:"ReadLatency,omitempty"`
	// The timestamp that is used to query the near real-time monitoring data of the disk. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-06-01T08:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The write bandwidth of the disk. Unit: MByte/s.
	//
	// example:
	//
	// 204
	WriteBPS *int64 `json:"WriteBPS,omitempty" xml:"WriteBPS,omitempty"`
	// Write IO block size. Unit: Bytes
	//
	// example:
	//
	// 4096
	WriteBlockSize *int64 `json:"WriteBlockSize,omitempty" xml:"WriteBlockSize,omitempty"`
	// The maximum number of write IOPS.
	//
	// example:
	//
	// 2000
	WriteIOPS *int64 `json:"WriteIOPS,omitempty" xml:"WriteIOPS,omitempty"`
	// Write IO latency. Unit: microsecond
	//
	// example:
	//
	// 100
	WriteLatency *int64 `json:"WriteLatency,omitempty" xml:"WriteLatency,omitempty"`
}

func (s DescribeDiskMonitorDataResponseBodyMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetBPSPercent() *int64 {
	return s.BPSPercent
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetBurstIOCount() *int64 {
	return s.BurstIOCount
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetIOPSPercent() *int64 {
	return s.IOPSPercent
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetReadBPS() *int64 {
	return s.ReadBPS
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetReadBlockSize() *int64 {
	return s.ReadBlockSize
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetReadIOPS() *int64 {
	return s.ReadIOPS
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetReadLatency() *int64 {
	return s.ReadLatency
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetWriteBPS() *int64 {
	return s.WriteBPS
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetWriteBlockSize() *int64 {
	return s.WriteBlockSize
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetWriteIOPS() *int64 {
	return s.WriteIOPS
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetWriteLatency() *int64 {
	return s.WriteLatency
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetBPSPercent(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.BPSPercent = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetBurstIOCount(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.BurstIOCount = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetDiskId(v string) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetIOPSPercent(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.IOPSPercent = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetReadBPS(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.ReadBPS = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetReadBlockSize(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.ReadBlockSize = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetReadIOPS(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.ReadIOPS = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetReadLatency(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.ReadLatency = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetTimestamp(v string) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.Timestamp = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetWriteBPS(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.WriteBPS = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetWriteBlockSize(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.WriteBlockSize = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetWriteIOPS(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.WriteIOPS = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetWriteLatency(v int64) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.WriteLatency = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) Validate() error {
	return dara.Validate(s)
}
