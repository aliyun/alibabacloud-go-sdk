// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcludeSystemPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExcludePaths(v []*DescribeExcludeSystemPathResponseBodyExcludePaths) *DescribeExcludeSystemPathResponseBody
	GetExcludePaths() []*DescribeExcludeSystemPathResponseBodyExcludePaths
	SetPageInfo(v *DescribeExcludeSystemPathResponseBodyPageInfo) *DescribeExcludeSystemPathResponseBody
	GetPageInfo() *DescribeExcludeSystemPathResponseBodyPageInfo
	SetRequestId(v string) *DescribeExcludeSystemPathResponseBody
	GetRequestId() *string
}

type DescribeExcludeSystemPathResponseBody struct {
	// An array consisting of the directories that are excluded.
	ExcludePaths []*DescribeExcludeSystemPathResponseBodyExcludePaths `json:"ExcludePaths,omitempty" xml:"ExcludePaths,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeExcludeSystemPathResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// FBBEB173-1F43-505F-A876-C03ECDF6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExcludeSystemPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcludeSystemPathResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExcludeSystemPathResponseBody) GetExcludePaths() []*DescribeExcludeSystemPathResponseBodyExcludePaths {
	return s.ExcludePaths
}

func (s *DescribeExcludeSystemPathResponseBody) GetPageInfo() *DescribeExcludeSystemPathResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeExcludeSystemPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExcludeSystemPathResponseBody) SetExcludePaths(v []*DescribeExcludeSystemPathResponseBodyExcludePaths) *DescribeExcludeSystemPathResponseBody {
	s.ExcludePaths = v
	return s
}

func (s *DescribeExcludeSystemPathResponseBody) SetPageInfo(v *DescribeExcludeSystemPathResponseBodyPageInfo) *DescribeExcludeSystemPathResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeExcludeSystemPathResponseBody) SetRequestId(v string) *DescribeExcludeSystemPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExcludeSystemPathResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExcludeSystemPathResponseBodyExcludePaths struct {
	// The operating system of the server. Valid values:
	//
	// 	- **linux**: Linux
	//
	// 	- **windows**: Windows
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The absolute path to the directory.
	//
	// example:
	//
	// /bin/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeExcludeSystemPathResponseBodyExcludePaths) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcludeSystemPathResponseBodyExcludePaths) GoString() string {
	return s.String()
}

func (s *DescribeExcludeSystemPathResponseBodyExcludePaths) GetOs() *string {
	return s.Os
}

func (s *DescribeExcludeSystemPathResponseBodyExcludePaths) GetPath() *string {
	return s.Path
}

func (s *DescribeExcludeSystemPathResponseBodyExcludePaths) SetOs(v string) *DescribeExcludeSystemPathResponseBodyExcludePaths {
	s.Os = &v
	return s
}

func (s *DescribeExcludeSystemPathResponseBodyExcludePaths) SetPath(v string) *DescribeExcludeSystemPathResponseBodyExcludePaths {
	s.Path = &v
	return s
}

func (s *DescribeExcludeSystemPathResponseBodyExcludePaths) Validate() error {
	return dara.Validate(s)
}

type DescribeExcludeSystemPathResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExcludeSystemPathResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcludeSystemPathResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeExcludeSystemPathResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeExcludeSystemPathResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeExcludeSystemPathResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExcludeSystemPathResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeExcludeSystemPathResponseBodyPageInfo) SetCount(v int32) *DescribeExcludeSystemPathResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeExcludeSystemPathResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeExcludeSystemPathResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeExcludeSystemPathResponseBodyPageInfo) SetPageSize(v int32) *DescribeExcludeSystemPathResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeExcludeSystemPathResponseBodyPageInfo) SetTotalCount(v int32) *DescribeExcludeSystemPathResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeExcludeSystemPathResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
