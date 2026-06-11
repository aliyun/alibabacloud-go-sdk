// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkOverlap(v int64) *UploadDocumentRequest
	GetChunkOverlap() *int64
	SetChunkSize(v int64) *UploadDocumentRequest
	GetChunkSize() *int64
	SetDescription(v string) *UploadDocumentRequest
	GetDescription() *string
	SetDocumentLoaderName(v string) *UploadDocumentRequest
	GetDocumentLoaderName() *string
	SetFileName(v string) *UploadDocumentRequest
	GetFileName() *string
	SetKbUuid(v string) *UploadDocumentRequest
	GetKbUuid() *string
	SetLocation(v string) *UploadDocumentRequest
	GetLocation() *string
	SetSeparators(v []*string) *UploadDocumentRequest
	GetSeparators() []*string
	SetSplitterModel(v string) *UploadDocumentRequest
	GetSplitterModel() *string
	SetTextSplitterName(v string) *UploadDocumentRequest
	GetTextSplitterName() *string
	SetVlEnhance(v bool) *UploadDocumentRequest
	GetVlEnhance() *bool
	SetZhTitleEnhance(v bool) *UploadDocumentRequest
	GetZhTitleEnhance() *bool
}

type UploadDocumentRequest struct {
	// The number of overlapping characters between adjacent chunks. This value cannot exceed `ChunkSize`. The default is 50.
	//
	// example:
	//
	// 50
	ChunkOverlap *int64 `json:"ChunkOverlap,omitempty" xml:"ChunkOverlap,omitempty"`
	// The size of each document chunk. The default is 250, and the maximum is 2,048.
	//
	// example:
	//
	// 250
	ChunkSize *int64 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// The description of the document.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the document loader. The default is `ADBPGLoader`.
	//
	// example:
	//
	// ADBPGLoader
	DocumentLoaderName *string `json:"DocumentLoaderName,omitempty" xml:"DocumentLoaderName,omitempty"`
	// The name of the document.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// The OSS location of the input file. Construct this path by appending the file name to the `UploadDir` value returned by the `DescribeKnowledgeBaseUploadSignature` operation.
	//
	// This parameter is required.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// An array of strings used to split text.
	//
	// > - This critical parameter affects data chunking results and is related to the splitter specified by `TextSplitterName`.
	//
	// >
	//
	// > - In most cases, you can omit this parameter. The service automatically assigns default separators based on `TextSplitterName`.
	Separators []*string `json:"Separators,omitempty" xml:"Separators,omitempty" type:"Repeated"`
	// The splitter model to use. The default is `qwen3-8b`.
	//
	// example:
	//
	// qwen3-8b
	SplitterModel *string `json:"SplitterModel,omitempty" xml:"SplitterModel,omitempty"`
	// The name of the text splitter.
	//
	// example:
	//
	// ChineseRecursiveTextSplitter
	TextSplitterName *string `json:"TextSplitterName,omitempty" xml:"TextSplitterName,omitempty"`
	// Specifies whether to enable visual-linguistic (VL) enhanced content recognition for complex documents. The default is false.
	//
	// example:
	//
	// false
	VlEnhance *bool `json:"VlEnhance,omitempty" xml:"VlEnhance,omitempty"`
	// Specifies whether to enable title enhancement.
	//
	// example:
	//
	// false
	ZhTitleEnhance *bool `json:"ZhTitleEnhance,omitempty" xml:"ZhTitleEnhance,omitempty"`
}

func (s UploadDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentRequest) GetChunkOverlap() *int64 {
	return s.ChunkOverlap
}

func (s *UploadDocumentRequest) GetChunkSize() *int64 {
	return s.ChunkSize
}

func (s *UploadDocumentRequest) GetDescription() *string {
	return s.Description
}

func (s *UploadDocumentRequest) GetDocumentLoaderName() *string {
	return s.DocumentLoaderName
}

func (s *UploadDocumentRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadDocumentRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *UploadDocumentRequest) GetLocation() *string {
	return s.Location
}

func (s *UploadDocumentRequest) GetSeparators() []*string {
	return s.Separators
}

func (s *UploadDocumentRequest) GetSplitterModel() *string {
	return s.SplitterModel
}

func (s *UploadDocumentRequest) GetTextSplitterName() *string {
	return s.TextSplitterName
}

func (s *UploadDocumentRequest) GetVlEnhance() *bool {
	return s.VlEnhance
}

func (s *UploadDocumentRequest) GetZhTitleEnhance() *bool {
	return s.ZhTitleEnhance
}

func (s *UploadDocumentRequest) SetChunkOverlap(v int64) *UploadDocumentRequest {
	s.ChunkOverlap = &v
	return s
}

func (s *UploadDocumentRequest) SetChunkSize(v int64) *UploadDocumentRequest {
	s.ChunkSize = &v
	return s
}

func (s *UploadDocumentRequest) SetDescription(v string) *UploadDocumentRequest {
	s.Description = &v
	return s
}

func (s *UploadDocumentRequest) SetDocumentLoaderName(v string) *UploadDocumentRequest {
	s.DocumentLoaderName = &v
	return s
}

func (s *UploadDocumentRequest) SetFileName(v string) *UploadDocumentRequest {
	s.FileName = &v
	return s
}

func (s *UploadDocumentRequest) SetKbUuid(v string) *UploadDocumentRequest {
	s.KbUuid = &v
	return s
}

func (s *UploadDocumentRequest) SetLocation(v string) *UploadDocumentRequest {
	s.Location = &v
	return s
}

func (s *UploadDocumentRequest) SetSeparators(v []*string) *UploadDocumentRequest {
	s.Separators = v
	return s
}

func (s *UploadDocumentRequest) SetSplitterModel(v string) *UploadDocumentRequest {
	s.SplitterModel = &v
	return s
}

func (s *UploadDocumentRequest) SetTextSplitterName(v string) *UploadDocumentRequest {
	s.TextSplitterName = &v
	return s
}

func (s *UploadDocumentRequest) SetVlEnhance(v bool) *UploadDocumentRequest {
	s.VlEnhance = &v
	return s
}

func (s *UploadDocumentRequest) SetZhTitleEnhance(v bool) *UploadDocumentRequest {
	s.ZhTitleEnhance = &v
	return s
}

func (s *UploadDocumentRequest) Validate() error {
	return dara.Validate(s)
}
