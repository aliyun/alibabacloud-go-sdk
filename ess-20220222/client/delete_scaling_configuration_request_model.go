// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteScalingConfigurationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteScalingConfigurationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteScalingConfigurationRequest
	GetResourceOwnerAccount() *string
	SetScalingConfigurationId(v string) *DeleteScalingConfigurationRequest
	GetScalingConfigurationId() *string
}

type DeleteScalingConfigurationRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling configuration that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// asc-bp1bx8mzur534edp****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s DeleteScalingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigurationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteScalingConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteScalingConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteScalingConfigurationRequest) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DeleteScalingConfigurationRequest) SetOwnerAccount(v string) *DeleteScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) SetOwnerId(v int64) *DeleteScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) SetResourceOwnerAccount(v string) *DeleteScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) SetScalingConfigurationId(v string) *DeleteScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
