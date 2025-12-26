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
	SetMaxResults(v int32) *ListProjectsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListProjectsResponseBody
	GetTotalCount() *int32
}

type ListProjectsResponseBody struct {
	List *ListProjectsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 9892074a2a89600ae4b0d5a34fb99a3f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 20250401102332e68e3d0b04ab4904
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListProjectsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProjectsResponseBody) SetList(v *ListProjectsResponseBodyList) *ListProjectsResponseBody {
	s.List = v
	return s
}

func (s *ListProjectsResponseBody) SetMaxResults(v int32) *ListProjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
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

func (s *ListProjectsResponseBody) SetTotalCount(v int32) *ListProjectsResponseBody {
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
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1708171905000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1048133943212399
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// poc_test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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
