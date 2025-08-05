// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVaultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateVaultRequest
	GetDescription() *string
	SetResourceGroupId(v string) *UpdateVaultRequest
	GetResourceGroupId() *string
	SetVaultId(v string) *UpdateVaultRequest
	GetVaultId() *string
	SetVaultName(v string) *UpdateVaultRequest
	GetVaultName() *string
	SetWormEnabled(v bool) *UpdateVaultRequest
	GetWormEnabled() *bool
}

type UpdateVaultRequest struct {
	// The description of the backup vault. The description must be 0 to 255 characters in length.
	//
	// example:
	//
	// vault description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm2fa2xeiebyy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	// The name of the backup vault. The name must be 1 to 64 characters in length.
	//
	// example:
	//
	// vaultname
	VaultName *string `json:"VaultName,omitempty" xml:"VaultName,omitempty"`
	// Specifies whether to enable the immutable backup feature for storage vaults. After the immutable backup feature is enabled, backup vaults and all backup data cannot be deleted until the retention period expires. The immutable backup feature cannot be disabled after it is enabled. Only standard backup vaults and archive vaults support the immutable backup feature.
	//
	// example:
	//
	// true
	WormEnabled *bool `json:"WormEnabled,omitempty" xml:"WormEnabled,omitempty"`
}

func (s UpdateVaultRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVaultRequest) GoString() string {
	return s.String()
}

func (s *UpdateVaultRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateVaultRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateVaultRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *UpdateVaultRequest) GetVaultName() *string {
	return s.VaultName
}

func (s *UpdateVaultRequest) GetWormEnabled() *bool {
	return s.WormEnabled
}

func (s *UpdateVaultRequest) SetDescription(v string) *UpdateVaultRequest {
	s.Description = &v
	return s
}

func (s *UpdateVaultRequest) SetResourceGroupId(v string) *UpdateVaultRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateVaultRequest) SetVaultId(v string) *UpdateVaultRequest {
	s.VaultId = &v
	return s
}

func (s *UpdateVaultRequest) SetVaultName(v string) *UpdateVaultRequest {
	s.VaultName = &v
	return s
}

func (s *UpdateVaultRequest) SetWormEnabled(v bool) *UpdateVaultRequest {
	s.WormEnabled = &v
	return s
}

func (s *UpdateVaultRequest) Validate() error {
	return dara.Validate(s)
}
