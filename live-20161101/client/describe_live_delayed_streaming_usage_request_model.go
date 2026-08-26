// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDelayedStreamingUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDelayedStreamingUsageRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDelayedStreamingUsageRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveDelayedStreamingUsageRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeLiveDelayedStreamingUsageRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeLiveDelayedStreamingUsageRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeLiveDelayedStreamingUsageRequest
	GetRegionId() *string
	SetSplitBy(v string) *DescribeLiveDelayedStreamingUsageRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeLiveDelayedStreamingUsageRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveDelayedStreamingUsageRequest
	GetStreamName() *string
}

type DescribeLiveDelayedStreamingUsageRequest struct {
	// The streaming domain name to query.
	//
	// - You can specify a single domain name or multiple domain names. Separate multiple domain names with commas (,).
	//
	// - If this parameter is left empty, the aggregated data of all live streaming domain names is returned by default.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The time span cannot exceed 10 hours. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-10-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the queried data. Unit: seconds. Valid values:
	//
	// - 300
	//
	// - 3600
	//
	// - 86400
	//
	// If this parameter is left empty or set to an unsupported value, the default value 3600 is used.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The live center to query. You can specify multiple regions. Separate multiple regions with commas (,). Valid values:
	//
	// - cn-beijing: Beijing
	//
	// - cn-shanghai: Shanghai
	//
	// - cn-shenzhen: Shenzhen
	//
	// - cn-qingdao: Qingdao
	//
	// - ap-southeast-1: Singapore
	//
	// - eu-central-1: Germany
	//
	// - ap-northeast-1: Tokyo
	//
	// - ap-southeast-5: Jakarta
	//
	// If this parameter is left empty, the aggregated data of all regions is returned by default.
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
	// The grouping key. If this parameter is left empty, user data is aggregated. Valid values:
	//
	// - domain: domain name. If the SplitBy (grouping key) parameter is set to domain, the Domain response parameter takes effect.
	//
	// - region: live center region. If the SplitBy (grouping key) parameter is set to region, the Region response parameter takes effect.
	//
	// - stream: stream name. If the SplitBy (grouping key) parameter is set to stream, the stream response parameter takes effect.
	//
	// > You can query data only when the parameter corresponding to the grouping key is not empty. Otherwise, an error is returned. For example, when DomainName is empty, you cannot specify domain as the grouping key.
	//
	// example:
	//
	// domain
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC. By default, data of the last seven days is returned.
	//
	// example:
	//
	// 2022-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The stream name. Separate multiple stream names with commas (,). By default, the data of all stream names is aggregated.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveDelayedStreamingUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDelayedStreamingUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDelayedStreamingUsageRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDelayedStreamingUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDelayedStreamingUsageRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDelayedStreamingUsageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDelayedStreamingUsageRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDelayedStreamingUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDelayedStreamingUsageRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeLiveDelayedStreamingUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDelayedStreamingUsageRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveDelayedStreamingUsageRequest) SetDomainName(v string) *DescribeLiveDelayedStreamingUsageRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageRequest) SetEndTime(v string) *DescribeLiveDelayedStreamingUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageRequest) SetInterval(v string) *DescribeLiveDelayedStreamingUsageRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageRequest) SetOwnerId(v int64) *DescribeLiveDelayedStreamingUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageRequest) SetRegion(v string) *DescribeLiveDelayedStreamingUsageRequest {
	s.Region = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageRequest) SetRegionId(v string) *DescribeLiveDelayedStreamingUsageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageRequest) SetSplitBy(v string) *DescribeLiveDelayedStreamingUsageRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageRequest) SetStartTime(v string) *DescribeLiveDelayedStreamingUsageRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageRequest) SetStreamName(v string) *DescribeLiveDelayedStreamingUsageRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveDelayedStreamingUsageRequest) Validate() error {
	return dara.Validate(s)
}
