// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeConsumerGroupRequest
	GetAccountId() *string
	SetOwnerId(v string) *DescribeConsumerGroupRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeConsumerGroupRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeConsumerGroupRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeConsumerGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeConsumerGroupRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *DescribeConsumerGroupRequest
	GetSubscriptionInstanceId() *string
}

type DescribeConsumerGroupRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than **0*	- and does not exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the change tracking instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aekz4us4iruleja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the change tracking instance. You can call the DescribeSubscriptionInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtso5xx5t9u19e****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s DescribeConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeConsumerGroupRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeConsumerGroupRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeConsumerGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConsumerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeConsumerGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeConsumerGroupRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *DescribeConsumerGroupRequest) SetAccountId(v string) *DescribeConsumerGroupRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetOwnerId(v string) *DescribeConsumerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetPageNum(v int32) *DescribeConsumerGroupRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetPageSize(v int32) *DescribeConsumerGroupRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetRegionId(v string) *DescribeConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetResourceGroupId(v string) *DescribeConsumerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetSubscriptionInstanceId(v string) *DescribeConsumerGroupRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *DescribeConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}
