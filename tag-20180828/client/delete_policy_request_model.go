// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeletePolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeletePolicyRequest
	GetOwnerId() *int64
	SetPolicyId(v string) *DeletePolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *DeletePolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeletePolicyRequest
	GetResourceOwnerAccount() *string
}

type DeletePolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the tag policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// p-557cb141331f41c7****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeletePolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeletePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePolicyRequest) SetOwnerAccount(v string) *DeletePolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeletePolicyRequest) SetOwnerId(v int64) *DeletePolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePolicyRequest) SetPolicyId(v string) *DeletePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeletePolicyRequest) SetRegionId(v string) *DeletePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolicyRequest) SetResourceOwnerAccount(v string) *DeletePolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePolicyRequest) Validate() error {
	return dara.Validate(s)
}
