// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeSubDomainRequest
	GetRegionId() *string
	SetZoneId(v string) *DescribeSubDomainRequest
	GetZoneId() *string
}

type DescribeSubDomainRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSubDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubDomainRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSubDomainRequest) SetRegionId(v string) *DescribeSubDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubDomainRequest) SetZoneId(v string) *DescribeSubDomainRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeSubDomainRequest) Validate() error {
	return dara.Validate(s)
}
