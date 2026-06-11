// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDocumentResponseBodyData) *DescribeDocumentResponseBody
	GetData() *DescribeDocumentResponseBodyData
	SetErrorCode(v string) *DescribeDocumentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeDocumentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDocumentResponseBody
	GetSuccess() *bool
}

type DescribeDocumentResponseBody struct {
	// The details of the document.
	Data *DescribeDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// KnowledgeBaseNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// Resource not found kb-***
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique request ID. Provide this ID for troubleshooting if an error occurs.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Valid values:
	//
	// - **true**: The request succeeded.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocumentResponseBody) GetData() *DescribeDocumentResponseBodyData {
	return s.Data
}

func (s *DescribeDocumentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDocumentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDocumentResponseBody) SetData(v *DescribeDocumentResponseBodyData) *DescribeDocumentResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDocumentResponseBody) SetErrorCode(v string) *DescribeDocumentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetErrorMessage(v string) *DescribeDocumentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetRequestId(v string) *DescribeDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDocumentResponseBody) SetSuccess(v bool) *DescribeDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDocumentResponseBodyData struct {
	// The description of the document.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of chunks.
	//
	// example:
	//
	// 123
	DocsCount *int64 `json:"DocsCount,omitempty" xml:"DocsCount,omitempty"`
	// The name of the document loader.
	//
	// example:
	//
	// ADBPGLoader
	DocumentLoaderName *string `json:"DocumentLoaderName,omitempty" xml:"DocumentLoaderName,omitempty"`
	// The file extension of the document.
	//
	// example:
	//
	// md
	FileExt *string `json:"FileExt,omitempty" xml:"FileExt,omitempty"`
	// The size of the document in bytes.
	//
	// example:
	//
	// 20307
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The creation time of the document, in UTC.
	//
	// example:
	//
	// 2026-04-22 22:59:35
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time of the document, in UTC.
	//
	// example:
	//
	// 2026-04-24 21:22:53
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the knowledge base.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// The keywords of the document.
	//
	// example:
	//
	// ["test","abc"]
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// The name of the document.
	//
	// example:
	//
	// test.md
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The document state. Possible values are:
	//
	// - **0**: Parsing complete.
	//
	// - **-1**: Not parsed.
	//
	// - **-2**: Parsing in progress.
	//
	// - **-3**: Parsing failed.
	//
	// - **-4**: Parsing canceled.
	//
	// example:
	//
	// 0
	State *int64 `json:"State,omitempty" xml:"State,omitempty"`
	// The summary of the document.
	//
	// example:
	//
	// This is a test document.
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The name of the text splitter.
	//
	// example:
	//
	// ChineseRecursiveTextSplitter
	TextSplitterName *string `json:"TextSplitterName,omitempty" xml:"TextSplitterName,omitempty"`
}

func (s DescribeDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDocumentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeDocumentResponseBodyData) GetDocsCount() *int64 {
	return s.DocsCount
}

func (s *DescribeDocumentResponseBodyData) GetDocumentLoaderName() *string {
	return s.DocumentLoaderName
}

func (s *DescribeDocumentResponseBodyData) GetFileExt() *string {
	return s.FileExt
}

func (s *DescribeDocumentResponseBodyData) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeDocumentResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeDocumentResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDocumentResponseBodyData) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DescribeDocumentResponseBodyData) GetKeywords() *string {
	return s.Keywords
}

func (s *DescribeDocumentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeDocumentResponseBodyData) GetState() *int64 {
	return s.State
}

func (s *DescribeDocumentResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *DescribeDocumentResponseBodyData) GetTextSplitterName() *string {
	return s.TextSplitterName
}

func (s *DescribeDocumentResponseBodyData) SetDescription(v string) *DescribeDocumentResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetDocsCount(v int64) *DescribeDocumentResponseBodyData {
	s.DocsCount = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetDocumentLoaderName(v string) *DescribeDocumentResponseBodyData {
	s.DocumentLoaderName = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetFileExt(v string) *DescribeDocumentResponseBodyData {
	s.FileExt = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetFileSize(v int64) *DescribeDocumentResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetGmtCreate(v string) *DescribeDocumentResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetGmtModified(v string) *DescribeDocumentResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetKbUuid(v string) *DescribeDocumentResponseBodyData {
	s.KbUuid = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetKeywords(v string) *DescribeDocumentResponseBodyData {
	s.Keywords = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetName(v string) *DescribeDocumentResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetState(v int64) *DescribeDocumentResponseBodyData {
	s.State = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetSummary(v string) *DescribeDocumentResponseBodyData {
	s.Summary = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) SetTextSplitterName(v string) *DescribeDocumentResponseBodyData {
	s.TextSplitterName = &v
	return s
}

func (s *DescribeDocumentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
