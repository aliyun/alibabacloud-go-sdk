// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainBpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDomainBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainBpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v string) *DescribeDomainBpsDataRequest
	GetOwnerId() *string
	SetResourceOwnerId(v string) *DescribeDomainBpsDataRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v string) *DescribeDomainBpsDataRequest
	GetResourceRealOwnerId() *string
	SetStartTime(v string) *DescribeDomainBpsDataRequest
	GetStartTime() *string
	SetTimeMerge(v string) *DescribeDomainBpsDataRequest
	GetTimeMerge() *string
}

type DescribeDomainBpsDataRequest struct {
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval            *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn           *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn      *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId             *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerId     *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId *string `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeMerge           *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDomainBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainBpsDataRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeDomainBpsDataRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DescribeDomainBpsDataRequest) GetResourceRealOwnerId() *string {
	return s.ResourceRealOwnerId
}

func (s *DescribeDomainBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainBpsDataRequest) GetTimeMerge() *string {
	return s.TimeMerge
}

func (s *DescribeDomainBpsDataRequest) SetDomainName(v string) *DescribeDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetEndTime(v string) *DescribeDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetInterval(v string) *DescribeDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetIspNameEn(v string) *DescribeDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetOwnerId(v string) *DescribeDomainBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetResourceOwnerId(v string) *DescribeDomainBpsDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetResourceRealOwnerId(v string) *DescribeDomainBpsDataRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetStartTime(v string) *DescribeDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetTimeMerge(v string) *DescribeDomainBpsDataRequest {
	s.TimeMerge = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
