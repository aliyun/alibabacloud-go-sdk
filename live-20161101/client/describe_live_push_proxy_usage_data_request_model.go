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
	// The ingest domain name of the streamer to query.
	//
	// - You can specify a single domain name or multiple domain names separated by commas (,).
	//
	// - If this parameter is left empty, the aggregated data of all live streaming domain names is returned by default.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-10-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The live center to query. You can specify multiple regions separated by commas (,). Valid values:
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
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The grouping key. If this parameter is left empty, the default value is region, and the aggregated data is returned. You can specify multiple values separated by commas (,). Valid values:
	//
	// - domain: the domain name. If SplitBy is set to domain, the Domain field in the response takes effect.
	//
	// - region (default): the live center region. If SplitBy is set to region, the Region field in the response takes effect.
	//
	// example:
	//
	// region
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC. By default, data from the last seven days is returned.
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
