// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ConvertClusterRequest
	GetClusterId() *string
	SetDuration(v int32) *ConvertClusterRequest
	GetDuration() *int32
	SetOwnerId(v int64) *ConvertClusterRequest
	GetOwnerId() *int64
	SetPricingCycle(v string) *ConvertClusterRequest
	GetPricingCycle() *string
	SetResourceOwnerAccount(v string) *ConvertClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ConvertClusterRequest
	GetResourceOwnerId() *int64
}

type ConvertClusterRequest struct {
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

func (s ConvertClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertClusterRequest) GoString() string {
	return s.String()
}

func (s *ConvertClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ConvertClusterRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ConvertClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ConvertClusterRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *ConvertClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ConvertClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ConvertClusterRequest) SetClusterId(v string) *ConvertClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ConvertClusterRequest) SetDuration(v int32) *ConvertClusterRequest {
	s.Duration = &v
	return s
}

func (s *ConvertClusterRequest) SetOwnerId(v int64) *ConvertClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ConvertClusterRequest) SetPricingCycle(v string) *ConvertClusterRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConvertClusterRequest) SetResourceOwnerAccount(v string) *ConvertClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConvertClusterRequest) SetResourceOwnerId(v int64) *ConvertClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConvertClusterRequest) Validate() error {
	return dara.Validate(s)
}
