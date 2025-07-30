// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSubscriptionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *StartSubscriptionInstanceRequest
	GetAccountId() *string
	SetOwnerId(v string) *StartSubscriptionInstanceRequest
	GetOwnerId() *string
	SetRegionId(v string) *StartSubscriptionInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StartSubscriptionInstanceRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *StartSubscriptionInstanceRequest
	GetSubscriptionInstanceId() *string
}

type StartSubscriptionInstanceRequest struct {
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
	// rg-aekz4us4iruleja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the change tracking instance. You can call the DescribeSubscriptionInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtso6m11cxt26q****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s StartSubscriptionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartSubscriptionInstanceRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *StartSubscriptionInstanceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *StartSubscriptionInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartSubscriptionInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StartSubscriptionInstanceRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *StartSubscriptionInstanceRequest) SetAccountId(v string) *StartSubscriptionInstanceRequest {
	s.AccountId = &v
	return s
}

func (s *StartSubscriptionInstanceRequest) SetOwnerId(v string) *StartSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *StartSubscriptionInstanceRequest) SetRegionId(v string) *StartSubscriptionInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartSubscriptionInstanceRequest) SetResourceGroupId(v string) *StartSubscriptionInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StartSubscriptionInstanceRequest) SetSubscriptionInstanceId(v string) *StartSubscriptionInstanceRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *StartSubscriptionInstanceRequest) Validate() error {
	return dara.Validate(s)
}
