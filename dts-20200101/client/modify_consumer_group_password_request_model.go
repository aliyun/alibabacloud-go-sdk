// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerGroupPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ModifyConsumerGroupPasswordRequest
	GetAccountId() *string
	SetConsumerGroupID(v string) *ModifyConsumerGroupPasswordRequest
	GetConsumerGroupID() *string
	SetConsumerGroupName(v string) *ModifyConsumerGroupPasswordRequest
	GetConsumerGroupName() *string
	SetConsumerGroupPassword(v string) *ModifyConsumerGroupPasswordRequest
	GetConsumerGroupPassword() *string
	SetConsumerGroupUserName(v string) *ModifyConsumerGroupPasswordRequest
	GetConsumerGroupUserName() *string
	SetOwnerId(v string) *ModifyConsumerGroupPasswordRequest
	GetOwnerId() *string
	SetRegionId(v string) *ModifyConsumerGroupPasswordRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyConsumerGroupPasswordRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *ModifyConsumerGroupPasswordRequest
	GetSubscriptionInstanceId() *string
	SetConsumerGroupNewPassword(v string) *ModifyConsumerGroupPasswordRequest
	GetConsumerGroupNewPassword() *string
}

type ModifyConsumerGroupPasswordRequest struct {
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
	// dtswc411cg617p****
	ConsumerGroupID *string `json:"ConsumerGroupID,omitempty" xml:"ConsumerGroupID,omitempty"`
	// The name of the consumer group. The name cannot exceed 128 characters in length. We recommend that you use an informative name for easy identification.
	//
	// example:
	//
	// dtstest
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// The new password of the consumer group.
	//
	// 	- A password must contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- A password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test123456
	ConsumerGroupPassword *string `json:"ConsumerGroupPassword,omitempty" xml:"ConsumerGroupPassword,omitempty"`
	// The username of the consumer group. You can call the [DescribeConsumerGroup](https://help.aliyun.com/document_detail/122886.html) operation to query the username.
	//
	// example:
	//
	// dtstest
	ConsumerGroupUserName *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
	OwnerId               *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The new password of the consumer group.
	//
	//
	//
	// 	- A password must contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- A password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test654321
	ConsumerGroupNewPassword *string `json:"consumerGroupNewPassword,omitempty" xml:"consumerGroupNewPassword,omitempty"`
}

func (s ModifyConsumerGroupPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerGroupPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyConsumerGroupPasswordRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ModifyConsumerGroupPasswordRequest) GetConsumerGroupID() *string {
	return s.ConsumerGroupID
}

func (s *ModifyConsumerGroupPasswordRequest) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *ModifyConsumerGroupPasswordRequest) GetConsumerGroupPassword() *string {
	return s.ConsumerGroupPassword
}

func (s *ModifyConsumerGroupPasswordRequest) GetConsumerGroupUserName() *string {
	return s.ConsumerGroupUserName
}

func (s *ModifyConsumerGroupPasswordRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyConsumerGroupPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyConsumerGroupPasswordRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyConsumerGroupPasswordRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *ModifyConsumerGroupPasswordRequest) GetConsumerGroupNewPassword() *string {
	return s.ConsumerGroupNewPassword
}

func (s *ModifyConsumerGroupPasswordRequest) SetAccountId(v string) *ModifyConsumerGroupPasswordRequest {
	s.AccountId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetConsumerGroupID(v string) *ModifyConsumerGroupPasswordRequest {
	s.ConsumerGroupID = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetConsumerGroupName(v string) *ModifyConsumerGroupPasswordRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetConsumerGroupPassword(v string) *ModifyConsumerGroupPasswordRequest {
	s.ConsumerGroupPassword = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetConsumerGroupUserName(v string) *ModifyConsumerGroupPasswordRequest {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetOwnerId(v string) *ModifyConsumerGroupPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetRegionId(v string) *ModifyConsumerGroupPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetResourceGroupId(v string) *ModifyConsumerGroupPasswordRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetSubscriptionInstanceId(v string) *ModifyConsumerGroupPasswordRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetConsumerGroupNewPassword(v string) *ModifyConsumerGroupPasswordRequest {
	s.ConsumerGroupNewPassword = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) Validate() error {
	return dara.Validate(s)
}
