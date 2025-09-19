// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupFiles(v []*DescribeBackupFilesResponseBodyBackupFiles) *DescribeBackupFilesResponseBody
	GetBackupFiles() []*DescribeBackupFilesResponseBodyBackupFiles
	SetPageInfo(v *DescribeBackupFilesResponseBodyPageInfo) *DescribeBackupFilesResponseBody
	GetPageInfo() *DescribeBackupFilesResponseBodyPageInfo
	SetRequestId(v string) *DescribeBackupFilesResponseBody
	GetRequestId() *string
}

type DescribeBackupFilesResponseBody struct {
	// An array that consists of the backup files returned.
	BackupFiles []*DescribeBackupFilesResponseBodyBackupFiles `json:"BackupFiles,omitempty" xml:"BackupFiles,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeBackupFilesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 00A60A6D-33E0-5D5A-9B7C-E5D4DCA88148
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesResponseBody) GetBackupFiles() []*DescribeBackupFilesResponseBodyBackupFiles {
	return s.BackupFiles
}

func (s *DescribeBackupFilesResponseBody) GetPageInfo() *DescribeBackupFilesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeBackupFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupFilesResponseBody) SetBackupFiles(v []*DescribeBackupFilesResponseBodyBackupFiles) *DescribeBackupFilesResponseBody {
	s.BackupFiles = v
	return s
}

func (s *DescribeBackupFilesResponseBody) SetPageInfo(v *DescribeBackupFilesResponseBodyPageInfo) *DescribeBackupFilesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeBackupFilesResponseBody) SetRequestId(v string) *DescribeBackupFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupFilesResponseBodyBackupFiles struct {
	// The name of the anti-ransomware policy.
	//
	// example:
	//
	// Group 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the backup file. Unit: bytes.
	//
	// example:
	//
	// 100
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The path to the subdirectory of the backup file.
	//
	// example:
	//
	// Python27\\
	Subtree *string `json:"Subtree,omitempty" xml:"Subtree,omitempty"`
	// The type of the protected file. Valid values:
	//
	// 	- **file**: files
	//
	// 	- **dir**: folders
	//
	// example:
	//
	// dir
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeBackupFilesResponseBodyBackupFiles) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupFilesResponseBodyBackupFiles) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesResponseBodyBackupFiles) GetName() *string {
	return s.Name
}

func (s *DescribeBackupFilesResponseBodyBackupFiles) GetSize() *int64 {
	return s.Size
}

func (s *DescribeBackupFilesResponseBodyBackupFiles) GetSubtree() *string {
	return s.Subtree
}

func (s *DescribeBackupFilesResponseBodyBackupFiles) GetType() *string {
	return s.Type
}

func (s *DescribeBackupFilesResponseBodyBackupFiles) SetName(v string) *DescribeBackupFilesResponseBodyBackupFiles {
	s.Name = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyBackupFiles) SetSize(v int64) *DescribeBackupFilesResponseBodyBackupFiles {
	s.Size = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyBackupFiles) SetSubtree(v string) *DescribeBackupFilesResponseBodyBackupFiles {
	s.Subtree = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyBackupFiles) SetType(v string) *DescribeBackupFilesResponseBodyBackupFiles {
	s.Type = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyBackupFiles) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupFilesResponseBodyPageInfo struct {
	// The number of backup files returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of backup files returned.
	//
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupFilesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupFilesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeBackupFilesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBackupFilesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupFilesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBackupFilesResponseBodyPageInfo) SetCount(v int32) *DescribeBackupFilesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeBackupFilesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyPageInfo) SetPageSize(v int32) *DescribeBackupFilesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeBackupFilesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
