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
	// This parameter is required.
	CatalogUuid  *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	DocFormat *string `json:"DocFormat,omitempty" xml:"DocFormat,omitempty"`
	// This parameter is required.
	Document *string `json:"Document,omitempty" xml:"Document,omitempty"`
	// This parameter is required.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Tag    *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
