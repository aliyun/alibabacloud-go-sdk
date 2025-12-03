// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ReleaseSubscriptionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ReleaseSubscriptionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ReleaseSubscriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseSubscriptionRequest
	GetResourceOwnerId() *int64
	SetSubscriptionId(v string) *ReleaseSubscriptionRequest
	GetSubscriptionId() *string
}

type ReleaseSubscriptionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s ReleaseSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseSubscriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseSubscriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseSubscriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseSubscriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseSubscriptionRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ReleaseSubscriptionRequest) SetOwnerId(v int64) *ReleaseSubscriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseSubscriptionRequest) SetRegionId(v string) *ReleaseSubscriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseSubscriptionRequest) SetResourceOwnerAccount(v string) *ReleaseSubscriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseSubscriptionRequest) SetResourceOwnerId(v int64) *ReleaseSubscriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseSubscriptionRequest) SetSubscriptionId(v string) *ReleaseSubscriptionRequest {
	s.SubscriptionId = &v
	return s
}

func (s *ReleaseSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
