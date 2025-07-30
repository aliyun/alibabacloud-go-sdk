// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIndexAddDocumentsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryIds(v []*string) *SubmitIndexAddDocumentsJobRequest
	GetCategoryIds() []*string
	SetChunkMode(v string) *SubmitIndexAddDocumentsJobRequest
	GetChunkMode() *string
	SetChunkSize(v int32) *SubmitIndexAddDocumentsJobRequest
	GetChunkSize() *int32
	SetDocumentIds(v []*string) *SubmitIndexAddDocumentsJobRequest
	GetDocumentIds() []*string
	SetEnableHeaders(v bool) *SubmitIndexAddDocumentsJobRequest
	GetEnableHeaders() *bool
	SetIndexId(v string) *SubmitIndexAddDocumentsJobRequest
	GetIndexId() *string
	SetOverlapSize(v int32) *SubmitIndexAddDocumentsJobRequest
	GetOverlapSize() *int32
	SetSeparator(v string) *SubmitIndexAddDocumentsJobRequest
	GetSeparator() *string
	SetSourceType(v string) *SubmitIndexAddDocumentsJobRequest
	GetSourceType() *string
}

type SubmitIndexAddDocumentsJobRequest struct {
	// The list of primary key IDs of the category.
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	ChunkMode   *string   `json:"ChunkMode,omitempty" xml:"ChunkMode,omitempty"`
	ChunkSize   *int32    `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// The list of the primary key IDs of the documents.
	DocumentIds   []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	EnableHeaders *bool     `json:"EnableHeaders,omitempty" xml:"EnableHeaders,omitempty"`
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId     *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	OverlapSize *int32  `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	Separator   *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The data type of [Data Management](https://bailian.console.aliyun.com/#/data-center). For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values:
	//
	// 	- DATA_CENTER_CATEGORY: The category type. Import all documents from one or more categories in Data Center.
	//
	// 	- DATA_CENTER_FILE: The document type. Import one or more documents from Data Center.
	//
	// >  If this parameter is set to DATA_CENTER_CATEGORY, you must specify the `CategoryIds` parameter. If this parameter is set to DATA_CENTER_FILE, you must specify the `DocumentIds` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s SubmitIndexAddDocumentsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobRequest) GetCategoryIds() []*string {
	return s.CategoryIds
}

func (s *SubmitIndexAddDocumentsJobRequest) GetChunkMode() *string {
	return s.ChunkMode
}

func (s *SubmitIndexAddDocumentsJobRequest) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *SubmitIndexAddDocumentsJobRequest) GetDocumentIds() []*string {
	return s.DocumentIds
}

func (s *SubmitIndexAddDocumentsJobRequest) GetEnableHeaders() *bool {
	return s.EnableHeaders
}

func (s *SubmitIndexAddDocumentsJobRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *SubmitIndexAddDocumentsJobRequest) GetOverlapSize() *int32 {
	return s.OverlapSize
}

func (s *SubmitIndexAddDocumentsJobRequest) GetSeparator() *string {
	return s.Separator
}

func (s *SubmitIndexAddDocumentsJobRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *SubmitIndexAddDocumentsJobRequest) SetCategoryIds(v []*string) *SubmitIndexAddDocumentsJobRequest {
	s.CategoryIds = v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetChunkMode(v string) *SubmitIndexAddDocumentsJobRequest {
	s.ChunkMode = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetChunkSize(v int32) *SubmitIndexAddDocumentsJobRequest {
	s.ChunkSize = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetDocumentIds(v []*string) *SubmitIndexAddDocumentsJobRequest {
	s.DocumentIds = v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetEnableHeaders(v bool) *SubmitIndexAddDocumentsJobRequest {
	s.EnableHeaders = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetIndexId(v string) *SubmitIndexAddDocumentsJobRequest {
	s.IndexId = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetOverlapSize(v int32) *SubmitIndexAddDocumentsJobRequest {
	s.OverlapSize = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetSeparator(v string) *SubmitIndexAddDocumentsJobRequest {
	s.Separator = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetSourceType(v string) *SubmitIndexAddDocumentsJobRequest {
	s.SourceType = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) Validate() error {
	return dara.Validate(s)
}
