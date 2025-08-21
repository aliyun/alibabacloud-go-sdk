// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*ListFilesResponseBodyFiles) *ListFilesResponseBody
	GetFiles() []*ListFilesResponseBodyFiles
	SetPageNumber(v int64) *ListFilesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListFilesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListFilesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListFilesResponseBody
	GetTotalCount() *int64
}

type ListFilesResponseBody struct {
	Files []*ListFilesResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBody) GetFiles() []*ListFilesResponseBodyFiles {
	return s.Files
}

func (s *ListFilesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListFilesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFilesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFilesResponseBody) SetFiles(v []*ListFilesResponseBodyFiles) *ListFilesResponseBody {
	s.Files = v
	return s
}

func (s *ListFilesResponseBody) SetPageNumber(v int64) *ListFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFilesResponseBody) SetPageSize(v int64) *ListFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFilesResponseBody) SetRequestId(v string) *ListFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFilesResponseBody) SetTotalCount(v int64) *ListFilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFilesResponseBodyFiles struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// f-1671accd4dafdag3er256cvgewt13f7141db2f7
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// myfile
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// upload success
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
	// example:
	//
	// /data/tmp/test/xxx.tar
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// example:
	//
	// 2024-03-28T14:15:08+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 2024-03-28T14:10:12+08:00
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s ListFilesResponseBodyFiles) String() string {
	return dara.Prettify(s)
}

func (s ListFilesResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBodyFiles) GetDescription() *string {
	return s.Description
}

func (s *ListFilesResponseBodyFiles) GetFileId() *string {
	return s.FileId
}

func (s *ListFilesResponseBodyFiles) GetFileName() *string {
	return s.FileName
}

func (s *ListFilesResponseBodyFiles) GetStatus() *string {
	return s.Status
}

func (s *ListFilesResponseBodyFiles) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *ListFilesResponseBodyFiles) GetTargetPath() *string {
	return s.TargetPath
}

func (s *ListFilesResponseBodyFiles) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListFilesResponseBodyFiles) GetUploadTime() *string {
	return s.UploadTime
}

func (s *ListFilesResponseBodyFiles) SetDescription(v string) *ListFilesResponseBodyFiles {
	s.Description = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetFileId(v string) *ListFilesResponseBodyFiles {
	s.FileId = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetFileName(v string) *ListFilesResponseBodyFiles {
	s.FileName = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetStatus(v string) *ListFilesResponseBodyFiles {
	s.Status = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetStatusDescription(v string) *ListFilesResponseBodyFiles {
	s.StatusDescription = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetTargetPath(v string) *ListFilesResponseBodyFiles {
	s.TargetPath = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetUpdateTime(v string) *ListFilesResponseBodyFiles {
	s.UpdateTime = &v
	return s
}

func (s *ListFilesResponseBodyFiles) SetUploadTime(v string) *ListFilesResponseBodyFiles {
	s.UploadTime = &v
	return s
}

func (s *ListFilesResponseBodyFiles) Validate() error {
	return dara.Validate(s)
}
