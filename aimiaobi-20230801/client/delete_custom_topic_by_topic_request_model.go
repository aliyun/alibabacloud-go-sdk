// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTopicByTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteCustomTopicByTopicRequest
	GetAgentKey() *string
	SetTopic(v string) *DeleteCustomTopicByTopicRequest
	GetTopic() *string
}

type DeleteCustomTopicByTopicRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 话题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DeleteCustomTopicByTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTopicByTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicByTopicRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteCustomTopicByTopicRequest) GetTopic() *string {
	return s.Topic
}

func (s *DeleteCustomTopicByTopicRequest) SetAgentKey(v string) *DeleteCustomTopicByTopicRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteCustomTopicByTopicRequest) SetTopic(v string) *DeleteCustomTopicByTopicRequest {
	s.Topic = &v
	return s
}

func (s *DeleteCustomTopicByTopicRequest) Validate() error {
	return dara.Validate(s)
}
