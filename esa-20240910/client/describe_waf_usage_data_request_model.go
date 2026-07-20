// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWafUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeWafUsageDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeWafUsageDataRequest
	GetInterval() *string
	SetRecordName(v string) *DescribeWafUsageDataRequest
	GetRecordName() *string
	SetSiteId(v int64) *DescribeWafUsageDataRequest
	GetSiteId() *int64
	SetSplitBy(v string) *DescribeWafUsageDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeWafUsageDataRequest
	GetStartTime() *string
}

type DescribeWafUsageDataRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// >The end time must be later than the start time.
	//
	// example:
	//
	// 2022-08-10T23:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for the query data, in seconds.
	//
	// The valid values vary based on the time range specified by **StartTime*	- and **EndTime**:
	//
	// - Less than 3 days: Valid values are **300**, **3600**, and **86400**. Default value: **300**.
	//
	// - 3 to 31 days (exclusive of 31 days): Valid values are **3600*	- and **86400**. Default value: **3600**.
	//
	// - 31 days or more: The only valid value is **86400**. Default value: **86400**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The domain record name. You can call the [ListSites](~~ListSites~~) operation to obtain the domain record name.
	//
	// example:
	//
	// test.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The grouping key. You can set this parameter to **domain**.
	//
	// - **domain**: groups the data by domain name.
	//
	// example:
	//
	// domain
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2022-08-10T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeWafUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeWafUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeWafUsageDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeWafUsageDataRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *DescribeWafUsageDataRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeWafUsageDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeWafUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeWafUsageDataRequest) SetEndTime(v string) *DescribeWafUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeWafUsageDataRequest) SetInterval(v string) *DescribeWafUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeWafUsageDataRequest) SetRecordName(v string) *DescribeWafUsageDataRequest {
	s.RecordName = &v
	return s
}

func (s *DescribeWafUsageDataRequest) SetSiteId(v int64) *DescribeWafUsageDataRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeWafUsageDataRequest) SetSplitBy(v string) *DescribeWafUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeWafUsageDataRequest) SetStartTime(v string) *DescribeWafUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeWafUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
