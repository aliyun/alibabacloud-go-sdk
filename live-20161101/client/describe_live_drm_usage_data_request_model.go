// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDrmUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDrmUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDrmUsageDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveDrmUsageDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeLiveDrmUsageDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDrmUsageDataRequest
	GetRegionId() *string
	SetSplitBy(v string) *DescribeLiveDrmUsageDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeLiveDrmUsageDataRequest
	GetStartTime() *string
}

type DescribeLiveDrmUsageDataRequest struct {
	// The domain name.
	//
	// 	- You can query one or more domain names. If you specify multiple domain names, separate them with commas (,).
	//
	// 	- If you leave this parameter empty, the data of all domain names within your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-02T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- 300
	//
	// 	- 3600
	//
	// 	- 86400
	//
	// Default value: 3600.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The key that is used to group data. The following keys are supported: domain, region, and drm_type. If you want to specify multiple keys, separate them with commas (,). Default value: domain,region,drm_type. If you leave this parameter empty or set it to null, the returned data is not grouped.
	//
	// example:
	//
	// domain,region,drm_type
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The minimum time granularity is 5 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-01T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDrmUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDrmUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDrmUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDrmUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDrmUsageDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDrmUsageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDrmUsageDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDrmUsageDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeLiveDrmUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDrmUsageDataRequest) SetDomainName(v string) *DescribeLiveDrmUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDrmUsageDataRequest) SetEndTime(v string) *DescribeLiveDrmUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDrmUsageDataRequest) SetInterval(v string) *DescribeLiveDrmUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDrmUsageDataRequest) SetOwnerId(v int64) *DescribeLiveDrmUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDrmUsageDataRequest) SetRegionId(v string) *DescribeLiveDrmUsageDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDrmUsageDataRequest) SetSplitBy(v string) *DescribeLiveDrmUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeLiveDrmUsageDataRequest) SetStartTime(v string) *DescribeLiveDrmUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDrmUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
