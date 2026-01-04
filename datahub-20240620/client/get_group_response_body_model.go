// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *GetGroupResponseBody
	GetComment() *string
	SetCreateTime(v int64) *GetGroupResponseBody
	GetCreateTime() *int64
	SetCreator(v string) *GetGroupResponseBody
	GetCreator() *string
	SetGroupName(v string) *GetGroupResponseBody
	GetGroupName() *string
	SetProjectName(v string) *GetGroupResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGroupResponseBody
	GetSuccess() *bool
	SetTopicList(v []*string) *GetGroupResponseBody
	GetTopicList() []*string
	SetUpdateTime(v int64) *GetGroupResponseBody
	GetUpdateTime() *int64
}

type GetGroupResponseBody struct {
	// example:
	//
	// test_comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1724041098000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 270523390948438349
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// consumer_group1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	TopicList []*string `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Repeated"`
	// example:
	//
	// 1724041098000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) GetComment() *string {
	return s.Comment
}

func (s *GetGroupResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetGroupResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGroupResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGroupResponseBody) GetTopicList() []*string {
	return s.TopicList
}

func (s *GetGroupResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetGroupResponseBody) SetComment(v string) *GetGroupResponseBody {
	s.Comment = &v
	return s
}

func (s *GetGroupResponseBody) SetCreateTime(v int64) *GetGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetGroupResponseBody) SetCreator(v string) *GetGroupResponseBody {
	s.Creator = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupName(v string) *GetGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *GetGroupResponseBody) SetProjectName(v string) *GetGroupResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetGroupResponseBody) SetRequestId(v string) *GetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupResponseBody) SetSuccess(v bool) *GetGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetGroupResponseBody) SetTopicList(v []*string) *GetGroupResponseBody {
	s.TopicList = v
	return s
}

func (s *GetGroupResponseBody) SetUpdateTime(v int64) *GetGroupResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
