// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRangeDataByLocateAndIspServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest
	GetStartTime() *string
}

type DescribeVodRangeDataByLocateAndIspServiceRequest struct {
	// The accelerated domain name to query.
	//
	// You can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com), and choose **Configuration Management > CDN Configuration > Domain Names*	- in the left-side navigation pane to view the accelerated domain names that you have added to ApsaraVideo VOD. You can also call the [DescribeVodUserDomains](~~DescribeVodUserDomains~~) operation to query the list of accelerated domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time. The maximum time range between the start time and end time is 1 hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-29T03:55:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ISP name in English. If you do not specify this parameter, data of all ISPs is queried. You can specify only one ISP for each request.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The region name in English. If you do not specify this parameter, data of all regions is queried. Separate multiple region names with commas (,).
	//
	// example:
	//
	// shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-29T03:50:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodRangeDataByLocateAndIspServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRangeDataByLocateAndIspServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetDomainName(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetEndTime(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetIspNameEn(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetLocationNameEn(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetOwnerId(v int64) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) SetStartTime(v string) *DescribeVodRangeDataByLocateAndIspServiceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceRequest) Validate() error {
	return dara.Validate(s)
}
