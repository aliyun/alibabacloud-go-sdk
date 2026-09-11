// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphDraftAssembledResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGraphDraftAssembledResponseBody
	GetCode() *string
	SetGraphName(v string) *GetGraphDraftAssembledResponseBody
	GetGraphName() *string
	SetMessage(v string) *GetGraphDraftAssembledResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGraphDraftAssembledResponseBody
	GetRequestId() *string
	SetSchemaVersion(v string) *GetGraphDraftAssembledResponseBody
	GetSchemaVersion() *string
	SetYamlEdit(v string) *GetGraphDraftAssembledResponseBody
	GetYamlEdit() *string
}

type GetGraphDraftAssembledResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The knowledge graph name.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// The current zone list is illegal.
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
	// 1.2.0
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The raw Graph Schema YAML text trimmed by READ permissions, with $ref references within the authorized subgraph retained.
	//
	// This parameter is required.
	//
	// example:
	//
	// objects:\\n  - name: customer\\n
	YamlEdit *string `json:"yamlEdit,omitempty" xml:"yamlEdit,omitempty"`
}

func (s GetGraphDraftAssembledResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGraphDraftAssembledResponseBody) GoString() string {
	return s.String()
}

func (s *GetGraphDraftAssembledResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGraphDraftAssembledResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *GetGraphDraftAssembledResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGraphDraftAssembledResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGraphDraftAssembledResponseBody) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *GetGraphDraftAssembledResponseBody) GetYamlEdit() *string {
	return s.YamlEdit
}

func (s *GetGraphDraftAssembledResponseBody) SetCode(v string) *GetGraphDraftAssembledResponseBody {
	s.Code = &v
	return s
}

func (s *GetGraphDraftAssembledResponseBody) SetGraphName(v string) *GetGraphDraftAssembledResponseBody {
	s.GraphName = &v
	return s
}

func (s *GetGraphDraftAssembledResponseBody) SetMessage(v string) *GetGraphDraftAssembledResponseBody {
	s.Message = &v
	return s
}

func (s *GetGraphDraftAssembledResponseBody) SetRequestId(v string) *GetGraphDraftAssembledResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGraphDraftAssembledResponseBody) SetSchemaVersion(v string) *GetGraphDraftAssembledResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *GetGraphDraftAssembledResponseBody) SetYamlEdit(v string) *GetGraphDraftAssembledResponseBody {
	s.YamlEdit = &v
	return s
}

func (s *GetGraphDraftAssembledResponseBody) Validate() error {
	return dara.Validate(s)
}
