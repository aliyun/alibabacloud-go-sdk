// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *UninstallClientRequest
	GetClientId() *string
	SetResourceGroupId(v string) *UninstallClientRequest
	GetResourceGroupId() *string
	SetVaultId(v string) *UninstallClientRequest
	GetVaultId() *string
}

type UninstallClientRequest struct {
	// The ID of the HBR client.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-000iuqo******zi3rn
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm3erpwweavki
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-0008n2q******ax3
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UninstallClientRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallClientRequest) GoString() string {
	return s.String()
}

func (s *UninstallClientRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UninstallClientRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UninstallClientRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *UninstallClientRequest) SetClientId(v string) *UninstallClientRequest {
	s.ClientId = &v
	return s
}

func (s *UninstallClientRequest) SetResourceGroupId(v string) *UninstallClientRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UninstallClientRequest) SetVaultId(v string) *UninstallClientRequest {
	s.VaultId = &v
	return s
}

func (s *UninstallClientRequest) Validate() error {
	return dara.Validate(s)
}
