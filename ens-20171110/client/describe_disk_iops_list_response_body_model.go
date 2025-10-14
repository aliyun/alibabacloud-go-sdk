// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskIopsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiskIopsList(v []*DescribeDiskIopsListResponseBodyDiskIopsList) *DescribeDiskIopsListResponseBody
	GetDiskIopsList() []*DescribeDiskIopsListResponseBodyDiskIopsList
	SetRequestId(v string) *DescribeDiskIopsListResponseBody
	GetRequestId() *string
}

type DescribeDiskIopsListResponseBody struct {
	// The IOPS monitoring data of the cloud disk.
	DiskIopsList []*DescribeDiskIopsListResponseBodyDiskIopsList `json:"DiskIopsList,omitempty" xml:"DiskIopsList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A7814CAB-DB4E-140A-9D6F-7C8210C1DAC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiskIopsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskIopsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskIopsListResponseBody) GetDiskIopsList() []*DescribeDiskIopsListResponseBodyDiskIopsList {
	return s.DiskIopsList
}

func (s *DescribeDiskIopsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskIopsListResponseBody) SetDiskIopsList(v []*DescribeDiskIopsListResponseBodyDiskIopsList) *DescribeDiskIopsListResponseBody {
	s.DiskIopsList = v
	return s
}

func (s *DescribeDiskIopsListResponseBody) SetRequestId(v string) *DescribeDiskIopsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskIopsListResponseBody) Validate() error {
	if s.DiskIopsList != nil {
		for _, item := range s.DiskIopsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiskIopsListResponseBodyDiskIopsList struct {
	// The business time . The time is displayed in the yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2023-12-14 00:00:00
	BizTime *string `json:"BizTime,omitempty" xml:"BizTime,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-5tzm9wnhzlhjzcbtxo465****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The read throughput. Unit: bytes.
	//
	// example:
	//
	// 10054
	ReadBytes *int64 `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	// The read latency. Unit: ms.
	//
	// example:
	//
	// 15646532
	ReadLatency *int64 `json:"ReadLatency,omitempty" xml:"ReadLatency,omitempty"`
	// The read IOPS.
	//
	// example:
	//
	// 4
	ReadOps *int64 `json:"ReadOps,omitempty" xml:"ReadOps,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cn-hangzhou-3
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The write throughput. Unit: bytes.
	//
	// example:
	//
	// 0
	WriteBytes *int64 `json:"WriteBytes,omitempty" xml:"WriteBytes,omitempty"`
	// The write latency. Unit: microseconds.
	//
	// example:
	//
	// 0
	WriteLatency *int64 `json:"WriteLatency,omitempty" xml:"WriteLatency,omitempty"`
	// The write IOPS.
	//
	// example:
	//
	// 0
	WriteOps *int64 `json:"WriteOps,omitempty" xml:"WriteOps,omitempty"`
}

func (s DescribeDiskIopsListResponseBodyDiskIopsList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskIopsListResponseBodyDiskIopsList) GoString() string {
	return s.String()
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) GetBizTime() *string {
	return s.BizTime
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) GetReadBytes() *int64 {
	return s.ReadBytes
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) GetReadLatency() *int64 {
	return s.ReadLatency
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) GetReadOps() *int64 {
	return s.ReadOps
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) GetWriteBytes() *int64 {
	return s.WriteBytes
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) GetWriteLatency() *int64 {
	return s.WriteLatency
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) GetWriteOps() *int64 {
	return s.WriteOps
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) SetBizTime(v string) *DescribeDiskIopsListResponseBodyDiskIopsList {
	s.BizTime = &v
	return s
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) SetDiskId(v string) *DescribeDiskIopsListResponseBodyDiskIopsList {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) SetReadBytes(v int64) *DescribeDiskIopsListResponseBodyDiskIopsList {
	s.ReadBytes = &v
	return s
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) SetReadLatency(v int64) *DescribeDiskIopsListResponseBodyDiskIopsList {
	s.ReadLatency = &v
	return s
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) SetReadOps(v int64) *DescribeDiskIopsListResponseBodyDiskIopsList {
	s.ReadOps = &v
	return s
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) SetRegionId(v string) *DescribeDiskIopsListResponseBodyDiskIopsList {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) SetWriteBytes(v int64) *DescribeDiskIopsListResponseBodyDiskIopsList {
	s.WriteBytes = &v
	return s
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) SetWriteLatency(v int64) *DescribeDiskIopsListResponseBodyDiskIopsList {
	s.WriteLatency = &v
	return s
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) SetWriteOps(v int64) *DescribeDiskIopsListResponseBodyDiskIopsList {
	s.WriteOps = &v
	return s
}

func (s *DescribeDiskIopsListResponseBodyDiskIopsList) Validate() error {
	return dara.Validate(s)
}
