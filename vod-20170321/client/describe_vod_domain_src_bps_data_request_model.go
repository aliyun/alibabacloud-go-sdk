// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainSrcBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainSrcBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainSrcBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainSrcBpsDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeVodDomainSrcBpsDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainSrcBpsDataRequest
	GetStartTime() *string
}

type DescribeVodDomainSrcBpsDataRequest struct {
	// The accelerated domain name to query.
	//
	// - If you do not specify this parameter, the pooled data of all accelerated domain names is returned by default.
	//
	// - Batch queries are supported. Separate multiple domain names with commas (,). You can specify up to 500 domain names at a time.
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com), and choose **Configuration Management > CDN Configuration > Domain Names*	- in the left-side navigation pane to view the accelerated domain names that you have added to ApsaraVideo VOD. You can also call the [DescribeVodUserDomains](~~DescribeVodUserDomains~~) operation to query the list of accelerated domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2022-04-26T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds. Valid values: **300**, **3600**, and **86400**. If you do not specify this parameter or specify an unsupported value, the default value is used. Based on the time span specified by `StartTime` and `EndTime`, the supported time granularity values are as follows:
	//
	// - Less than 3 days (exclusive): **300*	- (default), **3600**, and **86400**.
	//
	// - 3 to 31 days (exclusive): **3600*	- (default) and **86400**.
	//
	// - 31 days or more: **86400*	- (default).
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-04-25T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainSrcBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainSrcBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainSrcBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainSrcBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainSrcBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainSrcBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainSrcBpsDataRequest) SetDomainName(v string) *DescribeVodDomainSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataRequest) SetEndTime(v string) *DescribeVodDomainSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataRequest) SetInterval(v string) *DescribeVodDomainSrcBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataRequest) SetOwnerId(v int64) *DescribeVodDomainSrcBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataRequest) SetStartTime(v string) *DescribeVodDomainSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainSrcBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
