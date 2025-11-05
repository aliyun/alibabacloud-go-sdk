// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskCategory(v string) *DescribeDiskEventsRequest
	GetDiskCategory() *string
	SetDiskId(v string) *DescribeDiskEventsRequest
	GetDiskId() *string
	SetEndTime(v string) *DescribeDiskEventsRequest
	GetEndTime() *string
	SetMaxResults(v int64) *DescribeDiskEventsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeDiskEventsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDiskEventsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDiskEventsRequest
	GetStartTime() *string
	SetType(v string) *DescribeDiskEventsRequest
	GetType() *string
}

type DescribeDiskEventsRequest struct {
	// The type of the disk. Valid values:
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud_essd: enhanced SSD (ESSD).
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-06-01T05:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100.
	//
	// Default values:
	//
	// 	- If this parameter is not specified or is set to a value smaller than 10, the default value is 10.
	//
	// 	- If this parameter is set to a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in this request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the disk. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the list of regions that support CloudLens for EBS.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-06-01T03:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event type. Set the value to DataNeedProtect, which indicates that the disk data needs to be protected.
	//
	// example:
	//
	// DataNeedProtect
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDiskEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskEventsRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *DescribeDiskEventsRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiskEventsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeDiskEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiskEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiskEventsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDiskEventsRequest) SetDiskCategory(v string) *DescribeDiskEventsRequest {
	s.DiskCategory = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetDiskId(v string) *DescribeDiskEventsRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetEndTime(v string) *DescribeDiskEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetMaxResults(v int64) *DescribeDiskEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetNextToken(v string) *DescribeDiskEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetRegionId(v string) *DescribeDiskEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetStartTime(v string) *DescribeDiskEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiskEventsRequest) SetType(v string) *DescribeDiskEventsRequest {
	s.Type = &v
	return s
}

func (s *DescribeDiskEventsRequest) Validate() error {
	return dara.Validate(s)
}
