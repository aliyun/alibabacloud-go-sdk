// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInventorySchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *GetInventorySchemaResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *GetInventorySchemaResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetInventorySchemaResponseBody
	GetRequestId() *string
	SetSchemas(v []*GetInventorySchemaResponseBodySchemas) *GetInventorySchemaResponseBody
	GetSchemas() []*GetInventorySchemaResponseBodySchemas
}

type GetInventorySchemaResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 1
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that was used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfh8MVLQI9AuKGACLgjbsXbWs-Mna47IDM6tr6wK7TZ1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 89117642-7167-4F4D-B7F1-876582279E3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed configurations of the configuration list.
	Schemas []*GetInventorySchemaResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
}

func (s GetInventorySchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *GetInventorySchemaResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetInventorySchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInventorySchemaResponseBody) GetSchemas() []*GetInventorySchemaResponseBodySchemas {
	return s.Schemas
}

func (s *GetInventorySchemaResponseBody) SetMaxResults(v string) *GetInventorySchemaResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetInventorySchemaResponseBody) SetNextToken(v string) *GetInventorySchemaResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetInventorySchemaResponseBody) SetRequestId(v string) *GetInventorySchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInventorySchemaResponseBody) SetSchemas(v []*GetInventorySchemaResponseBodySchemas) *GetInventorySchemaResponseBody {
	s.Schemas = v
	return s
}

func (s *GetInventorySchemaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInventorySchemaResponseBodySchemas struct {
	// The properties of the configuration list.
	Attributes []*GetInventorySchemaResponseBodySchemasAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The name of the configuration list.
	//
	// example:
	//
	// ACS:Application
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// The version of the configuration list.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetInventorySchemaResponseBodySchemas) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySchemaResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponseBodySchemas) GetAttributes() []*GetInventorySchemaResponseBodySchemasAttributes {
	return s.Attributes
}

func (s *GetInventorySchemaResponseBodySchemas) GetTypeName() *string {
	return s.TypeName
}

func (s *GetInventorySchemaResponseBodySchemas) GetVersion() *string {
	return s.Version
}

func (s *GetInventorySchemaResponseBodySchemas) SetAttributes(v []*GetInventorySchemaResponseBodySchemasAttributes) *GetInventorySchemaResponseBodySchemas {
	s.Attributes = v
	return s
}

func (s *GetInventorySchemaResponseBodySchemas) SetTypeName(v string) *GetInventorySchemaResponseBodySchemas {
	s.TypeName = &v
	return s
}

func (s *GetInventorySchemaResponseBodySchemas) SetVersion(v string) *GetInventorySchemaResponseBodySchemas {
	s.Version = &v
	return s
}

func (s *GetInventorySchemaResponseBodySchemas) Validate() error {
	return dara.Validate(s)
}

type GetInventorySchemaResponseBodySchemasAttributes struct {
	// The data type of the property.
	//
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The name of the property.
	//
	// example:
	//
	// ApplicationType
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetInventorySchemaResponseBodySchemasAttributes) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySchemaResponseBodySchemasAttributes) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponseBodySchemasAttributes) GetDataType() *string {
	return s.DataType
}

func (s *GetInventorySchemaResponseBodySchemasAttributes) GetName() *string {
	return s.Name
}

func (s *GetInventorySchemaResponseBodySchemasAttributes) SetDataType(v string) *GetInventorySchemaResponseBodySchemasAttributes {
	s.DataType = &v
	return s
}

func (s *GetInventorySchemaResponseBodySchemasAttributes) SetName(v string) *GetInventorySchemaResponseBodySchemasAttributes {
	s.Name = &v
	return s
}

func (s *GetInventorySchemaResponseBodySchemasAttributes) Validate() error {
	return dara.Validate(s)
}
