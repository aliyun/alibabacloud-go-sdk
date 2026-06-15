// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListValidateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*ListValidateFileResponseBodyFiles) *ListValidateFileResponseBody
	GetFiles() []*ListValidateFileResponseBodyFiles
	SetHasNext(v bool) *ListValidateFileResponseBody
	GetHasNext() *bool
	SetPage(v int32) *ListValidateFileResponseBody
	GetPage() *int32
	SetPageSize(v int32) *ListValidateFileResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListValidateFileResponseBody
	GetRequestId() *string
	SetTotalPages(v int32) *ListValidateFileResponseBody
	GetTotalPages() *int32
	SetTotalSize(v int32) *ListValidateFileResponseBody
	GetTotalSize() *int32
}

type ListValidateFileResponseBody struct {
	// The list of files.
	Files []*ListValidateFileResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// Indicates whether a next page of data exists.
	//
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 10
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListValidateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListValidateFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListValidateFileResponseBody) GetFiles() []*ListValidateFileResponseBodyFiles {
	return s.Files
}

func (s *ListValidateFileResponseBody) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListValidateFileResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *ListValidateFileResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListValidateFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListValidateFileResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListValidateFileResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListValidateFileResponseBody) SetFiles(v []*ListValidateFileResponseBodyFiles) *ListValidateFileResponseBody {
	s.Files = v
	return s
}

func (s *ListValidateFileResponseBody) SetHasNext(v bool) *ListValidateFileResponseBody {
	s.HasNext = &v
	return s
}

func (s *ListValidateFileResponseBody) SetPage(v int32) *ListValidateFileResponseBody {
	s.Page = &v
	return s
}

func (s *ListValidateFileResponseBody) SetPageSize(v int32) *ListValidateFileResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListValidateFileResponseBody) SetRequestId(v string) *ListValidateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListValidateFileResponseBody) SetTotalPages(v int32) *ListValidateFileResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListValidateFileResponseBody) SetTotalSize(v int32) *ListValidateFileResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListValidateFileResponseBody) Validate() error {
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

type ListValidateFileResponseBodyFiles struct {
	// The number of addresses with the validation result \\"CatchAll\\".
	//
	// example:
	//
	// 1
	CatchAllNum *string `json:"CatchAllNum,omitempty" xml:"CatchAllNum,omitempty"`
	// The time when the task was completed.
	//
	// example:
	//
	// 2000-01-01T00:00:00Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The number of addresses with the validation result \\"DoNotMail\\".
	//
	// example:
	//
	// 0
	DoNotMailNum *string `json:"DoNotMailNum,omitempty" xml:"DoNotMailNum,omitempty"`
	// The file ID.
	//
	// example:
	//
	// xxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file name.
	//
	// example:
	//
	// test.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The number of addresses with the validation result \\"Invalid\\".
	//
	// example:
	//
	// 4
	InvalidNum *string `json:"InvalidNum,omitempty" xml:"InvalidNum,omitempty"`
	// Indicates whether the result can be downloaded.
	//
	// example:
	//
	// true
	IsDownloadable *bool `json:"IsDownloadable,omitempty" xml:"IsDownloadable,omitempty"`
	// The task execution progress.
	//
	// example:
	//
	// 100%
	Percentage *string `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The number of addresses that have been validated.
	//
	// example:
	//
	// 10
	ProcessedNum *string `json:"ProcessedNum,omitempty" xml:"ProcessedNum,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of addresses to validate in the task.
	//
	// example:
	//
	// 10
	TotalNum *string `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The number of addresses with the validation result \\"Unknown\\".
	//
	// example:
	//
	// 0
	UnknownNum *string `json:"UnknownNum,omitempty" xml:"UnknownNum,omitempty"`
	// The time when the file was submitted.
	//
	// example:
	//
	// 2000-01-01T00:00:00Z
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
	// The number of addresses with the validation result \\"Valid\\".
	//
	// example:
	//
	// 5
	ValidNum *string `json:"ValidNum,omitempty" xml:"ValidNum,omitempty"`
}

func (s ListValidateFileResponseBodyFiles) String() string {
	return dara.Prettify(s)
}

func (s ListValidateFileResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *ListValidateFileResponseBodyFiles) GetCatchAllNum() *string {
	return s.CatchAllNum
}

func (s *ListValidateFileResponseBodyFiles) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *ListValidateFileResponseBodyFiles) GetDoNotMailNum() *string {
	return s.DoNotMailNum
}

func (s *ListValidateFileResponseBodyFiles) GetFileId() *string {
	return s.FileId
}

func (s *ListValidateFileResponseBodyFiles) GetFileName() *string {
	return s.FileName
}

func (s *ListValidateFileResponseBodyFiles) GetInvalidNum() *string {
	return s.InvalidNum
}

func (s *ListValidateFileResponseBodyFiles) GetIsDownloadable() *bool {
	return s.IsDownloadable
}

func (s *ListValidateFileResponseBodyFiles) GetPercentage() *string {
	return s.Percentage
}

func (s *ListValidateFileResponseBodyFiles) GetProcessedNum() *string {
	return s.ProcessedNum
}

func (s *ListValidateFileResponseBodyFiles) GetStatus() *string {
	return s.Status
}

func (s *ListValidateFileResponseBodyFiles) GetTotalNum() *string {
	return s.TotalNum
}

func (s *ListValidateFileResponseBodyFiles) GetUnknownNum() *string {
	return s.UnknownNum
}

func (s *ListValidateFileResponseBodyFiles) GetUploadTime() *string {
	return s.UploadTime
}

func (s *ListValidateFileResponseBodyFiles) GetValidNum() *string {
	return s.ValidNum
}

func (s *ListValidateFileResponseBodyFiles) SetCatchAllNum(v string) *ListValidateFileResponseBodyFiles {
	s.CatchAllNum = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetCompleteTime(v string) *ListValidateFileResponseBodyFiles {
	s.CompleteTime = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetDoNotMailNum(v string) *ListValidateFileResponseBodyFiles {
	s.DoNotMailNum = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetFileId(v string) *ListValidateFileResponseBodyFiles {
	s.FileId = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetFileName(v string) *ListValidateFileResponseBodyFiles {
	s.FileName = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetInvalidNum(v string) *ListValidateFileResponseBodyFiles {
	s.InvalidNum = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetIsDownloadable(v bool) *ListValidateFileResponseBodyFiles {
	s.IsDownloadable = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetPercentage(v string) *ListValidateFileResponseBodyFiles {
	s.Percentage = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetProcessedNum(v string) *ListValidateFileResponseBodyFiles {
	s.ProcessedNum = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetStatus(v string) *ListValidateFileResponseBodyFiles {
	s.Status = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetTotalNum(v string) *ListValidateFileResponseBodyFiles {
	s.TotalNum = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetUnknownNum(v string) *ListValidateFileResponseBodyFiles {
	s.UnknownNum = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetUploadTime(v string) *ListValidateFileResponseBodyFiles {
	s.UploadTime = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) SetValidNum(v string) *ListValidateFileResponseBodyFiles {
	s.ValidNum = &v
	return s
}

func (s *ListValidateFileResponseBodyFiles) Validate() error {
	return dara.Validate(s)
}
