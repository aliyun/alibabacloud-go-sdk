// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserVipsByDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailable(v string) *DescribeVodUserVipsByDomainRequest
	GetAvailable() *string
	SetDomainName(v string) *DescribeVodUserVipsByDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVodUserVipsByDomainRequest
	GetOwnerId() *int64
}

type DescribeVodUserVipsByDomainRequest struct {
	// example:
	//
	// on
	Available *string `json:"Available,omitempty" xml:"Available,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodUserVipsByDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserVipsByDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodUserVipsByDomainRequest) GetAvailable() *string {
	return s.Available
}

func (s *DescribeVodUserVipsByDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodUserVipsByDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodUserVipsByDomainRequest) SetAvailable(v string) *DescribeVodUserVipsByDomainRequest {
	s.Available = &v
	return s
}

func (s *DescribeVodUserVipsByDomainRequest) SetDomainName(v string) *DescribeVodUserVipsByDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodUserVipsByDomainRequest) SetOwnerId(v int64) *DescribeVodUserVipsByDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodUserVipsByDomainRequest) Validate() error {
	return dara.Validate(s)
}
