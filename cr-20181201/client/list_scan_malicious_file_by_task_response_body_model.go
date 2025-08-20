// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanMaliciousFileByTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListScanMaliciousFileByTaskResponseBody
	GetCode() *int32
	SetIsSuccess(v bool) *ListScanMaliciousFileByTaskResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListScanMaliciousFileByTaskResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListScanMaliciousFileByTaskResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListScanMaliciousFileByTaskResponseBody
	GetRequestId() *string
	SetScanMaliciousFiles(v []*ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) *ListScanMaliciousFileByTaskResponseBody
	GetScanMaliciousFiles() []*ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles
	SetTotalCount(v int32) *ListScanMaliciousFileByTaskResponseBody
	GetTotalCount() *int32
}

type ListScanMaliciousFileByTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 52AE49C8-B91A-5C1A-821F-C34324B42F7C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried malicious files.
	ScanMaliciousFiles []*ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles `json:"ScanMaliciousFiles,omitempty" xml:"ScanMaliciousFiles,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScanMaliciousFileByTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScanMaliciousFileByTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListScanMaliciousFileByTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListScanMaliciousFileByTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListScanMaliciousFileByTaskResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListScanMaliciousFileByTaskResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScanMaliciousFileByTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScanMaliciousFileByTaskResponseBody) GetScanMaliciousFiles() []*ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	return s.ScanMaliciousFiles
}

func (s *ListScanMaliciousFileByTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetCode(v int32) *ListScanMaliciousFileByTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetIsSuccess(v bool) *ListScanMaliciousFileByTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetPageNo(v int32) *ListScanMaliciousFileByTaskResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetPageSize(v int32) *ListScanMaliciousFileByTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetRequestId(v string) *ListScanMaliciousFileByTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetScanMaliciousFiles(v []*ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) *ListScanMaliciousFileByTaskResponseBody {
	s.ScanMaliciousFiles = v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetTotalCount(v int32) *ListScanMaliciousFileByTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles struct {
	// The time when the image was created.
	//
	// example:
	//
	// 2023-04-10 11:42:06
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The path of the file.
	//
	// example:
	//
	// tenant/0000000000000000/
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The time when the image was first scanned.
	//
	// example:
	//
	// 2023-04-10 11:42:06
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The severity of the malicious file.
	//
	// example:
	//
	// remind
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The MD5 hash value of the malicious file.
	//
	// example:
	//
	// e76c9759783cbbc9be0ff91ca3xxxxxx
	MaliciousMd5 *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	// The type of the malicious file.
	//
	// example:
	//
	// Suspected to contain Webshell code
	MaliciousName *string `json:"MaliciousName,omitempty" xml:"MaliciousName,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// fe2d8980-de45-4f55-b57d-e438e6d2e972
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The time when the image was updated.
	//
	// example:
	//
	// 2023-04-10 11:42:06
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) String() string {
	return dara.Prettify(s)
}

func (s ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) GoString() string {
	return s.String()
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) GetFilePath() *string {
	return s.FilePath
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) GetLevel() *string {
	return s.Level
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) GetMaliciousMd5() *string {
	return s.MaliciousMd5
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) GetMaliciousName() *string {
	return s.MaliciousName
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) GetScanTaskId() *string {
	return s.ScanTaskId
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetCreateTime(v int64) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.CreateTime = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetFilePath(v string) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.FilePath = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetFirstScanTime(v int64) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.FirstScanTime = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetLevel(v string) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.Level = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetMaliciousMd5(v string) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.MaliciousMd5 = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetMaliciousName(v string) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.MaliciousName = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetScanTaskId(v string) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.ScanTaskId = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetUpdateTime(v int64) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.UpdateTime = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) Validate() error {
	return dara.Validate(s)
}
