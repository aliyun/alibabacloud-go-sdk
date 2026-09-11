// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphDraftAssembledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGraphName(v string) *GetGraphDraftAssembledRequest
	GetGraphName() *string
	SetTenantId(v string) *GetGraphDraftAssembledRequest
	GetTenantId() *string
}

type GetGraphDraftAssembledRequest struct {
	// The knowledge graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetGraphDraftAssembledRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGraphDraftAssembledRequest) GoString() string {
	return s.String()
}

func (s *GetGraphDraftAssembledRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *GetGraphDraftAssembledRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGraphDraftAssembledRequest) SetGraphName(v string) *GetGraphDraftAssembledRequest {
	s.GraphName = &v
	return s
}

func (s *GetGraphDraftAssembledRequest) SetTenantId(v string) *GetGraphDraftAssembledRequest {
	s.TenantId = &v
	return s
}

func (s *GetGraphDraftAssembledRequest) Validate() error {
	return dara.Validate(s)
}
