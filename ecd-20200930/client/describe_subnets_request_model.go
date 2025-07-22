// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubnetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeSubnetsRequest
	GetMaxResults() *int32
	SetName(v string) *DescribeSubnetsRequest
	GetName() *string
	SetNextToken(v string) *DescribeSubnetsRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeSubnetsRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DescribeSubnetsRequest
	GetRegionId() *string
	SetSubnetId(v string) *DescribeSubnetsRequest
	GetSubnetId() *string
}

type DescribeSubnetsRequest struct {
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s DescribeSubnetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubnetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubnetsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSubnetsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSubnetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSubnetsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeSubnetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubnetsRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *DescribeSubnetsRequest) SetMaxResults(v int32) *DescribeSubnetsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSubnetsRequest) SetName(v string) *DescribeSubnetsRequest {
	s.Name = &v
	return s
}

func (s *DescribeSubnetsRequest) SetNextToken(v string) *DescribeSubnetsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSubnetsRequest) SetOfficeSiteId(v string) *DescribeSubnetsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeSubnetsRequest) SetRegionId(v string) *DescribeSubnetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubnetsRequest) SetSubnetId(v string) *DescribeSubnetsRequest {
	s.SubnetId = &v
	return s
}

func (s *DescribeSubnetsRequest) Validate() error {
	return dara.Validate(s)
}
