// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RenewClusterRequest
	GetClusterId() *string
	SetDuration(v int32) *RenewClusterRequest
	GetDuration() *int32
	SetOwnerId(v int64) *RenewClusterRequest
	GetOwnerId() *int64
	SetPricingCycle(v string) *RenewClusterRequest
	GetPricingCycle() *string
	SetResourceOwnerAccount(v string) *RenewClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewClusterRequest
	GetResourceOwnerId() *int64
}

type RenewClusterRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RenewClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewClusterRequest) GoString() string {
	return s.String()
}

func (s *RenewClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RenewClusterRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *RenewClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewClusterRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *RenewClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewClusterRequest) SetClusterId(v string) *RenewClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *RenewClusterRequest) SetDuration(v int32) *RenewClusterRequest {
	s.Duration = &v
	return s
}

func (s *RenewClusterRequest) SetOwnerId(v int64) *RenewClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewClusterRequest) SetPricingCycle(v string) *RenewClusterRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewClusterRequest) SetResourceOwnerAccount(v string) *RenewClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewClusterRequest) SetResourceOwnerId(v int64) *RenewClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewClusterRequest) Validate() error {
	return dara.Validate(s)
}
