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
	CustomEndTimePeriod   *int64 `json:"CustomEndTimePeriod,omitempty" xml:"CustomEndTimePeriod,omitempty"`
	CustomStartTimePeriod *int64 `json:"CustomStartTimePeriod,omitempty" xml:"CustomStartTimePeriod,omitempty"`
	// The ID of the cloud computer share.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// Specifies whether to ignore deletion flags.
	//
	// Default value: true. Valid values:
	//
	// 	- true: ignores deletion flags. The cloud computers that were deleted are returned.
	//
	// 	- false: does not ignore deletion flags. The cloud computers that were deleted are not returned.
	//
	// example:
	//
	// true
	IgnoreDeleted *bool `json:"IgnoreDeleted,omitempty" xml:"IgnoreDeleted,omitempty"`
	// The maximum number of entries per page.
	//
	// 	- Default value: 10.
	//
	// 	- Maximum value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The billing method of the cloud computer share.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
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
