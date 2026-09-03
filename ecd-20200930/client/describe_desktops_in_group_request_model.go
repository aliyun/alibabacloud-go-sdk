// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopsInGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomEndTimePeriod(v int64) *DescribeDesktopsInGroupRequest
	GetCustomEndTimePeriod() *int64
	SetCustomStartTimePeriod(v int64) *DescribeDesktopsInGroupRequest
	GetCustomStartTimePeriod() *int64
	SetDesktopGroupId(v string) *DescribeDesktopsInGroupRequest
	GetDesktopGroupId() *string
	SetIgnoreDeleted(v bool) *DescribeDesktopsInGroupRequest
	GetIgnoreDeleted() *bool
	SetMaxResults(v int32) *DescribeDesktopsInGroupRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDesktopsInGroupRequest
	GetNextToken() *string
	SetPayType(v string) *DescribeDesktopsInGroupRequest
	GetPayType() *string
	SetRegionId(v string) *DescribeDesktopsInGroupRequest
	GetRegionId() *string
}

type DescribeDesktopsInGroupRequest struct {
	// The end time of the custom time range.
	//
	// example:
	//
	// 18
	CustomEndTimePeriod *int64 `json:"CustomEndTimePeriod,omitempty" xml:"CustomEndTimePeriod,omitempty"`
	// The start time of the custom time range.
	//
	// example:
	//
	// 8
	CustomStartTimePeriod *int64 `json:"CustomStartTimePeriod,omitempty" xml:"CustomStartTimePeriod,omitempty"`
	// The ID of the shared cloud desktop.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// Specifies whether to ignore the deletion flag.
	//
	// example:
	//
	// true
	IgnoreDeleted *bool `json:"IgnoreDeleted,omitempty" xml:"IgnoreDeleted,omitempty"`
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. An empty value indicates that no more results exist.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDesktopsInGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsInGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupRequest) GetCustomEndTimePeriod() *int64 {
	return s.CustomEndTimePeriod
}

func (s *DescribeDesktopsInGroupRequest) GetCustomStartTimePeriod() *int64 {
	return s.CustomStartTimePeriod
}

func (s *DescribeDesktopsInGroupRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeDesktopsInGroupRequest) GetIgnoreDeleted() *bool {
	return s.IgnoreDeleted
}

func (s *DescribeDesktopsInGroupRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDesktopsInGroupRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopsInGroupRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDesktopsInGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDesktopsInGroupRequest) SetCustomEndTimePeriod(v int64) *DescribeDesktopsInGroupRequest {
	s.CustomEndTimePeriod = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetCustomStartTimePeriod(v int64) *DescribeDesktopsInGroupRequest {
	s.CustomStartTimePeriod = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetDesktopGroupId(v string) *DescribeDesktopsInGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetIgnoreDeleted(v bool) *DescribeDesktopsInGroupRequest {
	s.IgnoreDeleted = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetMaxResults(v int32) *DescribeDesktopsInGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetNextToken(v string) *DescribeDesktopsInGroupRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetPayType(v string) *DescribeDesktopsInGroupRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetRegionId(v string) *DescribeDesktopsInGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) Validate() error {
	return dara.Validate(s)
}
