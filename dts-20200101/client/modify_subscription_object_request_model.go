// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ModifySubscriptionObjectRequest
	GetAccountId() *string
	SetOwnerId(v string) *ModifySubscriptionObjectRequest
	GetOwnerId() *string
	SetRegionId(v string) *ModifySubscriptionObjectRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifySubscriptionObjectRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *ModifySubscriptionObjectRequest
	GetSubscriptionInstanceId() *string
	SetSubscriptionObject(v string) *ModifySubscriptionObjectRequest
	GetSubscriptionObject() *string
}

type ModifySubscriptionObjectRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// dtsl8zl9ek6292****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	// The objects from which you want to track data changes. The value is a JSON string and can contain regular expressions. For more information, see [SubscriptionObjects](https://help.aliyun.com/document_detail/141902.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"DBName":"dtstestdata"}]
	SubscriptionObject *string `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty"`
}

func (s ModifySubscriptionObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionObjectRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionObjectRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ModifySubscriptionObjectRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifySubscriptionObjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySubscriptionObjectRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifySubscriptionObjectRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *ModifySubscriptionObjectRequest) GetSubscriptionObject() *string {
	return s.SubscriptionObject
}

func (s *ModifySubscriptionObjectRequest) SetAccountId(v string) *ModifySubscriptionObjectRequest {
	s.AccountId = &v
	return s
}

func (s *ModifySubscriptionObjectRequest) SetOwnerId(v string) *ModifySubscriptionObjectRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySubscriptionObjectRequest) SetRegionId(v string) *ModifySubscriptionObjectRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySubscriptionObjectRequest) SetResourceGroupId(v string) *ModifySubscriptionObjectRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySubscriptionObjectRequest) SetSubscriptionInstanceId(v string) *ModifySubscriptionObjectRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *ModifySubscriptionObjectRequest) SetSubscriptionObject(v string) *ModifySubscriptionObjectRequest {
	s.SubscriptionObject = &v
	return s
}

func (s *ModifySubscriptionObjectRequest) Validate() error {
	return dara.Validate(s)
}
