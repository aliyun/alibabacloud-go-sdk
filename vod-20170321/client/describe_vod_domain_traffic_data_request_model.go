// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainTrafficDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVodDomainTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainTrafficDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainTrafficDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainTrafficDataRequest
	GetStartTime() *string
}

type DescribeVodDomainTrafficDataRequest struct {
	// The accelerated domain name to query.
	//
	// - If you do not specify this parameter, the pooled data of all accelerated domain names is returned by default.
	//
	// - Batch queries are supported. Separate multiple domain names with commas (,). You can query up to 500 domain names at a time.
	//
	// - You can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management > Distribution Acceleration > Domain Names*	- in the left-side navigation pane to view the accelerated domain names added to ApsaraVideo VOD. Alternatively, you can call the [DescribeVodUserDomains](https://help.aliyun.com/document_detail/455030.html) operation to query the list of accelerated domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of the data to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-01-20T14:59:58Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data to query. Unit: seconds. Valid values: **300**, **3600**, and **86400**. If you do not specify this parameter or specify an unsupported value, the default value is used. The supported time granularity varies based on the time span between StartTime and EndTime:
	//
	// - Less than 3 days: **300*	- (default), **3600**, and **86400**.
	//
	// - 3 to 31 days (excluding 31 days): **3600*	- (default) and **86400**.
	//
	// - More than 31 days: **86400*	- (default).
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The English name of the Internet service provider (ISP). If you do not specify this parameter, all ISPs are queried by default.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The English name of the region. If you do not specify this parameter, all regions are queried by default. Currently, only the China (Shanghai) region is supported.
	//
	// example:
	//
	// shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time of the data to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-01-20T13:59:58Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainTrafficDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainTrafficDataRequest) SetDomainName(v string) *DescribeVodDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetEndTime(v string) *DescribeVodDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetInterval(v string) *DescribeVodDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeVodDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeVodDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetOwnerId(v int64) *DescribeVodDomainTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetStartTime(v string) *DescribeVodDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
