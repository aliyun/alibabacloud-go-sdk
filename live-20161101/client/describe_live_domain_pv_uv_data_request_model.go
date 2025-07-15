// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPvUvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainPvUvDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainPvUvDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveDomainPvUvDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainPvUvDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainPvUvDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainPvUvDataRequest struct {
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ssZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2018-03-20T16:00:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ssZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2018-03-17T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainPvUvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPvUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPvUvDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainPvUvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainPvUvDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainPvUvDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainPvUvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainPvUvDataRequest) SetDomainName(v string) *DescribeLiveDomainPvUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataRequest) SetEndTime(v string) *DescribeLiveDomainPvUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataRequest) SetOwnerId(v int64) *DescribeLiveDomainPvUvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataRequest) SetRegionId(v string) *DescribeLiveDomainPvUvDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataRequest) SetStartTime(v string) *DescribeLiveDomainPvUvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataRequest) Validate() error {
	return dara.Validate(s)
}
