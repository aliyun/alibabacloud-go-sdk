// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateDocRequest
	GetAgentKey() *string
	SetCategoryId(v int64) *UpdateDocRequest
	GetCategoryId() *int64
	SetConfig(v string) *UpdateDocRequest
	GetConfig() *string
	SetContent(v string) *UpdateDocRequest
	GetContent() *string
	SetDocMetadata(v []*UpdateDocRequestDocMetadata) *UpdateDocRequest
	GetDocMetadata() []*UpdateDocRequestDocMetadata
	SetDocName(v string) *UpdateDocRequest
	GetDocName() *string
	SetEndDate(v string) *UpdateDocRequest
	GetEndDate() *string
	SetKnowledgeId(v int64) *UpdateDocRequest
	GetKnowledgeId() *int64
	SetMeta(v string) *UpdateDocRequest
	GetMeta() *string
	SetStartDate(v string) *UpdateDocRequest
	GetStartDate() *string
	SetTagIds(v []*int64) *UpdateDocRequest
	GetTagIds() []*int64
	SetTitle(v string) *UpdateDocRequest
	GetTitle() *string
}

type UpdateDocRequest struct {
	// The business space key. If this parameter is not specified, the default business space is accessed. You can obtain the key value on the business management page of the primary account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The document category ID.
	//
	// example:
	//
	// 231001028593
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The document parsing/chunking configuration.
	//
	// Configuration item: Splitter
	//
	// KEY: Splitter
	//
	// Valid values: paragraphSplitter - hierarchy recognition (default)
	//
	// Valid values: treeSplitter - rule-based hierarchy
	//
	// Configuration item: Document chunk size
	//
	// KEY: ChunkSize
	//
	// VALUE: Default 500. Range: [200,800]
	//
	// Configuration item: Rule-based hierarchy
	//
	// KEY: TreePatterns
	//
	// VALUE: Default []
	//
	// Configuration item: Document title source
	//
	// KEY: TitleSource
	//
	// Valid values: ocrTitle - OCR-recognized title (default)
	//
	// Valid values: docName - document name as title.
	//
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The document content.
	//
	// example:
	//
	// Document content
	Content     *string                        `json:"Content,omitempty" xml:"Content,omitempty"`
	DocMetadata []*UpdateDocRequestDocMetadata `json:"DocMetadata,omitempty" xml:"DocMetadata,omitempty" type:"Repeated"`
	// The document name.
	//
	// example:
	//
	// Document QA Module Feature Description
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// The end of the effective period.
	//
	// example:
	//
	// 2023-03-11T23:59:59Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The document ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The document metadata.
	//
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The start of the effective period.
	//
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The list of tag IDs. If this parameter is left empty, the tag bindings are removed.
	TagIds []*int64 `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The document title.
	//
	// example:
	//
	// Document QA Feature Description
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateDocRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateDocRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateDocRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateDocRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateDocRequest) GetDocMetadata() []*UpdateDocRequestDocMetadata {
	return s.DocMetadata
}

func (s *UpdateDocRequest) GetDocName() *string {
	return s.DocName
}

func (s *UpdateDocRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateDocRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *UpdateDocRequest) GetMeta() *string {
	return s.Meta
}

func (s *UpdateDocRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateDocRequest) GetTagIds() []*int64 {
	return s.TagIds
}

func (s *UpdateDocRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateDocRequest) SetAgentKey(v string) *UpdateDocRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDocRequest) SetCategoryId(v int64) *UpdateDocRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateDocRequest) SetConfig(v string) *UpdateDocRequest {
	s.Config = &v
	return s
}

func (s *UpdateDocRequest) SetContent(v string) *UpdateDocRequest {
	s.Content = &v
	return s
}

func (s *UpdateDocRequest) SetDocMetadata(v []*UpdateDocRequestDocMetadata) *UpdateDocRequest {
	s.DocMetadata = v
	return s
}

func (s *UpdateDocRequest) SetDocName(v string) *UpdateDocRequest {
	s.DocName = &v
	return s
}

func (s *UpdateDocRequest) SetEndDate(v string) *UpdateDocRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateDocRequest) SetKnowledgeId(v int64) *UpdateDocRequest {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateDocRequest) SetMeta(v string) *UpdateDocRequest {
	s.Meta = &v
	return s
}

func (s *UpdateDocRequest) SetStartDate(v string) *UpdateDocRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateDocRequest) SetTagIds(v []*int64) *UpdateDocRequest {
	s.TagIds = v
	return s
}

func (s *UpdateDocRequest) SetTitle(v string) *UpdateDocRequest {
	s.Title = &v
	return s
}

func (s *UpdateDocRequest) Validate() error {
	if s.DocMetadata != nil {
		for _, item := range s.DocMetadata {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDocRequestDocMetadata struct {
	BusinessViewId      *string                                           `json:"BusinessViewId,omitempty" xml:"BusinessViewId,omitempty"`
	BusinessViewName    *string                                           `json:"BusinessViewName,omitempty" xml:"BusinessViewName,omitempty"`
	MetaCellInfoDTOList []*UpdateDocRequestDocMetadataMetaCellInfoDTOList `json:"MetaCellInfoDTOList,omitempty" xml:"MetaCellInfoDTOList,omitempty" type:"Repeated"`
}

func (s UpdateDocRequestDocMetadata) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocRequestDocMetadata) GoString() string {
	return s.String()
}

func (s *UpdateDocRequestDocMetadata) GetBusinessViewId() *string {
	return s.BusinessViewId
}

func (s *UpdateDocRequestDocMetadata) GetBusinessViewName() *string {
	return s.BusinessViewName
}

func (s *UpdateDocRequestDocMetadata) GetMetaCellInfoDTOList() []*UpdateDocRequestDocMetadataMetaCellInfoDTOList {
	return s.MetaCellInfoDTOList
}

func (s *UpdateDocRequestDocMetadata) SetBusinessViewId(v string) *UpdateDocRequestDocMetadata {
	s.BusinessViewId = &v
	return s
}

func (s *UpdateDocRequestDocMetadata) SetBusinessViewName(v string) *UpdateDocRequestDocMetadata {
	s.BusinessViewName = &v
	return s
}

func (s *UpdateDocRequestDocMetadata) SetMetaCellInfoDTOList(v []*UpdateDocRequestDocMetadataMetaCellInfoDTOList) *UpdateDocRequestDocMetadata {
	s.MetaCellInfoDTOList = v
	return s
}

func (s *UpdateDocRequestDocMetadata) Validate() error {
	if s.MetaCellInfoDTOList != nil {
		for _, item := range s.MetaCellInfoDTOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDocRequestDocMetadataMetaCellInfoDTOList struct {
	FieldCode *string `json:"FieldCode,omitempty" xml:"FieldCode,omitempty"`
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateDocRequestDocMetadataMetaCellInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocRequestDocMetadataMetaCellInfoDTOList) GoString() string {
	return s.String()
}

func (s *UpdateDocRequestDocMetadataMetaCellInfoDTOList) GetFieldCode() *string {
	return s.FieldCode
}

func (s *UpdateDocRequestDocMetadataMetaCellInfoDTOList) GetFieldName() *string {
	return s.FieldName
}

func (s *UpdateDocRequestDocMetadataMetaCellInfoDTOList) GetValue() *string {
	return s.Value
}

func (s *UpdateDocRequestDocMetadataMetaCellInfoDTOList) SetFieldCode(v string) *UpdateDocRequestDocMetadataMetaCellInfoDTOList {
	s.FieldCode = &v
	return s
}

func (s *UpdateDocRequestDocMetadataMetaCellInfoDTOList) SetFieldName(v string) *UpdateDocRequestDocMetadataMetaCellInfoDTOList {
	s.FieldName = &v
	return s
}

func (s *UpdateDocRequestDocMetadataMetaCellInfoDTOList) SetValue(v string) *UpdateDocRequestDocMetadataMetaCellInfoDTOList {
	s.Value = &v
	return s
}

func (s *UpdateDocRequestDocMetadataMetaCellInfoDTOList) Validate() error {
	return dara.Validate(s)
}
