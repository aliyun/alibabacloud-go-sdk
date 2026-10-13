// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBotSessionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetBotSessionDataRequest
	GetAgentKey() *string
	SetEndTime(v string) *GetBotSessionDataRequest
	GetEndTime() *string
	SetRobotInstanceId(v string) *GetBotSessionDataRequest
	GetRobotInstanceId() *string
	SetStartTime(v string) *GetBotSessionDataRequest
	GetStartTime() *string
}

type GetBotSessionDataRequest struct {
	// The business space key. If this parameter is not specified, the default business space is accessed. You can obtain the key value on the business management page of the Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The end time of the query. The time format is yyyyMMdd, such as 20240605.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20240605
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the chatbot instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-7QuUfaqMQe
	RobotInstanceId *string `json:"RobotInstanceId,omitempty" xml:"RobotInstanceId,omitempty"`
	// The start time of the query. The time format is yyyyMMdd, such as 20240505.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20240505
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetBotSessionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBotSessionDataRequest) GoString() string {
	return s.String()
}

func (s *GetBotSessionDataRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetBotSessionDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetBotSessionDataRequest) GetRobotInstanceId() *string {
	return s.RobotInstanceId
}

func (s *GetBotSessionDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetBotSessionDataRequest) SetAgentKey(v string) *GetBotSessionDataRequest {
	s.AgentKey = &v
	return s
}

func (s *GetBotSessionDataRequest) SetEndTime(v string) *GetBotSessionDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetBotSessionDataRequest) SetRobotInstanceId(v string) *GetBotSessionDataRequest {
	s.RobotInstanceId = &v
	return s
}

func (s *GetBotSessionDataRequest) SetStartTime(v string) *GetBotSessionDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetBotSessionDataRequest) Validate() error {
	return dara.Validate(s)
}
