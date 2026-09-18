// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGraphDraftBatchDefineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDraftChangeIdsShrink(v string) *SaveGraphDraftBatchDefineShrinkRequest
	GetDraftChangeIdsShrink() *string
	SetGraphName(v string) *SaveGraphDraftBatchDefineShrinkRequest
	GetGraphName() *string
	SetSaveMode(v string) *SaveGraphDraftBatchDefineShrinkRequest
	GetSaveMode() *string
	SetTenantId(v string) *SaveGraphDraftBatchDefineShrinkRequest
	GetTenantId() *string
	SetYamlEdit(v string) *SaveGraphDraftBatchDefineShrinkRequest
	GetYamlEdit() *string
}

type SaveGraphDraftBatchDefineShrinkRequest struct {
	// The list of draft change IDs.
	//
	// example:
	//
	// [401001, 401002]
	DraftChangeIdsShrink *string `json:"draftChangeIds,omitempty" xml:"draftChangeIds,omitempty"`
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

func (s SaveGraphDraftBatchDefineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGraphDraftBatchDefineShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) GetDraftChangeIdsShrink() *string {
	return s.DraftChangeIdsShrink
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) GetSaveMode() *string {
	return s.SaveMode
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) GetYamlEdit() *string {
	return s.YamlEdit
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) SetDraftChangeIdsShrink(v string) *SaveGraphDraftBatchDefineShrinkRequest {
	s.DraftChangeIdsShrink = &v
	return s
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) SetGraphName(v string) *SaveGraphDraftBatchDefineShrinkRequest {
	s.GraphName = &v
	return s
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) SetSaveMode(v string) *SaveGraphDraftBatchDefineShrinkRequest {
	s.SaveMode = &v
	return s
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) SetTenantId(v string) *SaveGraphDraftBatchDefineShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) SetYamlEdit(v string) *SaveGraphDraftBatchDefineShrinkRequest {
	s.YamlEdit = &v
	return s
}

func (s *SaveGraphDraftBatchDefineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
