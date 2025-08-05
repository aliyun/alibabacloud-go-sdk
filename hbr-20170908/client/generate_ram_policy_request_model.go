// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateRamPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *GenerateRamPolicyRequest
	GetActionType() *string
	SetRequireBasePolicy(v bool) *GenerateRamPolicyRequest
	GetRequireBasePolicy() *bool
	SetResourceGroupId(v string) *GenerateRamPolicyRequest
	GetResourceGroupId() *string
	SetVaultId(v string) *GenerateRamPolicyRequest
	GetVaultId() *string
}

type GenerateRamPolicyRequest struct {
	// The type of policy that you want to generate. Valid values:
	//
	// 	- BACKUP: the permission to back up data to a backup vault
	//
	// 	- RESTORE: the permission to restore data from a backup vault
	//
	// This parameter is required.
	//
	// example:
	//
	// system
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// Specifies whether to generate the policy based on an existing instance-specific rule. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	RequireBasePolicy *bool `json:"RequireBasePolicy,omitempty" xml:"RequireBasePolicy,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-*********************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-0007al3m******7ao
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s GenerateRamPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateRamPolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateRamPolicyRequest) GetActionType() *string {
	return s.ActionType
}

func (s *GenerateRamPolicyRequest) GetRequireBasePolicy() *bool {
	return s.RequireBasePolicy
}

func (s *GenerateRamPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GenerateRamPolicyRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *GenerateRamPolicyRequest) SetActionType(v string) *GenerateRamPolicyRequest {
	s.ActionType = &v
	return s
}

func (s *GenerateRamPolicyRequest) SetRequireBasePolicy(v bool) *GenerateRamPolicyRequest {
	s.RequireBasePolicy = &v
	return s
}

func (s *GenerateRamPolicyRequest) SetResourceGroupId(v string) *GenerateRamPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GenerateRamPolicyRequest) SetVaultId(v string) *GenerateRamPolicyRequest {
	s.VaultId = &v
	return s
}

func (s *GenerateRamPolicyRequest) Validate() error {
	return dara.Validate(s)
}
