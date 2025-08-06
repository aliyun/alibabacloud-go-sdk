// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTieringStorageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodTieringStorageDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodTieringStorageDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodTieringStorageDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeVodTieringStorageDataRequest
	GetRegion() *string
	SetStartTime(v string) *DescribeVodTieringStorageDataRequest
	GetStartTime() *string
	SetStorageClass(v string) *DescribeVodTieringStorageDataRequest
	GetStorageClass() *string
}

type DescribeVodTieringStorageDataRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time at which data is obtained. The end time must be later than the start time. The difference cannot exceed 31 days. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2023-05-29T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region in which you want to query data. If you leave this parameter empty, data in all regions is returned. Separate multiple regions with commas (,).
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. If you leave this parameter empty, data in the last 24 hours is queried.
	//
	// example:
	//
	// 2023-05-29T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The storage type. By default, all storage types are returned. Valid values:
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

func (s DescribeVodTieringStorageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTieringStorageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodTieringStorageDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodTieringStorageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodTieringStorageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodTieringStorageDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodTieringStorageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodTieringStorageDataRequest) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeVodTieringStorageDataRequest) SetAppId(v string) *DescribeVodTieringStorageDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodTieringStorageDataRequest) SetEndTime(v string) *DescribeVodTieringStorageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodTieringStorageDataRequest) SetOwnerId(v int64) *DescribeVodTieringStorageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodTieringStorageDataRequest) SetRegion(v string) *DescribeVodTieringStorageDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodTieringStorageDataRequest) SetStartTime(v string) *DescribeVodTieringStorageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodTieringStorageDataRequest) SetStorageClass(v string) *DescribeVodTieringStorageDataRequest {
	s.StorageClass = &v
	return s
}

func (s *DescribeVodTieringStorageDataRequest) Validate() error {
	return dara.Validate(s)
}
