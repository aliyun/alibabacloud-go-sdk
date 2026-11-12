// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSharePerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDataSharePerformanceRequest
	GetEndTime() *string
	SetKey(v string) *DescribeDataSharePerformanceRequest
	GetKey() *string
	SetRegionId(v string) *DescribeDataSharePerformanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDataSharePerformanceRequest
	GetResourceGroupId() *string
	SetStartTime(v string) *DescribeDataSharePerformanceRequest
	GetStartTime() *string
}

type DescribeDataSharePerformanceRequest struct {
	// The end time of the query. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm</i>Z format (UTC).
	//
	// example:
	//
	// 2022-08-03T15:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the performance metric. Separate multiple metrics with commas (,). Valid values:
	//
	// - **adbpg_datashare_topic_count**: the number of shared topics.
	//
	// - **adbpg_datashare_data_size_mb**: the data volume of data sharing.
	//
	// This parameter is required.
	//
	// example:
	//
	// adbpg_datashare_topic_count
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time of the query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm</i>Z format (UTC).
	//
	// example:
	//
	// 2022-08-03T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataSharePerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSharePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDataSharePerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDataSharePerformanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataSharePerformanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDataSharePerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDataSharePerformanceRequest) SetEndTime(v string) *DescribeDataSharePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetKey(v string) *DescribeDataSharePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetRegionId(v string) *DescribeDataSharePerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetResourceGroupId(v string) *DescribeDataSharePerformanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) SetStartTime(v string) *DescribeDataSharePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataSharePerformanceRequest) Validate() error {
	return dara.Validate(s)
}
