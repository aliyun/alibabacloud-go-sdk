// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOneMetaOssieModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *UpdateOneMetaOssieModelRequest
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *UpdateOneMetaOssieModelRequest
	GetDatabaseUuid() *string
	SetDescription(v string) *UpdateOneMetaOssieModelRequest
	GetDescription() *string
	SetDocFormat(v string) *UpdateOneMetaOssieModelRequest
	GetDocFormat() *string
	SetDocument(v string) *UpdateOneMetaOssieModelRequest
	GetDocument() *string
	SetKnowledgeUuid(v string) *UpdateOneMetaOssieModelRequest
	GetKnowledgeUuid() *string
	SetTag(v string) *UpdateOneMetaOssieModelRequest
	GetTag() *string
	SetTitle(v string) *UpdateOneMetaOssieModelRequest
	GetTitle() *string
}

type UpdateOneMetaOssieModelRequest struct {
	CatalogUuid  *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocFormat    *string `json:"DocFormat,omitempty" xml:"DocFormat,omitempty"`
	Document     *string `json:"Document,omitempty" xml:"Document,omitempty"`
	// This parameter is required.
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	Tag           *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateOneMetaOssieModelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOneMetaOssieModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateOneMetaOssieModelRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *UpdateOneMetaOssieModelRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *UpdateOneMetaOssieModelRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOneMetaOssieModelRequest) GetDocFormat() *string {
	return s.DocFormat
}

func (s *UpdateOneMetaOssieModelRequest) GetDocument() *string {
	return s.Document
}

func (s *UpdateOneMetaOssieModelRequest) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *UpdateOneMetaOssieModelRequest) GetTag() *string {
	return s.Tag
}

func (s *UpdateOneMetaOssieModelRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateOneMetaOssieModelRequest) SetCatalogUuid(v string) *UpdateOneMetaOssieModelRequest {
	s.CatalogUuid = &v
	return s
}

func (s *UpdateOneMetaOssieModelRequest) SetDatabaseUuid(v string) *UpdateOneMetaOssieModelRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *UpdateOneMetaOssieModelRequest) SetDescription(v string) *UpdateOneMetaOssieModelRequest {
	s.Description = &v
	return s
}

func (s *UpdateOneMetaOssieModelRequest) SetDocFormat(v string) *UpdateOneMetaOssieModelRequest {
	s.DocFormat = &v
	return s
}

func (s *UpdateOneMetaOssieModelRequest) SetDocument(v string) *UpdateOneMetaOssieModelRequest {
	s.Document = &v
	return s
}

func (s *UpdateOneMetaOssieModelRequest) SetKnowledgeUuid(v string) *UpdateOneMetaOssieModelRequest {
	s.KnowledgeUuid = &v
	return s
}

func (s *UpdateOneMetaOssieModelRequest) SetTag(v string) *UpdateOneMetaOssieModelRequest {
	s.Tag = &v
	return s
}

func (s *UpdateOneMetaOssieModelRequest) SetTitle(v string) *UpdateOneMetaOssieModelRequest {
	s.Title = &v
	return s
}

func (s *UpdateOneMetaOssieModelRequest) Validate() error {
	return dara.Validate(s)
}
