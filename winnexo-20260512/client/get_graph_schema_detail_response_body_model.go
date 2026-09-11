// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphSchemaDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessProfile(v string) *GetGraphSchemaDetailResponseBody
	GetBusinessProfile() *string
	SetCode(v string) *GetGraphSchemaDetailResponseBody
	GetCode() *string
	SetContentHash(v string) *GetGraphSchemaDetailResponseBody
	GetContentHash() *string
	SetCreatedBy(v string) *GetGraphSchemaDetailResponseBody
	GetCreatedBy() *string
	SetDisplayName(v string) *GetGraphSchemaDetailResponseBody
	GetDisplayName() *string
	SetGmtCreate(v string) *GetGraphSchemaDetailResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetGraphSchemaDetailResponseBody
	GetGmtModified() *string
	SetGraphName(v string) *GetGraphSchemaDetailResponseBody
	GetGraphName() *string
	SetGraphStatus(v string) *GetGraphSchemaDetailResponseBody
	GetGraphStatus() *string
	SetHasDraft(v bool) *GetGraphSchemaDetailResponseBody
	GetHasDraft() *bool
	SetMessage(v string) *GetGraphSchemaDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGraphSchemaDetailResponseBody
	GetRequestId() *string
	SetSchemaVersion(v string) *GetGraphSchemaDetailResponseBody
	GetSchemaVersion() *string
	SetYamlEdit(v string) *GetGraphSchemaDetailResponseBody
	GetYamlEdit() *string
}

type GetGraphSchemaDetailResponseBody struct {
	// The business description of the graph. An empty string is returned if this parameter is not configured.
	//
	// example:
	//
	// Customer domain semantic graph
	BusinessProfile *string `json:"businessProfile,omitempty" xml:"businessProfile,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The hash fingerprint of the schema content.
	//
	// example:
	//
	// a1b2c3
	ContentHash *string `json:"contentHash,omitempty" xml:"contentHash,omitempty"`
	// The creator.
	//
	// example:
	//
	// u001
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The display name.
	//
	// example:
	//
	// CRM Graph
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-09-08T10:00:00+00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 2026-09-08T11:30:00+00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The name of the graph.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The status of the semantic graph.
	//
	// example:
	//
	// PUBLISHING: A publish task is in progress for this graph.
	//
	// DEVELOPING: An active draft exists for this graph (being edited, not yet published).
	//
	// PUBLISHED: Normal status
	GraphStatus *string `json:"graphStatus,omitempty" xml:"graphStatus,omitempty"`
	// Indicates whether the graph contains a draft.
	//
	// example:
	//
	// false
	HasDraft *bool `json:"hasDraft,omitempty" xml:"hasDraft,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The version.
	//
	// example:
	//
	// 0.0.0
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The original YAML text of the graph schema trimmed by READ permission. The $ref references within the authorized subgraph are retained.
	//
	// example:
	//
	// name: crm_graph
	YamlEdit *string `json:"yamlEdit,omitempty" xml:"yamlEdit,omitempty"`
}

func (s GetGraphSchemaDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGraphSchemaDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaDetailResponseBody) GetBusinessProfile() *string {
	return s.BusinessProfile
}

func (s *GetGraphSchemaDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGraphSchemaDetailResponseBody) GetContentHash() *string {
	return s.ContentHash
}

func (s *GetGraphSchemaDetailResponseBody) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetGraphSchemaDetailResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetGraphSchemaDetailResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGraphSchemaDetailResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGraphSchemaDetailResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *GetGraphSchemaDetailResponseBody) GetGraphStatus() *string {
	return s.GraphStatus
}

func (s *GetGraphSchemaDetailResponseBody) GetHasDraft() *bool {
	return s.HasDraft
}

func (s *GetGraphSchemaDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGraphSchemaDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGraphSchemaDetailResponseBody) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *GetGraphSchemaDetailResponseBody) GetYamlEdit() *string {
	return s.YamlEdit
}

func (s *GetGraphSchemaDetailResponseBody) SetBusinessProfile(v string) *GetGraphSchemaDetailResponseBody {
	s.BusinessProfile = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetCode(v string) *GetGraphSchemaDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetContentHash(v string) *GetGraphSchemaDetailResponseBody {
	s.ContentHash = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetCreatedBy(v string) *GetGraphSchemaDetailResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetDisplayName(v string) *GetGraphSchemaDetailResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetGmtCreate(v string) *GetGraphSchemaDetailResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetGmtModified(v string) *GetGraphSchemaDetailResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetGraphName(v string) *GetGraphSchemaDetailResponseBody {
	s.GraphName = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetGraphStatus(v string) *GetGraphSchemaDetailResponseBody {
	s.GraphStatus = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetHasDraft(v bool) *GetGraphSchemaDetailResponseBody {
	s.HasDraft = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetMessage(v string) *GetGraphSchemaDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetRequestId(v string) *GetGraphSchemaDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetSchemaVersion(v string) *GetGraphSchemaDetailResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) SetYamlEdit(v string) *GetGraphSchemaDetailResponseBody {
	s.YamlEdit = &v
	return s
}

func (s *GetGraphSchemaDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
