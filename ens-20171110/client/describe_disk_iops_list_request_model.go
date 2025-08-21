// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskIopsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DescribeDiskIopsListRequest
	GetDiskId() *string
	SetEndTime(v string) *DescribeDiskIopsListRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDiskIopsListRequest
	GetStartTime() *string
}

type DescribeDiskIopsListRequest struct {
	// The ID of the disk. Format: d-\\*\\*\\*\\*\\*\\*\\*\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-5tzm9wnhzlhjzcbtxo465****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The beginning of the time range to query. Specify the time in the format of yyyy-MM-dd HH:mm:ss. The time range specified by the StartTime and EndTime parameters cannot exceed one day for a query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-12-14 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-12-14 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiskIopsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskIopsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskIopsListRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskIopsListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiskIopsListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiskIopsListRequest) SetDiskId(v string) *DescribeDiskIopsListRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskIopsListRequest) SetEndTime(v string) *DescribeDiskIopsListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiskIopsListRequest) SetStartTime(v string) *DescribeDiskIopsListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiskIopsListRequest) Validate() error {
	return dara.Validate(s)
}
