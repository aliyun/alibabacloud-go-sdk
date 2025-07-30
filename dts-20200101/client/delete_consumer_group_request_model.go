// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeleteConsumerGroupRequest
	GetAccountId() *string
	SetConsumerGroupID(v string) *DeleteConsumerGroupRequest
	GetConsumerGroupID() *string
	SetOwnerId(v string) *DeleteConsumerGroupRequest
	GetOwnerId() *string
	SetRegionId(v string) *DeleteConsumerGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteConsumerGroupRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *DeleteConsumerGroupRequest
	GetSubscriptionInstanceId() *string
}

type DeleteConsumerGroupRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the consumer group. You can call the [DescribeConsumerGroup](https://help.aliyun.com/document_detail/122886.html) operation to query the consumer group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtssb911ydd192****
	ConsumerGroupID *string `json:"ConsumerGroupID,omitempty" xml:"ConsumerGroupID,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the change tracking instance. You can call the **DescribeSubscriptionInstances*	- operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsg2m10r1x15a****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s DeleteConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteConsumerGroupRequest) GetConsumerGroupID() *string {
	return s.ConsumerGroupID
}

func (s *DeleteConsumerGroupRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DeleteConsumerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteConsumerGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteConsumerGroupRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *DeleteConsumerGroupRequest) SetAccountId(v string) *DeleteConsumerGroupRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetConsumerGroupID(v string) *DeleteConsumerGroupRequest {
	s.ConsumerGroupID = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetOwnerId(v string) *DeleteConsumerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetRegionId(v string) *DeleteConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetResourceGroupId(v string) *DeleteConsumerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetSubscriptionInstanceId(v string) *DeleteConsumerGroupRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}
