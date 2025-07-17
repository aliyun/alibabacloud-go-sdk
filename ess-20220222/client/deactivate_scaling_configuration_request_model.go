// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateScalingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeactivateScalingConfigurationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeactivateScalingConfigurationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeactivateScalingConfigurationRequest
	GetResourceOwnerAccount() *string
	SetScalingConfigurationId(v string) *DeactivateScalingConfigurationRequest
	GetScalingConfigurationId() *string
}

type DeactivateScalingConfigurationRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// asc-bp1ahp2ud7qkzt2a****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s DeactivateScalingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeactivateScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeactivateScalingConfigurationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeactivateScalingConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeactivateScalingConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeactivateScalingConfigurationRequest) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DeactivateScalingConfigurationRequest) SetOwnerAccount(v string) *DeactivateScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeactivateScalingConfigurationRequest) SetOwnerId(v int64) *DeactivateScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeactivateScalingConfigurationRequest) SetResourceOwnerAccount(v string) *DeactivateScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeactivateScalingConfigurationRequest) SetScalingConfigurationId(v string) *DeactivateScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DeactivateScalingConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
