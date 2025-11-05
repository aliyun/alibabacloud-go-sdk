// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskMonitorDataListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskIds(v string) *DescribeDiskMonitorDataListRequest
	GetDiskIds() *string
	SetEndTime(v string) *DescribeDiskMonitorDataListRequest
	GetEndTime() *string
	SetMaxResults(v string) *DescribeDiskMonitorDataListRequest
	GetMaxResults() *string
	SetNextToken(v string) *DescribeDiskMonitorDataListRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDiskMonitorDataListRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDiskMonitorDataListRequest
	GetStartTime() *string
	SetType(v string) *DescribeDiskMonitorDataListRequest
	GetType() *string
}

type DescribeDiskMonitorDataListRequest struct {
	// The IDs of the disks. The value is a JSON array that contains multiple disk IDs. Separate the IDs with commas (,).
	//
	// example:
	//
	// ["d-bp67acfmxazb4p****","d-bp67acfmxazs5t****"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The end of the time range during which you want to query the near real-time monitoring data of the disks. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-06-01T05:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries per page. If you specify this parameter, both `MaxResults` and `NextToken` are used for a paged query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in this request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// e71d8a535bd9c****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the list of regions that support CloudLens for EBS.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range during which you want to query the near real-time monitoring data of the disks. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-06-01T03:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the monitoring data. Set the value to pro.
	//
	// pro: burst performance data, such as burst I/O operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// pro
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDiskMonitorDataListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataListRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *DescribeDiskMonitorDataListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiskMonitorDataListRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *DescribeDiskMonitorDataListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiskMonitorDataListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskMonitorDataListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiskMonitorDataListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDiskMonitorDataListRequest) SetDiskIds(v string) *DescribeDiskMonitorDataListRequest {
	s.DiskIds = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetEndTime(v string) *DescribeDiskMonitorDataListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetMaxResults(v string) *DescribeDiskMonitorDataListRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetNextToken(v string) *DescribeDiskMonitorDataListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetRegionId(v string) *DescribeDiskMonitorDataListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetStartTime(v string) *DescribeDiskMonitorDataListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) SetType(v string) *DescribeDiskMonitorDataListRequest {
	s.Type = &v
	return s
}

func (s *DescribeDiskMonitorDataListRequest) Validate() error {
	return dara.Validate(s)
}
