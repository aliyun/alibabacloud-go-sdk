// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGraphSchemaResponseBody
	GetCode() *string
	SetGraphName(v string) *GetGraphSchemaResponseBody
	GetGraphName() *string
	SetMessage(v string) *GetGraphSchemaResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGraphSchemaResponseBody
	GetRequestId() *string
	SetSchemaId(v string) *GetGraphSchemaResponseBody
	GetSchemaId() *string
	SetSchemaVersion(v string) *GetGraphSchemaResponseBody
	GetSchemaVersion() *string
	SetYamlEdit(v string) *GetGraphSchemaResponseBody
	GetYamlEdit() *string
}

type GetGraphSchemaResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The status code description.
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
	// The active QueryAgent registered schema ID corresponding to the graph. The value is null if not yet registered.
	//
	// example:
	//
	// schema_123456
	SchemaId *string `json:"schemaId,omitempty" xml:"schemaId,omitempty"`
	// The version.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The raw YAML text of the Graph Schema trimmed by READ permissions, retaining $ref references within the authorized subgraph.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	YamlEdit *string `json:"yamlEdit,omitempty" xml:"yamlEdit,omitempty"`
}

func (s GetGraphSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGraphSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGraphSchemaResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *GetGraphSchemaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGraphSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGraphSchemaResponseBody) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetGraphSchemaResponseBody) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *GetGraphSchemaResponseBody) GetYamlEdit() *string {
	return s.YamlEdit
}

func (s *GetGraphSchemaResponseBody) SetCode(v string) *GetGraphSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetGraphName(v string) *GetGraphSchemaResponseBody {
	s.GraphName = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetMessage(v string) *GetGraphSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetRequestId(v string) *GetGraphSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetSchemaId(v string) *GetGraphSchemaResponseBody {
	s.SchemaId = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetSchemaVersion(v string) *GetGraphSchemaResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetYamlEdit(v string) *GetGraphSchemaResponseBody {
	s.YamlEdit = &v
	return s
}

func (s *GetGraphSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
