// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeZonesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeZonesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeZonesRequest
	GetRegionId() *string
}

type DescribeZonesRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 40831b4f-d91d-4796-9589-ad306ec528d5
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeZonesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesRequest) SetMaxResults(v int32) *DescribeZonesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeZonesRequest) SetNextToken(v string) *DescribeZonesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
