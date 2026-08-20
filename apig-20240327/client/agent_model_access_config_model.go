// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentModelAccessConfig interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerIds(v []*string) *AgentModelAccessConfig
	GetConsumerIds() []*string
	SetModelApiId(v string) *AgentModelAccessConfig
	GetModelApiId() *string
}

type AgentModelAccessConfig struct {
	// The list of consumer IDs that represent the Agent to access the Model API. The Model API ID and consumer ID together identify the Agent identity, and the configuration takes effect for all current and future routes of the Model API. Specify at least one consumer. The consumer must be enabled and must have direct Consumer authorization for the Model API in the default environment of the target gateway, with the authorization publish status being Success. Different Agents cannot bind the same consumer to the same Model API. ConsumerGroup is not supported.
	//
	// This parameter is required.
	ConsumerIds []*string `json:"consumerIds,omitempty" xml:"consumerIds,omitempty" type:"Repeated"`
	// The ID of the Model API to associate. The Model API must belong to the specified gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// model-api-1
	ModelApiId *string `json:"modelApiId,omitempty" xml:"modelApiId,omitempty"`
}

func (s AgentModelAccessConfig) String() string {
	return dara.Prettify(s)
}

func (s AgentModelAccessConfig) GoString() string {
	return s.String()
}

func (s *AgentModelAccessConfig) GetConsumerIds() []*string {
	return s.ConsumerIds
}

func (s *AgentModelAccessConfig) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *AgentModelAccessConfig) SetConsumerIds(v []*string) *AgentModelAccessConfig {
	s.ConsumerIds = v
	return s
}

func (s *AgentModelAccessConfig) SetModelApiId(v string) *AgentModelAccessConfig {
	s.ModelApiId = &v
	return s
}

func (s *AgentModelAccessConfig) Validate() error {
	return dara.Validate(s)
}
