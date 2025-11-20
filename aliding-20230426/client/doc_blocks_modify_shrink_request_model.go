// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksModifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockId(v string) *DocBlocksModifyShrinkRequest
	GetBlockId() *string
	SetDentryUuid(v string) *DocBlocksModifyShrinkRequest
	GetDentryUuid() *string
	SetElementShrink(v string) *DocBlocksModifyShrinkRequest
	GetElementShrink() *string
	SetTenantContextShrink(v string) *DocBlocksModifyShrinkRequest
	GetTenantContextShrink() *string
}

type DocBlocksModifyShrinkRequest struct {
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
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"blockType":"paragraph","paragraph":{"text":"ok"}}
	ElementShrink       *string `json:"Element,omitempty" xml:"Element,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DocBlocksModifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksModifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DocBlocksModifyShrinkRequest) GetBlockId() *string {
	return s.BlockId
}

func (s *DocBlocksModifyShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *DocBlocksModifyShrinkRequest) GetElementShrink() *string {
	return s.ElementShrink
}

func (s *DocBlocksModifyShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DocBlocksModifyShrinkRequest) SetBlockId(v string) *DocBlocksModifyShrinkRequest {
	s.BlockId = &v
	return s
}

func (s *DocBlocksModifyShrinkRequest) SetDentryUuid(v string) *DocBlocksModifyShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *DocBlocksModifyShrinkRequest) SetElementShrink(v string) *DocBlocksModifyShrinkRequest {
	s.ElementShrink = &v
	return s
}

func (s *DocBlocksModifyShrinkRequest) SetTenantContextShrink(v string) *DocBlocksModifyShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DocBlocksModifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
