// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocumentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkOverlap(v int64) *UploadDocumentShrinkRequest
	GetChunkOverlap() *int64
	SetChunkSize(v int64) *UploadDocumentShrinkRequest
	GetChunkSize() *int64
	SetDescription(v string) *UploadDocumentShrinkRequest
	GetDescription() *string
	SetDocumentLoaderName(v string) *UploadDocumentShrinkRequest
	GetDocumentLoaderName() *string
	SetFileName(v string) *UploadDocumentShrinkRequest
	GetFileName() *string
	SetKbUuid(v string) *UploadDocumentShrinkRequest
	GetKbUuid() *string
	SetLocation(v string) *UploadDocumentShrinkRequest
	GetLocation() *string
	SetSeparatorsShrink(v string) *UploadDocumentShrinkRequest
	GetSeparatorsShrink() *string
	SetSplitterModel(v string) *UploadDocumentShrinkRequest
	GetSplitterModel() *string
	SetTextSplitterName(v string) *UploadDocumentShrinkRequest
	GetTextSplitterName() *string
	SetVlEnhance(v bool) *UploadDocumentShrinkRequest
	GetVlEnhance() *bool
	SetZhTitleEnhance(v bool) *UploadDocumentShrinkRequest
	GetZhTitleEnhance() *bool
}

type UploadDocumentShrinkRequest struct {
	// example:
	//
	// 50
	ChunkOverlap *int64 `json:"ChunkOverlap,omitempty" xml:"ChunkOverlap,omitempty"`
	// example:
	//
	// 250
	ChunkSize *int64 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ADBPGLoader
	DocumentLoaderName *string `json:"DocumentLoaderName,omitempty" xml:"DocumentLoaderName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// This parameter is required.
	Location         *string `json:"Location,omitempty" xml:"Location,omitempty"`
	SeparatorsShrink *string `json:"Separators,omitempty" xml:"Separators,omitempty"`
	// example:
	//
	// qwen3-8b
	SplitterModel *string `json:"SplitterModel,omitempty" xml:"SplitterModel,omitempty"`
	// example:
	//
	// ChineseRecursiveTextSplitter
	TextSplitterName *string `json:"TextSplitterName,omitempty" xml:"TextSplitterName,omitempty"`
	// example:
	//
	// false
	VlEnhance *bool `json:"VlEnhance,omitempty" xml:"VlEnhance,omitempty"`
	// example:
	//
	// false
	ZhTitleEnhance *bool `json:"ZhTitleEnhance,omitempty" xml:"ZhTitleEnhance,omitempty"`
}

func (s UploadDocumentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentShrinkRequest) GetChunkOverlap() *int64 {
	return s.ChunkOverlap
}

func (s *UploadDocumentShrinkRequest) GetChunkSize() *int64 {
	return s.ChunkSize
}

func (s *UploadDocumentShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UploadDocumentShrinkRequest) GetDocumentLoaderName() *string {
	return s.DocumentLoaderName
}

func (s *UploadDocumentShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadDocumentShrinkRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *UploadDocumentShrinkRequest) GetLocation() *string {
	return s.Location
}

func (s *UploadDocumentShrinkRequest) GetSeparatorsShrink() *string {
	return s.SeparatorsShrink
}

func (s *UploadDocumentShrinkRequest) GetSplitterModel() *string {
	return s.SplitterModel
}

func (s *UploadDocumentShrinkRequest) GetTextSplitterName() *string {
	return s.TextSplitterName
}

func (s *UploadDocumentShrinkRequest) GetVlEnhance() *bool {
	return s.VlEnhance
}

func (s *UploadDocumentShrinkRequest) GetZhTitleEnhance() *bool {
	return s.ZhTitleEnhance
}

func (s *UploadDocumentShrinkRequest) SetChunkOverlap(v int64) *UploadDocumentShrinkRequest {
	s.ChunkOverlap = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetChunkSize(v int64) *UploadDocumentShrinkRequest {
	s.ChunkSize = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetDescription(v string) *UploadDocumentShrinkRequest {
	s.Description = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetDocumentLoaderName(v string) *UploadDocumentShrinkRequest {
	s.DocumentLoaderName = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetFileName(v string) *UploadDocumentShrinkRequest {
	s.FileName = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetKbUuid(v string) *UploadDocumentShrinkRequest {
	s.KbUuid = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetLocation(v string) *UploadDocumentShrinkRequest {
	s.Location = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetSeparatorsShrink(v string) *UploadDocumentShrinkRequest {
	s.SeparatorsShrink = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetSplitterModel(v string) *UploadDocumentShrinkRequest {
	s.SplitterModel = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetTextSplitterName(v string) *UploadDocumentShrinkRequest {
	s.TextSplitterName = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetVlEnhance(v bool) *UploadDocumentShrinkRequest {
	s.VlEnhance = &v
	return s
}

func (s *UploadDocumentShrinkRequest) SetZhTitleEnhance(v bool) *UploadDocumentShrinkRequest {
	s.ZhTitleEnhance = &v
	return s
}

func (s *UploadDocumentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
