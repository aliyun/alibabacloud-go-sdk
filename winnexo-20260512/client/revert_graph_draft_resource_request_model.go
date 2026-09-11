// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertGraphDraftResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDraftChangeId(v int64) *RevertGraphDraftResourceRequest
	GetDraftChangeId() *int64
	SetGraphName(v string) *RevertGraphDraftResourceRequest
	GetGraphName() *string
	SetTenantId(v string) *RevertGraphDraftResourceRequest
	GetTenantId() *string
}

type RevertGraphDraftResourceRequest struct {
	// The draft change ID (the draftChangeId returned by listGraphDraftResources).
	//
	// This parameter is required.
	//
	// example:
	//
	// 401001
	DraftChangeId *int64 `json:"draftChangeId,omitempty" xml:"draftChangeId,omitempty"`
	// The knowledge graph name.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The tenant ID. This is a common parameter. Pass it explicitly by using --tenant-id in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RevertGraphDraftResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RevertGraphDraftResourceRequest) GoString() string {
	return s.String()
}

func (s *RevertGraphDraftResourceRequest) GetDraftChangeId() *int64 {
	return s.DraftChangeId
}

func (s *RevertGraphDraftResourceRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *RevertGraphDraftResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RevertGraphDraftResourceRequest) SetDraftChangeId(v int64) *RevertGraphDraftResourceRequest {
	s.DraftChangeId = &v
	return s
}

func (s *RevertGraphDraftResourceRequest) SetGraphName(v string) *RevertGraphDraftResourceRequest {
	s.GraphName = &v
	return s
}

func (s *RevertGraphDraftResourceRequest) SetTenantId(v string) *RevertGraphDraftResourceRequest {
	s.TenantId = &v
	return s
}

func (s *RevertGraphDraftResourceRequest) Validate() error {
	return dara.Validate(s)
}
