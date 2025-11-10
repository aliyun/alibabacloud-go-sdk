// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRepoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddTargetCount(v int32) *DescribeImageRepoListResponseBody
	GetAddTargetCount() *int32
	SetAllTargetCount(v int32) *DescribeImageRepoListResponseBody
	GetAllTargetCount() *int32
	SetDelTargetCount(v int32) *DescribeImageRepoListResponseBody
	GetDelTargetCount() *int32
	SetImageRepoList(v []*DescribeImageRepoListResponseBodyImageRepoList) *DescribeImageRepoListResponseBody
	GetImageRepoList() []*DescribeImageRepoListResponseBodyImageRepoList
	SetPageInfo(v *DescribeImageRepoListResponseBodyPageInfo) *DescribeImageRepoListResponseBody
	GetPageInfo() *DescribeImageRepoListResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageRepoListResponseBody
	GetRequestId() *string
}

type DescribeImageRepoListResponseBody struct {
	// The number of image repositories that are added to Security Center.
	//
	// example:
	//
	// 20
	AddTargetCount *int32 `json:"AddTargetCount,omitempty" xml:"AddTargetCount,omitempty"`
	// The total number of image repositories.
	//
	// example:
	//
	// 25
	AllTargetCount *int32 `json:"AllTargetCount,omitempty" xml:"AllTargetCount,omitempty"`
	// The number of excluded image repositories.
	//
	// example:
	//
	// 5
	DelTargetCount *int32 `json:"DelTargetCount,omitempty" xml:"DelTargetCount,omitempty"`
	// An array that consists of the information about image repositories.
	ImageRepoList []*DescribeImageRepoListResponseBodyImageRepoList `json:"ImageRepoList,omitempty" xml:"ImageRepoList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageRepoListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageRepoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoListResponseBody) GetAddTargetCount() *int32 {
	return s.AddTargetCount
}

func (s *DescribeImageRepoListResponseBody) GetAllTargetCount() *int32 {
	return s.AllTargetCount
}

func (s *DescribeImageRepoListResponseBody) GetDelTargetCount() *int32 {
	return s.DelTargetCount
}

func (s *DescribeImageRepoListResponseBody) GetImageRepoList() []*DescribeImageRepoListResponseBodyImageRepoList {
	return s.ImageRepoList
}

func (s *DescribeImageRepoListResponseBody) GetPageInfo() *DescribeImageRepoListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageRepoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageRepoListResponseBody) SetAddTargetCount(v int32) *DescribeImageRepoListResponseBody {
	s.AddTargetCount = &v
	return s
}

func (s *DescribeImageRepoListResponseBody) SetAllTargetCount(v int32) *DescribeImageRepoListResponseBody {
	s.AllTargetCount = &v
	return s
}

func (s *DescribeImageRepoListResponseBody) SetDelTargetCount(v int32) *DescribeImageRepoListResponseBody {
	s.DelTargetCount = &v
	return s
}

func (s *DescribeImageRepoListResponseBody) SetImageRepoList(v []*DescribeImageRepoListResponseBodyImageRepoList) *DescribeImageRepoListResponseBody {
	s.ImageRepoList = v
	return s
}

func (s *DescribeImageRepoListResponseBody) SetPageInfo(v *DescribeImageRepoListResponseBodyPageInfo) *DescribeImageRepoListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageRepoListResponseBody) SetRequestId(v string) *DescribeImageRepoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageRepoListResponseBody) Validate() error {
	if s.ImageRepoList != nil {
		for _, item := range s.ImageRepoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageRepoListResponseBodyImageRepoList struct {
	// Indicates whether the feature takes effect on the image repository. Valid values:
	//
	// 	- **add**: yes
	//
	// 	- **del**: no
	//
	// example:
	//
	// add
	Flag       *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	ImageCount *int32  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// zeus
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// sas-script
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
}

func (s DescribeImageRepoListResponseBodyImageRepoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoListResponseBodyImageRepoList) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoListResponseBodyImageRepoList) GetFlag() *string {
	return s.Flag
}

func (s *DescribeImageRepoListResponseBodyImageRepoList) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *DescribeImageRepoListResponseBodyImageRepoList) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageRepoListResponseBodyImageRepoList) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageRepoListResponseBodyImageRepoList) SetFlag(v string) *DescribeImageRepoListResponseBodyImageRepoList {
	s.Flag = &v
	return s
}

func (s *DescribeImageRepoListResponseBodyImageRepoList) SetImageCount(v int32) *DescribeImageRepoListResponseBodyImageRepoList {
	s.ImageCount = &v
	return s
}

func (s *DescribeImageRepoListResponseBodyImageRepoList) SetRepoName(v string) *DescribeImageRepoListResponseBodyImageRepoList {
	s.RepoName = &v
	return s
}

func (s *DescribeImageRepoListResponseBodyImageRepoList) SetRepoNamespace(v string) *DescribeImageRepoListResponseBodyImageRepoList {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageRepoListResponseBodyImageRepoList) Validate() error {
	return dara.Validate(s)
}

type DescribeImageRepoListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
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
	// The total number of image repositories.
	//
	// example:
	//
	// 83
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageRepoListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageRepoListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageRepoListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageRepoListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageRepoListResponseBodyPageInfo) SetCount(v int32) *DescribeImageRepoListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageRepoListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageRepoListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageRepoListResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageRepoListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageRepoListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageRepoListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageRepoListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
