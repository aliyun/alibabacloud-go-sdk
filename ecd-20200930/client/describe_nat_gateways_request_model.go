// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeNatGatewaysRequest
	GetMaxResults() *int32
	SetName(v string) *DescribeNatGatewaysRequest
	GetName() *string
	SetNatGatewayId(v string) *DescribeNatGatewaysRequest
	GetNatGatewayId() *string
	SetNextToken(v string) *DescribeNatGatewaysRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeNatGatewaysRequest
	GetRegionId() *string
	SetVpcId(v string) *DescribeNatGatewaysRequest
	GetVpcId() *string
}

type DescribeNatGatewaysRequest struct {
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId    *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNatGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNatGatewaysRequest) GetName() *string {
	return s.Name
}

func (s *DescribeNatGatewaysRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewaysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNatGatewaysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNatGatewaysRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNatGatewaysRequest) SetMaxResults(v int32) *DescribeNatGatewaysRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetName(v string) *DescribeNatGatewaysRequest {
	s.Name = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetNatGatewayId(v string) *DescribeNatGatewaysRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetNextToken(v string) *DescribeNatGatewaysRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetRegionId(v string) *DescribeNatGatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetVpcId(v string) *DescribeNatGatewaysRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) Validate() error {
	return dara.Validate(s)
}
