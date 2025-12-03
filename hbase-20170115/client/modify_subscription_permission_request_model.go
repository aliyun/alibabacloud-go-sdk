// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ModifySubscriptionPermissionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySubscriptionPermissionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySubscriptionPermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySubscriptionPermissionRequest
	GetResourceOwnerId() *int64
	SetStatus(v int32) *ModifySubscriptionPermissionRequest
	GetStatus() *int32
}

type ModifySubscriptionPermissionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifySubscriptionPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionPermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionPermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySubscriptionPermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySubscriptionPermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySubscriptionPermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySubscriptionPermissionRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifySubscriptionPermissionRequest) SetOwnerId(v int64) *ModifySubscriptionPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySubscriptionPermissionRequest) SetRegionId(v string) *ModifySubscriptionPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySubscriptionPermissionRequest) SetResourceOwnerAccount(v string) *ModifySubscriptionPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySubscriptionPermissionRequest) SetResourceOwnerId(v int64) *ModifySubscriptionPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySubscriptionPermissionRequest) SetStatus(v int32) *ModifySubscriptionPermissionRequest {
	s.Status = &v
	return s
}

func (s *ModifySubscriptionPermissionRequest) Validate() error {
	return dara.Validate(s)
}
