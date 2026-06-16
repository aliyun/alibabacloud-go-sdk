// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMem0MemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *AddMem0MemoriesRequest
	GetAgentSpace() *string
	SetBody(v map[string]interface{}) *AddMem0MemoriesRequest
	GetBody() map[string]interface{}
}

type AddMem0MemoriesRequest struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// {"messages":[{"role":"user","content":"我喜欢喝拿铁"}],"user_id":"alice","metadata":{"channel":"app"}}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMem0MemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMem0MemoriesRequest) GoString() string {
	return s.String()
}

func (s *AddMem0MemoriesRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *AddMem0MemoriesRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *AddMem0MemoriesRequest) SetAgentSpace(v string) *AddMem0MemoriesRequest {
	s.AgentSpace = &v
	return s
}

func (s *AddMem0MemoriesRequest) SetBody(v map[string]interface{}) *AddMem0MemoriesRequest {
	s.Body = v
	return s
}

func (s *AddMem0MemoriesRequest) Validate() error {
	return dara.Validate(s)
}
