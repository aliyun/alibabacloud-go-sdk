// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGraphDraftBatchDefineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDraftChangeIds(v []*int64) *SaveGraphDraftBatchDefineRequest
	GetDraftChangeIds() []*int64
	SetGraphName(v string) *SaveGraphDraftBatchDefineRequest
	GetGraphName() *string
	SetSaveMode(v string) *SaveGraphDraftBatchDefineRequest
	GetSaveMode() *string
	SetTenantId(v string) *SaveGraphDraftBatchDefineRequest
	GetTenantId() *string
	SetYamlEdit(v string) *SaveGraphDraftBatchDefineRequest
	GetYamlEdit() *string
}

type SaveGraphDraftBatchDefineRequest struct {
	// The list of draft change IDs.
	//
	// example:
	//
	// [401001, 401002]
	DraftChangeIds []*int64 `json:"draftChangeIds,omitempty" xml:"draftChangeIds,omitempty" type:"Repeated"`
	// The graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The save mode.
	//
	// example:
	//
	// FULL_YAML
	SaveMode *string `json:"saveMode,omitempty" xml:"saveMode,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The raw YAML text of the graph schema trimmed by READ permissions, with $ref references retained within the authorized subgraph.
	//
	// This parameter is required.
	//
	// example:
	//
	// name: crm_graph
	YamlEdit *string `json:"yamlEdit,omitempty" xml:"yamlEdit,omitempty"`
}

func (s SaveGraphDraftBatchDefineRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGraphDraftBatchDefineRequest) GoString() string {
	return s.String()
}

func (s *SaveGraphDraftBatchDefineRequest) GetDraftChangeIds() []*int64 {
	return s.DraftChangeIds
}

func (s *SaveGraphDraftBatchDefineRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *SaveGraphDraftBatchDefineRequest) GetSaveMode() *string {
	return s.SaveMode
}

func (s *SaveGraphDraftBatchDefineRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGraphDraftBatchDefineRequest) GetYamlEdit() *string {
	return s.YamlEdit
}

func (s *SaveGraphDraftBatchDefineRequest) SetDraftChangeIds(v []*int64) *SaveGraphDraftBatchDefineRequest {
	s.DraftChangeIds = v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetGraphName(v string) *SaveGraphDraftBatchDefineRequest {
	s.GraphName = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetSaveMode(v string) *SaveGraphDraftBatchDefineRequest {
	s.SaveMode = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetTenantId(v string) *SaveGraphDraftBatchDefineRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetYamlEdit(v string) *SaveGraphDraftBatchDefineRequest {
	s.YamlEdit = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) Validate() error {
	return dara.Validate(s)
}
