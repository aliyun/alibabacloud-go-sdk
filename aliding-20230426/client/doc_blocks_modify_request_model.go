// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksModifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockId(v string) *DocBlocksModifyRequest
	GetBlockId() *string
	SetDentryUuid(v string) *DocBlocksModifyRequest
	GetDentryUuid() *string
	SetElement(v map[string]interface{}) *DocBlocksModifyRequest
	GetElement() map[string]interface{}
	SetTenantContext(v *DocBlocksModifyRequestTenantContext) *DocBlocksModifyRequest
	GetTenantContext() *DocBlocksModifyRequestTenantContext
}

type DocBlocksModifyRequest struct {
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
	Element       map[string]interface{}               `json:"Element,omitempty" xml:"Element,omitempty"`
	TenantContext *DocBlocksModifyRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DocBlocksModifyRequest) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksModifyRequest) GoString() string {
	return s.String()
}

func (s *DocBlocksModifyRequest) GetBlockId() *string {
	return s.BlockId
}

func (s *DocBlocksModifyRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *DocBlocksModifyRequest) GetElement() map[string]interface{} {
	return s.Element
}

func (s *DocBlocksModifyRequest) GetTenantContext() *DocBlocksModifyRequestTenantContext {
	return s.TenantContext
}

func (s *DocBlocksModifyRequest) SetBlockId(v string) *DocBlocksModifyRequest {
	s.BlockId = &v
	return s
}

func (s *DocBlocksModifyRequest) SetDentryUuid(v string) *DocBlocksModifyRequest {
	s.DentryUuid = &v
	return s
}

func (s *DocBlocksModifyRequest) SetElement(v map[string]interface{}) *DocBlocksModifyRequest {
	s.Element = v
	return s
}

func (s *DocBlocksModifyRequest) SetTenantContext(v *DocBlocksModifyRequestTenantContext) *DocBlocksModifyRequest {
	s.TenantContext = v
	return s
}

func (s *DocBlocksModifyRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocBlocksModifyRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DocBlocksModifyRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksModifyRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DocBlocksModifyRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DocBlocksModifyRequestTenantContext) SetTenantId(v string) *DocBlocksModifyRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DocBlocksModifyRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
