// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DeleteClientRequest
	GetClientId() *string
	SetResourceGroupId(v string) *DeleteClientRequest
	GetResourceGroupId() *string
	SetVaultId(v string) *DeleteClientRequest
	GetVaultId() *string
}

type DeleteClientRequest struct {
	// The ID of the client.
	//
	// example:
	//
	// c-000************f3h
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acf************kwy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-000************gs3
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteClientRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DeleteClientRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteClientRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DeleteClientRequest) SetClientId(v string) *DeleteClientRequest {
	s.ClientId = &v
	return s
}

func (s *DeleteClientRequest) SetResourceGroupId(v string) *DeleteClientRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteClientRequest) SetVaultId(v string) *DeleteClientRequest {
	s.VaultId = &v
	return s
}

func (s *DeleteClientRequest) Validate() error {
	return dara.Validate(s)
}
