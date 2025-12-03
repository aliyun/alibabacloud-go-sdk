// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMapping(v string) *ModifySubscriptionMappingRequest
	GetMapping() *string
	SetOwnerId(v int64) *ModifySubscriptionMappingRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySubscriptionMappingRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySubscriptionMappingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySubscriptionMappingRequest
	GetResourceOwnerId() *int64
	SetSubscriptionId(v string) *ModifySubscriptionMappingRequest
	GetSubscriptionId() *string
}

type ModifySubscriptionMappingRequest struct {
	// This parameter is required.
	Mapping              *string `json:"Mapping,omitempty" xml:"Mapping,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s ModifySubscriptionMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionMappingRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionMappingRequest) GetMapping() *string {
	return s.Mapping
}

func (s *ModifySubscriptionMappingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySubscriptionMappingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySubscriptionMappingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySubscriptionMappingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySubscriptionMappingRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ModifySubscriptionMappingRequest) SetMapping(v string) *ModifySubscriptionMappingRequest {
	s.Mapping = &v
	return s
}

func (s *ModifySubscriptionMappingRequest) SetOwnerId(v int64) *ModifySubscriptionMappingRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySubscriptionMappingRequest) SetRegionId(v string) *ModifySubscriptionMappingRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySubscriptionMappingRequest) SetResourceOwnerAccount(v string) *ModifySubscriptionMappingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySubscriptionMappingRequest) SetResourceOwnerId(v int64) *ModifySubscriptionMappingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySubscriptionMappingRequest) SetSubscriptionId(v string) *ModifySubscriptionMappingRequest {
	s.SubscriptionId = &v
	return s
}

func (s *ModifySubscriptionMappingRequest) Validate() error {
	return dara.Validate(s)
}
