// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListGroupResponseBody
	GetCount() *int64
	SetGroups(v []*ListGroupResponseBodyGroups) *ListGroupResponseBody
	GetGroups() []*ListGroupResponseBodyGroups
	SetPageNumber(v int64) *ListGroupResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListGroupResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListGroupResponseBody
	GetRequestId() *string
}

type ListGroupResponseBody struct {
	// example:
	//
	// 3
	Count  *int64                         `json:"count,omitempty" xml:"count,omitempty"`
	Groups []*ListGroupResponseBodyGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListGroupResponseBody) GetGroups() []*ListGroupResponseBodyGroups {
	return s.Groups
}

func (s *ListGroupResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListGroupResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupResponseBody) SetCount(v int64) *ListGroupResponseBody {
	s.Count = &v
	return s
}

func (s *ListGroupResponseBody) SetGroups(v []*ListGroupResponseBodyGroups) *ListGroupResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupResponseBody) SetPageNumber(v int64) *ListGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGroupResponseBody) SetPageSize(v int64) *ListGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGroupResponseBody) SetRequestId(v string) *ListGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupResponseBodyGroups struct {
	// example:
	//
	// 2022-09-14T07:19:13Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// g-148e7853433574fffe9fec72ed9b73
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// example:
	//
	// 1
	ModuleCnt *int64 `json:"moduleCnt,omitempty" xml:"moduleCnt,omitempty"`
	// example:
	//
	// 1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-4267dcfbf1b6d126edcadf0e949
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 1
	SceneTestingTaskCnt *int64                             `json:"sceneTestingTaskCnt,omitempty" xml:"sceneTestingTaskCnt,omitempty"`
	Tags                []*ListGroupResponseBodyGroupsTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	TaskCnt *int64 `json:"taskCnt,omitempty" xml:"taskCnt,omitempty"`
}

func (s ListGroupResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupResponseBodyGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListGroupResponseBodyGroups) GetDescription() *string {
	return s.Description
}

func (s *ListGroupResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupResponseBodyGroups) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListGroupResponseBodyGroups) GetModuleCnt() *int64 {
	return s.ModuleCnt
}

func (s *ListGroupResponseBodyGroups) GetName() *string {
	return s.Name
}

func (s *ListGroupResponseBodyGroups) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListGroupResponseBodyGroups) GetSceneTestingTaskCnt() *int64 {
	return s.SceneTestingTaskCnt
}

func (s *ListGroupResponseBodyGroups) GetTags() []*ListGroupResponseBodyGroupsTags {
	return s.Tags
}

func (s *ListGroupResponseBodyGroups) GetTaskCnt() *int64 {
	return s.TaskCnt
}

func (s *ListGroupResponseBodyGroups) SetCreateTime(v string) *ListGroupResponseBodyGroups {
	s.CreateTime = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetDescription(v string) *ListGroupResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetGroupId(v string) *ListGroupResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetIsDefault(v bool) *ListGroupResponseBodyGroups {
	s.IsDefault = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetModuleCnt(v int64) *ListGroupResponseBodyGroups {
	s.ModuleCnt = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetName(v string) *ListGroupResponseBodyGroups {
	s.Name = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetProjectId(v string) *ListGroupResponseBodyGroups {
	s.ProjectId = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetSceneTestingTaskCnt(v int64) *ListGroupResponseBodyGroups {
	s.SceneTestingTaskCnt = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetTags(v []*ListGroupResponseBodyGroupsTags) *ListGroupResponseBodyGroups {
	s.Tags = v
	return s
}

func (s *ListGroupResponseBodyGroups) SetTaskCnt(v int64) *ListGroupResponseBodyGroups {
	s.TaskCnt = &v
	return s
}

func (s *ListGroupResponseBodyGroups) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupResponseBodyGroupsTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGroupResponseBodyGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s ListGroupResponseBodyGroupsTags) GoString() string {
	return s.String()
}

func (s *ListGroupResponseBodyGroupsTags) GetKey() *string {
	return s.Key
}

func (s *ListGroupResponseBodyGroupsTags) GetValue() *string {
	return s.Value
}

func (s *ListGroupResponseBodyGroupsTags) SetKey(v string) *ListGroupResponseBodyGroupsTags {
	s.Key = &v
	return s
}

func (s *ListGroupResponseBodyGroupsTags) SetValue(v string) *ListGroupResponseBodyGroupsTags {
	s.Value = &v
	return s
}

func (s *ListGroupResponseBodyGroupsTags) Validate() error {
	return dara.Validate(s)
}
