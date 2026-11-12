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
	// The URL of the chunked file. The URL is valid for 2 hours. The file is in JSONL format, and each line is in the following format: {"page_content":"*****", "metadata": {"**":"***","**":"***"}
	//
	// example:
	//
	// http://oss.xxx/music_chunk.jsonl
	ChunkFileUrl *string `json:"ChunkFileUrl,omitempty" xml:"ChunkFileUrl,omitempty"`
	// The number of chunks generated from the document.
	//
	// example:
	//
	// 100
	DocsCount *int32 `json:"DocsCount,omitempty" xml:"DocsCount,omitempty"`
	// The name of the document loader.
	//
	// example:
	//
	// RapidOCRPDFLoader
	DocumentLoader *string `json:"DocumentLoader,omitempty" xml:"DocumentLoader,omitempty"`
	// The URL of the Markdown result file parsed by ADBPGLoader. The URL is valid for 2 hours.
	//
	// example:
	//
	// http://oss.xxx/adbpg_loader_result.md
	DocumentLoaderResultFileUrl *string `json:"DocumentLoaderResultFileUrl,omitempty" xml:"DocumentLoaderResultFileUrl,omitempty"`
	// The file name extension.
	//
	// example:
	//
	// txt
	FileExt *string `json:"FileExt,omitempty" xml:"FileExt,omitempty"`
	// The MD5 hash of the file.
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
	// The file name.
	//
	// example:
	//
	// music.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file size in bytes.
	//
	// example:
	//
	// 10000
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The download URL of the document. The URL is valid for 2 hours.
	//
	// example:
	//
	// http://oss.xxx/music.txt
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The document version. This value is incremented by 1 each time the same document is updated and uploaded.
	//
	// example:
	//
	// 1
	FileVersion *int32 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// The detailed information returned by the API operation.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The download URL of the plain text chunks without metadata. Each line represents a chunk. The URL is valid for 2 hours.
	//
	// example:
	//
	// http://oss.xxx/music_plain_chunk.txt
	PlainChunkFileUrl *string `json:"PlainChunkFileUrl,omitempty" xml:"PlainChunkFileUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source of the document.
	//
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The execution status of the API operation. Valid values:
	//
	// - **success**: Execution succeeded.
	//
	// - **fail**: Execution failed.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the text splitter.
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
