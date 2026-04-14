// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaKnowledgeBaseDocument interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *OneMetaKnowledgeBaseDocument
	GetDescription() *string
	SetDocsCount(v int32) *OneMetaKnowledgeBaseDocument
	GetDocsCount() *int32
	SetDocumentLoaderName(v string) *OneMetaKnowledgeBaseDocument
	GetDocumentLoaderName() *string
	SetFileExt(v string) *OneMetaKnowledgeBaseDocument
	GetFileExt() *string
	SetFileSize(v int64) *OneMetaKnowledgeBaseDocument
	GetFileSize() *int64
	SetGmtCreate(v string) *OneMetaKnowledgeBaseDocument
	GetGmtCreate() *string
	SetGmtModified(v string) *OneMetaKnowledgeBaseDocument
	GetGmtModified() *string
	SetKbUuid(v string) *OneMetaKnowledgeBaseDocument
	GetKbUuid() *string
	SetKeywords(v string) *OneMetaKnowledgeBaseDocument
	GetKeywords() *string
	SetName(v string) *OneMetaKnowledgeBaseDocument
	GetName() *string
	SetState(v int32) *OneMetaKnowledgeBaseDocument
	GetState() *int32
	SetSummary(v string) *OneMetaKnowledgeBaseDocument
	GetSummary() *string
	SetTextSplitterName(v string) *OneMetaKnowledgeBaseDocument
	GetTextSplitterName() *string
}

type OneMetaKnowledgeBaseDocument struct {
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocsCount          *int32  `json:"DocsCount,omitempty" xml:"DocsCount,omitempty"`
	DocumentLoaderName *string `json:"DocumentLoaderName,omitempty" xml:"DocumentLoaderName,omitempty"`
	FileExt            *string `json:"FileExt,omitempty" xml:"FileExt,omitempty"`
	FileSize           *int64  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	KbUuid             *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	Keywords           *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	State              *int32  `json:"State,omitempty" xml:"State,omitempty"`
	Summary            *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextSplitterName   *string `json:"TextSplitterName,omitempty" xml:"TextSplitterName,omitempty"`
}

func (s OneMetaKnowledgeBaseDocument) String() string {
	return dara.Prettify(s)
}

func (s OneMetaKnowledgeBaseDocument) GoString() string {
	return s.String()
}

func (s *OneMetaKnowledgeBaseDocument) GetDescription() *string {
	return s.Description
}

func (s *OneMetaKnowledgeBaseDocument) GetDocsCount() *int32 {
	return s.DocsCount
}

func (s *OneMetaKnowledgeBaseDocument) GetDocumentLoaderName() *string {
	return s.DocumentLoaderName
}

func (s *OneMetaKnowledgeBaseDocument) GetFileExt() *string {
	return s.FileExt
}

func (s *OneMetaKnowledgeBaseDocument) GetFileSize() *int64 {
	return s.FileSize
}

func (s *OneMetaKnowledgeBaseDocument) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *OneMetaKnowledgeBaseDocument) GetGmtModified() *string {
	return s.GmtModified
}

func (s *OneMetaKnowledgeBaseDocument) GetKbUuid() *string {
	return s.KbUuid
}

func (s *OneMetaKnowledgeBaseDocument) GetKeywords() *string {
	return s.Keywords
}

func (s *OneMetaKnowledgeBaseDocument) GetName() *string {
	return s.Name
}

func (s *OneMetaKnowledgeBaseDocument) GetState() *int32 {
	return s.State
}

func (s *OneMetaKnowledgeBaseDocument) GetSummary() *string {
	return s.Summary
}

func (s *OneMetaKnowledgeBaseDocument) GetTextSplitterName() *string {
	return s.TextSplitterName
}

func (s *OneMetaKnowledgeBaseDocument) SetDescription(v string) *OneMetaKnowledgeBaseDocument {
	s.Description = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetDocsCount(v int32) *OneMetaKnowledgeBaseDocument {
	s.DocsCount = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetDocumentLoaderName(v string) *OneMetaKnowledgeBaseDocument {
	s.DocumentLoaderName = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetFileExt(v string) *OneMetaKnowledgeBaseDocument {
	s.FileExt = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetFileSize(v int64) *OneMetaKnowledgeBaseDocument {
	s.FileSize = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetGmtCreate(v string) *OneMetaKnowledgeBaseDocument {
	s.GmtCreate = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetGmtModified(v string) *OneMetaKnowledgeBaseDocument {
	s.GmtModified = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetKbUuid(v string) *OneMetaKnowledgeBaseDocument {
	s.KbUuid = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetKeywords(v string) *OneMetaKnowledgeBaseDocument {
	s.Keywords = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetName(v string) *OneMetaKnowledgeBaseDocument {
	s.Name = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetState(v int32) *OneMetaKnowledgeBaseDocument {
	s.State = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetSummary(v string) *OneMetaKnowledgeBaseDocument {
	s.Summary = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) SetTextSplitterName(v string) *OneMetaKnowledgeBaseDocument {
	s.TextSplitterName = &v
	return s
}

func (s *OneMetaKnowledgeBaseDocument) Validate() error {
	return dara.Validate(s)
}
