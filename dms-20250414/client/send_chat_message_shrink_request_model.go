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
	// **[Optimized]*	- This field is now automatically retrieved in the background. You do not need to specify this field.
	//
	// example:
	//
	// agent_***
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// **[Optimized]*	- This field is now automatically retrieved in the background. You do not need to specify this field when you call the API operation.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The data source information. This parameter can be left empty. You can specify only one data source for this parameter. We recommend that you use the DataSources parameter instead.
	//
	// example:
	//
	// null
	DataSourceShrink *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The detailed data source information. This parameter can be left empty.
	DataSourcesShrink *string `json:"DataSources,omitempty" xml:"DataSources,omitempty"`
	// The content of the message sent to the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// what can you do?
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The message type. Default value: primary.
	//
	// - In normal interactions with the agent, the message type is primary.
	//
	// - When the message responds to a human-in-the-loop question from the agent, the type must be additional.
	//
	// - When the message triggers report generation, the type must be report.
	//
	// - When the message cancels the current session, the type must be cancel.
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
	// The specific question asked by the agent to the user through human-in-the-loop. This field is required when the message type is additional.
	//
	// example:
	//
	// Provide the calculation criteria for GMV
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// The current quoted content. This parameter is typically used during interactions with the agent.
	//
	// example:
	//
	// {"version":"v0"}
	QuotedMessage *string `json:"QuotedMessage,omitempty" xml:"QuotedMessage,omitempty"`
	// **Important**
	//
	// If this message is a reply to an agent message, such as an agent clarifying a question through ASK_HUMAN, you must accurately pass the checkpoint sequence number carried in the agent message to reply_to. If this message is not a targeted reply, such as requesting the agent to perform further in-depth analysis after the analysis is complete, you can leave reply_to empty or set it to 0.
	//
	// This field affects how the agent decides to process the message. Passing an incorrect value may lead to suboptimal analysis results.
	//
	// example:
	//
	// 0
	ReplyTo *string `json:"ReplyTo,omitempty" xml:"ReplyTo,omitempty"`
	// The special configurations for the current session. For the same session, only the configurations passed during the first SendMessage call take effect.
	//
	// if can be null:
	// true
	SessionConfigShrink *string `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty"`
	// The session ID. This field is optional and is used for multi-turn sessions.
	//
	// - You can start a session without specifying this field. The response will contain the session ID of the current session.
	//
	// - You can also manually create a session ID by calling the CreateDataAgentSession operation and include the ID when you initiate the session.
	//
	// - If you need multi-turn conversations, such as follow-up questions or confirming execution plans, include the session ID returned by the previous SendChatMessage operation in this operation.
	//
	// example:
	//
	// sess_***
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The configuration items that affect only the current task.
	TaskConfigShrink *string `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty"`
	// The OSS bucket of the user. If you leave this parameter empty, the analysis process is securely stored in the built-in storage.
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
