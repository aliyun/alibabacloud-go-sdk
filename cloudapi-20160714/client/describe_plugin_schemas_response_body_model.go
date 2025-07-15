// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPluginSchemas(v *DescribePluginSchemasResponseBodyPluginSchemas) *DescribePluginSchemasResponseBody
	GetPluginSchemas() *DescribePluginSchemasResponseBodyPluginSchemas
	SetRequestId(v string) *DescribePluginSchemasResponseBody
	GetRequestId() *string
}

type DescribePluginSchemasResponseBody struct {
	PluginSchemas *DescribePluginSchemasResponseBodyPluginSchemas `json:"PluginSchemas,omitempty" xml:"PluginSchemas,omitempty" type:"Struct"`
	// example:
	//
	// 2D39D1B3-8548-508A-9CE2-7F4A3F2A7989
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePluginSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginSchemasResponseBody) GetPluginSchemas() *DescribePluginSchemasResponseBodyPluginSchemas {
	return s.PluginSchemas
}

func (s *DescribePluginSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePluginSchemasResponseBody) SetPluginSchemas(v *DescribePluginSchemasResponseBodyPluginSchemas) *DescribePluginSchemasResponseBody {
	s.PluginSchemas = v
	return s
}

func (s *DescribePluginSchemasResponseBody) SetRequestId(v string) *DescribePluginSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginSchemasResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePluginSchemasResponseBodyPluginSchemas struct {
	PluginSchema []*DescribePluginSchemasResponseBodyPluginSchemasPluginSchema `json:"PluginSchema,omitempty" xml:"PluginSchema,omitempty" type:"Repeated"`
}

func (s DescribePluginSchemasResponseBodyPluginSchemas) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginSchemasResponseBodyPluginSchemas) GoString() string {
	return s.String()
}

func (s *DescribePluginSchemasResponseBodyPluginSchemas) GetPluginSchema() []*DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	return s.PluginSchema
}

func (s *DescribePluginSchemasResponseBodyPluginSchemas) SetPluginSchema(v []*DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) *DescribePluginSchemasResponseBodyPluginSchemas {
	s.PluginSchema = v
	return s
}

func (s *DescribePluginSchemasResponseBodyPluginSchemas) Validate() error {
	return dara.Validate(s)
}

type DescribePluginSchemasResponseBodyPluginSchemasPluginSchema struct {
	// example:
	//
	// plugin scheme description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 4107**
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// example:
	//
	// VPC_C
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	SupportClassic *bool `json:"SupportClassic,omitempty" xml:"SupportClassic,omitempty"`
	// example:
	//
	// plugin schema title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) GoString() string {
	return s.String()
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) GetDescription() *string {
	return s.Description
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) GetDocumentId() *string {
	return s.DocumentId
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) GetName() *string {
	return s.Name
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) GetSupportClassic() *bool {
	return s.SupportClassic
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) GetTitle() *string {
	return s.Title
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) SetDescription(v string) *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	s.Description = &v
	return s
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) SetDocumentId(v string) *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	s.DocumentId = &v
	return s
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) SetName(v string) *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	s.Name = &v
	return s
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) SetSupportClassic(v bool) *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	s.SupportClassic = &v
	return s
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) SetTitle(v string) *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	s.Title = &v
	return s
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) Validate() error {
	return dara.Validate(s)
}
