// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListGroupsResponseBodyList) *ListGroupsResponseBody
	GetList() []*ListGroupsResponseBodyList
	SetMaxResults(v int32) *ListGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGroupsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListGroupsResponseBody
	GetTotalCount() *int32
}

type ListGroupsResponseBody struct {
	List []*ListGroupsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 9892074a2a89600ae4b0d5a34fb99a3f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
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

func (s ListGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) GetList() []*ListGroupsResponseBodyList {
	return s.List
}

func (s *ListGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListGroupsResponseBody) SetList(v []*ListGroupsResponseBodyList) *ListGroupsResponseBody {
	s.List = v
	return s
}

func (s *ListGroupsResponseBody) SetMaxResults(v int32) *ListGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsResponseBody) SetNextToken(v string) *ListGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetSuccess(v bool) *ListGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotalCount(v int32) *ListGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsResponseBodyList struct {
	// example:
	//
	// test_comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1708171905000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 276887103073464052
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// test_project
	ProjectName *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TopicList   []*string `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Repeated"`
	// example:
	//
	// 1708171905000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGroupsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyList) GetComment() *string {
	return s.Comment
}

func (s *ListGroupsResponseBodyList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListGroupsResponseBodyList) GetCreator() *string {
	return s.Creator
}

func (s *ListGroupsResponseBodyList) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGroupsResponseBodyList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListGroupsResponseBodyList) GetTopicList() []*string {
	return s.TopicList
}

func (s *ListGroupsResponseBodyList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListGroupsResponseBodyList) SetComment(v string) *ListGroupsResponseBodyList {
	s.Comment = &v
	return s
}

func (s *ListGroupsResponseBodyList) SetCreateTime(v int64) *ListGroupsResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *ListGroupsResponseBodyList) SetCreator(v string) *ListGroupsResponseBodyList {
	s.Creator = &v
	return s
}

func (s *ListGroupsResponseBodyList) SetGroupName(v string) *ListGroupsResponseBodyList {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyList) SetProjectName(v string) *ListGroupsResponseBodyList {
	s.ProjectName = &v
	return s
}

func (s *ListGroupsResponseBodyList) SetTopicList(v []*string) *ListGroupsResponseBodyList {
	s.TopicList = v
	return s
}

func (s *ListGroupsResponseBodyList) SetUpdateTime(v int64) *ListGroupsResponseBodyList {
	s.UpdateTime = &v
	return s
}

func (s *ListGroupsResponseBodyList) Validate() error {
	return dara.Validate(s)
}
