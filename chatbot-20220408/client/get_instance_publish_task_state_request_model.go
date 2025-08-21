// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancePublishTaskStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetInstancePublishTaskStateRequest
	GetAgentKey() *string
	SetId(v int64) *GetInstancePublishTaskStateRequest
	GetId() *int64
	SetInstanceId(v string) *GetInstancePublishTaskStateRequest
	GetInstanceId() *string
}

type GetInstancePublishTaskStateRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstancePublishTaskStateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstancePublishTaskStateRequest) GoString() string {
	return s.String()
}

func (s *GetInstancePublishTaskStateRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetInstancePublishTaskStateRequest) GetId() *int64 {
	return s.Id
}

func (s *GetInstancePublishTaskStateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstancePublishTaskStateRequest) SetAgentKey(v string) *GetInstancePublishTaskStateRequest {
	s.AgentKey = &v
	return s
}

func (s *GetInstancePublishTaskStateRequest) SetId(v int64) *GetInstancePublishTaskStateRequest {
	s.Id = &v
	return s
}

func (s *GetInstancePublishTaskStateRequest) SetInstanceId(v string) *GetInstancePublishTaskStateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstancePublishTaskStateRequest) Validate() error {
	return dara.Validate(s)
}
