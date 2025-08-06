// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodStorageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodStorageDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodStorageDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodStorageDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeVodStorageDataRequest
	GetRegion() *string
	SetStartTime(v string) *DescribeVodStorageDataRequest
	GetStartTime() *string
	SetStorage(v string) *DescribeVodStorageDataRequest
	GetStorage() *string
	SetStorageType(v string) *DescribeVodStorageDataRequest
	GetStorageType() *string
}

type DescribeVodStorageDataRequest struct {
	// The ID of the application.
	//
	// 	- Default value: **app-1000000**.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-02-01T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region in which you want to query storage data. If you leave this parameter empty, data in all regions is returned. Separate multiple regions with commas (,). Valid values:
	//
	// 	- **cn-shanghai**: China (Shanghai)
	//
	// 	- **cn-beijing**: China (Beijing)
	//
	// 	- **eu-central-1**: Germany (Frankfurt)
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-02-01T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the Object Storage Service (OSS) bucket. If you leave this parameter empty, data of all buckets is returned. Separate multiple transcoding specifications with commas (,).
	//
	// example:
	//
	// bucket
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The storage class. Set the value to **OSS**.
	//
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeVodStorageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStorageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodStorageDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodStorageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodStorageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodStorageDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodStorageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodStorageDataRequest) GetStorage() *string {
	return s.Storage
}

func (s *DescribeVodStorageDataRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeVodStorageDataRequest) SetAppId(v string) *DescribeVodStorageDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetEndTime(v string) *DescribeVodStorageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetOwnerId(v int64) *DescribeVodStorageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetRegion(v string) *DescribeVodStorageDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetStartTime(v string) *DescribeVodStorageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetStorage(v string) *DescribeVodStorageDataRequest {
	s.Storage = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetStorageType(v string) *DescribeVodStorageDataRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeVodStorageDataRequest) Validate() error {
	return dara.Validate(s)
}
