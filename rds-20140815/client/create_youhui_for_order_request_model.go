// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYouhuiForOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *CreateYouhuiForOrderRequest
	GetActivityId() *int64
	SetOwnerId(v string) *CreateYouhuiForOrderRequest
	GetOwnerId() *string
	SetPromotionId(v int64) *CreateYouhuiForOrderRequest
	GetPromotionId() *int64
	SetRegionId(v string) *CreateYouhuiForOrderRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateYouhuiForOrderRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateYouhuiForOrderRequest
	GetResourceOwnerId() *int64
}

type CreateYouhuiForOrderRequest struct {
	// The activity ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1711510887******
	ActivityId *int64  `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The promotion ID. You can call the GetResourcePrice operation to query the promotion ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2000001******
	PromotionId *int64 `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateYouhuiForOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateYouhuiForOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateYouhuiForOrderRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *CreateYouhuiForOrderRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateYouhuiForOrderRequest) GetPromotionId() *int64 {
	return s.PromotionId
}

func (s *CreateYouhuiForOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateYouhuiForOrderRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateYouhuiForOrderRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateYouhuiForOrderRequest) SetActivityId(v int64) *CreateYouhuiForOrderRequest {
	s.ActivityId = &v
	return s
}

func (s *CreateYouhuiForOrderRequest) SetOwnerId(v string) *CreateYouhuiForOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateYouhuiForOrderRequest) SetPromotionId(v int64) *CreateYouhuiForOrderRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateYouhuiForOrderRequest) SetRegionId(v string) *CreateYouhuiForOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CreateYouhuiForOrderRequest) SetResourceOwnerAccount(v string) *CreateYouhuiForOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateYouhuiForOrderRequest) SetResourceOwnerId(v int64) *CreateYouhuiForOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateYouhuiForOrderRequest) Validate() error {
	return dara.Validate(s)
}
