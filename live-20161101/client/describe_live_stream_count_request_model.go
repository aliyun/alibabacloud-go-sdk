// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamCountRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveStreamCountRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamCountRequest
	GetRegionId() *string
}

type DescribeLiveStreamCountRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveStreamCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamCountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamCountRequest) SetDomainName(v string) *DescribeLiveStreamCountRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamCountRequest) SetOwnerId(v int64) *DescribeLiveStreamCountRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamCountRequest) SetRegionId(v string) *DescribeLiveStreamCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamCountRequest) Validate() error {
	return dara.Validate(s)
}
