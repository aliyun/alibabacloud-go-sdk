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
	// This parameter is required.
	ConsumerIds []*string `json:"consumerIds,omitempty" xml:"consumerIds,omitempty" type:"Repeated"`
	// This parameter is required.
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
