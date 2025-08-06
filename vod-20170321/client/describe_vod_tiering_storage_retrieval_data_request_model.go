// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTieringStorageRetrievalDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodTieringStorageRetrievalDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodTieringStorageRetrievalDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodTieringStorageRetrievalDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeVodTieringStorageRetrievalDataRequest
	GetRegion() *string
	SetStartTime(v string) *DescribeVodTieringStorageRetrievalDataRequest
	GetStartTime() *string
	SetStorageClass(v string) *DescribeVodTieringStorageRetrievalDataRequest
	GetStorageClass() *string
}

type DescribeVodTieringStorageRetrievalDataRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The time range cannot exceed 31 days. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2023-06-02T11:20:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region in which you want to query data. If you leave this parameter empty, data in all regions is returned. Separate multiple regions with commas (,).
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The minimum time range is 5 minutes. If you leave this parameter empty, data in the last 24 hours is queried.
	//
	// example:
	//
	// 2023-06-02T10:20:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The storage type. Valid values:
	//
	// 	- **IA**
	//
	// 	- **Archive**
	//
	// 	- **ColdArchive**
	//
	// example:
	//
	// IA
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s DescribeVodTieringStorageRetrievalDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTieringStorageRetrievalDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) SetAppId(v string) *DescribeVodTieringStorageRetrievalDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) SetEndTime(v string) *DescribeVodTieringStorageRetrievalDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) SetOwnerId(v int64) *DescribeVodTieringStorageRetrievalDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) SetRegion(v string) *DescribeVodTieringStorageRetrievalDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) SetStartTime(v string) *DescribeVodTieringStorageRetrievalDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) SetStorageClass(v string) *DescribeVodTieringStorageRetrievalDataRequest {
	s.StorageClass = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataRequest) Validate() error {
	return dara.Validate(s)
}
