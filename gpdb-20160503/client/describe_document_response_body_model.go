// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChunkFileUrl(v string) *DescribeDocumentResponseBody
	GetChunkFileUrl() *string
	SetDocsCount(v int32) *DescribeDocumentResponseBody
	GetDocsCount() *int32
	SetDocumentLoader(v string) *DescribeDocumentResponseBody
	GetDocumentLoader() *string
	SetDocumentLoaderResultFileUrl(v string) *DescribeDocumentResponseBody
	GetDocumentLoaderResultFileUrl() *string
	SetFileExt(v string) *DescribeDocumentResponseBody
	GetFileExt() *string
	SetFileMd5(v string) *DescribeDocumentResponseBody
	GetFileMd5() *string
	SetFileMtime(v string) *DescribeDocumentResponseBody
	GetFileMtime() *string
	SetFileName(v string) *DescribeDocumentResponseBody
	GetFileName() *string
	SetFileSize(v int64) *DescribeDocumentResponseBody
	GetFileSize() *int64
	SetFileUrl(v string) *DescribeDocumentResponseBody
	GetFileUrl() *string
	SetFileVersion(v int32) *DescribeDocumentResponseBody
	GetFileVersion() *int32
	SetMessage(v string) *DescribeDocumentResponseBody
	GetMessage() *string
	SetPlainChunkFileUrl(v string) *DescribeDocumentResponseBody
	GetPlainChunkFileUrl() *string
	SetRequestId(v string) *DescribeDocumentResponseBody
	GetRequestId() *string
	SetSource(v string) *DescribeDocumentResponseBody
	GetSource() *string
	SetStatus(v string) *DescribeDocumentResponseBody
	GetStatus() *string
	SetTextSplitter(v string) *DescribeDocumentResponseBody
	GetTextSplitter() *string
}

type DescribeDocumentResponseBody struct {
	// URL of the split file, valid for 2 hours. The file format is JSONL, with each line formatted as `{"page_content":"*****", "metadata": {"**":"***","**":"***"}`.
	//
	// example:
	//
	// http://oss.xxx/music_chunk.jsonl
	ChunkFileUrl *string `json:"ChunkFileUrl,omitempty" xml:"ChunkFileUrl,omitempty"`
	// Number of documents after splitting.
	//
	// example:
	//
	// 100
	DocsCount *int32 `json:"DocsCount,omitempty" xml:"DocsCount,omitempty"`
	// Name of the document loader.
	//
	// example:
	//
	// RapidOCRPDFLoader
	DocumentLoader              *string `json:"DocumentLoader,omitempty" xml:"DocumentLoader,omitempty"`
	DocumentLoaderResultFileUrl *string `json:"DocumentLoaderResultFileUrl,omitempty" xml:"DocumentLoaderResultFileUrl,omitempty"`
	// File extension.
	//
	// example:
	//
	// txt
	FileExt *string `json:"FileExt,omitempty" xml:"FileExt,omitempty"`
	// MD5 value of the file.
	//
	// example:
	//
	// b8078c9591413550f8963e37e24abcea
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// The last modified time of the document.
	//
	// example:
	//
	// 2023-11-01 10:01:01.123456
	FileMtime *string `json:"FileMtime,omitempty" xml:"FileMtime,omitempty"`
	// File name.
	//
	// example:
	//
	// music.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// File size, in bytes.
	//
	// example:
	//
	// 10000
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// Download URL of the document, valid for 2 hours.
	//
	// example:
	//
	// http://oss.xxx/music.txt
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Document version. This value increments by 1 each time the same document is updated and uploaded.
	//
	// example:
	//
	// 1
	FileVersion *int32 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// Detailed information returned by the API.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Download URL for the plain text (without metadata) after splitting, each line is a chunk, valid for 2 hours.
	//
	// example:
	//
	// http://oss.xxx/music_plain_chunk.txt
	PlainChunkFileUrl *string `json:"PlainChunkFileUrl,omitempty" xml:"PlainChunkFileUrl,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Source of the document.
	//
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// API execution status, with values as follows:
	//
	// - **success**: Execution succeeded.
	//
	// - **fail**: Execution failed.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Name of the text splitter.
	//
	// example:
	//
	// ChineseRecursiveTextSplitter
	TextSplitter *string `json:"TextSplitter,omitempty" xml:"TextSplitter,omitempty"`
}

func (s DescribeDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocumentResponseBody) GetChunkFileUrl() *string {
	return s.ChunkFileUrl
}

func (s *DescribeDocumentResponseBody) GetDocsCount() *int32 {
	return s.DocsCount
}

func (s *DescribeDocumentResponseBody) GetDocumentLoader() *string {
	return s.DocumentLoader
}

func (s *DescribeDocumentResponseBody) GetDocumentLoaderResultFileUrl() *string {
	return s.DocumentLoaderResultFileUrl
}

func (s *DescribeDocumentResponseBody) GetFileExt() *string {
	return s.FileExt
}

func (s *DescribeDocumentResponseBody) GetFileMd5() *string {
	return s.FileMd5
}

func (s *DescribeDocumentResponseBody) GetFileMtime() *string {
	return s.FileMtime
}

func (s *DescribeDocumentResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *DescribeDocumentResponseBody) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeDocumentResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *DescribeDocumentResponseBody) GetFileVersion() *int32 {
	return s.FileVersion
}

func (s *DescribeDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDocumentResponseBody) GetPlainChunkFileUrl() *string {
	return s.PlainChunkFileUrl
}

func (s *DescribeDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDocumentResponseBody) GetSource() *string {
	return s.Source
}

func (s *DescribeDocumentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDocumentResponseBody) GetTextSplitter() *string {
	return s.TextSplitter
}

func (s *DescribeDocumentResponseBody) SetChunkFileUrl(v string) *DescribeDocumentResponseBody {
	s.ChunkFileUrl = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetDocsCount(v int32) *DescribeDocumentResponseBody {
	s.DocsCount = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetDocumentLoader(v string) *DescribeDocumentResponseBody {
	s.DocumentLoader = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetDocumentLoaderResultFileUrl(v string) *DescribeDocumentResponseBody {
	s.DocumentLoaderResultFileUrl = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetFileExt(v string) *DescribeDocumentResponseBody {
	s.FileExt = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetFileMd5(v string) *DescribeDocumentResponseBody {
	s.FileMd5 = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetFileMtime(v string) *DescribeDocumentResponseBody {
	s.FileMtime = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetFileName(v string) *DescribeDocumentResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetFileSize(v int64) *DescribeDocumentResponseBody {
	s.FileSize = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetFileUrl(v string) *DescribeDocumentResponseBody {
	s.FileUrl = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetFileVersion(v int32) *DescribeDocumentResponseBody {
	s.FileVersion = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetMessage(v string) *DescribeDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetPlainChunkFileUrl(v string) *DescribeDocumentResponseBody {
	s.PlainChunkFileUrl = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetRequestId(v string) *DescribeDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetSource(v string) *DescribeDocumentResponseBody {
	s.Source = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetStatus(v string) *DescribeDocumentResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetTextSplitter(v string) *DescribeDocumentResponseBody {
	s.TextSplitter = &v
	return s
}

func (s *DescribeDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
