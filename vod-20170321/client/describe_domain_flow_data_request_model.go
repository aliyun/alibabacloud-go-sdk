// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainFlowDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainFlowDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainFlowDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainFlowDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDomainFlowDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainFlowDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v string) *DescribeDomainFlowDataRequest
	GetOwnerId() *string
	SetResourceOwnerId(v string) *DescribeDomainFlowDataRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v string) *DescribeDomainFlowDataRequest
	GetResourceRealOwnerId() *string
	SetStartTime(v string) *DescribeDomainFlowDataRequest
	GetStartTime() *string
	SetTimeMerge(v string) *DescribeDomainFlowDataRequest
	GetTimeMerge() *string
}

type DescribeDomainFlowDataRequest struct {
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

func (s DescribeDomainFlowDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainFlowDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainFlowDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainFlowDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainFlowDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainFlowDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainFlowDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainFlowDataRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeDomainFlowDataRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DescribeDomainFlowDataRequest) GetResourceRealOwnerId() *string {
	return s.ResourceRealOwnerId
}

func (s *DescribeDomainFlowDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainFlowDataRequest) GetTimeMerge() *string {
	return s.TimeMerge
}

func (s *DescribeDomainFlowDataRequest) SetDomainName(v string) *DescribeDomainFlowDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetEndTime(v string) *DescribeDomainFlowDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetInterval(v string) *DescribeDomainFlowDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetIspNameEn(v string) *DescribeDomainFlowDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetLocationNameEn(v string) *DescribeDomainFlowDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetOwnerId(v string) *DescribeDomainFlowDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetResourceOwnerId(v string) *DescribeDomainFlowDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetResourceRealOwnerId(v string) *DescribeDomainFlowDataRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetStartTime(v string) *DescribeDomainFlowDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetTimeMerge(v string) *DescribeDomainFlowDataRequest {
	s.TimeMerge = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) Validate() error {
	return dara.Validate(s)
}
