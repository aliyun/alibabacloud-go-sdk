// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamDelayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamDelayConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveStreamDelayConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamDelayConfigRequest
	GetRegionId() *string
}

type DescribeLiveStreamDelayConfigRequest struct {
	// The streaming domain.
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

func (s DescribeLiveStreamDelayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamDelayConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamDelayConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamDelayConfigRequest) SetDomainName(v string) *DescribeLiveStreamDelayConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigRequest) SetOwnerId(v int64) *DescribeLiveStreamDelayConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigRequest) SetRegionId(v string) *DescribeLiveStreamDelayConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigRequest) Validate() error {
	return dara.Validate(s)
}
