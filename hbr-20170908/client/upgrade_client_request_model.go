// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *UpgradeClientRequest
	GetClientId() *string
	SetResourceGroupId(v string) *UpgradeClientRequest
	GetResourceGroupId() *string
	SetVaultId(v string) *UpgradeClientRequest
	GetVaultId() *string
}

type UpgradeClientRequest struct {
	// The ID of the Cloud Backup client.
	//
	// example:
	//
	// c-000boklw******63a9
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy6uja5wyc2i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-000djw8ci******3ic
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpgradeClientRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClientRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClientRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UpgradeClientRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpgradeClientRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *UpgradeClientRequest) SetClientId(v string) *UpgradeClientRequest {
	s.ClientId = &v
	return s
}

func (s *UpgradeClientRequest) SetResourceGroupId(v string) *UpgradeClientRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpgradeClientRequest) SetVaultId(v string) *UpgradeClientRequest {
	s.VaultId = &v
	return s
}

func (s *UpgradeClientRequest) Validate() error {
	return dara.Validate(s)
}
