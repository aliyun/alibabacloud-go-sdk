// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockId(v string) *DocBlocksDeleteRequest
	GetBlockId() *string
	SetDentryUuid(v string) *DocBlocksDeleteRequest
	GetDentryUuid() *string
	SetTenantContext(v *DocBlocksDeleteRequestTenantContext) *DocBlocksDeleteRequest
	GetTenantContext() *DocBlocksDeleteRequestTenantContext
}

type DocBlocksDeleteRequest struct {
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
	DentryUuid    *string                              `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	TenantContext *DocBlocksDeleteRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DocBlocksDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksDeleteRequest) GoString() string {
	return s.String()
}

func (s *DocBlocksDeleteRequest) GetBlockId() *string {
	return s.BlockId
}

func (s *DocBlocksDeleteRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *DocBlocksDeleteRequest) GetTenantContext() *DocBlocksDeleteRequestTenantContext {
	return s.TenantContext
}

func (s *DocBlocksDeleteRequest) SetBlockId(v string) *DocBlocksDeleteRequest {
	s.BlockId = &v
	return s
}

func (s *DocBlocksDeleteRequest) SetDentryUuid(v string) *DocBlocksDeleteRequest {
	s.DentryUuid = &v
	return s
}

func (s *DocBlocksDeleteRequest) SetTenantContext(v *DocBlocksDeleteRequestTenantContext) *DocBlocksDeleteRequest {
	s.TenantContext = v
	return s
}

func (s *DocBlocksDeleteRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocBlocksDeleteRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DocBlocksDeleteRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksDeleteRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DocBlocksDeleteRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DocBlocksDeleteRequestTenantContext) SetTenantId(v string) *DocBlocksDeleteRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DocBlocksDeleteRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
