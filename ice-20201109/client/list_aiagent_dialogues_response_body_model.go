// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentDialoguesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDialogues(v []*ListAIAgentDialoguesResponseBodyDialogues) *ListAIAgentDialoguesResponseBody
	GetDialogues() []*ListAIAgentDialoguesResponseBodyDialogues
	SetRequestId(v string) *ListAIAgentDialoguesResponseBody
	GetRequestId() *string
}

type ListAIAgentDialoguesResponseBody struct {
	// The dialog records.
	Dialogues []*ListAIAgentDialoguesResponseBodyDialogues `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7B117AF5-***************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIAgentDialoguesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentDialoguesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIAgentDialoguesResponseBody) GetDialogues() []*ListAIAgentDialoguesResponseBodyDialogues {
	return s.Dialogues
}

func (s *ListAIAgentDialoguesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIAgentDialoguesResponseBody) SetDialogues(v []*ListAIAgentDialoguesResponseBodyDialogues) *ListAIAgentDialoguesResponseBody {
	s.Dialogues = v
	return s
}

func (s *ListAIAgentDialoguesResponseBody) SetRequestId(v string) *ListAIAgentDialoguesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBody) Validate() error {
	if s.Dialogues != nil {
		for _, item := range s.Dialogues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAIAgentDialoguesResponseBodyDialogues struct {
	AttachedFileList []*ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList `json:"AttachedFileList,omitempty" xml:"AttachedFileList,omitempty" type:"Repeated"`
	// The unique ID of the dialog.
	//
	// example:
	//
	// 19de81b3b3d94abda22****
	DialogueId *string `json:"DialogueId,omitempty" xml:"DialogueId,omitempty"`
	Extend     *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The speaker. Valid values:
	//
	// - user
	//
	// - agent
	//
	// example:
	//
	// user
	Producer *string `json:"Producer,omitempty" xml:"Producer,omitempty"`
	// The reasoning trace.
	//
	// example:
	//
	// I\\"m thinking
	ReasoningText *string `json:"ReasoningText,omitempty" xml:"ReasoningText,omitempty"`
	// The ID of the conversational turn.
	//
	// example:
	//
	// f27f9b9be28642a88e18****
	RoundId *string `json:"RoundId,omitempty" xml:"RoundId,omitempty"`
	// The source of the message. Valid values:
	//
	// chat: messaging conversations.
	//
	// call: voice calls.
	//
	// example:
	//
	// chat
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The specific content.
	//
	// example:
	//
	// Hello
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The UNIX timestamp, measured in milliseconds, which indicates the time when the message was generated.
	//
	// example:
	//
	// 1734511087000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The message type. Valid values:
	//
	// Voice calls:
	//
	// 1.  greeting: the welcome message.
	//
	// 2.  normal: the voice response.
	//
	// 3.  speech: the proactive message.
	//
	// Messaging conversations:
	//
	// 1.  normal: the text reply.
	//
	// 2.  announcement: the proactive text message.
	//
	// 3.  custom: the custom message.
	//
	// example:
	//
	// announcement
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAIAgentDialoguesResponseBodyDialogues) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentDialoguesResponseBodyDialogues) GoString() string {
	return s.String()
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetAttachedFileList() []*ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList {
	return s.AttachedFileList
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetDialogueId() *string {
	return s.DialogueId
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetExtend() *string {
	return s.Extend
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetNodeId() *string {
	return s.NodeId
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetProducer() *string {
	return s.Producer
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetReasoningText() *string {
	return s.ReasoningText
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetRoundId() *string {
	return s.RoundId
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetSource() *string {
	return s.Source
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetText() *string {
	return s.Text
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetTime() *int64 {
	return s.Time
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) GetType() *string {
	return s.Type
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetAttachedFileList(v []*ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) *ListAIAgentDialoguesResponseBodyDialogues {
	s.AttachedFileList = v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetDialogueId(v string) *ListAIAgentDialoguesResponseBodyDialogues {
	s.DialogueId = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetExtend(v string) *ListAIAgentDialoguesResponseBodyDialogues {
	s.Extend = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetNodeId(v string) *ListAIAgentDialoguesResponseBodyDialogues {
	s.NodeId = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetProducer(v string) *ListAIAgentDialoguesResponseBodyDialogues {
	s.Producer = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetReasoningText(v string) *ListAIAgentDialoguesResponseBodyDialogues {
	s.ReasoningText = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetRoundId(v string) *ListAIAgentDialoguesResponseBodyDialogues {
	s.RoundId = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetSource(v string) *ListAIAgentDialoguesResponseBodyDialogues {
	s.Source = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetText(v string) *ListAIAgentDialoguesResponseBodyDialogues {
	s.Text = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetTime(v int64) *ListAIAgentDialoguesResponseBodyDialogues {
	s.Time = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) SetType(v string) *ListAIAgentDialoguesResponseBodyDialogues {
	s.Type = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialogues) Validate() error {
	if s.AttachedFileList != nil {
		for _, item := range s.AttachedFileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList struct {
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type   *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) GoString() string {
	return s.String()
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) GetFormat() *string {
	return s.Format
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) GetId() *string {
	return s.Id
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) GetName() *string {
	return s.Name
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) GetType() *int32 {
	return s.Type
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) GetUrl() *string {
	return s.Url
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) SetFormat(v string) *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList {
	s.Format = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) SetId(v string) *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList {
	s.Id = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) SetName(v string) *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList {
	s.Name = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) SetType(v int32) *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList {
	s.Type = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) SetUrl(v string) *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList {
	s.Url = &v
	return s
}

func (s *ListAIAgentDialoguesResponseBodyDialoguesAttachedFileList) Validate() error {
	return dara.Validate(s)
}
