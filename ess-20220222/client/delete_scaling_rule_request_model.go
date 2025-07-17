// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteScalingRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteScalingRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteScalingRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteScalingRuleRequest
	GetResourceOwnerAccount() *string
	SetScalingRuleId(v string) *DeleteScalingRuleRequest
	GetScalingRuleId() *string
}

type DeleteScalingRuleRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling rule that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// asr-bp163l21e07uhnyt****
	ScalingRuleId *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
}

func (s DeleteScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteScalingRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteScalingRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteScalingRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteScalingRuleRequest) GetScalingRuleId() *string {
	return s.ScalingRuleId
}

func (s *DeleteScalingRuleRequest) SetOwnerAccount(v string) *DeleteScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetOwnerId(v int64) *DeleteScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetRegionId(v string) *DeleteScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetResourceOwnerAccount(v string) *DeleteScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetScalingRuleId(v string) *DeleteScalingRuleRequest {
	s.ScalingRuleId = &v
	return s
}

func (s *DeleteScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
