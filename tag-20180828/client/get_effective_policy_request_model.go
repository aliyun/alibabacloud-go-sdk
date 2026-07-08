// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEffectivePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetEffectivePolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetEffectivePolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetEffectivePolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetEffectivePolicyRequest
	GetResourceOwnerAccount() *string
	SetTagKeys(v []*string) *GetEffectivePolicyRequest
	GetTagKeys() []*string
	SetTargetId(v string) *GetEffectivePolicyRequest
	GetTargetId() *string
	SetTargetType(v string) *GetEffectivePolicyRequest
	GetTargetType() *string
}

type GetEffectivePolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Only `cn-shanghai` is supported.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TagKeys              []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The ID of the target object.
	//
	// > This parameter is optional in Single-Account Mode and required in Multi-Account Mode.
	//
	// example:
	//
	// 154950938137****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the target object. Valid values:
	//
	// - USER: queries the effective policy for the current logon account. This value applies only to Single-Account Mode.
	//
	// - ROOT: queries the effective policy for the Root Folder in a Resource Directory. This value applies only to Multi-Account Mode.
	//
	// - FOLDER: queries the effective policy for a Folder in a Resource Directory. This value applies only to Multi-Account Mode.
	//
	// - ACCOUNT: queries the effective policy for a Member in a Resource Directory. This value applies only to Multi-Account Mode.
	//
	// > This parameter is optional in Single-Account Mode and required in Multi-Account Mode. The value is case-insensitive.
	//
	// example:
	//
	// ACCOUNT
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetEffectivePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEffectivePolicyRequest) GoString() string {
	return s.String()
}

func (s *GetEffectivePolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetEffectivePolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetEffectivePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEffectivePolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetEffectivePolicyRequest) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *GetEffectivePolicyRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *GetEffectivePolicyRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *GetEffectivePolicyRequest) SetOwnerAccount(v string) *GetEffectivePolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetOwnerId(v int64) *GetEffectivePolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetRegionId(v string) *GetEffectivePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetResourceOwnerAccount(v string) *GetEffectivePolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetTagKeys(v []*string) *GetEffectivePolicyRequest {
	s.TagKeys = v
	return s
}

func (s *GetEffectivePolicyRequest) SetTargetId(v string) *GetEffectivePolicyRequest {
	s.TargetId = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetTargetType(v string) *GetEffectivePolicyRequest {
	s.TargetType = &v
	return s
}

func (s *GetEffectivePolicyRequest) Validate() error {
	return dara.Validate(s)
}
