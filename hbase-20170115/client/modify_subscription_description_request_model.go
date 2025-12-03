// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ModifySubscriptionDescriptionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySubscriptionDescriptionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySubscriptionDescriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySubscriptionDescriptionRequest
	GetResourceOwnerId() *int64
	SetSubscriptionDescription(v string) *ModifySubscriptionDescriptionRequest
	GetSubscriptionDescription() *string
	SetSubscriptionId(v string) *ModifySubscriptionDescriptionRequest
	GetSubscriptionId() *string
}

type ModifySubscriptionDescriptionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SubscriptionDescription *string `json:"SubscriptionDescription,omitempty" xml:"SubscriptionDescription,omitempty"`
	// This parameter is required.
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s ModifySubscriptionDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionDescriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySubscriptionDescriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySubscriptionDescriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySubscriptionDescriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySubscriptionDescriptionRequest) GetSubscriptionDescription() *string {
	return s.SubscriptionDescription
}

func (s *ModifySubscriptionDescriptionRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ModifySubscriptionDescriptionRequest) SetOwnerId(v int64) *ModifySubscriptionDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySubscriptionDescriptionRequest) SetRegionId(v string) *ModifySubscriptionDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySubscriptionDescriptionRequest) SetResourceOwnerAccount(v string) *ModifySubscriptionDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySubscriptionDescriptionRequest) SetResourceOwnerId(v int64) *ModifySubscriptionDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySubscriptionDescriptionRequest) SetSubscriptionDescription(v string) *ModifySubscriptionDescriptionRequest {
	s.SubscriptionDescription = &v
	return s
}

func (s *ModifySubscriptionDescriptionRequest) SetSubscriptionId(v string) *ModifySubscriptionDescriptionRequest {
	s.SubscriptionId = &v
	return s
}

func (s *ModifySubscriptionDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
