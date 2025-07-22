// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEipId(v string) *DescribeIpAddressesRequest
	GetEipId() *string
	SetInstanceId(v []*string) *DescribeIpAddressesRequest
	GetInstanceId() []*string
	SetMaxResults(v int32) *DescribeIpAddressesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeIpAddressesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeIpAddressesRequest
	GetRegionId() *string
}

type DescribeIpAddressesRequest struct {
	EipId      *string   `json:"EipId,omitempty" xml:"EipId,omitempty"`
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeIpAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpAddressesRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpAddressesRequest) GetEipId() *string {
	return s.EipId
}

func (s *DescribeIpAddressesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeIpAddressesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeIpAddressesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeIpAddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpAddressesRequest) SetEipId(v string) *DescribeIpAddressesRequest {
	s.EipId = &v
	return s
}

func (s *DescribeIpAddressesRequest) SetInstanceId(v []*string) *DescribeIpAddressesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeIpAddressesRequest) SetMaxResults(v int32) *DescribeIpAddressesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeIpAddressesRequest) SetNextToken(v string) *DescribeIpAddressesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeIpAddressesRequest) SetRegionId(v string) *DescribeIpAddressesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIpAddressesRequest) Validate() error {
	return dara.Validate(s)
}
