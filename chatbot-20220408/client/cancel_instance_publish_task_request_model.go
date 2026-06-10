// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelInstancePublishTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CancelInstancePublishTaskRequest
	GetAgentKey() *string
	SetId(v int64) *CancelInstancePublishTaskRequest
	GetId() *int64
	SetInstanceId(v string) *CancelInstancePublishTaskRequest
	GetInstanceId() *string
}

type CancelInstancePublishTaskRequest struct {
	// The key of the business space. You can obtain the key from the Business Management page of your Alibaba Cloud account. If you omit this parameter, the default business space is used.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique ID of the chatbot.
	//
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CancelInstancePublishTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelInstancePublishTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelInstancePublishTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CancelInstancePublishTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *CancelInstancePublishTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelInstancePublishTaskRequest) SetAgentKey(v string) *CancelInstancePublishTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CancelInstancePublishTaskRequest) SetId(v int64) *CancelInstancePublishTaskRequest {
	s.Id = &v
	return s
}

func (s *CancelInstancePublishTaskRequest) SetInstanceId(v string) *CancelInstancePublishTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelInstancePublishTaskRequest) Validate() error {
	return dara.Validate(s)
}
