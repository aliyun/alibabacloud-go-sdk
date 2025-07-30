// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIndexAddDocumentsJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryIdsShrink(v string) *SubmitIndexAddDocumentsJobShrinkRequest
	GetCategoryIdsShrink() *string
	SetChunkMode(v string) *SubmitIndexAddDocumentsJobShrinkRequest
	GetChunkMode() *string
	SetChunkSize(v int32) *SubmitIndexAddDocumentsJobShrinkRequest
	GetChunkSize() *int32
	SetDocumentIdsShrink(v string) *SubmitIndexAddDocumentsJobShrinkRequest
	GetDocumentIdsShrink() *string
	SetEnableHeaders(v bool) *SubmitIndexAddDocumentsJobShrinkRequest
	GetEnableHeaders() *bool
	SetIndexId(v string) *SubmitIndexAddDocumentsJobShrinkRequest
	GetIndexId() *string
	SetOverlapSize(v int32) *SubmitIndexAddDocumentsJobShrinkRequest
	GetOverlapSize() *int32
	SetSeparator(v string) *SubmitIndexAddDocumentsJobShrinkRequest
	GetSeparator() *string
	SetSourceType(v string) *SubmitIndexAddDocumentsJobShrinkRequest
	GetSourceType() *string
}

type SubmitIndexAddDocumentsJobShrinkRequest struct {
	// The list of primary key IDs of the category.
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	ChunkMode         *string `json:"ChunkMode,omitempty" xml:"ChunkMode,omitempty"`
	ChunkSize         *int32  `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// The list of the primary key IDs of the documents.
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
	EnableHeaders     *bool   `json:"EnableHeaders,omitempty" xml:"EnableHeaders,omitempty"`
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

func (s SubmitIndexAddDocumentsJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) GetCategoryIdsShrink() *string {
	return s.CategoryIdsShrink
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) GetChunkMode() *string {
	return s.ChunkMode
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) GetDocumentIdsShrink() *string {
	return s.DocumentIdsShrink
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) GetEnableHeaders() *bool {
	return s.EnableHeaders
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) GetOverlapSize() *int32 {
	return s.OverlapSize
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) GetSeparator() *string {
	return s.Separator
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetCategoryIdsShrink(v string) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.CategoryIdsShrink = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetChunkMode(v string) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.ChunkMode = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetChunkSize(v int32) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.ChunkSize = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetDocumentIdsShrink(v string) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.DocumentIdsShrink = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetEnableHeaders(v bool) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.EnableHeaders = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetIndexId(v string) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.IndexId = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetOverlapSize(v int32) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.OverlapSize = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetSeparator(v string) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.Separator = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetSourceType(v string) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
