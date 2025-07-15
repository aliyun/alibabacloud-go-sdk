// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainOnlineUserNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainOnlineUserNumRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveDomainOnlineUserNumRequest
	GetOwnerId() *int64
	SetQueryTime(v string) *DescribeLiveDomainOnlineUserNumRequest
	GetQueryTime() *string
	SetRegionId(v string) *DescribeLiveDomainOnlineUserNumRequest
	GetRegionId() *string
}

type DescribeLiveDomainOnlineUserNumRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The point of time to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2018-12-27T13:09:21Z
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveDomainOnlineUserNumRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainOnlineUserNumRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainOnlineUserNumRequest) GetQueryTime() *string {
	return s.QueryTime
}

func (s *DescribeLiveDomainOnlineUserNumRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainOnlineUserNumRequest) SetDomainName(v string) *DescribeLiveDomainOnlineUserNumRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumRequest) SetOwnerId(v int64) *DescribeLiveDomainOnlineUserNumRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumRequest) SetQueryTime(v string) *DescribeLiveDomainOnlineUserNumRequest {
	s.QueryTime = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumRequest) SetRegionId(v string) *DescribeLiveDomainOnlineUserNumRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumRequest) Validate() error {
	return dara.Validate(s)
}
