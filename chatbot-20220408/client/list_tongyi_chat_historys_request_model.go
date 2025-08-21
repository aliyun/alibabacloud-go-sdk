// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTongyiChatHistorysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListTongyiChatHistorysRequest
	GetAgentKey() *string
	SetEndTime(v string) *ListTongyiChatHistorysRequest
	GetEndTime() *string
	SetLimit(v int32) *ListTongyiChatHistorysRequest
	GetLimit() *int32
	SetRobotInstanceId(v string) *ListTongyiChatHistorysRequest
	GetRobotInstanceId() *string
	SetStartTime(v string) *ListTongyiChatHistorysRequest
	GetStartTime() *string
}

type ListTongyiChatHistorysRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-04-01 08:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 30
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
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
	// 2024-04-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListTongyiChatHistorysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTongyiChatHistorysRequest) GoString() string {
	return s.String()
}

func (s *ListTongyiChatHistorysRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListTongyiChatHistorysRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTongyiChatHistorysRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListTongyiChatHistorysRequest) GetRobotInstanceId() *string {
	return s.RobotInstanceId
}

func (s *ListTongyiChatHistorysRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTongyiChatHistorysRequest) SetAgentKey(v string) *ListTongyiChatHistorysRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTongyiChatHistorysRequest) SetEndTime(v string) *ListTongyiChatHistorysRequest {
	s.EndTime = &v
	return s
}

func (s *ListTongyiChatHistorysRequest) SetLimit(v int32) *ListTongyiChatHistorysRequest {
	s.Limit = &v
	return s
}

func (s *ListTongyiChatHistorysRequest) SetRobotInstanceId(v string) *ListTongyiChatHistorysRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *ListTongyiChatHistorysRequest) SetStartTime(v string) *ListTongyiChatHistorysRequest {
	s.StartTime = &v
	return s
}

func (s *ListTongyiChatHistorysRequest) Validate() error {
	return dara.Validate(s)
}
