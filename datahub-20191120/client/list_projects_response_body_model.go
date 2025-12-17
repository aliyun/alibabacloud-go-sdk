// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v *ListProjectsResponseBodyList) *ListProjectsResponseBody
	GetList() *ListProjectsResponseBodyList
	SetPageCount(v string) *ListProjectsResponseBody
	GetPageCount() *string
	SetPageNumber(v string) *ListProjectsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *ListProjectsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectsResponseBody
	GetSuccess() *bool
	SetTotalCount(v string) *ListProjectsResponseBody
	GetTotalCount() *string
}

type ListProjectsResponseBody struct {
	// Project name
	List *ListProjectsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// Total number of pages
	//
	// example:
	//
	// 1
	PageCount *string `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// Current page number
	//
	// example:
	//
	// 84
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 20250401102332e68e3d0b04ab4904
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of query results
	//
	// example:
	//
	// 50
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetList() *ListProjectsResponseBodyList {
	return s.List
}

func (s *ListProjectsResponseBody) GetPageCount() *string {
	return s.PageCount
}

func (s *ListProjectsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListProjectsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListProjectsResponseBody) SetList(v *ListProjectsResponseBodyList) *ListProjectsResponseBody {
	s.List = v
	return s
}

func (s *ListProjectsResponseBody) SetPageCount(v string) *ListProjectsResponseBody {
	s.PageCount = &v
	return s
}

func (s *ListProjectsResponseBody) SetPageNumber(v string) *ListProjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsResponseBody) SetPageSize(v string) *ListProjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetSuccess(v bool) *ListProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalCount(v string) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyList struct {
	Project []*ListProjectsResponseBodyListProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Repeated"`
}

func (s ListProjectsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyList) GetProject() []*ListProjectsResponseBodyListProject {
	return s.Project
}

func (s *ListProjectsResponseBodyList) SetProject(v []*ListProjectsResponseBodyListProject) *ListProjectsResponseBodyList {
	s.Project = v
	return s
}

func (s *ListProjectsResponseBodyList) Validate() error {
	if s.Project != nil {
		for _, item := range s.Project {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsResponseBodyListProject struct {
	// Text comment for the field
	//
	// example:
	//
	// 测试专用
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Creation time
	//
	// example:
	//
	// 1708171905000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Creator ID
	//
	// example:
	//
	// 1048133943212399
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Creator\\"s name
	//
	// if can be null:
	// true
	//
	// example:
	//
	// user_name
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// Project name
	//
	// example:
	//
	// poc_test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Update time
	//
	// example:
	//
	// 1708171905000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListProjectsResponseBodyListProject) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyListProject) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyListProject) GetComment() *string {
	return s.Comment
}

func (s *ListProjectsResponseBodyListProject) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListProjectsResponseBodyListProject) GetCreator() *string {
	return s.Creator
}

func (s *ListProjectsResponseBodyListProject) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListProjectsResponseBodyListProject) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListProjectsResponseBodyListProject) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListProjectsResponseBodyListProject) SetComment(v string) *ListProjectsResponseBodyListProject {
	s.Comment = &v
	return s
}

func (s *ListProjectsResponseBodyListProject) SetCreateTime(v int64) *ListProjectsResponseBodyListProject {
	s.CreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyListProject) SetCreator(v string) *ListProjectsResponseBodyListProject {
	s.Creator = &v
	return s
}

func (s *ListProjectsResponseBodyListProject) SetCreatorName(v string) *ListProjectsResponseBodyListProject {
	s.CreatorName = &v
	return s
}

func (s *ListProjectsResponseBodyListProject) SetProjectName(v string) *ListProjectsResponseBodyListProject {
	s.ProjectName = &v
	return s
}

func (s *ListProjectsResponseBodyListProject) SetUpdateTime(v int64) *ListProjectsResponseBodyListProject {
	s.UpdateTime = &v
	return s
}

func (s *ListProjectsResponseBodyListProject) Validate() error {
	return dara.Validate(s)
}
