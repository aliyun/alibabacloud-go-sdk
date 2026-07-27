// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *SendChatMessageShrinkRequest
	GetAgentId() *string
	SetDMSUnit(v string) *SendChatMessageShrinkRequest
	GetDMSUnit() *string
	SetDataSourceShrink(v string) *SendChatMessageShrinkRequest
	GetDataSourceShrink() *string
	SetDataSourcesShrink(v string) *SendChatMessageShrinkRequest
	GetDataSourcesShrink() *string
	SetMessage(v string) *SendChatMessageShrinkRequest
	GetMessage() *string
	SetMessageType(v string) *SendChatMessageShrinkRequest
	GetMessageType() *string
	SetParentSessionId(v string) *SendChatMessageShrinkRequest
	GetParentSessionId() *string
	SetQuestion(v string) *SendChatMessageShrinkRequest
	GetQuestion() *string
	SetQuotedMessage(v string) *SendChatMessageShrinkRequest
	GetQuotedMessage() *string
	SetReplyTo(v string) *SendChatMessageShrinkRequest
	GetReplyTo() *string
	SetSessionConfigShrink(v string) *SendChatMessageShrinkRequest
	GetSessionConfigShrink() *string
	SetSessionId(v string) *SendChatMessageShrinkRequest
	GetSessionId() *string
	SetTaskConfigShrink(v string) *SendChatMessageShrinkRequest
	GetTaskConfigShrink() *string
	SetUserOssBucket(v string) *SendChatMessageShrinkRequest
	GetUserOssBucket() *string
	SetWorkspaceId(v string) *SendChatMessageShrinkRequest
	GetWorkspaceId() *string
}

type SendChatMessageShrinkRequest struct {
	// The agent ID. This parameter is required. You can obtain the current AgentId from the response of the CreateAgentSession operation. Agent resources have a lifecycle, so the AgentId you need to specify may change with each request.
	//
	// example:
	//
	// agent_***
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The Data Management unit you are currently in. If you select an analytics database, this information is used to correctly connect to your Data Management instance through Data Management. You can go to the DAS console to view your current Data Management unit. If you are a user of Alibaba Cloud China Website (www.aliyun.com), set this parameter to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The data source information. This parameter can be left empty. Only one data source can be specified for this parameter. Use the DataSources parameter instead.
	//
	// example:
	//
	// null
	DataSourceShrink *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The detailed data source information. This parameter can be left empty.
	DataSourcesShrink *string `json:"DataSources,omitempty" xml:"DataSources,omitempty"`
	// The message content to send to the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// what can you do?
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The message type. Default value: `[primary]`.
	//
	// example:
	//
	// primary
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The parent session ID.
	//
	// example:
	//
	// 20qrliuoo7p2vlsfg*****
	ParentSessionId *string `json:"ParentSessionId,omitempty" xml:"ParentSessionId,omitempty"`
	// The specific question that the agent asks the user through Human-in-Loop. This parameter is required when the message type is `additional`.
	//
	// example:
	//
	// Please provide the criteria for calculating GMV.
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// The quoted content. This parameter is typically used when interacting with the agent.
	//
	// example:
	//
	// {"version":"v0"}
	QuotedMessage *string `json:"QuotedMessage,omitempty" xml:"QuotedMessage,omitempty"`
	// **Important**
	//
	// example:
	//
	// 0
	ReplyTo *string `json:"ReplyTo,omitempty" xml:"ReplyTo,omitempty"`
	// The special configuration for the current session. For the same session, only the configuration specified in the first SendMessage call takes effect.
	//
	// if can be null:
	// true
	SessionConfigShrink *string `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty"`
	// The session ID. This parameter is required. You can obtain the SessionId by calling the CreateAgentSession operation.
	//
	// example:
	//
	// sess_***
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The configuration items that affect only the current task.
	TaskConfigShrink *string `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty"`
	// The OSS bucket of the user. If this parameter is not specified, the analysis data is securely stored in the built-in storage.
	//
	// example:
	//
	// my-bucket
	UserOssBucket *string `json:"UserOssBucket,omitempty" xml:"UserOssBucket,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// c1p71ne***baexrt3o
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SendChatMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendChatMessageShrinkRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *SendChatMessageShrinkRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *SendChatMessageShrinkRequest) GetDataSourceShrink() *string {
	return s.DataSourceShrink
}

func (s *SendChatMessageShrinkRequest) GetDataSourcesShrink() *string {
	return s.DataSourcesShrink
}

func (s *SendChatMessageShrinkRequest) GetMessage() *string {
	return s.Message
}

func (s *SendChatMessageShrinkRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *SendChatMessageShrinkRequest) GetParentSessionId() *string {
	return s.ParentSessionId
}

func (s *SendChatMessageShrinkRequest) GetQuestion() *string {
	return s.Question
}

func (s *SendChatMessageShrinkRequest) GetQuotedMessage() *string {
	return s.QuotedMessage
}

func (s *SendChatMessageShrinkRequest) GetReplyTo() *string {
	return s.ReplyTo
}

func (s *SendChatMessageShrinkRequest) GetSessionConfigShrink() *string {
	return s.SessionConfigShrink
}

func (s *SendChatMessageShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendChatMessageShrinkRequest) GetTaskConfigShrink() *string {
	return s.TaskConfigShrink
}

func (s *SendChatMessageShrinkRequest) GetUserOssBucket() *string {
	return s.UserOssBucket
}

func (s *SendChatMessageShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SendChatMessageShrinkRequest) SetAgentId(v string) *SendChatMessageShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetDMSUnit(v string) *SendChatMessageShrinkRequest {
	s.DMSUnit = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetDataSourceShrink(v string) *SendChatMessageShrinkRequest {
	s.DataSourceShrink = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetDataSourcesShrink(v string) *SendChatMessageShrinkRequest {
	s.DataSourcesShrink = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetMessage(v string) *SendChatMessageShrinkRequest {
	s.Message = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetMessageType(v string) *SendChatMessageShrinkRequest {
	s.MessageType = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetParentSessionId(v string) *SendChatMessageShrinkRequest {
	s.ParentSessionId = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetQuestion(v string) *SendChatMessageShrinkRequest {
	s.Question = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetQuotedMessage(v string) *SendChatMessageShrinkRequest {
	s.QuotedMessage = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetReplyTo(v string) *SendChatMessageShrinkRequest {
	s.ReplyTo = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetSessionConfigShrink(v string) *SendChatMessageShrinkRequest {
	s.SessionConfigShrink = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetSessionId(v string) *SendChatMessageShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetTaskConfigShrink(v string) *SendChatMessageShrinkRequest {
	s.TaskConfigShrink = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetUserOssBucket(v string) *SendChatMessageShrinkRequest {
	s.UserOssBucket = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetWorkspaceId(v string) *SendChatMessageShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SendChatMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
