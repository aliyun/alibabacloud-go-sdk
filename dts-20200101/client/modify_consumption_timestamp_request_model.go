// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumptionTimestampRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ModifyConsumptionTimestampRequest
	GetAccountId() *string
	SetConsumptionTimestamp(v string) *ModifyConsumptionTimestampRequest
	GetConsumptionTimestamp() *string
	SetOwnerId(v string) *ModifyConsumptionTimestampRequest
	GetOwnerId() *string
	SetRegionId(v string) *ModifyConsumptionTimestampRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyConsumptionTimestampRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *ModifyConsumptionTimestampRequest
	GetSubscriptionInstanceId() *string
}

type ModifyConsumptionTimestampRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The consumption checkpoint. The format is *yyyy-MM-dd*T*HH:mm:ss*Z. The time is displayed in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-10-15T17:20:03Z
	ConsumptionTimestamp *string `json:"ConsumptionTimestamp,omitempty" xml:"ConsumptionTimestamp,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s ModifyConsumptionTimestampRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumptionTimestampRequest) GoString() string {
	return s.String()
}

func (s *ModifyConsumptionTimestampRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ModifyConsumptionTimestampRequest) GetConsumptionTimestamp() *string {
	return s.ConsumptionTimestamp
}

func (s *ModifyConsumptionTimestampRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyConsumptionTimestampRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyConsumptionTimestampRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyConsumptionTimestampRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *ModifyConsumptionTimestampRequest) SetAccountId(v string) *ModifyConsumptionTimestampRequest {
	s.AccountId = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetConsumptionTimestamp(v string) *ModifyConsumptionTimestampRequest {
	s.ConsumptionTimestamp = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetOwnerId(v string) *ModifyConsumptionTimestampRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetRegionId(v string) *ModifyConsumptionTimestampRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetResourceGroupId(v string) *ModifyConsumptionTimestampRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetSubscriptionInstanceId(v string) *ModifyConsumptionTimestampRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) Validate() error {
	return dara.Validate(s)
}
