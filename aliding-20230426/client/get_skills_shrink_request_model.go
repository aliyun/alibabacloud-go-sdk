// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIdsShrink(v string) *GetSkillsShrinkRequest
	GetGroupIdsShrink() *string
	SetSkillIdsShrink(v string) *GetSkillsShrinkRequest
	GetSkillIdsShrink() *string
	SetSourceIdOfAssistantId(v string) *GetSkillsShrinkRequest
	GetSourceIdOfAssistantId() *string
}

type GetSkillsShrinkRequest struct {
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	SkillIdsShrink *string `json:"SkillIds,omitempty" xml:"SkillIds,omitempty"`
	// example:
	//
	// xxx
	SourceIdOfAssistantId *string `json:"SourceIdOfAssistantId,omitempty" xml:"SourceIdOfAssistantId,omitempty"`
}

func (s GetSkillsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillsShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetSkillsShrinkRequest) GetSkillIdsShrink() *string {
	return s.SkillIdsShrink
}

func (s *GetSkillsShrinkRequest) GetSourceIdOfAssistantId() *string {
	return s.SourceIdOfAssistantId
}

func (s *GetSkillsShrinkRequest) SetGroupIdsShrink(v string) *GetSkillsShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillsShrinkRequest) SetSkillIdsShrink(v string) *GetSkillsShrinkRequest {
	s.SkillIdsShrink = &v
	return s
}

func (s *GetSkillsShrinkRequest) SetSourceIdOfAssistantId(v string) *GetSkillsShrinkRequest {
	s.SourceIdOfAssistantId = &v
	return s
}

func (s *GetSkillsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
