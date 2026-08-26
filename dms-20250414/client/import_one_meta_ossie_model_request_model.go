// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportOneMetaOssieModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *ImportOneMetaOssieModelRequest
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *ImportOneMetaOssieModelRequest
	GetDatabaseUuid() *string
	SetDescription(v string) *ImportOneMetaOssieModelRequest
	GetDescription() *string
	SetDocFormat(v string) *ImportOneMetaOssieModelRequest
	GetDocFormat() *string
	SetDocument(v string) *ImportOneMetaOssieModelRequest
	GetDocument() *string
	SetSource(v string) *ImportOneMetaOssieModelRequest
	GetSource() *string
	SetTag(v string) *ImportOneMetaOssieModelRequest
	GetTag() *string
	SetTitle(v string) *ImportOneMetaOssieModelRequest
	GetTitle() *string
}

type ImportOneMetaOssieModelRequest struct {
	// The UUID of the associated folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The UUID of the associated database.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The semantic description.
	//
	// example:
	//
	// Order summary
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The semantic model document type. Valid values: JSON and YAML.
	//
	// This parameter is required.
	//
	// example:
	//
	// JSON
	DocFormat *string `json:"DocFormat,omitempty" xml:"DocFormat,omitempty"`
	// The semantic model document definition.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "version": "0.2.0.dev0",
	//
	//   "semantic_model": [
	//
	//     {
	//
	//       "name": "sales",
	//
	//       "datasets": [
	//
	//         {
	//
	//           "name": "orders",
	//
	//           "source": "analytics.public.orders"
	//
	//         }
	//
	//       ]
	//
	//     }
	//
	//   ]
	//
	// }
	Document *string `json:"Document,omitempty" xml:"Document,omitempty"`
	// The source of the semantic model.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA_AGENT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The semantic model tag.
	//
	// example:
	//
	// new_sales
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The semantic title. If the value is empty, the title is extracted from the document.
	//
	// example:
	//
	// Order total
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ImportOneMetaOssieModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportOneMetaOssieModelRequest) GoString() string {
	return s.String()
}

func (s *ImportOneMetaOssieModelRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *ImportOneMetaOssieModelRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *ImportOneMetaOssieModelRequest) GetDescription() *string {
	return s.Description
}

func (s *ImportOneMetaOssieModelRequest) GetDocFormat() *string {
	return s.DocFormat
}

func (s *ImportOneMetaOssieModelRequest) GetDocument() *string {
	return s.Document
}

func (s *ImportOneMetaOssieModelRequest) GetSource() *string {
	return s.Source
}

func (s *ImportOneMetaOssieModelRequest) GetTag() *string {
	return s.Tag
}

func (s *ImportOneMetaOssieModelRequest) GetTitle() *string {
	return s.Title
}

func (s *ImportOneMetaOssieModelRequest) SetCatalogUuid(v string) *ImportOneMetaOssieModelRequest {
	s.CatalogUuid = &v
	return s
}

func (s *ImportOneMetaOssieModelRequest) SetDatabaseUuid(v string) *ImportOneMetaOssieModelRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *ImportOneMetaOssieModelRequest) SetDescription(v string) *ImportOneMetaOssieModelRequest {
	s.Description = &v
	return s
}

func (s *ImportOneMetaOssieModelRequest) SetDocFormat(v string) *ImportOneMetaOssieModelRequest {
	s.DocFormat = &v
	return s
}

func (s *ImportOneMetaOssieModelRequest) SetDocument(v string) *ImportOneMetaOssieModelRequest {
	s.Document = &v
	return s
}

func (s *ImportOneMetaOssieModelRequest) SetSource(v string) *ImportOneMetaOssieModelRequest {
	s.Source = &v
	return s
}

func (s *ImportOneMetaOssieModelRequest) SetTag(v string) *ImportOneMetaOssieModelRequest {
	s.Tag = &v
	return s
}

func (s *ImportOneMetaOssieModelRequest) SetTitle(v string) *ImportOneMetaOssieModelRequest {
	s.Title = &v
	return s
}

func (s *ImportOneMetaOssieModelRequest) Validate() error {
	return dara.Validate(s)
}
