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
	// The application ID. If you have activated the multi-application feature, you can specify this parameter to query the storage usage of a specific application. If you do not specify this parameter, the total storage usage of all applications is returned. You can obtain the value of this parameter from the AppId response parameter of the [CreateAppInfo](~~CreateAppInfo~~) operation. For more information, see [Multi-application](https://help.aliyun.com/document_detail/113601.html).
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-02-01T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The storage region. By default, data of all regions is returned. You can specify multiple regions separated by commas (,). Valid values:
	//
	// - **cn-shanghai**: Shanghai.
	//
	// - **cn-beijing**: Beijing.
	//
	// - **eu-central-1**: Germany.
	//
	// - **ap-southeast-1**: Singapore.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The start of the time range to query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-02-01T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The storage name (Alibaba Cloud OSS bucket name). By default, data of all storage buckets is returned. You can specify multiple storage names separated by commas (,).
	//
	// example:
	//
	// bucket
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The storage type. Set the value to **OSS**.
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
