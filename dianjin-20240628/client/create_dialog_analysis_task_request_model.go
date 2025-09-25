// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDialogAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisNodes(v []*string) *CreateDialogAnalysisTaskRequest
	GetAnalysisNodes() []*string
	SetConversationList(v []*CreateDialogAnalysisTaskRequestConversationList) *CreateDialogAnalysisTaskRequest
	GetConversationList() []*CreateDialogAnalysisTaskRequestConversationList
	SetMetaData(v map[string]interface{}) *CreateDialogAnalysisTaskRequest
	GetMetaData() map[string]interface{}
	SetPlayCode(v string) *CreateDialogAnalysisTaskRequest
	GetPlayCode() *string
	SetRequestId(v string) *CreateDialogAnalysisTaskRequest
	GetRequestId() *string
}

type CreateDialogAnalysisTaskRequest struct {
	AnalysisNodes []*string `json:"analysisNodes,omitempty" xml:"analysisNodes,omitempty" type:"Repeated"`
	// This parameter is required.
	ConversationList []*CreateDialogAnalysisTaskRequestConversationList `json:"conversationList,omitempty" xml:"conversationList,omitempty" type:"Repeated"`
	// example:
	//
	// {
	//
	// "labels": "XXX",
	//
	// "summaryConstraints": "XXX",
	//
	// "sopInfo": "XXX"
	//
	// }
	MetaData map[string]interface{} `json:"metaData,omitempty" xml:"metaData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// common
	PlayCode *string `json:"playCode,omitempty" xml:"playCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDialogAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDialogAnalysisTaskRequest) GetAnalysisNodes() []*string {
	return s.AnalysisNodes
}

func (s *CreateDialogAnalysisTaskRequest) GetConversationList() []*CreateDialogAnalysisTaskRequestConversationList {
	return s.ConversationList
}

func (s *CreateDialogAnalysisTaskRequest) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *CreateDialogAnalysisTaskRequest) GetPlayCode() *string {
	return s.PlayCode
}

func (s *CreateDialogAnalysisTaskRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDialogAnalysisTaskRequest) SetAnalysisNodes(v []*string) *CreateDialogAnalysisTaskRequest {
	s.AnalysisNodes = v
	return s
}

func (s *CreateDialogAnalysisTaskRequest) SetConversationList(v []*CreateDialogAnalysisTaskRequestConversationList) *CreateDialogAnalysisTaskRequest {
	s.ConversationList = v
	return s
}

func (s *CreateDialogAnalysisTaskRequest) SetMetaData(v map[string]interface{}) *CreateDialogAnalysisTaskRequest {
	s.MetaData = v
	return s
}

func (s *CreateDialogAnalysisTaskRequest) SetPlayCode(v string) *CreateDialogAnalysisTaskRequest {
	s.PlayCode = &v
	return s
}

func (s *CreateDialogAnalysisTaskRequest) SetRequestId(v string) *CreateDialogAnalysisTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateDialogAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDialogAnalysisTaskRequestConversationList struct {
	// This parameter is required.
	DialogueList []*CreateDialogAnalysisTaskRequestConversationListDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
}

func (s CreateDialogAnalysisTaskRequestConversationList) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogAnalysisTaskRequestConversationList) GoString() string {
	return s.String()
}

func (s *CreateDialogAnalysisTaskRequestConversationList) GetDialogueList() []*CreateDialogAnalysisTaskRequestConversationListDialogueList {
	return s.DialogueList
}

func (s *CreateDialogAnalysisTaskRequestConversationList) SetDialogueList(v []*CreateDialogAnalysisTaskRequestConversationListDialogueList) *CreateDialogAnalysisTaskRequestConversationList {
	s.DialogueList = v
	return s
}

func (s *CreateDialogAnalysisTaskRequestConversationList) Validate() error {
	return dara.Validate(s)
}

type CreateDialogAnalysisTaskRequestConversationListDialogueList struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s CreateDialogAnalysisTaskRequestConversationListDialogueList) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogAnalysisTaskRequestConversationListDialogueList) GoString() string {
	return s.String()
}

func (s *CreateDialogAnalysisTaskRequestConversationListDialogueList) GetContent() *string {
	return s.Content
}

func (s *CreateDialogAnalysisTaskRequestConversationListDialogueList) GetRole() *string {
	return s.Role
}

func (s *CreateDialogAnalysisTaskRequestConversationListDialogueList) SetContent(v string) *CreateDialogAnalysisTaskRequestConversationListDialogueList {
	s.Content = &v
	return s
}

func (s *CreateDialogAnalysisTaskRequestConversationListDialogueList) SetRole(v string) *CreateDialogAnalysisTaskRequestConversationListDialogueList {
	s.Role = &v
	return s
}

func (s *CreateDialogAnalysisTaskRequestConversationListDialogueList) Validate() error {
	return dara.Validate(s)
}
