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
	// The application ID.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query. The end time must be later than the start time. The time range cannot exceed 31 days. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2023-05-29T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The storage region. By default, data of all regions is returned. You can specify multiple regions separated by commas (,).
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time must be in UTC. If this parameter is left empty, data of the last 24 hours is returned by default.
	//
	// example:
	//
	// 2023-05-29T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The storage class. By default, data of all storage classes is returned. Valid values:
	//
	// - **IA**: Infrequent Access.
	//
	// - **Archive**: Archive.
	//
	// - **ColdArchive**: Cold Archive.
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
