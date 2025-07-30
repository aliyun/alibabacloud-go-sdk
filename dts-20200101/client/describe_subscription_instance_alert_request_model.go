// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInstanceAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeSubscriptionInstanceAlertRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeSubscriptionInstanceAlertRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeSubscriptionInstanceAlertRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeSubscriptionInstanceAlertRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSubscriptionInstanceAlertRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *DescribeSubscriptionInstanceAlertRequest
	GetSubscriptionInstanceId() *string
}

type DescribeSubscriptionInstanceAlertRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
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
	// dtsl8zl9ek6292****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s DescribeSubscriptionInstanceAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstanceAlertRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceAlertRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeSubscriptionInstanceAlertRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSubscriptionInstanceAlertRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeSubscriptionInstanceAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubscriptionInstanceAlertRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSubscriptionInstanceAlertRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetAccountId(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetClientToken(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetOwnerId(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetRegionId(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetResourceGroupId(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetSubscriptionInstanceId(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertRequest) Validate() error {
	return dara.Validate(s)
}
