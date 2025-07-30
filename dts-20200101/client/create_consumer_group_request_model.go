// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateConsumerGroupRequest
	GetAccountId() *string
	SetConsumerGroupName(v string) *CreateConsumerGroupRequest
	GetConsumerGroupName() *string
	SetConsumerGroupPassword(v string) *CreateConsumerGroupRequest
	GetConsumerGroupPassword() *string
	SetConsumerGroupUserName(v string) *CreateConsumerGroupRequest
	GetConsumerGroupUserName() *string
	SetOwnerId(v string) *CreateConsumerGroupRequest
	GetOwnerId() *string
	SetRegionId(v string) *CreateConsumerGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateConsumerGroupRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *CreateConsumerGroupRequest
	GetSubscriptionInstanceId() *string
}

type CreateConsumerGroupRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the consumer group. The name cannot exceed 128 characters in length. We recommend that you use an informative name for easy identification.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试订阅组
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// The password that corresponds to the username of the consumer group.
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
	// The username of the consumer group.
	//
	// 	- A username must contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// 	- A username cannot exceed 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtstest
	ConsumerGroupUserName *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
	OwnerId               *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The ID of the change tracking instance. You can call the DescribeSubscriptionInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsg2m10r1x15a****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s CreateConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateConsumerGroupRequest) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *CreateConsumerGroupRequest) GetConsumerGroupPassword() *string {
	return s.ConsumerGroupPassword
}

func (s *CreateConsumerGroupRequest) GetConsumerGroupUserName() *string {
	return s.ConsumerGroupUserName
}

func (s *CreateConsumerGroupRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateConsumerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateConsumerGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateConsumerGroupRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *CreateConsumerGroupRequest) SetAccountId(v string) *CreateConsumerGroupRequest {
	s.AccountId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetConsumerGroupName(v string) *CreateConsumerGroupRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetConsumerGroupPassword(v string) *CreateConsumerGroupRequest {
	s.ConsumerGroupPassword = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetConsumerGroupUserName(v string) *CreateConsumerGroupRequest {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetOwnerId(v string) *CreateConsumerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRegionId(v string) *CreateConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetResourceGroupId(v string) *CreateConsumerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetSubscriptionInstanceId(v string) *CreateConsumerGroupRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *CreateConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}
