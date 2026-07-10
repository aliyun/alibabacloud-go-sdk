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
	SetConsumerName(v string) *DescribeConsumersRequest
	GetConsumerName() *string
	SetConsumerNameList(v string) *DescribeConsumersRequest
	GetConsumerNameList() *string
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
	// The user group ID.
	//
	// example:
	//
	// cg-xxxxxxxx
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// c-mqveroemc***
	ConsumerId       *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	ConsumerName     *string `json:"ConsumerName,omitempty" xml:"ConsumerName,omitempty"`
	ConsumerNameList *string `json:"ConsumerNameList,omitempty" xml:"ConsumerNameList,omitempty"`
	// The gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
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

func (s *DescribeConsumersRequest) GetConsumerName() *string {
	return s.ConsumerName
}

func (s *DescribeConsumersRequest) GetConsumerNameList() *string {
	return s.ConsumerNameList
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

func (s *DescribeConsumersRequest) SetConsumerName(v string) *DescribeConsumersRequest {
	s.ConsumerName = &v
	return s
}

func (s *DescribeConsumersRequest) SetConsumerNameList(v string) *DescribeConsumersRequest {
	s.ConsumerNameList = &v
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
