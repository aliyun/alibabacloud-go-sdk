// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemoryDetailResponseBody
	GetCode() *string
	SetGroups(v []*GetMemoryDetailResponseBodyGroups) *GetMemoryDetailResponseBody
	GetGroups() []*GetMemoryDetailResponseBodyGroups
	SetMessage(v string) *GetMemoryDetailResponseBody
	GetMessage() *string
	SetProfile(v *GetMemoryDetailResponseBodyProfile) *GetMemoryDetailResponseBody
	GetProfile() *GetMemoryDetailResponseBodyProfile
	SetRequestId(v string) *GetMemoryDetailResponseBody
	GetRequestId() *string
}

type GetMemoryDetailResponseBody struct {
	// 业务状态码；成功为 200，失败时为对应错误码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 自动拾取记忆的分组列表；无自动记忆时为空数组
	//
	// example:
	//
	// [{"type":"user","items":[{"id":"user/example.md","title":"示例偏好","content":"喜欢简洁的回答","description":"来自近期会话的偏好总结"}],"gmtModified":1720000000}]
	Groups []*GetMemoryDetailResponseBodyGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	// 错误描述；成功时为空字符串，失败时返回具体原因
	//
	// example:
	//
	// 当前用户不是群组 cg_12345 的活跃成员
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 手动添加的记忆，来自 profile.md；包含完整内容、存在状态和修改时间
	//
	// example:
	//
	// {"content":"示例手动记忆","exists":true,"gmtModified":1720000000}
	Profile *GetMemoryDetailResponseBodyProfile `json:"profile,omitempty" xml:"profile,omitempty" type:"Struct"`
	// 请求追踪 ID，用于定位本次调用
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMemoryDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemoryDetailResponseBody) GetGroups() []*GetMemoryDetailResponseBodyGroups {
	return s.Groups
}

func (s *GetMemoryDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMemoryDetailResponseBody) GetProfile() *GetMemoryDetailResponseBodyProfile {
	return s.Profile
}

func (s *GetMemoryDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryDetailResponseBody) SetCode(v string) *GetMemoryDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemoryDetailResponseBody) SetGroups(v []*GetMemoryDetailResponseBodyGroups) *GetMemoryDetailResponseBody {
	s.Groups = v
	return s
}

func (s *GetMemoryDetailResponseBody) SetMessage(v string) *GetMemoryDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetMemoryDetailResponseBody) SetProfile(v *GetMemoryDetailResponseBodyProfile) *GetMemoryDetailResponseBody {
	s.Profile = v
	return s
}

func (s *GetMemoryDetailResponseBody) SetRequestId(v string) *GetMemoryDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryDetailResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryDetailResponseBodyGroups struct {
	// 该分组最近修改时间，Unix 秒级时间戳；无文件时为 null
	//
	// example:
	//
	// 1720000000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 该类型下的自动记忆条目列表；没有条目时为空数组
	//
	// example:
	//
	// [{"id":"user/example.md","title":"示例偏好","content":"喜欢简洁的回答","description":"来自近期会话的偏好总结"}]
	Items []*GetMemoryDetailResponseBodyGroupsItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 自动记忆类型：user、feedback、reference 或 project
	//
	// example:
	//
	// user
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMemoryDetailResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailResponseBodyGroups) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetMemoryDetailResponseBodyGroups) GetItems() []*GetMemoryDetailResponseBodyGroupsItems {
	return s.Items
}

func (s *GetMemoryDetailResponseBodyGroups) GetType() *string {
	return s.Type
}

func (s *GetMemoryDetailResponseBodyGroups) SetGmtModified(v int64) *GetMemoryDetailResponseBodyGroups {
	s.GmtModified = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroups) SetItems(v []*GetMemoryDetailResponseBodyGroupsItems) *GetMemoryDetailResponseBodyGroups {
	s.Items = v
	return s
}

func (s *GetMemoryDetailResponseBodyGroups) SetType(v string) *GetMemoryDetailResponseBodyGroups {
	s.Type = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroups) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMemoryDetailResponseBodyGroupsItems struct {
	// 自动记忆正文，不包含 YAML frontmatter
	//
	// example:
	//
	// 喜欢简洁的回答
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 自动记忆的附加描述；没有时为空字符串
	//
	// example:
	//
	// 来自近期会话的偏好总结
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 自动记忆条目的唯一标识，为相对 OSS 路径
	//
	// example:
	//
	// user/example.md
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 自动记忆标题，来自文件元数据或文件名
	//
	// example:
	//
	// 示例偏好
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetMemoryDetailResponseBodyGroupsItems) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailResponseBodyGroupsItems) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailResponseBodyGroupsItems) GetContent() *string {
	return s.Content
}

func (s *GetMemoryDetailResponseBodyGroupsItems) GetDescription() *string {
	return s.Description
}

func (s *GetMemoryDetailResponseBodyGroupsItems) GetId() *string {
	return s.Id
}

func (s *GetMemoryDetailResponseBodyGroupsItems) GetTitle() *string {
	return s.Title
}

func (s *GetMemoryDetailResponseBodyGroupsItems) SetContent(v string) *GetMemoryDetailResponseBodyGroupsItems {
	s.Content = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroupsItems) SetDescription(v string) *GetMemoryDetailResponseBodyGroupsItems {
	s.Description = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroupsItems) SetId(v string) *GetMemoryDetailResponseBodyGroupsItems {
	s.Id = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroupsItems) SetTitle(v string) *GetMemoryDetailResponseBodyGroupsItems {
	s.Title = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroupsItems) Validate() error {
	return dara.Validate(s)
}

type GetMemoryDetailResponseBodyProfile struct {
	// 手动记忆的完整正文；未保存过时为空字符串
	//
	// example:
	//
	// 示例手动记忆
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// profile.md 是否存在；未保存过时为 false
	//
	// example:
	//
	// true
	Exists *bool `json:"exists,omitempty" xml:"exists,omitempty"`
	// profile.md 最近修改时间，Unix 秒级时间戳；文件不存在或无法取得元数据时为 null
	//
	// example:
	//
	// 1720000000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
}

func (s GetMemoryDetailResponseBodyProfile) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailResponseBodyProfile) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailResponseBodyProfile) GetContent() *string {
	return s.Content
}

func (s *GetMemoryDetailResponseBodyProfile) GetExists() *bool {
	return s.Exists
}

func (s *GetMemoryDetailResponseBodyProfile) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetMemoryDetailResponseBodyProfile) SetContent(v string) *GetMemoryDetailResponseBodyProfile {
	s.Content = &v
	return s
}

func (s *GetMemoryDetailResponseBodyProfile) SetExists(v bool) *GetMemoryDetailResponseBodyProfile {
	s.Exists = &v
	return s
}

func (s *GetMemoryDetailResponseBodyProfile) SetGmtModified(v int64) *GetMemoryDetailResponseBodyProfile {
	s.GmtModified = &v
	return s
}

func (s *GetMemoryDetailResponseBodyProfile) Validate() error {
	return dara.Validate(s)
}
