// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstancePublishTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateInstancePublishTaskRequest
	GetAgentKey() *string
	SetInstanceId(v string) *CreateInstancePublishTaskRequest
	GetInstanceId() *string
}

type CreateInstancePublishTaskRequest struct {
	// The key for the business space. If you omit this parameter, the default business space is used. You can obtain the key from the Business Management page of your primary account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The unique instance ID of the robot.
	//
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateInstancePublishTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstancePublishTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateInstancePublishTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateInstancePublishTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstancePublishTaskRequest) SetAgentKey(v string) *CreateInstancePublishTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateInstancePublishTaskRequest) SetInstanceId(v string) *CreateInstancePublishTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstancePublishTaskRequest) Validate() error {
	return dara.Validate(s)
}
