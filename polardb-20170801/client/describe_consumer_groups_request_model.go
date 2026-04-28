// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupId(v string) *DescribeConsumerGroupsRequest
	GetConsumerGroupId() *string
	SetGwClusterId(v string) *DescribeConsumerGroupsRequest
	GetGwClusterId() *string
	SetPageNumber(v int32) *DescribeConsumerGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeConsumerGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeConsumerGroupsRequest
	GetRegionId() *string
}

type DescribeConsumerGroupsRequest struct {
	// example:
	//
	// cg-xxxxxxxx
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeConsumerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupsRequest) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *DescribeConsumerGroupsRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeConsumerGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeConsumerGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConsumerGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeConsumerGroupsRequest) SetConsumerGroupId(v string) *DescribeConsumerGroupsRequest {
	s.ConsumerGroupId = &v
	return s
}

func (s *DescribeConsumerGroupsRequest) SetGwClusterId(v string) *DescribeConsumerGroupsRequest {
	s.GwClusterId = &v
	return s
}

func (s *DescribeConsumerGroupsRequest) SetPageNumber(v int32) *DescribeConsumerGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumerGroupsRequest) SetPageSize(v int32) *DescribeConsumerGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConsumerGroupsRequest) SetRegionId(v string) *DescribeConsumerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeConsumerGroupsRequest) Validate() error {
	return dara.Validate(s)
}
