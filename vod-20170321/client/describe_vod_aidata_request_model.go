// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodAIDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIType(v string) *DescribeVodAIDataRequest
	GetAIType() *string
	SetEndTime(v string) *DescribeVodAIDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodAIDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeVodAIDataRequest
	GetRegion() *string
	SetStartTime(v string) *DescribeVodAIDataRequest
	GetStartTime() *string
}

type DescribeVodAIDataRequest struct {
	// The type of video AI. If you leave this parameter empty, statistics on video AI of all types are returned. Separate multiple types with commas (,). Valid values:
	//
	// 	- **AIVideoCensor**: automated review
	//
	// 	- **AIVideoFPShot**: media fingerprinting
	//
	// 	- **AIVideoTag**: smart tagging
	//
	// example:
	//
	// AIVideoCensor
	AIType *string `json:"AIType,omitempty" xml:"AIType,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-02-01T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region in which you want to query data. If you leave this parameter empty, data in all regions is returned. Separate multiple regions with commas (,). Valid values:
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
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-02-01T13:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodAIDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAIDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataRequest) GetAIType() *string {
	return s.AIType
}

func (s *DescribeVodAIDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodAIDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodAIDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodAIDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodAIDataRequest) SetAIType(v string) *DescribeVodAIDataRequest {
	s.AIType = &v
	return s
}

func (s *DescribeVodAIDataRequest) SetEndTime(v string) *DescribeVodAIDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodAIDataRequest) SetOwnerId(v int64) *DescribeVodAIDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodAIDataRequest) SetRegion(v string) *DescribeVodAIDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodAIDataRequest) SetStartTime(v string) *DescribeVodAIDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodAIDataRequest) Validate() error {
	return dara.Validate(s)
}
