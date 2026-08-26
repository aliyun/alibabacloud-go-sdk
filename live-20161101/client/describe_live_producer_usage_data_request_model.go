// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveProducerUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveProducerUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveProducerUsageDataRequest
	GetEndTime() *string
	SetInstance(v string) *DescribeLiveProducerUsageDataRequest
	GetInstance() *string
	SetInterval(v string) *DescribeLiveProducerUsageDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeLiveProducerUsageDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeLiveProducerUsageDataRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeLiveProducerUsageDataRequest
	GetRegionId() *string
	SetSplitBy(v string) *DescribeLiveProducerUsageDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeLiveProducerUsageDataRequest
	GetStartTime() *string
	SetType(v string) *DescribeLiveProducerUsageDataRequest
	GetType() *string
	SetApp(v string) *DescribeLiveProducerUsageDataRequest
	GetApp() *string
}

type DescribeLiveProducerUsageDataRequest struct {
	// The streaming domain name of the cloud producer studio.
	//
	// - Supports single or batch domain name queries. Separate multiple domain names with commas (,) for batch queries.
	//
	// - If this parameter is left empty, merged data of all live streaming domain names is returned by default.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2018-10-31T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance to query. Batch query is supported. Separate multiple instances with commas (,).
	//
	// > If this parameter is left empty, merged data of all instances is returned by default.
	//
	// example:
	//
	// a17d0184-462d-4630-b2a6-8c26dde2****
	Instance *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
	// The time granularity of the queried data. Valid values: 3600 (1 hour) and 86400 (1 day). Unit: seconds.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region to which the domain name belongs. If this parameter is left empty, merged data of all regions is returned by default. Batch query is supported. Separate multiple regions with commas (,).
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
	// The grouping key. You can specify one or more of the following: domain, region, instance, or type. Separate multiple values with commas (,). The specified fields will be grouped in the output.
	//
	//
	// > If this parameter is left empty, only aggregated data is returned.
	//
	// example:
	//
	// type
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The producer type. Batch query is supported. Separate multiple types with commas (,). Valid values:
	//
	// - **slidelive**: playlist-based.
	//
	// - **universal**: general-purpose.
	//
	// > If this parameter is left empty, merged data of all producer types is returned by default.
	//
	// example:
	//
	// slidelive
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the application to which the stream belongs.
	//
	// example:
	//
	// liveApp****
	App *string `json:"app,omitempty" xml:"app,omitempty"`
}

func (s DescribeLiveProducerUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveProducerUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveProducerUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveProducerUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveProducerUsageDataRequest) GetInstance() *string {
	return s.Instance
}

func (s *DescribeLiveProducerUsageDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveProducerUsageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveProducerUsageDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveProducerUsageDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveProducerUsageDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeLiveProducerUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveProducerUsageDataRequest) GetType() *string {
	return s.Type
}

func (s *DescribeLiveProducerUsageDataRequest) GetApp() *string {
	return s.App
}

func (s *DescribeLiveProducerUsageDataRequest) SetDomainName(v string) *DescribeLiveProducerUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) SetEndTime(v string) *DescribeLiveProducerUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) SetInstance(v string) *DescribeLiveProducerUsageDataRequest {
	s.Instance = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) SetInterval(v string) *DescribeLiveProducerUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) SetOwnerId(v int64) *DescribeLiveProducerUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) SetRegion(v string) *DescribeLiveProducerUsageDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) SetRegionId(v string) *DescribeLiveProducerUsageDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) SetSplitBy(v string) *DescribeLiveProducerUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) SetStartTime(v string) *DescribeLiveProducerUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) SetType(v string) *DescribeLiveProducerUsageDataRequest {
	s.Type = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) SetApp(v string) *DescribeLiveProducerUsageDataRequest {
	s.App = &v
	return s
}

func (s *DescribeLiveProducerUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
