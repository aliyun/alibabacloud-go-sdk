// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTongyiConversationLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListTongyiConversationLogsRequest
	GetAgentKey() *string
	SetRobotInstanceId(v string) *ListTongyiConversationLogsRequest
	GetRobotInstanceId() *string
	SetSessionId(v string) *ListTongyiConversationLogsRequest
	GetSessionId() *string
}

type ListTongyiConversationLogsRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-7QuUfaqMQe
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7c3cec23cc8940bc9db4a318c8f4f0aa
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ListTongyiConversationLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTongyiConversationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListTongyiConversationLogsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListTongyiConversationLogsRequest) GetRobotInstanceId() *string {
	return s.RobotInstanceId
}

func (s *ListTongyiConversationLogsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListTongyiConversationLogsRequest) SetAgentKey(v string) *ListTongyiConversationLogsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTongyiConversationLogsRequest) SetRobotInstanceId(v string) *ListTongyiConversationLogsRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListTongyiConversationLogsRequest) SetSessionId(v string) *ListTongyiConversationLogsRequest {
	s.SessionId = &v
	return s
}

func (s *ListTongyiConversationLogsRequest) Validate() error {
	return dara.Validate(s)
}
