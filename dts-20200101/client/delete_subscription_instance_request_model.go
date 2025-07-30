// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscriptionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeleteSubscriptionInstanceRequest
	GetAccountId() *string
	SetOwnerId(v string) *DeleteSubscriptionInstanceRequest
	GetOwnerId() *string
	SetRegionId(v string) *DeleteSubscriptionInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteSubscriptionInstanceRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *DeleteSubscriptionInstanceRequest
	GetSubscriptionInstanceId() *string
}

type DeleteSubscriptionInstanceRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// dtsmxg11pfp231****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s DeleteSubscriptionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionInstanceRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteSubscriptionInstanceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DeleteSubscriptionInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSubscriptionInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteSubscriptionInstanceRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *DeleteSubscriptionInstanceRequest) SetAccountId(v string) *DeleteSubscriptionInstanceRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteSubscriptionInstanceRequest) SetOwnerId(v string) *DeleteSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSubscriptionInstanceRequest) SetRegionId(v string) *DeleteSubscriptionInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSubscriptionInstanceRequest) SetResourceGroupId(v string) *DeleteSubscriptionInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteSubscriptionInstanceRequest) SetSubscriptionInstanceId(v string) *DeleteSubscriptionInstanceRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *DeleteSubscriptionInstanceRequest) Validate() error {
	return dara.Validate(s)
}
