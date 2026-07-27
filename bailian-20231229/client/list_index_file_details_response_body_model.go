// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexFileDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIndexFileDetailsResponseBody
	GetCode() *string
	SetData(v *ListIndexFileDetailsResponseBodyData) *ListIndexFileDetailsResponseBody
	GetData() *ListIndexFileDetailsResponseBodyData
	SetMessage(v string) *ListIndexFileDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIndexFileDetailsResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListIndexFileDetailsResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListIndexFileDetailsResponseBody
	GetSuccess() *bool
}

type ListIndexFileDetailsResponseBody struct {
	// The error status code.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data field of the operation.
	Data *ListIndexFileDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35A267BF-xxxx-54DB-8394-AA3B0742D833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the operation.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIndexFileDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIndexFileDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexFileDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIndexFileDetailsResponseBody) GetData() *ListIndexFileDetailsResponseBodyData {
	return s.Data
}

func (s *ListIndexFileDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIndexFileDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIndexFileDetailsResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListIndexFileDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIndexFileDetailsResponseBody) SetCode(v string) *ListIndexFileDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ListIndexFileDetailsResponseBody) SetData(v *ListIndexFileDetailsResponseBodyData) *ListIndexFileDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListIndexFileDetailsResponseBody) SetMessage(v string) *ListIndexFileDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListIndexFileDetailsResponseBody) SetRequestId(v string) *ListIndexFileDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndexFileDetailsResponseBody) SetStatus(v string) *ListIndexFileDetailsResponseBody {
	s.Status = &v
	return s
}

func (s *ListIndexFileDetailsResponseBody) SetSuccess(v bool) *ListIndexFileDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *ListIndexFileDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIndexFileDetailsResponseBodyData struct {
	// The list of files in the knowledge base, sorted by file import time in descending order (consistent with the console).
	Documents []*ListIndexFileDetailsResponseBodyDataDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// The knowledge base ID.
	//
	// example:
	//
	// 79c0alxxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The specified page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The specified number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned results.
	//
	// example:
	//
	// 2437
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIndexFileDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIndexFileDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIndexFileDetailsResponseBodyData) GetDocuments() []*ListIndexFileDetailsResponseBodyDataDocuments {
	return s.Documents
}

func (s *ListIndexFileDetailsResponseBodyData) GetIndexId() *string {
	return s.IndexId
}

func (s *ListIndexFileDetailsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIndexFileDetailsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIndexFileDetailsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIndexFileDetailsResponseBodyData) SetDocuments(v []*ListIndexFileDetailsResponseBodyDataDocuments) *ListIndexFileDetailsResponseBodyData {
	s.Documents = v
	return s
}

func (s *ListIndexFileDetailsResponseBodyData) SetIndexId(v string) *ListIndexFileDetailsResponseBodyData {
	s.IndexId = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyData) SetPageNumber(v int32) *ListIndexFileDetailsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyData) SetPageSize(v int32) *ListIndexFileDetailsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyData) SetTotalCount(v int64) *ListIndexFileDetailsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyData) Validate() error {
	if s.Documents != nil {
		for _, item := range s.Documents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIndexFileDetailsResponseBodyDataDocuments struct {
	// The custom chunking mode.
	//
	// example:
	//
	// DashSplitter
	ChunkMode *string `json:"ChunkMode,omitempty" xml:"ChunkMode,omitempty"`
	// The chunk length, which is the number of characters in a text chunk.
	//
	// example:
	//
	// 600
	ChunkSize *string `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// The file import error status code.
	//
	// example:
	//
	// 110002
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The file format type. Valid values: pdf, docx, doc, txt, md, pptx, ppt, png, jpg, jpeg, bmp, gif, EXCEL.
	//
	// example:
	//
	// pdf
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// Indicates whether Excel file headers support concatenation.
	//
	// example:
	//
	// true
	EnableHeaders *string `json:"EnableHeaders,omitempty" xml:"EnableHeaders,omitempty"`
	// The time when the file was imported to the knowledge base, in Unix timestamp format.
	//
	// example:
	//
	// 1744856423000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The file ID.
	//
	// example:
	//
	// doc_c134aa2073204a5d936d870bf960f56axxxxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The file import error message.
	//
	// example:
	//
	// check fileUrlKey[file_path] / fileNameKey[null] / fileExtensionKey[file_extension] is invalid
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MetaExtractInfo *string `json:"MetaExtractInfo,omitempty" xml:"MetaExtractInfo,omitempty"`
	// The file name.
	//
	// example:
	//
	// TranslationPlatformO&MDocument.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The chunk overlap length.
	//
	// example:
	//
	// 100
	OverlapSize *string `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// The file size, in bytes.
	//
	// example:
	//
	// 996764
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The category ID.
	//
	// example:
	//
	// cate_21a407a3372c4ba7aedc649709143f0cxxxxxxxx
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The file import status. Valid values:
	//
	// - INSERT_ERROR: File import failed.
	//
	// - RUNNING: File import in progress.
	//
	// - DELETED: File deleted.
	//
	// - FINISH: File import succeeded.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sentence separator.
	//
	// example:
	//
	// " "
	Separator *string `json:"separator,omitempty" xml:"separator,omitempty"`
}

func (s ListIndexFileDetailsResponseBodyDataDocuments) String() string {
	return dara.Prettify(s)
}

func (s ListIndexFileDetailsResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetChunkMode() *string {
	return s.ChunkMode
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetChunkSize() *string {
	return s.ChunkSize
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetCode() *string {
	return s.Code
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetDocumentType() *string {
	return s.DocumentType
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetEnableHeaders() *string {
	return s.EnableHeaders
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetId() *string {
	return s.Id
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetMessage() *string {
	return s.Message
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetMetaExtractInfo() *string {
	return s.MetaExtractInfo
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetName() *string {
	return s.Name
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetOverlapSize() *string {
	return s.OverlapSize
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetSize() *int32 {
	return s.Size
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetSourceId() *string {
	return s.SourceId
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetStatus() *string {
	return s.Status
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) GetSeparator() *string {
	return s.Separator
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetChunkMode(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.ChunkMode = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetChunkSize(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.ChunkSize = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetCode(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.Code = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetDocumentType(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.DocumentType = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetEnableHeaders(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.EnableHeaders = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetGmtModified(v int64) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.GmtModified = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetId(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.Id = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetMessage(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.Message = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetMetaExtractInfo(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.MetaExtractInfo = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetName(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.Name = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetOverlapSize(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.OverlapSize = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetSize(v int32) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.Size = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetSourceId(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.SourceId = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetStatus(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.Status = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) SetSeparator(v string) *ListIndexFileDetailsResponseBodyDataDocuments {
	s.Separator = &v
	return s
}

func (s *ListIndexFileDetailsResponseBodyDataDocuments) Validate() error {
	return dara.Validate(s)
}
