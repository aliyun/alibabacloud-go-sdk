// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEciScalingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteEciScalingConfigurationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteEciScalingConfigurationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteEciScalingConfigurationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteEciScalingConfigurationRequest
	GetResourceOwnerAccount() *string
	SetScalingConfigurationId(v string) *DeleteEciScalingConfigurationRequest
	GetScalingConfigurationId() *string
}

type DeleteEciScalingConfigurationRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s DeleteEciScalingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEciScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteEciScalingConfigurationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteEciScalingConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteEciScalingConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEciScalingConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteEciScalingConfigurationRequest) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DeleteEciScalingConfigurationRequest) SetOwnerAccount(v string) *DeleteEciScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteEciScalingConfigurationRequest) SetOwnerId(v int64) *DeleteEciScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteEciScalingConfigurationRequest) SetRegionId(v string) *DeleteEciScalingConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEciScalingConfigurationRequest) SetResourceOwnerAccount(v string) *DeleteEciScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteEciScalingConfigurationRequest) SetScalingConfigurationId(v string) *DeleteEciScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DeleteEciScalingConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
