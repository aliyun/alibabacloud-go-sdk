// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTask interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *TextTask
	GetAgentId() *string
	SetAgentName(v string) *TextTask
	GetAgentName() *string
	SetContentRequirement(v string) *TextTask
	GetContentRequirement() *string
	SetGmtCreate(v string) *TextTask
	GetGmtCreate() *string
	SetGmtModified(v string) *TextTask
	GetGmtModified() *string
	SetIntroduction(v string) *TextTask
	GetIntroduction() *string
	SetNums(v int32) *TextTask
	GetNums() *int32
	SetPoint(v string) *TextTask
	GetPoint() *string
	SetReferenceTag(v *ReferenceTag) *TextTask
	GetReferenceTag() *ReferenceTag
	SetRelatedRagIds(v []*int64) *TextTask
	GetRelatedRagIds() []*int64
	SetStyle(v string) *TextTask
	GetStyle() *string
	SetTarget(v string) *TextTask
	GetTarget() *string
	SetTextIds(v []*int64) *TextTask
	GetTextIds() []*int64
	SetTextModeType(v string) *TextTask
	GetTextModeType() *string
	SetTextTaskId(v int64) *TextTask
	GetTextTaskId() *int64
	SetTextTaskStatus(v string) *TextTask
	GetTextTaskStatus() *string
	SetTexts(v []*Text) *TextTask
	GetTexts() []*Text
	SetTheme(v string) *TextTask
	GetTheme() *string
	SetThemeDesc(v string) *TextTask
	GetThemeDesc() *string
}

type TextTask struct {
	AgentId   *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// example:
	//
	// 九寨沟三日游攻略
	ContentRequirement *string `json:"contentRequirement,omitempty" xml:"contentRequirement,omitempty"`
	GmtCreate          *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Introduction       *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Nums *int32 `json:"nums,omitempty" xml:"nums,omitempty"`
	// example:
	//
	// xxx
	Point         *string       `json:"point,omitempty" xml:"point,omitempty"`
	ReferenceTag  *ReferenceTag `json:"referenceTag,omitempty" xml:"referenceTag,omitempty"`
	RelatedRagIds []*int64      `json:"relatedRagIds,omitempty" xml:"relatedRagIds,omitempty" type:"Repeated"`
	// This parameter is required.
	Style   *string  `json:"style,omitempty" xml:"style,omitempty"`
	Target  *string  `json:"target,omitempty" xml:"target,omitempty"`
	TextIds []*int64 `json:"textIds,omitempty" xml:"textIds,omitempty" type:"Repeated"`
	// This parameter is required.
	TextModeType   *string `json:"textModeType,omitempty" xml:"textModeType,omitempty"`
	TextTaskId     *int64  `json:"textTaskId,omitempty" xml:"textTaskId,omitempty"`
	TextTaskStatus *string `json:"textTaskStatus,omitempty" xml:"textTaskStatus,omitempty"`
	Texts          []*Text `json:"texts,omitempty" xml:"texts,omitempty" type:"Repeated"`
	// example:
	//
	// 旅游路线
	Theme     *string `json:"theme,omitempty" xml:"theme,omitempty"`
	ThemeDesc *string `json:"themeDesc,omitempty" xml:"themeDesc,omitempty"`
}

func (s TextTask) String() string {
	return dara.Prettify(s)
}

func (s TextTask) GoString() string {
	return s.String()
}

func (s *TextTask) GetAgentId() *string {
	return s.AgentId
}

func (s *TextTask) GetAgentName() *string {
	return s.AgentName
}

func (s *TextTask) GetContentRequirement() *string {
	return s.ContentRequirement
}

func (s *TextTask) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TextTask) GetGmtModified() *string {
	return s.GmtModified
}

func (s *TextTask) GetIntroduction() *string {
	return s.Introduction
}

func (s *TextTask) GetNums() *int32 {
	return s.Nums
}

func (s *TextTask) GetPoint() *string {
	return s.Point
}

func (s *TextTask) GetReferenceTag() *ReferenceTag {
	return s.ReferenceTag
}

func (s *TextTask) GetRelatedRagIds() []*int64 {
	return s.RelatedRagIds
}

func (s *TextTask) GetStyle() *string {
	return s.Style
}

func (s *TextTask) GetTarget() *string {
	return s.Target
}

func (s *TextTask) GetTextIds() []*int64 {
	return s.TextIds
}

func (s *TextTask) GetTextModeType() *string {
	return s.TextModeType
}

func (s *TextTask) GetTextTaskId() *int64 {
	return s.TextTaskId
}

func (s *TextTask) GetTextTaskStatus() *string {
	return s.TextTaskStatus
}

func (s *TextTask) GetTexts() []*Text {
	return s.Texts
}

func (s *TextTask) GetTheme() *string {
	return s.Theme
}

func (s *TextTask) GetThemeDesc() *string {
	return s.ThemeDesc
}

func (s *TextTask) SetAgentId(v string) *TextTask {
	s.AgentId = &v
	return s
}

func (s *TextTask) SetAgentName(v string) *TextTask {
	s.AgentName = &v
	return s
}

func (s *TextTask) SetContentRequirement(v string) *TextTask {
	s.ContentRequirement = &v
	return s
}

func (s *TextTask) SetGmtCreate(v string) *TextTask {
	s.GmtCreate = &v
	return s
}

func (s *TextTask) SetGmtModified(v string) *TextTask {
	s.GmtModified = &v
	return s
}

func (s *TextTask) SetIntroduction(v string) *TextTask {
	s.Introduction = &v
	return s
}

func (s *TextTask) SetNums(v int32) *TextTask {
	s.Nums = &v
	return s
}

func (s *TextTask) SetPoint(v string) *TextTask {
	s.Point = &v
	return s
}

func (s *TextTask) SetReferenceTag(v *ReferenceTag) *TextTask {
	s.ReferenceTag = v
	return s
}

func (s *TextTask) SetRelatedRagIds(v []*int64) *TextTask {
	s.RelatedRagIds = v
	return s
}

func (s *TextTask) SetStyle(v string) *TextTask {
	s.Style = &v
	return s
}

func (s *TextTask) SetTarget(v string) *TextTask {
	s.Target = &v
	return s
}

func (s *TextTask) SetTextIds(v []*int64) *TextTask {
	s.TextIds = v
	return s
}

func (s *TextTask) SetTextModeType(v string) *TextTask {
	s.TextModeType = &v
	return s
}

func (s *TextTask) SetTextTaskId(v int64) *TextTask {
	s.TextTaskId = &v
	return s
}

func (s *TextTask) SetTextTaskStatus(v string) *TextTask {
	s.TextTaskStatus = &v
	return s
}

func (s *TextTask) SetTexts(v []*Text) *TextTask {
	s.Texts = v
	return s
}

func (s *TextTask) SetTheme(v string) *TextTask {
	s.Theme = &v
	return s
}

func (s *TextTask) SetThemeDesc(v string) *TextTask {
	s.ThemeDesc = &v
	return s
}

func (s *TextTask) Validate() error {
	if s.ReferenceTag != nil {
		if err := s.ReferenceTag.Validate(); err != nil {
			return err
		}
	}
	if s.Texts != nil {
		for _, item := range s.Texts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
