// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRuleTestDialogue interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*RuleTestDialogueContent) *RuleTestDialogue
	GetContent() []*RuleTestDialogueContent
	SetId(v int64) *RuleTestDialogue
	GetId() *int64
	SetName(v string) *RuleTestDialogue
	GetName() *string
	SetUserGroup(v string) *RuleTestDialogue
	GetUserGroup() *string
}

type RuleTestDialogue struct {
	Content   []*RuleTestDialogueContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	Id        *int64                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string                    `json:"Name,omitempty" xml:"Name,omitempty"`
	UserGroup *string                    `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
}

func (s RuleTestDialogue) String() string {
	return dara.Prettify(s)
}

func (s RuleTestDialogue) GoString() string {
	return s.String()
}

func (s *RuleTestDialogue) GetContent() []*RuleTestDialogueContent {
	return s.Content
}

func (s *RuleTestDialogue) GetId() *int64 {
	return s.Id
}

func (s *RuleTestDialogue) GetName() *string {
	return s.Name
}

func (s *RuleTestDialogue) GetUserGroup() *string {
	return s.UserGroup
}

func (s *RuleTestDialogue) SetContent(v []*RuleTestDialogueContent) *RuleTestDialogue {
	s.Content = v
	return s
}

func (s *RuleTestDialogue) SetId(v int64) *RuleTestDialogue {
	s.Id = &v
	return s
}

func (s *RuleTestDialogue) SetName(v string) *RuleTestDialogue {
	s.Name = &v
	return s
}

func (s *RuleTestDialogue) SetUserGroup(v string) *RuleTestDialogue {
	s.UserGroup = &v
	return s
}

func (s *RuleTestDialogue) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RuleTestDialogueContent struct {
	Begin           *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	BeginTime       *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EmotionValue    *int32  `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End             *int64  `json:"End,omitempty" xml:"End,omitempty"`
	HourMinSec      *string `json:"HourMinSec,omitempty" xml:"HourMinSec,omitempty"`
	Identity        *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SilenceDuration *int64  `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	SpeechRate      *int64  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words           *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s RuleTestDialogueContent) String() string {
	return dara.Prettify(s)
}

func (s RuleTestDialogueContent) GoString() string {
	return s.String()
}

func (s *RuleTestDialogueContent) GetBegin() *int64 {
	return s.Begin
}

func (s *RuleTestDialogueContent) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *RuleTestDialogueContent) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *RuleTestDialogueContent) GetEnd() *int64 {
	return s.End
}

func (s *RuleTestDialogueContent) GetHourMinSec() *string {
	return s.HourMinSec
}

func (s *RuleTestDialogueContent) GetIdentity() *string {
	return s.Identity
}

func (s *RuleTestDialogueContent) GetRole() *string {
	return s.Role
}

func (s *RuleTestDialogueContent) GetSilenceDuration() *int64 {
	return s.SilenceDuration
}

func (s *RuleTestDialogueContent) GetSpeechRate() *int64 {
	return s.SpeechRate
}

func (s *RuleTestDialogueContent) GetWords() *string {
	return s.Words
}

func (s *RuleTestDialogueContent) SetBegin(v int64) *RuleTestDialogueContent {
	s.Begin = &v
	return s
}

func (s *RuleTestDialogueContent) SetBeginTime(v int64) *RuleTestDialogueContent {
	s.BeginTime = &v
	return s
}

func (s *RuleTestDialogueContent) SetEmotionValue(v int32) *RuleTestDialogueContent {
	s.EmotionValue = &v
	return s
}

func (s *RuleTestDialogueContent) SetEnd(v int64) *RuleTestDialogueContent {
	s.End = &v
	return s
}

func (s *RuleTestDialogueContent) SetHourMinSec(v string) *RuleTestDialogueContent {
	s.HourMinSec = &v
	return s
}

func (s *RuleTestDialogueContent) SetIdentity(v string) *RuleTestDialogueContent {
	s.Identity = &v
	return s
}

func (s *RuleTestDialogueContent) SetRole(v string) *RuleTestDialogueContent {
	s.Role = &v
	return s
}

func (s *RuleTestDialogueContent) SetSilenceDuration(v int64) *RuleTestDialogueContent {
	s.SilenceDuration = &v
	return s
}

func (s *RuleTestDialogueContent) SetSpeechRate(v int64) *RuleTestDialogueContent {
	s.SpeechRate = &v
	return s
}

func (s *RuleTestDialogueContent) SetWords(v string) *RuleTestDialogueContent {
	s.Words = &v
	return s
}

func (s *RuleTestDialogueContent) Validate() error {
	return dara.Validate(s)
}
