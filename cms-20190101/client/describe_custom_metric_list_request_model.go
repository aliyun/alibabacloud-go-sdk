// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomMetricListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimension(v string) *DescribeCustomMetricListRequest
	GetDimension() *string
	SetGroupId(v string) *DescribeCustomMetricListRequest
	GetGroupId() *string
	SetMd5(v string) *DescribeCustomMetricListRequest
	GetMd5() *string
	SetMetricName(v string) *DescribeCustomMetricListRequest
	GetMetricName() *string
	SetPageNumber(v string) *DescribeCustomMetricListRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeCustomMetricListRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeCustomMetricListRequest
	GetRegionId() *string
}

type DescribeCustomMetricListRequest struct {
	// The dimensions based on which the resources are queried.
	//
	// example:
	//
	// {sampleName1=value1&amp;sampleName2=value2}
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The ID of the application group.
	//
	// For information about how to query the IDs of application groups, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// example:
	//
	// 7378****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The MD5 value of the HTTP request body. The MD5 value is a 128-bit hash value used to verify the uniqueness of the reported monitoring data.
	//
	// example:
	//
	// 97c25982d9745a231276bff27469****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the custom metric.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Pages start from page 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomMetricListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomMetricListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomMetricListRequest) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeCustomMetricListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeCustomMetricListRequest) GetMd5() *string {
	return s.Md5
}

func (s *DescribeCustomMetricListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeCustomMetricListRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeCustomMetricListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeCustomMetricListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomMetricListRequest) SetDimension(v string) *DescribeCustomMetricListRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetGroupId(v string) *DescribeCustomMetricListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetMd5(v string) *DescribeCustomMetricListRequest {
	s.Md5 = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetMetricName(v string) *DescribeCustomMetricListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetPageNumber(v string) *DescribeCustomMetricListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetPageSize(v string) *DescribeCustomMetricListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetRegionId(v string) *DescribeCustomMetricListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomMetricListRequest) Validate() error {
	return dara.Validate(s)
}
