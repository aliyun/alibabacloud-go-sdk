// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRecordUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainRecordUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainRecordUsageDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveDomainRecordUsageDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeLiveDomainRecordUsageDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeLiveDomainRecordUsageDataRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeLiveDomainRecordUsageDataRequest
	GetRegionId() *string
	SetSplitBy(v string) *DescribeLiveDomainRecordUsageDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeLiveDomainRecordUsageDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainRecordUsageDataRequest struct {
	// The main streaming domain to query.
	//
	// 	- You can query one or more domain names. If you specify multiple domain names, separate them with commas (,).
	//
	// 	- If you leave this parameter empty, the data of all domain names within your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC. Value requirements:
	//
	// 	- The end time is later than the start time.
	//
	// 	- The time range between the start time and end time is up to 31 days. If the time range is more than 31 days, the request fails and an error is reported.
	//
	// example:
	//
	// 2021-05-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- 60
	//
	// 	- 300
	//
	// 	- 3600
	//
	// 	- 86400
	//
	// > If you do not specify this parameter or specify an invalid value: The time granularity of the query for a time range that is less than or equal to 31 days is 300 seconds by default. The time granularity of the query for a time range that is more than 31 days is 86400 seconds by default.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. Valid values:
	//
	// 	- **cn-beijing**: China (Beijing)
	//
	// 	- **cn-shanghai**: China (Shanghai)
	//
	// 	- **cn-shenzhen**: China (Shenzhen)
	//
	// 	- **cn-qingdao**: China (Qingdao)
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// 	- **eu-central-1**: Germany (Frankfurt)
	//
	// 	- **ap-northeast-1**: Japan (Tokyo)
	//
	// 	- **ap-southeast-5**: Indonesia (Jakarta)
	//
	// example:
	//
	// cn-shanghai
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The key that is used to group data. Valid values:
	//
	// 	- **domain**: groups results by domain name.
	//
	// 	- **record_fmt**: groups results by recording type.
	//
	// >  You can select one option or both. If you want to select both options, separate them with a comma (,). The default value is `domain,record_fmt`. If you leave this parameter empty or set the value to `null`, this parameter is ignored.
	//
	// example:
	//
	// domain,record_fmt
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// 	- The minimum data granularity is 5 minutes.
	//
	// 	- If you do not specify this parameter, the data in the last 24 hours is returned.
	//
	// > The earliest start time that you can specify is 90 days back from the current time, accurate to seconds.
	//
	// example:
	//
	// 2021-05-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainRecordUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRecordUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRecordUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainRecordUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainRecordUsageDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDomainRecordUsageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainRecordUsageDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDomainRecordUsageDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainRecordUsageDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeLiveDomainRecordUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainRecordUsageDataRequest) SetDomainName(v string) *DescribeLiveDomainRecordUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataRequest) SetEndTime(v string) *DescribeLiveDomainRecordUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataRequest) SetInterval(v string) *DescribeLiveDomainRecordUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataRequest) SetOwnerId(v int64) *DescribeLiveDomainRecordUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataRequest) SetRegion(v string) *DescribeLiveDomainRecordUsageDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataRequest) SetRegionId(v string) *DescribeLiveDomainRecordUsageDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataRequest) SetSplitBy(v string) *DescribeLiveDomainRecordUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataRequest) SetStartTime(v string) *DescribeLiveDomainRecordUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
