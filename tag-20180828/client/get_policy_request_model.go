// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetPolicyRequest
	GetOwnerId() *int64
	SetPolicyId(v string) *GetPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *GetPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetPolicyRequest
	GetResourceOwnerAccount() *string
}

type GetPolicyRequest struct {
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

func (s GetPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPolicyRequest) SetOwnerAccount(v string) *GetPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetPolicyRequest) SetOwnerId(v int64) *GetPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPolicyRequest) SetPolicyId(v string) *GetPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyRequest) SetRegionId(v string) *GetPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyRequest) SetResourceOwnerAccount(v string) *GetPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPolicyRequest) Validate() error {
	return dara.Validate(s)
}
