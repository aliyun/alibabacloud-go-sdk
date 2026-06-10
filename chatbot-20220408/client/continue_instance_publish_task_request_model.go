// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueInstancePublishTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ContinueInstancePublishTaskRequest
	GetAgentKey() *string
	SetId(v int64) *ContinueInstancePublishTaskRequest
	GetId() *int64
	SetInstanceId(v string) *ContinueInstancePublishTaskRequest
	GetInstanceId() *string
}

type ContinueInstancePublishTaskRequest struct {
	// The key for the business space. If you omit this parameter, the default business space is used. You can find this key on the Business Management page of your Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the chatbot.
	//
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ContinueInstancePublishTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ContinueInstancePublishTaskRequest) GoString() string {
	return s.String()
}

func (s *ContinueInstancePublishTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ContinueInstancePublishTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *ContinueInstancePublishTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ContinueInstancePublishTaskRequest) SetAgentKey(v string) *ContinueInstancePublishTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *ContinueInstancePublishTaskRequest) SetId(v int64) *ContinueInstancePublishTaskRequest {
	s.Id = &v
	return s
}

func (s *ContinueInstancePublishTaskRequest) SetInstanceId(v string) *ContinueInstancePublishTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ContinueInstancePublishTaskRequest) Validate() error {
	return dara.Validate(s)
}
