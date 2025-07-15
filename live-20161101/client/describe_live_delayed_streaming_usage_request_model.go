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
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The end time must be later than the start time. We recommend that you specify a time range that is less than or equal to 10 hours.
	//
	// example:
	//
	// 2022-10-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- 300
	//
	// 	- 3600
	//
	// 	- 86400
	//
	// If you specify an invalid value or do not specify this parameter, the default value 3600 is used.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. Separate multiple region IDs with commas (,). Valid values:
	//
	// 	- cn-beijing: China (Beijing)
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// 	- cn-shenzhen: China (Shenzhen)
	//
	// 	- cn-qingdao: China (Qingdao)
	//
	// 	- ap-southeast-1: Singapore
	//
	// 	- eu-central-1: Germany (Frankfurt)
	//
	// 	- ap-northeast-1: Japan (Tokyo)
	//
	// 	- ap-southeast-5: Indonesia (Jakarta)
	//
	// If you leave this parameter empty, data of all regions is aggregated and returned by default.
	//
	// example:
	//
	// cn-shanghai
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The key that is used to group data. If you leave this parameter empty, data is aggregated and returned. Valid values:
	//
	// 	- domain: The DomainName parameter in the response takes effect only if SplitBy is set to domain.
	//
	// 	- region: The Region parameter in the response takes effect only if SplitBy is set to region.
	//
	// 	- stream: The StreamName parameter in the response takes effect only if SplitBy is set to stream.
	//
	// >  This parameter takes effect only if the parameter corresponding to the value of this parameter is not left empty. Otherwise, an error is returned. For example, you cannot set this parameter to domain if the DomainName parameter is left empty.
	//
	// example:
	//
	// domain
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. By default, data in the last seven days is returned.
	//
	// example:
	//
	// 2022-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the stream. Separate multiple stream names with commas (,). By default, data of all streams is aggregated and returned.
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
