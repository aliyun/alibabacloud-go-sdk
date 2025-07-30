// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInstanceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeSubscriptionInstanceStatusRequest
	GetAccountId() *string
	SetOwnerId(v string) *DescribeSubscriptionInstanceStatusRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeSubscriptionInstanceStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSubscriptionInstanceStatusRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *DescribeSubscriptionInstanceStatusRequest
	GetSubscriptionInstanceId() *string
}

type DescribeSubscriptionInstanceStatusRequest struct {
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
	// The ID of the change tracking instance. You can call the [DescribeSubscriptionInstances](https://help.aliyun.com/document_detail/49442.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsy0zz3t13h7d****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeSubscriptionInstanceStatusRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeSubscriptionInstanceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubscriptionInstanceStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSubscriptionInstanceStatusRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetAccountId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetOwnerId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetRegionId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetResourceGroupId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetSubscriptionInstanceId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusRequest) Validate() error {
	return dara.Validate(s)
}
