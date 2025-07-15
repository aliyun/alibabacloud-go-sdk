// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePushProxyUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLivePushProxyUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLivePushProxyUsageDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLivePushProxyUsageDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeLivePushProxyUsageDataRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeLivePushProxyUsageDataRequest
	GetRegionId() *string
	SetSplitBy(v string) *DescribeLivePushProxyUsageDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeLivePushProxyUsageDataRequest
	GetStartTime() *string
}

type DescribeLivePushProxyUsageDataRequest struct {
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
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The end time must be later than the start time.
	//
	// example:
	//
	// 2022-10-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// If you do not specify this parameter, data of all regions is aggregated and returned by default.
	//
	// example:
	//
	// cn-beijing
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The key that is used to group data. If you do not specify this parameter, the default value region is used. Data is aggregated and returned. Separate multiple keys with commas (,). Valid values:
	//
	// 	- domain: The value of DomainName in the response takes effect only if SplitBy is set to domain.
	//
	// 	- region: This is the default value. The value of Region in the response takes effect only if SplitBy is set to region.
	//
	// example:
	//
	// region
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. By default, data in the last seven days is returned.
	//
	// example:
	//
	// 2022-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLivePushProxyUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePushProxyUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLivePushProxyUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLivePushProxyUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLivePushProxyUsageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLivePushProxyUsageDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeLivePushProxyUsageDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLivePushProxyUsageDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeLivePushProxyUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLivePushProxyUsageDataRequest) SetDomainName(v string) *DescribeLivePushProxyUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataRequest) SetEndTime(v string) *DescribeLivePushProxyUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataRequest) SetOwnerId(v int64) *DescribeLivePushProxyUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataRequest) SetRegion(v string) *DescribeLivePushProxyUsageDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataRequest) SetRegionId(v string) *DescribeLivePushProxyUsageDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataRequest) SetSplitBy(v string) *DescribeLivePushProxyUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataRequest) SetStartTime(v string) *DescribeLivePushProxyUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
