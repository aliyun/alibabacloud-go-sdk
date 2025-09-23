// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsWithCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainQpsWithCacheRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainQpsWithCacheRequest
	GetEndTime() *int64
	SetResourceGroupId(v string) *DescribeDomainQpsWithCacheRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeDomainQpsWithCacheRequest
	GetSourceIp() *string
	SetStartTime(v int64) *DescribeDomainQpsWithCacheRequest
	GetStartTime() *int64
}

type DescribeDomainQpsWithCacheRequest struct {
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1577796336
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1577794536
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainQpsWithCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsWithCacheRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainQpsWithCacheRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainQpsWithCacheRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainQpsWithCacheRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainQpsWithCacheRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainQpsWithCacheRequest) SetDomain(v string) *DescribeDomainQpsWithCacheRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetEndTime(v int64) *DescribeDomainQpsWithCacheRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetResourceGroupId(v string) *DescribeDomainQpsWithCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetSourceIp(v string) *DescribeDomainQpsWithCacheRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetStartTime(v int64) *DescribeDomainQpsWithCacheRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) Validate() error {
	return dara.Validate(s)
}
