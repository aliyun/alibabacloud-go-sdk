// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVaultReplicationRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *DescribeVaultReplicationRegionsRequest
	GetToken() *string
	SetVaultId(v string) *DescribeVaultReplicationRegionsRequest
	GetVaultId() *string
}

type DescribeVaultReplicationRegionsRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 01W3ZZOQ
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// example:
	//
	// v-00030j3c******sn
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeVaultReplicationRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultReplicationRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVaultReplicationRegionsRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeVaultReplicationRegionsRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeVaultReplicationRegionsRequest) SetToken(v string) *DescribeVaultReplicationRegionsRequest {
	s.Token = &v
	return s
}

func (s *DescribeVaultReplicationRegionsRequest) SetVaultId(v string) *DescribeVaultReplicationRegionsRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeVaultReplicationRegionsRequest) Validate() error {
	return dara.Validate(s)
}
