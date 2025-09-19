// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockFileChangeStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWebLockFileChangeStatisticsResponseBody
	GetCurrentPage() *int32
	SetList(v []*DescribeWebLockFileChangeStatisticsResponseBodyList) *DescribeWebLockFileChangeStatisticsResponseBody
	GetList() []*DescribeWebLockFileChangeStatisticsResponseBodyList
	SetPageSize(v int32) *DescribeWebLockFileChangeStatisticsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWebLockFileChangeStatisticsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWebLockFileChangeStatisticsResponseBody
	GetTotalCount() *int32
}

type DescribeWebLockFileChangeStatisticsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array consisting of the files that are changed.
	List []*DescribeWebLockFileChangeStatisticsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 709A8C3D-A543-5B79-AB75-361B206F71D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of files that are attempted to change.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebLockFileChangeStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileChangeStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) GetList() []*DescribeWebLockFileChangeStatisticsResponseBodyList {
	return s.List
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) SetCurrentPage(v int32) *DescribeWebLockFileChangeStatisticsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) SetList(v []*DescribeWebLockFileChangeStatisticsResponseBodyList) *DescribeWebLockFileChangeStatisticsResponseBody {
	s.List = v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) SetPageSize(v int32) *DescribeWebLockFileChangeStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) SetRequestId(v string) *DescribeWebLockFileChangeStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) SetTotalCount(v int32) *DescribeWebLockFileChangeStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWebLockFileChangeStatisticsResponseBodyList struct {
	// The number of attempts.
	//
	// example:
	//
	// 33
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The file path.
	//
	// example:
	//
	// /tmp
	File *string `json:"File,omitempty" xml:"File,omitempty"`
}

func (s DescribeWebLockFileChangeStatisticsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileChangeStatisticsResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileChangeStatisticsResponseBodyList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeWebLockFileChangeStatisticsResponseBodyList) GetFile() *string {
	return s.File
}

func (s *DescribeWebLockFileChangeStatisticsResponseBodyList) SetCount(v int32) *DescribeWebLockFileChangeStatisticsResponseBodyList {
	s.Count = &v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsResponseBodyList) SetFile(v string) *DescribeWebLockFileChangeStatisticsResponseBodyList {
	s.File = &v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsResponseBodyList) Validate() error {
	return dara.Validate(s)
}
