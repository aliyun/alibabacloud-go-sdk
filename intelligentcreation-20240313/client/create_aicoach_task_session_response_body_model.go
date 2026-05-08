// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICoachTaskSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelToken(v string) *CreateAICoachTaskSessionResponseBody
	GetChannelToken() *string
	SetRequestId(v string) *CreateAICoachTaskSessionResponseBody
	GetRequestId() *string
	SetScriptInfo(v *CreateAICoachTaskSessionResponseBodyScriptInfo) *CreateAICoachTaskSessionResponseBody
	GetScriptInfo() *CreateAICoachTaskSessionResponseBodyScriptInfo
	SetSessionId(v string) *CreateAICoachTaskSessionResponseBody
	GetSessionId() *string
	SetSessionStatus(v int64) *CreateAICoachTaskSessionResponseBody
	GetSessionStatus() *int64
	SetToken(v string) *CreateAICoachTaskSessionResponseBody
	GetToken() *string
	SetWebSocketUrl(v string) *CreateAICoachTaskSessionResponseBody
	GetWebSocketUrl() *string
}

type CreateAICoachTaskSessionResponseBody struct {
	// rtctoken
	//
	// example:
	//
	// 11
	ChannelToken *string `json:"channelToken,omitempty" xml:"channelToken,omitempty"`
	// example:
	//
	// 4830493A-728F-5F19-BBCC-1443292E9C49
	RequestId  *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScriptInfo *CreateAICoachTaskSessionResponseBodyScriptInfo `json:"scriptInfo,omitempty" xml:"scriptInfo,omitempty" type:"Struct"`
	// example:
	//
	// 111
	SessionId     *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	SessionStatus *int64  `json:"sessionStatus,omitempty" xml:"sessionStatus,omitempty"`
	// Token
	//
	// example:
	//
	// 11
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// 11
	WebSocketUrl *string `json:"webSocketUrl,omitempty" xml:"webSocketUrl,omitempty"`
}

func (s CreateAICoachTaskSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskSessionResponseBody) GetChannelToken() *string {
	return s.ChannelToken
}

func (s *CreateAICoachTaskSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAICoachTaskSessionResponseBody) GetScriptInfo() *CreateAICoachTaskSessionResponseBodyScriptInfo {
	return s.ScriptInfo
}

func (s *CreateAICoachTaskSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateAICoachTaskSessionResponseBody) GetSessionStatus() *int64 {
	return s.SessionStatus
}

func (s *CreateAICoachTaskSessionResponseBody) GetToken() *string {
	return s.Token
}

func (s *CreateAICoachTaskSessionResponseBody) GetWebSocketUrl() *string {
	return s.WebSocketUrl
}

func (s *CreateAICoachTaskSessionResponseBody) SetChannelToken(v string) *CreateAICoachTaskSessionResponseBody {
	s.ChannelToken = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetRequestId(v string) *CreateAICoachTaskSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetScriptInfo(v *CreateAICoachTaskSessionResponseBodyScriptInfo) *CreateAICoachTaskSessionResponseBody {
	s.ScriptInfo = v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetSessionId(v string) *CreateAICoachTaskSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetSessionStatus(v int64) *CreateAICoachTaskSessionResponseBody {
	s.SessionStatus = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetToken(v string) *CreateAICoachTaskSessionResponseBody {
	s.Token = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetWebSocketUrl(v string) *CreateAICoachTaskSessionResponseBody {
	s.WebSocketUrl = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) Validate() error {
	if s.ScriptInfo != nil {
		if err := s.ScriptInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAICoachTaskSessionResponseBodyScriptInfo struct {
	AgentIconUrl     *string   `json:"agentIconUrl,omitempty" xml:"agentIconUrl,omitempty"`
	CharacterName    *string   `json:"characterName,omitempty" xml:"characterName,omitempty"`
	DialogueTextFlag *bool     `json:"dialogueTextFlag,omitempty" xml:"dialogueTextFlag,omitempty"`
	DialogueTipFlag  *bool     `json:"dialogueTipFlag,omitempty" xml:"dialogueTipFlag,omitempty"`
	Initiator        *string   `json:"initiator,omitempty" xml:"initiator,omitempty"`
	InputTypeList    []*string `json:"inputTypeList,omitempty" xml:"inputTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 11
	MaxDuration *int64 `json:"maxDuration,omitempty" xml:"maxDuration,omitempty"`
	// example:
	//
	// test
	ScriptDesc            *string `json:"scriptDesc,omitempty" xml:"scriptDesc,omitempty"`
	ScriptName            *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	ScriptRecordId        *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	ScriptType            *int64  `json:"scriptType,omitempty" xml:"scriptType,omitempty"`
	SparringTipContent    *string `json:"sparringTipContent,omitempty" xml:"sparringTipContent,omitempty"`
	SparringTipTitle      *string `json:"sparringTipTitle,omitempty" xml:"sparringTipTitle,omitempty"`
	StudentThinkTimeFlag  *bool   `json:"studentThinkTimeFlag,omitempty" xml:"studentThinkTimeFlag,omitempty"`
	StudentThinkTimeLimit *int64  `json:"studentThinkTimeLimit,omitempty" xml:"studentThinkTimeLimit,omitempty"`
}

func (s CreateAICoachTaskSessionResponseBodyScriptInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskSessionResponseBodyScriptInfo) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetAgentIconUrl() *string {
	return s.AgentIconUrl
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetCharacterName() *string {
	return s.CharacterName
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetDialogueTextFlag() *bool {
	return s.DialogueTextFlag
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetDialogueTipFlag() *bool {
	return s.DialogueTipFlag
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetInitiator() *string {
	return s.Initiator
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetInputTypeList() []*string {
	return s.InputTypeList
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetMaxDuration() *int64 {
	return s.MaxDuration
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetScriptDesc() *string {
	return s.ScriptDesc
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetScriptName() *string {
	return s.ScriptName
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetScriptType() *int64 {
	return s.ScriptType
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetSparringTipContent() *string {
	return s.SparringTipContent
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetSparringTipTitle() *string {
	return s.SparringTipTitle
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetStudentThinkTimeFlag() *bool {
	return s.StudentThinkTimeFlag
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) GetStudentThinkTimeLimit() *int64 {
	return s.StudentThinkTimeLimit
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetAgentIconUrl(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.AgentIconUrl = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetCharacterName(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.CharacterName = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetDialogueTextFlag(v bool) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.DialogueTextFlag = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetDialogueTipFlag(v bool) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.DialogueTipFlag = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetInitiator(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.Initiator = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetInputTypeList(v []*string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.InputTypeList = v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetMaxDuration(v int64) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.MaxDuration = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetScriptDesc(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.ScriptDesc = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetScriptName(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.ScriptName = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetScriptRecordId(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.ScriptRecordId = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetScriptType(v int64) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.ScriptType = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetSparringTipContent(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.SparringTipContent = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetSparringTipTitle(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.SparringTipTitle = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetStudentThinkTimeFlag(v bool) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.StudentThinkTimeFlag = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetStudentThinkTimeLimit(v int64) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.StudentThinkTimeLimit = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) Validate() error {
	return dara.Validate(s)
}
