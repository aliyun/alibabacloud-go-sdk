// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVaultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DeleteVaultRequest
	GetResourceGroupId() *string
	SetToken(v string) *DeleteVaultRequest
	GetToken() *string
	SetVaultId(v string) *DeleteVaultRequest
	GetVaultId() *string
}

type DeleteVaultRequest struct {
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmoiyerpacj4q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The token.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a*
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteVaultRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVaultRequest) GoString() string {
	return s.String()
}

func (s *DeleteVaultRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteVaultRequest) GetToken() *string {
	return s.Token
}

func (s *DeleteVaultRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DeleteVaultRequest) SetResourceGroupId(v string) *DeleteVaultRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteVaultRequest) SetToken(v string) *DeleteVaultRequest {
	s.Token = &v
	return s
}

func (s *DeleteVaultRequest) SetVaultId(v string) *DeleteVaultRequest {
	s.VaultId = &v
	return s
}

func (s *DeleteVaultRequest) Validate() error {
	return dara.Validate(s)
}
