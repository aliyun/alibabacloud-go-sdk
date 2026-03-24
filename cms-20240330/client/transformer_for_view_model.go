// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformerForView interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*TransformAction) *TransformerForView
	GetActions() []*TransformAction
	SetCreateTime(v string) *TransformerForView
	GetCreateTime() *string
	SetDescription(v string) *TransformerForView
	GetDescription() *string
	SetEnable(v bool) *TransformerForView
	GetEnable() *bool
	SetFilterSetting(v *FilterSetting) *TransformerForView
	GetFilterSetting() *FilterSetting
	SetQuitAfterMatch(v bool) *TransformerForView
	GetQuitAfterMatch() *bool
	SetSortId(v int32) *TransformerForView
	GetSortId() *int32
	SetTransformerId(v string) *TransformerForView
	GetTransformerId() *string
	SetTransformerName(v string) *TransformerForView
	GetTransformerName() *string
	SetUpdateTime(v string) *TransformerForView
	GetUpdateTime() *string
	SetUserId(v string) *TransformerForView
	GetUserId() *string
	SetWorkspace(v string) *TransformerForView
	GetWorkspace() *string
}

type TransformerForView struct {
	// 转换操作
	Actions []*TransformAction `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// 创建时间
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 描述
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 是否启用
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 筛选设置
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// 匹配处理后退出
	//
	// example:
	//
	// false
	QuitAfterMatch *bool `json:"quitAfterMatch,omitempty" xml:"quitAfterMatch,omitempty"`
	// 排序数
	//
	// example:
	//
	// 1
	SortId *int32 `json:"sortId,omitempty" xml:"sortId,omitempty"`
	// transformer Id
	//
	// example:
	//
	// 1123123123123
	TransformerId *string `json:"transformerId,omitempty" xml:"transformerId,omitempty"`
	// 名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试处理流
	TransformerName *string `json:"transformerName,omitempty" xml:"transformerName,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 2025-01-16T02:27:01Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// 用户id
	//
	// example:
	//
	// 123123123**
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// workspace
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s TransformerForView) String() string {
	return dara.Prettify(s)
}

func (s TransformerForView) GoString() string {
	return s.String()
}

func (s *TransformerForView) GetActions() []*TransformAction {
	return s.Actions
}

func (s *TransformerForView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *TransformerForView) GetDescription() *string {
	return s.Description
}

func (s *TransformerForView) GetEnable() *bool {
	return s.Enable
}

func (s *TransformerForView) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *TransformerForView) GetQuitAfterMatch() *bool {
	return s.QuitAfterMatch
}

func (s *TransformerForView) GetSortId() *int32 {
	return s.SortId
}

func (s *TransformerForView) GetTransformerId() *string {
	return s.TransformerId
}

func (s *TransformerForView) GetTransformerName() *string {
	return s.TransformerName
}

func (s *TransformerForView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *TransformerForView) GetUserId() *string {
	return s.UserId
}

func (s *TransformerForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *TransformerForView) SetActions(v []*TransformAction) *TransformerForView {
	s.Actions = v
	return s
}

func (s *TransformerForView) SetCreateTime(v string) *TransformerForView {
	s.CreateTime = &v
	return s
}

func (s *TransformerForView) SetDescription(v string) *TransformerForView {
	s.Description = &v
	return s
}

func (s *TransformerForView) SetEnable(v bool) *TransformerForView {
	s.Enable = &v
	return s
}

func (s *TransformerForView) SetFilterSetting(v *FilterSetting) *TransformerForView {
	s.FilterSetting = v
	return s
}

func (s *TransformerForView) SetQuitAfterMatch(v bool) *TransformerForView {
	s.QuitAfterMatch = &v
	return s
}

func (s *TransformerForView) SetSortId(v int32) *TransformerForView {
	s.SortId = &v
	return s
}

func (s *TransformerForView) SetTransformerId(v string) *TransformerForView {
	s.TransformerId = &v
	return s
}

func (s *TransformerForView) SetTransformerName(v string) *TransformerForView {
	s.TransformerName = &v
	return s
}

func (s *TransformerForView) SetUpdateTime(v string) *TransformerForView {
	s.UpdateTime = &v
	return s
}

func (s *TransformerForView) SetUserId(v string) *TransformerForView {
	s.UserId = &v
	return s
}

func (s *TransformerForView) SetWorkspace(v string) *TransformerForView {
	s.Workspace = &v
	return s
}

func (s *TransformerForView) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FilterSetting != nil {
		if err := s.FilterSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}
