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
	// > - Make sure that the specified domain name is a streaming domain and that you have the required permissions to operate on it.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. The time must be later than the start time. The format is *yyyy-MM-dd*T*HH:mm:ssZ	- (UTC).
	//
	// example:
	//
	// 2018-03-20T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. The format is *yyyy-MM-dd*T*HH:mm:ssZ	- (UTC).
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
