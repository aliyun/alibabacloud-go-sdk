// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeHttpCodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRealTimeHttpCodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeHttpCodeDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeVodDomainRealTimeHttpCodeDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainRealTimeHttpCodeDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainRealTimeHttpCodeDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainRealTimeHttpCodeDataRequest
	GetStartTime() *string
}

type DescribeVodDomainRealTimeHttpCodeDataRequest struct {
	// The accelerated domain name to query.
	//
	// - Separate multiple domain names with commas (,). You can query up to 100 domain names at a time.
	//
	// - When you query multiple domain names, the returned data is the aggregated data of all specified domain names.
	//
	// - You can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management > Distribution and Acceleration > Domain Names*	- in the left-side navigation pane to view the accelerated domain names added to ApsaraVideo VOD. Alternatively, you can call the [DescribeVodUserDomains](~~DescribeVodUserDomains~~) operation to query the list of accelerated domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of the data to retrieve. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The English name of the Internet service provider (ISP). If you do not specify this parameter, all ISPs are queried by default.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The English name of the region. If you do not specify this parameter, all regions are queried by default.
	//
	// example:
	//
	// shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time of the data to retrieve. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-11-30T05:39:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeHttpCodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) SetDomainName(v string) *DescribeVodDomainRealTimeHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) SetEndTime(v string) *DescribeVodDomainRealTimeHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) SetIspNameEn(v string) *DescribeVodDomainRealTimeHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeVodDomainRealTimeHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) SetOwnerId(v int64) *DescribeVodDomainRealTimeHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) SetStartTime(v string) *DescribeVodDomainRealTimeHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeHttpCodeDataRequest) Validate() error {
	return dara.Validate(s)
}
