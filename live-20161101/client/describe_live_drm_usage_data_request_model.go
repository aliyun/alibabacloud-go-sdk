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
	// The live streaming domain name.
	//
	// - You can specify a single domain name or multiple domain names. Separate multiple domain names with commas (,).
	//
	// - If this parameter is left empty, the merged data of all live streaming domain names is returned by default.
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
	// The time granularity of the queried data. Unit: seconds. Valid values:
	//
	// - 300
	//
	// - 3600
	//
	// - 86400
	//
	// If you do not set this parameter or set it to an unsupported value, the default value 3600 is used.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The grouping key. Default value: domain,region,drm_type. You can specify one or more of the following values: domain, region, and drm_type. Separate multiple values with commas (,). Set this parameter to an empty string or null to disable grouping by these keys.
	//
	// example:
	//
	// domain,region,drm_type
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC. The minimum data granularity is 5 minutes.
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
