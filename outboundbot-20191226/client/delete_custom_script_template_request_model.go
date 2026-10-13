// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomScriptTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProfileId(v string) *DeleteCustomScriptTemplateRequest
	GetAgentProfileId() *string
	SetInstanceId(v string) *DeleteCustomScriptTemplateRequest
	GetInstanceId() *string
}

type DeleteCustomScriptTemplateRequest struct {
	// The ID of the custom scenario template.
	//
	// example:
	//
	// d31794e2a51f47d2901b4094d88311d7
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 0ec0c897-b92c-40e4-9ad7-e6e4f5ce13bb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteCustomScriptTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomScriptTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomScriptTemplateRequest) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *DeleteCustomScriptTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCustomScriptTemplateRequest) SetAgentProfileId(v string) *DeleteCustomScriptTemplateRequest {
	s.AgentProfileId = &v
	return s
}

func (s *DeleteCustomScriptTemplateRequest) SetInstanceId(v string) *DeleteCustomScriptTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCustomScriptTemplateRequest) Validate() error {
	return dara.Validate(s)
}
