// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTaskCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *TextTaskCreateCmd
	GetAgentId() *string
	SetContentRequirement(v string) *TextTaskCreateCmd
	GetContentRequirement() *string
	SetIdempotentId(v string) *TextTaskCreateCmd
	GetIdempotentId() *string
	SetIndustry(v string) *TextTaskCreateCmd
	GetIndustry() *string
	SetIntroduction(v string) *TextTaskCreateCmd
	GetIntroduction() *string
	SetNumber(v int32) *TextTaskCreateCmd
	GetNumber() *int32
	SetPoint(v string) *TextTaskCreateCmd
	GetPoint() *string
	SetReferenceTag(v *ReferenceTag) *TextTaskCreateCmd
	GetReferenceTag() *ReferenceTag
	SetRelatedRagIds(v []*int64) *TextTaskCreateCmd
	GetRelatedRagIds() []*int64
	SetStreamApi(v bool) *TextTaskCreateCmd
	GetStreamApi() *bool
	SetStyle(v string) *TextTaskCreateCmd
	GetStyle() *string
	SetTarget(v string) *TextTaskCreateCmd
	GetTarget() *string
	SetTextModeType(v string) *TextTaskCreateCmd
	GetTextModeType() *string
	SetTheme(v string) *TextTaskCreateCmd
	GetTheme() *string
	SetThemes(v []*string) *TextTaskCreateCmd
	GetThemes() []*string
}

type TextTaskCreateCmd struct {
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 极氪007新车上市
	ContentRequirement *string `json:"contentRequirement,omitempty" xml:"contentRequirement,omitempty"`
	// example:
	//
	// 28274623764834
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	Industry     *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// example:
	//
	// xxx
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// example:
	//
	// 超强续航
	Point        *string       `json:"point,omitempty" xml:"point,omitempty"`
	ReferenceTag *ReferenceTag `json:"referenceTag,omitempty" xml:"referenceTag,omitempty"`
	// example:
	//
	// 1
	RelatedRagIds []*int64 `json:"relatedRagIds,omitempty" xml:"relatedRagIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	StreamApi *bool `json:"streamApi,omitempty" xml:"streamApi,omitempty"`
	// This parameter is required.
	Style  *string `json:"style,omitempty" xml:"style,omitempty"`
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// This parameter is required.
	TextModeType *string `json:"textModeType,omitempty" xml:"textModeType,omitempty"`
	// example:
	//
	// 旅游路线
	Theme  *string   `json:"theme,omitempty" xml:"theme,omitempty"`
	Themes []*string `json:"themes,omitempty" xml:"themes,omitempty" type:"Repeated"`
}

func (s TextTaskCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s TextTaskCreateCmd) GoString() string {
	return s.String()
}

func (s *TextTaskCreateCmd) GetAgentId() *string {
	return s.AgentId
}

func (s *TextTaskCreateCmd) GetContentRequirement() *string {
	return s.ContentRequirement
}

func (s *TextTaskCreateCmd) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *TextTaskCreateCmd) GetIndustry() *string {
	return s.Industry
}

func (s *TextTaskCreateCmd) GetIntroduction() *string {
	return s.Introduction
}

func (s *TextTaskCreateCmd) GetNumber() *int32 {
	return s.Number
}

func (s *TextTaskCreateCmd) GetPoint() *string {
	return s.Point
}

func (s *TextTaskCreateCmd) GetReferenceTag() *ReferenceTag {
	return s.ReferenceTag
}

func (s *TextTaskCreateCmd) GetRelatedRagIds() []*int64 {
	return s.RelatedRagIds
}

func (s *TextTaskCreateCmd) GetStreamApi() *bool {
	return s.StreamApi
}

func (s *TextTaskCreateCmd) GetStyle() *string {
	return s.Style
}

func (s *TextTaskCreateCmd) GetTarget() *string {
	return s.Target
}

func (s *TextTaskCreateCmd) GetTextModeType() *string {
	return s.TextModeType
}

func (s *TextTaskCreateCmd) GetTheme() *string {
	return s.Theme
}

func (s *TextTaskCreateCmd) GetThemes() []*string {
	return s.Themes
}

func (s *TextTaskCreateCmd) SetAgentId(v string) *TextTaskCreateCmd {
	s.AgentId = &v
	return s
}

func (s *TextTaskCreateCmd) SetContentRequirement(v string) *TextTaskCreateCmd {
	s.ContentRequirement = &v
	return s
}

func (s *TextTaskCreateCmd) SetIdempotentId(v string) *TextTaskCreateCmd {
	s.IdempotentId = &v
	return s
}

func (s *TextTaskCreateCmd) SetIndustry(v string) *TextTaskCreateCmd {
	s.Industry = &v
	return s
}

func (s *TextTaskCreateCmd) SetIntroduction(v string) *TextTaskCreateCmd {
	s.Introduction = &v
	return s
}

func (s *TextTaskCreateCmd) SetNumber(v int32) *TextTaskCreateCmd {
	s.Number = &v
	return s
}

func (s *TextTaskCreateCmd) SetPoint(v string) *TextTaskCreateCmd {
	s.Point = &v
	return s
}

func (s *TextTaskCreateCmd) SetReferenceTag(v *ReferenceTag) *TextTaskCreateCmd {
	s.ReferenceTag = v
	return s
}

func (s *TextTaskCreateCmd) SetRelatedRagIds(v []*int64) *TextTaskCreateCmd {
	s.RelatedRagIds = v
	return s
}

func (s *TextTaskCreateCmd) SetStreamApi(v bool) *TextTaskCreateCmd {
	s.StreamApi = &v
	return s
}

func (s *TextTaskCreateCmd) SetStyle(v string) *TextTaskCreateCmd {
	s.Style = &v
	return s
}

func (s *TextTaskCreateCmd) SetTarget(v string) *TextTaskCreateCmd {
	s.Target = &v
	return s
}

func (s *TextTaskCreateCmd) SetTextModeType(v string) *TextTaskCreateCmd {
	s.TextModeType = &v
	return s
}

func (s *TextTaskCreateCmd) SetTheme(v string) *TextTaskCreateCmd {
	s.Theme = &v
	return s
}

func (s *TextTaskCreateCmd) SetThemes(v []*string) *TextTaskCreateCmd {
	s.Themes = v
	return s
}

func (s *TextTaskCreateCmd) Validate() error {
	if s.ReferenceTag != nil {
		if err := s.ReferenceTag.Validate(); err != nil {
			return err
		}
	}
	return nil
}
