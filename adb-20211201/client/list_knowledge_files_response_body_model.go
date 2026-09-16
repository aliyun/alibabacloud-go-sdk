// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListKnowledgeFilesResponseBodyData) *ListKnowledgeFilesResponseBody
	GetData() *ListKnowledgeFilesResponseBodyData
	SetRequestId(v string) *ListKnowledgeFilesResponseBody
	GetRequestId() *string
}

type ListKnowledgeFilesResponseBody struct {
	Data *ListKnowledgeFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListKnowledgeFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeFilesResponseBody) GetData() *ListKnowledgeFilesResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeFilesResponseBody) SetData(v *ListKnowledgeFilesResponseBodyData) *ListKnowledgeFilesResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeFilesResponseBody) SetRequestId(v string) *ListKnowledgeFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeFilesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeFilesResponseBodyData struct {
	Files []*ListKnowledgeFilesResponseBodyDataFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListKnowledgeFilesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeFilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeFilesResponseBodyData) GetFiles() []*ListKnowledgeFilesResponseBodyDataFiles {
	return s.Files
}

func (s *ListKnowledgeFilesResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeFilesResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListKnowledgeFilesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKnowledgeFilesResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeFilesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListKnowledgeFilesResponseBodyData) SetFiles(v []*ListKnowledgeFilesResponseBodyDataFiles) *ListKnowledgeFilesResponseBodyData {
	s.Files = v
	return s
}

func (s *ListKnowledgeFilesResponseBodyData) SetMessage(v string) *ListKnowledgeFilesResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyData) SetPage(v int32) *ListKnowledgeFilesResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyData) SetPageSize(v int32) *ListKnowledgeFilesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyData) SetSuccess(v bool) *ListKnowledgeFilesResponseBodyData {
	s.Success = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyData) SetTotal(v int64) *ListKnowledgeFilesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyData) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeFilesResponseBodyDataFiles struct {
	// example:
	//
	// 2026-06-09 10:27:35
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 137
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// 1024
	FileSizeBytes *int64 `json:"FileSizeBytes,omitempty" xml:"FileSizeBytes,omitempty"`
	// example:
	//
	// oss://bucketName/path/to/file
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// false
	IsDirectory *bool `json:"IsDirectory,omitempty" xml:"IsDirectory,omitempty"`
	// example:
	//
	// 122
	OwnerFileId *int64 `json:"OwnerFileId,omitempty" xml:"OwnerFileId,omitempty"`
	// example:
	//
	// 2
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// successful
	ProcessMessage *string `json:"ProcessMessage,omitempty" xml:"ProcessMessage,omitempty"`
	// example:
	//
	// Finished
	ProcessStatus *string `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// example:
	//
	// 2026-06-10 10:23:46
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListKnowledgeFilesResponseBodyDataFiles) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeFilesResponseBodyDataFiles) GoString() string {
	return s.String()
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetFileId() *int64 {
	return s.FileId
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetFileSizeBytes() *int64 {
	return s.FileSizeBytes
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetFormat() *string {
	return s.Format
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetIsDirectory() *bool {
	return s.IsDirectory
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetOwnerFileId() *int64 {
	return s.OwnerFileId
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetPageCount() *int32 {
	return s.PageCount
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetProcessMessage() *string {
	return s.ProcessMessage
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetProcessStatus() *string {
	return s.ProcessStatus
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetCreatedAt(v string) *ListKnowledgeFilesResponseBodyDataFiles {
	s.CreatedAt = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetFileId(v int64) *ListKnowledgeFilesResponseBodyDataFiles {
	s.FileId = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetFileSizeBytes(v int64) *ListKnowledgeFilesResponseBodyDataFiles {
	s.FileSizeBytes = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetFileUrl(v string) *ListKnowledgeFilesResponseBodyDataFiles {
	s.FileUrl = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetFormat(v string) *ListKnowledgeFilesResponseBodyDataFiles {
	s.Format = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetIsDirectory(v bool) *ListKnowledgeFilesResponseBodyDataFiles {
	s.IsDirectory = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetOwnerFileId(v int64) *ListKnowledgeFilesResponseBodyDataFiles {
	s.OwnerFileId = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetPageCount(v int32) *ListKnowledgeFilesResponseBodyDataFiles {
	s.PageCount = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetProcessMessage(v string) *ListKnowledgeFilesResponseBodyDataFiles {
	s.ProcessMessage = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetProcessStatus(v string) *ListKnowledgeFilesResponseBodyDataFiles {
	s.ProcessStatus = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) SetUpdatedAt(v string) *ListKnowledgeFilesResponseBodyDataFiles {
	s.UpdatedAt = &v
	return s
}

func (s *ListKnowledgeFilesResponseBodyDataFiles) Validate() error {
	return dara.Validate(s)
}
