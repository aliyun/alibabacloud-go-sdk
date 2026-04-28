// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupId(v string) *DescribeConsumersRequest
	GetConsumerGroupId() *string
	SetConsumerId(v string) *DescribeConsumersRequest
	GetConsumerId() *string
	SetGwClusterId(v string) *DescribeConsumersRequest
	GetGwClusterId() *string
	SetPageNumber(v int32) *DescribeConsumersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeConsumersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeConsumersRequest
	GetRegionId() *string
}

type DescribeConsumersRequest struct {
	// example:
	//
	// cg-xxxxxxxx
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// example:
	//
	// c-mqveroemc***
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
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

func (s DescribeConsumersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumersRequest) GoString() string {
	return s.String()
}

func (s *DescribeConsumersRequest) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *DescribeConsumersRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *DescribeConsumersRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeConsumersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeConsumersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConsumersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeConsumersRequest) SetConsumerGroupId(v string) *DescribeConsumersRequest {
	s.ConsumerGroupId = &v
	return s
}

func (s *DescribeConsumersRequest) SetConsumerId(v string) *DescribeConsumersRequest {
	s.ConsumerId = &v
	return s
}

func (s *DescribeConsumersRequest) SetGwClusterId(v string) *DescribeConsumersRequest {
	s.GwClusterId = &v
	return s
}

func (s *DescribeConsumersRequest) SetPageNumber(v int32) *DescribeConsumersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumersRequest) SetPageSize(v int32) *DescribeConsumersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConsumersRequest) SetRegionId(v string) *DescribeConsumersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeConsumersRequest) Validate() error {
	return dara.Validate(s)
}
