// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSseChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SseChatRequest
	GetAppId() *string
	SetCommand(v string) *SseChatRequest
	GetCommand() *string
	SetSenderId(v string) *SseChatRequest
	GetSenderId() *string
	SetSenderNick(v string) *SseChatRequest
	GetSenderNick() *string
	SetSessionId(v string) *SseChatRequest
	GetSessionId() *string
	SetUtterance(v string) *SseChatRequest
	GetUtterance() *string
	SetVendorParam(v string) *SseChatRequest
	GetVendorParam() *string
	SetWorkspaceId(v string) *SseChatRequest
	GetWorkspaceId() *string
}

type SseChatRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-dDmF3jcdVf
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// TIMEOUT
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// example:
	//
	// uid129312098593
	SenderId   *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// example:
	//
	// 15e04f27-acd7-489d-872f-1d68f7535e33
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	Utterance   *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	VendorParam *string `json:"VendorParam,omitempty" xml:"VendorParam,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-g7jspxljq8k909h3
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SseChatRequest) String() string {
	return dara.Prettify(s)
}

func (s SseChatRequest) GoString() string {
	return s.String()
}

func (s *SseChatRequest) GetAppId() *string {
	return s.AppId
}

func (s *SseChatRequest) GetCommand() *string {
	return s.Command
}

func (s *SseChatRequest) GetSenderId() *string {
	return s.SenderId
}

func (s *SseChatRequest) GetSenderNick() *string {
	return s.SenderNick
}

func (s *SseChatRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SseChatRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *SseChatRequest) GetVendorParam() *string {
	return s.VendorParam
}

func (s *SseChatRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SseChatRequest) SetAppId(v string) *SseChatRequest {
	s.AppId = &v
	return s
}

func (s *SseChatRequest) SetCommand(v string) *SseChatRequest {
	s.Command = &v
	return s
}

func (s *SseChatRequest) SetSenderId(v string) *SseChatRequest {
	s.SenderId = &v
	return s
}

func (s *SseChatRequest) SetSenderNick(v string) *SseChatRequest {
	s.SenderNick = &v
	return s
}

func (s *SseChatRequest) SetSessionId(v string) *SseChatRequest {
	s.SessionId = &v
	return s
}

func (s *SseChatRequest) SetUtterance(v string) *SseChatRequest {
	s.Utterance = &v
	return s
}

func (s *SseChatRequest) SetVendorParam(v string) *SseChatRequest {
	s.VendorParam = &v
	return s
}

func (s *SseChatRequest) SetWorkspaceId(v string) *SseChatRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SseChatRequest) Validate() error {
	return dara.Validate(s)
}
