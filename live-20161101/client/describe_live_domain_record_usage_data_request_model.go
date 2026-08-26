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
	// The streaming domain name to query.
	//
	// - Supports single or batch domain queries. Separate multiple domain names with commas (,).
	//
	// - If this parameter is left empty, the merged data of all live streaming domain names is returned by default.
	//
	// - When you specify DomainName, make sure that the specified domain name is a live streaming domain name and that the caller has the required permissions on the domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format in UTC. Requirements:
	//
	// - The end time must be later than the start time (StartTime).
	//
	// - The maximum time span between the end time and the start time is 31 days. Requests that exceed 31 days fail and return an error.
	//
	// example:
	//
	// 2021-05-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the queried data. Unit: seconds. Valid values:
	//
	// - 60.
	//
	// - 300.
	//
	// - 3600.
	//
	// - 86400.
	//
	// >If this parameter is not specified or an unsupported value is specified, the default time granularity is 300 seconds for query spans within 31 days and 86400 seconds for query spans longer than 31 days.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region. Valid values:
	//
	// - **cn-beijing**: Beijing.
	//
	// - **cn-shanghai**: Shanghai.
	//
	// - **cn-shenzhen**: Shenzhen.
	//
	// - **cn-qingdao**: Qingdao.
	//
	// - **ap-southeast-1**: Singapore.
	//
	// - **eu-central-1**: Germany.
	//
	// - **ap-northeast-1**: Tokyo.
	//
	// - **ap-southeast-5**: Jakarta.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The grouping key. Valid values:
	//
	// - **domain**: groups query results by domain name.
	//
	// - **record_fmt**: groups query results by recording type.
	//
	// > You can specify one or more values. Separate multiple values with commas (,). Default value: `domain,record_fmt`. If this parameter is set to empty or `null`, the results are not grouped by the preceding keys.
	//
	// example:
	//
	// domain,record_fmt
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format in UTC.
	//
	// - The minimum data granularity is 5 minutes.
	//
	// - If this parameter is not specified, data of the last 24 hours is returned by default.
	//
	// >The start time can be set to a point in time within the last 90 days from the current time, accurate to the second.
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
