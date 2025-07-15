// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePullStreamConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLivePullStreamConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLivePullStreamConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLivePullStreamConfigRequest
	GetRegionId() *string
}

type DescribeLivePullStreamConfigRequest struct {
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

func (s DescribeLivePullStreamConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullStreamConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLivePullStreamConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLivePullStreamConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLivePullStreamConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLivePullStreamConfigRequest) SetDomainName(v string) *DescribeLivePullStreamConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePullStreamConfigRequest) SetOwnerId(v int64) *DescribeLivePullStreamConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLivePullStreamConfigRequest) SetRegionId(v string) *DescribeLivePullStreamConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLivePullStreamConfigRequest) Validate() error {
	return dara.Validate(s)
}
