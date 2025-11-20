// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksDeleteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockId(v string) *DocBlocksDeleteShrinkRequest
	GetBlockId() *string
	SetDentryUuid(v string) *DocBlocksDeleteShrinkRequest
	GetDentryUuid() *string
	SetTenantContextShrink(v string) *DocBlocksDeleteShrinkRequest
	GetTenantContextShrink() *string
}

type DocBlocksDeleteShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mgokkwiovcq5eu02le8
	BlockId *string `json:"BlockId,omitempty" xml:"BlockId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Obva6QBXJwxNZoMOC9bE11Zb8n4qY5Pr
	DentryUuid          *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DocBlocksDeleteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksDeleteShrinkRequest) GoString() string {
	return s.String()
}

func (s *DocBlocksDeleteShrinkRequest) GetBlockId() *string {
	return s.BlockId
}

func (s *DocBlocksDeleteShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *DocBlocksDeleteShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DocBlocksDeleteShrinkRequest) SetBlockId(v string) *DocBlocksDeleteShrinkRequest {
	s.BlockId = &v
	return s
}

func (s *DocBlocksDeleteShrinkRequest) SetDentryUuid(v string) *DocBlocksDeleteShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *DocBlocksDeleteShrinkRequest) SetTenantContextShrink(v string) *DocBlocksDeleteShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DocBlocksDeleteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
