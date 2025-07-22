// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardTableEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardEntryId(v string) *DescribeForwardTableEntriesRequest
	GetForwardEntryId() *string
	SetForwardTableId(v string) *DescribeForwardTableEntriesRequest
	GetForwardTableId() *string
	SetMaxResults(v int32) *DescribeForwardTableEntriesRequest
	GetMaxResults() *int32
	SetNatGatewayId(v string) *DescribeForwardTableEntriesRequest
	GetNatGatewayId() *string
	SetNextToken(v string) *DescribeForwardTableEntriesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeForwardTableEntriesRequest
	GetRegionId() *string
}

type DescribeForwardTableEntriesRequest struct {
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NatGatewayId   *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeForwardTableEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DescribeForwardTableEntriesRequest) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *DescribeForwardTableEntriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeForwardTableEntriesRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeForwardTableEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeForwardTableEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeForwardTableEntriesRequest) SetForwardEntryId(v string) *DescribeForwardTableEntriesRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetForwardTableId(v string) *DescribeForwardTableEntriesRequest {
	s.ForwardTableId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetMaxResults(v int32) *DescribeForwardTableEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetNatGatewayId(v string) *DescribeForwardTableEntriesRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetNextToken(v string) *DescribeForwardTableEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetRegionId(v string) *DescribeForwardTableEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) Validate() error {
	return dara.Validate(s)
}
