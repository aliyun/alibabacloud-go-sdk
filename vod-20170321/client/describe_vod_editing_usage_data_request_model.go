// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodEditingUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodEditingUsageDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodEditingUsageDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodEditingUsageDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeVodEditingUsageDataRequest
	GetRegion() *string
	SetSpecification(v string) *DescribeVodEditingUsageDataRequest
	GetSpecification() *string
	SetStartTime(v string) *DescribeVodEditingUsageDataRequest
	GetStartTime() *string
}

type DescribeVodEditingUsageDataRequest struct {
	// The application ID. For more information, see [Multi-application](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2024-11-07T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region. By default, data from all regions is returned. You can specify multiple regions separated by commas (,).
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The output specification.
	//
	// example:
	//
	// H264.SD
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time must be in UTC.
	//
	// If you leave this parameter empty, data from the last 24 hours is queried by default.
	//
	// example:
	//
	// 2024-11-06T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodEditingUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodEditingUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodEditingUsageDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodEditingUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodEditingUsageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodEditingUsageDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodEditingUsageDataRequest) GetSpecification() *string {
	return s.Specification
}

func (s *DescribeVodEditingUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodEditingUsageDataRequest) SetAppId(v string) *DescribeVodEditingUsageDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodEditingUsageDataRequest) SetEndTime(v string) *DescribeVodEditingUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodEditingUsageDataRequest) SetOwnerId(v int64) *DescribeVodEditingUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodEditingUsageDataRequest) SetRegion(v string) *DescribeVodEditingUsageDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodEditingUsageDataRequest) SetSpecification(v string) *DescribeVodEditingUsageDataRequest {
	s.Specification = &v
	return s
}

func (s *DescribeVodEditingUsageDataRequest) SetStartTime(v string) *DescribeVodEditingUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodEditingUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
